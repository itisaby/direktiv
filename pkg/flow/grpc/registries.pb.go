// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: pkg/flow/grpc/registries.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_registries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_registries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_registries_proto_rawDescGZIP(), []int{0}
}

func (x *Registry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegistryEdge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node   *Registry `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Cursor string    `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *RegistryEdge) Reset() {
	*x = RegistryEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_registries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryEdge) ProtoMessage() {}

func (x *RegistryEdge) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_registries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryEdge.ProtoReflect.Descriptor instead.
func (*RegistryEdge) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_registries_proto_rawDescGZIP(), []int{1}
}

func (x *RegistryEdge) GetNode() *Registry {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *RegistryEdge) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type Registries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int32           `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	PageInfo   *PageInfo       `protobuf:"bytes,2,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
	Edges      []*RegistryEdge `protobuf:"bytes,3,rep,name=edges,proto3" json:"edges,omitempty"`
}

func (x *Registries) Reset() {
	*x = Registries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_registries_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registries) ProtoMessage() {}

func (x *Registries) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_registries_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registries.ProtoReflect.Descriptor instead.
func (*Registries) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_registries_proto_rawDescGZIP(), []int{2}
}

func (x *Registries) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *Registries) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *Registries) GetEdges() []*RegistryEdge {
	if x != nil {
		return x.Edges
	}
	return nil
}

type NamespaceRegistriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Namespace  string      `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *NamespaceRegistriesRequest) Reset() {
	*x = NamespaceRegistriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_registries_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceRegistriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceRegistriesRequest) ProtoMessage() {}

func (x *NamespaceRegistriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_registries_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceRegistriesRequest.ProtoReflect.Descriptor instead.
func (*NamespaceRegistriesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_registries_proto_rawDescGZIP(), []int{3}
}

func (x *NamespaceRegistriesRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *NamespaceRegistriesRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type NamespaceRegistriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string      `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Registries *Registries `protobuf:"bytes,2,opt,name=registries,proto3" json:"registries,omitempty"`
}

func (x *NamespaceRegistriesResponse) Reset() {
	*x = NamespaceRegistriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_registries_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceRegistriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceRegistriesResponse) ProtoMessage() {}

func (x *NamespaceRegistriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_registries_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceRegistriesResponse.ProtoReflect.Descriptor instead.
func (*NamespaceRegistriesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_registries_proto_rawDescGZIP(), []int{4}
}

func (x *NamespaceRegistriesResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *NamespaceRegistriesResponse) GetRegistries() *Registries {
	if x != nil {
		return x.Registries
	}
	return nil
}

type DeleteNamespaceRegistryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Registry  string `protobuf:"bytes,2,opt,name=registry,proto3" json:"registry,omitempty"`
}

func (x *DeleteNamespaceRegistryRequest) Reset() {
	*x = DeleteNamespaceRegistryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_registries_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNamespaceRegistryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNamespaceRegistryRequest) ProtoMessage() {}

func (x *DeleteNamespaceRegistryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_registries_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNamespaceRegistryRequest.ProtoReflect.Descriptor instead.
func (*DeleteNamespaceRegistryRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_registries_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteNamespaceRegistryRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeleteNamespaceRegistryRequest) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

type SetNamespaceRegistryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Registry  string `protobuf:"bytes,2,opt,name=registry,proto3" json:"registry,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SetNamespaceRegistryRequest) Reset() {
	*x = SetNamespaceRegistryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_registries_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNamespaceRegistryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNamespaceRegistryRequest) ProtoMessage() {}

func (x *SetNamespaceRegistryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_registries_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNamespaceRegistryRequest.ProtoReflect.Descriptor instead.
func (*SetNamespaceRegistryRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_registries_proto_rawDescGZIP(), []int{6}
}

func (x *SetNamespaceRegistryRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SetNamespaceRegistryRequest) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

func (x *SetNamespaceRegistryRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pkg_flow_grpc_registries_proto protoreflect.FileDescriptor

var file_pkg_flow_grpc_registries_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x1a,
	0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1e, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x53, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x64, 0x67, 0x65, 0x12,
	0x2b, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x22, 0x94, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x64, 0x67, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74,
	0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x45, 0x64, 0x67, 0x65, 0x52, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x22, 0x75, 0x0a, 0x1a, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x22, 0x76, 0x0a, 0x1b, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x0a,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x5a, 0x0a, 0x1e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x22, 0x6b, 0x0a, 0x1b, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74,
	0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_flow_grpc_registries_proto_rawDescOnce sync.Once
	file_pkg_flow_grpc_registries_proto_rawDescData = file_pkg_flow_grpc_registries_proto_rawDesc
)

func file_pkg_flow_grpc_registries_proto_rawDescGZIP() []byte {
	file_pkg_flow_grpc_registries_proto_rawDescOnce.Do(func() {
		file_pkg_flow_grpc_registries_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_flow_grpc_registries_proto_rawDescData)
	})
	return file_pkg_flow_grpc_registries_proto_rawDescData
}

var file_pkg_flow_grpc_registries_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_flow_grpc_registries_proto_goTypes = []interface{}{
	(*Registry)(nil),                       // 0: direktiv_flow.Registry
	(*RegistryEdge)(nil),                   // 1: direktiv_flow.RegistryEdge
	(*Registries)(nil),                     // 2: direktiv_flow.Registries
	(*NamespaceRegistriesRequest)(nil),     // 3: direktiv_flow.NamespaceRegistriesRequest
	(*NamespaceRegistriesResponse)(nil),    // 4: direktiv_flow.NamespaceRegistriesResponse
	(*DeleteNamespaceRegistryRequest)(nil), // 5: direktiv_flow.DeleteNamespaceRegistryRequest
	(*SetNamespaceRegistryRequest)(nil),    // 6: direktiv_flow.SetNamespaceRegistryRequest
	(*PageInfo)(nil),                       // 7: direktiv_flow.PageInfo
	(*Pagination)(nil),                     // 8: direktiv_flow.Pagination
}
var file_pkg_flow_grpc_registries_proto_depIdxs = []int32{
	0, // 0: direktiv_flow.RegistryEdge.node:type_name -> direktiv_flow.Registry
	7, // 1: direktiv_flow.Registries.pageInfo:type_name -> direktiv_flow.PageInfo
	1, // 2: direktiv_flow.Registries.edges:type_name -> direktiv_flow.RegistryEdge
	8, // 3: direktiv_flow.NamespaceRegistriesRequest.pagination:type_name -> direktiv_flow.Pagination
	2, // 4: direktiv_flow.NamespaceRegistriesResponse.registries:type_name -> direktiv_flow.Registries
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_flow_grpc_registries_proto_init() }
func file_pkg_flow_grpc_registries_proto_init() {
	if File_pkg_flow_grpc_registries_proto != nil {
		return
	}
	file_pkg_flow_grpc_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_flow_grpc_registries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry); i {
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
		file_pkg_flow_grpc_registries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryEdge); i {
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
		file_pkg_flow_grpc_registries_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registries); i {
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
		file_pkg_flow_grpc_registries_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceRegistriesRequest); i {
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
		file_pkg_flow_grpc_registries_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceRegistriesResponse); i {
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
		file_pkg_flow_grpc_registries_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNamespaceRegistryRequest); i {
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
		file_pkg_flow_grpc_registries_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNamespaceRegistryRequest); i {
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
			RawDescriptor: file_pkg_flow_grpc_registries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_flow_grpc_registries_proto_goTypes,
		DependencyIndexes: file_pkg_flow_grpc_registries_proto_depIdxs,
		MessageInfos:      file_pkg_flow_grpc_registries_proto_msgTypes,
	}.Build()
	File_pkg_flow_grpc_registries_proto = out.File
	file_pkg_flow_grpc_registries_proto_rawDesc = nil
	file_pkg_flow_grpc_registries_proto_goTypes = nil
	file_pkg_flow_grpc_registries_proto_depIdxs = nil
}
