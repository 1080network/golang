// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: partner/paymenttoken/v1/payment_token.proto

package paymenttokenv1

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	v1 "ten80/proto/common/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RemovePaymentTokenResponse_Status int32

const (
	RemovePaymentTokenResponse_STATUS_UNSPECIFIED RemovePaymentTokenResponse_Status = 0
	RemovePaymentTokenResponse_STATUS_SUCCESS     RemovePaymentTokenResponse_Status = 1
	RemovePaymentTokenResponse_STATUS_ERROR       RemovePaymentTokenResponse_Status = 2
	RemovePaymentTokenResponse_STATUS_NOT_FOUND   RemovePaymentTokenResponse_Status = 3
)

// Enum value maps for RemovePaymentTokenResponse_Status.
var (
	RemovePaymentTokenResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	RemovePaymentTokenResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x RemovePaymentTokenResponse_Status) Enum() *RemovePaymentTokenResponse_Status {
	p := new(RemovePaymentTokenResponse_Status)
	*p = x
	return p
}

func (x RemovePaymentTokenResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RemovePaymentTokenResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_partner_paymenttoken_v1_payment_token_proto_enumTypes[0].Descriptor()
}

func (RemovePaymentTokenResponse_Status) Type() protoreflect.EnumType {
	return &file_partner_paymenttoken_v1_payment_token_proto_enumTypes[0]
}

func (x RemovePaymentTokenResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RemovePaymentTokenResponse_Status.Descriptor instead.
func (RemovePaymentTokenResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_partner_paymenttoken_v1_payment_token_proto_rawDescGZIP(), []int{2, 0}
}

type ExchangePaymentTokenResponse_Status int32

const (
	ExchangePaymentTokenResponse_STATUS_UNSPECIFIED ExchangePaymentTokenResponse_Status = 0
	ExchangePaymentTokenResponse_STATUS_SUCCESS     ExchangePaymentTokenResponse_Status = 1
	ExchangePaymentTokenResponse_STATUS_ERROR       ExchangePaymentTokenResponse_Status = 2
	ExchangePaymentTokenResponse_STATUS_NOT_FOUND   ExchangePaymentTokenResponse_Status = 3
)

// Enum value maps for ExchangePaymentTokenResponse_Status.
var (
	ExchangePaymentTokenResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	ExchangePaymentTokenResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x ExchangePaymentTokenResponse_Status) Enum() *ExchangePaymentTokenResponse_Status {
	p := new(ExchangePaymentTokenResponse_Status)
	*p = x
	return p
}

func (x ExchangePaymentTokenResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExchangePaymentTokenResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_partner_paymenttoken_v1_payment_token_proto_enumTypes[1].Descriptor()
}

func (ExchangePaymentTokenResponse_Status) Type() protoreflect.EnumType {
	return &file_partner_paymenttoken_v1_payment_token_proto_enumTypes[1]
}

func (x ExchangePaymentTokenResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExchangePaymentTokenResponse_Status.Descriptor instead.
func (ExchangePaymentTokenResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_partner_paymenttoken_v1_payment_token_proto_rawDescGZIP(), []int{4, 0}
}

type SearchPartnerPaymentTokenResponse_Status int32

const (
	SearchPartnerPaymentTokenResponse_STATUS_UNSPECIFIED SearchPartnerPaymentTokenResponse_Status = 0
	SearchPartnerPaymentTokenResponse_STATUS_SUCCESS     SearchPartnerPaymentTokenResponse_Status = 1
	SearchPartnerPaymentTokenResponse_STATUS_ERROR       SearchPartnerPaymentTokenResponse_Status = 2
)

// Enum value maps for SearchPartnerPaymentTokenResponse_Status.
var (
	SearchPartnerPaymentTokenResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
	}
	SearchPartnerPaymentTokenResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
	}
)

func (x SearchPartnerPaymentTokenResponse_Status) Enum() *SearchPartnerPaymentTokenResponse_Status {
	p := new(SearchPartnerPaymentTokenResponse_Status)
	*p = x
	return p
}

func (x SearchPartnerPaymentTokenResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchPartnerPaymentTokenResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_partner_paymenttoken_v1_payment_token_proto_enumTypes[2].Descriptor()
}

func (SearchPartnerPaymentTokenResponse_Status) Type() protoreflect.EnumType {
	return &file_partner_paymenttoken_v1_payment_token_proto_enumTypes[2]
}

func (x SearchPartnerPaymentTokenResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchPartnerPaymentTokenResponse_Status.Descriptor instead.
func (SearchPartnerPaymentTokenResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_partner_paymenttoken_v1_payment_token_proto_rawDescGZIP(), []int{6, 0}
}

type PaymentToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartnerInstrumentRef string                 `protobuf:"bytes,1,opt,name=partner_instrument_ref,json=partnerInstrumentRef,proto3" json:"partner_instrument_ref,omitempty"`
	PaymentToken         *v1.CommonPaymentToken `protobuf:"bytes,2,opt,name=payment_token,json=paymentToken,proto3" json:"payment_token,omitempty"`
}

func (x *PaymentToken) Reset() {
	*x = PaymentToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentToken) ProtoMessage() {}

func (x *PaymentToken) ProtoReflect() protoreflect.Message {
	mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentToken.ProtoReflect.Descriptor instead.
func (*PaymentToken) Descriptor() ([]byte, []int) {
	return file_partner_paymenttoken_v1_payment_token_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentToken) GetPartnerInstrumentRef() string {
	if x != nil {
		return x.PartnerInstrumentRef
	}
	return ""
}

func (x *PaymentToken) GetPaymentToken() *v1.CommonPaymentToken {
	if x != nil {
		return x.PaymentToken
	}
	return nil
}

type RemovePaymentTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The payment token that the Partner can use to obtain or return funds for this user.
	PartnerPaymentToken string `protobuf:"bytes,2,opt,name=partner_payment_token,json=partnerPaymentToken,proto3" json:"partner_payment_token,omitempty"`
}

func (x *RemovePaymentTokenRequest) Reset() {
	*x = RemovePaymentTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePaymentTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePaymentTokenRequest) ProtoMessage() {}

func (x *RemovePaymentTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePaymentTokenRequest.ProtoReflect.Descriptor instead.
func (*RemovePaymentTokenRequest) Descriptor() ([]byte, []int) {
	return file_partner_paymenttoken_v1_payment_token_proto_rawDescGZIP(), []int{1}
}

func (x *RemovePaymentTokenRequest) GetPartnerPaymentToken() string {
	if x != nil {
		return x.PartnerPaymentToken
	}
	return ""
}

type RemovePaymentTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status RemovePaymentTokenResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=partner.paymenttoken.v1.RemovePaymentTokenResponse_Status" json:"status,omitempty"`
	Error  *v1.Error                         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RemovePaymentTokenResponse) Reset() {
	*x = RemovePaymentTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePaymentTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePaymentTokenResponse) ProtoMessage() {}

