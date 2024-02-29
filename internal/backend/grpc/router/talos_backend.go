// Copyright (c) 2024 Sidero Labs, Inc.
//
// Use of this software is governed by the Business Source License
// included in the LICENSE file.

package router

import (
	"context"

	"github.com/blang/semver"
	"github.com/siderolabs/gen/xslices"
	"github.com/siderolabs/go-api-signature/pkg/message"
	"github.com/siderolabs/talos/pkg/machinery/api/machine"
	"github.com/siderolabs/talos/pkg/machinery/constants"
	talosrole "github.com/siderolabs/talos/pkg/machinery/role"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/siderolabs/omni/internal/backend/dns"
	"github.com/siderolabs/omni/internal/pkg/auth"
	"github.com/siderolabs/omni/internal/pkg/auth/role"
	"github.com/siderolabs/omni/internal/pkg/grpcutil"
)

// operatorMethodSet is the set of methods that are allowed to be called by the minimum role of os:operator.
var operatorMethodSet = xslices.ToSet([]string{
	grpcutil.MustFullMethodName(&machine.MachineService_ServiceDesc, "EtcdAlarmList"),
	grpcutil.MustFullMethodName(&machine.MachineService_ServiceDesc, "EtcdAlarmDisarm"),
	grpcutil.MustFullMethodName(&machine.MachineService_ServiceDesc, "EtcdDefragment"),
	grpcutil.MustFullMethodName(&machine.MachineService_ServiceDesc, "EtcdStatus"),
	grpcutil.MustFullMethodName(&machine.MachineService_ServiceDesc, "PacketCapture"),
	grpcutil.MustFullMethodName(&machine.MachineService_ServiceDesc, "Reboot"),
	grpcutil.MustFullMethodName(&machine.MachineService_ServiceDesc, "Restart"),
	grpcutil.MustFullMethodName(&machine.MachineService_ServiceDesc, "ServiceRestart"),
	grpcutil.MustFullMethodName(&machine.MachineService_ServiceDesc, "ServiceStart"),
	grpcutil.MustFullMethodName(&machine.MachineService_ServiceDesc, "ServiceStop"),
	grpcutil.MustFullMethodName(&machine.MachineService_ServiceDesc, "Shutdown"),
})

// TalosBackend implements a backend (proxying one2one to a Talos node).
type TalosBackend struct {
	conn         *grpc.ClientConn
	nodeResolver NodeResolver
	verifier     grpc.UnaryServerInterceptor
	name         string
	authEnabled  bool
}

// NewTalosBackend builds new Talos API backend.
func NewTalosBackend(name string, nodeResolver NodeResolver, conn *grpc.ClientConn, authEnabled bool, verifier grpc.UnaryServerInterceptor) *TalosBackend {
	backend := &TalosBackend{
		name:         name,
		nodeResolver: nodeResolver,
		conn:         conn,
		authEnabled:  authEnabled,
		verifier:     verifier,
	}

	return backend
}

func (l *TalosBackend) String() string {
	return l.name
}

// GetConnection returns a grpc connection to the backend.
func (l *TalosBackend) GetConnection(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
	md, _ := metadata.FromIncomingContext(ctx)

	// we can't use regular gRPC server interceptors here, as proxy interface is a bit different

	// prepare context values for the verifier
	ctx = context.WithValue(ctx, auth.EnabledAuthContextKey{}, l.authEnabled)
	msg := message.NewGRPC(md, fullMethodName)
	ctx = context.WithValue(ctx, auth.GRPCMessageContextKey{}, msg)

	grpcutil.SetShouldLog(ctx, "talos-backend")
	grpcutil.AddLogPair(ctx, "cluster", l.name)

	// perform authentication, result of the authentication should be written to ctx
	_, err := l.verifier(ctx, nil, nil,
		func(innerCtx context.Context, _ any) (any, error) {
			// save enhanced context
			ctx = innerCtx

			return nil, nil //nolint:nilnil
		},
	)
	if err != nil {
		// authentication failed
		return ctx, nil, err
	}

	hasModifyAccess := false

	_, modifyErr := auth.Check(ctx, auth.WithRole(role.Operator))
	if modifyErr == nil {
		hasModifyAccess = true
	}

	if !hasModifyAccess {
		// at least read access is required
		if _, err := auth.CheckGRPC(ctx, auth.WithRole(role.Reader)); err != nil {
			return ctx, nil, err
		}
	}

	// overwrite the node headers with the resolved ones
	resolved := resolveNodes(l.nodeResolver, md)
	if resolved.node.Address != "" {
		md = md.Copy()

		setHeaderData(ctx, md, "node", resolved.node.Address)
	} else if len(resolved.nodes) > 0 {
		md = md.Copy()

		setHeaderData(ctx, md, "nodes", xslices.Map(resolved.nodes, func(info dns.Info) string {
			return info.Address
		})...)
	}

	l.setRoleHeaders(ctx, md, fullMethodName, resolved, hasModifyAccess)

	outCtx := metadata.NewOutgoingContext(ctx, md)

	return outCtx, l.conn, nil
}

func (l *TalosBackend) setRoleHeaders(ctx context.Context, md metadata.MD, fullMethodName string, resolvedInfo resolvedNodeInfo, hasModifyAccess bool) {
	if !hasModifyAccess {
		setHeaderData(ctx, md, constants.APIAuthzRoleMetadataKey, talosrole.MakeSet(talosrole.Reader).Strings()...)

		return
	}

	minTalosVersion := l.minTalosVersion(resolvedInfo)

	// min Talos version is >= 1.4.0, we can use Operator role
	if minTalosVersion != nil && minTalosVersion.GTE(semver.MustParse("1.4.0")) {
		setHeaderData(ctx, md, constants.APIAuthzRoleMetadataKey, talosrole.MakeSet(talosrole.Operator).Strings()...)

		return
	}

	// min Talos version is unknown or < 1.4.0, fallback to backwards-compatibility logic
	if _, ok := operatorMethodSet[fullMethodName]; ok {
		setHeaderData(ctx, md, constants.APIAuthzRoleMetadataKey, talosrole.MakeSet(talosrole.Admin).Strings()...)
	} else {
		setHeaderData(ctx, md, constants.APIAuthzRoleMetadataKey, talosrole.MakeSet(talosrole.Reader).Strings()...)
	}
}

func (l *TalosBackend) minTalosVersion(info resolvedNodeInfo) *semver.Version {
	ver := takePtr(semver.ParseTolerant(info.node.TalosVersion))

	for _, node := range info.nodes {
		nodeVer := takePtr(semver.ParseTolerant(node.TalosVersion))
		if nodeVer != nil && (ver == nil || nodeVer.LT(*ver)) {
			ver = nodeVer
		}
	}

	return ver
}

func takePtr[T any](v T, err error) *T {
	if err != nil {
		return nil
	}

	return &v
}

// AppendInfo is called to enhance response from the backend with additional data.
func (l *TalosBackend) AppendInfo(_ bool, resp []byte) ([]byte, error) {
	return resp, nil
}

// BuildError is called to convert error from upstream into response field.
func (l *TalosBackend) BuildError(bool, error) ([]byte, error) {
	return nil, nil
}

func setHeaderData(ctx context.Context, md metadata.MD, k string, v ...string) {
	if len(v) == 0 {
		return
	}

	md.Set(k, v...)

	if len(v) == 1 {
		grpcutil.AddLogPair(ctx, k, v[0])
	} else {
		grpcutil.AddLogPair(ctx, k, v)
	}
}
