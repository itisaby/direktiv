// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: pkg/flow/grpc/nodes.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Name         string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Path         string                 `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Parent       string                 `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
	Type         string                 `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Attributes   []string               `protobuf:"bytes,7,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Oid          string                 `protobuf:"bytes,8,opt,name=oid,proto3" json:"oid,omitempty"`
	ReadOnly     bool                   `protobuf:"varint,9,opt,name=readOnly,proto3" json:"readOnly,omitempty"`
	ExpandedType string                 `protobuf:"bytes,10,opt,name=expandedType,proto3" json:"expandedType,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_nodes_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Node) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Node) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *Node) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Node) GetAttributes() []string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Node) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

func (x *Node) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

func (x *Node) GetExpandedType() string {
	if x != nil {
		return x.ExpandedType
	}
	return ""
}

type DirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Namespace  string      `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path       string      `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DirectoryRequest) Reset() {
	*x = DirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryRequest) ProtoMessage() {}

func (x *DirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryRequest.ProtoReflect.Descriptor instead.
func (*DirectoryRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_nodes_proto_rawDescGZIP(), []int{1}
}

func (x *DirectoryRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *DirectoryRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DirectoryRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DirectoryResponseEdge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node   *Node  `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *DirectoryResponseEdge) Reset() {
	*x = DirectoryResponseEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryResponseEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryResponseEdge) ProtoMessage() {}

func (x *DirectoryResponseEdge) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryResponseEdge.ProtoReflect.Descriptor instead.
func (*DirectoryResponseEdge) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_nodes_proto_rawDescGZIP(), []int{2}
}

func (x *DirectoryResponseEdge) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *DirectoryResponseEdge) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type DirectoryChildren struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int32                    `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	PageInfo   *PageInfo                `protobuf:"bytes,2,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
	Edges      []*DirectoryResponseEdge `protobuf:"bytes,3,rep,name=edges,proto3" json:"edges,omitempty"`
}

func (x *DirectoryChildren) Reset() {
	*x = DirectoryChildren{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryChildren) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryChildren) ProtoMessage() {}

func (x *DirectoryChildren) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryChildren.ProtoReflect.Descriptor instead.
func (*DirectoryChildren) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_nodes_proto_rawDescGZIP(), []int{3}
}

func (x *DirectoryChildren) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *DirectoryChildren) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *DirectoryChildren) GetEdges() []*DirectoryResponseEdge {
	if x != nil {
		return x.Edges
	}
	return nil
}

type DirectoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string             `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Node      *Node              `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	Children  *DirectoryChildren `protobuf:"bytes,3,opt,name=children,proto3" json:"children,omitempty"`
}

func (x *DirectoryResponse) Reset() {
	*x = DirectoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryResponse) ProtoMessage() {}

func (x *DirectoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryResponse.ProtoReflect.Descriptor instead.
func (*DirectoryResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_nodes_proto_rawDescGZIP(), []int{4}
}

func (x *DirectoryResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DirectoryResponse) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *DirectoryResponse) GetChildren() *DirectoryChildren {
	if x != nil {
		return x.Children
	}
	return nil
}

type CreateDirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path       string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Idempotent bool   `protobuf:"varint,3,opt,name=idempotent,proto3" json:"idempotent,omitempty"`
	Parents    bool   `protobuf:"varint,4,opt,name=parents,proto3" json:"parents,omitempty"`
}

func (x *CreateDirectoryRequest) Reset() {
	*x = CreateDirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDirectoryRequest) ProtoMessage() {}

