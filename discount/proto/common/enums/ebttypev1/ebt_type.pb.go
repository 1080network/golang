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
// source: common/enums/ebttype/v1/ebt_type.proto

package ebttypev1

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

type EbtType int32

const (
	EbtType_EBT_TYPE_UNSPECIFIED EbtType = 0
	// US only - Supplemental Nutrition Assistance Program
	EbtType_EBT_TYPE_SNAP EbtType = 1
	// US only - Temporary Assistance for Needy Families
	EbtType_EBT_TYPE_TANF EbtType = 2
	// US only - Special Supplemental Nutrition Program for Women, Infants, and Children
	EbtType_EBT_TYPE_WIC EbtType = 3
)

// Enum value maps for EbtType.
var (
	EbtType_name = map[int32]string{
		0: "EBT_TYPE_UNSPECIFIED",
		1: "EBT_TYPE_SNAP",
		2: "EBT_TYPE_TANF",
		3: "EBT_TYPE_WIC",
	}
	EbtType_value = map[string]int32{
		"EBT_TYPE_UNSPECIFIED": 0,
		"EBT_TYPE_SNAP":        1,
		"EBT_TYPE_TANF":        2,
		"EBT_TYPE_WIC":         3,
	}
)

func (x EbtType) Enum() *EbtType {
	p := new(EbtType)
	*p = x
	return p
}

func (x EbtType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EbtType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enums_ebttype_v1_ebt_type_proto_enumTypes[0].Descriptor()
}

func (EbtType) Type() protoreflect.EnumType {
	return &file_common_enums_ebttype_v1_ebt_type_proto_enumTypes[0]
}

func (x EbtType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EbtType.Descriptor instead.
func (EbtType) EnumDescriptor() ([]byte, []int) {
	return file_common_enums_ebttype_v1_ebt_type_proto_rawDescGZIP(), []int{0}
}

var File_common_enums_ebttype_v1_ebt_type_proto protoreflect.FileDescriptor

var file_common_enums_ebttype_v1_ebt_type_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65,
	0x62, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x62, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x65, 0x62, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x31, 0x2a, 0x5b, 0x0a, 0x07, 0x45, 0x62, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14,
	0x45, 0x42, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x42, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x42, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x41, 0x4e, 0x46, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x45, 0x42, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x49, 0x43, 0x10, 0x03, 0x42, 0x55,
	0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x45, 0x62, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x21, 0x6d, 0x69, 0x63,
	0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x62, 0x74, 0x74, 0x79, 0x70, 0x65, 0x76, 0x31, 0xa2, 0x02,
	0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_enums_ebttype_v1_ebt_type_proto_rawDescOnce sync.Once
	file_common_enums_ebttype_v1_ebt_type_proto_rawDescData = file_common_enums_ebttype_v1_ebt_type_proto_rawDesc
)

func file_common_enums_ebttype_v1_ebt_type_proto_rawDescGZIP() []byte {
	file_common_enums_ebttype_v1_ebt_type_proto_rawDescOnce.Do(func() {
		file_common_enums_ebttype_v1_ebt_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_enums_ebttype_v1_ebt_type_proto_rawDescData)
	})
	return file_common_enums_ebttype_v1_ebt_type_proto_rawDescData
}

var file_common_enums_ebttype_v1_ebt_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_enums_ebttype_v1_ebt_type_proto_goTypes = []interface{}{
	(EbtType)(0), // 0: common.enums.ebttype.v1.EbtType
}
var file_common_enums_ebttype_v1_ebt_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_enums_ebttype_v1_ebt_type_proto_init() }
func file_common_enums_ebttype_v1_ebt_type_proto_init() {
	if File_common_enums_ebttype_v1_ebt_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_enums_ebttype_v1_ebt_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_enums_ebttype_v1_ebt_type_proto_goTypes,
		DependencyIndexes: file_common_enums_ebttype_v1_ebt_type_proto_depIdxs,
		EnumInfos:         file_common_enums_ebttype_v1_ebt_type_proto_enumTypes,
	}.Build()
	File_common_enums_ebttype_v1_ebt_type_proto = out.File
	file_common_enums_ebttype_v1_ebt_type_proto_rawDesc = nil
	file_common_enums_ebttype_v1_ebt_type_proto_goTypes = nil
	file_common_enums_ebttype_v1_ebt_type_proto_depIdxs = nil
}
