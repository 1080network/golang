// Copyright (c) 2022 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.12
// source: micashared/common/v1/receipt.proto

package commonv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	barcodelocationv1 "github.com/1080network/golang/discount/proto/micashared/common/enums/barcodelocationv1"
	barcodetypev1 "github.com/1080network/golang/discount/proto/micashared/common/enums/barcodetypev1"
	currencyv1 "github.com/1080network/golang/discount/proto/micashared/common/enums/currencyv1"
	organizationcategoryv1 "github.com/1080network/golang/discount/proto/micashared/common/enums/organizationcategoryv1"
	valueoperationtypev1 "github.com/1080network/golang/discount/proto/micashared/common/enums/valueoperationtypev1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetReceiptResponse_Status int32

const (
	GetReceiptResponse_STATUS_UNSPECIFIED GetReceiptResponse_Status = 0
	GetReceiptResponse_STATUS_SUCCESS     GetReceiptResponse_Status = 1
	GetReceiptResponse_STATUS_ERROR       GetReceiptResponse_Status = 2
	GetReceiptResponse_STATUS_NOT_FOUND   GetReceiptResponse_Status = 3
)

// Enum value maps for GetReceiptResponse_Status.
var (
	GetReceiptResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	GetReceiptResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x GetReceiptResponse_Status) Enum() *GetReceiptResponse_Status {
	p := new(GetReceiptResponse_Status)
	*p = x
	return p
}

func (x GetReceiptResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetReceiptResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_micashared_common_v1_receipt_proto_enumTypes[0].Descriptor()
}

func (GetReceiptResponse_Status) Type() protoreflect.EnumType {
	return &file_micashared_common_v1_receipt_proto_enumTypes[0]
}

