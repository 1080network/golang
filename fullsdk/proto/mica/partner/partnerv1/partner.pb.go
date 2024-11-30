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
// 	protoc        v4.24.4
// source: mica/partner/partner/v1/partner.proto

package partnerv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type GetPartnerResponse_Status int32

const (
	GetPartnerResponse_STATUS_UNSPECIFIED GetPartnerResponse_Status = 0
	GetPartnerResponse_STATUS_SUCCESS     GetPartnerResponse_Status = 1
	GetPartnerResponse_STATUS_ERROR       GetPartnerResponse_Status = 2
	GetPartnerResponse_STATUS_NOT_FOUND   GetPartnerResponse_Status = 3
)

// Enum value maps for GetPartnerResponse_Status.
var (
	GetPartnerResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	GetPartnerResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x GetPartnerResponse_Status) Enum() *GetPartnerResponse_Status {
	p := new(GetPartnerResponse_Status)
	*p = x
	return p
}

func (x GetPartnerResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetPartnerResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_partner_partner_v1_partner_proto_enumTypes[0].Descriptor()
}

func (GetPartnerResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_partner_partner_v1_partner_proto_enumTypes[0]
}

func (x GetPartnerResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetPartnerResponse_Status.Descriptor instead.
func (GetPartnerResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_partner_partner_v1_partner_proto_rawDescGZIP(), []int{2, 0}
}

type Partner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartnerKey      string                 `protobuf:"bytes,1,opt,name=partner_key,json=partnerKey,proto3" json:"partner_key,omitempty"`
	Version         int64                  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Created         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated         *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name            string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	TaxIdentifier   string                 `protobuf:"bytes,6,opt,name=tax_identifier,json=taxIdentifier,proto3" json:"tax_identifier,omitempty"`
	Address         *v1.Address            `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	CallbackAddress string                 `protobuf:"bytes,8,opt,name=callback_address,json=callbackAddress,proto3" json:"callback_address,omitempty"`
	NetworkGroup    string                 `protobuf:"bytes,9,opt,name=network_group,json=networkGroup,proto3" json:"network_group,omitempty"`
	Features        []string               `protobuf:"bytes,10,rep,name=features,proto3" json:"features,omitempty"`
}

func (x *Partner) Reset() {
	*x = Partner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_partner_v1_partner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Partner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partner) ProtoMessage() {}

func (x *Partner) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_partner_v1_partner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partner.ProtoReflect.Descriptor instead.
func (*Partner) Descriptor() ([]byte, []int) {
	return file_mica_partner_partner_v1_partner_proto_rawDescGZIP(), []int{0}
}

func (x *Partner) GetPartnerKey() string {
	if x != nil {
		return x.PartnerKey
	}
	return ""
}

func (x *Partner) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Partner) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Partner) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *Partner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Partner) GetTaxIdentifier() string {
	if x != nil {
		return x.TaxIdentifier
	}
	return ""
}

func (x *Partner) GetAddress() *v1.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Partner) GetCallbackAddress() string {
	if x != nil {
		return x.CallbackAddress
	}
	return ""
}

func (x *Partner) GetNetworkGroup() string {
	if x != nil {
		return x.NetworkGroup
	}
	return ""
}

func (x *Partner) GetFeatures() []string {
	if x != nil {
		return x.Features
	}
	return nil
}

type GetPartnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPartnerRequest) Reset() {
	*x = GetPartnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_partner_v1_partner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartnerRequest) ProtoMessage() {}

func (x *GetPartnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_partner_v1_partner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartnerRequest.ProtoReflect.Descriptor instead.
func (*GetPartnerRequest) Descriptor() ([]byte, []int) {
	return file_mica_partner_partner_v1_partner_proto_rawDescGZIP(), []int{1}
}

type GetPartnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  GetPartnerResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.partner.partner.v1.GetPartnerResponse_Status" json:"status,omitempty"`
	Error   *v1.Error                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Partner *Partner                  `protobuf:"bytes,3,opt,name=partner,proto3" json:"partner,omitempty"`
}

func (x *GetPartnerResponse) Reset() {
	*x = GetPartnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_partner_v1_partner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartnerResponse) ProtoMessage() {}

