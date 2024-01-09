// Copyright (c) 2022 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: mica/partner/uuek/v1/uuek.proto

package uuekv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "github.com/1080network/golang/partner/proto/micashared/common/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RemoveUUEKResponse_Status int32

const (
	RemoveUUEKResponse_STATUS_UNSPECIFIED RemoveUUEKResponse_Status = 0
	RemoveUUEKResponse_STATUS_SUCCESS     RemoveUUEKResponse_Status = 1
	RemoveUUEKResponse_STATUS_ERROR       RemoveUUEKResponse_Status = 2
	RemoveUUEKResponse_STATUS_NOT_FOUND   RemoveUUEKResponse_Status = 3
)

// Enum value maps for RemoveUUEKResponse_Status.
var (
	RemoveUUEKResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	RemoveUUEKResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x RemoveUUEKResponse_Status) Enum() *RemoveUUEKResponse_Status {
	p := new(RemoveUUEKResponse_Status)
	*p = x
	return p
}

func (x RemoveUUEKResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RemoveUUEKResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_partner_uuek_v1_uuek_proto_enumTypes[0].Descriptor()
}

func (RemoveUUEKResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_partner_uuek_v1_uuek_proto_enumTypes[0]
}

func (x RemoveUUEKResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RemoveUUEKResponse_Status.Descriptor instead.
func (RemoveUUEKResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_partner_uuek_v1_uuek_proto_rawDescGZIP(), []int{2, 0}
}

type ExchangeUUEKResponse_Status int32

const (
	ExchangeUUEKResponse_STATUS_UNSPECIFIED ExchangeUUEKResponse_Status = 0
	ExchangeUUEKResponse_STATUS_SUCCESS     ExchangeUUEKResponse_Status = 1
	ExchangeUUEKResponse_STATUS_ERROR       ExchangeUUEKResponse_Status = 2
	ExchangeUUEKResponse_STATUS_NOT_FOUND   ExchangeUUEKResponse_Status = 3
)

// Enum value maps for ExchangeUUEKResponse_Status.
var (
	ExchangeUUEKResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	ExchangeUUEKResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x ExchangeUUEKResponse_Status) Enum() *ExchangeUUEKResponse_Status {
	p := new(ExchangeUUEKResponse_Status)
	*p = x
	return p
}

func (x ExchangeUUEKResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExchangeUUEKResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_partner_uuek_v1_uuek_proto_enumTypes[1].Descriptor()
}

func (ExchangeUUEKResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_partner_uuek_v1_uuek_proto_enumTypes[1]
}

func (x ExchangeUUEKResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExchangeUUEKResponse_Status.Descriptor instead.
func (ExchangeUUEKResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_partner_uuek_v1_uuek_proto_rawDescGZIP(), []int{4, 0}
}

type SearchUUEKResponse_Status int32

const (
	SearchUUEKResponse_STATUS_UNSPECIFIED SearchUUEKResponse_Status = 0
	SearchUUEKResponse_STATUS_SUCCESS     SearchUUEKResponse_Status = 1
	SearchUUEKResponse_STATUS_ERROR       SearchUUEKResponse_Status = 2
)

// Enum value maps for SearchUUEKResponse_Status.
var (
	SearchUUEKResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
	}
	SearchUUEKResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
	}
)

func (x SearchUUEKResponse_Status) Enum() *SearchUUEKResponse_Status {
	p := new(SearchUUEKResponse_Status)
	*p = x
	return p
}

func (x SearchUUEKResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchUUEKResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_partner_uuek_v1_uuek_proto_enumTypes[2].Descriptor()
}

func (SearchUUEKResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_partner_uuek_v1_uuek_proto_enumTypes[2]
}

