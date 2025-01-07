// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: metastore/v1/query.proto

package metastorev1

import (
	v1 "github.com/grafana/pyroscope/api/gen/proto/go/types/v1"
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

type QueryMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId  []string `protobuf:"bytes,1,rep,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	StartTime int64    `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64    `protobuf:"varint,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Query     string   `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *QueryMetadataRequest) Reset() {
	*x = QueryMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMetadataRequest) ProtoMessage() {}

func (x *QueryMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMetadataRequest.ProtoReflect.Descriptor instead.
func (*QueryMetadataRequest) Descriptor() ([]byte, []int) {
	return file_metastore_v1_query_proto_rawDescGZIP(), []int{0}
}

func (x *QueryMetadataRequest) GetTenantId() []string {
	if x != nil {
		return x.TenantId
	}
	return nil
}

func (x *QueryMetadataRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *QueryMetadataRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *QueryMetadataRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type QueryMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*BlockMeta `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *QueryMetadataResponse) Reset() {
	*x = QueryMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMetadataResponse) ProtoMessage() {}

func (x *QueryMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMetadataResponse.ProtoReflect.Descriptor instead.
func (*QueryMetadataResponse) Descriptor() ([]byte, []int) {
	return file_metastore_v1_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryMetadataResponse) GetBlocks() []*BlockMeta {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type QueryMetadataLabelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId  []string `protobuf:"bytes,1,rep,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	StartTime int64    `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64    `protobuf:"varint,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Query     string   `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	Labels    []string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *QueryMetadataLabelsRequest) Reset() {
	*x = QueryMetadataLabelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMetadataLabelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMetadataLabelsRequest) ProtoMessage() {}

func (x *QueryMetadataLabelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMetadataLabelsRequest.ProtoReflect.Descriptor instead.
func (*QueryMetadataLabelsRequest) Descriptor() ([]byte, []int) {
	return file_metastore_v1_query_proto_rawDescGZIP(), []int{2}
}

func (x *QueryMetadataLabelsRequest) GetTenantId() []string {
	if x != nil {
		return x.TenantId
	}
	return nil
}

func (x *QueryMetadataLabelsRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *QueryMetadataLabelsRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *QueryMetadataLabelsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *QueryMetadataLabelsRequest) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type QueryMetadataLabelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels []*v1.Labels `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *QueryMetadataLabelsResponse) Reset() {
	*x = QueryMetadataLabelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMetadataLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMetadataLabelsResponse) ProtoMessage() {}

func (x *QueryMetadataLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMetadataLabelsResponse.ProtoReflect.Descriptor instead.
func (*QueryMetadataLabelsResponse) Descriptor() ([]byte, []int) {
	return file_metastore_v1_query_proto_rawDescGZIP(), []int{3}
}

func (x *QueryMetadataLabelsResponse) GetLabels() []*v1.Labels {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_metastore_v1_query_proto protoreflect.FileDescriptor

var file_metastore_v1_query_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x74, 0x61,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x14, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x48,
	0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0xa1, 0x01, 0x0a, 0x1a, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x47, 0x0a, 0x1b,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x32, 0xe0, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a,
	0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x22, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x13, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x12, 0x28, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb7, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f,
	0x70, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x4d, 0x65, 0x74, 0x61,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metastore_v1_query_proto_rawDescOnce sync.Once
	file_metastore_v1_query_proto_rawDescData = file_metastore_v1_query_proto_rawDesc
)

func file_metastore_v1_query_proto_rawDescGZIP() []byte {
	file_metastore_v1_query_proto_rawDescOnce.Do(func() {
		file_metastore_v1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_metastore_v1_query_proto_rawDescData)
	})
	return file_metastore_v1_query_proto_rawDescData
}

var file_metastore_v1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_metastore_v1_query_proto_goTypes = []any{
	(*QueryMetadataRequest)(nil),        // 0: metastore.v1.QueryMetadataRequest
	(*QueryMetadataResponse)(nil),       // 1: metastore.v1.QueryMetadataResponse
	(*QueryMetadataLabelsRequest)(nil),  // 2: metastore.v1.QueryMetadataLabelsRequest
	(*QueryMetadataLabelsResponse)(nil), // 3: metastore.v1.QueryMetadataLabelsResponse
	(*BlockMeta)(nil),                   // 4: metastore.v1.BlockMeta
	(*v1.Labels)(nil),                   // 5: types.v1.Labels
}
var file_metastore_v1_query_proto_depIdxs = []int32{
	4, // 0: metastore.v1.QueryMetadataResponse.blocks:type_name -> metastore.v1.BlockMeta
	5, // 1: metastore.v1.QueryMetadataLabelsResponse.labels:type_name -> types.v1.Labels
	0, // 2: metastore.v1.MetadataQueryService.QueryMetadata:input_type -> metastore.v1.QueryMetadataRequest
	2, // 3: metastore.v1.MetadataQueryService.QueryMetadataLabels:input_type -> metastore.v1.QueryMetadataLabelsRequest
	1, // 4: metastore.v1.MetadataQueryService.QueryMetadata:output_type -> metastore.v1.QueryMetadataResponse
	3, // 5: metastore.v1.MetadataQueryService.QueryMetadataLabels:output_type -> metastore.v1.QueryMetadataLabelsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_metastore_v1_query_proto_init() }
func file_metastore_v1_query_proto_init() {
	if File_metastore_v1_query_proto != nil {
		return
	}
	file_metastore_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_metastore_v1_query_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*QueryMetadataRequest); i {
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
		file_metastore_v1_query_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*QueryMetadataResponse); i {
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
		file_metastore_v1_query_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*QueryMetadataLabelsRequest); i {
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
		file_metastore_v1_query_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*QueryMetadataLabelsResponse); i {
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
			RawDescriptor: file_metastore_v1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metastore_v1_query_proto_goTypes,
		DependencyIndexes: file_metastore_v1_query_proto_depIdxs,
		MessageInfos:      file_metastore_v1_query_proto_msgTypes,
	}.Build()
	File_metastore_v1_query_proto = out.File
	file_metastore_v1_query_proto_rawDesc = nil
	file_metastore_v1_query_proto_goTypes = nil
	file_metastore_v1_query_proto_depIdxs = nil
}