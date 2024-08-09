// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: app/config/grpc/to_do_list.proto

package grpc

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

type UserCredentialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	RequestId string `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Token     string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserCredentialsRequest) Reset() {
	*x = UserCredentialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_grpc_to_do_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCredentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCredentialsRequest) ProtoMessage() {}

func (x *UserCredentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_grpc_to_do_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCredentialsRequest.ProtoReflect.Descriptor instead.
func (*UserCredentialsRequest) Descriptor() ([]byte, []int) {
	return file_app_config_grpc_to_do_list_proto_rawDescGZIP(), []int{0}
}

func (x *UserCredentialsRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserCredentialsRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserCredentialsRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *UserCredentialsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_grpc_to_do_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_grpc_to_do_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_app_config_grpc_to_do_list_proto_rawDescGZIP(), []int{1}
}

func (x *UserId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type TaskDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId          string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	TaskMessage     string `protobuf:"bytes,2,opt,name=taskMessage,proto3" json:"taskMessage,omitempty"`
	CreatedAt       string `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	IsTaskCompleted bool   `protobuf:"varint,4,opt,name=isTaskCompleted,proto3" json:"isTaskCompleted,omitempty"`
	UserId          string `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *TaskDomain) Reset() {
	*x = TaskDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_grpc_to_do_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDomain) ProtoMessage() {}

func (x *TaskDomain) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_grpc_to_do_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDomain.ProtoReflect.Descriptor instead.
func (*TaskDomain) Descriptor() ([]byte, []int) {
	return file_app_config_grpc_to_do_list_proto_rawDescGZIP(), []int{2}
}

func (x *TaskDomain) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *TaskDomain) GetTaskMessage() string {
	if x != nil {
		return x.TaskMessage
	}
	return ""
}

func (x *TaskDomain) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *TaskDomain) GetIsTaskCompleted() bool {
	if x != nil {
		return x.IsTaskCompleted
	}
	return false
}

func (x *TaskDomain) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type TaskDomainList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskDomain []*TaskDomain `protobuf:"bytes,1,rep,name=taskDomain,proto3" json:"taskDomain,omitempty"`
}

func (x *TaskDomainList) Reset() {
	*x = TaskDomainList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_grpc_to_do_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDomainList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDomainList) ProtoMessage() {}

func (x *TaskDomainList) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_grpc_to_do_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDomainList.ProtoReflect.Descriptor instead.
func (*TaskDomainList) Descriptor() ([]byte, []int) {
	return file_app_config_grpc_to_do_list_proto_rawDescGZIP(), []int{3}
}

func (x *TaskDomainList) GetTaskDomain() []*TaskDomain {
	if x != nil {
		return x.TaskDomain
	}
	return nil
}

type GetAllTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetAllTasksRequest) Reset() {
	*x = GetAllTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_grpc_to_do_list_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTasksRequest) ProtoMessage() {}

func (x *GetAllTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_grpc_to_do_list_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTasksRequest.ProtoReflect.Descriptor instead.
func (*GetAllTasksRequest) Descriptor() ([]byte, []int) {
	return file_app_config_grpc_to_do_list_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllTasksRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetAllTasksRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *GetAllTasksRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	TaskMessage string `protobuf:"bytes,2,opt,name=taskMessage,proto3" json:"taskMessage,omitempty"`
	RequestId   string `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Token       string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_grpc_to_do_list_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_grpc_to_do_list_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_app_config_grpc_to_do_list_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTaskRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateTaskRequest) GetTaskMessage() string {
	if x != nil {
		return x.TaskMessage
	}
	return ""
}

