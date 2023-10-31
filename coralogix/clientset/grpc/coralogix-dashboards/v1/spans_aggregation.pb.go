// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogixapis/dashboards/v1/common/spans_aggregation.proto

package __

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SpansAggregation_MetricAggregation_MetricField int32

const (
	SpansAggregation_MetricAggregation_METRIC_FIELD_UNSPECIFIED SpansAggregation_MetricAggregation_MetricField = 0
	SpansAggregation_MetricAggregation_METRIC_FIELD_DURATION    SpansAggregation_MetricAggregation_MetricField = 1
)

// Enum value maps for SpansAggregation_MetricAggregation_MetricField.
var (
	SpansAggregation_MetricAggregation_MetricField_name = map[int32]string{
		0: "METRIC_FIELD_UNSPECIFIED",
		1: "METRIC_FIELD_DURATION",
	}
	SpansAggregation_MetricAggregation_MetricField_value = map[string]int32{
		"METRIC_FIELD_UNSPECIFIED": 0,
		"METRIC_FIELD_DURATION":    1,
	}
)

func (x SpansAggregation_MetricAggregation_MetricField) Enum() *SpansAggregation_MetricAggregation_MetricField {
	p := new(SpansAggregation_MetricAggregation_MetricField)
	*p = x
	return p
}

func (x SpansAggregation_MetricAggregation_MetricField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpansAggregation_MetricAggregation_MetricField) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_enumTypes[0].Descriptor()
}

func (SpansAggregation_MetricAggregation_MetricField) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_enumTypes[0]
}

func (x SpansAggregation_MetricAggregation_MetricField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpansAggregation_MetricAggregation_MetricField.Descriptor instead.
func (SpansAggregation_MetricAggregation_MetricField) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescGZIP(), []int{0, 0, 0}
}

type SpansAggregation_MetricAggregation_MetricAggregationType int32

const (
	SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_UNSPECIFIED   SpansAggregation_MetricAggregation_MetricAggregationType = 0
	SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_MIN           SpansAggregation_MetricAggregation_MetricAggregationType = 1
	SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_MAX           SpansAggregation_MetricAggregation_MetricAggregationType = 2
	SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_AVERAGE       SpansAggregation_MetricAggregation_MetricAggregationType = 3
	SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_SUM           SpansAggregation_MetricAggregation_MetricAggregationType = 4
	SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_PERCENTILE_99 SpansAggregation_MetricAggregation_MetricAggregationType = 5
	SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_PERCENTILE_95 SpansAggregation_MetricAggregation_MetricAggregationType = 6
	SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_PERCENTILE_50 SpansAggregation_MetricAggregation_MetricAggregationType = 7
)

// Enum value maps for SpansAggregation_MetricAggregation_MetricAggregationType.
var (
	SpansAggregation_MetricAggregation_MetricAggregationType_name = map[int32]string{
		0: "METRIC_AGGREGATION_TYPE_UNSPECIFIED",
		1: "METRIC_AGGREGATION_TYPE_MIN",
		2: "METRIC_AGGREGATION_TYPE_MAX",
		3: "METRIC_AGGREGATION_TYPE_AVERAGE",
		4: "METRIC_AGGREGATION_TYPE_SUM",
		5: "METRIC_AGGREGATION_TYPE_PERCENTILE_99",
		6: "METRIC_AGGREGATION_TYPE_PERCENTILE_95",
		7: "METRIC_AGGREGATION_TYPE_PERCENTILE_50",
	}
	SpansAggregation_MetricAggregation_MetricAggregationType_value = map[string]int32{
		"METRIC_AGGREGATION_TYPE_UNSPECIFIED":   0,
		"METRIC_AGGREGATION_TYPE_MIN":           1,
		"METRIC_AGGREGATION_TYPE_MAX":           2,
		"METRIC_AGGREGATION_TYPE_AVERAGE":       3,
		"METRIC_AGGREGATION_TYPE_SUM":           4,
		"METRIC_AGGREGATION_TYPE_PERCENTILE_99": 5,
		"METRIC_AGGREGATION_TYPE_PERCENTILE_95": 6,
		"METRIC_AGGREGATION_TYPE_PERCENTILE_50": 7,
	}
)

