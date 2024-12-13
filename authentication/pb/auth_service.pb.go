// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: proto/auth_service.proto

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

type AuthLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthLoginRequest) Reset() {
	*x = AuthLoginRequest{}
	mi := &file_proto_auth_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLoginRequest) ProtoMessage() {}

func (x *AuthLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLoginRequest.ProtoReflect.Descriptor instead.
func (*AuthLoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *AuthLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken    string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	AccessExpired  string `protobuf:"bytes,2,opt,name=access_expired,json=accessExpired,proto3" json:"access_expired,omitempty"`
	RefreshToken   string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	RefreshExpired string `protobuf:"bytes,4,opt,name=refresh_expired,json=refreshExpired,proto3" json:"refresh_expired,omitempty"`
	Username       string `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AuthLoginResponse) Reset() {
	*x = AuthLoginResponse{}
	mi := &file_proto_auth_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLoginResponse) ProtoMessage() {}

func (x *AuthLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLoginResponse.ProtoReflect.Descriptor instead.
func (*AuthLoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *AuthLoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthLoginResponse) GetAccessExpired() string {
	if x != nil {
		return x.AccessExpired
	}
	return ""
}

func (x *AuthLoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *AuthLoginResponse) GetRefreshExpired() string {
	if x != nil {
		return x.RefreshExpired
	}
	return ""
}

func (x *AuthLoginResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type AuthLogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthLogoutRequest) Reset() {
	*x = AuthLogoutRequest{}
	mi := &file_proto_auth_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthLogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLogoutRequest) ProtoMessage() {}

func (x *AuthLogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLogoutRequest.ProtoReflect.Descriptor instead.
func (*AuthLogoutRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_service_proto_rawDescGZIP(), []int{2}
}

type AuthLogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthLogoutResponse) Reset() {
	*x = AuthLogoutResponse{}
	mi := &file_proto_auth_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthLogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLogoutResponse) ProtoMessage() {}

func (x *AuthLogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLogoutResponse.ProtoReflect.Descriptor instead.
func (*AuthLogoutResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_service_proto_rawDescGZIP(), []int{3}
}

type AuthMeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthMeRequest) Reset() {
	*x = AuthMeRequest{}
	mi := &file_proto_auth_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthMeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthMeRequest) ProtoMessage() {}

func (x *AuthMeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthMeRequest.ProtoReflect.Descriptor instead.
func (*AuthMeRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_service_proto_rawDescGZIP(), []int{4}
}

type AuthMeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	RoleCode    string   `protobuf:"bytes,3,opt,name=role_code,json=roleCode,proto3" json:"role_code,omitempty"`
	RoleName    string   `protobuf:"bytes,4,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	Permissions []string `protobuf:"bytes,5,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *AuthMeResponse) Reset() {
	*x = AuthMeResponse{}
	mi := &file_proto_auth_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthMeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthMeResponse) ProtoMessage() {}

func (x *AuthMeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthMeResponse.ProtoReflect.Descriptor instead.
func (*AuthMeResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_service_proto_rawDescGZIP(), []int{5}
}

func (x *AuthMeResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthMeResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthMeResponse) GetRoleCode() string {
	if x != nil {
		return x.RoleCode
	}
	return ""
}

func (x *AuthMeResponse) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *AuthMeResponse) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type AuthRefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthRefreshRequest) Reset() {
	*x = AuthRefreshRequest{}
	mi := &file_proto_auth_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthRefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRefreshRequest) ProtoMessage() {}

