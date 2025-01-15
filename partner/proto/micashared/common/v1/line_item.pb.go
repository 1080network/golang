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
// source: micashared/common/v1/line_item.proto

package commonv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	lineitemstatusv1 "github.com/1080network/golang/partner/proto/micashared/common/enums/lineitemstatusv1"
	unitv1 "github.com/1080network/golang/partner/proto/micashared/common/enums/unitv1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AdjustmentType int32

const (
	AdjustmentType_ADJUSTMENT_TYPE_UNSPECIFIED           AdjustmentType = 0
	AdjustmentType_ADJUSTMENT_TYPE_MANUFACTURER_DISCOUNT AdjustmentType = 1
	AdjustmentType_ADJUSTMENT_TYPE_STORE_DISCOUNT        AdjustmentType = 2
	AdjustmentType_ADJUSTMENT_TYPE_RETURN                AdjustmentType = 3
	AdjustmentType_ADJUSTMENT_TYPE_PAYMENT_CASH          AdjustmentType = 4
	AdjustmentType_ADJUSTMENT_TYPE_PAYMENT_GIFT_CARD     AdjustmentType = 5
	AdjustmentType_ADJUSTMENT_TYPE_PAYMENT_OTHER         AdjustmentType = 6
)

// Enum value maps for AdjustmentType.
var (
	AdjustmentType_name = map[int32]string{
		0: "ADJUSTMENT_TYPE_UNSPECIFIED",
		1: "ADJUSTMENT_TYPE_MANUFACTURER_DISCOUNT",
		2: "ADJUSTMENT_TYPE_STORE_DISCOUNT",
		3: "ADJUSTMENT_TYPE_RETURN",
		4: "ADJUSTMENT_TYPE_PAYMENT_CASH",
		5: "ADJUSTMENT_TYPE_PAYMENT_GIFT_CARD",
		6: "ADJUSTMENT_TYPE_PAYMENT_OTHER",
	}
	AdjustmentType_value = map[string]int32{
		"ADJUSTMENT_TYPE_UNSPECIFIED":           0,
		"ADJUSTMENT_TYPE_MANUFACTURER_DISCOUNT": 1,
		"ADJUSTMENT_TYPE_STORE_DISCOUNT":        2,
		"ADJUSTMENT_TYPE_RETURN":                3,
		"ADJUSTMENT_TYPE_PAYMENT_CASH":          4,
		"ADJUSTMENT_TYPE_PAYMENT_GIFT_CARD":     5,
		"ADJUSTMENT_TYPE_PAYMENT_OTHER":         6,
	}
)

func (x AdjustmentType) Enum() *AdjustmentType {
	p := new(AdjustmentType)
	*p = x
	return p
}

func (x AdjustmentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdjustmentType) Descriptor() protoreflect.EnumDescriptor {
	return file_micashared_common_v1_line_item_proto_enumTypes[0].Descriptor()
}

func (AdjustmentType) Type() protoreflect.EnumType {
	return &file_micashared_common_v1_line_item_proto_enumTypes[0]
}

