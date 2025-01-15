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
// source: micashared/common/v1/uuek.proto

package commonv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	currencyv1 "github.com/1080network/golang/discount/proto/micashared/common/enums/currencyv1"
	uuektypev1 "github.com/1080network/golang/discount/proto/micashared/common/enums/uuektypev1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommonUUEK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UuekType uuektypev1.UUEKType `protobuf:"varint,1,opt,name=uuek_type,json=uuekType,proto3,enum=micashared.common.enums.uuektype.v1.UUEKType" json:"uuek_type,omitempty"`
	Uuek     string              `protobuf:"bytes,2,opt,name=uuek,proto3" json:"uuek,omitempty"`
	// version of the UUEK record, used for optimistic locking
	Version int64 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	// Date that the UUEK was created at Mica.
	Created *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	// Date that the UUEK was last updated at Mica.
	Updated   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"`
	ValidFrom *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=valid_from,json=validFrom,proto3" json:"valid_from,omitempty"`
	ValidTo   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=valid_to,json=validTo,proto3" json:"valid_to,omitempty"`
	// a maximum transaction that is allowed for this UUEK if the value request exceeds this
	// amount it will fail. If not set any transaction amount is allowed.
	SpendLimit string `protobuf:"bytes,10,opt,name=spend_limit,json=spendLimit,proto3" json:"spend_limit,omitempty"`
	// Types that are assignable to UseCriteria:
	//
	//	*CommonUUEK_NumberOfUses
	//	*CommonUUEK_Unlimited
	UseCriteria isCommonUUEK_UseCriteria `protobuf_oneof:"use_criteria"`
	// The 3-letter currency code defined in ISO 4217.
	Currency currencyv1.Currency `protobuf:"varint,8,opt,name=currency,proto3,enum=micashared.common.enums.currency.v1.Currency" json:"currency,omitempty"`
}

func (x *CommonUUEK) Reset() {
	*x = CommonUUEK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_uuek_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonUUEK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonUUEK) ProtoMessage() {}

func (x *CommonUUEK) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_uuek_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonUUEK.ProtoReflect.Descriptor instead.
func (*CommonUUEK) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_uuek_proto_rawDescGZIP(), []int{0}
}

func (x *CommonUUEK) GetUuekType() uuektypev1.UUEKType {
	if x != nil {
		return x.UuekType
	}
	return uuektypev1.UUEKType(0)
}

func (x *CommonUUEK) GetUuek() string {
	if x != nil {
		return x.Uuek
	}
	return ""
}

func (x *CommonUUEK) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CommonUUEK) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *CommonUUEK) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *CommonUUEK) GetValidFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidFrom
	}
	return nil
}

func (x *CommonUUEK) GetValidTo() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidTo
	}
	return nil
}

func (x *CommonUUEK) GetSpendLimit() string {
	if x != nil {
		return x.SpendLimit
	}
	return ""
}

func (m *CommonUUEK) GetUseCriteria() isCommonUUEK_UseCriteria {
	if m != nil {
		return m.UseCriteria
	}
	return nil
}

func (x *CommonUUEK) GetNumberOfUses() uint32 {
	if x, ok := x.GetUseCriteria().(*CommonUUEK_NumberOfUses); ok {
		return x.NumberOfUses
	}
	return 0
}

func (x *CommonUUEK) GetUnlimited() bool {
	if x, ok := x.GetUseCriteria().(*CommonUUEK_Unlimited); ok {
		return x.Unlimited
	}
	return false
}

func (x *CommonUUEK) GetCurrency() currencyv1.Currency {
	if x != nil {
		return x.Currency
	}
	return currencyv1.Currency(0)
}

type isCommonUUEK_UseCriteria interface {
	isCommonUUEK_UseCriteria()
}

type CommonUUEK_NumberOfUses struct {
	NumberOfUses uint32 `protobuf:"varint,11,opt,name=number_of_uses,json=numberOfUses,proto3,oneof"`
}

type CommonUUEK_Unlimited struct {
	Unlimited bool `protobuf:"varint,12,opt,name=unlimited,proto3,oneof"`
}

func (*CommonUUEK_NumberOfUses) isCommonUUEK_UseCriteria() {}

func (*CommonUUEK_Unlimited) isCommonUUEK_UseCriteria() {}

// This message holds immutable details about the instruments that back UUEKs. These attributes can be used to have
// stable references or knowledge of the underlying instrument that was used to transact.
type ImmutableInstrumentDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// non operable fix identifier guaranteed to be consistent across transactions
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *ImmutableInstrumentDetails) Reset() {
	*x = ImmutableInstrumentDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_uuek_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImmutableInstrumentDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImmutableInstrumentDetails) ProtoMessage() {}

