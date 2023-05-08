// Copyright (c) 2022 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: micashared/common/enums/unit/v1/unit.proto

package unitv1

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

type Unit int32

const (
	Unit_UNIT_UNSPECIFIED Unit = 0
	Unit_UNIT_ITEM        Unit = 1
	// weight
	Unit_UNIT_GRAM     Unit = 2
	Unit_UNIT_KILOGRAM Unit = 3
	Unit_UNIT_POUND    Unit = 4
	Unit_UNIT_OUNCE    Unit = 5
	// liquid
	Unit_UNIT_LITRE       Unit = 6
	Unit_UNIT_MILLILITRE  Unit = 7
	Unit_UNIT_GALLON      Unit = 8
	Unit_UNIT_QUART       Unit = 9
	Unit_UNIT_PINT        Unit = 10
	Unit_UNIT_FLUID_OUNCE Unit = 11
	// length
	Unit_UNIT_METRE      Unit = 12
	Unit_UNIT_CENTIMETRE Unit = 13
	Unit_UNIT_MILLIMETRE Unit = 14
	Unit_UNIT_YARD       Unit = 15
	Unit_UNIT_FOOT       Unit = 16
	Unit_UNIT_INCH       Unit = 17
)

// Enum value maps for Unit.
var (
	Unit_name = map[int32]string{
		0:  "UNIT_UNSPECIFIED",
		1:  "UNIT_ITEM",
		2:  "UNIT_GRAM",
		3:  "UNIT_KILOGRAM",
		4:  "UNIT_POUND",
		5:  "UNIT_OUNCE",
		6:  "UNIT_LITRE",
		7:  "UNIT_MILLILITRE",
		8:  "UNIT_GALLON",
		9:  "UNIT_QUART",
		10: "UNIT_PINT",
		11: "UNIT_FLUID_OUNCE",
		12: "UNIT_METRE",
		13: "UNIT_CENTIMETRE",
		14: "UNIT_MILLIMETRE",
		15: "UNIT_YARD",
		16: "UNIT_FOOT",
		17: "UNIT_INCH",
	}
	Unit_value = map[string]int32{
		"UNIT_UNSPECIFIED": 0,
		"UNIT_ITEM":        1,
		"UNIT_GRAM":        2,
		"UNIT_KILOGRAM":    3,
		"UNIT_POUND":       4,
		"UNIT_OUNCE":       5,
		"UNIT_LITRE":       6,
		"UNIT_MILLILITRE":  7,
		"UNIT_GALLON":      8,
		"UNIT_QUART":       9,
		"UNIT_PINT":        10,
		"UNIT_FLUID_OUNCE": 11,
		"UNIT_METRE":       12,
		"UNIT_CENTIMETRE":  13,
		"UNIT_MILLIMETRE":  14,
		"UNIT_YARD":        15,
		"UNIT_FOOT":        16,
		"UNIT_INCH":        17,
	}
)

func (x Unit) Enum() *Unit {
	p := new(Unit)
	*p = x
	return p
}

func (x Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_micashared_common_enums_unit_v1_unit_proto_enumTypes[0].Descriptor()
}

func (Unit) Type() protoreflect.EnumType {
	return &file_micashared_common_enums_unit_v1_unit_proto_enumTypes[0]
}

func (x Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unit.Descriptor instead.
func (Unit) EnumDescriptor() ([]byte, []int) {
	return file_micashared_common_enums_unit_v1_unit_proto_rawDescGZIP(), []int{0}
}

var File_micashared_common_enums_unit_v1_unit_proto protoreflect.FileDescriptor

var file_micashared_common_enums_unit_v1_unit_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2a, 0xbf, 0x02,
	0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x55, 0x4e, 0x49, 0x54, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x55,
	0x4e, 0x49, 0x54, 0x5f, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e,
	0x49, 0x54, 0x5f, 0x4b, 0x49, 0x4c, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x03, 0x12, 0x0e, 0x0a,
	0x0a, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x50, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x04, 0x12, 0x0e, 0x0a,
	0x0a, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4f, 0x55, 0x4e, 0x43, 0x45, 0x10, 0x05, 0x12, 0x0e, 0x0a,
	0x0a, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4c, 0x49, 0x54, 0x52, 0x45, 0x10, 0x06, 0x12, 0x13, 0x0a,
	0x0f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4c, 0x49, 0x54, 0x52, 0x45,
	0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x47, 0x41, 0x4c, 0x4c, 0x4f,
	0x4e, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x51, 0x55, 0x41, 0x52,
	0x54, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x50, 0x49, 0x4e, 0x54,
	0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x46, 0x4c, 0x55, 0x49, 0x44,
	0x5f, 0x4f, 0x55, 0x4e, 0x43, 0x45, 0x10, 0x0b, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x49, 0x54,
	0x5f, 0x4d, 0x45, 0x54, 0x52, 0x45, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x49, 0x54,
	0x5f, 0x43, 0x45, 0x4e, 0x54, 0x49, 0x4d, 0x45, 0x54, 0x52, 0x45, 0x10, 0x0d, 0x12, 0x13, 0x0a,
	0x0f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4d, 0x45, 0x54, 0x52, 0x45,
	0x10, 0x0e, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x59, 0x41, 0x52, 0x44, 0x10,
	0x0f, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x46, 0x4f, 0x4f, 0x54, 0x10, 0x10,
	0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x49, 0x4e, 0x43, 0x48, 0x10, 0x11, 0x42,
	0x55, 0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x0d, 0x55, 0x6e, 0x69, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a,
	0x1e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x76, 0x31, 0xa2,
	0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_enums_unit_v1_unit_proto_rawDescOnce sync.Once
	file_micashared_common_enums_unit_v1_unit_proto_rawDescData = file_micashared_common_enums_unit_v1_unit_proto_rawDesc
)

func file_micashared_common_enums_unit_v1_unit_proto_rawDescGZIP() []byte {
	file_micashared_common_enums_unit_v1_unit_proto_rawDescOnce.Do(func() {
		file_micashared_common_enums_unit_v1_unit_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_enums_unit_v1_unit_proto_rawDescData)
	})
	return file_micashared_common_enums_unit_v1_unit_proto_rawDescData
}

var file_micashared_common_enums_unit_v1_unit_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_micashared_common_enums_unit_v1_unit_proto_goTypes = []interface{}{
	(Unit)(0), // 0: micashared.common.enums.unit.v1.Unit
}
var file_micashared_common_enums_unit_v1_unit_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_micashared_common_enums_unit_v1_unit_proto_init() }
func file_micashared_common_enums_unit_v1_unit_proto_init() {
	if File_micashared_common_enums_unit_v1_unit_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_micashared_common_enums_unit_v1_unit_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_enums_unit_v1_unit_proto_goTypes,
		DependencyIndexes: file_micashared_common_enums_unit_v1_unit_proto_depIdxs,
		EnumInfos:         file_micashared_common_enums_unit_v1_unit_proto_enumTypes,
	}.Build()
	File_micashared_common_enums_unit_v1_unit_proto = out.File
	file_micashared_common_enums_unit_v1_unit_proto_rawDesc = nil
	file_micashared_common_enums_unit_v1_unit_proto_goTypes = nil
	file_micashared_common_enums_unit_v1_unit_proto_depIdxs = nil
}