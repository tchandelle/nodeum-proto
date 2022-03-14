// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: protobuf/scheduler.micro.proto

package schedulerpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Schedule_Status int32

const (
	// Waiting for the scheduler to find the next schedule
	Schedule_PENDING Schedule_Status = 0
	// Next scheduled has been defined
	Schedule_SCHEDULED Schedule_Status = 1
	// Schedule is finished, no more occurences
	Schedule_DONE Schedule_Status = 2
)

// Enum value maps for Schedule_Status.
var (
	Schedule_Status_name = map[int32]string{
		0: "PENDING",
		1: "SCHEDULED",
		2: "DONE",
	}
	Schedule_Status_value = map[string]int32{
		"PENDING":   0,
		"SCHEDULED": 1,
		"DONE":      2,
	}
)

func (x Schedule_Status) Enum() *Schedule_Status {
	p := new(Schedule_Status)
	*p = x
	return p
}

func (x Schedule_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Schedule_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_scheduler_micro_proto_enumTypes[0].Descriptor()
}

func (Schedule_Status) Type() protoreflect.EnumType {
	return &file_protobuf_scheduler_micro_proto_enumTypes[0]
}

func (x Schedule_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Schedule_Status.Descriptor instead.
func (Schedule_Status) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_scheduler_micro_proto_rawDescGZIP(), []int{1, 0}
}

type ScheduleIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,2,opt,name=task_id,proto3" json:"task_id,omitempty"`
}

func (x *ScheduleIdRequest) Reset() {
	*x = ScheduleIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_scheduler_micro_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleIdRequest) ProtoMessage() {}

