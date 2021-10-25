// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: pkg/functions/grpc/watch-pods.proto

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

type WatchPodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *string   `protobuf:"bytes,1,opt,name=event,proto3,oneof" json:"event,omitempty"`
	Pod   *PodsInfo `protobuf:"bytes,2,opt,name=pod,proto3,oneof" json:"pod,omitempty"`
}

func (x *WatchPodsResponse) Reset() {
	*x = WatchPodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_watch_pods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchPodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchPodsResponse) ProtoMessage() {}

func (x *WatchPodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_watch_pods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchPodsResponse.ProtoReflect.Descriptor instead.
func (*WatchPodsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_watch_pods_proto_rawDescGZIP(), []int{0}
}

func (x *WatchPodsResponse) GetEvent() string {
	if x != nil && x.Event != nil {
		return *x.Event
	}
	return ""
}

func (x *WatchPodsResponse) GetPod() *PodsInfo {
	if x != nil {
		return x.Pod
	}
	return nil
}

type WatchPodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName  *string `protobuf:"bytes,1,opt,name=ServiceName,proto3,oneof" json:"ServiceName,omitempty"`
	RevisionName *string `protobuf:"bytes,2,opt,name=RevisionName,proto3,oneof" json:"RevisionName,omitempty"`
}

func (x *WatchPodsRequest) Reset() {
	*x = WatchPodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_watch_pods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchPodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchPodsRequest) ProtoMessage() {}

func (x *WatchPodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_watch_pods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchPodsRequest.ProtoReflect.Descriptor instead.
func (*WatchPodsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_watch_pods_proto_rawDescGZIP(), []int{1}
}

func (x *WatchPodsRequest) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

func (x *WatchPodsRequest) GetRevisionName() string {
	if x != nil && x.RevisionName != nil {
		return *x.RevisionName
	}
	return ""
}

var File_pkg_functions_grpc_watch_pods_proto protoreflect.FileDescriptor

var file_pkg_functions_grpc_watch_pods_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2d, 0x70, 0x6f, 0x64, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x11, 0x57, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x01, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x70, 0x6f, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x10,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x0c, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_functions_grpc_watch_pods_proto_rawDescOnce sync.Once
	file_pkg_functions_grpc_watch_pods_proto_rawDescData = file_pkg_functions_grpc_watch_pods_proto_rawDesc
)

func file_pkg_functions_grpc_watch_pods_proto_rawDescGZIP() []byte {
	file_pkg_functions_grpc_watch_pods_proto_rawDescOnce.Do(func() {
		file_pkg_functions_grpc_watch_pods_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_functions_grpc_watch_pods_proto_rawDescData)
	})
	return file_pkg_functions_grpc_watch_pods_proto_rawDescData
}

var file_pkg_functions_grpc_watch_pods_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_functions_grpc_watch_pods_proto_goTypes = []interface{}{
	(*WatchPodsResponse)(nil), // 0: direktiv_functions.WatchPodsResponse
	(*WatchPodsRequest)(nil),  // 1: direktiv_functions.WatchPodsRequest
	(*PodsInfo)(nil),          // 2: direktiv_functions.PodsInfo
}
var file_pkg_functions_grpc_watch_pods_proto_depIdxs = []int32{
	2, // 0: direktiv_functions.WatchPodsResponse.pod:type_name -> direktiv_functions.PodsInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_functions_grpc_watch_pods_proto_init() }
func file_pkg_functions_grpc_watch_pods_proto_init() {
	if File_pkg_functions_grpc_watch_pods_proto != nil {
		return
	}
	file_pkg_functions_grpc_functions_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_functions_grpc_watch_pods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchPodsResponse); i {
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
		file_pkg_functions_grpc_watch_pods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchPodsRequest); i {
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
	file_pkg_functions_grpc_watch_pods_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_functions_grpc_watch_pods_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_functions_grpc_watch_pods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_functions_grpc_watch_pods_proto_goTypes,
		DependencyIndexes: file_pkg_functions_grpc_watch_pods_proto_depIdxs,
		MessageInfos:      file_pkg_functions_grpc_watch_pods_proto_msgTypes,
	}.Build()
	File_pkg_functions_grpc_watch_pods_proto = out.File
	file_pkg_functions_grpc_watch_pods_proto_rawDesc = nil
	file_pkg_functions_grpc_watch_pods_proto_goTypes = nil
	file_pkg_functions_grpc_watch_pods_proto_depIdxs = nil
}
