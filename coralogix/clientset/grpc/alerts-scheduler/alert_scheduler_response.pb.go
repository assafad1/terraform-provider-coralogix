// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_response.proto

package __

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

type GetAlertSchedulerRuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlertSchedulerRule *AlertSchedulerRule `protobuf:"bytes,1,opt,name=alert_scheduler_rule,json=alertSchedulerRule,proto3" json:"alert_scheduler_rule,omitempty"`
}

func (x *GetAlertSchedulerRuleResponse) Reset() {
	*x = GetAlertSchedulerRuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlertSchedulerRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertSchedulerRuleResponse) ProtoMessage() {}

func (x *GetAlertSchedulerRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlertSchedulerRuleResponse.ProtoReflect.Descriptor instead.
func (*GetAlertSchedulerRuleResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescGZIP(), []int{0}
}

func (x *GetAlertSchedulerRuleResponse) GetAlertSchedulerRule() *AlertSchedulerRule {
	if x != nil {
		return x.AlertSchedulerRule
	}
	return nil
}

type CreateAlertSchedulerRuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlertSchedulerRule *AlertSchedulerRule `protobuf:"bytes,1,opt,name=alert_scheduler_rule,json=alertSchedulerRule,proto3" json:"alert_scheduler_rule,omitempty"`
}

func (x *CreateAlertSchedulerRuleResponse) Reset() {
	*x = CreateAlertSchedulerRuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlertSchedulerRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlertSchedulerRuleResponse) ProtoMessage() {}

func (x *CreateAlertSchedulerRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlertSchedulerRuleResponse.ProtoReflect.Descriptor instead.
func (*CreateAlertSchedulerRuleResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAlertSchedulerRuleResponse) GetAlertSchedulerRule() *AlertSchedulerRule {
	if x != nil {
		return x.AlertSchedulerRule
	}
	return nil
}

type UpdateAlertSchedulerRuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlertSchedulerRule *AlertSchedulerRule `protobuf:"bytes,1,opt,name=alert_scheduler_rule,json=alertSchedulerRule,proto3" json:"alert_scheduler_rule,omitempty"`
}

func (x *UpdateAlertSchedulerRuleResponse) Reset() {
	*x = UpdateAlertSchedulerRuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAlertSchedulerRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAlertSchedulerRuleResponse) ProtoMessage() {}

func (x *UpdateAlertSchedulerRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAlertSchedulerRuleResponse.ProtoReflect.Descriptor instead.
func (*UpdateAlertSchedulerRuleResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAlertSchedulerRuleResponse) GetAlertSchedulerRule() *AlertSchedulerRule {
	if x != nil {
		return x.AlertSchedulerRule
	}
	return nil
}

type DeleteAlertSchedulerRuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAlertSchedulerRuleResponse) Reset() {
	*x = DeleteAlertSchedulerRuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlertSchedulerRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlertSchedulerRuleResponse) ProtoMessage() {}

func (x *DeleteAlertSchedulerRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAlertSchedulerRuleResponse.ProtoReflect.Descriptor instead.
func (*DeleteAlertSchedulerRuleResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescGZIP(), []int{3}
}

type CreateBulkAlertSchedulerRuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateSuppressionResponses []*CreateAlertSchedulerRuleResponse `protobuf:"bytes,1,rep,name=create_suppression_responses,json=createSuppressionResponses,proto3" json:"create_suppression_responses,omitempty"`
}

func (x *CreateBulkAlertSchedulerRuleResponse) Reset() {
	*x = CreateBulkAlertSchedulerRuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBulkAlertSchedulerRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBulkAlertSchedulerRuleResponse) ProtoMessage() {}

func (x *CreateBulkAlertSchedulerRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBulkAlertSchedulerRuleResponse.ProtoReflect.Descriptor instead.
func (*CreateBulkAlertSchedulerRuleResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBulkAlertSchedulerRuleResponse) GetCreateSuppressionResponses() []*CreateAlertSchedulerRuleResponse {
	if x != nil {
		return x.CreateSuppressionResponses
	}
	return nil
}

type UpdateBulkAlertSchedulerRuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateSuppressionResponses []*UpdateAlertSchedulerRuleResponse `protobuf:"bytes,1,rep,name=update_suppression_responses,json=updateSuppressionResponses,proto3" json:"update_suppression_responses,omitempty"`
}

func (x *UpdateBulkAlertSchedulerRuleResponse) Reset() {
	*x = UpdateBulkAlertSchedulerRuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBulkAlertSchedulerRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBulkAlertSchedulerRuleResponse) ProtoMessage() {}

func (x *UpdateBulkAlertSchedulerRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBulkAlertSchedulerRuleResponse.ProtoReflect.Descriptor instead.
func (*UpdateBulkAlertSchedulerRuleResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBulkAlertSchedulerRuleResponse) GetUpdateSuppressionResponses() []*UpdateAlertSchedulerRuleResponse {
	if x != nil {
		return x.UpdateSuppressionResponses
	}
	return nil
}

type GetBulkAlertSchedulerRuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlertSchedulerRules []*AlertSchedulerRuleWithActiveTimeframe `protobuf:"bytes,1,rep,name=alert_scheduler_rules,json=alertSchedulerRules,proto3" json:"alert_scheduler_rules,omitempty"`
	NextPageToken       string                                   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *GetBulkAlertSchedulerRuleResponse) Reset() {
	*x = GetBulkAlertSchedulerRuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBulkAlertSchedulerRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBulkAlertSchedulerRuleResponse) ProtoMessage() {}

func (x *GetBulkAlertSchedulerRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBulkAlertSchedulerRuleResponse.ProtoReflect.Descriptor instead.
func (*GetBulkAlertSchedulerRuleResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescGZIP(), []int{6}
}

func (x *GetBulkAlertSchedulerRuleResponse) GetAlertSchedulerRules() []*AlertSchedulerRuleWithActiveTimeframe {
	if x != nil {
		return x.AlertSchedulerRules
	}
	return nil
}

func (x *GetBulkAlertSchedulerRuleResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type AlertSchedulerRuleWithActiveTimeframe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlertSchedulerRule   *AlertSchedulerRule `protobuf:"bytes,1,opt,name=alert_scheduler_rule,json=alertSchedulerRule,proto3" json:"alert_scheduler_rule,omitempty"`
	NextActiveTimeframes []*ActiveTimeframe  `protobuf:"bytes,2,rep,name=next_active_timeframes,json=nextActiveTimeframes,proto3" json:"next_active_timeframes,omitempty"`
}

func (x *AlertSchedulerRuleWithActiveTimeframe) Reset() {
	*x = AlertSchedulerRuleWithActiveTimeframe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertSchedulerRuleWithActiveTimeframe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertSchedulerRuleWithActiveTimeframe) ProtoMessage() {}

func (x *AlertSchedulerRuleWithActiveTimeframe) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertSchedulerRuleWithActiveTimeframe.ProtoReflect.Descriptor instead.
func (*AlertSchedulerRuleWithActiveTimeframe) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescGZIP(), []int{7}
}

func (x *AlertSchedulerRuleWithActiveTimeframe) GetAlertSchedulerRule() *AlertSchedulerRule {
	if x != nil {
		return x.AlertSchedulerRule
	}
	return nil
}

func (x *AlertSchedulerRuleWithActiveTimeframe) GetNextActiveTimeframes() []*ActiveTimeframe {
	if x != nil {
		return x.NextActiveTimeframes
	}
	return nil
}

var File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDesc = []byte{
	0x0a, 0x5a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x56, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x52, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x14, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x12, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x20,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x81, 0x01, 0x0a, 0x14, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x4f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x12, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x52, 0x75, 0x6c, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x14, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x12, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x22, 0x22, 0x0a,
	0x20, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xc8, 0x01, 0x0a, 0x24, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6c, 0x6b,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x1c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x5d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x1a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0xc8, 0x01, 0x0a,
	0x24, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6c, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x1c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x5d, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x1a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x6c, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x96, 0x01,
	0x0a, 0x15, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x52, 0x13, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb0,
	0x02, 0x0a, 0x25, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x52, 0x75, 0x6c, 0x65, 0x57, 0x69, 0x74, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x14, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x12, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x82, 0x01, 0x0a,
	0x16, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4c, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x14, 0x6e, 0x65, 0x78,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x73, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescData = file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDesc
)

func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescData)
	})
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDescData
}

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_goTypes = []interface{}{
	(*GetAlertSchedulerRuleResponse)(nil),         // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetAlertSchedulerRuleResponse
	(*CreateAlertSchedulerRuleResponse)(nil),      // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleResponse
	(*UpdateAlertSchedulerRuleResponse)(nil),      // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleResponse
	(*DeleteAlertSchedulerRuleResponse)(nil),      // 3: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleResponse
	(*CreateBulkAlertSchedulerRuleResponse)(nil),  // 4: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleResponse
	(*UpdateBulkAlertSchedulerRuleResponse)(nil),  // 5: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleResponse
	(*GetBulkAlertSchedulerRuleResponse)(nil),     // 6: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleResponse
	(*AlertSchedulerRuleWithActiveTimeframe)(nil), // 7: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleWithActiveTimeframe
	(*AlertSchedulerRule)(nil),                    // 8: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRule
	(*ActiveTimeframe)(nil),                       // 9: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.ActiveTimeframe
}
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_depIdxs = []int32{
	8, // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetAlertSchedulerRuleResponse.alert_scheduler_rule:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRule
	8, // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleResponse.alert_scheduler_rule:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRule
	8, // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleResponse.alert_scheduler_rule:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRule
	1, // 3: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleResponse.create_suppression_responses:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleResponse
	2, // 4: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleResponse.update_suppression_responses:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleResponse
	7, // 5: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleResponse.alert_scheduler_rules:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleWithActiveTimeframe
	8, // 6: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleWithActiveTimeframe.alert_scheduler_rule:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRule
	9, // 7: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleWithActiveTimeframe.next_active_timeframes:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.ActiveTimeframe
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_init()
}
func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_init() {
	if File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto != nil {
		return
	}
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_init()
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_active_timeframe_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlertSchedulerRuleResponse); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlertSchedulerRuleResponse); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAlertSchedulerRuleResponse); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAlertSchedulerRuleResponse); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBulkAlertSchedulerRuleResponse); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBulkAlertSchedulerRuleResponse); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBulkAlertSchedulerRuleResponse); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertSchedulerRuleWithActiveTimeframe); i {
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
			RawDescriptor: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto = out.File
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_rawDesc = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_goTypes = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_depIdxs = nil
}
