// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: com/coralogixapis/dashboards/v1/common/ordering_field.proto

package golang

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type OrderingField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field          *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	OrderDirection OrderDirection          `protobuf:"varint,2,opt,name=order_direction,json=orderDirection,proto3,enum=com.coralogixapis.dashboards.v1.common.OrderDirection" json:"order_direction,omitempty"`
}

func (x *OrderingField) Reset() {
	*x = OrderingField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderingField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderingField) ProtoMessage() {}

func (x *OrderingField) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderingField.ProtoReflect.Descriptor instead.
func (*OrderingField) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_rawDescGZIP(), []int{0}
}

func (x *OrderingField) GetField() *wrapperspb.StringValue {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *OrderingField) GetOrderDirection() OrderDirection {
	if x != nil {
		return x.OrderDirection
	}
	return OrderDirection_ORDER_DIRECTION_UNSPECIFIED
}

var File_com_coralogixapis_dashboards_v1_common_ordering_field_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x3c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x5f, 0x0a, 0x0f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_rawDescData = file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_goTypes = []interface{}{
	(*OrderingField)(nil),          // 0: com.coralogixapis.dashboards.v1.common.OrderingField
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(OrderDirection)(0),            // 2: com.coralogixapis.dashboards.v1.common.OrderDirection
}
var file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.dashboards.v1.common.OrderingField.field:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogixapis.dashboards.v1.common.OrderingField.order_direction:type_name -> com.coralogixapis.dashboards.v1.common.OrderDirection
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_init() }
func file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_init() {
	if File_com_coralogixapis_dashboards_v1_common_ordering_field_proto != nil {
		return
	}
	file_com_coralogixapis_dashboards_v1_common_order_direction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderingField); i {
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
			RawDescriptor: file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_ordering_field_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_ordering_field_proto_depIdxs = nil
}
