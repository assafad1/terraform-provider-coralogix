// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: com/coralogixapis/dashboards/v1/services/dashboard_catalog_service.proto

package golang

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

type GetDashboardCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDashboardCatalogRequest) Reset() {
	*x = GetDashboardCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDashboardCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardCatalogRequest) ProtoMessage() {}

func (x *GetDashboardCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardCatalogRequest.ProtoReflect.Descriptor instead.
func (*GetDashboardCatalogRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDescGZIP(), []int{0}
}

type GetDashboardCatalogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*DashboardCatalogItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetDashboardCatalogResponse) Reset() {
	*x = GetDashboardCatalogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDashboardCatalogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardCatalogResponse) ProtoMessage() {}

func (x *GetDashboardCatalogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardCatalogResponse.ProtoReflect.Descriptor instead.
func (*GetDashboardCatalogResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetDashboardCatalogResponse) GetItems() []*DashboardCatalogItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type DashboardCatalogItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IsDefault   *wrapperspb.BoolValue   `protobuf:"bytes,4,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	IsPinned    *wrapperspb.BoolValue   `protobuf:"bytes,5,opt,name=is_pinned,json=isPinned,proto3" json:"is_pinned,omitempty"`
	CreateTime  *timestamppb.Timestamp  `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime  *timestamppb.Timestamp  `protobuf:"bytes,7,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *DashboardCatalogItem) Reset() {
	*x = DashboardCatalogItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DashboardCatalogItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardCatalogItem) ProtoMessage() {}

func (x *DashboardCatalogItem) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DashboardCatalogItem.ProtoReflect.Descriptor instead.
func (*DashboardCatalogItem) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDescGZIP(), []int{2}
}

func (x *DashboardCatalogItem) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *DashboardCatalogItem) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *DashboardCatalogItem) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *DashboardCatalogItem) GetIsDefault() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsDefault
	}
	return nil
}

func (x *DashboardCatalogItem) GetIsPinned() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsPinned
	}
	return nil
}

func (x *DashboardCatalogItem) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *DashboardCatalogItem) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDesc = []byte{
	0x0a, 0x48, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x33,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x73, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xa4, 0x03, 0x0a, 0x14, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x37,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x69,
	0x73, 0x50, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x32, 0xdb, 0x01, 0x0a, 0x17, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbf, 0x01,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x44, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1b, 0xba, 0xb8, 0x02, 0x17, 0x0a, 0x15, 0x67, 0x65, 0x74, 0x20, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x20, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x42,
	0x09, 0x5a, 0x07, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDescData = file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_goTypes = []interface{}{
	(*GetDashboardCatalogRequest)(nil),  // 0: com.coralogixapis.dashboards.v1.services.GetDashboardCatalogRequest
	(*GetDashboardCatalogResponse)(nil), // 1: com.coralogixapis.dashboards.v1.services.GetDashboardCatalogResponse
	(*DashboardCatalogItem)(nil),        // 2: com.coralogixapis.dashboards.v1.services.DashboardCatalogItem
	(*wrapperspb.StringValue)(nil),      // 3: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),        // 4: google.protobuf.BoolValue
	(*timestamppb.Timestamp)(nil),       // 5: google.protobuf.Timestamp
}
var file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.dashboards.v1.services.GetDashboardCatalogResponse.items:type_name -> com.coralogixapis.dashboards.v1.services.DashboardCatalogItem
	3, // 1: com.coralogixapis.dashboards.v1.services.DashboardCatalogItem.id:type_name -> google.protobuf.StringValue
	3, // 2: com.coralogixapis.dashboards.v1.services.DashboardCatalogItem.name:type_name -> google.protobuf.StringValue
	3, // 3: com.coralogixapis.dashboards.v1.services.DashboardCatalogItem.description:type_name -> google.protobuf.StringValue
	4, // 4: com.coralogixapis.dashboards.v1.services.DashboardCatalogItem.is_default:type_name -> google.protobuf.BoolValue
	4, // 5: com.coralogixapis.dashboards.v1.services.DashboardCatalogItem.is_pinned:type_name -> google.protobuf.BoolValue
	5, // 6: com.coralogixapis.dashboards.v1.services.DashboardCatalogItem.create_time:type_name -> google.protobuf.Timestamp
	5, // 7: com.coralogixapis.dashboards.v1.services.DashboardCatalogItem.update_time:type_name -> google.protobuf.Timestamp
	0, // 8: com.coralogixapis.dashboards.v1.services.DashboardCatalogService.GetDashboardCatalog:input_type -> com.coralogixapis.dashboards.v1.services.GetDashboardCatalogRequest
	1, // 9: com.coralogixapis.dashboards.v1.services.DashboardCatalogService.GetDashboardCatalog:output_type -> com.coralogixapis.dashboards.v1.services.GetDashboardCatalogResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_init() }
func file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_init() {
	if File_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto != nil {
		return
	}
	file_com_coralogixapis_dashboards_v1_audit_log3_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDashboardCatalogRequest); i {
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
		file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDashboardCatalogResponse); i {
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
		file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DashboardCatalogItem); i {
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
			RawDescriptor: file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto = out.File
	file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_services_dashboard_catalog_service_proto_depIdxs = nil
}
