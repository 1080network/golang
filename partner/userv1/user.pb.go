// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: partner/user/v1/user.proto

package userv1

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type EnrollUserInstrumentResponse_Status int32

const (
	EnrollUserInstrumentResponse_STATUS_UNSPECIFIED         EnrollUserInstrumentResponse_Status = 0
	EnrollUserInstrumentResponse_STATUS_SUCCESS             EnrollUserInstrumentResponse_Status = 1
	EnrollUserInstrumentResponse_STATUS_ERROR               EnrollUserInstrumentResponse_Status = 2
	EnrollUserInstrumentResponse_STATUS_PARTIAL_APPROVAL    EnrollUserInstrumentResponse_Status = 3
	EnrollUserInstrumentResponse_STATUS_INSUFFICIENT_FUNDS  EnrollUserInstrumentResponse_Status = 4
	EnrollUserInstrumentResponse_STATUS_INELIGIBLE_PRODUCTS EnrollUserInstrumentResponse_Status = 5
	EnrollUserInstrumentResponse_STATUS_USER_CLOSED         EnrollUserInstrumentResponse_Status = 6
)

// Enum value maps for EnrollUserInstrumentResponse_Status.
var (
	EnrollUserInstrumentResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_PARTIAL_APPROVAL",
		4: "STATUS_INSUFFICIENT_FUNDS",
		5: "STATUS_INELIGIBLE_PRODUCTS",
		6: "STATUS_USER_CLOSED",
	}
	EnrollUserInstrumentResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED":         0,
		"STATUS_SUCCESS":             1,
		"STATUS_ERROR":               2,
		"STATUS_PARTIAL_APPROVAL":    3,
		"STATUS_INSUFFICIENT_FUNDS":  4,
		"STATUS_INELIGIBLE_PRODUCTS": 5,
		"STATUS_USER_CLOSED":         6,
	}
)

func (x EnrollUserInstrumentResponse_Status) Enum() *EnrollUserInstrumentResponse_Status {
	p := new(EnrollUserInstrumentResponse_Status)
	*p = x
	return p
}

func (x EnrollUserInstrumentResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnrollUserInstrumentResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_partner_user_v1_user_proto_enumTypes[0].Descriptor()
}

func (EnrollUserInstrumentResponse_Status) Type() protoreflect.EnumType {
	return &file_partner_user_v1_user_proto_enumTypes[0]
}

func (x EnrollUserInstrumentResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnrollUserInstrumentResponse_Status.Descriptor instead.
func (EnrollUserInstrumentResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_partner_user_v1_user_proto_rawDescGZIP(), []int{2, 0}
}

type SearchUserInstrumentResponse_Status int32

const (
	SearchUserInstrumentResponse_STATUS_UNSPECIFIED SearchUserInstrumentResponse_Status = 0
	SearchUserInstrumentResponse_STATUS_SUCCESS     SearchUserInstrumentResponse_Status = 1
	SearchUserInstrumentResponse_STATUS_ERROR       SearchUserInstrumentResponse_Status = 2
)

// Enum value maps for SearchUserInstrumentResponse_Status.
var (
	SearchUserInstrumentResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
	}
	SearchUserInstrumentResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
	}
)

func (x SearchUserInstrumentResponse_Status) Enum() *SearchUserInstrumentResponse_Status {
	p := new(SearchUserInstrumentResponse_Status)
	*p = x
	return p
}

func (x SearchUserInstrumentResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchUserInstrumentResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_partner_user_v1_user_proto_enumTypes[1].Descriptor()
}

func (SearchUserInstrumentResponse_Status) Type() protoreflect.EnumType {
	return &file_partner_user_v1_user_proto_enumTypes[1]
}

func (x SearchUserInstrumentResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchUserInstrumentResponse_Status.Descriptor instead.
func (SearchUserInstrumentResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_partner_user_v1_user_proto_rawDescGZIP(), []int{4, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserKey string               `protobuf:"bytes,1,opt,name=user_key,json=userKey,proto3" json:"user_key,omitempty"`
	Created *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Version int64                `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	// Details of the user.
	UserDemographic *v1.UserDemographic `protobuf:"bytes,5,opt,name=user_demographic,json=userDemographic,proto3" json:"user_demographic,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_partner_user_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_partner_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserKey() string {
	if x != nil {
		return x.UserKey
	}
	return ""
}

func (x *User) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *User) GetUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *User) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *User) GetUserDemographic() *v1.UserDemographic {
	if x != nil {
		return x.UserDemographic
	}
	return nil
}

type EnrollUserInstrumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The payment token that the Partner can use to obtain or return funds for this user.
	PartnerPaymentToken string              `protobuf:"bytes,1,opt,name=partner_payment_token,json=partnerPaymentToken,proto3" json:"partner_payment_token,omitempty"`
	User                *v1.UserDemographic `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *EnrollUserInstrumentRequest) Reset() {
	*x = EnrollUserInstrumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_user_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollUserInstrumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollUserInstrumentRequest) ProtoMessage() {}

func (x *EnrollUserInstrumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_partner_user_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollUserInstrumentRequest.ProtoReflect.Descriptor instead.
func (*EnrollUserInstrumentRequest) Descriptor() ([]byte, []int) {
	return file_partner_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *EnrollUserInstrumentRequest) GetPartnerPaymentToken() string {
	if x != nil {
		return x.PartnerPaymentToken
	}
	return ""
}

func (x *EnrollUserInstrumentRequest) GetUser() *v1.UserDemographic {
	if x != nil {
		return x.User
	}
	return nil
}

type EnrollUserInstrumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status EnrollUserInstrumentResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=partner.user.v1.EnrollUserInstrumentResponse_Status" json:"status,omitempty"`
	Error  *v1.Error                           `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// Primary key on the external system.
	PartnerInstrumentRef string `protobuf:"bytes,3,opt,name=partner_instrument_ref,json=partnerInstrumentRef,proto3" json:"partner_instrument_ref,omitempty"`
}

func (x *EnrollUserInstrumentResponse) Reset() {
	*x = EnrollUserInstrumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_user_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollUserInstrumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollUserInstrumentResponse) ProtoMessage() {}

func (x *EnrollUserInstrumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_partner_user_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollUserInstrumentResponse.ProtoReflect.Descriptor instead.
func (*EnrollUserInstrumentResponse) Descriptor() ([]byte, []int) {
	return file_partner_user_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *EnrollUserInstrumentResponse) GetStatus() EnrollUserInstrumentResponse_Status {
	if x != nil {
		return x.Status
	}
	return EnrollUserInstrumentResponse_STATUS_UNSPECIFIED
}

func (x *EnrollUserInstrumentResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *EnrollUserInstrumentResponse) GetPartnerInstrumentRef() string {
	if x != nil {
		return x.PartnerInstrumentRef
	}
	return ""
}

type SearchUserInstrumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Primary key on the external system.
	PartnerInstrumentRef string `protobuf:"bytes,1,opt,name=partner_instrument_ref,json=partnerInstrumentRef,proto3" json:"partner_instrument_ref,omitempty"`
}

func (x *SearchUserInstrumentRequest) Reset() {
	*x = SearchUserInstrumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_user_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserInstrumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserInstrumentRequest) ProtoMessage() {}

func (x *SearchUserInstrumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_partner_user_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserInstrumentRequest.ProtoReflect.Descriptor instead.
func (*SearchUserInstrumentRequest) Descriptor() ([]byte, []int) {
	return file_partner_user_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *SearchUserInstrumentRequest) GetPartnerInstrumentRef() string {
	if x != nil {
		return x.PartnerInstrumentRef
	}
	return ""
}

type SearchUserInstrumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status SearchUserInstrumentResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=partner.user.v1.SearchUserInstrumentResponse_Status" json:"status,omitempty"`
	Error  *v1.Error                           `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Users  []*User                             `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *SearchUserInstrumentResponse) Reset() {
	*x = SearchUserInstrumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_user_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserInstrumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserInstrumentResponse) ProtoMessage() {}

