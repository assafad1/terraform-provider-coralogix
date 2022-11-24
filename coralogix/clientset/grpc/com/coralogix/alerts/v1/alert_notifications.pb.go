// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogix/alerts/v1/alert_notifications.proto

package __

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

type AlertNotifications struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emails       []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=emails,proto3" json:"emails,omitempty"`
	Integrations []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=integrations,proto3" json:"integrations,omitempty"`
}

func (x *AlertNotifications) Reset() {
	*x = AlertNotifications{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_alerts_v1_alert_notifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertNotifications) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertNotifications) ProtoMessage() {}

func (x *AlertNotifications) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_alerts_v1_alert_notifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertNotifications.ProtoReflect.Descriptor instead.
func (*AlertNotifications) Descriptor() ([]byte, []int) {
	return file_com_coralogix_alerts_v1_alert_notifications_proto_rawDescGZIP(), []int{0}
}

func (x *AlertNotifications) GetEmails() []*wrapperspb.StringValue {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *AlertNotifications) GetIntegrations() []*wrapperspb.StringValue {
	if x != nil {
		return x.Integrations
	}
	return nil
}

var File_com_coralogix_alerts_v1_alert_notifications_proto protoreflect.FileDescriptor

var file_com_coralogix_alerts_v1_alert_notifications_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a,
	0x12, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x04, 0x5a, 0x02, 0x2e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_alerts_v1_alert_notifications_proto_rawDescOnce sync.Once
	file_com_coralogix_alerts_v1_alert_notifications_proto_rawDescData = file_com_coralogix_alerts_v1_alert_notifications_proto_rawDesc
)

func file_com_coralogix_alerts_v1_alert_notifications_proto_rawDescGZIP() []byte {
	file_com_coralogix_alerts_v1_alert_notifications_proto_rawDescOnce.Do(func() {
		file_com_coralogix_alerts_v1_alert_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_alerts_v1_alert_notifications_proto_rawDescData)
	})
	return file_com_coralogix_alerts_v1_alert_notifications_proto_rawDescData
}

var file_com_coralogix_alerts_v1_alert_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_alerts_v1_alert_notifications_proto_goTypes = []interface{}{
	(*AlertNotifications)(nil),     // 0: com.coralogix.alerts.v1.AlertNotifications
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_com_coralogix_alerts_v1_alert_notifications_proto_depIdxs = []int32{
	1, // 0: com.coralogix.alerts.v1.AlertNotifications.emails:type_name -> google.protobuf.StringValue
	1, // 1: com.coralogix.alerts.v1.AlertNotifications.integrations:type_name -> google.protobuf.StringValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogix_alerts_v1_alert_notifications_proto_init() }
func file_com_coralogix_alerts_v1_alert_notifications_proto_init() {
	if File_com_coralogix_alerts_v1_alert_notifications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_alerts_v1_alert_notifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertNotifications); i {
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
			RawDescriptor: file_com_coralogix_alerts_v1_alert_notifications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_alerts_v1_alert_notifications_proto_goTypes,
		DependencyIndexes: file_com_coralogix_alerts_v1_alert_notifications_proto_depIdxs,
		MessageInfos:      file_com_coralogix_alerts_v1_alert_notifications_proto_msgTypes,
	}.Build()
	File_com_coralogix_alerts_v1_alert_notifications_proto = out.File
	file_com_coralogix_alerts_v1_alert_notifications_proto_rawDesc = nil
	file_com_coralogix_alerts_v1_alert_notifications_proto_goTypes = nil
	file_com_coralogix_alerts_v1_alert_notifications_proto_depIdxs = nil
}