func (x *CreateDirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDirectoryRequest.ProtoReflect.Descriptor instead.
func (*CreateDirectoryRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_nodes_proto_rawDescGZIP(), []int{5}
}

func (x *CreateDirectoryRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateDirectoryRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateDirectoryRequest) GetIdempotent() bool {
	if x != nil {
		return x.Idempotent
	}
	return false
}

func (x *CreateDirectoryRequest) GetParents() bool {
	if x != nil {
		return x.Parents
	}
	return false
}

type CreateDirectoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Node      *Node  `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *CreateDirectoryResponse) Reset() {
	*x = CreateDirectoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDirectoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDirectoryResponse) ProtoMessage() {}

func (x *CreateDirectoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDirectoryResponse.ProtoReflect.Descriptor instead.
func (*CreateDirectoryResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_nodes_proto_rawDescGZIP(), []int{6}
}

func (x *CreateDirectoryResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateDirectoryResponse) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

type DeleteNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path       string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Idempotent bool   `protobuf:"varint,3,opt,name=idempotent,proto3" json:"idempotent,omitempty"`
	Recursive  bool   `protobuf:"varint,4,opt,name=recursive,proto3" json:"recursive,omitempty"`
}

func (x *DeleteNodeRequest) Reset() {
	*x = DeleteNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeRequest) ProtoMessage() {}

func (x *DeleteNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeRequest.ProtoReflect.Descriptor instead.
func (*DeleteNodeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_nodes_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteNodeRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeleteNodeRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *DeleteNodeRequest) GetIdempotent() bool {
	if x != nil {
		return x.Idempotent
	}
	return false
}

func (x *DeleteNodeRequest) GetRecursive() bool {
	if x != nil {
		return x.Recursive
	}
	return false
}

type RenameNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Old       string `protobuf:"bytes,2,opt,name=old,proto3" json:"old,omitempty"`
	New       string `protobuf:"bytes,3,opt,name=new,proto3" json:"new,omitempty"`
}

func (x *RenameNodeRequest) Reset() {
	*x = RenameNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameNodeRequest) ProtoMessage() {}

func (x *RenameNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameNodeRequest.ProtoReflect.Descriptor instead.
func (*RenameNodeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_nodes_proto_rawDescGZIP(), []int{8}
}

func (x *RenameNodeRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RenameNodeRequest) GetOld() string {
	if x != nil {
		return x.Old
	}
	return ""
}

func (x *RenameNodeRequest) GetNew() string {
	if x != nil {
		return x.New
	}
	return ""
}

type RenameNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Node      *Node  `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *RenameNodeResponse) Reset() {
	*x = RenameNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameNodeResponse) ProtoMessage() {}

func (x *RenameNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameNodeResponse.ProtoReflect.Descriptor instead.
func (*RenameNodeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_nodes_proto_rawDescGZIP(), []int{9}
}

func (x *RenameNodeResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RenameNodeResponse) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

type CreateNodeAttributesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path       string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Attributes []string `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *CreateNodeAttributesRequest) Reset() {
	*x = CreateNodeAttributesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeAttributesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeAttributesRequest) ProtoMessage() {}

func (x *CreateNodeAttributesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeAttributesRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeAttributesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_nodes_proto_rawDescGZIP(), []int{10}
}

func (x *CreateNodeAttributesRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateNodeAttributesRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateNodeAttributesRequest) GetAttributes() []string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type DeleteNodeAttributesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path       string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Attributes []string `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *DeleteNodeAttributesRequest) Reset() {
	*x = DeleteNodeAttributesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNodeAttributesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeAttributesRequest) ProtoMessage() {}

func (x *DeleteNodeAttributesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_nodes_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeAttributesRequest.ProtoReflect.Descriptor instead.
func (*DeleteNodeAttributesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_nodes_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteNodeAttributesRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeleteNodeAttributesRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *DeleteNodeAttributesRequest) GetAttributes() []string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

var File_pkg_flow_grpc_nodes_proto protoreflect.FileDescriptor

var file_pkg_flow_grpc_nodes_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x6b, 0x67,
	0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x02, 0x0a, 0x04,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c,
	0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x7f, 0x0a, 0x10, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b,
	0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x22, 0x58, 0x0a, 0x15, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x64, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b,
	0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0xa4, 0x01, 0x0a, 0x11,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x64, 0x67, 0x65, 0x52, 0x05, 0x65, 0x64, 0x67,
	0x65, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x11, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12,
	0x3c, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x72, 0x65, 0x6e, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x84, 0x01,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64,
	0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x60, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x27, 0x0a,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x69,
	0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x22, 0x55, 0x0a, 0x11,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x6c,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6e, 0x65, 0x77, 0x22, 0x5b, 0x0a, 0x12, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65,
	0x22, 0x6f, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x22, 0x6f, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74,
	0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_flow_grpc_nodes_proto_rawDescOnce sync.Once
	file_pkg_flow_grpc_nodes_proto_rawDescData = file_pkg_flow_grpc_nodes_proto_rawDesc
)

func file_pkg_flow_grpc_nodes_proto_rawDescGZIP() []byte {
	file_pkg_flow_grpc_nodes_proto_rawDescOnce.Do(func() {
		file_pkg_flow_grpc_nodes_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_flow_grpc_nodes_proto_rawDescData)
	})
	return file_pkg_flow_grpc_nodes_proto_rawDescData
}

