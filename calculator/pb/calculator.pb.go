// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: calculator/pb/calculator.proto

package calculatorpb

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

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber  int32 `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber int32 `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_pb_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_pb_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calculator_pb_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *SumRequest) GetSecondNumber() int32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SumResult int32 `protobuf:"varint,1,opt,name=sum_result,json=sumResult,proto3" json:"sum_result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_pb_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_pb_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calculator_pb_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *SumResponse) GetSumResult() int32 {
	if x != nil {
		return x.SumResult
	}
	return 0
}

type SumOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber int32 `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	LastNumber  int32 `protobuf:"varint,2,opt,name=last_number,json=lastNumber,proto3" json:"last_number,omitempty"`
}

func (x *SumOneRequest) Reset() {
	*x = SumOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_pb_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumOneRequest) ProtoMessage() {}

func (x *SumOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_pb_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumOneRequest.ProtoReflect.Descriptor instead.
func (*SumOneRequest) Descriptor() ([]byte, []int) {
	return file_calculator_pb_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *SumOneRequest) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *SumOneRequest) GetLastNumber() int32 {
	if x != nil {
		return x.LastNumber
	}
	return 0
}

type SumManyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SumResult int32 `protobuf:"varint,1,opt,name=sum_result,json=sumResult,proto3" json:"sum_result,omitempty"`
}

func (x *SumManyResponse) Reset() {
	*x = SumManyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_pb_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumManyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumManyResponse) ProtoMessage() {}

func (x *SumManyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_pb_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumManyResponse.ProtoReflect.Descriptor instead.
func (*SumManyResponse) Descriptor() ([]byte, []int) {
	return file_calculator_pb_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *SumManyResponse) GetSumResult() int32 {
	if x != nil {
		return x.SumResult
	}
	return 0
}

type SumManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber int32 `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	LastNumber  int32 `protobuf:"varint,2,opt,name=last_number,json=lastNumber,proto3" json:"last_number,omitempty"`
}

func (x *SumManyRequest) Reset() {
	*x = SumManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_pb_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumManyRequest) ProtoMessage() {}

func (x *SumManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_pb_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumManyRequest.ProtoReflect.Descriptor instead.
func (*SumManyRequest) Descriptor() ([]byte, []int) {
	return file_calculator_pb_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *SumManyRequest) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *SumManyRequest) GetLastNumber() int32 {
	if x != nil {
		return x.LastNumber
	}
	return 0
}

type SumOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SumResult int32 `protobuf:"varint,1,opt,name=sum_result,json=sumResult,proto3" json:"sum_result,omitempty"`
}

func (x *SumOneResponse) Reset() {
	*x = SumOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_pb_calculator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumOneResponse) ProtoMessage() {}

func (x *SumOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_pb_calculator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumOneResponse.ProtoReflect.Descriptor instead.
func (*SumOneResponse) Descriptor() ([]byte, []int) {
	return file_calculator_pb_calculator_proto_rawDescGZIP(), []int{5}
}

func (x *SumOneResponse) GetSumResult() int32 {
	if x != nil {
		return x.SumResult
	}
	return 0
}

// for average client streaming
type AverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *AverageRequest) Reset() {
	*x = AverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_pb_calculator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageRequest) ProtoMessage() {}

func (x *AverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_pb_calculator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageRequest.ProtoReflect.Descriptor instead.
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return file_calculator_pb_calculator_proto_rawDescGZIP(), []int{6}
}

func (x *AverageRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type AverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average float64 `protobuf:"fixed64,1,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *AverageResponse) Reset() {
	*x = AverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_pb_calculator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageResponse) ProtoMessage() {}

func (x *AverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_pb_calculator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageResponse.ProtoReflect.Descriptor instead.
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return file_calculator_pb_calculator_proto_rawDescGZIP(), []int{7}
}

func (x *AverageResponse) GetAverage() float64 {
	if x != nil {
		return x.Average
	}
	return 0
}

