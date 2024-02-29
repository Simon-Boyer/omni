// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: omni/specs/virtual.proto

package specs

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CurrentUserSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Role     string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CurrentUserSpec) Reset() {
	*x = CurrentUserSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omni_specs_virtual_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentUserSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentUserSpec) ProtoMessage() {}

func (x *CurrentUserSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_virtual_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentUserSpec.ProtoReflect.Descriptor instead.
func (*CurrentUserSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_virtual_proto_rawDescGZIP(), []int{0}
}

func (x *CurrentUserSpec) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *CurrentUserSpec) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type PermissionsSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanReadClusters               bool `protobuf:"varint,1,opt,name=can_read_clusters,json=canReadClusters,proto3" json:"can_read_clusters,omitempty"`
	CanCreateClusters             bool `protobuf:"varint,2,opt,name=can_create_clusters,json=canCreateClusters,proto3" json:"can_create_clusters,omitempty"`
	CanManageUsers                bool `protobuf:"varint,3,opt,name=can_manage_users,json=canManageUsers,proto3" json:"can_manage_users,omitempty"`
	CanReadMachines               bool `protobuf:"varint,4,opt,name=can_read_machines,json=canReadMachines,proto3" json:"can_read_machines,omitempty"`
	CanRemoveMachines             bool `protobuf:"varint,5,opt,name=can_remove_machines,json=canRemoveMachines,proto3" json:"can_remove_machines,omitempty"`
	CanReadMachineLogs            bool `protobuf:"varint,6,opt,name=can_read_machine_logs,json=canReadMachineLogs,proto3" json:"can_read_machine_logs,omitempty"`
	CanReadMachineConfigPatches   bool `protobuf:"varint,9,opt,name=can_read_machine_config_patches,json=canReadMachineConfigPatches,proto3" json:"can_read_machine_config_patches,omitempty"`
	CanManageMachineConfigPatches bool `protobuf:"varint,10,opt,name=can_manage_machine_config_patches,json=canManageMachineConfigPatches,proto3" json:"can_manage_machine_config_patches,omitempty"`
	CanManageBackupStore          bool `protobuf:"varint,11,opt,name=can_manage_backup_store,json=canManageBackupStore,proto3" json:"can_manage_backup_store,omitempty"`
}

func (x *PermissionsSpec) Reset() {
	*x = PermissionsSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omni_specs_virtual_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionsSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionsSpec) ProtoMessage() {}