func (x SearchUUEKResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchUUEKResponse_Status.Descriptor instead.
func (SearchUUEKResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_partner_uuek_v1_uuek_proto_rawDescGZIP(), []int{6, 0}
}

type UUEK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// reference to the user at the Partner
	PartnerUserRef string `protobuf:"bytes,3,opt,name=partner_user_ref,json=partnerUserRef,proto3" json:"partner_user_ref,omitempty"`
	// reference to the user's instrument at the Partner
	PartnerInstrumentRef string `protobuf:"bytes,1,opt,name=partner_instrument_ref,json=partnerInstrumentRef,proto3" json:"partner_instrument_ref,omitempty"`
	// depending on how the Partner handles instruments and UUEKs, this will either be their reference to this
	// UUEK or a random UUID value if they don't have this concept
	PartnerUuekRef string         `protobuf:"bytes,4,opt,name=partner_uuek_ref,json=partnerUuekRef,proto3" json:"partner_uuek_ref,omitempty"`
	Uuek           *v1.CommonUUEK `protobuf:"bytes,2,opt,name=uuek,proto3" json:"uuek,omitempty"`
}

func (x *UUEK) Reset() {
	*x = UUEK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUEK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUEK) ProtoMessage() {}

func (x *UUEK) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUEK.ProtoReflect.Descriptor instead.
func (*UUEK) Descriptor() ([]byte, []int) {
	return file_mica_partner_uuek_v1_uuek_proto_rawDescGZIP(), []int{0}
}

func (x *UUEK) GetPartnerUserRef() string {
	if x != nil {
		return x.PartnerUserRef
	}
	return ""
}

func (x *UUEK) GetPartnerInstrumentRef() string {
	if x != nil {
		return x.PartnerInstrumentRef
	}
	return ""
}

func (x *UUEK) GetPartnerUuekRef() string {
	if x != nil {
		return x.PartnerUuekRef
	}
	return ""
}

func (x *UUEK) GetUuek() *v1.CommonUUEK {
	if x != nil {
		return x.Uuek
	}
	return nil
}

type RemoveUUEKRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The UUEK that the Partner can use to obtain or return value for this user.
	PartnerUuek string `protobuf:"bytes,2,opt,name=partner_uuek,json=partnerUuek,proto3" json:"partner_uuek,omitempty"`
}

func (x *RemoveUUEKRequest) Reset() {
	*x = RemoveUUEKRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUUEKRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUUEKRequest) ProtoMessage() {}

func (x *RemoveUUEKRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUUEKRequest.ProtoReflect.Descriptor instead.
func (*RemoveUUEKRequest) Descriptor() ([]byte, []int) {
	return file_mica_partner_uuek_v1_uuek_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveUUEKRequest) GetPartnerUuek() string {
	if x != nil {
		return x.PartnerUuek
	}
	return ""
}

type RemoveUUEKResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status RemoveUUEKResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.partner.uuek.v1.RemoveUUEKResponse_Status" json:"status,omitempty"`
	Error  *v1.Error                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RemoveUUEKResponse) Reset() {
	*x = RemoveUUEKResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUUEKResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUUEKResponse) ProtoMessage() {}

