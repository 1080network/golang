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
// source: micashared/common/enums/instrumenttype/v1/instrument_type.proto

package instrumenttypev1

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

type InstrumentType int32

const (
	InstrumentType_INSTRUMENT_TYPE_UNSPECIFIED    InstrumentType = 0
	InstrumentType_INSTRUMENT_TYPE_CHECKING       InstrumentType = 1
	InstrumentType_INSTRUMENT_TYPE_SAVINGS        InstrumentType = 2
	InstrumentType_INSTRUMENT_TYPE_MONEY_MARKET   InstrumentType = 3
	InstrumentType_INSTRUMENT_TYPE_LINE_OF_CREDIT InstrumentType = 4
	// US only - Supplemental Nutrition Assistance Program
	InstrumentType_INSTRUMENT_TYPE_SNAP InstrumentType = 5
	// US only - Temporary Assistance for Needy Families
	InstrumentType_INSTRUMENT_TYPE_TANF InstrumentType = 6
	// US only - Special Supplemental Nutrition Program for Women, Infants, and Children
	InstrumentType_INSTRUMENT_TYPE_WIC InstrumentType = 7
	// Program that helps to pay for nycloop
	InstrumentType_INSTRUMENT_TYPE_NYCLOOP        InstrumentType = 8
	InstrumentType_INSTRUMENT_TYPE_POSITIVE_VALUE InstrumentType = 9
	InstrumentType_INSTRUMENT_TYPE_ADJUDICATED    InstrumentType = 10
)

// Enum value maps for InstrumentType.
var (
	InstrumentType_name = map[int32]string{
		0:  "INSTRUMENT_TYPE_UNSPECIFIED",
		1:  "INSTRUMENT_TYPE_CHECKING",
		2:  "INSTRUMENT_TYPE_SAVINGS",
		3:  "INSTRUMENT_TYPE_MONEY_MARKET",
		4:  "INSTRUMENT_TYPE_LINE_OF_CREDIT",
		5:  "INSTRUMENT_TYPE_SNAP",
		6:  "INSTRUMENT_TYPE_TANF",
		7:  "INSTRUMENT_TYPE_WIC",
		8:  "INSTRUMENT_TYPE_NYCLOOP",
		9:  "INSTRUMENT_TYPE_POSITIVE_VALUE",
		10: "INSTRUMENT_TYPE_ADJUDICATED",
	}
	InstrumentType_value = map[string]int32{
		"INSTRUMENT_TYPE_UNSPECIFIED":    0,
		"INSTRUMENT_TYPE_CHECKING":       1,
		"INSTRUMENT_TYPE_SAVINGS":        2,
		"INSTRUMENT_TYPE_MONEY_MARKET":   3,
		"INSTRUMENT_TYPE_LINE_OF_CREDIT": 4,
		"INSTRUMENT_TYPE_SNAP":           5,
		"INSTRUMENT_TYPE_TANF":           6,
		"INSTRUMENT_TYPE_WIC":            7,
		"INSTRUMENT_TYPE_NYCLOOP":        8,
		"INSTRUMENT_TYPE_POSITIVE_VALUE": 9,
		"INSTRUMENT_TYPE_ADJUDICATED":    10,
	}
)

func (x InstrumentType) Enum() *InstrumentType {
	p := new(InstrumentType)
	*p = x
	return p
}

func (x InstrumentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstrumentType) Descriptor() protoreflect.EnumDescriptor {
	return file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_enumTypes[0].Descriptor()
}

func (InstrumentType) Type() protoreflect.EnumType {
	return &file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_enumTypes[0]
}

func (x InstrumentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstrumentType.Descriptor instead.
func (InstrumentType) EnumDescriptor() ([]byte, []int) {
	return file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_rawDescGZIP(), []int{0}
}

var File_micashared_common_enums_instrumenttype_v1_instrument_type_proto protoreflect.FileDescriptor

var file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x29, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2a, 0xe1, 0x02, 0x0a,
	0x0e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1b,
	0x0a, 0x17, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x41, 0x56, 0x49, 0x4e, 0x47, 0x53, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x49,
	0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d,
	0x4f, 0x4e, 0x45, 0x59, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x10, 0x03, 0x12, 0x22, 0x0a,
	0x1e, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x4f, 0x46, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x10,
	0x04, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x49,
	0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x41, 0x4e, 0x46, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x49, 0x43, 0x10, 0x07, 0x12, 0x1b,
	0x0a, 0x17, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4e, 0x59, 0x43, 0x4c, 0x4f, 0x4f, 0x50, 0x10, 0x08, 0x12, 0x22, 0x0a, 0x1e, 0x49,
	0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x4f, 0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x09, 0x12,
	0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x44, 0x4a, 0x55, 0x44, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x0a,
	0x42, 0x69, 0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x42, 0x17, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x28, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x74, 0x79,
	0x70, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_rawDescOnce sync.Once
	file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_rawDescData = file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_rawDesc
)

func file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_rawDescGZIP() []byte {
	file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_rawDescOnce.Do(func() {
		file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_rawDescData)
	})
	return file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_rawDescData
}

var file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_goTypes = []interface{}{
	(InstrumentType)(0), // 0: micashared.common.enums.instrumenttype.v1.InstrumentType
}
var file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_init() }
func file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_init() {
	if File_micashared_common_enums_instrumenttype_v1_instrument_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_goTypes,
		DependencyIndexes: file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_depIdxs,
		EnumInfos:         file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_enumTypes,
	}.Build()
	File_micashared_common_enums_instrumenttype_v1_instrument_type_proto = out.File
	file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_rawDesc = nil
	file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_goTypes = nil
	file_micashared_common_enums_instrumenttype_v1_instrument_type_proto_depIdxs = nil
}
