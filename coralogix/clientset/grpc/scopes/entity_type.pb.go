// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: com/coralogixapis/scopes/v1/entity_type.proto

package v1

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

type EntityType int32

const (
	EntityType_UNSPECIFIED EntityType = 0
	EntityType_LOGS        EntityType = 1
	EntityType_SPANS       EntityType = 2
)

// Enum value maps for EntityType.
var (
	EntityType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "LOGS",
		2: "SPANS",
	}
	EntityType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"LOGS":        1,
		"SPANS":       2,
	}
)

func (x EntityType) Enum() *EntityType {
	p := new(EntityType)
	*p = x
	return p
}

func (x EntityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_scopes_v1_entity_type_proto_enumTypes[0].Descriptor()
}

func (EntityType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_scopes_v1_entity_type_proto_enumTypes[0]
}

func (x EntityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityType.Descriptor instead.
func (EntityType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_scopes_v1_entity_type_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_scopes_v1_entity_type_proto protoreflect.FileDescriptor

var file_com_coralogixapis_scopes_v1_entity_type_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x32, 0x0a, 0x0a,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c,
	0x4f, 0x47, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x50, 0x41, 0x4e, 0x53, 0x10, 0x02,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_scopes_v1_entity_type_proto_rawDescOnce sync.Once
	file_com_coralogixapis_scopes_v1_entity_type_proto_rawDescData = file_com_coralogixapis_scopes_v1_entity_type_proto_rawDesc
)

func file_com_coralogixapis_scopes_v1_entity_type_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_scopes_v1_entity_type_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_scopes_v1_entity_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_scopes_v1_entity_type_proto_rawDescData)
	})
	return file_com_coralogixapis_scopes_v1_entity_type_proto_rawDescData
}

var file_com_coralogixapis_scopes_v1_entity_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_scopes_v1_entity_type_proto_goTypes = []interface{}{
	(EntityType)(0), // 0: com.coralogixapis.scopes.v1.EntityType
}
var file_com_coralogixapis_scopes_v1_entity_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_scopes_v1_entity_type_proto_init() }
func file_com_coralogixapis_scopes_v1_entity_type_proto_init() {
	if File_com_coralogixapis_scopes_v1_entity_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_scopes_v1_entity_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_scopes_v1_entity_type_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_scopes_v1_entity_type_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_scopes_v1_entity_type_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_scopes_v1_entity_type_proto = out.File
	file_com_coralogixapis_scopes_v1_entity_type_proto_rawDesc = nil
	file_com_coralogixapis_scopes_v1_entity_type_proto_goTypes = nil
	file_com_coralogixapis_scopes_v1_entity_type_proto_depIdxs = nil
}
