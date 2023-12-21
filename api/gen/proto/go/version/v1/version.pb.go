// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: version/v1/version.proto

package versionv1

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

type VersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VersionRequest) Reset() {
	*x = VersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_version_v1_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionRequest) ProtoMessage() {}

func (x *VersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_version_v1_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionRequest.ProtoReflect.Descriptor instead.
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return file_version_v1_version_proto_rawDescGZIP(), []int{0}
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuerierAPI uint64 `protobuf:"varint,1,opt,name=QuerierAPI,proto3" json:"QuerierAPI,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_version_v1_version_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_version_v1_version_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_version_v1_version_proto_rawDescGZIP(), []int{1}
}

func (x *VersionResponse) GetQuerierAPI() uint64 {
	if x != nil {
		return x.QuerierAPI
	}
	return 0
}

type InstanceVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	// Unix timestamp (with nanoseconds precision) of the last heartbeat sent
	// by this instance.
	Timestamp int64 `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	// Querier Service API version
	QuerierAPI uint64 `protobuf:"varint,4,opt,name=QuerierAPI,proto3" json:"QuerierAPI,omitempty"`
	// Tells if the instance is running or has left cluster.
	Left bool `protobuf:"varint,5,opt,name=left,proto3" json:"left,omitempty"`
}

func (x *InstanceVersion) Reset() {
	*x = InstanceVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_version_v1_version_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceVersion) ProtoMessage() {}

func (x *InstanceVersion) ProtoReflect() protoreflect.Message {
	mi := &file_version_v1_version_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceVersion.ProtoReflect.Descriptor instead.
func (*InstanceVersion) Descriptor() ([]byte, []int) {
	return file_version_v1_version_proto_rawDescGZIP(), []int{2}
}

func (x *InstanceVersion) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *InstanceVersion) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *InstanceVersion) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *InstanceVersion) GetQuerierAPI() uint64 {
	if x != nil {
		return x.QuerierAPI
	}
	return 0
}

func (x *InstanceVersion) GetLeft() bool {
	if x != nil {
		return x.Left
	}
	return false
}

// Versions is the top-level type used to model version for all instances, containing information for individual instances.
type Versions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instances map[string]*InstanceVersion `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Versions) Reset() {
	*x = Versions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_version_v1_version_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Versions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Versions) ProtoMessage() {}

func (x *Versions) ProtoReflect() protoreflect.Message {
	mi := &file_version_v1_version_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Versions.ProtoReflect.Descriptor instead.
func (*Versions) Descriptor() ([]byte, []int) {
	return file_version_v1_version_proto_rawDescGZIP(), []int{3}
}

func (x *Versions) GetInstances() map[string]*InstanceVersion {
	if x != nil {
		return x.Instances
	}
	return nil
}

var File_version_v1_version_proto protoreflect.FileDescriptor

var file_version_v1_version_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x10, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x51,
	0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x41, 0x50, 0x49, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x41, 0x50, 0x49, 0x22, 0x87, 0x01, 0x0a, 0x0f,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x41, 0x50, 0x49, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x41, 0x50,
	0x49, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x6c, 0x65, 0x66, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x08, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x41, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x59, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x32, 0x4f, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x07, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0xab, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02,
	0x0a, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_version_v1_version_proto_rawDescOnce sync.Once
	file_version_v1_version_proto_rawDescData = file_version_v1_version_proto_rawDesc
)

func file_version_v1_version_proto_rawDescGZIP() []byte {
	file_version_v1_version_proto_rawDescOnce.Do(func() {
		file_version_v1_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_version_v1_version_proto_rawDescData)
	})
	return file_version_v1_version_proto_rawDescData
}

var file_version_v1_version_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_version_v1_version_proto_goTypes = []interface{}{
	(*VersionRequest)(nil),  // 0: version.v1.VersionRequest
	(*VersionResponse)(nil), // 1: version.v1.VersionResponse
	(*InstanceVersion)(nil), // 2: version.v1.InstanceVersion
	(*Versions)(nil),        // 3: version.v1.Versions
	nil,                     // 4: version.v1.Versions.InstancesEntry
}
var file_version_v1_version_proto_depIdxs = []int32{
	4, // 0: version.v1.Versions.instances:type_name -> version.v1.Versions.InstancesEntry
	2, // 1: version.v1.Versions.InstancesEntry.value:type_name -> version.v1.InstanceVersion
	0, // 2: version.v1.Version.Version:input_type -> version.v1.VersionRequest
	1, // 3: version.v1.Version.Version:output_type -> version.v1.VersionResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_version_v1_version_proto_init() }
func file_version_v1_version_proto_init() {
	if File_version_v1_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_version_v1_version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionRequest); i {
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
		file_version_v1_version_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_version_v1_version_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceVersion); i {
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
		file_version_v1_version_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Versions); i {
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
			RawDescriptor: file_version_v1_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_version_v1_version_proto_goTypes,
		DependencyIndexes: file_version_v1_version_proto_depIdxs,
		MessageInfos:      file_version_v1_version_proto_msgTypes,
	}.Build()
	File_version_v1_version_proto = out.File
	file_version_v1_version_proto_rawDesc = nil
	file_version_v1_version_proto_goTypes = nil
	file_version_v1_version_proto_depIdxs = nil
}