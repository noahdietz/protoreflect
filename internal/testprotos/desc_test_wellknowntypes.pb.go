// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.7
// source: desc_test_wellknowntypes.proto

package testprotos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestWellKnownTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime *timestamppb.Timestamp  `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Elapsed   *durationpb.Duration    `protobuf:"bytes,2,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	Dbl       *wrapperspb.DoubleValue `protobuf:"bytes,3,opt,name=dbl,proto3" json:"dbl,omitempty"`
	Flt       *wrapperspb.FloatValue  `protobuf:"bytes,4,opt,name=flt,proto3" json:"flt,omitempty"`
	Bl        *wrapperspb.BoolValue   `protobuf:"bytes,5,opt,name=bl,proto3" json:"bl,omitempty"`
	I32       *wrapperspb.Int32Value  `protobuf:"bytes,6,opt,name=i32,proto3" json:"i32,omitempty"`
	I64       *wrapperspb.Int64Value  `protobuf:"bytes,7,opt,name=i64,proto3" json:"i64,omitempty"`
	U32       *wrapperspb.UInt32Value `protobuf:"bytes,8,opt,name=u32,proto3" json:"u32,omitempty"`
	U64       *wrapperspb.UInt64Value `protobuf:"bytes,9,opt,name=u64,proto3" json:"u64,omitempty"`
	Str       *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=str,proto3" json:"str,omitempty"`
	Byt       *wrapperspb.BytesValue  `protobuf:"bytes,11,opt,name=byt,proto3" json:"byt,omitempty"`
	Json      []*structpb.Value       `protobuf:"bytes,12,rep,name=json,proto3" json:"json,omitempty"`
	Extras    []*anypb.Any            `protobuf:"bytes,13,rep,name=extras,proto3" json:"extras,omitempty"`
}

func (x *TestWellKnownTypes) Reset() {
	*x = TestWellKnownTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_test_wellknowntypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestWellKnownTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestWellKnownTypes) ProtoMessage() {}

func (x *TestWellKnownTypes) ProtoReflect() protoreflect.Message {
	mi := &file_desc_test_wellknowntypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestWellKnownTypes.ProtoReflect.Descriptor instead.
func (*TestWellKnownTypes) Descriptor() ([]byte, []int) {
	return file_desc_test_wellknowntypes_proto_rawDescGZIP(), []int{0}
}

func (x *TestWellKnownTypes) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *TestWellKnownTypes) GetElapsed() *durationpb.Duration {
	if x != nil {
		return x.Elapsed
	}
	return nil
}

func (x *TestWellKnownTypes) GetDbl() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Dbl
	}
	return nil
}

func (x *TestWellKnownTypes) GetFlt() *wrapperspb.FloatValue {
	if x != nil {
		return x.Flt
	}
	return nil
}

func (x *TestWellKnownTypes) GetBl() *wrapperspb.BoolValue {
	if x != nil {
		return x.Bl
	}
	return nil
}

func (x *TestWellKnownTypes) GetI32() *wrapperspb.Int32Value {
	if x != nil {
		return x.I32
	}
	return nil
}

func (x *TestWellKnownTypes) GetI64() *wrapperspb.Int64Value {
	if x != nil {
		return x.I64
	}
	return nil
}

func (x *TestWellKnownTypes) GetU32() *wrapperspb.UInt32Value {
	if x != nil {
		return x.U32
	}
	return nil
}

func (x *TestWellKnownTypes) GetU64() *wrapperspb.UInt64Value {
	if x != nil {
		return x.U64
	}
	return nil
}

func (x *TestWellKnownTypes) GetStr() *wrapperspb.StringValue {
	if x != nil {
		return x.Str
	}
	return nil
}

func (x *TestWellKnownTypes) GetByt() *wrapperspb.BytesValue {
	if x != nil {
		return x.Byt
	}
	return nil
}

func (x *TestWellKnownTypes) GetJson() []*structpb.Value {
	if x != nil {
		return x.Json
	}
	return nil
}

func (x *TestWellKnownTypes) GetExtras() []*anypb.Any {
	if x != nil {
		return x.Extras
	}
	return nil
}

var File_desc_test_wellknowntypes_proto protoreflect.FileDescriptor