func (x *RemovePaymentTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePaymentTokenResponse.ProtoReflect.Descriptor instead.
func (*RemovePaymentTokenResponse) Descriptor() ([]byte, []int) {
	return file_partner_paymenttoken_v1_payment_token_proto_rawDescGZIP(), []int{2}
}

func (x *RemovePaymentTokenResponse) GetStatus() RemovePaymentTokenResponse_Status {
	if x != nil {
		return x.Status
	}
	return RemovePaymentTokenResponse_STATUS_UNSPECIFIED
}

func (x *RemovePaymentTokenResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type ExchangePaymentTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The payment token that the Partner can use to obtain or return funds for this user.
	PartnerPaymentToken string `protobuf:"bytes,1,opt,name=partner_payment_token,json=partnerPaymentToken,proto3" json:"partner_payment_token,omitempty"`
	// flag to indicate what to do with the payment token that's being replaced
	RemoveExisting bool `protobuf:"varint,2,opt,name=remove_existing,json=removeExisting,proto3" json:"remove_existing,omitempty"`
}

func (x *ExchangePaymentTokenRequest) Reset() {
	*x = ExchangePaymentTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangePaymentTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangePaymentTokenRequest) ProtoMessage() {}

func (x *ExchangePaymentTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangePaymentTokenRequest.ProtoReflect.Descriptor instead.
func (*ExchangePaymentTokenRequest) Descriptor() ([]byte, []int) {
	return file_partner_paymenttoken_v1_payment_token_proto_rawDescGZIP(), []int{3}
}

func (x *ExchangePaymentTokenRequest) GetPartnerPaymentToken() string {
	if x != nil {
		return x.PartnerPaymentToken
	}
	return ""
}

func (x *ExchangePaymentTokenRequest) GetRemoveExisting() bool {
	if x != nil {
		return x.RemoveExisting
	}
	return false
}

type ExchangePaymentTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ExchangePaymentTokenResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=partner.paymenttoken.v1.ExchangePaymentTokenResponse_Status" json:"status,omitempty"`
	Error  *v1.Error                           `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// The payment token that the Partner can use to obtain or return funds for this user.
	PartnerPaymentToken string `protobuf:"bytes,3,opt,name=partner_payment_token,json=partnerPaymentToken,proto3" json:"partner_payment_token,omitempty"`
}

func (x *ExchangePaymentTokenResponse) Reset() {
	*x = ExchangePaymentTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangePaymentTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangePaymentTokenResponse) ProtoMessage() {}

func (x *ExchangePaymentTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangePaymentTokenResponse.ProtoReflect.Descriptor instead.
func (*ExchangePaymentTokenResponse) Descriptor() ([]byte, []int) {
	return file_partner_paymenttoken_v1_payment_token_proto_rawDescGZIP(), []int{4}
}

func (x *ExchangePaymentTokenResponse) GetStatus() ExchangePaymentTokenResponse_Status {
	if x != nil {
		return x.Status
	}
	return ExchangePaymentTokenResponse_STATUS_UNSPECIFIED
}

func (x *ExchangePaymentTokenResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ExchangePaymentTokenResponse) GetPartnerPaymentToken() string {
	if x != nil {
		return x.PartnerPaymentToken
	}
	return ""
}

type SearchPartnerPaymentTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartnerInstrumentRef string `protobuf:"bytes,1,opt,name=partner_instrument_ref,json=partnerInstrumentRef,proto3" json:"partner_instrument_ref,omitempty"`
}

func (x *SearchPartnerPaymentTokenRequest) Reset() {
	*x = SearchPartnerPaymentTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPartnerPaymentTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPartnerPaymentTokenRequest) ProtoMessage() {}

func (x *SearchPartnerPaymentTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPartnerPaymentTokenRequest.ProtoReflect.Descriptor instead.
func (*SearchPartnerPaymentTokenRequest) Descriptor() ([]byte, []int) {
	return file_partner_paymenttoken_v1_payment_token_proto_rawDescGZIP(), []int{5}
}

func (x *SearchPartnerPaymentTokenRequest) GetPartnerInstrumentRef() string {
	if x != nil {
		return x.PartnerInstrumentRef
	}
	return ""
}

type SearchPartnerPaymentTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        SearchPartnerPaymentTokenResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=partner.paymenttoken.v1.SearchPartnerPaymentTokenResponse_Status" json:"status,omitempty"`
	Error         *v1.Error                                `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	PaymentTokens []*PaymentToken                          `protobuf:"bytes,3,rep,name=payment_tokens,json=paymentTokens,proto3" json:"payment_tokens,omitempty"`
}

func (x *SearchPartnerPaymentTokenResponse) Reset() {
	*x = SearchPartnerPaymentTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPartnerPaymentTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPartnerPaymentTokenResponse) ProtoMessage() {}

func (x *SearchPartnerPaymentTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_partner_paymenttoken_v1_payment_token_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPartnerPaymentTokenResponse.ProtoReflect.Descriptor instead.
func (*SearchPartnerPaymentTokenResponse) Descriptor() ([]byte, []int) {
	return file_partner_paymenttoken_v1_payment_token_proto_rawDescGZIP(), []int{6}
}

func (x *SearchPartnerPaymentTokenResponse) GetStatus() SearchPartnerPaymentTokenResponse_Status {
	if x != nil {
		return x.Status
	}
	return SearchPartnerPaymentTokenResponse_STATUS_UNSPECIFIED
}

func (x *SearchPartnerPaymentTokenResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SearchPartnerPaymentTokenResponse) GetPaymentTokens() []*PaymentToken {
	if x != nil {
		return x.PaymentTokens
	}
	return nil
}

var File_partner_paymenttoken_v1_payment_token_proto protoreflect.FileDescriptor

var file_partner_paymenttoken_v1_payment_token_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a,
	0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x34, 0x0a,
	0x16, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x66, 0x12, 0x42, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4f, 0x0a, 0x19, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xf6, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x03, 0x22, 0x7a, 0x0a, 0x1b, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xae, 0x02,
	0x0a, 0x1c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x15,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x22, 0x58,
	0x0a, 0x20, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x14, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x22, 0xbc, 0x02, 0x0a, 0x21, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x4c, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22,
	0x46, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x42, 0x66, 0x0a, 0x25, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x74, 0x65, 0x6e, 0x38, 0x30, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31,
	0x42, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x22, 0x74, 0x65, 0x6e, 0x38, 0x30, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x54, 0x45, 0x4e, 0x38, 0x30, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_partner_paymenttoken_v1_payment_token_proto_rawDescOnce sync.Once
	file_partner_paymenttoken_v1_payment_token_proto_rawDescData = file_partner_paymenttoken_v1_payment_token_proto_rawDesc
)

func file_partner_paymenttoken_v1_payment_token_proto_rawDescGZIP() []byte {
	file_partner_paymenttoken_v1_payment_token_proto_rawDescOnce.Do(func() {
		file_partner_paymenttoken_v1_payment_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_partner_paymenttoken_v1_payment_token_proto_rawDescData)
	})
	return file_partner_paymenttoken_v1_payment_token_proto_rawDescData
}

var file_partner_paymenttoken_v1_payment_token_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_partner_paymenttoken_v1_payment_token_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_partner_paymenttoken_v1_payment_token_proto_goTypes = []interface{}{
	(RemovePaymentTokenResponse_Status)(0),        // 0: partner.paymenttoken.v1.RemovePaymentTokenResponse.Status
	(ExchangePaymentTokenResponse_Status)(0),      // 1: partner.paymenttoken.v1.ExchangePaymentTokenResponse.Status
	(SearchPartnerPaymentTokenResponse_Status)(0), // 2: partner.paymenttoken.v1.SearchPartnerPaymentTokenResponse.Status
	(*PaymentToken)(nil),                          // 3: partner.paymenttoken.v1.PaymentToken
	(*RemovePaymentTokenRequest)(nil),             // 4: partner.paymenttoken.v1.RemovePaymentTokenRequest
	(*RemovePaymentTokenResponse)(nil),            // 5: partner.paymenttoken.v1.RemovePaymentTokenResponse
	(*ExchangePaymentTokenRequest)(nil),           // 6: partner.paymenttoken.v1.ExchangePaymentTokenRequest
	(*ExchangePaymentTokenResponse)(nil),          // 7: partner.paymenttoken.v1.ExchangePaymentTokenResponse
	(*SearchPartnerPaymentTokenRequest)(nil),      // 8: partner.paymenttoken.v1.SearchPartnerPaymentTokenRequest
	(*SearchPartnerPaymentTokenResponse)(nil),     // 9: partner.paymenttoken.v1.SearchPartnerPaymentTokenResponse
	(*v1.CommonPaymentToken)(nil),                 // 10: common.v1.CommonPaymentToken
	(*v1.Error)(nil),                              // 11: common.v1.Error
}
var file_partner_paymenttoken_v1_payment_token_proto_depIdxs = []int32{
	10, // 0: partner.paymenttoken.v1.PaymentToken.payment_token:type_name -> common.v1.CommonPaymentToken
	0,  // 1: partner.paymenttoken.v1.RemovePaymentTokenResponse.status:type_name -> partner.paymenttoken.v1.RemovePaymentTokenResponse.Status
	11, // 2: partner.paymenttoken.v1.RemovePaymentTokenResponse.error:type_name -> common.v1.Error
	1,  // 3: partner.paymenttoken.v1.ExchangePaymentTokenResponse.status:type_name -> partner.paymenttoken.v1.ExchangePaymentTokenResponse.Status
	11, // 4: partner.paymenttoken.v1.ExchangePaymentTokenResponse.error:type_name -> common.v1.Error
	2,  // 5: partner.paymenttoken.v1.SearchPartnerPaymentTokenResponse.status:type_name -> partner.paymenttoken.v1.SearchPartnerPaymentTokenResponse.Status
	11, // 6: partner.paymenttoken.v1.SearchPartnerPaymentTokenResponse.error:type_name -> common.v1.Error
	3,  // 7: partner.paymenttoken.v1.SearchPartnerPaymentTokenResponse.payment_tokens:type_name -> partner.paymenttoken.v1.PaymentToken
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_partner_paymenttoken_v1_payment_token_proto_init() }
func file_partner_paymenttoken_v1_payment_token_proto_init() {
	if File_partner_paymenttoken_v1_payment_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_partner_paymenttoken_v1_payment_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentToken); i {
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
		file_partner_paymenttoken_v1_payment_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePaymentTokenRequest); i {
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
		file_partner_paymenttoken_v1_payment_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePaymentTokenResponse); i {
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
		file_partner_paymenttoken_v1_payment_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangePaymentTokenRequest); i {
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
		file_partner_paymenttoken_v1_payment_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangePaymentTokenResponse); i {
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
		file_partner_paymenttoken_v1_payment_token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPartnerPaymentTokenRequest); i {
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
		file_partner_paymenttoken_v1_payment_token_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPartnerPaymentTokenResponse); i {
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
			RawDescriptor: file_partner_paymenttoken_v1_payment_token_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_partner_paymenttoken_v1_payment_token_proto_goTypes,
		DependencyIndexes: file_partner_paymenttoken_v1_payment_token_proto_depIdxs,
		EnumInfos:         file_partner_paymenttoken_v1_payment_token_proto_enumTypes,
		MessageInfos:      file_partner_paymenttoken_v1_payment_token_proto_msgTypes,
	}.Build()
	File_partner_paymenttoken_v1_payment_token_proto = out.File
	file_partner_paymenttoken_v1_payment_token_proto_rawDesc = nil
	file_partner_paymenttoken_v1_payment_token_proto_goTypes = nil
	file_partner_paymenttoken_v1_payment_token_proto_depIdxs = nil
}