func (x SpansAggregation_MetricAggregation_MetricAggregationType) Enum() *SpansAggregation_MetricAggregation_MetricAggregationType {
	p := new(SpansAggregation_MetricAggregation_MetricAggregationType)
	*p = x
	return p
}

func (x SpansAggregation_MetricAggregation_MetricAggregationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpansAggregation_MetricAggregation_MetricAggregationType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_enumTypes[1].Descriptor()
}

func (SpansAggregation_MetricAggregation_MetricAggregationType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_enumTypes[1]
}

func (x SpansAggregation_MetricAggregation_MetricAggregationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpansAggregation_MetricAggregation_MetricAggregationType.Descriptor instead.
func (SpansAggregation_MetricAggregation_MetricAggregationType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescGZIP(), []int{0, 0, 1}
}

type SpansAggregation_DimensionAggregation_DimensionField int32

const (
	SpansAggregation_DimensionAggregation_DIMENSION_FIELD_UNSPECIFIED SpansAggregation_DimensionAggregation_DimensionField = 0
	SpansAggregation_DimensionAggregation_DIMENSION_FIELD_TRACE_ID    SpansAggregation_DimensionAggregation_DimensionField = 1
)

// Enum value maps for SpansAggregation_DimensionAggregation_DimensionField.
var (
	SpansAggregation_DimensionAggregation_DimensionField_name = map[int32]string{
		0: "DIMENSION_FIELD_UNSPECIFIED",
		1: "DIMENSION_FIELD_TRACE_ID",
	}
	SpansAggregation_DimensionAggregation_DimensionField_value = map[string]int32{
		"DIMENSION_FIELD_UNSPECIFIED": 0,
		"DIMENSION_FIELD_TRACE_ID":    1,
	}
)

func (x SpansAggregation_DimensionAggregation_DimensionField) Enum() *SpansAggregation_DimensionAggregation_DimensionField {
	p := new(SpansAggregation_DimensionAggregation_DimensionField)
	*p = x
	return p
}

func (x SpansAggregation_DimensionAggregation_DimensionField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpansAggregation_DimensionAggregation_DimensionField) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_enumTypes[2].Descriptor()
}

func (SpansAggregation_DimensionAggregation_DimensionField) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_enumTypes[2]
}

func (x SpansAggregation_DimensionAggregation_DimensionField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpansAggregation_DimensionAggregation_DimensionField.Descriptor instead.
func (SpansAggregation_DimensionAggregation_DimensionField) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescGZIP(), []int{0, 1, 0}
}

type SpansAggregation_DimensionAggregation_DimensionAggregationType int32

const (
	SpansAggregation_DimensionAggregation_DIMENSION_AGGREGATION_TYPE_UNSPECIFIED  SpansAggregation_DimensionAggregation_DimensionAggregationType = 0
	SpansAggregation_DimensionAggregation_DIMENSION_AGGREGATION_TYPE_UNIQUE_COUNT SpansAggregation_DimensionAggregation_DimensionAggregationType = 1
	SpansAggregation_DimensionAggregation_DIMENSION_AGGREGATION_TYPE_ERROR_COUNT  SpansAggregation_DimensionAggregation_DimensionAggregationType = 2
)

// Enum value maps for SpansAggregation_DimensionAggregation_DimensionAggregationType.
var (
	SpansAggregation_DimensionAggregation_DimensionAggregationType_name = map[int32]string{
		0: "DIMENSION_AGGREGATION_TYPE_UNSPECIFIED",
		1: "DIMENSION_AGGREGATION_TYPE_UNIQUE_COUNT",
		2: "DIMENSION_AGGREGATION_TYPE_ERROR_COUNT",
	}
	SpansAggregation_DimensionAggregation_DimensionAggregationType_value = map[string]int32{
		"DIMENSION_AGGREGATION_TYPE_UNSPECIFIED":  0,
		"DIMENSION_AGGREGATION_TYPE_UNIQUE_COUNT": 1,
		"DIMENSION_AGGREGATION_TYPE_ERROR_COUNT":  2,
	}
)

