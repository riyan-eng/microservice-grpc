// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: proto/background_service.proto

package pb

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

type BackgroundCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string                       `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Body *BackgroundCreateBodyRequest `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *BackgroundCreateRequest) Reset() {
	*x = BackgroundCreateRequest{}
	mi := &file_proto_background_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackgroundCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackgroundCreateRequest) ProtoMessage() {}

func (x *BackgroundCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_background_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackgroundCreateRequest.ProtoReflect.Descriptor instead.
func (*BackgroundCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_background_service_proto_rawDescGZIP(), []int{0}
}

func (x *BackgroundCreateRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *BackgroundCreateRequest) GetBody() *BackgroundCreateBodyRequest {
	if x != nil {
		return x.Body
	}
	return nil
}

type BackgroundCreateBodyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To   string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	V1   string `protobuf:"bytes,4,opt,name=v1,proto3" json:"v1,omitempty"`
	V2   string `protobuf:"bytes,5,opt,name=v2,proto3" json:"v2,omitempty"`
	V3   string `protobuf:"bytes,6,opt,name=v3,proto3" json:"v3,omitempty"`
	V4   string `protobuf:"bytes,7,opt,name=v4,proto3" json:"v4,omitempty"`
	V5   string `protobuf:"bytes,8,opt,name=v5,proto3" json:"v5,omitempty"`
}

func (x *BackgroundCreateBodyRequest) Reset() {
	*x = BackgroundCreateBodyRequest{}
	mi := &file_proto_background_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackgroundCreateBodyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackgroundCreateBodyRequest) ProtoMessage() {}

func (x *BackgroundCreateBodyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_background_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackgroundCreateBodyRequest.ProtoReflect.Descriptor instead.
func (*BackgroundCreateBodyRequest) Descriptor() ([]byte, []int) {
	return file_proto_background_service_proto_rawDescGZIP(), []int{1}
}

func (x *BackgroundCreateBodyRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *BackgroundCreateBodyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BackgroundCreateBodyRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *BackgroundCreateBodyRequest) GetV1() string {
	if x != nil {
		return x.V1
	}
	return ""
}

func (x *BackgroundCreateBodyRequest) GetV2() string {
	if x != nil {
		return x.V2
	}
	return ""
}

func (x *BackgroundCreateBodyRequest) GetV3() string {
	if x != nil {
		return x.V3
	}
	return ""
}

func (x *BackgroundCreateBodyRequest) GetV4() string {
	if x != nil {
		return x.V4
	}
	return ""
}

func (x *BackgroundCreateBodyRequest) GetV5() string {
	if x != nil {
		return x.V5
	}
	return ""
}

type BackgroundCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BackgroundCreateResponse) Reset() {
	*x = BackgroundCreateResponse{}
	mi := &file_proto_background_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackgroundCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackgroundCreateResponse) ProtoMessage() {}

func (x *BackgroundCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_background_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackgroundCreateResponse.ProtoReflect.Descriptor instead.
func (*BackgroundCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_background_service_proto_rawDescGZIP(), []int{2}
}

var File_proto_background_service_proto protoreflect.FileDescriptor

var file_proto_background_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x17, 0x42, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x64,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xa5,
	0x01, 0x0a, 0x1b, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x31, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x76, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x32, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x76, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x33, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x76, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x34, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x76, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x35, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x76, 0x35, 0x22, 0x1a, 0x0a, 0x18, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x68, 0x0a, 0x11, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x42, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_background_service_proto_rawDescOnce sync.Once
	file_proto_background_service_proto_rawDescData = file_proto_background_service_proto_rawDesc
)

func file_proto_background_service_proto_rawDescGZIP() []byte {
	file_proto_background_service_proto_rawDescOnce.Do(func() {
		file_proto_background_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_background_service_proto_rawDescData)
	})
	return file_proto_background_service_proto_rawDescData
}

var file_proto_background_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_background_service_proto_goTypes = []any{
	(*BackgroundCreateRequest)(nil),     // 0: proto.BackgroundCreateRequest
	(*BackgroundCreateBodyRequest)(nil), // 1: proto.BackgroundCreateBodyRequest
	(*BackgroundCreateResponse)(nil),    // 2: proto.BackgroundCreateResponse
}
var file_proto_background_service_proto_depIdxs = []int32{
	1, // 0: proto.BackgroundCreateRequest.body:type_name -> proto.BackgroundCreateBodyRequest
	0, // 1: proto.BackgroundService.BackgroundCreate:input_type -> proto.BackgroundCreateRequest
	2, // 2: proto.BackgroundService.BackgroundCreate:output_type -> proto.BackgroundCreateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_background_service_proto_init() }
func file_proto_background_service_proto_init() {
	if File_proto_background_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_background_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_background_service_proto_goTypes,
		DependencyIndexes: file_proto_background_service_proto_depIdxs,
		MessageInfos:      file_proto_background_service_proto_msgTypes,
	}.Build()
	File_proto_background_service_proto = out.File
	file_proto_background_service_proto_rawDesc = nil
	file_proto_background_service_proto_goTypes = nil
	file_proto_background_service_proto_depIdxs = nil
}