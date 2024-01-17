// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogixapis/apm/services/v1/service_slo.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type CompareType int32

const (
	CompareType_COMPARE_TYPE_UNSPECIFIED CompareType = 0
	CompareType_COMPARE_TYPE_IS          CompareType = 1
	CompareType_COMPARE_TYPE_START_WITH  CompareType = 2
	CompareType_COMPARE_TYPE_ENDS_WITH   CompareType = 3
	CompareType_COMPARE_TYPE_INCLUDES    CompareType = 4
)

// Enum value maps for CompareType.
var (
	CompareType_name = map[int32]string{
		0: "COMPARE_TYPE_UNSPECIFIED",
		1: "COMPARE_TYPE_IS",
		2: "COMPARE_TYPE_START_WITH",
		3: "COMPARE_TYPE_ENDS_WITH",
		4: "COMPARE_TYPE_INCLUDES",
	}
	CompareType_value = map[string]int32{
		"COMPARE_TYPE_UNSPECIFIED": 0,
		"COMPARE_TYPE_IS":          1,
		"COMPARE_TYPE_START_WITH":  2,
		"COMPARE_TYPE_ENDS_WITH":   3,
		"COMPARE_TYPE_INCLUDES":    4,
	}
)

func (x CompareType) Enum() *CompareType {
	p := new(CompareType)
	*p = x
	return p
}

func (x CompareType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompareType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_enumTypes[0].Descriptor()
}

func (CompareType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_apm_services_v1_service_slo_proto_enumTypes[0]
}

func (x CompareType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompareType.Descriptor instead.
func (CompareType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescGZIP(), []int{0}
}

type SloStatus int32

const (
	SloStatus_SLO_STATUS_UNSPECIFIED SloStatus = 0
	SloStatus_SLO_STATUS_OK          SloStatus = 1
	SloStatus_SLO_STATUS_BREACHED    SloStatus = 2
)

// Enum value maps for SloStatus.
var (
	SloStatus_name = map[int32]string{
		0: "SLO_STATUS_UNSPECIFIED",
		1: "SLO_STATUS_OK",
		2: "SLO_STATUS_BREACHED",
	}
	SloStatus_value = map[string]int32{
		"SLO_STATUS_UNSPECIFIED": 0,
		"SLO_STATUS_OK":          1,
		"SLO_STATUS_BREACHED":    2,
	}
)

func (x SloStatus) Enum() *SloStatus {
	p := new(SloStatus)
	*p = x
	return p
}

func (x SloStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SloStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_enumTypes[1].Descriptor()
}

func (SloStatus) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_apm_services_v1_service_slo_proto_enumTypes[1]
}

func (x SloStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SloStatus.Descriptor instead.
func (SloStatus) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescGZIP(), []int{1}
}

type SliMetricType int32

const (
	SliMetricType_SLI_METRIC_TYPE_UNSPECIFIED SliMetricType = 0
	SliMetricType_SLI_METRIC_TYPE_ERROR       SliMetricType = 1
	SliMetricType_SLI_METRIC_TYPE_LATENCY     SliMetricType = 2
	SliMetricType_SLI_METRIC_TYPE_CUSTOM      SliMetricType = 3
)

// Enum value maps for SliMetricType.
var (
	SliMetricType_name = map[int32]string{
		0: "SLI_METRIC_TYPE_UNSPECIFIED",
		1: "SLI_METRIC_TYPE_ERROR",
		2: "SLI_METRIC_TYPE_LATENCY",
		3: "SLI_METRIC_TYPE_CUSTOM",
	}
	SliMetricType_value = map[string]int32{
		"SLI_METRIC_TYPE_UNSPECIFIED": 0,
		"SLI_METRIC_TYPE_ERROR":       1,
		"SLI_METRIC_TYPE_LATENCY":     2,
		"SLI_METRIC_TYPE_CUSTOM":      3,
	}
)

func (x SliMetricType) Enum() *SliMetricType {
	p := new(SliMetricType)
	*p = x
	return p
}

func (x SliMetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SliMetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_enumTypes[2].Descriptor()
}

func (SliMetricType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_apm_services_v1_service_slo_proto_enumTypes[2]
}