func (x SpansAggregation_DimensionAggregation_DimensionAggregationType) Enum() *SpansAggregation_DimensionAggregation_DimensionAggregationType {
	p := new(SpansAggregation_DimensionAggregation_DimensionAggregationType)
	*p = x
	return p
}

func (x SpansAggregation_DimensionAggregation_DimensionAggregationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpansAggregation_DimensionAggregation_DimensionAggregationType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_enumTypes[3].Descriptor()
}

func (SpansAggregation_DimensionAggregation_DimensionAggregationType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_enumTypes[3]
}

func (x SpansAggregation_DimensionAggregation_DimensionAggregationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpansAggregation_DimensionAggregation_DimensionAggregationType.Descriptor instead.
func (SpansAggregation_DimensionAggregation_DimensionAggregationType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescGZIP(), []int{0, 1, 1}
}

type SpansAggregation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Aggregation:
	//	*SpansAggregation_MetricAggregation_
	//	*SpansAggregation_DimensionAggregation_
	Aggregation isSpansAggregation_Aggregation `protobuf_oneof:"aggregation"`
}

func (x *SpansAggregation) Reset() {
	*x = SpansAggregation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpansAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpansAggregation) ProtoMessage() {}

func (x *SpansAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpansAggregation.ProtoReflect.Descriptor instead.
func (*SpansAggregation) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescGZIP(), []int{0}
}

func (m *SpansAggregation) GetAggregation() isSpansAggregation_Aggregation {
	if m != nil {
		return m.Aggregation
	}
	return nil
}

func (x *SpansAggregation) GetMetricAggregation() *SpansAggregation_MetricAggregation {
	if x, ok := x.GetAggregation().(*SpansAggregation_MetricAggregation_); ok {
		return x.MetricAggregation
	}
	return nil
}

func (x *SpansAggregation) GetDimensionAggregation() *SpansAggregation_DimensionAggregation {
	if x, ok := x.GetAggregation().(*SpansAggregation_DimensionAggregation_); ok {
		return x.DimensionAggregation
	}
	return nil
}

type isSpansAggregation_Aggregation interface {
	isSpansAggregation_Aggregation()
}

type SpansAggregation_MetricAggregation_ struct {
	MetricAggregation *SpansAggregation_MetricAggregation `protobuf:"bytes,1,opt,name=metric_aggregation,json=metricAggregation,proto3,oneof"`
}

type SpansAggregation_DimensionAggregation_ struct {
	DimensionAggregation *SpansAggregation_DimensionAggregation `protobuf:"bytes,2,opt,name=dimension_aggregation,json=dimensionAggregation,proto3,oneof"`
}

func (*SpansAggregation_MetricAggregation_) isSpansAggregation_Aggregation() {}

func (*SpansAggregation_DimensionAggregation_) isSpansAggregation_Aggregation() {}

type SpansAggregation_MetricAggregation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricField     SpansAggregation_MetricAggregation_MetricField           `protobuf:"varint,1,opt,name=metric_field,json=metricField,proto3,enum=com.coralogixapis.dashboards.v1.common.SpansAggregation_MetricAggregation_MetricField" json:"metric_field,omitempty"`
	AggregationType SpansAggregation_MetricAggregation_MetricAggregationType `protobuf:"varint,2,opt,name=aggregation_type,json=aggregationType,proto3,enum=com.coralogixapis.dashboards.v1.common.SpansAggregation_MetricAggregation_MetricAggregationType" json:"aggregation_type,omitempty"`
}

func (x *SpansAggregation_MetricAggregation) Reset() {
	*x = SpansAggregation_MetricAggregation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpansAggregation_MetricAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpansAggregation_MetricAggregation) ProtoMessage() {}