func (x AdjustmentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdjustmentType.Descriptor instead.
func (AdjustmentType) EnumDescriptor() ([]byte, []int) {
	return file_micashared_common_v1_line_item_proto_rawDescGZIP(), []int{0}
}

type LineItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique number in a list of items that can be used to refer to a specific item
	Sequence int32 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// a header label that is shown on a receipt to group line items
	LineItemGroup string `protobuf:"bytes,2,opt,name=line_item_group,json=lineItemGroup,proto3" json:"line_item_group,omitempty"`
	ProductCode   string `protobuf:"bytes,3,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty"`
	Description   string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// an integer or decimal number
	Quantity      string      `protobuf:"bytes,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Unit          unitv1.Unit `protobuf:"varint,6,opt,name=unit,proto3,enum=micashared.common.enums.unit.v1.Unit" json:"unit,omitempty"`
	UnitAmount    string      `protobuf:"bytes,7,opt,name=unit_amount,json=unitAmount,proto3" json:"unit_amount,omitempty"`
	UnitTaxAmount string      `protobuf:"bytes,8,opt,name=unit_tax_amount,json=unitTaxAmount,proto3" json:"unit_tax_amount,omitempty"`
	// The total amount for this line item excluding tax.
	// Amount expressed as: ([0-9]*[.])?[0-9]+
	LineAmount string `protobuf:"bytes,9,opt,name=line_amount,json=lineAmount,proto3" json:"line_amount,omitempty"`
	// The amount of tax for this line item.
	// Amount expressed as: ([0-9]*[.])?[0-9]+
	LineTaxAmount string `protobuf:"bytes,10,opt,name=line_tax_amount,json=lineTaxAmount,proto3" json:"line_tax_amount,omitempty"`
}

func (x *LineItem) Reset() {
	*x = LineItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_line_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItem) ProtoMessage() {}

func (x *LineItem) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_line_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItem.ProtoReflect.Descriptor instead.
func (*LineItem) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_line_item_proto_rawDescGZIP(), []int{0}
}

func (x *LineItem) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *LineItem) GetLineItemGroup() string {
	if x != nil {
		return x.LineItemGroup
	}
	return ""
}

func (x *LineItem) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *LineItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LineItem) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *LineItem) GetUnit() unitv1.Unit {
	if x != nil {
		return x.Unit
	}
	return unitv1.Unit(0)
}

func (x *LineItem) GetUnitAmount() string {
	if x != nil {
		return x.UnitAmount
	}
	return ""
}

func (x *LineItem) GetUnitTaxAmount() string {
	if x != nil {
		return x.UnitTaxAmount
	}
	return ""
}

func (x *LineItem) GetLineAmount() string {
	if x != nil {
		return x.LineAmount
	}
	return ""
}

func (x *LineItem) GetLineTaxAmount() string {
	if x != nil {
		return x.LineTaxAmount
	}
	return ""
}

type LineItemWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LineItems []*LineItem `protobuf:"bytes,1,rep,name=line_items,json=lineItems,proto3" json:"line_items,omitempty"`
}

func (x *LineItemWrapper) Reset() {
	*x = LineItemWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_line_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItemWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItemWrapper) ProtoMessage() {}

func (x *LineItemWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_line_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItemWrapper.ProtoReflect.Descriptor instead.
func (*LineItemWrapper) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_line_item_proto_rawDescGZIP(), []int{1}
}

func (x *LineItemWrapper) GetLineItems() []*LineItem {
	if x != nil {
		return x.LineItems
	}
	return nil
}

type Adjustment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique number in a list of items that can be used to refer to a specific item
	Sequence       int32          `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	AdjustmentType AdjustmentType `protobuf:"varint,2,opt,name=adjustment_type,json=adjustmentType,proto3,enum=micashared.common.v1.AdjustmentType" json:"adjustment_type,omitempty"`
	ProductCode    string         `protobuf:"bytes,3,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty"`
	Description    string         `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// an integer or decimal number
	Quantity      string      `protobuf:"bytes,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Unit          unitv1.Unit `protobuf:"varint,6,opt,name=unit,proto3,enum=micashared.common.enums.unit.v1.Unit" json:"unit,omitempty"`
	UnitAmount    string      `protobuf:"bytes,7,opt,name=unit_amount,json=unitAmount,proto3" json:"unit_amount,omitempty"`
	UnitTaxAmount string      `protobuf:"bytes,8,opt,name=unit_tax_amount,json=unitTaxAmount,proto3" json:"unit_tax_amount,omitempty"`
	// The total amount for this line item excluding tax.
	// Amount expressed as: ([0-9]*[.])?[0-9]+
	LineAmount string `protobuf:"bytes,9,opt,name=line_amount,json=lineAmount,proto3" json:"line_amount,omitempty"`
	// The amount of tax for this line item.
	// Amount expressed as: ([0-9]*[.])?[0-9]+
	LineTaxAmount string `protobuf:"bytes,10,opt,name=line_tax_amount,json=lineTaxAmount,proto3" json:"line_tax_amount,omitempty"`
}

func (x *Adjustment) Reset() {
	*x = Adjustment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_line_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adjustment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adjustment) ProtoMessage() {}

func (x *Adjustment) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_line_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adjustment.ProtoReflect.Descriptor instead.
func (*Adjustment) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_line_item_proto_rawDescGZIP(), []int{2}
}

