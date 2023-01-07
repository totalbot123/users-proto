// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: location-processing.proto

package users_proto

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

// The request message containing the user's name.
type LocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string  `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	X        float32 `protobuf:"fixed32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y        float32 `protobuf:"fixed32,3,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *LocationRequest) Reset() {
	*x = LocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_processing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationRequest) ProtoMessage() {}

func (x *LocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_location_processing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationRequest.ProtoReflect.Descriptor instead.
func (*LocationRequest) Descriptor() ([]byte, []int) {
	return file_location_processing_proto_rawDescGZIP(), []int{0}
}

func (x *LocationRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LocationRequest) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *LocationRequest) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

// The response message containing the greetings
type LocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username          string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	StartDate         string `protobuf:"bytes,2,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate           string `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	DistanceTravelled string `protobuf:"bytes,4,opt,name=distanceTravelled,proto3" json:"distanceTravelled,omitempty"`
}

func (x *LocationResponse) Reset() {
	*x = LocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_processing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationResponse) ProtoMessage() {}

func (x *LocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_location_processing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationResponse.ProtoReflect.Descriptor instead.
func (*LocationResponse) Descriptor() ([]byte, []int) {
	return file_location_processing_proto_rawDescGZIP(), []int{1}
}

func (x *LocationResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LocationResponse) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *LocationResponse) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *LocationResponse) GetDistanceTravelled() string {
	if x != nil {
		return x.DistanceTravelled
	}
	return ""
}

var File_location_processing_proto protoreflect.FileDescriptor

var file_location_processing_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x49, 0x0a, 0x0f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12,
	0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x22, 0x94, 0x01,
	0x0a, 0x10, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x72, 0x61, 0x76, 0x65,
	0x6c, 0x6c, 0x65, 0x64, 0x32, 0x5a, 0x0a, 0x12, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x44, 0x0a, 0x11, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_location_processing_proto_rawDescOnce sync.Once
	file_location_processing_proto_rawDescData = file_location_processing_proto_rawDesc
)

func file_location_processing_proto_rawDescGZIP() []byte {
	file_location_processing_proto_rawDescOnce.Do(func() {
		file_location_processing_proto_rawDescData = protoimpl.X.CompressGZIP(file_location_processing_proto_rawDescData)
	})
	return file_location_processing_proto_rawDescData
}

var file_location_processing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_location_processing_proto_goTypes = []interface{}{
	(*LocationRequest)(nil),  // 0: proto.LocationRequest
	(*LocationResponse)(nil), // 1: proto.LocationResponse
}
var file_location_processing_proto_depIdxs = []int32{
	0, // 0: proto.LocationProcessing.StoreUserLocation:input_type -> proto.LocationRequest
	1, // 1: proto.LocationProcessing.StoreUserLocation:output_type -> proto.LocationResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_location_processing_proto_init() }
func file_location_processing_proto_init() {
	if File_location_processing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_location_processing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationRequest); i {
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
		file_location_processing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationResponse); i {
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
			RawDescriptor: file_location_processing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_location_processing_proto_goTypes,
		DependencyIndexes: file_location_processing_proto_depIdxs,
		MessageInfos:      file_location_processing_proto_msgTypes,
	}.Build()
	File_location_processing_proto = out.File
	file_location_processing_proto_rawDesc = nil
	file_location_processing_proto_goTypes = nil
	file_location_processing_proto_depIdxs = nil
}
