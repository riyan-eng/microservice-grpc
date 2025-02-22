// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto/task_message.proto

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

type TaskCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Detail string `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *TaskCreateRequest) Reset() {
	*x = TaskCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCreateRequest) ProtoMessage() {}

func (x *TaskCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCreateRequest.ProtoReflect.Descriptor instead.
func (*TaskCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_message_proto_rawDescGZIP(), []int{0}
}

func (x *TaskCreateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TaskCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskCreateRequest) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type TaskCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskCreateResponse) Reset() {
	*x = TaskCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCreateResponse) ProtoMessage() {}

func (x *TaskCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCreateResponse.ProtoReflect.Descriptor instead.
func (*TaskCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_message_proto_rawDescGZIP(), []int{1}
}

type TaskListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *TaskListRequest) Reset() {
	*x = TaskListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListRequest) ProtoMessage() {}

func (x *TaskListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListRequest.ProtoReflect.Descriptor instead.
func (*TaskListRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_message_proto_rawDescGZIP(), []int{2}
}

func (x *TaskListRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *TaskListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *TaskListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type TaskListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []*TaskList `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	TotalRows int32       `protobuf:"varint,2,opt,name=total_rows,json=totalRows,proto3" json:"total_rows,omitempty"`
}

func (x *TaskListResponse) Reset() {
	*x = TaskListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse) ProtoMessage() {}

func (x *TaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResponse.ProtoReflect.Descriptor instead.
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_message_proto_rawDescGZIP(), []int{3}
}

func (x *TaskListResponse) GetData() []*TaskList {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TaskListResponse) GetTotalRows() int32 {
	if x != nil {
		return x.TotalRows
	}
	return 0
}

type TaskDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TaskDetailRequest) Reset() {
	*x = TaskDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDetailRequest) ProtoMessage() {}

func (x *TaskDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDetailRequest.ProtoReflect.Descriptor instead.
func (*TaskDetailRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_message_proto_rawDescGZIP(), []int{4}
}

func (x *TaskDetailRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TaskDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *TaskDetail `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TaskDetailResponse) Reset() {
	*x = TaskDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDetailResponse) ProtoMessage() {}

func (x *TaskDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDetailResponse.ProtoReflect.Descriptor instead.
func (*TaskDetailResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_message_proto_rawDescGZIP(), []int{5}
}

func (x *TaskDetailResponse) GetData() *TaskDetail {
	if x != nil {
		return x.Data
	}
	return nil
}

type TaskPutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Detail string `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *TaskPutRequest) Reset() {
	*x = TaskPutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskPutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskPutRequest) ProtoMessage() {}

func (x *TaskPutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskPutRequest.ProtoReflect.Descriptor instead.
func (*TaskPutRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_message_proto_rawDescGZIP(), []int{6}
}

func (x *TaskPutRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TaskPutRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskPutRequest) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type TaskPutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskPutResponse) Reset() {
	*x = TaskPutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskPutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskPutResponse) ProtoMessage() {}

func (x *TaskPutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskPutResponse.ProtoReflect.Descriptor instead.
func (*TaskPutResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_message_proto_rawDescGZIP(), []int{7}
}

type TaskDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TaskDeleteRequest) Reset() {
	*x = TaskDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDeleteRequest) ProtoMessage() {}

func (x *TaskDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDeleteRequest.ProtoReflect.Descriptor instead.
func (*TaskDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_message_proto_rawDescGZIP(), []int{8}
}

func (x *TaskDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TaskDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskDeleteResponse) Reset() {
	*x = TaskDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDeleteResponse) ProtoMessage() {}

func (x *TaskDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDeleteResponse.ProtoReflect.Descriptor instead.
func (*TaskDeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_message_proto_rawDescGZIP(), []int{9}
}

var File_proto_task_message_proto protoreflect.FileDescriptor

var file_proto_task_message_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f,
	0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22,
	0x14, 0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x56,
	0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x52, 0x6f, 0x77, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x12, 0x54,
	0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4c, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b,
	0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x11, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x54, 0x61, 0x73,
	0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14,
	0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_task_message_proto_rawDescOnce sync.Once
	file_proto_task_message_proto_rawDescData = file_proto_task_message_proto_rawDesc
)

func file_proto_task_message_proto_rawDescGZIP() []byte {
	file_proto_task_message_proto_rawDescOnce.Do(func() {
		file_proto_task_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_task_message_proto_rawDescData)
	})
	return file_proto_task_message_proto_rawDescData
}

var file_proto_task_message_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_task_message_proto_goTypes = []any{
	(*TaskCreateRequest)(nil),  // 0: proto.TaskCreateRequest
	(*TaskCreateResponse)(nil), // 1: proto.TaskCreateResponse
	(*TaskListRequest)(nil),    // 2: proto.TaskListRequest
	(*TaskListResponse)(nil),   // 3: proto.TaskListResponse
	(*TaskDetailRequest)(nil),  // 4: proto.TaskDetailRequest
	(*TaskDetailResponse)(nil), // 5: proto.TaskDetailResponse
	(*TaskPutRequest)(nil),     // 6: proto.TaskPutRequest
	(*TaskPutResponse)(nil),    // 7: proto.TaskPutResponse
	(*TaskDeleteRequest)(nil),  // 8: proto.TaskDeleteRequest
	(*TaskDeleteResponse)(nil), // 9: proto.TaskDeleteResponse
	(*TaskList)(nil),           // 10: proto.TaskList
	(*TaskDetail)(nil),         // 11: proto.TaskDetail
}
var file_proto_task_message_proto_depIdxs = []int32{
	10, // 0: proto.TaskListResponse.data:type_name -> proto.TaskList
	11, // 1: proto.TaskDetailResponse.data:type_name -> proto.TaskDetail
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_task_message_proto_init() }
func file_proto_task_message_proto_init() {
	if File_proto_task_message_proto != nil {
		return
	}
	file_proto_task_datastruct_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_task_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TaskCreateRequest); i {
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
		file_proto_task_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TaskCreateResponse); i {
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
		file_proto_task_message_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TaskListRequest); i {
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
		file_proto_task_message_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TaskListResponse); i {
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
		file_proto_task_message_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*TaskDetailRequest); i {
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
		file_proto_task_message_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TaskDetailResponse); i {
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
		file_proto_task_message_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*TaskPutRequest); i {
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
		file_proto_task_message_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*TaskPutResponse); i {
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
		file_proto_task_message_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*TaskDeleteRequest); i {
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
		file_proto_task_message_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*TaskDeleteResponse); i {
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
			RawDescriptor: file_proto_task_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_task_message_proto_goTypes,
		DependencyIndexes: file_proto_task_message_proto_depIdxs,
		MessageInfos:      file_proto_task_message_proto_msgTypes,
	}.Build()
	File_proto_task_message_proto = out.File
	file_proto_task_message_proto_rawDesc = nil
	file_proto_task_message_proto_goTypes = nil
	file_proto_task_message_proto_depIdxs = nil
}
