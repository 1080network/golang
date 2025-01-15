// Copyright (c) 2022-2024 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.3
// source: mica/discount/service/v1/discount_service_test_support.proto

package servicev1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "github.com/1080network/golang/fullsdk/proto/micashared/common/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DetermineDiscountResponse_Status int32

const (
	DetermineDiscountResponse_STATUS_UNSPECIFIED DetermineDiscountResponse_Status = 0
	DetermineDiscountResponse_STATUS_SUCCESS     DetermineDiscountResponse_Status = 1
	DetermineDiscountResponse_STATUS_ERROR       DetermineDiscountResponse_Status = 2
)

// Enum value maps for DetermineDiscountResponse_Status.
var (
	DetermineDiscountResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
	}
	DetermineDiscountResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
	}
)

func (x DetermineDiscountResponse_Status) Enum() *DetermineDiscountResponse_Status {
	p := new(DetermineDiscountResponse_Status)
	*p = x
	return p
}

func (x DetermineDiscountResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DetermineDiscountResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_discount_service_v1_discount_service_test_support_proto_enumTypes[0].Descriptor()
}

func (DetermineDiscountResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_discount_service_v1_discount_service_test_support_proto_enumTypes[0]
}

func (x DetermineDiscountResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DetermineDiscountResponse_Status.Descriptor instead.
func (DetermineDiscountResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_discount_service_v1_discount_service_test_support_proto_rawDescGZIP(), []int{1, 0}
}

type ApplyDiscountResponse_Status int32

const (
	ApplyDiscountResponse_STATUS_UNSPECIFIED ApplyDiscountResponse_Status = 0
	ApplyDiscountResponse_STATUS_SUCCESS     ApplyDiscountResponse_Status = 1
	ApplyDiscountResponse_STATUS_ERROR       ApplyDiscountResponse_Status = 2
	ApplyDiscountResponse_STATUS_NOT_FOUND   ApplyDiscountResponse_Status = 3
)

// Enum value maps for ApplyDiscountResponse_Status.
var (
	ApplyDiscountResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	ApplyDiscountResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x ApplyDiscountResponse_Status) Enum() *ApplyDiscountResponse_Status {
	p := new(ApplyDiscountResponse_Status)
	*p = x
	return p
}

func (x ApplyDiscountResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApplyDiscountResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_discount_service_v1_discount_service_test_support_proto_enumTypes[1].Descriptor()
}

func (ApplyDiscountResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_discount_service_v1_discount_service_test_support_proto_enumTypes[1]
}

func (x ApplyDiscountResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApplyDiscountResponse_Status.Descriptor instead.
func (ApplyDiscountResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_discount_service_v1_discount_service_test_support_proto_rawDescGZIP(), []int{3, 0}
}

type ReleaseDiscountResponse_Status int32

const (
	ReleaseDiscountResponse_STATUS_UNSPECIFIED ReleaseDiscountResponse_Status = 0
	ReleaseDiscountResponse_STATUS_SUCCESS     ReleaseDiscountResponse_Status = 1
	ReleaseDiscountResponse_STATUS_ERROR       ReleaseDiscountResponse_Status = 2
	ReleaseDiscountResponse_STATUS_NOT_FOUND   ReleaseDiscountResponse_Status = 3
)

// Enum value maps for ReleaseDiscountResponse_Status.
var (
	ReleaseDiscountResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	ReleaseDiscountResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x ReleaseDiscountResponse_Status) Enum() *ReleaseDiscountResponse_Status {
	p := new(ReleaseDiscountResponse_Status)
	*p = x
	return p
}

func (x ReleaseDiscountResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReleaseDiscountResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_discount_service_v1_discount_service_test_support_proto_enumTypes[2].Descriptor()
}

func (ReleaseDiscountResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_discount_service_v1_discount_service_test_support_proto_enumTypes[2]
}

func (x ReleaseDiscountResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReleaseDiscountResponse_Status.Descriptor instead.
func (ReleaseDiscountResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_discount_service_v1_discount_service_test_support_proto_rawDescGZIP(), []int{5, 0}
}

type DetermineDiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FederatedUserKey string         `protobuf:"bytes,1,opt,name=federated_user_key,json=federatedUserKey,proto3" json:"federated_user_key,omitempty"`
	LineItems        []*v1.LineItem `protobuf:"bytes,2,rep,name=line_items,json=lineItems,proto3" json:"line_items,omitempty"`
}

func (x *DetermineDiscountRequest) Reset() {
	*x = DetermineDiscountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetermineDiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetermineDiscountRequest) ProtoMessage() {}

func (x *DetermineDiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetermineDiscountRequest.ProtoReflect.Descriptor instead.
func (*DetermineDiscountRequest) Descriptor() ([]byte, []int) {
	return file_mica_discount_service_v1_discount_service_test_support_proto_rawDescGZIP(), []int{0}
}

func (x *DetermineDiscountRequest) GetFederatedUserKey() string {
	if x != nil {
		return x.FederatedUserKey
	}
	return ""
}

func (x *DetermineDiscountRequest) GetLineItems() []*v1.LineItem {
	if x != nil {
		return x.LineItems
	}
	return nil
}

type DetermineDiscountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status DetermineDiscountResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.discount.service.v1.DetermineDiscountResponse_Status" json:"status,omitempty"`
	Error  *v1.Error                        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DetermineDiscountResponse) Reset() {
	*x = DetermineDiscountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetermineDiscountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetermineDiscountResponse) ProtoMessage() {}

func (x *DetermineDiscountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetermineDiscountResponse.ProtoReflect.Descriptor instead.
func (*DetermineDiscountResponse) Descriptor() ([]byte, []int) {
	return file_mica_discount_service_v1_discount_service_test_support_proto_rawDescGZIP(), []int{1}
}

func (x *DetermineDiscountResponse) GetStatus() DetermineDiscountResponse_Status {
	if x != nil {
		return x.Status
	}
	return DetermineDiscountResponse_STATUS_UNSPECIFIED
}

func (x *DetermineDiscountResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type ApplyDiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscountRef string `protobuf:"bytes,1,opt,name=discount_ref,json=discountRef,proto3" json:"discount_ref,omitempty"`
}

func (x *ApplyDiscountRequest) Reset() {
	*x = ApplyDiscountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyDiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyDiscountRequest) ProtoMessage() {}

func (x *ApplyDiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyDiscountRequest.ProtoReflect.Descriptor instead.
func (*ApplyDiscountRequest) Descriptor() ([]byte, []int) {
	return file_mica_discount_service_v1_discount_service_test_support_proto_rawDescGZIP(), []int{2}
}

func (x *ApplyDiscountRequest) GetDiscountRef() string {
	if x != nil {
		return x.DiscountRef
	}
	return ""
}

type ApplyDiscountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ApplyDiscountResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.discount.service.v1.ApplyDiscountResponse_Status" json:"status,omitempty"`
	Error  *v1.Error                    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ApplyDiscountResponse) Reset() {
	*x = ApplyDiscountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyDiscountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyDiscountResponse) ProtoMessage() {}

func (x *ApplyDiscountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyDiscountResponse.ProtoReflect.Descriptor instead.
func (*ApplyDiscountResponse) Descriptor() ([]byte, []int) {
	return file_mica_discount_service_v1_discount_service_test_support_proto_rawDescGZIP(), []int{3}
}

func (x *ApplyDiscountResponse) GetStatus() ApplyDiscountResponse_Status {
	if x != nil {
		return x.Status
	}
	return ApplyDiscountResponse_STATUS_UNSPECIFIED
}

func (x *ApplyDiscountResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type ReleaseDiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscountRef string `protobuf:"bytes,1,opt,name=discount_ref,json=discountRef,proto3" json:"discount_ref,omitempty"`
}

func (x *ReleaseDiscountRequest) Reset() {
	*x = ReleaseDiscountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseDiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseDiscountRequest) ProtoMessage() {}

func (x *ReleaseDiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseDiscountRequest.ProtoReflect.Descriptor instead.
func (*ReleaseDiscountRequest) Descriptor() ([]byte, []int) {
	return file_mica_discount_service_v1_discount_service_test_support_proto_rawDescGZIP(), []int{4}
}

func (x *ReleaseDiscountRequest) GetDiscountRef() string {
	if x != nil {
		return x.DiscountRef
	}
	return ""
}

type ReleaseDiscountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ReleaseDiscountResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.discount.service.v1.ReleaseDiscountResponse_Status" json:"status,omitempty"`
	Error  *v1.Error                      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ReleaseDiscountResponse) Reset() {
	*x = ReleaseDiscountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseDiscountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseDiscountResponse) ProtoMessage() {}

func (x *ReleaseDiscountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseDiscountResponse.ProtoReflect.Descriptor instead.
func (*ReleaseDiscountResponse) Descriptor() ([]byte, []int) {
	return file_mica_discount_service_v1_discount_service_test_support_proto_rawDescGZIP(), []int{5}
}

func (x *ReleaseDiscountResponse) GetStatus() ReleaseDiscountResponse_Status {
	if x != nil {
		return x.Status
	}
	return ReleaseDiscountResponse_STATUS_UNSPECIFIED
}

func (x *ReleaseDiscountResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_mica_discount_service_v1_discount_service_test_support_proto protoreflect.FileDescriptor

var file_mica_discount_service_v1_discount_service_test_support_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x6d, 0x69, 0x63, 0x61, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x18, 0x44, 0x65,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x12, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1b, 0x18, 0x32, 0x52, 0x10, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12,
	0x47, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x09, 0x6c,
	0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xea, 0x01, 0x0a, 0x19, 0x44, 0x65, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x46, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x02, 0x22, 0x39, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x66,
	0x22, 0xf8, 0x01, 0x0a, 0x15, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5c, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x22, 0x3b, 0x0a, 0x16, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x22, 0xfc, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x32, 0x97, 0x03, 0x0a, 0x1a, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x54, 0x65, 0x73, 0x74, 0x44,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x11, 0x54,
	0x65, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2e, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x13, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x5e, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x1f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x17, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43,
	0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mica_discount_service_v1_discount_service_test_support_proto_rawDescOnce sync.Once
	file_mica_discount_service_v1_discount_service_test_support_proto_rawDescData = file_mica_discount_service_v1_discount_service_test_support_proto_rawDesc
)

func file_mica_discount_service_v1_discount_service_test_support_proto_rawDescGZIP() []byte {
	file_mica_discount_service_v1_discount_service_test_support_proto_rawDescOnce.Do(func() {
		file_mica_discount_service_v1_discount_service_test_support_proto_rawDescData = protoimpl.X.CompressGZIP(file_mica_discount_service_v1_discount_service_test_support_proto_rawDescData)
	})
	return file_mica_discount_service_v1_discount_service_test_support_proto_rawDescData
}

var file_mica_discount_service_v1_discount_service_test_support_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mica_discount_service_v1_discount_service_test_support_proto_goTypes = []interface{}{
	(DetermineDiscountResponse_Status)(0), // 0: mica.discount.service.v1.DetermineDiscountResponse.Status
	(ApplyDiscountResponse_Status)(0),     // 1: mica.discount.service.v1.ApplyDiscountResponse.Status
	(ReleaseDiscountResponse_Status)(0),   // 2: mica.discount.service.v1.ReleaseDiscountResponse.Status
	(*DetermineDiscountRequest)(nil),      // 3: mica.discount.service.v1.DetermineDiscountRequest
	(*DetermineDiscountResponse)(nil),     // 4: mica.discount.service.v1.DetermineDiscountResponse
	(*ApplyDiscountRequest)(nil),          // 5: mica.discount.service.v1.ApplyDiscountRequest
	(*ApplyDiscountResponse)(nil),         // 6: mica.discount.service.v1.ApplyDiscountResponse
	(*ReleaseDiscountRequest)(nil),        // 7: mica.discount.service.v1.ReleaseDiscountRequest
	(*ReleaseDiscountResponse)(nil),       // 8: mica.discount.service.v1.ReleaseDiscountResponse
	(*v1.LineItem)(nil),                   // 9: micashared.common.v1.LineItem
	(*v1.Error)(nil),                      // 10: micashared.common.v1.Error
}
var file_mica_discount_service_v1_discount_service_test_support_proto_depIdxs = []int32{
	9,  // 0: mica.discount.service.v1.DetermineDiscountRequest.line_items:type_name -> micashared.common.v1.LineItem
	0,  // 1: mica.discount.service.v1.DetermineDiscountResponse.status:type_name -> mica.discount.service.v1.DetermineDiscountResponse.Status
	10, // 2: mica.discount.service.v1.DetermineDiscountResponse.error:type_name -> micashared.common.v1.Error
	1,  // 3: mica.discount.service.v1.ApplyDiscountResponse.status:type_name -> mica.discount.service.v1.ApplyDiscountResponse.Status
	10, // 4: mica.discount.service.v1.ApplyDiscountResponse.error:type_name -> micashared.common.v1.Error
	2,  // 5: mica.discount.service.v1.ReleaseDiscountResponse.status:type_name -> mica.discount.service.v1.ReleaseDiscountResponse.Status
	10, // 6: mica.discount.service.v1.ReleaseDiscountResponse.error:type_name -> micashared.common.v1.Error
	3,  // 7: mica.discount.service.v1.DiscountTestSupportService.TestDetermineDiscount:input_type -> mica.discount.service.v1.DetermineDiscountRequest
	5,  // 8: mica.discount.service.v1.DiscountTestSupportService.TestApplyDiscount:input_type -> mica.discount.service.v1.ApplyDiscountRequest
	7,  // 9: mica.discount.service.v1.DiscountTestSupportService.TestReleaseDiscount:input_type -> mica.discount.service.v1.ReleaseDiscountRequest
	4,  // 10: mica.discount.service.v1.DiscountTestSupportService.TestDetermineDiscount:output_type -> mica.discount.service.v1.DetermineDiscountResponse
	6,  // 11: mica.discount.service.v1.DiscountTestSupportService.TestApplyDiscount:output_type -> mica.discount.service.v1.ApplyDiscountResponse
	8,  // 12: mica.discount.service.v1.DiscountTestSupportService.TestReleaseDiscount:output_type -> mica.discount.service.v1.ReleaseDiscountResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_mica_discount_service_v1_discount_service_test_support_proto_init() }
func file_mica_discount_service_v1_discount_service_test_support_proto_init() {
	if File_mica_discount_service_v1_discount_service_test_support_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetermineDiscountRequest); i {
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
		file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetermineDiscountResponse); i {
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
		file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyDiscountRequest); i {
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
		file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyDiscountResponse); i {
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
		file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseDiscountRequest); i {
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
		file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseDiscountResponse); i {
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
			RawDescriptor: file_mica_discount_service_v1_discount_service_test_support_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mica_discount_service_v1_discount_service_test_support_proto_goTypes,
		DependencyIndexes: file_mica_discount_service_v1_discount_service_test_support_proto_depIdxs,
		EnumInfos:         file_mica_discount_service_v1_discount_service_test_support_proto_enumTypes,
		MessageInfos:      file_mica_discount_service_v1_discount_service_test_support_proto_msgTypes,
	}.Build()
	File_mica_discount_service_v1_discount_service_test_support_proto = out.File
	file_mica_discount_service_v1_discount_service_test_support_proto_rawDesc = nil
	file_mica_discount_service_v1_discount_service_test_support_proto_goTypes = nil
	file_mica_discount_service_v1_discount_service_test_support_proto_depIdxs = nil
}