func (x GetReceiptResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetReceiptResponse_Status.Descriptor instead.
func (GetReceiptResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_micashared_common_v1_receipt_proto_rawDescGZIP(), []int{2, 0}
}

type Receipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mica's transaction record primary key.
	TransactionKey string `protobuf:"bytes,1,opt,name=transaction_key,json=transactionKey,proto3" json:"transaction_key,omitempty"`
	// version of the user record, used for optimistic locking
	Version int64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// Date that the Store was created at Mica.
	Created *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	// Date that the Store was last updated at Mica.
	Updated       *timestamppb.Timestamp                  `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	OperationType valueoperationtypev1.ValueOperationType `protobuf:"varint,5,opt,name=operation_type,json=operationType,proto3,enum=micashared.common.enums.valueoperationtype.v1.ValueOperationType" json:"operation_type,omitempty"`
	// A unique identifier at the Partner for this transaction.
	PartnerTransactionRef string `protobuf:"bytes,6,opt,name=partner_transaction_ref,json=partnerTransactionRef,proto3" json:"partner_transaction_ref,omitempty"`
	// The Mica generated key for this instrument.
	ServiceProviderInstrumentKey string `protobuf:"bytes,7,opt,name=service_provider_instrument_key,json=serviceProviderInstrumentKey,proto3" json:"service_provider_instrument_key,omitempty"`
	// Service Provider's primary key for their instrument.
	ServiceProviderInstrumentRef string `protobuf:"bytes,8,opt,name=service_provider_instrument_ref,json=serviceProviderInstrumentRef,proto3" json:"service_provider_instrument_ref,omitempty"`
	// The 3-letter currency code defined in ISO 4217. Note all amounts are in this currency.
	Currency currencyv1.Currency `protobuf:"varint,10,opt,name=currency,proto3,enum=micashared.common.enums.currency.v1.Currency" json:"currency,omitempty"`
	// The organization this operation is being processed for.
	OrganizationKey string `protobuf:"bytes,11,opt,name=organization_key,json=organizationKey,proto3" json:"organization_key,omitempty"`
	// The organization name this operation is being processed for.
	OrganizationName string `protobuf:"bytes,12,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty"`
	// The organization address this operation is being processed for.
	OrganizationAddress *Address `protobuf:"bytes,13,opt,name=organization_address,json=organizationAddress,proto3" json:"organization_address,omitempty"`
	// The organizations category.
	Category organizationcategoryv1.OrganizationCategory `protobuf:"varint,14,opt,name=category,proto3,enum=micashared.common.enums.organizationcategory.v1.OrganizationCategory" json:"category,omitempty"`
	// The store where this operation occurred.
	StoreKey string `protobuf:"bytes,15,opt,name=store_key,json=storeKey,proto3" json:"store_key,omitempty"`
	// The number of this store (an alternative unique value).
	StoreNumber string `protobuf:"bytes,16,opt,name=store_number,json=storeNumber,proto3" json:"store_number,omitempty"`
	// The street address of this store.
	StoreAddress *Address `protobuf:"bytes,17,opt,name=store_address,json=storeAddress,proto3" json:"store_address,omitempty"`
	// The clerk that processed this operation.
	ClerkIdentifier string `protobuf:"bytes,18,opt,name=clerk_identifier,json=clerkIdentifier,proto3" json:"clerk_identifier,omitempty"`
	// Total amount for this transaction.
	//
	// When Items are present then: `total_amount` = sum(`items.amount + items.tax_amount`)
	// Amount expressed as: [+-]?([0-9]*[.])?[0-9]+
	TotalAmount string `protobuf:"bytes,19,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	// The amount requested for this transaction. Note that this will be the same as total_amount if all products are eligible
	// Amount expressed as: [+-]?([0-9]*[.])?[0-9]+
	RequestedAmount string `protobuf:"bytes,20,opt,name=requested_amount,json=requestedAmount,proto3" json:"requested_amount,omitempty"`
	// Amount expressed as: [+-]?([0-9]*[.])?[0-9]+
	ApprovedAmount string `protobuf:"bytes,21,opt,name=approved_amount,json=approvedAmount,proto3" json:"approved_amount,omitempty"`
	OrderNumber    string `protobuf:"bytes,22,opt,name=order_number,json=orderNumber,proto3" json:"order_number,omitempty"`
	// List of items bought or returned.
	LineItemAndStatuses []*LineItemAndStatus              `protobuf:"bytes,23,rep,name=line_item_and_statuses,json=lineItemAndStatuses,proto3" json:"line_item_and_statuses,omitempty"`
	AppliedDiscounts    []*AppliedDiscount                `protobuf:"bytes,24,rep,name=applied_discounts,json=appliedDiscounts,proto3" json:"applied_discounts,omitempty"`
	Adjustments         []*Adjustment                     `protobuf:"bytes,25,rep,name=adjustments,proto3" json:"adjustments,omitempty"`
	BarcodeType         barcodetypev1.BarcodeType         `protobuf:"varint,26,opt,name=barcode_type,json=barcodeType,proto3,enum=micashared.common.enums.barcodetype.v1.BarcodeType" json:"barcode_type,omitempty"`
	BarcodeLocation     barcodelocationv1.BarcodeLocation `protobuf:"varint,27,opt,name=barcode_location,json=barcodeLocation,proto3,enum=micashared.common.enums.barcodelocation.v1.BarcodeLocation" json:"barcode_location,omitempty"`
	BarcodeText         string                            `protobuf:"bytes,28,opt,name=barcode_text,json=barcodeText,proto3" json:"barcode_text,omitempty"`
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_receipt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_receipt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_receipt_proto_rawDescGZIP(), []int{0}
}

func (x *Receipt) GetTransactionKey() string {
	if x != nil {
		return x.TransactionKey
	}
	return ""
}

func (x *Receipt) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Receipt) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Receipt) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *Receipt) GetOperationType() valueoperationtypev1.ValueOperationType {
	if x != nil {
		return x.OperationType
	}
	return valueoperationtypev1.ValueOperationType(0)
}