func (x SliMetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SliMetricType.Descriptor instead.
func (SliMetricType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescGZIP(), []int{2}
}

type ThresholdSymbol int32

const (
	ThresholdSymbol_THRESHOLD_SYMBOL_UNSPECIFIED      ThresholdSymbol = 0
	ThresholdSymbol_THRESHOLD_SYMBOL_GREATER          ThresholdSymbol = 1
	ThresholdSymbol_THRESHOLD_SYMBOL_GREATER_OR_EQUAL ThresholdSymbol = 2
	ThresholdSymbol_THRESHOLD_SYMBOL_LESS             ThresholdSymbol = 3
	ThresholdSymbol_THRESHOLD_SYMBOL_LESS_OR_EQUAL    ThresholdSymbol = 4
	ThresholdSymbol_THRESHOLD_SYMBOL_EQUAL            ThresholdSymbol = 5
	ThresholdSymbol_THRESHOLD_SYMBOL_NOT_EQUAL        ThresholdSymbol = 6
)

// Enum value maps for ThresholdSymbol.
var (
	ThresholdSymbol_name = map[int32]string{
		0: "THRESHOLD_SYMBOL_UNSPECIFIED",
		1: "THRESHOLD_SYMBOL_GREATER",
		2: "THRESHOLD_SYMBOL_GREATER_OR_EQUAL",
		3: "THRESHOLD_SYMBOL_LESS",
		4: "THRESHOLD_SYMBOL_LESS_OR_EQUAL",
		5: "THRESHOLD_SYMBOL_EQUAL",
		6: "THRESHOLD_SYMBOL_NOT_EQUAL",
	}
	ThresholdSymbol_value = map[string]int32{
		"THRESHOLD_SYMBOL_UNSPECIFIED":      0,
		"THRESHOLD_SYMBOL_GREATER":          1,
		"THRESHOLD_SYMBOL_GREATER_OR_EQUAL": 2,
		"THRESHOLD_SYMBOL_LESS":             3,
		"THRESHOLD_SYMBOL_LESS_OR_EQUAL":    4,
		"THRESHOLD_SYMBOL_EQUAL":            5,
		"THRESHOLD_SYMBOL_NOT_EQUAL":        6,
	}
)

func (x ThresholdSymbol) Enum() *ThresholdSymbol {
	p := new(ThresholdSymbol)
	*p = x
	return p
}

func (x ThresholdSymbol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ThresholdSymbol) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_enumTypes[3].Descriptor()
}

func (ThresholdSymbol) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_apm_services_v1_service_slo_proto_enumTypes[3]
}

func (x ThresholdSymbol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ThresholdSymbol.Descriptor instead.
func (ThresholdSymbol) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescGZIP(), []int{3}
}

type SloPeriod int32

const (
	SloPeriod_SLO_PERIOD_UNSPECIFIED SloPeriod = 0
	SloPeriod_SLO_PERIOD_7_DAYS      SloPeriod = 1
	SloPeriod_SLO_PERIOD_14_DAYS     SloPeriod = 2
	SloPeriod_SLO_PERIOD_30_DAYS     SloPeriod = 3
)

// Enum value maps for SloPeriod.
var (
	SloPeriod_name = map[int32]string{
		0: "SLO_PERIOD_UNSPECIFIED",
		1: "SLO_PERIOD_7_DAYS",
		2: "SLO_PERIOD_14_DAYS",
		3: "SLO_PERIOD_30_DAYS",
	}
	SloPeriod_value = map[string]int32{
		"SLO_PERIOD_UNSPECIFIED": 0,
		"SLO_PERIOD_7_DAYS":      1,
		"SLO_PERIOD_14_DAYS":     2,
		"SLO_PERIOD_30_DAYS":     3,
	}
)

func (x SloPeriod) Enum() *SloPeriod {
	p := new(SloPeriod)
	*p = x
	return p
}

func (x SloPeriod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SloPeriod) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_enumTypes[4].Descriptor()
}

func (SloPeriod) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_apm_services_v1_service_slo_proto_enumTypes[4]
}

func (x SloPeriod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SloPeriod.Descriptor instead.
func (SloPeriod) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescGZIP(), []int{4}
}

type SliFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field       *wrapperspb.StringValue   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	CompareType CompareType               `protobuf:"varint,2,opt,name=compare_type,json=compareType,proto3,enum=com.coralogixapis.apm.services.v1.CompareType" json:"compare_type,omitempty"`
	FieldValues []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=field_values,json=fieldValues,proto3" json:"field_values,omitempty"`
}

func (x *SliFilter) Reset() {
	*x = SliFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SliFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SliFilter) ProtoMessage() {}

