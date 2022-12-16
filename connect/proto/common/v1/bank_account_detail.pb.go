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
// 	protoc        v3.21.8
// source: common/v1/bank_account_detail.proto

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

type BankAccountDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The organizations domestic ACH routing number (9 characters)
	DomesticAchRoutingNumber string `protobuf:"bytes,1,opt,name=domestic_ach_routing_number,json=domesticAchRoutingNumber,proto3" json:"domestic_ach_routing_number,omitempty"`
	// The organizations international ACH routing number (9 characters)
	InternationalAchRoutingNumber string `protobuf:"bytes,2,opt,name=international_ach_routing_number,json=internationalAchRoutingNumber,proto3" json:"international_ach_routing_number,omitempty"`
	// The organizations wire routing number (9 characters)
	WireRoutingNumber string `protobuf:"bytes,3,opt,name=wire_routing_number,json=wireRoutingNumber,proto3" json:"wire_routing_number,omitempty"`
	// The organizations swift routing number (9 characters)
	SwiftRoutingNumber string `protobuf:"bytes,4,opt,name=swift_routing_number,json=swiftRoutingNumber,proto3" json:"swift_routing_number,omitempty"`
	// The organizations bank account number (14 digits)
	BankAccountNumber string `protobuf:"bytes,5,opt,name=bank_account_number,json=bankAccountNumber,proto3" json:"bank_account_number,omitempty"`
}

func (x *BankAccountDetail) Reset() {
	*x = BankAccountDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_bank_account_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankAccountDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankAccountDetail) ProtoMessage() {}

func (x *BankAccountDetail) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_bank_account_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankAccountDetail.ProtoReflect.Descriptor instead.
func (*BankAccountDetail) Descriptor() ([]byte, []int) {
	return file_common_v1_bank_account_detail_proto_rawDescGZIP(), []int{0}
}

func (x *BankAccountDetail) GetDomesticAchRoutingNumber() string {
	if x != nil {
		return x.DomesticAchRoutingNumber
	}
	return ""
}

func (x *BankAccountDetail) GetInternationalAchRoutingNumber() string {
	if x != nil {
		return x.InternationalAchRoutingNumber
	}
	return ""
}

func (x *BankAccountDetail) GetWireRoutingNumber() string {
	if x != nil {
		return x.WireRoutingNumber
	}
	return ""
}

func (x *BankAccountDetail) GetSwiftRoutingNumber() string {
	if x != nil {
		return x.SwiftRoutingNumber
	}
	return ""
}

func (x *BankAccountDetail) GetBankAccountNumber() string {
	if x != nil {
		return x.BankAccountNumber
	}
	return ""
}

var File_common_v1_bank_account_detail_proto protoreflect.FileDescriptor

var file_common_v1_bank_account_detail_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6e, 0x6b,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x03, 0x0a, 0x11, 0x42, 0x61,
	0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x4e, 0x0a, 0x1b, 0x64, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x61, 0x63, 0x68, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x72, 0x0a, 0x32, 0x08, 0x5e, 0x7c, 0x5c,
	0x77, 0x7b, 0x39, 0x7d, 0x24, 0x52, 0x18, 0x64, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x69, 0x63, 0x41,
	0x63, 0x68, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x58, 0x0a, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x72, 0x0a,
	0x32, 0x08, 0x5e, 0x7c, 0x5c, 0x77, 0x7b, 0x39, 0x7d, 0x24, 0x52, 0x1d, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x68, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x13, 0x77, 0x69, 0x72,
	0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x72, 0x0a, 0x32, 0x08, 0x5e,
	0x7c, 0x5c, 0x77, 0x7b, 0x39, 0x7d, 0x24, 0x52, 0x11, 0x77, 0x69, 0x72, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x14, 0x73, 0x77,
	0x69, 0x66, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x72, 0x0a, 0x32,
	0x08, 0x5e, 0x7c, 0x5c, 0x77, 0x7b, 0x39, 0x7d, 0x24, 0x52, 0x12, 0x73, 0x77, 0x69, 0x66, 0x74,
	0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x40, 0x0a,
	0x13, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0x72,
	0x0b, 0x32, 0x09, 0x5e, 0x7c, 0x5c, 0x77, 0x7b, 0x31, 0x34, 0x7d, 0x24, 0x52, 0x11, 0x62, 0x61,
	0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42,
	0x51, 0x0a, 0x11, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x6d, 0x69,
	0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49,
	0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_v1_bank_account_detail_proto_rawDescOnce sync.Once
	file_common_v1_bank_account_detail_proto_rawDescData = file_common_v1_bank_account_detail_proto_rawDesc
)

func file_common_v1_bank_account_detail_proto_rawDescGZIP() []byte {
	file_common_v1_bank_account_detail_proto_rawDescOnce.Do(func() {
		file_common_v1_bank_account_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v1_bank_account_detail_proto_rawDescData)
	})
	return file_common_v1_bank_account_detail_proto_rawDescData
}

var file_common_v1_bank_account_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_v1_bank_account_detail_proto_goTypes = []interface{}{
	(*BankAccountDetail)(nil), // 0: common.v1.BankAccountDetail
}
var file_common_v1_bank_account_detail_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_v1_bank_account_detail_proto_init() }
func file_common_v1_bank_account_detail_proto_init() {
	if File_common_v1_bank_account_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_v1_bank_account_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankAccountDetail); i {
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
			RawDescriptor: file_common_v1_bank_account_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_bank_account_detail_proto_goTypes,
		DependencyIndexes: file_common_v1_bank_account_detail_proto_depIdxs,
		MessageInfos:      file_common_v1_bank_account_detail_proto_msgTypes,
	}.Build()
	File_common_v1_bank_account_detail_proto = out.File
	file_common_v1_bank_account_detail_proto_rawDesc = nil
	file_common_v1_bank_account_detail_proto_goTypes = nil
	file_common_v1_bank_account_detail_proto_depIdxs = nil
}