func (x *RemoveUUEKResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUUEKResponse.ProtoReflect.Descriptor instead.
func (*RemoveUUEKResponse) Descriptor() ([]byte, []int) {
	return file_mica_partner_uuek_v1_uuek_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveUUEKResponse) GetStatus() RemoveUUEKResponse_Status {
	if x != nil {
		return x.Status
	}
	return RemoveUUEKResponse_STATUS_UNSPECIFIED
}

func (x *RemoveUUEKResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type ExchangeUUEKRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The UUEK that the Partner can use to obtain or return value for this user.
	PartnerUuek string `protobuf:"bytes,1,opt,name=partner_uuek,json=partnerUuek,proto3" json:"partner_uuek,omitempty"`
	// flag to indicate what to do with the UUEK that's being replaced
	RemoveExisting bool `protobuf:"varint,2,opt,name=remove_existing,json=removeExisting,proto3" json:"remove_existing,omitempty"`
	// Primary key on the external system.
	PartnerUuekRef string `protobuf:"bytes,3,opt,name=partner_uuek_ref,json=partnerUuekRef,proto3" json:"partner_uuek_ref,omitempty"`
}

func (x *ExchangeUUEKRequest) Reset() {
	*x = ExchangeUUEKRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeUUEKRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeUUEKRequest) ProtoMessage() {}

func (x *ExchangeUUEKRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeUUEKRequest.ProtoReflect.Descriptor instead.
func (*ExchangeUUEKRequest) Descriptor() ([]byte, []int) {
	return file_mica_partner_uuek_v1_uuek_proto_rawDescGZIP(), []int{3}
}

func (x *ExchangeUUEKRequest) GetPartnerUuek() string {
	if x != nil {
		return x.PartnerUuek
	}
	return ""
}

func (x *ExchangeUUEKRequest) GetRemoveExisting() bool {
	if x != nil {
		return x.RemoveExisting
	}
	return false
}

func (x *ExchangeUUEKRequest) GetPartnerUuekRef() string {
	if x != nil {
		return x.PartnerUuekRef
	}
	return ""
}

type ExchangeUUEKResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ExchangeUUEKResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.partner.uuek.v1.ExchangeUUEKResponse_Status" json:"status,omitempty"`
	Error  *v1.Error                   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// The UUEK that the Partner can use to obtain or return value for this user.
	PartnerUuek string `protobuf:"bytes,3,opt,name=partner_uuek,json=partnerUuek,proto3" json:"partner_uuek,omitempty"`
}

func (x *ExchangeUUEKResponse) Reset() {
	*x = ExchangeUUEKResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeUUEKResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeUUEKResponse) ProtoMessage() {}

func (x *ExchangeUUEKResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeUUEKResponse.ProtoReflect.Descriptor instead.
func (*ExchangeUUEKResponse) Descriptor() ([]byte, []int) {
	return file_mica_partner_uuek_v1_uuek_proto_rawDescGZIP(), []int{4}
}

func (x *ExchangeUUEKResponse) GetStatus() ExchangeUUEKResponse_Status {
	if x != nil {
		return x.Status
	}
	return ExchangeUUEKResponse_STATUS_UNSPECIFIED
}

func (x *ExchangeUUEKResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ExchangeUUEKResponse) GetPartnerUuek() string {
	if x != nil {
		return x.PartnerUuek
	}
	return ""
}

type SearchUUEKRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// reference to the user at the Partner
	PartnerUserRef string `protobuf:"bytes,3,opt,name=partner_user_ref,json=partnerUserRef,proto3" json:"partner_user_ref,omitempty"`
	// reference to the user's instrument at the Partner
	PartnerInstrumentRef string `protobuf:"bytes,1,opt,name=partner_instrument_ref,json=partnerInstrumentRef,proto3" json:"partner_instrument_ref,omitempty"`
	// depending on how the Partner handles instruments and UUEKs, this will either be their reference to this
	// UUEK or a random UUID value if they don't have this concept
	PartnerUuekRef string `protobuf:"bytes,4,opt,name=partner_uuek_ref,json=partnerUuekRef,proto3" json:"partner_uuek_ref,omitempty"`
}

func (x *SearchUUEKRequest) Reset() {
	*x = SearchUUEKRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUUEKRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUUEKRequest) ProtoMessage() {}

func (x *SearchUUEKRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUUEKRequest.ProtoReflect.Descriptor instead.
func (*SearchUUEKRequest) Descriptor() ([]byte, []int) {
	return file_mica_partner_uuek_v1_uuek_proto_rawDescGZIP(), []int{5}
}

func (x *SearchUUEKRequest) GetPartnerUserRef() string {
	if x != nil {
		return x.PartnerUserRef
	}
	return ""
}

func (x *SearchUUEKRequest) GetPartnerInstrumentRef() string {
	if x != nil {
		return x.PartnerInstrumentRef
	}
	return ""
}

func (x *SearchUUEKRequest) GetPartnerUuekRef() string {
	if x != nil {
		return x.PartnerUuekRef
	}
	return ""
}

type SearchUUEKResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status SearchUUEKResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.partner.uuek.v1.SearchUUEKResponse_Status" json:"status,omitempty"`
	Error  *v1.Error                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Uueks  []*UUEK                   `protobuf:"bytes,3,rep,name=uueks,proto3" json:"uueks,omitempty"`
}

func (x *SearchUUEKResponse) Reset() {
	*x = SearchUUEKResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUUEKResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUUEKResponse) ProtoMessage() {}

func (x *SearchUUEKResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUUEKResponse.ProtoReflect.Descriptor instead.
func (*SearchUUEKResponse) Descriptor() ([]byte, []int) {
	return file_mica_partner_uuek_v1_uuek_proto_rawDescGZIP(), []int{6}
}

func (x *SearchUUEKResponse) GetStatus() SearchUUEKResponse_Status {
	if x != nil {
		return x.Status
	}
	return SearchUUEKResponse_STATUS_UNSPECIFIED
}

func (x *SearchUUEKResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SearchUUEKResponse) GetUueks() []*UUEK {
	if x != nil {
		return x.Uueks
	}
	return nil
}

var File_mica_partner_uuek_v1_uuek_proto protoreflect.FileDescriptor

var file_mica_partner_uuek_v1_uuek_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x75,
	0x75, 0x65, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x75, 0x65, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e,
	0x75, 0x75, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x75, 0x65, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x04, 0x55, 0x55, 0x45, 0x4b, 0x12, 0x33, 0x0a, 0x10,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x66,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18,
	0x64, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x66, 0x12, 0x3f, 0x0a, 0x16, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x14, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x66, 0x12, 0x33, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x75,
	0x65, 0x6b, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x55, 0x75, 0x65, 0x6b, 0x52, 0x65, 0x66, 0x12, 0x34, 0x0a, 0x04, 0x75, 0x75, 0x65, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x55, 0x55, 0x45, 0x4b, 0x52, 0x04, 0x75, 0x75, 0x65, 0x6b, 0x22, 0x41, 0x0a,
	0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x55, 0x45, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x75,
	0x65, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x1e, 0x18, 0x32, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x55, 0x75, 0x65, 0x6b,
	0x22, 0xee, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x55, 0x45, 0x4b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x75, 0x75, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x55, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x03, 0x22, 0xa1, 0x01, 0x0a, 0x13, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x55,
	0x45, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0c, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x65, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x55, 0x75, 0x65, 0x6b, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x33, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x65, 0x6b,
	0x5f, 0x72, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x55, 0x75,
	0x65, 0x6b, 0x52, 0x65, 0x66, 0x22, 0xa0, 0x02, 0x0a, 0x14, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x55, 0x55, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x75, 0x75,
	0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x55,
	0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x0c,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x65, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52, 0x0b, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x55, 0x75, 0x65, 0x6b, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x22, 0xb8, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x55, 0x55, 0x45, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31,
	0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18,
	0x64, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x66, 0x12, 0x3d, 0x0a, 0x16, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x64, 0x52, 0x14, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66,
	0x12, 0x31, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x65, 0x6b,
	0x5f, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x18, 0x64, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x55, 0x75, 0x65, 0x6b,
	0x52, 0x65, 0x66, 0x22, 0x8a, 0x02, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x55,
	0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x75, 0x75, 0x65, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x55, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x30, 0x0a, 0x05, 0x75, 0x75, 0x65, 0x6b, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2e, 0x75, 0x75, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x55, 0x45,
	0x4b, 0x52, 0x05, 0x75, 0x75, 0x65, 0x6b, 0x73, 0x22, 0x46, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02,
	0x42, 0x40, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2e, 0x75, 0x75, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x55, 0x55, 0x45,
	0x4b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x13, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2f, 0x75, 0x75, 0x65, 0x6b, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49,
	0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mica_partner_uuek_v1_uuek_proto_rawDescOnce sync.Once
	file_mica_partner_uuek_v1_uuek_proto_rawDescData = file_mica_partner_uuek_v1_uuek_proto_rawDesc
)

func file_mica_partner_uuek_v1_uuek_proto_rawDescGZIP() []byte {
	file_mica_partner_uuek_v1_uuek_proto_rawDescOnce.Do(func() {
		file_mica_partner_uuek_v1_uuek_proto_rawDescData = protoimpl.X.CompressGZIP(file_mica_partner_uuek_v1_uuek_proto_rawDescData)
	})
	return file_mica_partner_uuek_v1_uuek_proto_rawDescData
}

var file_mica_partner_uuek_v1_uuek_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_mica_partner_uuek_v1_uuek_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mica_partner_uuek_v1_uuek_proto_goTypes = []interface{}{
	(RemoveUUEKResponse_Status)(0),   // 0: mica.partner.uuek.v1.RemoveUUEKResponse.Status
	(ExchangeUUEKResponse_Status)(0), // 1: mica.partner.uuek.v1.ExchangeUUEKResponse.Status
	(SearchUUEKResponse_Status)(0),   // 2: mica.partner.uuek.v1.SearchUUEKResponse.Status
	(*UUEK)(nil),                     // 3: mica.partner.uuek.v1.UUEK
	(*RemoveUUEKRequest)(nil),        // 4: mica.partner.uuek.v1.RemoveUUEKRequest
	(*RemoveUUEKResponse)(nil),       // 5: mica.partner.uuek.v1.RemoveUUEKResponse
	(*ExchangeUUEKRequest)(nil),      // 6: mica.partner.uuek.v1.ExchangeUUEKRequest
	(*ExchangeUUEKResponse)(nil),     // 7: mica.partner.uuek.v1.ExchangeUUEKResponse
	(*SearchUUEKRequest)(nil),        // 8: mica.partner.uuek.v1.SearchUUEKRequest
	(*SearchUUEKResponse)(nil),       // 9: mica.partner.uuek.v1.SearchUUEKResponse
	(*v1.CommonUUEK)(nil),            // 10: micashared.common.v1.CommonUUEK
	(*v1.Error)(nil),                 // 11: micashared.common.v1.Error
}
var file_mica_partner_uuek_v1_uuek_proto_depIdxs = []int32{
	10, // 0: mica.partner.uuek.v1.UUEK.uuek:type_name -> micashared.common.v1.CommonUUEK
	0,  // 1: mica.partner.uuek.v1.RemoveUUEKResponse.status:type_name -> mica.partner.uuek.v1.RemoveUUEKResponse.Status
	11, // 2: mica.partner.uuek.v1.RemoveUUEKResponse.error:type_name -> micashared.common.v1.Error
	1,  // 3: mica.partner.uuek.v1.ExchangeUUEKResponse.status:type_name -> mica.partner.uuek.v1.ExchangeUUEKResponse.Status
	11, // 4: mica.partner.uuek.v1.ExchangeUUEKResponse.error:type_name -> micashared.common.v1.Error
	2,  // 5: mica.partner.uuek.v1.SearchUUEKResponse.status:type_name -> mica.partner.uuek.v1.SearchUUEKResponse.Status
	11, // 6: mica.partner.uuek.v1.SearchUUEKResponse.error:type_name -> micashared.common.v1.Error
	3,  // 7: mica.partner.uuek.v1.SearchUUEKResponse.uueks:type_name -> mica.partner.uuek.v1.UUEK
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_mica_partner_uuek_v1_uuek_proto_init() }
func file_mica_partner_uuek_v1_uuek_proto_init() {
	if File_mica_partner_uuek_v1_uuek_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mica_partner_uuek_v1_uuek_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUEK); i {
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
		file_mica_partner_uuek_v1_uuek_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUUEKRequest); i {
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
		file_mica_partner_uuek_v1_uuek_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUUEKResponse); i {
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
		file_mica_partner_uuek_v1_uuek_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeUUEKRequest); i {
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
		file_mica_partner_uuek_v1_uuek_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeUUEKResponse); i {
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
		file_mica_partner_uuek_v1_uuek_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUUEKRequest); i {
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
		file_mica_partner_uuek_v1_uuek_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUUEKResponse); i {
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
			RawDescriptor: file_mica_partner_uuek_v1_uuek_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mica_partner_uuek_v1_uuek_proto_goTypes,
		DependencyIndexes: file_mica_partner_uuek_v1_uuek_proto_depIdxs,
		EnumInfos:         file_mica_partner_uuek_v1_uuek_proto_enumTypes,
		MessageInfos:      file_mica_partner_uuek_v1_uuek_proto_msgTypes,
	}.Build()
	File_mica_partner_uuek_v1_uuek_proto = out.File
	file_mica_partner_uuek_v1_uuek_proto_rawDesc = nil
	file_mica_partner_uuek_v1_uuek_proto_goTypes = nil
	file_mica_partner_uuek_v1_uuek_proto_depIdxs = nil
}
