// Copyright (c) 2022 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: common/enums/accounttype/v1/account_type.proto

package accounttypev1

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

type AccountType int32

const (
	AccountType_ACCOUNT_TYPE_UNSPECIFIED AccountType = 0
	AccountType_ACCOUNT_TYPE_REVENUE     AccountType = 1
)

// Enum value maps for AccountType.
var (
	AccountType_name = map[int32]string{
		0: "ACCOUNT_TYPE_UNSPECIFIED",
		1: "ACCOUNT_TYPE_REVENUE",
	}
	AccountType_value = map[string]int32{
		"ACCOUNT_TYPE_UNSPECIFIED": 0,
		"ACCOUNT_TYPE_REVENUE":     1,
	}
)

func (x AccountType) Enum() *AccountType {
	p := new(AccountType)
	*p = x
	return p
}

func (x AccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enums_accounttype_v1_account_type_proto_enumTypes[0].Descriptor()
}

func (AccountType) Type() protoreflect.EnumType {
	return &file_common_enums_accounttype_v1_account_type_proto_enumTypes[0]
}

func (x AccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountType.Descriptor instead.
func (AccountType) EnumDescriptor() ([]byte, []int) {
	return file_common_enums_accounttype_v1_account_type_proto_rawDescGZIP(), []int{0}
}

var File_common_enums_accounttype_v1_account_type_proto protoreflect.FileDescriptor

var file_common_enums_accounttype_v1_account_type_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2a, 0x45, 0x0a,
	0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18,
	0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x43,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x56, 0x45, 0x4e,
	0x55, 0x45, 0x10, 0x01, 0x42, 0x5d, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42,
	0x14, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x25, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d,
	0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_enums_accounttype_v1_account_type_proto_rawDescOnce sync.Once
	file_common_enums_accounttype_v1_account_type_proto_rawDescData = file_common_enums_accounttype_v1_account_type_proto_rawDesc
)

func file_common_enums_accounttype_v1_account_type_proto_rawDescGZIP() []byte {
	file_common_enums_accounttype_v1_account_type_proto_rawDescOnce.Do(func() {
		file_common_enums_accounttype_v1_account_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_enums_accounttype_v1_account_type_proto_rawDescData)
	})
	return file_common_enums_accounttype_v1_account_type_proto_rawDescData
}

var file_common_enums_accounttype_v1_account_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_enums_accounttype_v1_account_type_proto_goTypes = []interface{}{
	(AccountType)(0), // 0: common.enums.accounttype.v1.AccountType
}
var file_common_enums_accounttype_v1_account_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_enums_accounttype_v1_account_type_proto_init() }
func file_common_enums_accounttype_v1_account_type_proto_init() {
	if File_common_enums_accounttype_v1_account_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_enums_accounttype_v1_account_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_enums_accounttype_v1_account_type_proto_goTypes,
		DependencyIndexes: file_common_enums_accounttype_v1_account_type_proto_depIdxs,
		EnumInfos:         file_common_enums_accounttype_v1_account_type_proto_enumTypes,
	}.Build()
	File_common_enums_accounttype_v1_account_type_proto = out.File
	file_common_enums_accounttype_v1_account_type_proto_rawDesc = nil
	file_common_enums_accounttype_v1_account_type_proto_goTypes = nil
	file_common_enums_accounttype_v1_account_type_proto_depIdxs = nil
}
