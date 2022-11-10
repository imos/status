// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: stacktrace.proto

package stpb

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

// Stack trace.
type StackTrace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stack frames of the stack trace.
	Frames []*StackTrace_Frame `protobuf:"bytes,1,rep,name=frames,proto3" json:"frames,omitempty"`
}

func (x *StackTrace) Reset() {
	*x = StackTrace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stacktrace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackTrace) ProtoMessage() {}

func (x *StackTrace) ProtoReflect() protoreflect.Message {
	mi := &file_stacktrace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackTrace.ProtoReflect.Descriptor instead.
func (*StackTrace) Descriptor() ([]byte, []int) {
	return file_stacktrace_proto_rawDescGZIP(), []int{0}
}

func (x *StackTrace) GetFrames() []*StackTrace_Frame {
	if x != nil {
		return x.Frames
	}
	return nil
}

// Frame represents a call frame.
type StackTrace_Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// File path.
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// Line number.
	Line int32 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	// Function name.
	Function string `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
	// The entry address of the function.
	Entry uint64 `protobuf:"varint,4,opt,name=entry,proto3" json:"entry,omitempty"`
	// Program counter of the call frame.
	ProgramCounter uint64 `protobuf:"varint,5,opt,name=program_counter,json=programCounter,proto3" json:"program_counter,omitempty"`
}

func (x *StackTrace_Frame) Reset() {
	*x = StackTrace_Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stacktrace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackTrace_Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackTrace_Frame) ProtoMessage() {}

func (x *StackTrace_Frame) ProtoReflect() protoreflect.Message {
	mi := &file_stacktrace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackTrace_Frame.ProtoReflect.Descriptor instead.
func (*StackTrace_Frame) Descriptor() ([]byte, []int) {
	return file_stacktrace_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StackTrace_Frame) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *StackTrace_Frame) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *StackTrace_Frame) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

func (x *StackTrace_Frame) GetEntry() uint64 {
	if x != nil {
		return x.Entry
	}
	return 0
}

func (x *StackTrace_Frame) GetProgramCounter() uint64 {
	if x != nil {
		return x.ProgramCounter
	}
	return 0
}

var File_stacktrace_proto protoreflect.FileDescriptor

var file_stacktrace_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x52, 0x06, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x8a, 0x01, 0x0a, 0x05,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x73, 0x74,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stacktrace_proto_rawDescOnce sync.Once
	file_stacktrace_proto_rawDescData = file_stacktrace_proto_rawDesc
)

func file_stacktrace_proto_rawDescGZIP() []byte {
	file_stacktrace_proto_rawDescOnce.Do(func() {
		file_stacktrace_proto_rawDescData = protoimpl.X.CompressGZIP(file_stacktrace_proto_rawDescData)
	})
	return file_stacktrace_proto_rawDescData
}

var file_stacktrace_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stacktrace_proto_goTypes = []interface{}{
	(*StackTrace)(nil),       // 0: go_status.status.proto.StackTrace
	(*StackTrace_Frame)(nil), // 1: go_status.status.proto.StackTrace.Frame
}
var file_stacktrace_proto_depIdxs = []int32{
	1, // 0: go_status.status.proto.StackTrace.frames:type_name -> go_status.status.proto.StackTrace.Frame
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_stacktrace_proto_init() }
func file_stacktrace_proto_init() {
	if File_stacktrace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stacktrace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackTrace); i {
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
		file_stacktrace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackTrace_Frame); i {
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
			RawDescriptor: file_stacktrace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stacktrace_proto_goTypes,
		DependencyIndexes: file_stacktrace_proto_depIdxs,
		MessageInfos:      file_stacktrace_proto_msgTypes,
	}.Build()
	File_stacktrace_proto = out.File
	file_stacktrace_proto_rawDesc = nil
	file_stacktrace_proto_goTypes = nil
	file_stacktrace_proto_depIdxs = nil
}