func (x *SearchUserInstrumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_partner_user_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserInstrumentResponse.ProtoReflect.Descriptor instead.
func (*SearchUserInstrumentResponse) Descriptor() ([]byte, []int) {
	return file_partner_user_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *SearchUserInstrumentResponse) GetStatus() SearchUserInstrumentResponse_Status {
	if x != nil {
		return x.Status
	}
	return SearchUserInstrumentResponse_STATUS_UNSPECIFIED
}

func (x *SearchUserInstrumentResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SearchUserInstrumentResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_partner_user_v1_user_proto protoreflect.FileDescriptor

var file_partner_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12,
	0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65,
	0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x52, 0x0f, 0x75, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x81, 0x01, 0x0a,
	0x1b, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x15,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x2e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x87, 0x03, 0x0a, 0x1c, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x34, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65,
	0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x22, 0xba, 0x01,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x41,
	0x4c, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e,
	0x53, 0x55, 0x46, 0x46, 0x49, 0x43, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x55, 0x4e, 0x44, 0x53,
	0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x45,
	0x4c, 0x49, 0x47, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x53,
	0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x06, 0x22, 0x53, 0x0a, 0x1b, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x22,
	0x89, 0x02, 0x0a, 0x1c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x34, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x22, 0x46, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x42, 0x4e, 0x0a, 0x1d, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x65, 0x6e, 0x38, 0x30, 0x2e, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1a, 0x74, 0x65, 0x6e, 0x38, 0x30, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x54, 0x45, 0x4e, 0x38, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_partner_user_v1_user_proto_rawDescOnce sync.Once
	file_partner_user_v1_user_proto_rawDescData = file_partner_user_v1_user_proto_rawDesc
)

func file_partner_user_v1_user_proto_rawDescGZIP() []byte {
	file_partner_user_v1_user_proto_rawDescOnce.Do(func() {
		file_partner_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_partner_user_v1_user_proto_rawDescData)
	})
	return file_partner_user_v1_user_proto_rawDescData
}

var file_partner_user_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_partner_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_partner_user_v1_user_proto_goTypes = []interface{}{
	(EnrollUserInstrumentResponse_Status)(0), // 0: partner.user.v1.EnrollUserInstrumentResponse.Status
	(SearchUserInstrumentResponse_Status)(0), // 1: partner.user.v1.SearchUserInstrumentResponse.Status
	(*User)(nil),                             // 2: partner.user.v1.User
	(*EnrollUserInstrumentRequest)(nil),      // 3: partner.user.v1.EnrollUserInstrumentRequest
	(*EnrollUserInstrumentResponse)(nil),     // 4: partner.user.v1.EnrollUserInstrumentResponse
	(*SearchUserInstrumentRequest)(nil),      // 5: partner.user.v1.SearchUserInstrumentRequest
	(*SearchUserInstrumentResponse)(nil),     // 6: partner.user.v1.SearchUserInstrumentResponse
	(*timestamp.Timestamp)(nil),              // 7: google.protobuf.Timestamp
	(*v1.UserDemographic)(nil),               // 8: common.v1.UserDemographic
	(*v1.Error)(nil),                         // 9: common.v1.Error
}
var file_partner_user_v1_user_proto_depIdxs = []int32{
	7, // 0: partner.user.v1.User.created:type_name -> google.protobuf.Timestamp
	7, // 1: partner.user.v1.User.updated:type_name -> google.protobuf.Timestamp
	8, // 2: partner.user.v1.User.user_demographic:type_name -> common.v1.UserDemographic
	8, // 3: partner.user.v1.EnrollUserInstrumentRequest.user:type_name -> common.v1.UserDemographic
	0, // 4: partner.user.v1.EnrollUserInstrumentResponse.status:type_name -> partner.user.v1.EnrollUserInstrumentResponse.Status
	9, // 5: partner.user.v1.EnrollUserInstrumentResponse.error:type_name -> common.v1.Error
	1, // 6: partner.user.v1.SearchUserInstrumentResponse.status:type_name -> partner.user.v1.SearchUserInstrumentResponse.Status
	9, // 7: partner.user.v1.SearchUserInstrumentResponse.error:type_name -> common.v1.Error
	2, // 8: partner.user.v1.SearchUserInstrumentResponse.users:type_name -> partner.user.v1.User
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_partner_user_v1_user_proto_init() }
func file_partner_user_v1_user_proto_init() {
	if File_partner_user_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_partner_user_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_partner_user_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollUserInstrumentRequest); i {
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
		file_partner_user_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollUserInstrumentResponse); i {
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
		file_partner_user_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserInstrumentRequest); i {
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
		file_partner_user_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserInstrumentResponse); i {
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
			RawDescriptor: file_partner_user_v1_user_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_partner_user_v1_user_proto_goTypes,
		DependencyIndexes: file_partner_user_v1_user_proto_depIdxs,
		EnumInfos:         file_partner_user_v1_user_proto_enumTypes,
		MessageInfos:      file_partner_user_v1_user_proto_msgTypes,
	}.Build()
	File_partner_user_v1_user_proto = out.File
	file_partner_user_v1_user_proto_rawDesc = nil
	file_partner_user_v1_user_proto_goTypes = nil
	file_partner_user_v1_user_proto_depIdxs = nil
}