var file_desc_test_wellknowntypes_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x77, 0x65, 0x6c, 0x6c,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x05, 0x0a, 0x12, 0x54, 0x65, 0x73, 0x74, 0x57,
	0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x65, 0x6c, 0x61, 0x70,
	0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x12, 0x2e, 0x0a,
	0x03, 0x64, 0x62, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x64, 0x62, 0x6c, 0x12, 0x2d, 0x0a,
	0x03, 0x66, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x66, 0x6c, 0x74, 0x12, 0x2a, 0x0a, 0x02,
	0x62, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x62, 0x6c, 0x12, 0x2d, 0x0a, 0x03, 0x69, 0x33, 0x32, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x03, 0x69, 0x33, 0x32, 0x12, 0x2d, 0x0a, 0x03, 0x69, 0x36, 0x34, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x03, 0x69, 0x36, 0x34, 0x12, 0x2e, 0x0a, 0x03, 0x75, 0x33, 0x32, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x03, 0x75, 0x33, 0x32, 0x12, 0x2e, 0x0a, 0x03, 0x75, 0x36, 0x34, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x03, 0x75, 0x36, 0x34, 0x12, 0x2e, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x2d, 0x0a, 0x03, 0x62, 0x79, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x03, 0x62, 0x79, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6a, 0x73, 0x6f,
	0x6e, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x78, 0x74, 0x72, 0x61, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x65, 0x78, 0x74, 0x72, 0x61, 0x73, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x68,
	0x75, 0x6d, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_desc_test_wellknowntypes_proto_rawDescOnce sync.Once
	file_desc_test_wellknowntypes_proto_rawDescData = file_desc_test_wellknowntypes_proto_rawDesc
)

func file_desc_test_wellknowntypes_proto_rawDescGZIP() []byte {
	file_desc_test_wellknowntypes_proto_rawDescOnce.Do(func() {
		file_desc_test_wellknowntypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_desc_test_wellknowntypes_proto_rawDescData)
	})
	return file_desc_test_wellknowntypes_proto_rawDescData
}

var file_desc_test_wellknowntypes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_desc_test_wellknowntypes_proto_goTypes = []interface{}{
	(*TestWellKnownTypes)(nil),     // 0: testprotos.TestWellKnownTypes
	(*timestamppb.Timestamp)(nil),  // 1: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),    // 2: google.protobuf.Duration
	(*wrapperspb.DoubleValue)(nil), // 3: google.protobuf.DoubleValue
	(*wrapperspb.FloatValue)(nil),  // 4: google.protobuf.FloatValue
	(*wrapperspb.BoolValue)(nil),   // 5: google.protobuf.BoolValue
	(*wrapperspb.Int32Value)(nil),  // 6: google.protobuf.Int32Value
	(*wrapperspb.Int64Value)(nil),  // 7: google.protobuf.Int64Value
	(*wrapperspb.UInt32Value)(nil), // 8: google.protobuf.UInt32Value
	(*wrapperspb.UInt64Value)(nil), // 9: google.protobuf.UInt64Value
	(*wrapperspb.StringValue)(nil), // 10: google.protobuf.StringValue
	(*wrapperspb.BytesValue)(nil),  // 11: google.protobuf.BytesValue
	(*structpb.Value)(nil),         // 12: google.protobuf.Value
	(*anypb.Any)(nil),              // 13: google.protobuf.Any
}
var file_desc_test_wellknowntypes_proto_depIdxs = []int32{
	1,  // 0: testprotos.TestWellKnownTypes.start_time:type_name -> google.protobuf.Timestamp
	2,  // 1: testprotos.TestWellKnownTypes.elapsed:type_name -> google.protobuf.Duration
	3,  // 2: testprotos.TestWellKnownTypes.dbl:type_name -> google.protobuf.DoubleValue
	4,  // 3: testprotos.TestWellKnownTypes.flt:type_name -> google.protobuf.FloatValue
	5,  // 4: testprotos.TestWellKnownTypes.bl:type_name -> google.protobuf.BoolValue
	6,  // 5: testprotos.TestWellKnownTypes.i32:type_name -> google.protobuf.Int32Value
	7,  // 6: testprotos.TestWellKnownTypes.i64:type_name -> google.protobuf.Int64Value
	8,  // 7: testprotos.TestWellKnownTypes.u32:type_name -> google.protobuf.UInt32Value
	9,  // 8: testprotos.TestWellKnownTypes.u64:type_name -> google.protobuf.UInt64Value
	10, // 9: testprotos.TestWellKnownTypes.str:type_name -> google.protobuf.StringValue
	11, // 10: testprotos.TestWellKnownTypes.byt:type_name -> google.protobuf.BytesValue
	12, // 11: testprotos.TestWellKnownTypes.json:type_name -> google.protobuf.Value
	13, // 12: testprotos.TestWellKnownTypes.extras:type_name -> google.protobuf.Any
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_desc_test_wellknowntypes_proto_init() }
func file_desc_test_wellknowntypes_proto_init() {
	if File_desc_test_wellknowntypes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_desc_test_wellknowntypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestWellKnownTypes); i {
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
			RawDescriptor: file_desc_test_wellknowntypes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_desc_test_wellknowntypes_proto_goTypes,
		DependencyIndexes: file_desc_test_wellknowntypes_proto_depIdxs,
		MessageInfos:      file_desc_test_wellknowntypes_proto_msgTypes,
	}.Build()
	File_desc_test_wellknowntypes_proto = out.File
	file_desc_test_wellknowntypes_proto_rawDesc = nil
	file_desc_test_wellknowntypes_proto_goTypes = nil
	file_desc_test_wellknowntypes_proto_depIdxs = nil
}