func (x *Receipt) GetPartnerTransactionRef() string {
	if x != nil {
		return x.PartnerTransactionRef
	}
	return ""
}

func (x *Receipt) GetServiceProviderInstrumentKey() string {
	if x != nil {
		return x.ServiceProviderInstrumentKey
	}
	return ""
}

func (x *Receipt) GetServiceProviderInstrumentRef() string {
	if x != nil {
		return x.ServiceProviderInstrumentRef
	}
	return ""
}

func (x *Receipt) GetCurrency() currencyv1.Currency {
	if x != nil {
		return x.Currency
	}
	return currencyv1.Currency(0)
}

func (x *Receipt) GetOrganizationKey() string {
	if x != nil {
		return x.OrganizationKey
	}
	return ""
}

func (x *Receipt) GetOrganizationName() string {
	if x != nil {
		return x.OrganizationName
	}
	return ""
}

func (x *Receipt) GetOrganizationAddress() *Address {
	if x != nil {
		return x.OrganizationAddress
	}
	return nil
}

func (x *Receipt) GetCategory() organizationcategoryv1.OrganizationCategory {
	if x != nil {
		return x.Category
	}
	return organizationcategoryv1.OrganizationCategory(0)
}

func (x *Receipt) GetStoreKey() string {
	if x != nil {
		return x.StoreKey
	}
	return ""
}

func (x *Receipt) GetStoreNumber() string {
	if x != nil {
		return x.StoreNumber
	}
	return ""
}

func (x *Receipt) GetStoreAddress() *Address {
	if x != nil {
		return x.StoreAddress
	}
	return nil
}

func (x *Receipt) GetClerkIdentifier() string {
	if x != nil {
		return x.ClerkIdentifier
	}
	return ""
}

func (x *Receipt) GetTotalAmount() string {
	if x != nil {
		return x.TotalAmount
	}
	return ""
}

func (x *Receipt) GetRequestedAmount() string {
	if x != nil {
		return x.RequestedAmount
	}
	return ""
}

func (x *Receipt) GetApprovedAmount() string {
	if x != nil {
		return x.ApprovedAmount
	}
	return ""
}

func (x *Receipt) GetOrderNumber() string {
	if x != nil {
		return x.OrderNumber
	}
	return ""
}

func (x *Receipt) GetLineItemAndStatuses() []*LineItemAndStatus {
	if x != nil {
		return x.LineItemAndStatuses
	}
	return nil
}

func (x *Receipt) GetAppliedDiscounts() []*AppliedDiscount {
	if x != nil {
		return x.AppliedDiscounts
	}
	return nil
}

func (x *Receipt) GetAdjustments() []*Adjustment {
	if x != nil {
		return x.Adjustments
	}
	return nil
}

func (x *Receipt) GetBarcodeType() barcodetypev1.BarcodeType {
	if x != nil {
		return x.BarcodeType
	}
	return barcodetypev1.BarcodeType(0)
}

func (x *Receipt) GetBarcodeLocation() barcodelocationv1.BarcodeLocation {
	if x != nil {
		return x.BarcodeLocation
	}
	return barcodelocationv1.BarcodeLocation(0)
}

func (x *Receipt) GetBarcodeText() string {
	if x != nil {
		return x.BarcodeText
	}
	return ""
}

type GetReceiptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionKey string `protobuf:"bytes,1,opt,name=transaction_key,json=transactionKey,proto3" json:"transaction_key,omitempty"`
}

func (x *GetReceiptRequest) Reset() {
	*x = GetReceiptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_receipt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptRequest) ProtoMessage() {}

func (x *GetReceiptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_receipt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptRequest.ProtoReflect.Descriptor instead.
func (*GetReceiptRequest) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_receipt_proto_rawDescGZIP(), []int{1}
}

func (x *GetReceiptRequest) GetTransactionKey() string {
	if x != nil {
		return x.TransactionKey
	}
	return ""
}

type GetReceiptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  GetReceiptResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=micashared.common.v1.GetReceiptResponse_Status" json:"status,omitempty"`
	Error   *Error                    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Receipt *Receipt                  `protobuf:"bytes,3,opt,name=receipt,proto3" json:"receipt,omitempty"`
}

func (x *GetReceiptResponse) Reset() {
	*x = GetReceiptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_receipt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptResponse) ProtoMessage() {}

func (x *GetReceiptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_receipt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptResponse.ProtoReflect.Descriptor instead.
func (*GetReceiptResponse) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_receipt_proto_rawDescGZIP(), []int{2}
}

func (x *GetReceiptResponse) GetStatus() GetReceiptResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetReceiptResponse_STATUS_UNSPECIFIED
}

func (x *GetReceiptResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetReceiptResponse) GetReceipt() *Receipt {
	if x != nil {
		return x.Receipt
	}
	return nil
}

var File_micashared_common_v1_receipt_proto protoreflect.FileDescriptor

var file_micashared_common_v1_receipt_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4b, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x48, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e,
	0x0d, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x32, 0x0a, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x68, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36,
	0x0a, 0x17, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x15, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x12, 0x50, 0x0a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52, 0x1c, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x45, 0x0a, 0x1f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x1c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x12,
	0x49, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2d, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x34, 0x0a, 0x10, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52,
	0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79,
	0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a,
	0x14, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x13, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x61, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x45, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x26, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32,
	0x52, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x42, 0x0a,
	0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6c, 0x65, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6c, 0x65,
	0x72, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x29, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x5c, 0x0a, 0x16, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73,
	0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x13, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x65, 0x73, 0x12, 0x52, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x42, 0x0a, 0x0b, 0x61, 0x64, 0x6a, 0x75,
	0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0b, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x56, 0x0a, 0x0c,
	0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1a, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x33, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x62, 0x61, 0x72,
	0x63, 0x6f, 0x64, 0x65, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x72, 0x63,
	0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x66, 0x0a, 0x10, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x72, 0x63,
	0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x62, 0x61, 0x72,
	0x63, 0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x1c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x78, 0x74, 0x22,
	0x47, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x22, 0xa7, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x07, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x07, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0x03, 0x42, 0x4d, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43,
	0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_v1_receipt_proto_rawDescOnce sync.Once
	file_micashared_common_v1_receipt_proto_rawDescData = file_micashared_common_v1_receipt_proto_rawDesc
)

func file_micashared_common_v1_receipt_proto_rawDescGZIP() []byte {
	file_micashared_common_v1_receipt_proto_rawDescOnce.Do(func() {
		file_micashared_common_v1_receipt_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_v1_receipt_proto_rawDescData)
	})
	return file_micashared_common_v1_receipt_proto_rawDescData
}

