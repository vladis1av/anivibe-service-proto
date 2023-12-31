// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: proto/common/filter/v1/filter.proto

package pb_filter

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

type StringFieldFilter_Operator int32

const (
	StringFieldFilter_OPERATOR_UNSPECIFIED StringFieldFilter_Operator = 0
	StringFieldFilter_OPERATOR_EQ          StringFieldFilter_Operator = 1
	StringFieldFilter_OPERATOR_NEQ         StringFieldFilter_Operator = 2
	StringFieldFilter_OPERATOR_LIKE        StringFieldFilter_Operator = 3
)

// Enum value maps for StringFieldFilter_Operator.
var (
	StringFieldFilter_Operator_name = map[int32]string{
		0: "OPERATOR_UNSPECIFIED",
		1: "OPERATOR_EQ",
		2: "OPERATOR_NEQ",
		3: "OPERATOR_LIKE",
	}
	StringFieldFilter_Operator_value = map[string]int32{
		"OPERATOR_UNSPECIFIED": 0,
		"OPERATOR_EQ":          1,
		"OPERATOR_NEQ":         2,
		"OPERATOR_LIKE":        3,
	}
)

func (x StringFieldFilter_Operator) Enum() *StringFieldFilter_Operator {
	p := new(StringFieldFilter_Operator)
	*p = x
	return p
}

func (x StringFieldFilter_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StringFieldFilter_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_common_filter_v1_filter_proto_enumTypes[0].Descriptor()
}

func (StringFieldFilter_Operator) Type() protoreflect.EnumType {
	return &file_proto_common_filter_v1_filter_proto_enumTypes[0]
}

func (x StringFieldFilter_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StringFieldFilter_Operator.Descriptor instead.
func (StringFieldFilter_Operator) EnumDescriptor() ([]byte, []int) {
	return file_proto_common_filter_v1_filter_proto_rawDescGZIP(), []int{1, 0}
}

type NumberFieldFilter_Operator int32

const (
	NumberFieldFilter_OPERATOR_UNSPECIFIED NumberFieldFilter_Operator = 0
	NumberFieldFilter_OPERATOR_EQ          NumberFieldFilter_Operator = 1
	NumberFieldFilter_OPERATOR_NEQ         NumberFieldFilter_Operator = 3
	NumberFieldFilter_OPERATOR_LT          NumberFieldFilter_Operator = 4
	NumberFieldFilter_OPERATOR_LTE         NumberFieldFilter_Operator = 5
	NumberFieldFilter_OPERATOR_GT          NumberFieldFilter_Operator = 6
	NumberFieldFilter_OPERATOR_GTE         NumberFieldFilter_Operator = 7
)

// Enum value maps for NumberFieldFilter_Operator.
var (
	NumberFieldFilter_Operator_name = map[int32]string{
		0: "OPERATOR_UNSPECIFIED",
		1: "OPERATOR_EQ",
		3: "OPERATOR_NEQ",
		4: "OPERATOR_LT",
		5: "OPERATOR_LTE",
		6: "OPERATOR_GT",
		7: "OPERATOR_GTE",
	}
	NumberFieldFilter_Operator_value = map[string]int32{
		"OPERATOR_UNSPECIFIED": 0,
		"OPERATOR_EQ":          1,
		"OPERATOR_NEQ":         3,
		"OPERATOR_LT":          4,
		"OPERATOR_LTE":         5,
		"OPERATOR_GT":          6,
		"OPERATOR_GTE":         7,
	}
)

func (x NumberFieldFilter_Operator) Enum() *NumberFieldFilter_Operator {
	p := new(NumberFieldFilter_Operator)
	*p = x
	return p
}

func (x NumberFieldFilter_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NumberFieldFilter_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_common_filter_v1_filter_proto_enumTypes[1].Descriptor()
}

func (NumberFieldFilter_Operator) Type() protoreflect.EnumType {
	return &file_proto_common_filter_v1_filter_proto_enumTypes[1]
}

func (x NumberFieldFilter_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NumberFieldFilter_Operator.Descriptor instead.
func (NumberFieldFilter_Operator) EnumDescriptor() ([]byte, []int) {
	return file_proto_common_filter_v1_filter_proto_rawDescGZIP(), []int{2, 0}
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_filter_v1_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_filter_v1_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_proto_common_filter_v1_filter_proto_rawDescGZIP(), []int{0}
}

func (x *Pagination) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Filter for string values, example: ?email.op=eq&email.val=me@example.com
type StringFieldFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op  StringFieldFilter_Operator `protobuf:"varint,1,opt,name=op,proto3,enum=common.filter.v1.StringFieldFilter_Operator" json:"op,omitempty"`
	Val string                     `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *StringFieldFilter) Reset() {
	*x = StringFieldFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_filter_v1_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringFieldFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringFieldFilter) ProtoMessage() {}

func (x *StringFieldFilter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_filter_v1_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringFieldFilter.ProtoReflect.Descriptor instead.
func (*StringFieldFilter) Descriptor() ([]byte, []int) {
	return file_proto_common_filter_v1_filter_proto_rawDescGZIP(), []int{1}
}

func (x *StringFieldFilter) GetOp() StringFieldFilter_Operator {
	if x != nil {
		return x.Op
	}
	return StringFieldFilter_OPERATOR_UNSPECIFIED
}

func (x *StringFieldFilter) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

// Filter for int values, example: ?age.op=gt&age.val=18
type NumberFieldFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op  NumberFieldFilter_Operator `protobuf:"varint,1,opt,name=op,proto3,enum=common.filter.v1.NumberFieldFilter_Operator" json:"op,omitempty"`
	Val string                     `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *NumberFieldFilter) Reset() {
	*x = NumberFieldFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_filter_v1_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberFieldFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberFieldFilter) ProtoMessage() {}

func (x *NumberFieldFilter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_filter_v1_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberFieldFilter.ProtoReflect.Descriptor instead.
func (*NumberFieldFilter) Descriptor() ([]byte, []int) {
	return file_proto_common_filter_v1_filter_proto_rawDescGZIP(), []int{2}
}

func (x *NumberFieldFilter) GetOp() NumberFieldFilter_Operator {
	if x != nil {
		return x.Op
	}
	return NumberFieldFilter_OPERATOR_UNSPECIFIED
}

func (x *NumberFieldFilter) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

var File_proto_common_filter_v1_filter_proto protoreflect.FileDescriptor

var file_proto_common_filter_v1_filter_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x3a, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0xbf, 0x01, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x02, 0x6f, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x5a, 0x0a, 0x08, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f,
	0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x45, 0x51, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x4e, 0x45, 0x51,
	0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x4c,
	0x49, 0x4b, 0x45, 0x10, 0x03, 0x22, 0xf3, 0x01, 0x0a, 0x11, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x02, 0x6f,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x8d, 0x01, 0x0a, 0x08,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x50, 0x45, 0x52,
	0x41, 0x54, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x45,
	0x51, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f,
	0x4e, 0x45, 0x51, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f,
	0x52, 0x5f, 0x4c, 0x54, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x4f, 0x52, 0x5f, 0x4c, 0x54, 0x45, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52,
	0x41, 0x54, 0x4f, 0x52, 0x5f, 0x47, 0x54, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x47, 0x54, 0x45, 0x10, 0x07, 0x42, 0x4e, 0x5a, 0x4c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6c, 0x61, 0x64, 0x69, 0x73,
	0x31, 0x61, 0x76, 0x2f, 0x61, 0x6e, 0x69, 0x76, 0x69, 0x62, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x62, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_common_filter_v1_filter_proto_rawDescOnce sync.Once
	file_proto_common_filter_v1_filter_proto_rawDescData = file_proto_common_filter_v1_filter_proto_rawDesc
)

func file_proto_common_filter_v1_filter_proto_rawDescGZIP() []byte {
	file_proto_common_filter_v1_filter_proto_rawDescOnce.Do(func() {
		file_proto_common_filter_v1_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_common_filter_v1_filter_proto_rawDescData)
	})
	return file_proto_common_filter_v1_filter_proto_rawDescData
}

