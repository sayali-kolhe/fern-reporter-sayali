// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.0
// source: gettestrunall.proto

package gettestrunall

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

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	mi := &file_gettestrunall_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gettestrunall_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_gettestrunall_proto_rawDescGZIP(), []int{0}
}

type TestRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"` // Add other fields based on your `TestRun` model.
}

func (x *TestRun) Reset() {
	*x = TestRun{}
	mi := &file_gettestrunall_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRun) ProtoMessage() {}

func (x *TestRun) ProtoReflect() protoreflect.Message {
	mi := &file_gettestrunall_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRun.ProtoReflect.Descriptor instead.
func (*TestRun) Descriptor() ([]byte, []int) {
	return file_gettestrunall_proto_rawDescGZIP(), []int{1}
}

func (x *TestRun) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestRun) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestRun) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type TestRunList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRuns []*TestRun `protobuf:"bytes,1,rep,name=test_runs,json=testRuns,proto3" json:"test_runs,omitempty"`
}

func (x *TestRunList) Reset() {
	*x = TestRunList{}
	mi := &file_gettestrunall_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestRunList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRunList) ProtoMessage() {}

func (x *TestRunList) ProtoReflect() protoreflect.Message {
	mi := &file_gettestrunall_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRunList.ProtoReflect.Descriptor instead.
func (*TestRunList) Descriptor() ([]byte, []int) {
	return file_gettestrunall_proto_rawDescGZIP(), []int{2}
}

func (x *TestRunList) GetTestRuns() []*TestRun {
	if x != nil {
		return x.TestRuns
	}
	return nil
}

var File_gettestrunall_proto protoreflect.FileDescriptor

var file_gettestrunall_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x65, 0x74, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x61, 0x6c, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x65, 0x74, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75,
	0x6e, 0x61, 0x6c, 0x6c, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x42, 0x0a, 0x0b, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x65, 0x74, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x61, 0x6c, 0x6c, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x73, 0x32,
	0x59, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x41, 0x6c, 0x6c, 0x12,
	0x1b, 0x2e, 0x67, 0x65, 0x74, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x61, 0x6c, 0x6c, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67,
	0x65, 0x74, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x61, 0x6c, 0x6c, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f,
	0x3b, 0x67, 0x65, 0x74, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x61, 0x6c, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gettestrunall_proto_rawDescOnce sync.Once
	file_gettestrunall_proto_rawDescData = file_gettestrunall_proto_rawDesc
)

func file_gettestrunall_proto_rawDescGZIP() []byte {
	file_gettestrunall_proto_rawDescOnce.Do(func() {
		file_gettestrunall_proto_rawDescData = protoimpl.X.CompressGZIP(file_gettestrunall_proto_rawDescData)
	})
	return file_gettestrunall_proto_rawDescData
}

var file_gettestrunall_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gettestrunall_proto_goTypes = []any{
	(*EmptyRequest)(nil), // 0: gettestrunall.EmptyRequest
	(*TestRun)(nil),      // 1: gettestrunall.TestRun
	(*TestRunList)(nil),  // 2: gettestrunall.TestRunList
}
var file_gettestrunall_proto_depIdxs = []int32{
	1, // 0: gettestrunall.TestRunList.test_runs:type_name -> gettestrunall.TestRun
	0, // 1: gettestrunall.TestService.GetTestRunAll:input_type -> gettestrunall.EmptyRequest
	2, // 2: gettestrunall.TestService.GetTestRunAll:output_type -> gettestrunall.TestRunList
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gettestrunall_proto_init() }
func file_gettestrunall_proto_init() {
	if File_gettestrunall_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gettestrunall_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gettestrunall_proto_goTypes,
		DependencyIndexes: file_gettestrunall_proto_depIdxs,
		MessageInfos:      file_gettestrunall_proto_msgTypes,
	}.Build()
	File_gettestrunall_proto = out.File
	file_gettestrunall_proto_rawDesc = nil
	file_gettestrunall_proto_goTypes = nil
	file_gettestrunall_proto_depIdxs = nil
}
