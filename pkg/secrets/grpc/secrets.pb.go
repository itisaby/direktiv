// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: pkg/secrets/grpc/secrets.proto

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

type SecretsStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Name      *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Data      []byte  `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *SecretsStoreRequest) Reset() {
	*x = SecretsStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsStoreRequest) ProtoMessage() {}

func (x *SecretsStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsStoreRequest.ProtoReflect.Descriptor instead.
func (*SecretsStoreRequest) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_grpc_secrets_proto_rawDescGZIP(), []int{0}
}

func (x *SecretsStoreRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *SecretsStoreRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *SecretsStoreRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SecretsRetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Name      *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *SecretsRetrieveRequest) Reset() {
	*x = SecretsRetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsRetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsRetrieveRequest) ProtoMessage() {}

func (x *SecretsRetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsRetrieveRequest.ProtoReflect.Descriptor instead.
func (*SecretsRetrieveRequest) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_grpc_secrets_proto_rawDescGZIP(), []int{1}
}

func (x *SecretsRetrieveRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *SecretsRetrieveRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type SecretsDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Name      *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *SecretsDeleteRequest) Reset() {
	*x = SecretsDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsDeleteRequest) ProtoMessage() {}

func (x *SecretsDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsDeleteRequest.ProtoReflect.Descriptor instead.
func (*SecretsDeleteRequest) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_grpc_secrets_proto_rawDescGZIP(), []int{2}
}

func (x *SecretsDeleteRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *SecretsDeleteRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type SecretsDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count *int32 `protobuf:"varint,1,opt,name=count,proto3,oneof" json:"count,omitempty"`
}

func (x *SecretsDeleteResponse) Reset() {
	*x = SecretsDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsDeleteResponse) ProtoMessage() {}

func (x *SecretsDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsDeleteResponse.ProtoReflect.Descriptor instead.
func (*SecretsDeleteResponse) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_grpc_secrets_proto_rawDescGZIP(), []int{3}
}

func (x *SecretsDeleteResponse) GetCount() int32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

type SecretsRetrieveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *SecretsRetrieveResponse) Reset() {
	*x = SecretsRetrieveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsRetrieveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsRetrieveResponse) ProtoMessage() {}

func (x *SecretsRetrieveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsRetrieveResponse.ProtoReflect.Descriptor instead.
func (*SecretsRetrieveResponse) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_grpc_secrets_proto_rawDescGZIP(), []int{4}
}

func (x *SecretsRetrieveResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetSecretsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
}

func (x *GetSecretsRequest) Reset() {
	*x = GetSecretsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretsRequest) ProtoMessage() {}

func (x *GetSecretsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretsRequest.ProtoReflect.Descriptor instead.
func (*GetSecretsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_grpc_secrets_proto_rawDescGZIP(), []int{5}
}

func (x *GetSecretsRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

type GetSecretsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secrets []*GetSecretsResponse_Secret `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty"`
}

func (x *GetSecretsResponse) Reset() {
	*x = GetSecretsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretsResponse) ProtoMessage() {}

func (x *GetSecretsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretsResponse.ProtoReflect.Descriptor instead.
func (*GetSecretsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_grpc_secrets_proto_rawDescGZIP(), []int{6}
}

func (x *GetSecretsResponse) GetSecrets() []*GetSecretsResponse_Secret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

type DeleteSecretsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
}

func (x *DeleteSecretsRequest) Reset() {
	*x = DeleteSecretsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSecretsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSecretsRequest) ProtoMessage() {}

func (x *DeleteSecretsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSecretsRequest.ProtoReflect.Descriptor instead.
func (*DeleteSecretsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_grpc_secrets_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSecretsRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

type GetSecretsResponse_Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *GetSecretsResponse_Secret) Reset() {
	*x = GetSecretsResponse_Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretsResponse_Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretsResponse_Secret) ProtoMessage() {}

func (x *GetSecretsResponse_Secret) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_grpc_secrets_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretsResponse_Secret.ProtoReflect.Descriptor instead.
func (*GetSecretsResponse_Secret) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_grpc_secrets_proto_rawDescGZIP(), []int{6, 0}
}

