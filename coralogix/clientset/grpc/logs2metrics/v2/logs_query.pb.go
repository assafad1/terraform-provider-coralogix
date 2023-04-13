// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogixapis/logs2metrics/v2/logs_query.proto

package __

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Severity int32

const (
	Severity_Unspecified Severity = 0
	Severity_Debug       Severity = 1
	Severity_Verbose     Severity = 2
	Severity_Info        Severity = 3
	Severity_Warning     Severity = 4
	Severity_Error       Severity = 5
	Severity_Critical    Severity = 6
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0: "Unspecified",
		1: "Debug",
		2: "Verbose",
		3: "Info",
		4: "Warning",
		5: "Error",
		6: "Critical",
	}
	Severity_value = map[string]int32{
		"Unspecified": 0,
		"Debug":       1,
		"Verbose":     2,
		"Info":        3,
		"Warning":     4,
		"Error":       5,
		"Critical":    6,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_logs2metrics_v2_logs_query_proto_enumTypes[0].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_logs2metrics_v2_logs_query_proto_enumTypes[0]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescGZIP(), []int{0}
}

type LogsQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lucene                 *wrapperspb.StringValue   `protobuf:"bytes,1,opt,name=lucene,proto3" json:"lucene,omitempty"`
	Alias                  *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	ApplicationnameFilters []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=applicationname_filters,json=applicationnameFilters,proto3" json:"applicationname_filters,omitempty"`
	SubsystemnameFilters   []*wrapperspb.StringValue `protobuf:"bytes,4,rep,name=subsystemname_filters,json=subsystemnameFilters,proto3" json:"subsystemname_filters,omitempty"`
	SeverityFilters        []Severity                `protobuf:"varint,5,rep,packed,name=severity_filters,json=severityFilters,proto3,enum=com.coralogixapis.logs2metrics.v2.Severity" json:"severity_filters,omitempty"`
}

func (x *LogsQuery) Reset() {
	*x = LogsQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_logs2metrics_v2_logs_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsQuery) ProtoMessage() {}

func (x *LogsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_logs2metrics_v2_logs_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsQuery.ProtoReflect.Descriptor instead.
func (*LogsQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescGZIP(), []int{0}
}

func (x *LogsQuery) GetLucene() *wrapperspb.StringValue {
	if x != nil {
		return x.Lucene
	}
	return nil
}

func (x *LogsQuery) GetAlias() *wrapperspb.StringValue {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *LogsQuery) GetApplicationnameFilters() []*wrapperspb.StringValue {
	if x != nil {
		return x.ApplicationnameFilters
	}
	return nil
}

func (x *LogsQuery) GetSubsystemnameFilters() []*wrapperspb.StringValue {
	if x != nil {
		return x.SubsystemnameFilters
	}
	return nil
}

func (x *LogsQuery) GetSeverityFilters() []Severity {
	if x != nil {
		return x.SeverityFilters
	}
	return nil
}

var File_com_coralogixapis_logs2metrics_v2_logs_query_proto protoreflect.FileDescriptor

var file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x32, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x02, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x73,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x34, 0x0a, 0x06, 0x6c, 0x75, 0x63, 0x65, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x6c, 0x75, 0x63, 0x65, 0x6e, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x61,
	0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12,
	0x55, 0x0a, 0x17, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x16,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x51, 0x0a, 0x15, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x14, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x6e, 0x61,
	0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x56, 0x0a, 0x10, 0x73, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x32, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x0f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2a, 0x63, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x65, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x04, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x72, 0x69, 0x74,
	0x69, 0x63, 0x61, 0x6c, 0x10, 0x06, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescOnce sync.Once
	file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescData = file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDesc
)

func file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescData)
	})
	return file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescData
}

var file_com_coralogixapis_logs2metrics_v2_logs_query_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_logs2metrics_v2_logs_query_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_logs2metrics_v2_logs_query_proto_goTypes = []interface{}{
	(Severity)(0),                  // 0: com.coralogixapis.logs2metrics.v2.Severity
	(*LogsQuery)(nil),              // 1: com.coralogixapis.logs2metrics.v2.LogsQuery
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_com_coralogixapis_logs2metrics_v2_logs_query_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.logs2metrics.v2.LogsQuery.lucene:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogixapis.logs2metrics.v2.LogsQuery.alias:type_name -> google.protobuf.StringValue
	2, // 2: com.coralogixapis.logs2metrics.v2.LogsQuery.applicationname_filters:type_name -> google.protobuf.StringValue
	2, // 3: com.coralogixapis.logs2metrics.v2.LogsQuery.subsystemname_filters:type_name -> google.protobuf.StringValue
	0, // 4: com.coralogixapis.logs2metrics.v2.LogsQuery.severity_filters:type_name -> com.coralogixapis.logs2metrics.v2.Severity
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { File_com_coralogixapis_logs2metrics_v2_logs_query_proto_init() }
func File_com_coralogixapis_logs2metrics_v2_logs_query_proto_init() {
	if File_com_coralogixapis_logs2metrics_v2_logs_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_logs2metrics_v2_logs_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsQuery); i {
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
			RawDescriptor: file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_logs2metrics_v2_logs_query_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_logs2metrics_v2_logs_query_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_logs2metrics_v2_logs_query_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_logs2metrics_v2_logs_query_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_logs2metrics_v2_logs_query_proto = out.File
	file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDesc = nil
	file_com_coralogixapis_logs2metrics_v2_logs_query_proto_goTypes = nil
	file_com_coralogixapis_logs2metrics_v2_logs_query_proto_depIdxs = nil
}