func (x *Adjustment) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Adjustment) GetAdjustmentType() AdjustmentType {
	if x != nil {
		return x.AdjustmentType
	}
	return AdjustmentType_ADJUSTMENT_TYPE_UNSPECIFIED
}

func (x *Adjustment) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *Adjustment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Adjustment) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *Adjustment) GetUnit() unitv1.Unit {
	if x != nil {
		return x.Unit
	}
	return unitv1.Unit(0)
}

func (x *Adjustment) GetUnitAmount() string {
	if x != nil {
		return x.UnitAmount
	}
	return ""
}

func (x *Adjustment) GetUnitTaxAmount() string {
	if x != nil {
		return x.UnitTaxAmount
	}
	return ""
}

func (x *Adjustment) GetLineAmount() string {
	if x != nil {
		return x.LineAmount
	}
	return ""
}

func (x *Adjustment) GetLineTaxAmount() string {
	if x != nil {
		return x.LineTaxAmount
	}
	return ""
}

type LineItemSequenceAndStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique number in a list of items that can be used to refer to a specific item
	Sequence int32                           `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Status   lineitemstatusv1.LineItemStatus `protobuf:"varint,2,opt,name=status,proto3,enum=micashared.common.enums.lineitemstatus.v1.LineItemStatus" json:"status,omitempty"`
}

func (x *LineItemSequenceAndStatus) Reset() {
	*x = LineItemSequenceAndStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_line_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItemSequenceAndStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItemSequenceAndStatus) ProtoMessage() {}

func (x *LineItemSequenceAndStatus) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_line_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItemSequenceAndStatus.ProtoReflect.Descriptor instead.
func (*LineItemSequenceAndStatus) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_line_item_proto_rawDescGZIP(), []int{3}
}

func (x *LineItemSequenceAndStatus) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *LineItemSequenceAndStatus) GetStatus() lineitemstatusv1.LineItemStatus {
	if x != nil {
		return x.Status
	}
	return lineitemstatusv1.LineItemStatus(0)
}

type LineItemSequenceAndStatusValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LineItemSequenceAndStatuses []*LineItemSequenceAndStatus `protobuf:"bytes,1,rep,name=line_item_sequence_and_statuses,json=lineItemSequenceAndStatuses,proto3" json:"line_item_sequence_and_statuses,omitempty"`
}

func (x *LineItemSequenceAndStatusValue) Reset() {
	*x = LineItemSequenceAndStatusValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_line_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItemSequenceAndStatusValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItemSequenceAndStatusValue) ProtoMessage() {}

func (x *LineItemSequenceAndStatusValue) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_line_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItemSequenceAndStatusValue.ProtoReflect.Descriptor instead.
func (*LineItemSequenceAndStatusValue) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_line_item_proto_rawDescGZIP(), []int{4}
}

func (x *LineItemSequenceAndStatusValue) GetLineItemSequenceAndStatuses() []*LineItemSequenceAndStatus {
	if x != nil {
		return x.LineItemSequenceAndStatuses
	}
	return nil
}

type LineItemAndStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LineItem *LineItem                       `protobuf:"bytes,1,opt,name=line_item,json=lineItem,proto3" json:"line_item,omitempty"`
	Status   lineitemstatusv1.LineItemStatus `protobuf:"varint,2,opt,name=status,proto3,enum=micashared.common.enums.lineitemstatus.v1.LineItemStatus" json:"status,omitempty"`
}

func (x *LineItemAndStatus) Reset() {
	*x = LineItemAndStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_line_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItemAndStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItemAndStatus) ProtoMessage() {}

func (x *LineItemAndStatus) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_line_item_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItemAndStatus.ProtoReflect.Descriptor instead.
func (*LineItemAndStatus) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_line_item_proto_rawDescGZIP(), []int{5}
}

func (x *LineItemAndStatus) GetLineItem() *LineItem {
	if x != nil {
		return x.LineItem
	}
	return nil
}

func (x *LineItemAndStatus) GetStatus() lineitemstatusv1.LineItemStatus {
	if x != nil {
		return x.Status
	}
	return lineitemstatusv1.LineItemStatus(0)
}

type LineItemResponseValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LineItemAndStatuses []*LineItemAndStatus `protobuf:"bytes,1,rep,name=line_item_and_statuses,json=lineItemAndStatuses,proto3" json:"line_item_and_statuses,omitempty"`
}

func (x *LineItemResponseValue) Reset() {
	*x = LineItemResponseValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_line_item_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItemResponseValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItemResponseValue) ProtoMessage() {}

func (x *LineItemResponseValue) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_line_item_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItemResponseValue.ProtoReflect.Descriptor instead.
func (*LineItemResponseValue) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_line_item_proto_rawDescGZIP(), []int{6}
}

func (x *LineItemResponseValue) GetLineItemAndStatuses() []*LineItemAndStatus {
	if x != nil {
		return x.LineItemAndStatuses
	}
	return nil
}

type Surcharge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Amount      string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Surcharge) Reset() {
	*x = Surcharge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_line_item_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Surcharge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Surcharge) ProtoMessage() {}

func (x *Surcharge) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_line_item_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Surcharge.ProtoReflect.Descriptor instead.
func (*Surcharge) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_line_item_proto_rawDescGZIP(), []int{7}
}

func (x *Surcharge) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Surcharge) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_micashared_common_v1_line_item_proto protoreflect.FileDescriptor

var file_micashared_common_v1_line_item_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6c,
	0x69, 0x6e, 0x65, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x02, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x75, 0x6e,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x74, 0x61, 0x78, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x6e, 0x69, 0x74,
	0x54, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6c, 0x69, 0x6e, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x74, 0x61, 0x78, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x50, 0x0a, 0x0f, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x57, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0xa5, 0x03, 0x0a, 0x0a, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x4d, 0x0a, 0x0f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e,
	0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x39, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x6e,
	0x69, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x75, 0x6e, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x75,
	0x6e, 0x69, 0x74, 0x5f, 0x74, 0x61, 0x78, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x6e, 0x69, 0x74, 0x54, 0x61, 0x78, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x61, 0x78,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c,
	0x69, 0x6e, 0x65, 0x54, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8a, 0x01, 0x0a,
	0x19, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x6c, 0x69, 0x6e, 0x65, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x1e, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x75, 0x0a, 0x1f,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x6e, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x1b, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x65, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x09, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x6c, 0x69,
	0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x51, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x6c, 0x69, 0x6e, 0x65, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x75, 0x0a, 0x15, 0x4c, 0x69, 0x6e,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x5c, 0x0a, 0x16, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x61, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x6c, 0x69, 0x6e,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73,
	0x22, 0x66, 0x0a, 0x09, 0x53, 0x75, 0x72, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x37, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1f, 0xfa, 0x42, 0x1c, 0x72, 0x1a, 0x10, 0x01, 0x18, 0x32, 0x32, 0x14, 0x5e, 0x28, 0x5b, 0x30,
	0x2d, 0x39, 0x5d, 0x2a, 0x5b, 0x2e, 0x5d, 0x29, 0x3f, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x2b, 0x24,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x88, 0x02, 0x0a, 0x0e, 0x41, 0x64, 0x6a,
	0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x41,
	0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x29, 0x0a, 0x25,
	0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4d, 0x41, 0x4e, 0x55, 0x46, 0x41, 0x43, 0x54, 0x55, 0x52, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x53,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x41, 0x44, 0x4a, 0x55, 0x53,
	0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45,
	0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x41,
	0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52,
	0x45, 0x54, 0x55, 0x52, 0x4e, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x44, 0x4a, 0x55, 0x53,
	0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x43, 0x41, 0x53, 0x48, 0x10, 0x04, 0x12, 0x25, 0x0a, 0x21, 0x41, 0x44, 0x4a,
	0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x59,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x47, 0x49, 0x46, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x05,
	0x12, 0x21, 0x0a, 0x1d, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4f, 0x54, 0x48, 0x45,
	0x52, 0x10, 0x06, 0x42, 0x4e, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0d,
	0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d,
	0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_v1_line_item_proto_rawDescOnce sync.Once
	file_micashared_common_v1_line_item_proto_rawDescData = file_micashared_common_v1_line_item_proto_rawDesc
)

func file_micashared_common_v1_line_item_proto_rawDescGZIP() []byte {
	file_micashared_common_v1_line_item_proto_rawDescOnce.Do(func() {
		file_micashared_common_v1_line_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_v1_line_item_proto_rawDescData)
	})
	return file_micashared_common_v1_line_item_proto_rawDescData
}

var file_micashared_common_v1_line_item_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_micashared_common_v1_line_item_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_micashared_common_v1_line_item_proto_goTypes = []interface{}{
	(AdjustmentType)(0),                    // 0: micashared.common.v1.AdjustmentType
	(*LineItem)(nil),                       // 1: micashared.common.v1.LineItem
	(*LineItemWrapper)(nil),                // 2: micashared.common.v1.LineItemWrapper
	(*Adjustment)(nil),                     // 3: micashared.common.v1.Adjustment
	(*LineItemSequenceAndStatus)(nil),      // 4: micashared.common.v1.LineItemSequenceAndStatus
	(*LineItemSequenceAndStatusValue)(nil), // 5: micashared.common.v1.LineItemSequenceAndStatusValue
	(*LineItemAndStatus)(nil),              // 6: micashared.common.v1.LineItemAndStatus
	(*LineItemResponseValue)(nil),          // 7: micashared.common.v1.LineItemResponseValue
	(*Surcharge)(nil),                      // 8: micashared.common.v1.Surcharge
	(unitv1.Unit)(0),                       // 9: micashared.common.enums.unit.v1.Unit
	(lineitemstatusv1.LineItemStatus)(0),   // 10: micashared.common.enums.lineitemstatus.v1.LineItemStatus
}
var file_micashared_common_v1_line_item_proto_depIdxs = []int32{
	9,  // 0: micashared.common.v1.LineItem.unit:type_name -> micashared.common.enums.unit.v1.Unit
	1,  // 1: micashared.common.v1.LineItemWrapper.line_items:type_name -> micashared.common.v1.LineItem
	0,  // 2: micashared.common.v1.Adjustment.adjustment_type:type_name -> micashared.common.v1.AdjustmentType
	9,  // 3: micashared.common.v1.Adjustment.unit:type_name -> micashared.common.enums.unit.v1.Unit
	10, // 4: micashared.common.v1.LineItemSequenceAndStatus.status:type_name -> micashared.common.enums.lineitemstatus.v1.LineItemStatus
	4,  // 5: micashared.common.v1.LineItemSequenceAndStatusValue.line_item_sequence_and_statuses:type_name -> micashared.common.v1.LineItemSequenceAndStatus
	1,  // 6: micashared.common.v1.LineItemAndStatus.line_item:type_name -> micashared.common.v1.LineItem
	10, // 7: micashared.common.v1.LineItemAndStatus.status:type_name -> micashared.common.enums.lineitemstatus.v1.LineItemStatus
	6,  // 8: micashared.common.v1.LineItemResponseValue.line_item_and_statuses:type_name -> micashared.common.v1.LineItemAndStatus
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_micashared_common_v1_line_item_proto_init() }
func file_micashared_common_v1_line_item_proto_init() {
	if File_micashared_common_v1_line_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_micashared_common_v1_line_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineItem); i {
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
		file_micashared_common_v1_line_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineItemWrapper); i {
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
		file_micashared_common_v1_line_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adjustment); i {
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
		file_micashared_common_v1_line_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineItemSequenceAndStatus); i {
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
		file_micashared_common_v1_line_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineItemSequenceAndStatusValue); i {
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
		file_micashared_common_v1_line_item_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineItemAndStatus); i {
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
		file_micashared_common_v1_line_item_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineItemResponseValue); i {
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
		file_micashared_common_v1_line_item_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Surcharge); i {
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
			RawDescriptor: file_micashared_common_v1_line_item_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_v1_line_item_proto_goTypes,
		DependencyIndexes: file_micashared_common_v1_line_item_proto_depIdxs,
		EnumInfos:         file_micashared_common_v1_line_item_proto_enumTypes,
		MessageInfos:      file_micashared_common_v1_line_item_proto_msgTypes,
	}.Build()
	File_micashared_common_v1_line_item_proto = out.File
	file_micashared_common_v1_line_item_proto_rawDesc = nil
	file_micashared_common_v1_line_item_proto_goTypes = nil
	file_micashared_common_v1_line_item_proto_depIdxs = nil
}