func (x *AuthRefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRefreshRequest.ProtoReflect.Descriptor instead.
func (*AuthRefreshRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_service_proto_rawDescGZIP(), []int{6}
}

func (x *AuthRefreshRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthRefreshResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken    string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	AccessExpired  string `protobuf:"bytes,2,opt,name=access_expired,json=accessExpired,proto3" json:"access_expired,omitempty"`
	RefreshToken   string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	RefreshExpired string `protobuf:"bytes,4,opt,name=refresh_expired,json=refreshExpired,proto3" json:"refresh_expired,omitempty"`
	Username       string `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AuthRefreshResponse) Reset() {
	*x = AuthRefreshResponse{}
	mi := &file_proto_auth_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthRefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRefreshResponse) ProtoMessage() {}

func (x *AuthRefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRefreshResponse.ProtoReflect.Descriptor instead.
func (*AuthRefreshResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_service_proto_rawDescGZIP(), []int{7}
}

func (x *AuthRefreshResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthRefreshResponse) GetAccessExpired() string {
	if x != nil {
		return x.AccessExpired
	}
	return ""
}

func (x *AuthRefreshResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *AuthRefreshResponse) GetRefreshExpired() string {
	if x != nil {
		return x.RefreshExpired
	}
	return ""
}

func (x *AuthRefreshResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type AuthValidateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthValidateTokenRequest) Reset() {
	*x = AuthValidateTokenRequest{}
	mi := &file_proto_auth_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthValidateTokenRequest) ProtoMessage() {}

func (x *AuthValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*AuthValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_service_proto_rawDescGZIP(), []int{8}
}

func (x *AuthValidateTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthValidateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleCode string `protobuf:"bytes,2,opt,name=role_code,json=roleCode,proto3" json:"role_code,omitempty"`
}

func (x *AuthValidateTokenResponse) Reset() {
	*x = AuthValidateTokenResponse{}
	mi := &file_proto_auth_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthValidateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthValidateTokenResponse) ProtoMessage() {}

func (x *AuthValidateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthValidateTokenResponse.ProtoReflect.Descriptor instead.
func (*AuthValidateTokenResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_service_proto_rawDescGZIP(), []int{9}
}

func (x *AuthValidateTokenResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthValidateTokenResponse) GetRoleCode() string {
	if x != nil {
		return x.RoleCode
	}
	return ""
}

type AuthValidatePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FullMethod string `protobuf:"bytes,2,opt,name=full_method,json=fullMethod,proto3" json:"full_method,omitempty"`
}

func (x *AuthValidatePermissionRequest) Reset() {
	*x = AuthValidatePermissionRequest{}
	mi := &file_proto_auth_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthValidatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthValidatePermissionRequest) ProtoMessage() {}

func (x *AuthValidatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthValidatePermissionRequest.ProtoReflect.Descriptor instead.
func (*AuthValidatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_service_proto_rawDescGZIP(), []int{10}
}

func (x *AuthValidatePermissionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthValidatePermissionRequest) GetFullMethod() string {
	if x != nil {
		return x.FullMethod
	}
	return ""
}

type AuthValidatePermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthValidatePermissionResponse) Reset() {
	*x = AuthValidatePermissionResponse{}
	mi := &file_proto_auth_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthValidatePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthValidatePermissionResponse) ProtoMessage() {}

func (x *AuthValidatePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthValidatePermissionResponse.ProtoReflect.Descriptor instead.
func (*AuthValidatePermissionResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_service_proto_rawDescGZIP(), []int{11}
}

var File_proto_auth_service_proto protoreflect.FileDescriptor

var file_proto_auth_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4a, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xc7, 0x01,
	0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12,
	0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x98, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2a,
	0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc9, 0x01, 0x0a, 0x13, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x18, 0x41, 0x75, 0x74, 0x68, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x51, 0x0a, 0x19, 0x41, 0x75, 0x74, 0x68,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x59, 0x0a, 0x1d, 0x41,
	0x75, 0x74, 0x68, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x75, 0x6c, 0x6c,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x20, 0x0a, 0x1e, 0x41, 0x75, 0x74, 0x68, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb4, 0x03, 0x0a, 0x0b, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x02, 0x4d, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x12,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_auth_service_proto_rawDescOnce sync.Once
	file_proto_auth_service_proto_rawDescData = file_proto_auth_service_proto_rawDesc
)

func file_proto_auth_service_proto_rawDescGZIP() []byte {
	file_proto_auth_service_proto_rawDescOnce.Do(func() {
		file_proto_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auth_service_proto_rawDescData)
	})
	return file_proto_auth_service_proto_rawDescData
}

var file_proto_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_auth_service_proto_goTypes = []any{
	(*AuthLoginRequest)(nil),               // 0: proto.AuthLoginRequest
	(*AuthLoginResponse)(nil),              // 1: proto.AuthLoginResponse
	(*AuthLogoutRequest)(nil),              // 2: proto.AuthLogoutRequest
	(*AuthLogoutResponse)(nil),             // 3: proto.AuthLogoutResponse
	(*AuthMeRequest)(nil),                  // 4: proto.AuthMeRequest
	(*AuthMeResponse)(nil),                 // 5: proto.AuthMeResponse
	(*AuthRefreshRequest)(nil),             // 6: proto.AuthRefreshRequest
	(*AuthRefreshResponse)(nil),            // 7: proto.AuthRefreshResponse
	(*AuthValidateTokenRequest)(nil),       // 8: proto.AuthValidateTokenRequest
	(*AuthValidateTokenResponse)(nil),      // 9: proto.AuthValidateTokenResponse
	(*AuthValidatePermissionRequest)(nil),  // 10: proto.AuthValidatePermissionRequest
	(*AuthValidatePermissionResponse)(nil), // 11: proto.AuthValidatePermissionResponse
}
var file_proto_auth_service_proto_depIdxs = []int32{
	0,  // 0: proto.AuthService.Login:input_type -> proto.AuthLoginRequest
	2,  // 1: proto.AuthService.Logout:input_type -> proto.AuthLogoutRequest
	4,  // 2: proto.AuthService.Me:input_type -> proto.AuthMeRequest
	6,  // 3: proto.AuthService.Refresh:input_type -> proto.AuthRefreshRequest
	8,  // 4: proto.AuthService.ValidateToken:input_type -> proto.AuthValidateTokenRequest
	10, // 5: proto.AuthService.ValidatePermission:input_type -> proto.AuthValidatePermissionRequest
	1,  // 6: proto.AuthService.Login:output_type -> proto.AuthLoginResponse
	3,  // 7: proto.AuthService.Logout:output_type -> proto.AuthLogoutResponse
	5,  // 8: proto.AuthService.Me:output_type -> proto.AuthMeResponse
	7,  // 9: proto.AuthService.Refresh:output_type -> proto.AuthRefreshResponse
	9,  // 10: proto.AuthService.ValidateToken:output_type -> proto.AuthValidateTokenResponse
	11, // 11: proto.AuthService.ValidatePermission:output_type -> proto.AuthValidatePermissionResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_auth_service_proto_init() }
func file_proto_auth_service_proto_init() {
	if File_proto_auth_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auth_service_proto_goTypes,
		DependencyIndexes: file_proto_auth_service_proto_depIdxs,
		MessageInfos:      file_proto_auth_service_proto_msgTypes,
	}.Build()
	File_proto_auth_service_proto = out.File
	file_proto_auth_service_proto_rawDesc = nil
	file_proto_auth_service_proto_goTypes = nil
	file_proto_auth_service_proto_depIdxs = nil
}