var file_proto_common_filter_v1_filter_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_common_filter_v1_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_common_filter_v1_filter_proto_goTypes = []interface{}{
	(StringFieldFilter_Operator)(0), // 0: common.filter.v1.StringFieldFilter.Operator
	(NumberFieldFilter_Operator)(0), // 1: common.filter.v1.NumberFieldFilter.Operator
	(*Pagination)(nil),              // 2: common.filter.v1.Pagination
	(*StringFieldFilter)(nil),       // 3: common.filter.v1.StringFieldFilter
	(*NumberFieldFilter)(nil),       // 4: common.filter.v1.NumberFieldFilter
}
var file_proto_common_filter_v1_filter_proto_depIdxs = []int32{
	0, // 0: common.filter.v1.StringFieldFilter.op:type_name -> common.filter.v1.StringFieldFilter.Operator
	1, // 1: common.filter.v1.NumberFieldFilter.op:type_name -> common.filter.v1.NumberFieldFilter.Operator
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_common_filter_v1_filter_proto_init() }
func file_proto_common_filter_v1_filter_proto_init() {
	if File_proto_common_filter_v1_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_common_filter_v1_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_proto_common_filter_v1_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringFieldFilter); i {
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
		file_proto_common_filter_v1_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumberFieldFilter); i {
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
			RawDescriptor: file_proto_common_filter_v1_filter_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_common_filter_v1_filter_proto_goTypes,
		DependencyIndexes: file_proto_common_filter_v1_filter_proto_depIdxs,
		EnumInfos:         file_proto_common_filter_v1_filter_proto_enumTypes,
		MessageInfos:      file_proto_common_filter_v1_filter_proto_msgTypes,
	}.Build()
	File_proto_common_filter_v1_filter_proto = out.File
	file_proto_common_filter_v1_filter_proto_rawDesc = nil
	file_proto_common_filter_v1_filter_proto_goTypes = nil
	file_proto_common_filter_v1_filter_proto_depIdxs = nil
}
