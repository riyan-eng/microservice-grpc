// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: proto/permission_service.proto

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

type PermissionRoleAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PermissionRoleAccessRequest) Reset() {
	*x = PermissionRoleAccessRequest{}
	mi := &file_proto_permission_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionRoleAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionRoleAccessRequest) ProtoMessage() {}

func (x *PermissionRoleAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_permission_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionRoleAccessRequest.ProtoReflect.Descriptor instead.
func (*PermissionRoleAccessRequest) Descriptor() ([]byte, []int) {
	return file_proto_permission_service_proto_rawDescGZIP(), []int{0}
}

type PermissionRoleAccessListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*PermissionRoleAccessResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *PermissionRoleAccessListResponse) Reset() {
	*x = PermissionRoleAccessListResponse{}
	mi := &file_proto_permission_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionRoleAccessListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionRoleAccessListResponse) ProtoMessage() {}

func (x *PermissionRoleAccessListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_permission_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionRoleAccessListResponse.ProtoReflect.Descriptor instead.
func (*PermissionRoleAccessListResponse) Descriptor() ([]byte, []int) {
	return file_proto_permission_service_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionRoleAccessListResponse) GetData() []*PermissionRoleAccessResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type PermissionRoleAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Permissions []*PermissionRoleAccessChildResponse `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *PermissionRoleAccessResponse) Reset() {
	*x = PermissionRoleAccessResponse{}
	mi := &file_proto_permission_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionRoleAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionRoleAccessResponse) ProtoMessage() {}

func (x *PermissionRoleAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_permission_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionRoleAccessResponse.ProtoReflect.Descriptor instead.
func (*PermissionRoleAccessResponse) Descriptor() ([]byte, []int) {
	return file_proto_permission_service_proto_rawDescGZIP(), []int{2}
}

func (x *PermissionRoleAccessResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PermissionRoleAccessResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionRoleAccessResponse) GetPermissions() []*PermissionRoleAccessChildResponse {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type PermissionRoleAccessChildResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code       string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Parent     string `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	FullMethod string `protobuf:"bytes,5,opt,name=full_method,json=fullMethod,proto3" json:"full_method,omitempty"`
}

func (x *PermissionRoleAccessChildResponse) Reset() {
	*x = PermissionRoleAccessChildResponse{}
	mi := &file_proto_permission_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionRoleAccessChildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionRoleAccessChildResponse) ProtoMessage() {}

func (x *PermissionRoleAccessChildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_permission_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionRoleAccessChildResponse.ProtoReflect.Descriptor instead.
func (*PermissionRoleAccessChildResponse) Descriptor() ([]byte, []int) {
	return file_proto_permission_service_proto_rawDescGZIP(), []int{3}
}

func (x *PermissionRoleAccessChildResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PermissionRoleAccessChildResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PermissionRoleAccessChildResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionRoleAccessChildResponse) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *PermissionRoleAccessChildResponse) GetFullMethod() string {
	if x != nil {
		return x.FullMethod
	}
	return ""
}

var File_proto_permission_service_proto protoreflect.FileDescriptor

var file_proto_permission_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x1b, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5b, 0x0a, 0x20, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x8e, 0x01, 0x0a, 0x1c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x6f, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x21, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x75,
	0x6c, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x66, 0x75, 0x6c, 0x6c, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x32, 0x72, 0x0a, 0x11, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5d, 0x0a, 0x0e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_permission_service_proto_rawDescOnce sync.Once
	file_proto_permission_service_proto_rawDescData = file_proto_permission_service_proto_rawDesc
)

func file_proto_permission_service_proto_rawDescGZIP() []byte {
	file_proto_permission_service_proto_rawDescOnce.Do(func() {
		file_proto_permission_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_permission_service_proto_rawDescData)
	})
	return file_proto_permission_service_proto_rawDescData
}

var file_proto_permission_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_permission_service_proto_goTypes = []any{
	(*PermissionRoleAccessRequest)(nil),       // 0: proto.PermissionRoleAccessRequest
	(*PermissionRoleAccessListResponse)(nil),  // 1: proto.PermissionRoleAccessListResponse
	(*PermissionRoleAccessResponse)(nil),      // 2: proto.PermissionRoleAccessResponse
	(*PermissionRoleAccessChildResponse)(nil), // 3: proto.PermissionRoleAccessChildResponse
}
var file_proto_permission_service_proto_depIdxs = []int32{
	2, // 0: proto.PermissionRoleAccessListResponse.data:type_name -> proto.PermissionRoleAccessResponse
	3, // 1: proto.PermissionRoleAccessResponse.permissions:type_name -> proto.PermissionRoleAccessChildResponse
	0, // 2: proto.PermissionService.RoleAccessList:input_type -> proto.PermissionRoleAccessRequest
	1, // 3: proto.PermissionService.RoleAccessList:output_type -> proto.PermissionRoleAccessListResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_permission_service_proto_init() }
func file_proto_permission_service_proto_init() {
	if File_proto_permission_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_permission_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_permission_service_proto_goTypes,
		DependencyIndexes: file_proto_permission_service_proto_depIdxs,
		MessageInfos:      file_proto_permission_service_proto_msgTypes,
	}.Build()
	File_proto_permission_service_proto = out.File
	file_proto_permission_service_proto_rawDesc = nil
	file_proto_permission_service_proto_goTypes = nil
	file_proto_permission_service_proto_depIdxs = nil
}
