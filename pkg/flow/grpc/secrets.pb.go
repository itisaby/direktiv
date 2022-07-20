// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: pkg/flow/grpc/secrets.proto

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

type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_secrets_proto_rawDescGZIP(), []int{0}
}

func (x *Secret) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SecretEdge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node   *Secret `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Cursor string  `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *SecretEdge) Reset() {
	*x = SecretEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretEdge) ProtoMessage() {}

func (x *SecretEdge) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretEdge.ProtoReflect.Descriptor instead.
func (*SecretEdge) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_secrets_proto_rawDescGZIP(), []int{1}
}

func (x *SecretEdge) GetNode() *Secret {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *SecretEdge) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type Secrets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int32         `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	PageInfo   *PageInfo     `protobuf:"bytes,2,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
	Edges      []*SecretEdge `protobuf:"bytes,3,rep,name=edges,proto3" json:"edges,omitempty"`
}

func (x *Secrets) Reset() {
	*x = Secrets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secrets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secrets) ProtoMessage() {}

func (x *Secrets) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secrets.ProtoReflect.Descriptor instead.
func (*Secrets) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_secrets_proto_rawDescGZIP(), []int{2}
}

func (x *Secrets) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *Secrets) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *Secrets) GetEdges() []*SecretEdge {
	if x != nil {
		return x.Edges
	}
	return nil
}

type SecretsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Namespace  string      `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *SecretsRequest) Reset() {
	*x = SecretsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsRequest) ProtoMessage() {}

func (x *SecretsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsRequest.ProtoReflect.Descriptor instead.
func (*SecretsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_secrets_proto_rawDescGZIP(), []int{3}
}

func (x *SecretsRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *SecretsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type SecretsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Secrets   *Secrets `protobuf:"bytes,2,opt,name=secrets,proto3" json:"secrets,omitempty"`
}

func (x *SecretsResponse) Reset() {
	*x = SecretsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsResponse) ProtoMessage() {}

func (x *SecretsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsResponse.ProtoReflect.Descriptor instead.
func (*SecretsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_secrets_proto_rawDescGZIP(), []int{4}
}

func (x *SecretsResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SecretsResponse) GetSecrets() *Secrets {
	if x != nil {
		return x.Secrets
	}
	return nil
}

type SetSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SetSecretRequest) Reset() {
	*x = SetSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSecretRequest) ProtoMessage() {}

func (x *SetSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSecretRequest.ProtoReflect.Descriptor instead.
func (*SetSecretRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_secrets_proto_rawDescGZIP(), []int{5}
}

func (x *SetSecretRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SetSecretRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetSecretRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SetSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *SetSecretResponse) Reset() {
	*x = SetSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSecretResponse) ProtoMessage() {}

func (x *SetSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSecretResponse.ProtoReflect.Descriptor instead.
func (*SetSecretResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_secrets_proto_rawDescGZIP(), []int{6}
}

func (x *SetSecretResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SetSecretResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteSecretRequest) Reset() {
	*x = DeleteSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSecretRequest) ProtoMessage() {}

func (x *DeleteSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_secrets_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSecretRequest.ProtoReflect.Descriptor instead.
func (*DeleteSecretRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_secrets_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSecretRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeleteSecretRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_pkg_flow_grpc_secrets_proto protoreflect.FileDescriptor

var file_pkg_flow_grpc_secrets_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64,
	0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x1e, 0x70, 0x6b,
	0x67, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x06,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x0a, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x45, 0x64, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69,
	0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x8f, 0x01, 0x0a, 0x07,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x69, 0x72, 0x65,
	0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x05,
	0x65, 0x64, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x69,
	0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x45, 0x64, 0x67, 0x65, 0x52, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x22, 0x69, 0x0a,
	0x0e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x61, 0x0a, 0x0f, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x22, 0x56, 0x0a, 0x10, 0x53,
	0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x43, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x45, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x42,
	0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69,
	0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_flow_grpc_secrets_proto_rawDescOnce sync.Once
	file_pkg_flow_grpc_secrets_proto_rawDescData = file_pkg_flow_grpc_secrets_proto_rawDesc
)

func file_pkg_flow_grpc_secrets_proto_rawDescGZIP() []byte {
	file_pkg_flow_grpc_secrets_proto_rawDescOnce.Do(func() {
		file_pkg_flow_grpc_secrets_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_flow_grpc_secrets_proto_rawDescData)
	})
	return file_pkg_flow_grpc_secrets_proto_rawDescData
}

var file_pkg_flow_grpc_secrets_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_flow_grpc_secrets_proto_goTypes = []interface{}{
	(*Secret)(nil),              // 0: direktiv_flow.Secret
	(*SecretEdge)(nil),          // 1: direktiv_flow.SecretEdge
	(*Secrets)(nil),             // 2: direktiv_flow.Secrets
	(*SecretsRequest)(nil),      // 3: direktiv_flow.SecretsRequest
	(*SecretsResponse)(nil),     // 4: direktiv_flow.SecretsResponse
	(*SetSecretRequest)(nil),    // 5: direktiv_flow.SetSecretRequest
	(*SetSecretResponse)(nil),   // 6: direktiv_flow.SetSecretResponse
	(*DeleteSecretRequest)(nil), // 7: direktiv_flow.DeleteSecretRequest
	(*PageInfo)(nil),            // 8: direktiv_flow.PageInfo
	(*Pagination)(nil),          // 9: direktiv_flow.Pagination
}
var file_pkg_flow_grpc_secrets_proto_depIdxs = []int32{
	0, // 0: direktiv_flow.SecretEdge.node:type_name -> direktiv_flow.Secret
	8, // 1: direktiv_flow.Secrets.pageInfo:type_name -> direktiv_flow.PageInfo
	1, // 2: direktiv_flow.Secrets.edges:type_name -> direktiv_flow.SecretEdge
	9, // 3: direktiv_flow.SecretsRequest.pagination:type_name -> direktiv_flow.Pagination
	2, // 4: direktiv_flow.SecretsResponse.secrets:type_name -> direktiv_flow.Secrets
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_flow_grpc_secrets_proto_init() }
func file_pkg_flow_grpc_secrets_proto_init() {
	if File_pkg_flow_grpc_secrets_proto != nil {
		return
	}
	file_pkg_flow_grpc_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_flow_grpc_secrets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret); i {
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
		file_pkg_flow_grpc_secrets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretEdge); i {
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
		file_pkg_flow_grpc_secrets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secrets); i {
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
		file_pkg_flow_grpc_secrets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsRequest); i {
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
		file_pkg_flow_grpc_secrets_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsResponse); i {
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
		file_pkg_flow_grpc_secrets_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSecretRequest); i {
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
		file_pkg_flow_grpc_secrets_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSecretResponse); i {
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
		file_pkg_flow_grpc_secrets_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSecretRequest); i {
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
			RawDescriptor: file_pkg_flow_grpc_secrets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_flow_grpc_secrets_proto_goTypes,
		DependencyIndexes: file_pkg_flow_grpc_secrets_proto_depIdxs,
		MessageInfos:      file_pkg_flow_grpc_secrets_proto_msgTypes,
	}.Build()
	File_pkg_flow_grpc_secrets_proto = out.File
	file_pkg_flow_grpc_secrets_proto_rawDesc = nil
	file_pkg_flow_grpc_secrets_proto_goTypes = nil
	file_pkg_flow_grpc_secrets_proto_depIdxs = nil
}
