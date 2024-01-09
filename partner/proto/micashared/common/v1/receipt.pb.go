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
// source: micashared/common/v1/receipt.proto

package commonv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

	// Types that are assignable to Image:
	//
	//	*Receipt_Data
	//	*Receipt_Url
	Image isReceipt_Image `protobuf_oneof:"image"`
	// The data used to generate the image.
	SourceData []byte `protobuf:"bytes,31,opt,name=source_data,json=sourceData,proto3" json:"source_data,omitempty"`
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

func (m *Receipt) GetImage() isReceipt_Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (x *Receipt) GetData() []byte {
	if x, ok := x.GetImage().(*Receipt_Data); ok {
		return x.Data
	}
	return nil
}

func (x *Receipt) GetUrl() string {
	if x, ok := x.GetImage().(*Receipt_Url); ok {
		return x.Url
	}
	return ""
}

func (x *Receipt) GetSourceData() []byte {
	if x != nil {
		return x.SourceData
	}
	return nil
}

type isReceipt_Image interface {
	isReceipt_Image()
}

type Receipt_Data struct {
	// The generated image bytes.
	Data []byte `protobuf:"bytes,29,opt,name=data,proto3,oneof"`
}

type Receipt_Url struct {
	// The url where to pull the generated image from, if applicable.
	Url string `protobuf:"bytes,30,opt,name=url,proto3,oneof"`
}

func (*Receipt_Data) isReceipt_Image() {}

func (*Receipt_Url) isReceipt_Image() {}

type GetReceiptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mica's transaction record primary key.
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

type ReceiptTransactionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// if added the receipt will encode this value into barcode configured for the organization this can be later scanned
	// to find the corresponding receipt in the store/merchant system
	BarcodeText string `protobuf:"bytes,1,opt,name=barcode_text,json=barcodeText,proto3" json:"barcode_text,omitempty"`
}

func (x *ReceiptTransactionData) Reset() {
	*x = ReceiptTransactionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_receipt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiptTransactionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiptTransactionData) ProtoMessage() {}

func (x *ReceiptTransactionData) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_receipt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiptTransactionData.ProtoReflect.Descriptor instead.
func (*ReceiptTransactionData) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_receipt_proto_rawDescGZIP(), []int{3}
}

func (x *ReceiptTransactionData) GetBarcodeText() string {
	if x != nil {
		return x.BarcodeText
	}
	return ""
}

var File_micashared_common_v1_receipt_proto protoreflect.FileDescriptor

var file_micashared_common_v1_receipt_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x07, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x22, 0xa7, 0x02,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x37, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x22, 0x3b, 0x0a, 0x16, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x54, 0x65, 0x78, 0x74, 0x42, 0x4d, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42,
	0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d,
	0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_micashared_common_v1_receipt_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_micashared_common_v1_receipt_proto_goTypes = []interface{}{
	(GetReceiptResponse_Status)(0), // 0: micashared.common.v1.GetReceiptResponse.Status
	(*Receipt)(nil),                // 1: micashared.common.v1.Receipt
	(*GetReceiptRequest)(nil),      // 2: micashared.common.v1.GetReceiptRequest
	(*GetReceiptResponse)(nil),     // 3: micashared.common.v1.GetReceiptResponse
	(*ReceiptTransactionData)(nil), // 4: micashared.common.v1.ReceiptTransactionData
	(*Error)(nil),                  // 5: micashared.common.v1.Error
}
var file_micashared_common_v1_receipt_proto_depIdxs = []int32{
	0, // 0: micashared.common.v1.GetReceiptResponse.status:type_name -> micashared.common.v1.GetReceiptResponse.Status
	5, // 1: micashared.common.v1.GetReceiptResponse.error:type_name -> micashared.common.v1.Error
	1, // 2: micashared.common.v1.GetReceiptResponse.receipt:type_name -> micashared.common.v1.Receipt
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_micashared_common_v1_receipt_proto_init() }
func file_micashared_common_v1_receipt_proto_init() {
	if File_micashared_common_v1_receipt_proto != nil {
		return
	}
	file_micashared_common_v1_error_proto_init()
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
		file_micashared_common_v1_receipt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiptTransactionData); i {
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
	file_micashared_common_v1_receipt_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Receipt_Data)(nil),
		(*Receipt_Url)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_micashared_common_v1_receipt_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
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