func (x *PermissionsSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_virtual_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionsSpec.ProtoReflect.Descriptor instead.
func (*PermissionsSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_virtual_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionsSpec) GetCanReadClusters() bool {
	if x != nil {
		return x.CanReadClusters
	}
	return false
}

func (x *PermissionsSpec) GetCanCreateClusters() bool {
	if x != nil {
		return x.CanCreateClusters
	}
	return false
}

func (x *PermissionsSpec) GetCanManageUsers() bool {
	if x != nil {
		return x.CanManageUsers
	}
	return false
}

func (x *PermissionsSpec) GetCanReadMachines() bool {
	if x != nil {
		return x.CanReadMachines
	}
	return false
}

func (x *PermissionsSpec) GetCanRemoveMachines() bool {
	if x != nil {
		return x.CanRemoveMachines
	}
	return false
}

func (x *PermissionsSpec) GetCanReadMachineLogs() bool {
	if x != nil {
		return x.CanReadMachineLogs
	}
	return false
}

func (x *PermissionsSpec) GetCanReadMachineConfigPatches() bool {
	if x != nil {
		return x.CanReadMachineConfigPatches
	}
	return false
}

func (x *PermissionsSpec) GetCanManageMachineConfigPatches() bool {
	if x != nil {
		return x.CanManageMachineConfigPatches
	}
	return false
}

func (x *PermissionsSpec) GetCanManageBackupStore() bool {
	if x != nil {
		return x.CanManageBackupStore
	}
	return false
}

type ClusterPermissionsSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanAddMachines             bool `protobuf:"varint,1,opt,name=can_add_machines,json=canAddMachines,proto3" json:"can_add_machines,omitempty"`
	CanRemoveMachines          bool `protobuf:"varint,2,opt,name=can_remove_machines,json=canRemoveMachines,proto3" json:"can_remove_machines,omitempty"`
	CanRebootMachines          bool `protobuf:"varint,3,opt,name=can_reboot_machines,json=canRebootMachines,proto3" json:"can_reboot_machines,omitempty"`
	CanUpdateKubernetes        bool `protobuf:"varint,4,opt,name=can_update_kubernetes,json=canUpdateKubernetes,proto3" json:"can_update_kubernetes,omitempty"`
	CanDownloadKubeconfig      bool `protobuf:"varint,5,opt,name=can_download_kubeconfig,json=canDownloadKubeconfig,proto3" json:"can_download_kubeconfig,omitempty"`
	CanSyncKubernetesManifests bool `protobuf:"varint,6,opt,name=can_sync_kubernetes_manifests,json=canSyncKubernetesManifests,proto3" json:"can_sync_kubernetes_manifests,omitempty"`
	CanUpdateTalos             bool `protobuf:"varint,7,opt,name=can_update_talos,json=canUpdateTalos,proto3" json:"can_update_talos,omitempty"`
	CanDownloadTalosconfig     bool `protobuf:"varint,8,opt,name=can_download_talosconfig,json=canDownloadTalosconfig,proto3" json:"can_download_talosconfig,omitempty"`
	CanReadConfigPatches       bool `protobuf:"varint,9,opt,name=can_read_config_patches,json=canReadConfigPatches,proto3" json:"can_read_config_patches,omitempty"`
	CanManageConfigPatches     bool `protobuf:"varint,10,opt,name=can_manage_config_patches,json=canManageConfigPatches,proto3" json:"can_manage_config_patches,omitempty"`
	CanManageClusterFeatures   bool `protobuf:"varint,11,opt,name=can_manage_cluster_features,json=canManageClusterFeatures,proto3" json:"can_manage_cluster_features,omitempty"`
}

func (x *ClusterPermissionsSpec) Reset() {
	*x = ClusterPermissionsSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omni_specs_virtual_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterPermissionsSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterPermissionsSpec) ProtoMessage() {}

