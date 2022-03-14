// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: protobuf/mover.micro.proto

package moverpb

import (
	storagespb "github.com/nodeum-io/nodeum-plugins/protobuf/types/storagespb"
	taskpb "github.com/nodeum-io/nodeum-plugins/protobuf/types/taskpb"
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

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GeneralQueue  string              `protobuf:"bytes,1,opt,name=general_queue,proto3" json:"general_queue,omitempty"`
	RequestsQueue string              `protobuf:"bytes,2,opt,name=requests_queue,proto3" json:"requests_queue,omitempty"`
	ResultsQueue  string              `protobuf:"bytes,3,opt,name=results_queue,proto3" json:"results_queue,omitempty"`
	Execution     *taskpb.Execution   `protobuf:"bytes,4,opt,name=execution,proto3" json:"execution,omitempty"`
	Source        *storagespb.Storage `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	Destination   *storagespb.Storage `protobuf:"bytes,6,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mover_micro_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mover_micro_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_mover_micro_proto_rawDescGZIP(), []int{0}
}

func (x *StartRequest) GetGeneralQueue() string {
	if x != nil {
		return x.GeneralQueue
	}
	return ""
}

func (x *StartRequest) GetRequestsQueue() string {
	if x != nil {
		return x.RequestsQueue
	}
	return ""
}

func (x *StartRequest) GetResultsQueue() string {
	if x != nil {
		return x.ResultsQueue
	}
	return ""
}

func (x *StartRequest) GetExecution() *taskpb.Execution {
	if x != nil {
		return x.Execution
	}
	return nil
}

func (x *StartRequest) GetSource() *storagespb.Storage {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *StartRequest) GetDestination() *storagespb.Storage {
	if x != nil {
		return x.Destination
	}
	return nil
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mover_micro_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mover_micro_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_mover_micro_proto_rawDescGZIP(), []int{1}
}

var File_protobuf_mover_micro_proto protoreflect.FileDescriptor

var file_protobuf_mover_micro_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x76, 0x65, 0x72,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6e, 0x6f,
	0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x6d, 0x6f,
	0x76, 0x65, 0x72, 0x1a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x69, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d,
	0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x69, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x75, 0x6d, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x5b, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x6d, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x6d, 0x6f, 0x76, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c,
	0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x75, 0x6d, 0x2d, 0x69, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x76, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_mover_micro_proto_rawDescOnce sync.Once
	file_protobuf_mover_micro_proto_rawDescData = file_protobuf_mover_micro_proto_rawDesc
)

func file_protobuf_mover_micro_proto_rawDescGZIP() []byte {
	file_protobuf_mover_micro_proto_rawDescOnce.Do(func() {
		file_protobuf_mover_micro_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_mover_micro_proto_rawDescData)
	})
	return file_protobuf_mover_micro_proto_rawDescData
}

var file_protobuf_mover_micro_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_mover_micro_proto_goTypes = []interface{}{
	(*StartRequest)(nil),       // 0: nodeum.protobuf.mover.StartRequest
	(*StartResponse)(nil),      // 1: nodeum.protobuf.mover.StartResponse
	(*taskpb.Execution)(nil),   // 2: nodeum.protobuf.Execution
	(*storagespb.Storage)(nil), // 3: nodeum.protobuf.Storage
}
var file_protobuf_mover_micro_proto_depIdxs = []int32{
	2, // 0: nodeum.protobuf.mover.StartRequest.execution:type_name -> nodeum.protobuf.Execution
	3, // 1: nodeum.protobuf.mover.StartRequest.source:type_name -> nodeum.protobuf.Storage
	3, // 2: nodeum.protobuf.mover.StartRequest.destination:type_name -> nodeum.protobuf.Storage
	0, // 3: nodeum.protobuf.mover.Mover.Start:input_type -> nodeum.protobuf.mover.StartRequest
	1, // 4: nodeum.protobuf.mover.Mover.Start:output_type -> nodeum.protobuf.mover.StartResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protobuf_mover_micro_proto_init() }
func file_protobuf_mover_micro_proto_init() {
	if File_protobuf_mover_micro_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_mover_micro_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_protobuf_mover_micro_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
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
			RawDescriptor: file_protobuf_mover_micro_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_mover_micro_proto_goTypes,
		DependencyIndexes: file_protobuf_mover_micro_proto_depIdxs,
		MessageInfos:      file_protobuf_mover_micro_proto_msgTypes,
	}.Build()
	File_protobuf_mover_micro_proto = out.File
	file_protobuf_mover_micro_proto_rawDesc = nil
	file_protobuf_mover_micro_proto_goTypes = nil
	file_protobuf_mover_micro_proto_depIdxs = nil
}
