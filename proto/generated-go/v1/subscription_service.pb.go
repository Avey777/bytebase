// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: v1/subscription_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
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

type PlanType int32

const (
	PlanType_PLAN_TYPE_UNSPECIFIED PlanType = 0
	PlanType_FREE                  PlanType = 1
	PlanType_TEAM                  PlanType = 2
	PlanType_ENTERPRISE            PlanType = 3
)

// Enum value maps for PlanType.
var (
	PlanType_name = map[int32]string{
		0: "PLAN_TYPE_UNSPECIFIED",
		1: "FREE",
		2: "TEAM",
		3: "ENTERPRISE",
	}
	PlanType_value = map[string]int32{
		"PLAN_TYPE_UNSPECIFIED": 0,
		"FREE":                  1,
		"TEAM":                  2,
		"ENTERPRISE":            3,
	}
)

func (x PlanType) Enum() *PlanType {
	p := new(PlanType)
	*p = x
	return p
}

func (x PlanType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_subscription_service_proto_enumTypes[0].Descriptor()
}

func (PlanType) Type() protoreflect.EnumType {
	return &file_v1_subscription_service_proto_enumTypes[0]
}

func (x PlanType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanType.Descriptor instead.
func (PlanType) EnumDescriptor() ([]byte, []int) {
	return file_v1_subscription_service_proto_rawDescGZIP(), []int{0}
}

type GetSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSubscriptionRequest) Reset() {
	*x = GetSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_subscription_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionRequest) ProtoMessage() {}

func (x *GetSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_subscription_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_v1_subscription_service_proto_rawDescGZIP(), []int{0}
}

type UpdateSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patch *PatchSubscription `protobuf:"bytes,1,opt,name=patch,proto3" json:"patch,omitempty"`
}

func (x *UpdateSubscriptionRequest) Reset() {
	*x = UpdateSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_subscription_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSubscriptionRequest) ProtoMessage() {}

func (x *UpdateSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_subscription_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*UpdateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_v1_subscription_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateSubscriptionRequest) GetPatch() *PatchSubscription {
	if x != nil {
		return x.Patch
	}
	return nil
}

type TrialSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trial *TrialSubscription `protobuf:"bytes,1,opt,name=trial,proto3" json:"trial,omitempty"`
}

func (x *TrialSubscriptionRequest) Reset() {
	*x = TrialSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_subscription_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrialSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrialSubscriptionRequest) ProtoMessage() {}

func (x *TrialSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_subscription_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrialSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*TrialSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_v1_subscription_service_proto_rawDescGZIP(), []int{2}
}

func (x *TrialSubscriptionRequest) GetTrial() *TrialSubscription {
	if x != nil {
		return x.Trial
	}
	return nil
}

type PatchSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	License string `protobuf:"bytes,1,opt,name=license,proto3" json:"license,omitempty"`
}

func (x *PatchSubscription) Reset() {
	*x = PatchSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_subscription_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchSubscription) ProtoMessage() {}

func (x *PatchSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_v1_subscription_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchSubscription.ProtoReflect.Descriptor instead.
func (*PatchSubscription) Descriptor() ([]byte, []int) {
	return file_v1_subscription_service_proto_rawDescGZIP(), []int{3}
}

func (x *PatchSubscription) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

type TrialSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plan          PlanType `protobuf:"varint,1,opt,name=plan,proto3,enum=bytebase.v1.PlanType" json:"plan,omitempty"`
	Days          int32    `protobuf:"varint,2,opt,name=days,proto3" json:"days,omitempty"`
	InstanceCount int32    `protobuf:"varint,4,opt,name=instance_count,json=instanceCount,proto3" json:"instance_count,omitempty"`
}

func (x *TrialSubscription) Reset() {
	*x = TrialSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_subscription_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrialSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrialSubscription) ProtoMessage() {}

func (x *TrialSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_v1_subscription_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrialSubscription.ProtoReflect.Descriptor instead.
func (*TrialSubscription) Descriptor() ([]byte, []int) {
	return file_v1_subscription_service_proto_rawDescGZIP(), []int{4}
}

func (x *TrialSubscription) GetPlan() PlanType {
	if x != nil {
		return x.Plan
	}
	return PlanType_PLAN_TYPE_UNSPECIFIED
}

func (x *TrialSubscription) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *TrialSubscription) GetInstanceCount() int32 {
	if x != nil {
		return x.InstanceCount
	}
	return 0
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceCount int32                  `protobuf:"varint,2,opt,name=instance_count,json=instanceCount,proto3" json:"instance_count,omitempty"`
	ExpiresTime   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expires_time,json=expiresTime,proto3" json:"expires_time,omitempty"`
	StartedTime   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=started_time,json=startedTime,proto3" json:"started_time,omitempty"`
	Plan          PlanType               `protobuf:"varint,5,opt,name=plan,proto3,enum=bytebase.v1.PlanType" json:"plan,omitempty"`
	Trialing      bool                   `protobuf:"varint,6,opt,name=trialing,proto3" json:"trialing,omitempty"`
	OrgId         string                 `protobuf:"bytes,7,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	OrgName       string                 `protobuf:"bytes,8,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_subscription_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_v1_subscription_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_v1_subscription_service_proto_rawDescGZIP(), []int{5}
}

func (x *Subscription) GetInstanceCount() int32 {
	if x != nil {
		return x.InstanceCount
	}
	return 0
}

func (x *Subscription) GetExpiresTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresTime
	}
	return nil
}

func (x *Subscription) GetStartedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedTime
	}
	return nil
}

