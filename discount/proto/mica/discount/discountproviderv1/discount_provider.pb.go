// Copyright (c) 2022 Mica, Inc. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: mica/discount/discountprovider/v1/discount_provider.proto

package discountproviderv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	v1 "github.com/1080network/golang/discount/proto/micashared/common/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetDiscountProviderResponse_Status int32

const (
	GetDiscountProviderResponse_STATUS_UNSPECIFIED GetDiscountProviderResponse_Status = 0
	GetDiscountProviderResponse_STATUS_SUCCESS     GetDiscountProviderResponse_Status = 1
	GetDiscountProviderResponse_STATUS_ERROR       GetDiscountProviderResponse_Status = 2
	GetDiscountProviderResponse_STATUS_NOT_FOUND   GetDiscountProviderResponse_Status = 3
)

// Enum value maps for GetDiscountProviderResponse_Status.
var (
	GetDiscountProviderResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	GetDiscountProviderResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x GetDiscountProviderResponse_Status) Enum() *GetDiscountProviderResponse_Status {
	p := new(GetDiscountProviderResponse_Status)
	*p = x
	return p
}

func (x GetDiscountProviderResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetDiscountProviderResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_discount_discountprovider_v1_discount_provider_proto_enumTypes[0].Descriptor()
}

func (GetDiscountProviderResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_discount_discountprovider_v1_discount_provider_proto_enumTypes[0]
}