func (x *SliFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SliFilter.ProtoReflect.Descriptor instead.
func (*SliFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescGZIP(), []int{0}
}

func (x *SliFilter) GetField() *wrapperspb.StringValue {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *SliFilter) GetCompareType() CompareType {
	if x != nil {
		return x.CompareType
	}
	return CompareType_COMPARE_TYPE_UNSPECIFIED
}

func (x *SliFilter) GetFieldValues() []*wrapperspb.StringValue {
	if x != nil {
		return x.FieldValues
	}
	return nil
}

type LatencySli struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThresholdMicroseconds *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=threshold_microseconds,json=thresholdMicroseconds,proto3" json:"threshold_microseconds,omitempty"`
	ThresholdSymbol       ThresholdSymbol         `protobuf:"varint,2,opt,name=threshold_symbol,json=thresholdSymbol,proto3,enum=com.coralogixapis.apm.services.v1.ThresholdSymbol" json:"threshold_symbol,omitempty"`
}

func (x *LatencySli) Reset() {
	*x = LatencySli{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatencySli) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatencySli) ProtoMessage() {}

func (x *LatencySli) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatencySli.ProtoReflect.Descriptor instead.
func (*LatencySli) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescGZIP(), []int{1}
}

func (x *LatencySli) GetThresholdMicroseconds() *wrapperspb.StringValue {
	if x != nil {
		return x.ThresholdMicroseconds
	}
	return nil
}

func (x *LatencySli) GetThresholdSymbol() ThresholdSymbol {
	if x != nil {
		return x.ThresholdSymbol
	}
	return ThresholdSymbol_THRESHOLD_SYMBOL_UNSPECIFIED
}

type ErrorSli struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ErrorSli) Reset() {
	*x = ErrorSli{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorSli) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorSli) ProtoMessage() {}

func (x *ErrorSli) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorSli.ProtoReflect.Descriptor instead.
func (*ErrorSli) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescGZIP(), []int{2}
}

