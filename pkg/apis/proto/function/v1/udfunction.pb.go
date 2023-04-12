// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: pkg/apis/proto/function/v1/udfunction.proto

package v1

import (
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

type EventTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// event_time is the time associated with each datum.
	EventTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
}

func (x *EventTime) Reset() {
	*x = EventTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTime) ProtoMessage() {}

func (x *EventTime) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTime.ProtoReflect.Descriptor instead.
func (*EventTime) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_function_v1_udfunction_proto_rawDescGZIP(), []int{0}
}

func (x *EventTime) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

type Watermark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// watermark is the monotonically increasing time which denotes completeness for the given time for the given vertex.
	Watermark *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=watermark,proto3" json:"watermark,omitempty"` // future we can add LATE, ON_TIME etc.
}

func (x *Watermark) Reset() {
	*x = Watermark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Watermark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Watermark) ProtoMessage() {}

func (x *Watermark) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Watermark.ProtoReflect.Descriptor instead.
func (*Watermark) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_function_v1_udfunction_proto_rawDescGZIP(), []int{1}
}

func (x *Watermark) GetWatermark() *timestamppb.Timestamp {
	if x != nil {
		return x.Watermark
	}
	return nil
}

// *
// Metadata of a datum element.
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NumDelivered uint64 `protobuf:"varint,2,opt,name=num_delivered,json=numDelivered,proto3" json:"num_delivered,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_function_v1_udfunction_proto_rawDescGZIP(), []int{2}
}

func (x *Metadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Metadata) GetNumDelivered() uint64 {
	if x != nil {
		return x.NumDelivered
	}
	return 0
}

// *
// Datum represents a datum element.
type DatumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys      []string   `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Value     []byte     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	EventTime *EventTime `protobuf:"bytes,3,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	Watermark *Watermark `protobuf:"bytes,4,opt,name=watermark,proto3" json:"watermark,omitempty"`
	Metadata  *Metadata  `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *DatumRequest) Reset() {
	*x = DatumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatumRequest) ProtoMessage() {}

func (x *DatumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatumRequest.ProtoReflect.Descriptor instead.
func (*DatumRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_function_v1_udfunction_proto_rawDescGZIP(), []int{3}
}

func (x *DatumRequest) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *DatumRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *DatumRequest) GetEventTime() *EventTime {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *DatumRequest) GetWatermark() *Watermark {
	if x != nil {
		return x.Watermark
	}
	return nil
}

func (x *DatumRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// *
// Datum represents a datum element.
type DatumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys      []string   `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Value     []byte     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	EventTime *EventTime `protobuf:"bytes,3,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	Watermark *Watermark `protobuf:"bytes,4,opt,name=watermark,proto3" json:"watermark,omitempty"`
	Tags      []string   `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *DatumResponse) Reset() {
	*x = DatumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatumResponse) ProtoMessage() {}

func (x *DatumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatumResponse.ProtoReflect.Descriptor instead.
func (*DatumResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_function_v1_udfunction_proto_rawDescGZIP(), []int{4}
}

func (x *DatumResponse) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *DatumResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *DatumResponse) GetEventTime() *EventTime {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *DatumResponse) GetWatermark() *Watermark {
	if x != nil {
		return x.Watermark
	}
	return nil
}

func (x *DatumResponse) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// *
// DatumList represents a list of datum elements.
type DatumResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elements []*DatumResponse `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *DatumResponseList) Reset() {
	*x = DatumResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatumResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatumResponseList) ProtoMessage() {}

func (x *DatumResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatumResponseList.ProtoReflect.Descriptor instead.
func (*DatumResponseList) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_function_v1_udfunction_proto_rawDescGZIP(), []int{5}
}

func (x *DatumResponseList) GetElements() []*DatumResponse {
	if x != nil {
		return x.Elements
	}
	return nil
}

