// Copyright (c) 2022 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: common/v1/line_item.proto

package commonv1

import (
	lineitemstatusv1 "github.com/1080network/golang/partner/proto/common/enums/lineitemstatusv1"
	unitv1 "github.com/1080network/golang/partner/proto/common/enums/unitv1"
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

type LineItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique number in a list of items that can be used to refer to a specific item
	Sequence    int32  `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	ProductCode string `protobuf:"bytes,2,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// an integer or decimal number
	Quantity      string      `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Unit          unitv1.Unit `protobuf:"varint,5,opt,name=unit,proto3,enum=common.enums.unit.v1.Unit" json:"unit,omitempty"`
	UnitAmount    string      `protobuf:"bytes,6,opt,name=unit_amount,json=unitAmount,proto3" json:"unit_amount,omitempty"`
	UnitTaxAmount string      `protobuf:"bytes,7,opt,name=unit_tax_amount,json=unitTaxAmount,proto3" json:"unit_tax_amount,omitempty"`
	// The total amount for this line item excluding tax.
	// Amount expressed as: [+-]?([0-9]*[.])?[0-9]+
	LineAmount string `protobuf:"bytes,8,opt,name=line_amount,json=lineAmount,proto3" json:"line_amount,omitempty"`
	// The amount of tax for this line item.
	// Amount expressed as: [+-]?([0-9]*[.])?[0-9]+
	LineTaxAmount string `protobuf:"bytes,9,opt,name=line_tax_amount,json=lineTaxAmount,proto3" json:"line_tax_amount,omitempty"`
}

func (x *LineItem) Reset() {
	*x = LineItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_line_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItem) ProtoMessage() {}

func (x *LineItem) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_line_item_proto_msgTypes[0]
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
	return file_common_v1_line_item_proto_rawDescGZIP(), []int{0}
}

func (x *LineItem) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
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

type LineItemSequenceAndStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique number in a list of items that can be used to refer to a specific item
	Sequence int32                           `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Status   lineitemstatusv1.LineItemStatus `protobuf:"varint,2,opt,name=status,proto3,enum=common.enums.lineitemstatus.v1.LineItemStatus" json:"status,omitempty"`
}

func (x *LineItemSequenceAndStatus) Reset() {
	*x = LineItemSequenceAndStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_line_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItemSequenceAndStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItemSequenceAndStatus) ProtoMessage() {}

func (x *LineItemSequenceAndStatus) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_line_item_proto_msgTypes[1]
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
	return file_common_v1_line_item_proto_rawDescGZIP(), []int{1}
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
		mi := &file_common_v1_line_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItemSequenceAndStatusValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItemSequenceAndStatusValue) ProtoMessage() {}

func (x *LineItemSequenceAndStatusValue) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_line_item_proto_msgTypes[2]
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
	return file_common_v1_line_item_proto_rawDescGZIP(), []int{2}
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
	Status   lineitemstatusv1.LineItemStatus `protobuf:"varint,2,opt,name=status,proto3,enum=common.enums.lineitemstatus.v1.LineItemStatus" json:"status,omitempty"`
}

func (x *LineItemAndStatus) Reset() {
	*x = LineItemAndStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_line_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItemAndStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItemAndStatus) ProtoMessage() {}

func (x *LineItemAndStatus) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_line_item_proto_msgTypes[3]
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
	return file_common_v1_line_item_proto_rawDescGZIP(), []int{3}
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
		mi := &file_common_v1_line_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItemResponseValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItemResponseValue) ProtoMessage() {}

func (x *LineItemResponseValue) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_line_item_proto_msgTypes[4]
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
	return file_common_v1_line_item_proto_rawDescGZIP(), []int{4}
}

func (x *LineItemResponseValue) GetLineItemAndStatuses() []*LineItemAndStatus {
	if x != nil {
		return x.LineItemAndStatuses
	}
	return nil
}

var File_common_v1_line_item_proto protoreflect.FileDescriptor

var file_common_v1_line_item_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x75, 0x6e, 0x69, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9,
	0x02, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e,
	0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x6e, 0x69, 0x74,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75,
	0x6e, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x6e, 0x69,
	0x74, 0x5f, 0x74, 0x61, 0x78, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x75, 0x6e, 0x69, 0x74, 0x54, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x61, 0x78, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x6e,
	0x65, 0x54, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7f, 0x0a, 0x19, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x65, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x1e,
	0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x6a,
	0x0a, 0x1f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x1b, 0x6c,
	0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x41,
	0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x4c,
	0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x30, 0x0a, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x65, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x6a, 0x0a, 0x15, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x51, 0x0a, 0x16, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x13, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x6e, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x42, 0x48, 0x0a, 0x11, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4c, 0x69, 0x6e,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x6d, 0x69, 0x63, 0x61,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_v1_line_item_proto_rawDescOnce sync.Once
	file_common_v1_line_item_proto_rawDescData = file_common_v1_line_item_proto_rawDesc
)

func file_common_v1_line_item_proto_rawDescGZIP() []byte {
	file_common_v1_line_item_proto_rawDescOnce.Do(func() {
		file_common_v1_line_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v1_line_item_proto_rawDescData)
	})
	return file_common_v1_line_item_proto_rawDescData
}

var file_common_v1_line_item_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_common_v1_line_item_proto_goTypes = []interface{}{
	(*LineItem)(nil),                       // 0: common.v1.LineItem
	(*LineItemSequenceAndStatus)(nil),      // 1: common.v1.LineItemSequenceAndStatus
	(*LineItemSequenceAndStatusValue)(nil), // 2: common.v1.LineItemSequenceAndStatusValue
	(*LineItemAndStatus)(nil),              // 3: common.v1.LineItemAndStatus
	(*LineItemResponseValue)(nil),          // 4: common.v1.LineItemResponseValue
	(unitv1.Unit)(0),                       // 5: common.enums.unit.v1.Unit
	(lineitemstatusv1.LineItemStatus)(0),   // 6: common.enums.lineitemstatus.v1.LineItemStatus
}
var file_common_v1_line_item_proto_depIdxs = []int32{
	5, // 0: common.v1.LineItem.unit:type_name -> common.enums.unit.v1.Unit
	6, // 1: common.v1.LineItemSequenceAndStatus.status:type_name -> common.enums.lineitemstatus.v1.LineItemStatus
	1, // 2: common.v1.LineItemSequenceAndStatusValue.line_item_sequence_and_statuses:type_name -> common.v1.LineItemSequenceAndStatus
	0, // 3: common.v1.LineItemAndStatus.line_item:type_name -> common.v1.LineItem
	6, // 4: common.v1.LineItemAndStatus.status:type_name -> common.enums.lineitemstatus.v1.LineItemStatus
	3, // 5: common.v1.LineItemResponseValue.line_item_and_statuses:type_name -> common.v1.LineItemAndStatus
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_common_v1_line_item_proto_init() }
func file_common_v1_line_item_proto_init() {
	if File_common_v1_line_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_v1_line_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_common_v1_line_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_common_v1_line_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_common_v1_line_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_common_v1_line_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_v1_line_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_line_item_proto_goTypes,
		DependencyIndexes: file_common_v1_line_item_proto_depIdxs,
		MessageInfos:      file_common_v1_line_item_proto_msgTypes,
	}.Build()
	File_common_v1_line_item_proto = out.File
	file_common_v1_line_item_proto_rawDesc = nil
	file_common_v1_line_item_proto_goTypes = nil
	file_common_v1_line_item_proto_depIdxs = nil
}
