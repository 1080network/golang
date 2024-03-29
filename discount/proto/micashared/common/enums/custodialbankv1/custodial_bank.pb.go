// Copyright (c) 2022-2024 Mica. All rights reserved. All software, including, without limitation, all source
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
// source: micashared/common/enums/custodialbank/v1/custodial_bank.proto

package custodialbankv1

import (
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

type CustodialBank int32

const (
	CustodialBank_CUSTODIAL_BANK_UNSPECIFIED     CustodialBank = 0
	CustodialBank_CUSTODIAL_BANK_TEST_BANK       CustodialBank = 1
	CustodialBank_CUSTODIAL_BANK_MAINSTREET_BANK CustodialBank = 2
)

// Enum value maps for CustodialBank.
var (
	CustodialBank_name = map[int32]string{
		0: "CUSTODIAL_BANK_UNSPECIFIED",
		1: "CUSTODIAL_BANK_TEST_BANK",
		2: "CUSTODIAL_BANK_MAINSTREET_BANK",
	}
	CustodialBank_value = map[string]int32{
		"CUSTODIAL_BANK_UNSPECIFIED":     0,
		"CUSTODIAL_BANK_TEST_BANK":       1,
		"CUSTODIAL_BANK_MAINSTREET_BANK": 2,
	}
)

func (x CustodialBank) Enum() *CustodialBank {
	p := new(CustodialBank)
	*p = x
	return p
}

func (x CustodialBank) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustodialBank) Descriptor() protoreflect.EnumDescriptor {
	return file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_enumTypes[0].Descriptor()
}

func (CustodialBank) Type() protoreflect.EnumType {
	return &file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_enumTypes[0]
}

func (x CustodialBank) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustodialBank.Descriptor instead.
func (CustodialBank) EnumDescriptor() ([]byte, []int) {
	return file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_rawDescGZIP(), []int{0}
}

var File_micashared_common_enums_custodialbank_v1_custodial_bank_proto protoreflect.FileDescriptor

var file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64,
	0x69, 0x61, 0x6c, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x64, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x28, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69,
	0x61, 0x6c, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2a, 0x71, 0x0a, 0x0d, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x64, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x55,
	0x53, 0x54, 0x4f, 0x44, 0x49, 0x41, 0x4c, 0x5f, 0x42, 0x41, 0x4e, 0x4b, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x55,
	0x53, 0x54, 0x4f, 0x44, 0x49, 0x41, 0x4c, 0x5f, 0x42, 0x41, 0x4e, 0x4b, 0x5f, 0x54, 0x45, 0x53,
	0x54, 0x5f, 0x42, 0x41, 0x4e, 0x4b, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x55, 0x53, 0x54,
	0x4f, 0x44, 0x49, 0x41, 0x4c, 0x5f, 0x42, 0x41, 0x4e, 0x4b, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x53,
	0x54, 0x52, 0x45, 0x45, 0x54, 0x5f, 0x42, 0x41, 0x4e, 0x4b, 0x10, 0x02, 0x42, 0x67, 0x0a, 0x1d,
	0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x6e, 0x6b, 0x45, 0x6e, 0x75, 0x6d,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x27, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6c, 0x62, 0x61, 0x6e, 0x6b, 0x76, 0x31, 0xa2, 0x02,
	0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_rawDescOnce sync.Once
	file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_rawDescData = file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_rawDesc
)

func file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_rawDescGZIP() []byte {
	file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_rawDescOnce.Do(func() {
		file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_rawDescData)
	})
	return file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_rawDescData
}

var file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_goTypes = []interface{}{
	(CustodialBank)(0), // 0: micashared.common.enums.custodialbank.v1.CustodialBank
}
var file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_init() }
func file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_init() {
	if File_micashared_common_enums_custodialbank_v1_custodial_bank_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_goTypes,
		DependencyIndexes: file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_depIdxs,
		EnumInfos:         file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_enumTypes,
	}.Build()
	File_micashared_common_enums_custodialbank_v1_custodial_bank_proto = out.File
	file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_rawDesc = nil
	file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_goTypes = nil
	file_micashared_common_enums_custodialbank_v1_custodial_bank_proto_depIdxs = nil
}