// *
// ReadyResponse is the health check result.
type ReadyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ready bool `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
}

func (x *ReadyResponse) Reset() {
	*x = ReadyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyResponse) ProtoMessage() {}

func (x *ReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyResponse.ProtoReflect.Descriptor instead.
func (*ReadyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_function_v1_udfunction_proto_rawDescGZIP(), []int{6}
}

func (x *ReadyResponse) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

var File_pkg_apis_proto_function_v1_udfunction_proto protoreflect.FileDescriptor

var file_pkg_apis_proto_function_v1_udfunction_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x64, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x45, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x38, 0x0a,
	0x09, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x77, 0x61,
	0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x3f, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x22, 0xd8, 0x01, 0x0a, 0x0c, 0x44, 0x61, 0x74,
	0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x77, 0x61,
	0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x65,
	0x72, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x09, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b,
	0x12, 0x31, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x22, 0xba, 0x01, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x35, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d,
	0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72,
	0x6b, 0x52, 0x09, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x22, 0x4b, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x25, 0x0a,
	0x0d, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72,
	0x65, 0x61, 0x64, 0x79, 0x32, 0xa8, 0x02, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x05,
	0x4d, 0x61, 0x70, 0x46, 0x6e, 0x12, 0x19, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x61, 0x74, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x43, 0x0a, 0x06, 0x4d, 0x61, 0x70, 0x54, 0x46, 0x6e, 0x12, 0x19, 0x2e, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x08, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x46,
	0x6e, 0x12, 0x19, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x61, 0x74, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x3d, 0x0a, 0x07, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75,
	0x6d, 0x61, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x6e, 0x75, 0x6d, 0x61, 0x66, 0x6c, 0x6f, 0x77, 0x2d,
	0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_apis_proto_function_v1_udfunction_proto_rawDescOnce sync.Once
	file_pkg_apis_proto_function_v1_udfunction_proto_rawDescData = file_pkg_apis_proto_function_v1_udfunction_proto_rawDesc
)

func file_pkg_apis_proto_function_v1_udfunction_proto_rawDescGZIP() []byte {
	file_pkg_apis_proto_function_v1_udfunction_proto_rawDescOnce.Do(func() {
		file_pkg_apis_proto_function_v1_udfunction_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_apis_proto_function_v1_udfunction_proto_rawDescData)
	})
	return file_pkg_apis_proto_function_v1_udfunction_proto_rawDescData
}

var file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_apis_proto_function_v1_udfunction_proto_goTypes = []interface{}{
	(*EventTime)(nil),             // 0: function.v1.EventTime
	(*Watermark)(nil),             // 1: function.v1.Watermark
	(*Metadata)(nil),              // 2: function.v1.Metadata
	(*DatumRequest)(nil),          // 3: function.v1.DatumRequest
	(*DatumResponse)(nil),         // 4: function.v1.DatumResponse
	(*DatumResponseList)(nil),     // 5: function.v1.DatumResponseList
	(*ReadyResponse)(nil),         // 6: function.v1.ReadyResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_pkg_apis_proto_function_v1_udfunction_proto_depIdxs = []int32{
	7,  // 0: function.v1.EventTime.event_time:type_name -> google.protobuf.Timestamp
	7,  // 1: function.v1.Watermark.watermark:type_name -> google.protobuf.Timestamp
	0,  // 2: function.v1.DatumRequest.event_time:type_name -> function.v1.EventTime
	1,  // 3: function.v1.DatumRequest.watermark:type_name -> function.v1.Watermark
	2,  // 4: function.v1.DatumRequest.metadata:type_name -> function.v1.Metadata
	0,  // 5: function.v1.DatumResponse.event_time:type_name -> function.v1.EventTime
	1,  // 6: function.v1.DatumResponse.watermark:type_name -> function.v1.Watermark
	4,  // 7: function.v1.DatumResponseList.elements:type_name -> function.v1.DatumResponse
	3,  // 8: function.v1.UserDefinedFunction.MapFn:input_type -> function.v1.DatumRequest
	3,  // 9: function.v1.UserDefinedFunction.MapTFn:input_type -> function.v1.DatumRequest
	3,  // 10: function.v1.UserDefinedFunction.ReduceFn:input_type -> function.v1.DatumRequest
	8,  // 11: function.v1.UserDefinedFunction.IsReady:input_type -> google.protobuf.Empty
	5,  // 12: function.v1.UserDefinedFunction.MapFn:output_type -> function.v1.DatumResponseList
	5,  // 13: function.v1.UserDefinedFunction.MapTFn:output_type -> function.v1.DatumResponseList
	5,  // 14: function.v1.UserDefinedFunction.ReduceFn:output_type -> function.v1.DatumResponseList
	6,  // 15: function.v1.UserDefinedFunction.IsReady:output_type -> function.v1.ReadyResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_pkg_apis_proto_function_v1_udfunction_proto_init() }
func file_pkg_apis_proto_function_v1_udfunction_proto_init() {
	if File_pkg_apis_proto_function_v1_udfunction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTime); i {
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
		file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Watermark); i {
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
		file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatumRequest); i {
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
		file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatumResponse); i {
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
		file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatumResponseList); i {
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
		file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyResponse); i {
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
			RawDescriptor: file_pkg_apis_proto_function_v1_udfunction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_apis_proto_function_v1_udfunction_proto_goTypes,
		DependencyIndexes: file_pkg_apis_proto_function_v1_udfunction_proto_depIdxs,
		MessageInfos:      file_pkg_apis_proto_function_v1_udfunction_proto_msgTypes,
	}.Build()
	File_pkg_apis_proto_function_v1_udfunction_proto = out.File
	file_pkg_apis_proto_function_v1_udfunction_proto_rawDesc = nil
	file_pkg_apis_proto_function_v1_udfunction_proto_goTypes = nil
	file_pkg_apis_proto_function_v1_udfunction_proto_depIdxs = nil
}