var file_pkg_flow_grpc_nodes_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_pkg_flow_grpc_nodes_proto_goTypes = []interface{}{
	(*Node)(nil),                        // 0: direktiv_flow.Node
	(*DirectoryRequest)(nil),            // 1: direktiv_flow.DirectoryRequest
	(*DirectoryResponseEdge)(nil),       // 2: direktiv_flow.DirectoryResponseEdge
	(*DirectoryChildren)(nil),           // 3: direktiv_flow.DirectoryChildren
	(*DirectoryResponse)(nil),           // 4: direktiv_flow.DirectoryResponse
	(*CreateDirectoryRequest)(nil),      // 5: direktiv_flow.CreateDirectoryRequest
	(*CreateDirectoryResponse)(nil),     // 6: direktiv_flow.CreateDirectoryResponse
	(*DeleteNodeRequest)(nil),           // 7: direktiv_flow.DeleteNodeRequest
	(*RenameNodeRequest)(nil),           // 8: direktiv_flow.RenameNodeRequest
	(*RenameNodeResponse)(nil),          // 9: direktiv_flow.RenameNodeResponse
	(*CreateNodeAttributesRequest)(nil), // 10: direktiv_flow.CreateNodeAttributesRequest
	(*DeleteNodeAttributesRequest)(nil), // 11: direktiv_flow.DeleteNodeAttributesRequest
	(*timestamppb.Timestamp)(nil),       // 12: google.protobuf.Timestamp
	(*Pagination)(nil),                  // 13: direktiv_flow.Pagination
	(*PageInfo)(nil),                    // 14: direktiv_flow.PageInfo
}
var file_pkg_flow_grpc_nodes_proto_depIdxs = []int32{
	12, // 0: direktiv_flow.Node.created_at:type_name -> google.protobuf.Timestamp
	12, // 1: direktiv_flow.Node.updated_at:type_name -> google.protobuf.Timestamp
	13, // 2: direktiv_flow.DirectoryRequest.pagination:type_name -> direktiv_flow.Pagination
	0,  // 3: direktiv_flow.DirectoryResponseEdge.node:type_name -> direktiv_flow.Node
	14, // 4: direktiv_flow.DirectoryChildren.pageInfo:type_name -> direktiv_flow.PageInfo
	2,  // 5: direktiv_flow.DirectoryChildren.edges:type_name -> direktiv_flow.DirectoryResponseEdge
	0,  // 6: direktiv_flow.DirectoryResponse.node:type_name -> direktiv_flow.Node
	3,  // 7: direktiv_flow.DirectoryResponse.children:type_name -> direktiv_flow.DirectoryChildren
	0,  // 8: direktiv_flow.CreateDirectoryResponse.node:type_name -> direktiv_flow.Node
	0,  // 9: direktiv_flow.RenameNodeResponse.node:type_name -> direktiv_flow.Node
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pkg_flow_grpc_nodes_proto_init() }
func file_pkg_flow_grpc_nodes_proto_init() {
	if File_pkg_flow_grpc_nodes_proto != nil {
		return
	}
	file_pkg_flow_grpc_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_flow_grpc_nodes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_pkg_flow_grpc_nodes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryRequest); i {
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
		file_pkg_flow_grpc_nodes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryResponseEdge); i {
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
		file_pkg_flow_grpc_nodes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryChildren); i {
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
		file_pkg_flow_grpc_nodes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryResponse); i {
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
		file_pkg_flow_grpc_nodes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDirectoryRequest); i {
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
		file_pkg_flow_grpc_nodes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDirectoryResponse); i {
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
		file_pkg_flow_grpc_nodes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNodeRequest); i {
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
		file_pkg_flow_grpc_nodes_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameNodeRequest); i {
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
		file_pkg_flow_grpc_nodes_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameNodeResponse); i {
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
		file_pkg_flow_grpc_nodes_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeAttributesRequest); i {
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
		file_pkg_flow_grpc_nodes_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNodeAttributesRequest); i {
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
			RawDescriptor: file_pkg_flow_grpc_nodes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_flow_grpc_nodes_proto_goTypes,
		DependencyIndexes: file_pkg_flow_grpc_nodes_proto_depIdxs,
		MessageInfos:      file_pkg_flow_grpc_nodes_proto_msgTypes,
	}.Build()
	File_pkg_flow_grpc_nodes_proto = out.File
	file_pkg_flow_grpc_nodes_proto_rawDesc = nil
	file_pkg_flow_grpc_nodes_proto_goTypes = nil
	file_pkg_flow_grpc_nodes_proto_depIdxs = nil
}
