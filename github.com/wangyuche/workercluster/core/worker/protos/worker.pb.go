// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: core/worker/protos/worker.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WorkerStatus int32

const (
	WorkerStatus_idle WorkerStatus = 0
	WorkerStatus_busy WorkerStatus = 1
)

// Enum value maps for WorkerStatus.
var (
	WorkerStatus_name = map[int32]string{
		0: "idle",
		1: "busy",
	}
	WorkerStatus_value = map[string]int32{
		"idle": 0,
		"busy": 1,
	}
)

func (x WorkerStatus) Enum() *WorkerStatus {
	p := new(WorkerStatus)
	*p = x
	return p
}

func (x WorkerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_core_worker_protos_worker_proto_enumTypes[0].Descriptor()
}

func (WorkerStatus) Type() protoreflect.EnumType {
	return &file_core_worker_protos_worker_proto_enumTypes[0]
}

func (x WorkerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkerStatus.Descriptor instead.
func (WorkerStatus) EnumDescriptor() ([]byte, []int) {
	return file_core_worker_protos_worker_proto_rawDescGZIP(), []int{0}
}

type GetStatusStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status WorkerStatus `protobuf:"varint,1,opt,name=status,proto3,enum=worker.WorkerStatus" json:"status,omitempty"`
}

func (x *GetStatusStruct) Reset() {
	*x = GetStatusStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_worker_protos_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusStruct) ProtoMessage() {}

func (x *GetStatusStruct) ProtoReflect() protoreflect.Message {
	mi := &file_core_worker_protos_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusStruct.ProtoReflect.Descriptor instead.
func (*GetStatusStruct) Descriptor() ([]byte, []int) {
	return file_core_worker_protos_worker_proto_rawDescGZIP(), []int{0}
}

func (x *GetStatusStruct) GetStatus() WorkerStatus {
	if x != nil {
		return x.Status
	}
	return WorkerStatus_idle
}

type PushStatusStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status WorkerStatus `protobuf:"varint,1,opt,name=status,proto3,enum=worker.WorkerStatus" json:"status,omitempty"`
}

func (x *PushStatusStruct) Reset() {
	*x = PushStatusStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_worker_protos_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushStatusStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushStatusStruct) ProtoMessage() {}

func (x *PushStatusStruct) ProtoReflect() protoreflect.Message {
	mi := &file_core_worker_protos_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushStatusStruct.ProtoReflect.Descriptor instead.
func (*PushStatusStruct) Descriptor() ([]byte, []int) {
	return file_core_worker_protos_worker_proto_rawDescGZIP(), []int{1}
}

func (x *PushStatusStruct) GetStatus() WorkerStatus {
	if x != nil {
		return x.Status
	}
	return WorkerStatus_idle
}

var File_core_worker_protos_worker_proto protoreflect.FileDescriptor

var file_core_worker_protos_worker_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x40, 0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x22, 0x0a, 0x0c, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x69, 0x64, 0x6c,
	0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x62, 0x75, 0x73, 0x79, 0x10, 0x01, 0x32, 0x90, 0x01,
	0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x52, 0x50, 0x43, 0x12, 0x3e, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x17, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0a,
	0x50, 0x75, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x61, 0x6e, 0x67, 0x79, 0x75, 0x63, 0x68, 0x65, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_core_worker_protos_worker_proto_rawDescOnce sync.Once
	file_core_worker_protos_worker_proto_rawDescData = file_core_worker_protos_worker_proto_rawDesc
)

func file_core_worker_protos_worker_proto_rawDescGZIP() []byte {
	file_core_worker_protos_worker_proto_rawDescOnce.Do(func() {
		file_core_worker_protos_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_worker_protos_worker_proto_rawDescData)
	})
	return file_core_worker_protos_worker_proto_rawDescData
}

var file_core_worker_protos_worker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_worker_protos_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_core_worker_protos_worker_proto_goTypes = []interface{}{
	(WorkerStatus)(0),        // 0: worker.WorkerStatus
	(*GetStatusStruct)(nil),  // 1: worker.GetStatusStruct
	(*PushStatusStruct)(nil), // 2: worker.PushStatusStruct
	(*emptypb.Empty)(nil),    // 3: google.protobuf.Empty
}
var file_core_worker_protos_worker_proto_depIdxs = []int32{
	0, // 0: worker.GetStatusStruct.status:type_name -> worker.WorkerStatus
	0, // 1: worker.PushStatusStruct.status:type_name -> worker.WorkerStatus
	3, // 2: worker.WorkerGRPC.GetStatus:input_type -> google.protobuf.Empty
	2, // 3: worker.WorkerGRPC.PushStatus:input_type -> worker.PushStatusStruct
	1, // 4: worker.WorkerGRPC.GetStatus:output_type -> worker.GetStatusStruct
	3, // 5: worker.WorkerGRPC.PushStatus:output_type -> google.protobuf.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_core_worker_protos_worker_proto_init() }
func file_core_worker_protos_worker_proto_init() {
	if File_core_worker_protos_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_worker_protos_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusStruct); i {
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
		file_core_worker_protos_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushStatusStruct); i {
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
			RawDescriptor: file_core_worker_protos_worker_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_worker_protos_worker_proto_goTypes,
		DependencyIndexes: file_core_worker_protos_worker_proto_depIdxs,
		EnumInfos:         file_core_worker_protos_worker_proto_enumTypes,
		MessageInfos:      file_core_worker_protos_worker_proto_msgTypes,
	}.Build()
	File_core_worker_protos_worker_proto = out.File
	file_core_worker_protos_worker_proto_rawDesc = nil
	file_core_worker_protos_worker_proto_goTypes = nil
	file_core_worker_protos_worker_proto_depIdxs = nil
}