var File_calculator_pb_calculator_proto protoreflect.FileDescriptor

var file_calculator_pb_calculator_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x62, 0x2f,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x54, 0x0a, 0x0a,
	0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x2c, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x53, 0x0a, 0x0d, 0x53, 0x75, 0x6d, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x0f, 0x53, 0x75, 0x6d, 0x4d, 0x61, 0x6e, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x6d, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x75,
	0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x54, 0x0a, 0x0e, 0x53, 0x75, 0x6d, 0x4d, 0x61,
	0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2f, 0x0a,
	0x0e, 0x53, 0x75, 0x6d, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x28,
	0x0a, 0x0e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0f, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x32, 0xad, 0x02, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x53,
	0x75, 0x6d, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x53, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x53, 0x75, 0x6d, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75, 0x6d, 0x4d,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x49, 0x0a,
	0x0d, 0x53, 0x75, 0x6d, 0x43, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x1a,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75, 0x6d, 0x4d,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75, 0x6d, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x4b, 0x0a, 0x0e, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_pb_calculator_proto_rawDescOnce sync.Once
	file_calculator_pb_calculator_proto_rawDescData = file_calculator_pb_calculator_proto_rawDesc
)

func file_calculator_pb_calculator_proto_rawDescGZIP() []byte {
	file_calculator_pb_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_pb_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_pb_calculator_proto_rawDescData)
	})
	return file_calculator_pb_calculator_proto_rawDescData
}

var file_calculator_pb_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_calculator_pb_calculator_proto_goTypes = []interface{}{
	(*SumRequest)(nil),      // 0: calculator.SumRequest
	(*SumResponse)(nil),     // 1: calculator.SumResponse
	(*SumOneRequest)(nil),   // 2: calculator.SumOneRequest
	(*SumManyResponse)(nil), // 3: calculator.SumManyResponse
	(*SumManyRequest)(nil),  // 4: calculator.SumManyRequest
	(*SumOneResponse)(nil),  // 5: calculator.SumOneResponse
	(*AverageRequest)(nil),  // 6: calculator.AverageRequest
	(*AverageResponse)(nil), // 7: calculator.AverageResponse
}
var file_calculator_pb_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.CalculatorService.Sum:input_type -> calculator.SumRequest
	2, // 1: calculator.CalculatorService.SumStreaming:input_type -> calculator.SumOneRequest
	4, // 2: calculator.CalculatorService.SumCStreaming:input_type -> calculator.SumManyRequest
	6, // 3: calculator.CalculatorService.AverageService:input_type -> calculator.AverageRequest
	1, // 4: calculator.CalculatorService.Sum:output_type -> calculator.SumResponse
	3, // 5: calculator.CalculatorService.SumStreaming:output_type -> calculator.SumManyResponse
	5, // 6: calculator.CalculatorService.SumCStreaming:output_type -> calculator.SumOneResponse
	7, // 7: calculator.CalculatorService.AverageService:output_type -> calculator.AverageResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_pb_calculator_proto_init() }
func file_calculator_pb_calculator_proto_init() {
	if File_calculator_pb_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_pb_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_calculator_pb_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_calculator_pb_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumOneRequest); i {
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
		file_calculator_pb_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumManyResponse); i {
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
		file_calculator_pb_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumManyRequest); i {
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
		file_calculator_pb_calculator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumOneResponse); i {
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
		file_calculator_pb_calculator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageRequest); i {
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
		file_calculator_pb_calculator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageResponse); i {
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
			RawDescriptor: file_calculator_pb_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_pb_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_pb_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_pb_calculator_proto_msgTypes,
	}.Build()
	File_calculator_pb_calculator_proto = out.File
	file_calculator_pb_calculator_proto_rawDesc = nil
	file_calculator_pb_calculator_proto_goTypes = nil
	file_calculator_pb_calculator_proto_depIdxs = nil
}