func (x *SpansAggregation_MetricAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpansAggregation_MetricAggregation.ProtoReflect.Descriptor instead.
func (*SpansAggregation_MetricAggregation) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SpansAggregation_MetricAggregation) GetMetricField() SpansAggregation_MetricAggregation_MetricField {
	if x != nil {
		return x.MetricField
	}
	return SpansAggregation_MetricAggregation_METRIC_FIELD_UNSPECIFIED
}

func (x *SpansAggregation_MetricAggregation) GetAggregationType() SpansAggregation_MetricAggregation_MetricAggregationType {
	if x != nil {
		return x.AggregationType
	}
	return SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_UNSPECIFIED
}

type SpansAggregation_DimensionAggregation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DimensionField  SpansAggregation_DimensionAggregation_DimensionField           `protobuf:"varint,1,opt,name=dimension_field,json=dimensionField,proto3,enum=com.coralogixapis.dashboards.v1.common.SpansAggregation_DimensionAggregation_DimensionField" json:"dimension_field,omitempty"`
	AggregationType SpansAggregation_DimensionAggregation_DimensionAggregationType `protobuf:"varint,2,opt,name=aggregation_type,json=aggregationType,proto3,enum=com.coralogixapis.dashboards.v1.common.SpansAggregation_DimensionAggregation_DimensionAggregationType" json:"aggregation_type,omitempty"`
}

func (x *SpansAggregation_DimensionAggregation) Reset() {
	*x = SpansAggregation_DimensionAggregation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpansAggregation_DimensionAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpansAggregation_DimensionAggregation) ProtoMessage() {}

func (x *SpansAggregation_DimensionAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpansAggregation_DimensionAggregation.ProtoReflect.Descriptor instead.
func (*SpansAggregation_DimensionAggregation) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SpansAggregation_DimensionAggregation) GetDimensionField() SpansAggregation_DimensionAggregation_DimensionField {
	if x != nil {
		return x.DimensionField
	}
	return SpansAggregation_DimensionAggregation_DIMENSION_FIELD_UNSPECIFIED
}

func (x *SpansAggregation_DimensionAggregation) GetAggregationType() SpansAggregation_DimensionAggregation_DimensionAggregationType {
	if x != nil {
		return x.AggregationType
	}
	return SpansAggregation_DimensionAggregation_DIMENSION_AGGREGATION_TYPE_UNSPECIFIED
}

