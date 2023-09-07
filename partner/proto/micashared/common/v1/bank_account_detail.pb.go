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
// source: micashared/common/v1/bank_account_detail.proto

package commonv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	custodialbankv1 "github.com/1080network/golang/partner/proto/micashared/common/enums/custodialbankv1"
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

	// the custodial bank where this transaction will be settled
	CustodialBank custodialbankv1.CustodialBank `protobuf:"varint,6,opt,name=custodial_bank,json=custodialBank,proto3,enum=micashared.common.enums.custodialbank.v1.CustodialBank" json:"custodial_bank,omitempty"`
	// The bank at which this account is held at
	Bank string `protobuf:"bytes,8,opt,name=bank,proto3" json:"bank,omitempty"`
	// The account reference for this account
	AccountRef string `protobuf:"bytes,7,opt,name=account_ref,json=accountRef,proto3" json:"account_ref,omitempty"`
}

func (x *BankAccountDetail) Reset() {
	*x = BankAccountDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_bank_account_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankAccountDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankAccountDetail) ProtoMessage() {}

func (x *BankAccountDetail) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_bank_account_detail_proto_msgTypes[0]
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
	return file_micashared_common_v1_bank_account_detail_proto_rawDescGZIP(), []int{0}
}

func (x *BankAccountDetail) GetCustodialBank() custodialbankv1.CustodialBank {
	if x != nil {
		return x.CustodialBank
	}
	return custodialbankv1.CustodialBank(0)
}

func (x *BankAccountDetail) GetBank() string {
	if x != nil {
		return x.Bank
	}
	return ""
}

func (x *BankAccountDetail) GetAccountRef() string {
	if x != nil {
		return x.AccountRef
	}
	return ""
}

type BankAccountDetailSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the custodial bank where this transaction will be settled
	CustodialBank custodialbankv1.CustodialBank `protobuf:"varint,1,opt,name=custodial_bank,json=custodialBank,proto3,enum=micashared.common.enums.custodialbank.v1.CustodialBank" json:"custodial_bank,omitempty"`
	// The bank at which this account is held at
	Bank string `protobuf:"bytes,2,opt,name=bank,proto3" json:"bank,omitempty"`
	// The account reference for this account
	AccountRef string `protobuf:"bytes,3,opt,name=account_ref,json=accountRef,proto3" json:"account_ref,omitempty"`
}

func (x *BankAccountDetailSearch) Reset() {
	*x = BankAccountDetailSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_bank_account_detail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankAccountDetailSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankAccountDetailSearch) ProtoMessage() {}

func (x *BankAccountDetailSearch) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_bank_account_detail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankAccountDetailSearch.ProtoReflect.Descriptor instead.
func (*BankAccountDetailSearch) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_bank_account_detail_proto_rawDescGZIP(), []int{1}
}

func (x *BankAccountDetailSearch) GetCustodialBank() custodialbankv1.CustodialBank {
	if x != nil {
		return x.CustodialBank
	}
	return custodialbankv1.CustodialBank(0)
}

func (x *BankAccountDetailSearch) GetBank() string {
	if x != nil {
		return x.Bank
	}
	return ""
}

func (x *BankAccountDetailSearch) GetAccountRef() string {
	if x != nil {
		return x.AccountRef
	}
	return ""
}

var File_micashared_common_v1_bank_account_detail_proto protoreflect.FileDescriptor

var file_micashared_common_v1_bank_account_detail_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x3d, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6c, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe,
	0x01, 0x0a, 0x11, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x5e, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61,
	0x6c, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6c,
	0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61,
	0x6c, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6c,
	0x42, 0x61, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x04, 0x62,
	0x61, 0x6e, 0x6b, 0x12, 0x2a, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72,
	0x65, 0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x01, 0x18, 0x32, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x22,
	0xae, 0x01, 0x0a, 0x17, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x5e, 0x0a, 0x0e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6c, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x0d, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x61, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x66,
	0x42, 0x57, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x42, 0x61, 0x6e,
	0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_micashared_common_v1_bank_account_detail_proto_rawDescOnce sync.Once
	file_micashared_common_v1_bank_account_detail_proto_rawDescData = file_micashared_common_v1_bank_account_detail_proto_rawDesc
)

func file_micashared_common_v1_bank_account_detail_proto_rawDescGZIP() []byte {
	file_micashared_common_v1_bank_account_detail_proto_rawDescOnce.Do(func() {
		file_micashared_common_v1_bank_account_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_v1_bank_account_detail_proto_rawDescData)
	})
	return file_micashared_common_v1_bank_account_detail_proto_rawDescData
}

var file_micashared_common_v1_bank_account_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_micashared_common_v1_bank_account_detail_proto_goTypes = []interface{}{
	(*BankAccountDetail)(nil),          // 0: micashared.common.v1.BankAccountDetail
	(*BankAccountDetailSearch)(nil),    // 1: micashared.common.v1.BankAccountDetailSearch
	(custodialbankv1.CustodialBank)(0), // 2: micashared.common.enums.custodialbank.v1.CustodialBank
}
var file_micashared_common_v1_bank_account_detail_proto_depIdxs = []int32{
	2, // 0: micashared.common.v1.BankAccountDetail.custodial_bank:type_name -> micashared.common.enums.custodialbank.v1.CustodialBank
	2, // 1: micashared.common.v1.BankAccountDetailSearch.custodial_bank:type_name -> micashared.common.enums.custodialbank.v1.CustodialBank
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_micashared_common_v1_bank_account_detail_proto_init() }
func file_micashared_common_v1_bank_account_detail_proto_init() {
	if File_micashared_common_v1_bank_account_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_micashared_common_v1_bank_account_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_micashared_common_v1_bank_account_detail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankAccountDetailSearch); i {
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
			RawDescriptor: file_micashared_common_v1_bank_account_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_v1_bank_account_detail_proto_goTypes,
		DependencyIndexes: file_micashared_common_v1_bank_account_detail_proto_depIdxs,
		MessageInfos:      file_micashared_common_v1_bank_account_detail_proto_msgTypes,
	}.Build()
	File_micashared_common_v1_bank_account_detail_proto = out.File
	file_micashared_common_v1_bank_account_detail_proto_rawDesc = nil
	file_micashared_common_v1_bank_account_detail_proto_goTypes = nil
	file_micashared_common_v1_bank_account_detail_proto_depIdxs = nil
}