func (x *ClusterPermissionsSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_virtual_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterPermissionsSpec.ProtoReflect.Descriptor instead.
func (*ClusterPermissionsSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_virtual_proto_rawDescGZIP(), []int{2}
}

func (x *ClusterPermissionsSpec) GetCanAddMachines() bool {
	if x != nil {
		return x.CanAddMachines
	}
	return false
}

func (x *ClusterPermissionsSpec) GetCanRemoveMachines() bool {
	if x != nil {
		return x.CanRemoveMachines
	}
	return false
}

func (x *ClusterPermissionsSpec) GetCanRebootMachines() bool {
	if x != nil {
		return x.CanRebootMachines
	}
	return false
}

func (x *ClusterPermissionsSpec) GetCanUpdateKubernetes() bool {
	if x != nil {
		return x.CanUpdateKubernetes
	}
	return false
}

func (x *ClusterPermissionsSpec) GetCanDownloadKubeconfig() bool {
	if x != nil {
		return x.CanDownloadKubeconfig
	}
	return false
}

func (x *ClusterPermissionsSpec) GetCanSyncKubernetesManifests() bool {
	if x != nil {
		return x.CanSyncKubernetesManifests
	}
	return false
}

func (x *ClusterPermissionsSpec) GetCanUpdateTalos() bool {
	if x != nil {
		return x.CanUpdateTalos
	}
	return false
}

func (x *ClusterPermissionsSpec) GetCanDownloadTalosconfig() bool {
	if x != nil {
		return x.CanDownloadTalosconfig
	}
	return false
}

func (x *ClusterPermissionsSpec) GetCanReadConfigPatches() bool {
	if x != nil {
		return x.CanReadConfigPatches
	}
	return false
}

func (x *ClusterPermissionsSpec) GetCanManageConfigPatches() bool {
	if x != nil {
		return x.CanManageConfigPatches
	}
	return false
}

func (x *ClusterPermissionsSpec) GetCanManageClusterFeatures() bool {
	if x != nil {
		return x.CanManageClusterFeatures
	}
	return false
}

var File_omni_specs_virtual_proto protoreflect.FileDescriptor

var file_omni_specs_virtual_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x70, 0x65, 0x63,
	0x73, 0x22, 0x47, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0xed, 0x03, 0x0a, 0x0f, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2a,
	0x0a, 0x11, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x61, 0x6e, 0x52, 0x65,
	0x61, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x61,
	0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63, 0x61, 0x6e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x61,
	0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x61, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73,
	0x12, 0x2e, 0x0a, 0x13, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63,
	0x61, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73,
	0x12, 0x31, 0x0a, 0x15, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x12, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4c,
	0x6f, 0x67, 0x73, 0x12, 0x44, 0x0a, 0x1f, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x63, 0x61,
	0x6e, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x50, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x21, 0x63, 0x61, 0x6e,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x63, 0x61, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x17, 0x63, 0x61, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x63, 0x61, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xe6, 0x04, 0x0a, 0x16, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x61, 0x6e, 0x5f, 0x61, 0x64, 0x64,
	0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x63, 0x61, 0x6e, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x12,
	0x2e, 0x0a, 0x13, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63, 0x61,
	0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x12,
	0x2e, 0x0a, 0x13, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63, 0x61,
	0x6e, 0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x12,
	0x32, 0x0a, 0x15, 0x63, 0x61, 0x6e, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13,
	0x63, 0x61, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x61, 0x6e, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x63, 0x61, 0x6e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x1d, 0x63,
	0x61, 0x6e, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x1a, 0x63, 0x61, 0x6e, 0x53, 0x79, 0x6e, 0x63, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x73, 0x12, 0x28,
	0x0a, 0x10, 0x63, 0x61, 0x6e, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x6c,
	0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x61, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x6c, 0x6f, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x63, 0x61, 0x6e, 0x5f,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x63, 0x61, 0x6e, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x61, 0x6c, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x35, 0x0a, 0x17, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x14, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x50, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x63, 0x61, 0x6e,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x63, 0x61,
	0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x63, 0x61, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x63, 0x61, 0x6e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x6f, 0x6d, 0x6e,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x6d, 0x6e,
	0x69, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_omni_specs_virtual_proto_rawDescOnce sync.Once
	file_omni_specs_virtual_proto_rawDescData = file_omni_specs_virtual_proto_rawDesc
)

func file_omni_specs_virtual_proto_rawDescGZIP() []byte {
	file_omni_specs_virtual_proto_rawDescOnce.Do(func() {
		file_omni_specs_virtual_proto_rawDescData = protoimpl.X.CompressGZIP(file_omni_specs_virtual_proto_rawDescData)
	})
	return file_omni_specs_virtual_proto_rawDescData
}

var file_omni_specs_virtual_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_omni_specs_virtual_proto_goTypes = []interface{}{
	(*CurrentUserSpec)(nil),        // 0: specs.CurrentUserSpec
	(*PermissionsSpec)(nil),        // 1: specs.PermissionsSpec
	(*ClusterPermissionsSpec)(nil), // 2: specs.ClusterPermissionsSpec
}
var file_omni_specs_virtual_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_omni_specs_virtual_proto_init() }
func file_omni_specs_virtual_proto_init() {
	if File_omni_specs_virtual_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_omni_specs_virtual_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentUserSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omni_specs_virtual_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionsSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omni_specs_virtual_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterPermissionsSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_omni_specs_virtual_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_omni_specs_virtual_proto_goTypes,
		DependencyIndexes: file_omni_specs_virtual_proto_depIdxs,
		MessageInfos:      file_omni_specs_virtual_proto_msgTypes,
	}.Build()
	File_omni_specs_virtual_proto = out.File
	file_omni_specs_virtual_proto_rawDesc = nil
	file_omni_specs_virtual_proto_goTypes = nil
	file_omni_specs_virtual_proto_depIdxs = nil
}