func (x GetDiscountProviderResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetDiscountProviderResponse_Status.Descriptor instead.
func (GetDiscountProviderResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescGZIP(), []int{2, 0}
}

type UpdateDiscountProviderResponse_Status int32

const (
	UpdateDiscountProviderResponse_STATUS_UNSPECIFIED UpdateDiscountProviderResponse_Status = 0
	UpdateDiscountProviderResponse_STATUS_SUCCESS     UpdateDiscountProviderResponse_Status = 1
	UpdateDiscountProviderResponse_STATUS_ERROR       UpdateDiscountProviderResponse_Status = 2
	UpdateDiscountProviderResponse_STATUS_NOT_FOUND   UpdateDiscountProviderResponse_Status = 3
)

// Enum value maps for UpdateDiscountProviderResponse_Status.
var (
	UpdateDiscountProviderResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	UpdateDiscountProviderResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x UpdateDiscountProviderResponse_Status) Enum() *UpdateDiscountProviderResponse_Status {
	p := new(UpdateDiscountProviderResponse_Status)
	*p = x
	return p
}

func (x UpdateDiscountProviderResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateDiscountProviderResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_discount_discountprovider_v1_discount_provider_proto_enumTypes[1].Descriptor()
}

func (UpdateDiscountProviderResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_discount_discountprovider_v1_discount_provider_proto_enumTypes[1]
}

func (x UpdateDiscountProviderResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateDiscountProviderResponse_Status.Descriptor instead.
func (UpdateDiscountProviderResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescGZIP(), []int{4, 0}
}

type DiscountProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscountProviderKey string                 `protobuf:"bytes,1,opt,name=discount_provider_key,json=discountProviderKey,proto3" json:"discount_provider_key,omitempty"`
	Version             int64                  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Created             *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated             *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name                string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Address             *v1.Address            `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	CallbackAddress     string                 `protobuf:"bytes,8,opt,name=callback_address,json=callbackAddress,proto3" json:"callback_address,omitempty"`
	Email               string                 `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
	Phone               string                 `protobuf:"bytes,10,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *DiscountProvider) Reset() {
	*x = DiscountProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountProvider) ProtoMessage() {}

func (x *DiscountProvider) ProtoReflect() protoreflect.Message {
	mi := &file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountProvider.ProtoReflect.Descriptor instead.
func (*DiscountProvider) Descriptor() ([]byte, []int) {
	return file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescGZIP(), []int{0}
}

func (x *DiscountProvider) GetDiscountProviderKey() string {
	if x != nil {
		return x.DiscountProviderKey
	}
	return ""
}

func (x *DiscountProvider) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *DiscountProvider) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *DiscountProvider) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *DiscountProvider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DiscountProvider) GetAddress() *v1.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *DiscountProvider) GetCallbackAddress() string {
	if x != nil {
		return x.CallbackAddress
	}
	return ""
}

func (x *DiscountProvider) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DiscountProvider) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type GetDiscountProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDiscountProviderRequest) Reset() {
	*x = GetDiscountProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDiscountProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDiscountProviderRequest) ProtoMessage() {}

func (x *GetDiscountProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDiscountProviderRequest.ProtoReflect.Descriptor instead.
func (*GetDiscountProviderRequest) Descriptor() ([]byte, []int) {
	return file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescGZIP(), []int{1}
}

type GetDiscountProviderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status           GetDiscountProviderResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.discount.discountprovider.v1.GetDiscountProviderResponse_Status" json:"status,omitempty"`
	Error            *v1.Error                          `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	DiscountProvider *DiscountProvider                  `protobuf:"bytes,3,opt,name=discount_provider,json=discountProvider,proto3" json:"discount_provider,omitempty"`
}

func (x *GetDiscountProviderResponse) Reset() {
	*x = GetDiscountProviderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDiscountProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDiscountProviderResponse) ProtoMessage() {}

func (x *GetDiscountProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDiscountProviderResponse.ProtoReflect.Descriptor instead.
func (*GetDiscountProviderResponse) Descriptor() ([]byte, []int) {
	return file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescGZIP(), []int{2}
}

func (x *GetDiscountProviderResponse) GetStatus() GetDiscountProviderResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetDiscountProviderResponse_STATUS_UNSPECIFIED
}

func (x *GetDiscountProviderResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetDiscountProviderResponse) GetDiscountProvider() *DiscountProvider {
	if x != nil {
		return x.DiscountProvider
	}
	return nil
}

type UpdateDiscountProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version         int64       `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Name            string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address         *v1.Address `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	CallbackAddress string      `protobuf:"bytes,4,opt,name=callback_address,json=callbackAddress,proto3" json:"callback_address,omitempty"`
	Email           string      `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone           string      `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UpdateDiscountProviderRequest) Reset() {
	*x = UpdateDiscountProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDiscountProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDiscountProviderRequest) ProtoMessage() {}

func (x *UpdateDiscountProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDiscountProviderRequest.ProtoReflect.Descriptor instead.
func (*UpdateDiscountProviderRequest) Descriptor() ([]byte, []int) {
	return file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDiscountProviderRequest) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *UpdateDiscountProviderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDiscountProviderRequest) GetAddress() *v1.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *UpdateDiscountProviderRequest) GetCallbackAddress() string {
	if x != nil {
		return x.CallbackAddress
	}
	return ""
}

func (x *UpdateDiscountProviderRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateDiscountProviderRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type UpdateDiscountProviderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  UpdateDiscountProviderResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.discount.discountprovider.v1.UpdateDiscountProviderResponse_Status" json:"status,omitempty"`
	Error   *v1.Error                             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Version int64                                 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *UpdateDiscountProviderResponse) Reset() {
	*x = UpdateDiscountProviderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDiscountProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDiscountProviderResponse) ProtoMessage() {}

func (x *UpdateDiscountProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDiscountProviderResponse.ProtoReflect.Descriptor instead.
func (*UpdateDiscountProviderResponse) Descriptor() ([]byte, []int) {
	return file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDiscountProviderResponse) GetStatus() UpdateDiscountProviderResponse_Status {
	if x != nil {
		return x.Status
	}
	return UpdateDiscountProviderResponse_STATUS_UNSPECIFIED
}

func (x *UpdateDiscountProviderResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *UpdateDiscountProviderResponse) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_mica_discount_discountprovider_v1_discount_provider_proto protoreflect.FileDescriptor

var file_mica_discount_discountprovider_v1_discount_provider_proto_rawDesc = []byte{
	0x0a, 0x39, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x6d, 0x69, 0x63,
	0x61, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab,
	0x03, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x15, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1b, 0x18, 0x32, 0x52, 0x13, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4b,
	0x65, 0x79, 0x12, 0x21, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x00, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x64, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x32, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xfa, 0x42, 0x19, 0x72, 0x17, 0x32, 0x15,
	0x5e, 0x24, 0x7c, 0x5e, 0x5c, 0x2b, 0x5b, 0x31, 0x2d, 0x39, 0x5d, 0x5c, 0x64, 0x7b, 0x31, 0x30,
	0x2c, 0x31, 0x34, 0x7d, 0x24, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x1c, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xef, 0x02, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x45, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x60, 0x0a, 0x11,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x10, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x5c,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x22, 0x8d, 0x02, 0x0a,
	0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x00, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x64, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x32, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xfa, 0x42, 0x19, 0x72, 0x17, 0x32, 0x15,
	0x5e, 0x24, 0x7c, 0x5e, 0x5c, 0x2b, 0x5b, 0x31, 0x2d, 0x39, 0x5d, 0x5c, 0x64, 0x7b, 0x31, 0x30,
	0x2c, 0x31, 0x34, 0x7d, 0x24, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0xb6, 0x02, 0x0a,
	0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x60, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x48, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x00, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a,
	0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12,
	0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x03, 0x42, 0x66, 0x0a, 0x24, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x20, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescOnce sync.Once
	file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescData = file_mica_discount_discountprovider_v1_discount_provider_proto_rawDesc
)

func file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescGZIP() []byte {
	file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescOnce.Do(func() {
		file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescData)
	})
	return file_mica_discount_discountprovider_v1_discount_provider_proto_rawDescData
}

var file_mica_discount_discountprovider_v1_discount_provider_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mica_discount_discountprovider_v1_discount_provider_proto_goTypes = []interface{}{
	(GetDiscountProviderResponse_Status)(0),    // 0: mica.discount.discountprovider.v1.GetDiscountProviderResponse.Status
	(UpdateDiscountProviderResponse_Status)(0), // 1: mica.discount.discountprovider.v1.UpdateDiscountProviderResponse.Status
	(*DiscountProvider)(nil),                   // 2: mica.discount.discountprovider.v1.DiscountProvider
	(*GetDiscountProviderRequest)(nil),         // 3: mica.discount.discountprovider.v1.GetDiscountProviderRequest
	(*GetDiscountProviderResponse)(nil),        // 4: mica.discount.discountprovider.v1.GetDiscountProviderResponse
	(*UpdateDiscountProviderRequest)(nil),      // 5: mica.discount.discountprovider.v1.UpdateDiscountProviderRequest
	(*UpdateDiscountProviderResponse)(nil),     // 6: mica.discount.discountprovider.v1.UpdateDiscountProviderResponse
	(*timestamppb.Timestamp)(nil),              // 7: google.protobuf.Timestamp
	(*v1.Address)(nil),                         // 8: micashared.common.v1.Address
	(*v1.Error)(nil),                           // 9: micashared.common.v1.Error
}
var file_mica_discount_discountprovider_v1_discount_provider_proto_depIdxs = []int32{
	7, // 0: mica.discount.discountprovider.v1.DiscountProvider.created:type_name -> google.protobuf.Timestamp
	7, // 1: mica.discount.discountprovider.v1.DiscountProvider.updated:type_name -> google.protobuf.Timestamp
	8, // 2: mica.discount.discountprovider.v1.DiscountProvider.address:type_name -> micashared.common.v1.Address
	0, // 3: mica.discount.discountprovider.v1.GetDiscountProviderResponse.status:type_name -> mica.discount.discountprovider.v1.GetDiscountProviderResponse.Status
	9, // 4: mica.discount.discountprovider.v1.GetDiscountProviderResponse.error:type_name -> micashared.common.v1.Error
	2, // 5: mica.discount.discountprovider.v1.GetDiscountProviderResponse.discount_provider:type_name -> mica.discount.discountprovider.v1.DiscountProvider
	8, // 6: mica.discount.discountprovider.v1.UpdateDiscountProviderRequest.address:type_name -> micashared.common.v1.Address
	1, // 7: mica.discount.discountprovider.v1.UpdateDiscountProviderResponse.status:type_name -> mica.discount.discountprovider.v1.UpdateDiscountProviderResponse.Status
	9, // 8: mica.discount.discountprovider.v1.UpdateDiscountProviderResponse.error:type_name -> micashared.common.v1.Error
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_mica_discount_discountprovider_v1_discount_provider_proto_init() }
func file_mica_discount_discountprovider_v1_discount_provider_proto_init() {
	if File_mica_discount_discountprovider_v1_discount_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountProvider); i {
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
		file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDiscountProviderRequest); i {
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
		file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDiscountProviderResponse); i {
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
		file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDiscountProviderRequest); i {
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
		file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDiscountProviderResponse); i {
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
			RawDescriptor: file_mica_discount_discountprovider_v1_discount_provider_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mica_discount_discountprovider_v1_discount_provider_proto_goTypes,
		DependencyIndexes: file_mica_discount_discountprovider_v1_discount_provider_proto_depIdxs,
		EnumInfos:         file_mica_discount_discountprovider_v1_discount_provider_proto_enumTypes,
		MessageInfos:      file_mica_discount_discountprovider_v1_discount_provider_proto_msgTypes,
	}.Build()
	File_mica_discount_discountprovider_v1_discount_provider_proto = out.File
	file_mica_discount_discountprovider_v1_discount_provider_proto_rawDesc = nil
	file_mica_discount_discountprovider_v1_discount_provider_proto_goTypes = nil
	file_mica_discount_discountprovider_v1_discount_provider_proto_depIdxs = nil
}