func (x *ScheduleIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_scheduler_micro_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleIdRequest.ProtoReflect.Descriptor instead.
func (*ScheduleIdRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_scheduler_micro_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleIdRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string                 `protobuf:"bytes,2,opt,name=task_id,proto3" json:"task_id,omitempty"`
	Rrule  string                 `protobuf:"bytes,4,opt,name=rrule,proto3" json:"rrule,omitempty"`
	Next   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=next,proto3" json:"next,omitempty"`
	Status Schedule_Status        `protobuf:"varint,6,opt,name=status,proto3,enum=nodeum.protobuf.scheduler.Schedule_Status" json:"status,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_scheduler_micro_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_scheduler_micro_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_protobuf_scheduler_micro_proto_rawDescGZIP(), []int{1}
}

func (x *Schedule) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Schedule) GetRrule() string {
	if x != nil {
		return x.Rrule
	}
	return ""
}

func (x *Schedule) GetNext() *timestamppb.Timestamp {
	if x != nil {
		return x.Next
	}
	return nil
}

func (x *Schedule) GetStatus() Schedule_Status {
	if x != nil {
		return x.Status
	}
	return Schedule_PENDING
}

type ReadScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schedule *Schedule `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty"`
}

func (x *ReadScheduleResponse) Reset() {
	*x = ReadScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_scheduler_micro_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadScheduleResponse) ProtoMessage() {}

func (x *ReadScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_scheduler_micro_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadScheduleResponse.ProtoReflect.Descriptor instead.
func (*ReadScheduleResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_scheduler_micro_proto_rawDescGZIP(), []int{2}
}

func (x *ReadScheduleResponse) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

type ListSchedulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schedules []*Schedule `protobuf:"bytes,1,rep,name=schedules,proto3" json:"schedules,omitempty"`
}

func (x *ListSchedulesResponse) Reset() {
	*x = ListSchedulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_scheduler_micro_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSchedulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSchedulesResponse) ProtoMessage() {}

func (x *ListSchedulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_scheduler_micro_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSchedulesResponse.ProtoReflect.Descriptor instead.
func (*ListSchedulesResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_scheduler_micro_proto_rawDescGZIP(), []int{3}
}

func (x *ListSchedulesResponse) GetSchedules() []*Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

type WriteScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schedule *Schedule `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty"`
}

func (x *WriteScheduleRequest) Reset() {
	*x = WriteScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_scheduler_micro_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteScheduleRequest) ProtoMessage() {}

func (x *WriteScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_scheduler_micro_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteScheduleRequest.ProtoReflect.Descriptor instead.
func (*WriteScheduleRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_scheduler_micro_proto_rawDescGZIP(), []int{4}
}

func (x *WriteScheduleRequest) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

type ScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ScheduleResponse) Reset() {
	*x = ScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_scheduler_micro_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleResponse) ProtoMessage() {}

func (x *ScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_scheduler_micro_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleResponse.ProtoReflect.Descriptor instead.
func (*ScheduleResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_scheduler_micro_proto_rawDescGZIP(), []int{5}
}

func (x *ScheduleResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_protobuf_scheduler_micro_proto protoreflect.FileDescriptor

var file_protobuf_scheduler_micro_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x11, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x22, 0xde, 0x01, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x72,
	0x75, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x6e,
	0x65, 0x78, 0x74, 0x12, 0x42, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x22, 0x57, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x22, 0x5a, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x14,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x24, 0x0a, 0x10, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x9b, 0x04, 0x0a, 0x09,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x8b, 0x01, 0x0a, 0x0c, 0x52, 0x65,
	0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2c, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x12, 0x14, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6d, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x30, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2f, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x89, 0x01,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x12, 0x2c, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f,
	0x7b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x69,
	0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protobuf_scheduler_micro_proto_rawDescOnce sync.Once
	file_protobuf_scheduler_micro_proto_rawDescData = file_protobuf_scheduler_micro_proto_rawDesc
)

func file_protobuf_scheduler_micro_proto_rawDescGZIP() []byte {
	file_protobuf_scheduler_micro_proto_rawDescOnce.Do(func() {
		file_protobuf_scheduler_micro_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_scheduler_micro_proto_rawDescData)
	})
	return file_protobuf_scheduler_micro_proto_rawDescData
}

var file_protobuf_scheduler_micro_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_scheduler_micro_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protobuf_scheduler_micro_proto_goTypes = []interface{}{
	(Schedule_Status)(0),          // 0: nodeum.protobuf.scheduler.Schedule.Status
	(*ScheduleIdRequest)(nil),     // 1: nodeum.protobuf.scheduler.ScheduleIdRequest
	(*Schedule)(nil),              // 2: nodeum.protobuf.scheduler.Schedule
	(*ReadScheduleResponse)(nil),  // 3: nodeum.protobuf.scheduler.ReadScheduleResponse
	(*ListSchedulesResponse)(nil), // 4: nodeum.protobuf.scheduler.ListSchedulesResponse
	(*WriteScheduleRequest)(nil),  // 5: nodeum.protobuf.scheduler.WriteScheduleRequest
	(*ScheduleResponse)(nil),      // 6: nodeum.protobuf.scheduler.ScheduleResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_protobuf_scheduler_micro_proto_depIdxs = []int32{
	7, // 0: nodeum.protobuf.scheduler.Schedule.next:type_name -> google.protobuf.Timestamp
	0, // 1: nodeum.protobuf.scheduler.Schedule.status:type_name -> nodeum.protobuf.scheduler.Schedule.Status
	2, // 2: nodeum.protobuf.scheduler.ReadScheduleResponse.schedule:type_name -> nodeum.protobuf.scheduler.Schedule
	2, // 3: nodeum.protobuf.scheduler.ListSchedulesResponse.schedules:type_name -> nodeum.protobuf.scheduler.Schedule
	2, // 4: nodeum.protobuf.scheduler.WriteScheduleRequest.schedule:type_name -> nodeum.protobuf.scheduler.Schedule
	1, // 5: nodeum.protobuf.scheduler.Scheduler.ReadSchedule:input_type -> nodeum.protobuf.scheduler.ScheduleIdRequest
	8, // 6: nodeum.protobuf.scheduler.Scheduler.ListSchedules:input_type -> google.protobuf.Empty
	5, // 7: nodeum.protobuf.scheduler.Scheduler.WriteSchedule:input_type -> nodeum.protobuf.scheduler.WriteScheduleRequest
	1, // 8: nodeum.protobuf.scheduler.Scheduler.DeleteSchedule:input_type -> nodeum.protobuf.scheduler.ScheduleIdRequest
	3, // 9: nodeum.protobuf.scheduler.Scheduler.ReadSchedule:output_type -> nodeum.protobuf.scheduler.ReadScheduleResponse
	4, // 10: nodeum.protobuf.scheduler.Scheduler.ListSchedules:output_type -> nodeum.protobuf.scheduler.ListSchedulesResponse
	6, // 11: nodeum.protobuf.scheduler.Scheduler.WriteSchedule:output_type -> nodeum.protobuf.scheduler.ScheduleResponse
	6, // 12: nodeum.protobuf.scheduler.Scheduler.DeleteSchedule:output_type -> nodeum.protobuf.scheduler.ScheduleResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protobuf_scheduler_micro_proto_init() }
func file_protobuf_scheduler_micro_proto_init() {
	if File_protobuf_scheduler_micro_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_scheduler_micro_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleIdRequest); i {
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
		file_protobuf_scheduler_micro_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
		file_protobuf_scheduler_micro_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadScheduleResponse); i {
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
		file_protobuf_scheduler_micro_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSchedulesResponse); i {
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
		file_protobuf_scheduler_micro_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteScheduleRequest); i {
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
		file_protobuf_scheduler_micro_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleResponse); i {
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
			RawDescriptor: file_protobuf_scheduler_micro_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_scheduler_micro_proto_goTypes,
		DependencyIndexes: file_protobuf_scheduler_micro_proto_depIdxs,
		EnumInfos:         file_protobuf_scheduler_micro_proto_enumTypes,
		MessageInfos:      file_protobuf_scheduler_micro_proto_msgTypes,
	}.Build()
	File_protobuf_scheduler_micro_proto = out.File
	file_protobuf_scheduler_micro_proto_rawDesc = nil
	file_protobuf_scheduler_micro_proto_goTypes = nil
	file_protobuf_scheduler_micro_proto_depIdxs = nil
}
