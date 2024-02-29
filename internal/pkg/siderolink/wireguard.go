// Copyright (c) 2024 Sidero Labs, Inc.
//
// Use of this software is governed by the Business Source License
// included in the LICENSE file.

package siderolink

import (
	"context"
	"fmt"
	"net/netip"

	"github.com/siderolabs/go-pointer"
	"github.com/siderolabs/siderolink/pkg/wireguard"
	"go.uber.org/zap"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"

	"github.com/siderolabs/omni/client/api/omni/specs"
)

// WireguardHandler abstraction around peer handler and wgDevice.
type WireguardHandler interface {
	SetupDevice(netip.Prefix, wgtypes.Key, string, *zap.Logger) error
	Shutdown() error
	Run(context.Context, *zap.Logger) error
	PeerEvent(context.Context, *specs.SiderolinkSpec, bool) error
	Peers() ([]wgtypes.Peer, error)
}

// PhysicalWireguardHandler implements WireguardHandler interface.
type PhysicalWireguardHandler struct {
	wgDevice *wireguard.Device
	events   chan wireguard.PeerEvent
}

// DefaultWireguardHandler is a default wireguard handler to be used with the siderolink manager.
var DefaultWireguardHandler = &PhysicalWireguardHandler{events: make(chan wireguard.PeerEvent)}

// PeerEvent implements WireguardHandler.
func (handler *PhysicalWireguardHandler) PeerEvent(ctx context.Context, spec *specs.SiderolinkSpec, deleted bool) error {
	address, err := netip.ParsePrefix(spec.NodeSubnet)
	if err != nil {
		return err
	}

	pubKey, err := wgtypes.ParseKey(spec.NodePublicKey)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
	case handler.events <- wireguard.PeerEvent{
		PubKey:                      pubKey,
		Address:                     address.Addr(),
		Remove:                      deleted,
		Endpoint:                    spec.LastEndpoint,
		PersistentKeepAliveInterval: pointer.To(wireguard.RecommendedPersistentKeepAliveInterval),
	}:
	}

	return nil
}

// EventCh implements the wireguard.PeerSource interface.
func (handler *PhysicalWireguardHandler) EventCh() <-chan wireguard.PeerEvent {
	return handler.events
}

// SetupDevice implements WireguardHandler.
func (handler *PhysicalWireguardHandler) SetupDevice(serverAddr netip.Prefix, key wgtypes.Key, endpoint string, logger *zap.Logger) error {
	wireguardEndpoint, err := netip.ParseAddrPort(endpoint)
	if err != nil {
		return fmt.Errorf("invalid Wireguard endpoint: %w", err)
	}

	handler.wgDevice, err = wireguard.NewDevice(serverAddr, key, wireguardEndpoint.Port(), false, logger)
	if err != nil {
		return fmt.Errorf("error initializing wgDevice: %w", err)
	}

	return nil
}

// Shutdown implements WireguardHandler.
func (handler *PhysicalWireguardHandler) Shutdown() error {
	return handler.wgDevice.Close()
}

// Run implements WireguardHandler.
func (handler *PhysicalWireguardHandler) Run(ctx context.Context, logger *zap.Logger) error {
	return handler.wgDevice.Run(ctx, logger, handler)
}

// Peers implements WireguardHandler.
func (handler *PhysicalWireguardHandler) Peers() ([]wgtypes.Peer, error) {
	if handler.wgDevice == nil {
		return []wgtypes.Peer{}, nil
	}

	return handler.wgDevice.Peers()
}