type ServiceSlo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                             *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                           *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ServiceName                    *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Status                         SloStatus               `protobuf:"varint,4,opt,name=status,proto3,enum=com.coralogixapis.apm.services.v1.SloStatus" json:"status,omitempty"`
	Description                    *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	TargetPercentage               *wrapperspb.UInt32Value `protobuf:"bytes,6,opt,name=target_percentage,json=targetPercentage,proto3" json:"target_percentage,omitempty"`
	CreatedAt                      *timestamppb.Timestamp  `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	RemainingErrorBudgetPercentage *wrapperspb.UInt32Value `protobuf:"bytes,8,opt,name=remaining_error_budget_percentage,json=remainingErrorBudgetPercentage,proto3" json:"remaining_error_budget_percentage,omitempty"`
	// Types that are assignable to SliType:
	//	*ServiceSlo_LatencySli
	//	*ServiceSlo_ErrorSli
	SliType isServiceSlo_SliType `protobuf_oneof:"sli_type"`
	Filters []*SliFilter         `protobuf:"bytes,11,rep,name=filters,proto3" json:"filters,omitempty"`
	Period  SloPeriod            `protobuf:"varint,12,opt,name=period,proto3,enum=com.coralogixapis.apm.services.v1.SloPeriod" json:"period,omitempty"`
}

func (x *ServiceSlo) Reset() {
	*x = ServiceSlo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceSlo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceSlo) ProtoMessage() {}

func (x *ServiceSlo) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceSlo.ProtoReflect.Descriptor instead.
func (*ServiceSlo) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceSlo) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ServiceSlo) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ServiceSlo) GetServiceName() *wrapperspb.StringValue {
	if x != nil {
		return x.ServiceName
	}
	return nil
}

func (x *ServiceSlo) GetStatus() SloStatus {
	if x != nil {
		return x.Status
	}
	return SloStatus_SLO_STATUS_UNSPECIFIED
}

func (x *ServiceSlo) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *ServiceSlo) GetTargetPercentage() *wrapperspb.UInt32Value {
	if x != nil {
		return x.TargetPercentage
	}
	return nil
}

func (x *ServiceSlo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ServiceSlo) GetRemainingErrorBudgetPercentage() *wrapperspb.UInt32Value {
	if x != nil {
		return x.RemainingErrorBudgetPercentage
	}
	return nil
}

func (m *ServiceSlo) GetSliType() isServiceSlo_SliType {
	if m != nil {
		return m.SliType
	}
	return nil
}

func (x *ServiceSlo) GetLatencySli() *LatencySli {
	if x, ok := x.GetSliType().(*ServiceSlo_LatencySli); ok {
		return x.LatencySli
	}
	return nil
}

func (x *ServiceSlo) GetErrorSli() *ErrorSli {
	if x, ok := x.GetSliType().(*ServiceSlo_ErrorSli); ok {
		return x.ErrorSli
	}
	return nil
}

func (x *ServiceSlo) GetFilters() []*SliFilter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ServiceSlo) GetPeriod() SloPeriod {
	if x != nil {
		return x.Period
	}
	return SloPeriod_SLO_PERIOD_UNSPECIFIED
}

type isServiceSlo_SliType interface {
	isServiceSlo_SliType()
}

type ServiceSlo_LatencySli struct {
	LatencySli *LatencySli `protobuf:"bytes,9,opt,name=latency_sli,json=latencySli,proto3,oneof"`
}

type ServiceSlo_ErrorSli struct {
	ErrorSli *ErrorSli `protobuf:"bytes,10,opt,name=error_sli,json=errorSli,proto3,oneof"`
}

func (*ServiceSlo_LatencySli) isServiceSlo_SliType() {}

func (*ServiceSlo_ErrorSli) isServiceSlo_SliType() {}

var File_com_coralogixapis_apm_services_v1_service_slo_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x09, 0x53, 0x6c,
	0x69, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x51, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3f,
	0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22,
	0xc0, 0x01, 0x0a, 0x0a, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x6c, 0x69, 0x12, 0x53,
	0x0a, 0x16, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x12, 0x5d, 0x0a, 0x10, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x52, 0x0f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x6c, 0x69, 0x22, 0xda,
	0x06, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x6c, 0x6f, 0x12, 0x2c, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6c, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x67, 0x0a, 0x21, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x62, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x1e, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73,
	0x6c, 0x69, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x53, 0x6c, 0x69, 0x48, 0x00, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x53, 0x6c, 0x69, 0x12, 0x4a, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73,
	0x6c, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x53, 0x6c, 0x69, 0x48, 0x00, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x6c,
	0x69, 0x12, 0x46, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x69, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x44, 0x0a, 0x06, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c,
	0x6f, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x42,
	0x0a, 0x0a, 0x08, 0x73, 0x6c, 0x69, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x94, 0x01, 0x0a, 0x0b,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x43,
	0x4f, 0x4d, 0x50, 0x41, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4d,
	0x50, 0x41, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x53, 0x10, 0x01, 0x12, 0x1b,
	0x0a, 0x17, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53,
	0x54, 0x41, 0x52, 0x54, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x43,
	0x4f, 0x4d, 0x50, 0x41, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x53,
	0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x50, 0x41,
	0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x53,
	0x10, 0x04, 0x2a, 0x53, 0x0a, 0x09, 0x53, 0x6c, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1a, 0x0a, 0x16, 0x53, 0x4c, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53,
	0x4c, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x13, 0x53, 0x4c, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x42, 0x52, 0x45,
	0x41, 0x43, 0x48, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x84, 0x01, 0x0a, 0x0d, 0x53, 0x6c, 0x69, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x4c, 0x49,
	0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x4c,
	0x49, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x4c, 0x49, 0x5f, 0x4d, 0x45, 0x54,
	0x52, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x41, 0x54, 0x45, 0x4e, 0x43, 0x59,
	0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x4c, 0x49, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x03, 0x2a, 0xf3,
	0x01, 0x0a, 0x0f, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x5f,
	0x53, 0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c,
	0x44, 0x5f, 0x53, 0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x5f, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52,
	0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x5f,
	0x53, 0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x5f, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x4f,
	0x52, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x48, 0x52,
	0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x53, 0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x5f, 0x4c, 0x45,
	0x53, 0x53, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c,
	0x44, 0x5f, 0x53, 0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x4f, 0x52,
	0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x48, 0x52, 0x45,
	0x53, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x53, 0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x5f, 0x45, 0x51, 0x55,
	0x41, 0x4c, 0x10, 0x05, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c,
	0x44, 0x5f, 0x53, 0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55,
	0x41, 0x4c, 0x10, 0x06, 0x2a, 0x6e, 0x0a, 0x09, 0x53, 0x6c, 0x6f, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x4c, 0x4f, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x53, 0x4c, 0x4f, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x5f, 0x37, 0x5f, 0x44, 0x41,
	0x59, 0x53, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x4c, 0x4f, 0x5f, 0x50, 0x45, 0x52, 0x49,
	0x4f, 0x44, 0x5f, 0x31, 0x34, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12,
	0x53, 0x4c, 0x4f, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x5f, 0x33, 0x30, 0x5f, 0x44, 0x41,
	0x59, 0x53, 0x10, 0x03, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescData = file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDesc
)

func file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDescData
}

var file_com_coralogixapis_apm_services_v1_service_slo_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogixapis_apm_services_v1_service_slo_proto_goTypes = []interface{}{
	(CompareType)(0),               // 0: com.coralogixapis.apm.services.v1.CompareType
	(SloStatus)(0),                 // 1: com.coralogixapis.apm.services.v1.SloStatus
	(SliMetricType)(0),             // 2: com.coralogixapis.apm.services.v1.SliMetricType
	(ThresholdSymbol)(0),           // 3: com.coralogixapis.apm.services.v1.ThresholdSymbol
	(SloPeriod)(0),                 // 4: com.coralogixapis.apm.services.v1.SloPeriod
	(*SliFilter)(nil),              // 5: com.coralogixapis.apm.services.v1.SliFilter
	(*LatencySli)(nil),             // 6: com.coralogixapis.apm.services.v1.LatencySli
	(*ErrorSli)(nil),               // 7: com.coralogixapis.apm.services.v1.ErrorSli
	(*ServiceSlo)(nil),             // 8: com.coralogixapis.apm.services.v1.ServiceSlo
	(*wrapperspb.StringValue)(nil), // 9: google.protobuf.StringValue
	(*wrapperspb.UInt32Value)(nil), // 10: google.protobuf.UInt32Value
	(*timestamppb.Timestamp)(nil),  // 11: google.protobuf.Timestamp
}
var file_com_coralogixapis_apm_services_v1_service_slo_proto_depIdxs = []int32{
	9,  // 0: com.coralogixapis.apm.services.v1.SliFilter.field:type_name -> google.protobuf.StringValue
	0,  // 1: com.coralogixapis.apm.services.v1.SliFilter.compare_type:type_name -> com.coralogixapis.apm.services.v1.CompareType
	9,  // 2: com.coralogixapis.apm.services.v1.SliFilter.field_values:type_name -> google.protobuf.StringValue
	9,  // 3: com.coralogixapis.apm.services.v1.LatencySli.threshold_microseconds:type_name -> google.protobuf.StringValue
	3,  // 4: com.coralogixapis.apm.services.v1.LatencySli.threshold_symbol:type_name -> com.coralogixapis.apm.services.v1.ThresholdSymbol
	9,  // 5: com.coralogixapis.apm.services.v1.ServiceSlo.id:type_name -> google.protobuf.StringValue
	9,  // 6: com.coralogixapis.apm.services.v1.ServiceSlo.name:type_name -> google.protobuf.StringValue
	9,  // 7: com.coralogixapis.apm.services.v1.ServiceSlo.service_name:type_name -> google.protobuf.StringValue
	1,  // 8: com.coralogixapis.apm.services.v1.ServiceSlo.status:type_name -> com.coralogixapis.apm.services.v1.SloStatus
	9,  // 9: com.coralogixapis.apm.services.v1.ServiceSlo.description:type_name -> google.protobuf.StringValue
	10, // 10: com.coralogixapis.apm.services.v1.ServiceSlo.target_percentage:type_name -> google.protobuf.UInt32Value
	11, // 11: com.coralogixapis.apm.services.v1.ServiceSlo.created_at:type_name -> google.protobuf.Timestamp
	10, // 12: com.coralogixapis.apm.services.v1.ServiceSlo.remaining_error_budget_percentage:type_name -> google.protobuf.UInt32Value
	6,  // 13: com.coralogixapis.apm.services.v1.ServiceSlo.latency_sli:type_name -> com.coralogixapis.apm.services.v1.LatencySli
	7,  // 14: com.coralogixapis.apm.services.v1.ServiceSlo.error_sli:type_name -> com.coralogixapis.apm.services.v1.ErrorSli
	5,  // 15: com.coralogixapis.apm.services.v1.ServiceSlo.filters:type_name -> com.coralogixapis.apm.services.v1.SliFilter
	4,  // 16: com.coralogixapis.apm.services.v1.ServiceSlo.period:type_name -> com.coralogixapis.apm.services.v1.SloPeriod
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_services_v1_service_slo_proto_init() }
func file_com_coralogixapis_apm_services_v1_service_slo_proto_init() {
	if File_com_coralogixapis_apm_services_v1_service_slo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SliFilter); i {
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
		file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatencySli); i {
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
		file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorSli); i {
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
		file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceSlo); i {
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
	file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ServiceSlo_LatencySli)(nil),
		(*ServiceSlo_ErrorSli)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_services_v1_service_slo_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_services_v1_service_slo_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_apm_services_v1_service_slo_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_apm_services_v1_service_slo_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_services_v1_service_slo_proto = out.File
	file_com_coralogixapis_apm_services_v1_service_slo_proto_rawDesc = nil
	file_com_coralogixapis_apm_services_v1_service_slo_proto_goTypes = nil
	file_com_coralogixapis_apm_services_v1_service_slo_proto_depIdxs = nil
}