var File_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x5f, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x80, 0x0c, 0x0a, 0x10, 0x53, 0x70, 0x61,
	0x6e, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7b, 0x0a,
	0x12, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x11, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x84, 0x01, 0x0a, 0x15, 0x64,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x14, 0x64, 0x69, 0x6d,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0xb0, 0x05, 0x0a, 0x11, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x79, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x56, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x8b, 0x01, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x60, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x46, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x1c, 0x0a, 0x18, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a,
	0x15, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x44, 0x55,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x22, 0xc9, 0x02, 0x0a, 0x15, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x27, 0x0a, 0x23, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x41, 0x47, 0x47,
	0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x4d,
	0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b,
	0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x58, 0x10, 0x02, 0x12, 0x23, 0x0a,
	0x1f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x56, 0x45, 0x52, 0x41, 0x47, 0x45,
	0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x41, 0x47, 0x47,
	0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x55,
	0x4d, 0x10, 0x04, 0x12, 0x29, 0x0a, 0x25, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x41, 0x47,
	0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x49, 0x4c, 0x45, 0x5f, 0x39, 0x39, 0x10, 0x05, 0x12, 0x29,
	0x0a, 0x25, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e,
	0x54, 0x49, 0x4c, 0x45, 0x5f, 0x39, 0x35, 0x10, 0x06, 0x12, 0x29, 0x0a, 0x25, 0x4d, 0x45, 0x54,
	0x52, 0x49, 0x43, 0x5f, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x49, 0x4c, 0x45, 0x5f,
	0x35, 0x30, 0x10, 0x07, 0x1a, 0xa5, 0x04, 0x0a, 0x14, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x85, 0x01,
	0x0a, 0x0f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0e, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x91, 0x01, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x66, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4f, 0x0a, 0x0e, 0x44, 0x69, 0x6d,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x1b, 0x44,
	0x49, 0x4d, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18,
	0x44, 0x49, 0x4d, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f,
	0x54, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x22, 0x9f, 0x01, 0x0a, 0x18, 0x44,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x26, 0x44, 0x49, 0x4d, 0x45, 0x4e,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x2b, 0x0a, 0x27, 0x44, 0x49, 0x4d, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01,
	0x12, 0x2a, 0x0a, 0x26, 0x44, 0x49, 0x4d, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x47,
	0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x02, 0x42, 0x0d, 0x0a, 0x0b,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x28, 0x5a, 0x26, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescData = file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_goTypes = []interface{}{
	(SpansAggregation_MetricAggregation_MetricField)(0),                 // 0: com.coralogixapis.dashboards.v1.common.SpansAggregation.MetricAggregation.MetricField
	(SpansAggregation_MetricAggregation_MetricAggregationType)(0),       // 1: com.coralogixapis.dashboards.v1.common.SpansAggregation.MetricAggregation.MetricAggregationType
	(SpansAggregation_DimensionAggregation_DimensionField)(0),           // 2: com.coralogixapis.dashboards.v1.common.SpansAggregation.DimensionAggregation.DimensionField
	(SpansAggregation_DimensionAggregation_DimensionAggregationType)(0), // 3: com.coralogixapis.dashboards.v1.common.SpansAggregation.DimensionAggregation.DimensionAggregationType
	(*SpansAggregation)(nil),                                            // 4: com.coralogixapis.dashboards.v1.common.SpansAggregation
	(*SpansAggregation_MetricAggregation)(nil),                          // 5: com.coralogixapis.dashboards.v1.common.SpansAggregation.MetricAggregation
	(*SpansAggregation_DimensionAggregation)(nil),                       // 6: com.coralogixapis.dashboards.v1.common.SpansAggregation.DimensionAggregation
}
var file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_depIdxs = []int32{
	5, // 0: com.coralogixapis.dashboards.v1.common.SpansAggregation.metric_aggregation:type_name -> com.coralogixapis.dashboards.v1.common.SpansAggregation.MetricAggregation
	6, // 1: com.coralogixapis.dashboards.v1.common.SpansAggregation.dimension_aggregation:type_name -> com.coralogixapis.dashboards.v1.common.SpansAggregation.DimensionAggregation
	0, // 2: com.coralogixapis.dashboards.v1.common.SpansAggregation.MetricAggregation.metric_field:type_name -> com.coralogixapis.dashboards.v1.common.SpansAggregation.MetricAggregation.MetricField
	1, // 3: com.coralogixapis.dashboards.v1.common.SpansAggregation.MetricAggregation.aggregation_type:type_name -> com.coralogixapis.dashboards.v1.common.SpansAggregation.MetricAggregation.MetricAggregationType
	2, // 4: com.coralogixapis.dashboards.v1.common.SpansAggregation.DimensionAggregation.dimension_field:type_name -> com.coralogixapis.dashboards.v1.common.SpansAggregation.DimensionAggregation.DimensionField
	3, // 5: com.coralogixapis.dashboards.v1.common.SpansAggregation.DimensionAggregation.aggregation_type:type_name -> com.coralogixapis.dashboards.v1.common.SpansAggregation.DimensionAggregation.DimensionAggregationType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_init() }
func file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_init() {
	if File_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpansAggregation); i {
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
		file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpansAggregation_MetricAggregation); i {
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
		file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpansAggregation_DimensionAggregation); i {
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
	file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SpansAggregation_MetricAggregation_)(nil),
		(*SpansAggregation_DimensionAggregation_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_depIdxs = nil
}