func (x *CreateTaskRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *CreateTaskRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpdateTaskMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId      string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	TaskMessage string `protobuf:"bytes,2,opt,name=taskMessage,proto3" json:"taskMessage,omitempty"`
	RequestId   string `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Token       string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UpdateTaskMessageRequest) Reset() {
	*x = UpdateTaskMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_grpc_to_do_list_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskMessageRequest) ProtoMessage() {}

func (x *UpdateTaskMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_grpc_to_do_list_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskMessageRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskMessageRequest) Descriptor() ([]byte, []int) {
	return file_app_config_grpc_to_do_list_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTaskMessageRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *UpdateTaskMessageRequest) GetTaskMessage() string {
	if x != nil {
		return x.TaskMessage
	}
	return ""
}

func (x *UpdateTaskMessageRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *UpdateTaskMessageRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpdateTaskCompletenessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UpdateTaskCompletenessRequest) Reset() {
	*x = UpdateTaskCompletenessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_grpc_to_do_list_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskCompletenessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskCompletenessRequest) ProtoMessage() {}

func (x *UpdateTaskCompletenessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_grpc_to_do_list_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskCompletenessRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskCompletenessRequest) Descriptor() ([]byte, []int) {
	return file_app_config_grpc_to_do_list_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTaskCompletenessRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *UpdateTaskCompletenessRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *UpdateTaskCompletenessRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DeleteMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DeleteMessageRequest) Reset() {
	*x = DeleteMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_grpc_to_do_list_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageRequest) ProtoMessage() {}

func (x *DeleteMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_grpc_to_do_list_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageRequest.ProtoReflect.Descriptor instead.
func (*DeleteMessageRequest) Descriptor() ([]byte, []int) {
	return file_app_config_grpc_to_do_list_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteMessageRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *DeleteMessageRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *DeleteMessageRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_grpc_to_do_list_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_grpc_to_do_list_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_app_config_grpc_to_do_list_proto_rawDescGZIP(), []int{9}
}

var File_app_config_grpc_to_do_list_proto protoreflect.FileDescriptor

var file_app_config_grpc_to_do_list_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x84, 0x01, 0x0a, 0x16, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x20, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0xa6, 0x01, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74,
	0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x73, 0x54, 0x61,
	0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x69, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0e, 0x54, 0x61,
	0x73, 0x6b, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0a,
	0x74, 0x61, 0x73, 0x6b, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x60,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x81, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x61, 0x73,
	0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x74, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x6b, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x62, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x32, 0x76, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x1c, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x00,
	0x32, 0xd3, 0x02, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_app_config_grpc_to_do_list_proto_rawDescOnce sync.Once
	file_app_config_grpc_to_do_list_proto_rawDescData = file_app_config_grpc_to_do_list_proto_rawDesc
)

func file_app_config_grpc_to_do_list_proto_rawDescGZIP() []byte {
	file_app_config_grpc_to_do_list_proto_rawDescOnce.Do(func() {
		file_app_config_grpc_to_do_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_config_grpc_to_do_list_proto_rawDescData)
	})
	return file_app_config_grpc_to_do_list_proto_rawDescData
}

var file_app_config_grpc_to_do_list_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_app_config_grpc_to_do_list_proto_goTypes = []interface{}{
	(*UserCredentialsRequest)(nil),        // 0: grpc.UserCredentialsRequest
	(*UserId)(nil),                        // 1: grpc.UserId
	(*TaskDomain)(nil),                    // 2: grpc.TaskDomain
	(*TaskDomainList)(nil),                // 3: grpc.TaskDomainList
	(*GetAllTasksRequest)(nil),            // 4: grpc.GetAllTasksRequest
	(*CreateTaskRequest)(nil),             // 5: grpc.CreateTaskRequest
	(*UpdateTaskMessageRequest)(nil),      // 6: grpc.UpdateTaskMessageRequest
	(*UpdateTaskCompletenessRequest)(nil), // 7: grpc.UpdateTaskCompletenessRequest
	(*DeleteMessageRequest)(nil),          // 8: grpc.DeleteMessageRequest
	(*Void)(nil),                          // 9: grpc.Void
}
var file_app_config_grpc_to_do_list_proto_depIdxs = []int32{
	2, // 0: grpc.TaskDomainList.taskDomain:type_name -> grpc.TaskDomain
	0, // 1: grpc.Account.Signup:input_type -> grpc.UserCredentialsRequest
	0, // 2: grpc.Account.Login:input_type -> grpc.UserCredentialsRequest
	4, // 3: grpc.Task.GetAllTasks:input_type -> grpc.GetAllTasksRequest
	5, // 4: grpc.Task.CreateTask:input_type -> grpc.CreateTaskRequest
	6, // 5: grpc.Task.UpdateTaskMessage:input_type -> grpc.UpdateTaskMessageRequest
	7, // 6: grpc.Task.UpdateTaskCompleteness:input_type -> grpc.UpdateTaskCompletenessRequest
	8, // 7: grpc.Task.DeleteMessage:input_type -> grpc.DeleteMessageRequest
	9, // 8: grpc.Account.Signup:output_type -> grpc.Void
	1, // 9: grpc.Account.Login:output_type -> grpc.UserId
	3, // 10: grpc.Task.GetAllTasks:output_type -> grpc.TaskDomainList
	2, // 11: grpc.Task.CreateTask:output_type -> grpc.TaskDomain
	2, // 12: grpc.Task.UpdateTaskMessage:output_type -> grpc.TaskDomain
	9, // 13: grpc.Task.UpdateTaskCompleteness:output_type -> grpc.Void
	9, // 14: grpc.Task.DeleteMessage:output_type -> grpc.Void
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_config_grpc_to_do_list_proto_init() }
func file_app_config_grpc_to_do_list_proto_init() {
	if File_app_config_grpc_to_do_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_config_grpc_to_do_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCredentialsRequest); i {
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
		file_app_config_grpc_to_do_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_app_config_grpc_to_do_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDomain); i {
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
		file_app_config_grpc_to_do_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDomainList); i {
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
		file_app_config_grpc_to_do_list_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTasksRequest); i {
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
		file_app_config_grpc_to_do_list_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequest); i {
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
		file_app_config_grpc_to_do_list_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskMessageRequest); i {
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
		file_app_config_grpc_to_do_list_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskCompletenessRequest); i {
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
		file_app_config_grpc_to_do_list_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessageRequest); i {
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
		file_app_config_grpc_to_do_list_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_app_config_grpc_to_do_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_app_config_grpc_to_do_list_proto_goTypes,
		DependencyIndexes: file_app_config_grpc_to_do_list_proto_depIdxs,
		MessageInfos:      file_app_config_grpc_to_do_list_proto_msgTypes,
	}.Build()
	File_app_config_grpc_to_do_list_proto = out.File
	file_app_config_grpc_to_do_list_proto_rawDesc = nil
	file_app_config_grpc_to_do_list_proto_goTypes = nil
	file_app_config_grpc_to_do_list_proto_depIdxs = nil
}
