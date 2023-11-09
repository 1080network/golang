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
// source: micashared/common/enums/transactiontype/v1/transaction_type.proto

package transactiontypev1

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

type TransactionType int32

const (
	TransactionType_TRANSACTION_TYPE_UNSPECIFIED TransactionType = 0
	TransactionType_TRANSACTION_TYPE_OBTAIN      TransactionType = 1
	TransactionType_TRANSACTION_TYPE_RETURN      TransactionType = 2
	TransactionType_TRANSACTION_TYPE_P2P         TransactionType = 3
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "TRANSACTION_TYPE_UNSPECIFIED",
		1: "TRANSACTION_TYPE_OBTAIN",
		2: "TRANSACTION_TYPE_RETURN",
		3: "TRANSACTION_TYPE_P2P",
	}
	TransactionType_value = map[string]int32{
		"TRANSACTION_TYPE_UNSPECIFIED": 0,
		"TRANSACTION_TYPE_OBTAIN":      1,
		"TRANSACTION_TYPE_RETURN":      2,
		"TRANSACTION_TYPE_P2P":         3,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_micashared_common_enums_transactiontype_v1_transaction_type_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_micashared_common_enums_transactiontype_v1_transaction_type_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_micashared_common_enums_transactiontype_v1_transaction_type_proto_rawDescGZIP(), []int{0}
}

var File_micashared_common_enums_transactiontype_v1_transaction_type_proto protoreflect.FileDescriptor

var file_micashared_common_enums_transactiontype_v1_transaction_type_proto_rawDesc = []byte{
	0x0a, 0x41, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2a,
	0x87, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x42, 0x54, 0x41, 0x49, 0x4e,
	0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x54, 0x55, 0x52, 0x4e, 0x10, 0x02, 0x12,
	0x18, 0x0a, 0x14, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x32, 0x50, 0x10, 0x03, 0x42, 0x6b, 0x0a, 0x1d, 0x69, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x18, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x29, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x76, 0x31, 0xa2,
	0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_enums_transactiontype_v1_transaction_type_proto_rawDescOnce sync.Once
	file_micashared_common_enums_transactiontype_v1_transaction_type_proto_rawDescData = file_micashared_common_enums_transactiontype_v1_transaction_type_proto_rawDesc
)

func file_micashared_common_enums_transactiontype_v1_transaction_type_proto_rawDescGZIP() []byte {
	file_micashared_common_enums_transactiontype_v1_transaction_type_proto_rawDescOnce.Do(func() {
		file_micashared_common_enums_transactiontype_v1_transaction_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_enums_transactiontype_v1_transaction_type_proto_rawDescData)
	})
	return file_micashared_common_enums_transactiontype_v1_transaction_type_proto_rawDescData
}

var file_micashared_common_enums_transactiontype_v1_transaction_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_micashared_common_enums_transactiontype_v1_transaction_type_proto_goTypes = []interface{}{
	(TransactionType)(0), // 0: micashared.common.enums.transactiontype.v1.TransactionType
}
var file_micashared_common_enums_transactiontype_v1_transaction_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_micashared_common_enums_transactiontype_v1_transaction_type_proto_init() }
func file_micashared_common_enums_transactiontype_v1_transaction_type_proto_init() {
	if File_micashared_common_enums_transactiontype_v1_transaction_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_micashared_common_enums_transactiontype_v1_transaction_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_enums_transactiontype_v1_transaction_type_proto_goTypes,
		DependencyIndexes: file_micashared_common_enums_transactiontype_v1_transaction_type_proto_depIdxs,
		EnumInfos:         file_micashared_common_enums_transactiontype_v1_transaction_type_proto_enumTypes,
	}.Build()
	File_micashared_common_enums_transactiontype_v1_transaction_type_proto = out.File
	file_micashared_common_enums_transactiontype_v1_transaction_type_proto_rawDesc = nil
	file_micashared_common_enums_transactiontype_v1_transaction_type_proto_goTypes = nil
	file_micashared_common_enums_transactiontype_v1_transaction_type_proto_depIdxs = nil
}