func (x *GetSecretsResponse_Secret) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

var File_pkg_secrets_grpc_secrets_proto protoreflect.FileDescriptor

var file_pkg_secrets_grpc_secrets_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x8a, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x02, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x6b, 0x0a, 0x16, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x69, 0x0a, 0x14, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x15, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x17, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x44, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x7b, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x1a, 0x2a, 0x0a,
	0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69,
	0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_secrets_grpc_secrets_proto_rawDescOnce sync.Once
	file_pkg_secrets_grpc_secrets_proto_rawDescData = file_pkg_secrets_grpc_secrets_proto_rawDesc
)

func file_pkg_secrets_grpc_secrets_proto_rawDescGZIP() []byte {
	file_pkg_secrets_grpc_secrets_proto_rawDescOnce.Do(func() {
		file_pkg_secrets_grpc_secrets_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_secrets_grpc_secrets_proto_rawDescData)
	})
	return file_pkg_secrets_grpc_secrets_proto_rawDescData
}

var file_pkg_secrets_grpc_secrets_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_secrets_grpc_secrets_proto_goTypes = []interface{}{
	(*SecretsStoreRequest)(nil),       // 0: grpc.SecretsStoreRequest
	(*SecretsRetrieveRequest)(nil),    // 1: grpc.SecretsRetrieveRequest
	(*SecretsDeleteRequest)(nil),      // 2: grpc.SecretsDeleteRequest
	(*SecretsDeleteResponse)(nil),     // 3: grpc.SecretsDeleteResponse
	(*SecretsRetrieveResponse)(nil),   // 4: grpc.SecretsRetrieveResponse
	(*GetSecretsRequest)(nil),         // 5: grpc.GetSecretsRequest
	(*GetSecretsResponse)(nil),        // 6: grpc.GetSecretsResponse
	(*DeleteSecretsRequest)(nil),      // 7: grpc.DeleteSecretsRequest
	(*GetSecretsResponse_Secret)(nil), // 8: grpc.GetSecretsResponse.Secret
}
var file_pkg_secrets_grpc_secrets_proto_depIdxs = []int32{
	8, // 0: grpc.GetSecretsResponse.secrets:type_name -> grpc.GetSecretsResponse.Secret
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_secrets_grpc_secrets_proto_init() }
func file_pkg_secrets_grpc_secrets_proto_init() {
	if File_pkg_secrets_grpc_secrets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_secrets_grpc_secrets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsStoreRequest); i {
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
		file_pkg_secrets_grpc_secrets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsRetrieveRequest); i {
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
		file_pkg_secrets_grpc_secrets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsDeleteRequest); i {
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
		file_pkg_secrets_grpc_secrets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsDeleteResponse); i {
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
		file_pkg_secrets_grpc_secrets_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsRetrieveResponse); i {
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
		file_pkg_secrets_grpc_secrets_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretsRequest); i {
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
		file_pkg_secrets_grpc_secrets_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretsResponse); i {
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
		file_pkg_secrets_grpc_secrets_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSecretsRequest); i {
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
		file_pkg_secrets_grpc_secrets_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretsResponse_Secret); i {
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
	file_pkg_secrets_grpc_secrets_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_secrets_grpc_secrets_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_pkg_secrets_grpc_secrets_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_pkg_secrets_grpc_secrets_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_pkg_secrets_grpc_secrets_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_pkg_secrets_grpc_secrets_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_pkg_secrets_grpc_secrets_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_pkg_secrets_grpc_secrets_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_secrets_grpc_secrets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_secrets_grpc_secrets_proto_goTypes,
		DependencyIndexes: file_pkg_secrets_grpc_secrets_proto_depIdxs,
		MessageInfos:      file_pkg_secrets_grpc_secrets_proto_msgTypes,
	}.Build()
	File_pkg_secrets_grpc_secrets_proto = out.File
	file_pkg_secrets_grpc_secrets_proto_rawDesc = nil
	file_pkg_secrets_grpc_secrets_proto_goTypes = nil
	file_pkg_secrets_grpc_secrets_proto_depIdxs = nil
}