func (x *GetPartnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_partner_v1_partner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartnerResponse.ProtoReflect.Descriptor instead.
func (*GetPartnerResponse) Descriptor() ([]byte, []int) {
	return file_mica_partner_partner_v1_partner_proto_rawDescGZIP(), []int{2}
}

func (x *GetPartnerResponse) GetStatus() GetPartnerResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetPartnerResponse_STATUS_UNSPECIFIED
}

func (x *GetPartnerResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetPartnerResponse) GetPartner() *Partner {
	if x != nil {
		return x.Partner
	}
	return nil
}

var File_mica_partner_partner_v1_partner_proto protoreflect.FileDescriptor

var file_mica_partner_partner_v1_partner_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa4, 0x03, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x0b,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02,
	0x28, 0x00, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x61, 0x78, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x61, 0x78, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x37, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x63,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xad, 0x02, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x3a, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x22, 0x5c,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x42, 0x49, 0x0a, 0x1a,
	0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x16, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x76, 0x31,
	0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mica_partner_partner_v1_partner_proto_rawDescOnce sync.Once
	file_mica_partner_partner_v1_partner_proto_rawDescData = file_mica_partner_partner_v1_partner_proto_rawDesc
)

func file_mica_partner_partner_v1_partner_proto_rawDescGZIP() []byte {
	file_mica_partner_partner_v1_partner_proto_rawDescOnce.Do(func() {
		file_mica_partner_partner_v1_partner_proto_rawDescData = protoimpl.X.CompressGZIP(file_mica_partner_partner_v1_partner_proto_rawDescData)
	})
	return file_mica_partner_partner_v1_partner_proto_rawDescData
}

var file_mica_partner_partner_v1_partner_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mica_partner_partner_v1_partner_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mica_partner_partner_v1_partner_proto_goTypes = []interface{}{
	(GetPartnerResponse_Status)(0), // 0: mica.partner.partner.v1.GetPartnerResponse.Status
	(*Partner)(nil),                // 1: mica.partner.partner.v1.Partner
	(*GetPartnerRequest)(nil),      // 2: mica.partner.partner.v1.GetPartnerRequest
	(*GetPartnerResponse)(nil),     // 3: mica.partner.partner.v1.GetPartnerResponse
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
	(*v1.Address)(nil),             // 5: micashared.common.v1.Address
	(*v1.Error)(nil),               // 6: micashared.common.v1.Error
}
var file_mica_partner_partner_v1_partner_proto_depIdxs = []int32{
	4, // 0: mica.partner.partner.v1.Partner.created:type_name -> google.protobuf.Timestamp
	4, // 1: mica.partner.partner.v1.Partner.updated:type_name -> google.protobuf.Timestamp
	5, // 2: mica.partner.partner.v1.Partner.address:type_name -> micashared.common.v1.Address
	0, // 3: mica.partner.partner.v1.GetPartnerResponse.status:type_name -> mica.partner.partner.v1.GetPartnerResponse.Status
	6, // 4: mica.partner.partner.v1.GetPartnerResponse.error:type_name -> micashared.common.v1.Error
	1, // 5: mica.partner.partner.v1.GetPartnerResponse.partner:type_name -> mica.partner.partner.v1.Partner
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_mica_partner_partner_v1_partner_proto_init() }
func file_mica_partner_partner_v1_partner_proto_init() {
	if File_mica_partner_partner_v1_partner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mica_partner_partner_v1_partner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Partner); i {
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
		file_mica_partner_partner_v1_partner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartnerRequest); i {
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
		file_mica_partner_partner_v1_partner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartnerResponse); i {
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
			RawDescriptor: file_mica_partner_partner_v1_partner_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mica_partner_partner_v1_partner_proto_goTypes,
		DependencyIndexes: file_mica_partner_partner_v1_partner_proto_depIdxs,
		EnumInfos:         file_mica_partner_partner_v1_partner_proto_enumTypes,
		MessageInfos:      file_mica_partner_partner_v1_partner_proto_msgTypes,
	}.Build()
	File_mica_partner_partner_v1_partner_proto = out.File
	file_mica_partner_partner_v1_partner_proto_rawDesc = nil
	file_mica_partner_partner_v1_partner_proto_goTypes = nil
	file_mica_partner_partner_v1_partner_proto_depIdxs = nil
}