func (x *ImmutableInstrumentDetails) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_uuek_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImmutableInstrumentDetails.ProtoReflect.Descriptor instead.
func (*ImmutableInstrumentDetails) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_uuek_proto_rawDescGZIP(), []int{1}
}

func (x *ImmutableInstrumentDetails) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

var File_micashared_common_v1_uuek_proto protoreflect.FileDescriptor

var file_micashared_common_v1_uuek_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x75, 0x65, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x75, 0x75, 0x65, 0x6b, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x75, 0x65, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x04, 0x0a, 0x0a, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x55, 0x55, 0x45, 0x4b, 0x12, 0x4a, 0x0a, 0x09, 0x75, 0x75, 0x65,
	0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x75, 0x75, 0x65, 0x6b, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x55, 0x45, 0x4b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x75, 0x75, 0x65,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x65, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x65, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x39, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x35, 0x0a, 0x08, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x54,
	0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x2f, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f,
	0x75, 0x73, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a,
	0x02, 0x20, 0x00, 0x48, 0x00, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x55,
	0x73, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x09, 0x75, 0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x75, 0x6e, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x65, 0x64, 0x12, 0x49, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x13,
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x03,
	0xf8, 0x42, 0x01, 0x22, 0x3c, 0x0a, 0x1a, 0x49, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x42, 0x4a, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x55, 0x55,
	0x45, 0x4b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_v1_uuek_proto_rawDescOnce sync.Once
	file_micashared_common_v1_uuek_proto_rawDescData = file_micashared_common_v1_uuek_proto_rawDesc
)

func file_micashared_common_v1_uuek_proto_rawDescGZIP() []byte {
	file_micashared_common_v1_uuek_proto_rawDescOnce.Do(func() {
		file_micashared_common_v1_uuek_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_v1_uuek_proto_rawDescData)
	})
	return file_micashared_common_v1_uuek_proto_rawDescData
}

var file_micashared_common_v1_uuek_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_micashared_common_v1_uuek_proto_goTypes = []interface{}{
	(*CommonUUEK)(nil),                 // 0: micashared.common.v1.CommonUUEK
	(*ImmutableInstrumentDetails)(nil), // 1: micashared.common.v1.ImmutableInstrumentDetails
	(uuektypev1.UUEKType)(0),           // 2: micashared.common.enums.uuektype.v1.UUEKType
	(*timestamppb.Timestamp)(nil),      // 3: google.protobuf.Timestamp
	(currencyv1.Currency)(0),           // 4: micashared.common.enums.currency.v1.Currency
}
var file_micashared_common_v1_uuek_proto_depIdxs = []int32{
	2, // 0: micashared.common.v1.CommonUUEK.uuek_type:type_name -> micashared.common.enums.uuektype.v1.UUEKType
	3, // 1: micashared.common.v1.CommonUUEK.created:type_name -> google.protobuf.Timestamp
	3, // 2: micashared.common.v1.CommonUUEK.updated:type_name -> google.protobuf.Timestamp
	3, // 3: micashared.common.v1.CommonUUEK.valid_from:type_name -> google.protobuf.Timestamp
	3, // 4: micashared.common.v1.CommonUUEK.valid_to:type_name -> google.protobuf.Timestamp
	4, // 5: micashared.common.v1.CommonUUEK.currency:type_name -> micashared.common.enums.currency.v1.Currency
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_micashared_common_v1_uuek_proto_init() }
func file_micashared_common_v1_uuek_proto_init() {
	if File_micashared_common_v1_uuek_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_micashared_common_v1_uuek_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonUUEK); i {
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
		file_micashared_common_v1_uuek_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImmutableInstrumentDetails); i {
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
	file_micashared_common_v1_uuek_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CommonUUEK_NumberOfUses)(nil),
		(*CommonUUEK_Unlimited)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_micashared_common_v1_uuek_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_v1_uuek_proto_goTypes,
		DependencyIndexes: file_micashared_common_v1_uuek_proto_depIdxs,
		MessageInfos:      file_micashared_common_v1_uuek_proto_msgTypes,
	}.Build()
	File_micashared_common_v1_uuek_proto = out.File
	file_micashared_common_v1_uuek_proto_rawDesc = nil
	file_micashared_common_v1_uuek_proto_goTypes = nil
	file_micashared_common_v1_uuek_proto_depIdxs = nil
}
