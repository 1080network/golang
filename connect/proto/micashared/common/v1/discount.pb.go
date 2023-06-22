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
// source: micashared/common/v1/discount.proto

package commonv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	discounttypev1 "github.com/1080network/golang/connect/proto/micashared/common/enums/discounttypev1"
	unitv1 "github.com/1080network/golang/connect/proto/micashared/common/enums/unitv1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DiscountLineItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sequence int32 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// quantity of this line item that was applicable to the given discount
	Quantity string `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// unit of this line item that was applicable to the given discount
	Unit unitv1.Unit `protobuf:"varint,3,opt,name=unit,proto3,enum=micashared.common.enums.unit.v1.Unit" json:"unit,omitempty"`
}

func (x *DiscountLineItem) Reset() {
	*x = DiscountLineItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_discount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountLineItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountLineItem) ProtoMessage() {}

func (x *DiscountLineItem) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_discount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountLineItem.ProtoReflect.Descriptor instead.
func (*DiscountLineItem) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_discount_proto_rawDescGZIP(), []int{0}
}

func (x *DiscountLineItem) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *DiscountLineItem) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *DiscountLineItem) GetUnit() unitv1.Unit {
	if x != nil {
		return x.Unit
	}
	return unitv1.Unit(0)
}

type AppliedDiscount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscountKey        string                      `protobuf:"bytes,1,opt,name=discount_key,json=discountKey,proto3" json:"discount_key,omitempty"`
	ReceiptDescription string                      `protobuf:"bytes,2,opt,name=receipt_description,json=receiptDescription,proto3" json:"receipt_description,omitempty"`
	DiscountType       discounttypev1.DiscountType `protobuf:"varint,3,opt,name=discount_type,json=discountType,proto3,enum=micashared.common.enums.discounttype.v1.DiscountType" json:"discount_type,omitempty"`
	DiscountAmount     string                      `protobuf:"bytes,4,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount,omitempty"`
	DiscountTaxAmount  string                      `protobuf:"bytes,5,opt,name=discount_tax_amount,json=discountTaxAmount,proto3" json:"discount_tax_amount,omitempty"`
	DiscountLineItems  []*DiscountLineItem         `protobuf:"bytes,6,rep,name=discount_line_items,json=discountLineItems,proto3" json:"discount_line_items,omitempty"`
}

func (x *AppliedDiscount) Reset() {
	*x = AppliedDiscount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_discount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppliedDiscount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppliedDiscount) ProtoMessage() {}

func (x *AppliedDiscount) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_discount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppliedDiscount.ProtoReflect.Descriptor instead.
func (*AppliedDiscount) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_discount_proto_rawDescGZIP(), []int{1}
}

func (x *AppliedDiscount) GetDiscountKey() string {
	if x != nil {
		return x.DiscountKey
	}
	return ""
}

func (x *AppliedDiscount) GetReceiptDescription() string {
	if x != nil {
		return x.ReceiptDescription
	}
	return ""
}

func (x *AppliedDiscount) GetDiscountType() discounttypev1.DiscountType {
	if x != nil {
		return x.DiscountType
	}
	return discounttypev1.DiscountType(0)
}

func (x *AppliedDiscount) GetDiscountAmount() string {
	if x != nil {
		return x.DiscountAmount
	}
	return ""
}

func (x *AppliedDiscount) GetDiscountTaxAmount() string {
	if x != nil {
		return x.DiscountTaxAmount
	}
	return ""
}

func (x *AppliedDiscount) GetDiscountLineItems() []*DiscountLineItem {
	if x != nil {
		return x.DiscountLineItems
	}
	return nil
}

var File_micashared_common_v1_discount_proto protoreflect.FileDescriptor

var file_micashared_common_v1_discount_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x3b, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01,
	0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0xfd, 0x02, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x64, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5a, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x35, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a,
	0x13, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x78, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x56, 0x0a,
	0x13, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x11, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x4e, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x42, 0x0d, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a,
	0x1d, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02,
	0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_v1_discount_proto_rawDescOnce sync.Once
	file_micashared_common_v1_discount_proto_rawDescData = file_micashared_common_v1_discount_proto_rawDesc
)

func file_micashared_common_v1_discount_proto_rawDescGZIP() []byte {
	file_micashared_common_v1_discount_proto_rawDescOnce.Do(func() {
		file_micashared_common_v1_discount_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_v1_discount_proto_rawDescData)
	})
	return file_micashared_common_v1_discount_proto_rawDescData
}

var file_micashared_common_v1_discount_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_micashared_common_v1_discount_proto_goTypes = []interface{}{
	(*DiscountLineItem)(nil),         // 0: micashared.common.v1.DiscountLineItem
	(*AppliedDiscount)(nil),          // 1: micashared.common.v1.AppliedDiscount
	(unitv1.Unit)(0),                 // 2: micashared.common.enums.unit.v1.Unit
	(discounttypev1.DiscountType)(0), // 3: micashared.common.enums.discounttype.v1.DiscountType
}
var file_micashared_common_v1_discount_proto_depIdxs = []int32{
	2, // 0: micashared.common.v1.DiscountLineItem.unit:type_name -> micashared.common.enums.unit.v1.Unit
	3, // 1: micashared.common.v1.AppliedDiscount.discount_type:type_name -> micashared.common.enums.discounttype.v1.DiscountType
	0, // 2: micashared.common.v1.AppliedDiscount.discount_line_items:type_name -> micashared.common.v1.DiscountLineItem
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_micashared_common_v1_discount_proto_init() }
func file_micashared_common_v1_discount_proto_init() {
	if File_micashared_common_v1_discount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_micashared_common_v1_discount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountLineItem); i {
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
		file_micashared_common_v1_discount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppliedDiscount); i {
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
			RawDescriptor: file_micashared_common_v1_discount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_v1_discount_proto_goTypes,
		DependencyIndexes: file_micashared_common_v1_discount_proto_depIdxs,
		MessageInfos:      file_micashared_common_v1_discount_proto_msgTypes,
	}.Build()
	File_micashared_common_v1_discount_proto = out.File
	file_micashared_common_v1_discount_proto_rawDesc = nil
	file_micashared_common_v1_discount_proto_goTypes = nil
	file_micashared_common_v1_discount_proto_depIdxs = nil
}