var file_micashared_common_v1_receipt_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_micashared_common_v1_receipt_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_micashared_common_v1_receipt_proto_goTypes = []interface{}{
	(GetReceiptResponse_Status)(0),                   // 0: micashared.common.v1.GetReceiptResponse.Status
	(*Receipt)(nil),                                  // 1: micashared.common.v1.Receipt
	(*GetReceiptRequest)(nil),                        // 2: micashared.common.v1.GetReceiptRequest
	(*GetReceiptResponse)(nil),                       // 3: micashared.common.v1.GetReceiptResponse
	(*timestamppb.Timestamp)(nil),                    // 4: google.protobuf.Timestamp
	(valueoperationtypev1.ValueOperationType)(0),     // 5: micashared.common.enums.valueoperationtype.v1.ValueOperationType
	(currencyv1.Currency)(0),                         // 6: micashared.common.enums.currency.v1.Currency
	(*Address)(nil),                                  // 7: micashared.common.v1.Address
	(organizationcategoryv1.OrganizationCategory)(0), // 8: micashared.common.enums.organizationcategory.v1.OrganizationCategory
	(*LineItemAndStatus)(nil),                        // 9: micashared.common.v1.LineItemAndStatus
	(*AppliedDiscount)(nil),                          // 10: micashared.common.v1.AppliedDiscount
	(*Adjustment)(nil),                               // 11: micashared.common.v1.Adjustment
	(barcodetypev1.BarcodeType)(0),                   // 12: micashared.common.enums.barcodetype.v1.BarcodeType
	(barcodelocationv1.BarcodeLocation)(0),           // 13: micashared.common.enums.barcodelocation.v1.BarcodeLocation
	(*Error)(nil),                                    // 14: micashared.common.v1.Error
}
var file_micashared_common_v1_receipt_proto_depIdxs = []int32{
	4,  // 0: micashared.common.v1.Receipt.created:type_name -> google.protobuf.Timestamp
	4,  // 1: micashared.common.v1.Receipt.updated:type_name -> google.protobuf.Timestamp
	5,  // 2: micashared.common.v1.Receipt.operation_type:type_name -> micashared.common.enums.valueoperationtype.v1.ValueOperationType
	6,  // 3: micashared.common.v1.Receipt.currency:type_name -> micashared.common.enums.currency.v1.Currency
	7,  // 4: micashared.common.v1.Receipt.organization_address:type_name -> micashared.common.v1.Address
	8,  // 5: micashared.common.v1.Receipt.category:type_name -> micashared.common.enums.organizationcategory.v1.OrganizationCategory
	7,  // 6: micashared.common.v1.Receipt.store_address:type_name -> micashared.common.v1.Address
	9,  // 7: micashared.common.v1.Receipt.line_item_and_statuses:type_name -> micashared.common.v1.LineItemAndStatus
	10, // 8: micashared.common.v1.Receipt.applied_discounts:type_name -> micashared.common.v1.AppliedDiscount
	11, // 9: micashared.common.v1.Receipt.adjustments:type_name -> micashared.common.v1.Adjustment
	12, // 10: micashared.common.v1.Receipt.barcode_type:type_name -> micashared.common.enums.barcodetype.v1.BarcodeType
	13, // 11: micashared.common.v1.Receipt.barcode_location:type_name -> micashared.common.enums.barcodelocation.v1.BarcodeLocation
	0,  // 12: micashared.common.v1.GetReceiptResponse.status:type_name -> micashared.common.v1.GetReceiptResponse.Status
	14, // 13: micashared.common.v1.GetReceiptResponse.error:type_name -> micashared.common.v1.Error
	1,  // 14: micashared.common.v1.GetReceiptResponse.receipt:type_name -> micashared.common.v1.Receipt
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_micashared_common_v1_receipt_proto_init() }
func file_micashared_common_v1_receipt_proto_init() {
	if File_micashared_common_v1_receipt_proto != nil {
		return
	}
	file_micashared_common_v1_address_proto_init()
	file_micashared_common_v1_discount_proto_init()
	file_micashared_common_v1_error_proto_init()
	file_micashared_common_v1_line_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_micashared_common_v1_receipt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Receipt); i {
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
		file_micashared_common_v1_receipt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptRequest); i {
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
		file_micashared_common_v1_receipt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptResponse); i {
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
			RawDescriptor: file_micashared_common_v1_receipt_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_v1_receipt_proto_goTypes,
		DependencyIndexes: file_micashared_common_v1_receipt_proto_depIdxs,
		EnumInfos:         file_micashared_common_v1_receipt_proto_enumTypes,
		MessageInfos:      file_micashared_common_v1_receipt_proto_msgTypes,
	}.Build()
	File_micashared_common_v1_receipt_proto = out.File
	file_micashared_common_v1_receipt_proto_rawDesc = nil
	file_micashared_common_v1_receipt_proto_goTypes = nil
	file_micashared_common_v1_receipt_proto_depIdxs = nil
}