func (x *Subscription) GetPlan() PlanType {
	if x != nil {
		return x.Plan
	}
	return PlanType_PLAN_TYPE_UNSPECIFIED
}

func (x *Subscription) GetTrialing() bool {
	if x != nil {
		return x.Trialing
	}
	return false
}

func (x *Subscription) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Subscription) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

var File_v1_subscription_service_proto protoreflect.FileDescriptor

var file_v1_subscription_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x19,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x22,
	0x50, 0x0a, 0x18, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x74,
	0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x74, 0x72, 0x69, 0x61,
	0x6c, 0x22, 0x2d, 0x0a, 0x11, 0x50, 0x61, 0x74, 0x63, 0x68, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x22, 0x79, 0x0a, 0x11, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x64, 0x61, 0x79, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xcf, 0x02, 0x0a, 0x0c,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x0e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e,
	0x12, 0x1f, 0x0a, 0x08, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x6e,
	0x67, 0x12, 0x1a, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x08, 0x6f, 0x72, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x49, 0x0a,
	0x08, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x4c, 0x41,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x52, 0x45, 0x45, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x54, 0x45, 0x41, 0x4d, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x4e, 0x54, 0x45,
	0x52, 0x50, 0x52, 0x49, 0x53, 0x45, 0x10, 0x03, 0x32, 0x8f, 0x03, 0x0a, 0x13, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x1b, 0xda, 0x41, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x80, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0xda, 0x41, 0x05, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x05, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x32, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x84, 0x01, 0x0a, 0x11, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0xda, 0x41, 0x05,
	0x74, 0x72, 0x69, 0x61, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x05, 0x74, 0x72, 0x69,
	0x61, 0x6c, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_subscription_service_proto_rawDescOnce sync.Once
	file_v1_subscription_service_proto_rawDescData = file_v1_subscription_service_proto_rawDesc
)

func file_v1_subscription_service_proto_rawDescGZIP() []byte {
	file_v1_subscription_service_proto_rawDescOnce.Do(func() {
		file_v1_subscription_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_subscription_service_proto_rawDescData)
	})
	return file_v1_subscription_service_proto_rawDescData
}

var file_v1_subscription_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_subscription_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_subscription_service_proto_goTypes = []interface{}{
	(PlanType)(0),                     // 0: bytebase.v1.PlanType
	(*GetSubscriptionRequest)(nil),    // 1: bytebase.v1.GetSubscriptionRequest
	(*UpdateSubscriptionRequest)(nil), // 2: bytebase.v1.UpdateSubscriptionRequest
	(*TrialSubscriptionRequest)(nil),  // 3: bytebase.v1.TrialSubscriptionRequest
	(*PatchSubscription)(nil),         // 4: bytebase.v1.PatchSubscription
	(*TrialSubscription)(nil),         // 5: bytebase.v1.TrialSubscription
	(*Subscription)(nil),              // 6: bytebase.v1.Subscription
	(*timestamppb.Timestamp)(nil),     // 7: google.protobuf.Timestamp
}
var file_v1_subscription_service_proto_depIdxs = []int32{
	4, // 0: bytebase.v1.UpdateSubscriptionRequest.patch:type_name -> bytebase.v1.PatchSubscription
	5, // 1: bytebase.v1.TrialSubscriptionRequest.trial:type_name -> bytebase.v1.TrialSubscription
	0, // 2: bytebase.v1.TrialSubscription.plan:type_name -> bytebase.v1.PlanType
	7, // 3: bytebase.v1.Subscription.expires_time:type_name -> google.protobuf.Timestamp
	7, // 4: bytebase.v1.Subscription.started_time:type_name -> google.protobuf.Timestamp
	0, // 5: bytebase.v1.Subscription.plan:type_name -> bytebase.v1.PlanType
	1, // 6: bytebase.v1.SubscriptionService.GetSubscription:input_type -> bytebase.v1.GetSubscriptionRequest
	2, // 7: bytebase.v1.SubscriptionService.UpdateSubscription:input_type -> bytebase.v1.UpdateSubscriptionRequest
	3, // 8: bytebase.v1.SubscriptionService.TrialSubscription:input_type -> bytebase.v1.TrialSubscriptionRequest
	6, // 9: bytebase.v1.SubscriptionService.GetSubscription:output_type -> bytebase.v1.Subscription
	6, // 10: bytebase.v1.SubscriptionService.UpdateSubscription:output_type -> bytebase.v1.Subscription
	6, // 11: bytebase.v1.SubscriptionService.TrialSubscription:output_type -> bytebase.v1.Subscription
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v1_subscription_service_proto_init() }
func file_v1_subscription_service_proto_init() {
	if File_v1_subscription_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_subscription_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionRequest); i {
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
		file_v1_subscription_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSubscriptionRequest); i {
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
		file_v1_subscription_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrialSubscriptionRequest); i {
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
		file_v1_subscription_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchSubscription); i {
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
		file_v1_subscription_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrialSubscription); i {
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
		file_v1_subscription_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
			RawDescriptor: file_v1_subscription_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_subscription_service_proto_goTypes,
		DependencyIndexes: file_v1_subscription_service_proto_depIdxs,
		EnumInfos:         file_v1_subscription_service_proto_enumTypes,
		MessageInfos:      file_v1_subscription_service_proto_msgTypes,
	}.Build()
	File_v1_subscription_service_proto = out.File
	file_v1_subscription_service_proto_rawDesc = nil
	file_v1_subscription_service_proto_goTypes = nil
	file_v1_subscription_service_proto_depIdxs = nil
}
