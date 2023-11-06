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
// source: micashared/common/enums/ageband/v1/age_band.proto

package agebandv1

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

type AgeBand int32

const (
	AgeBand_AGE_BAND_UNSPECIFIED AgeBand = 0
	AgeBand_AGE_BAND_UNDER_13    AgeBand = 1
	AgeBand_AGE_BAND_UNDER_14    AgeBand = 2
	AgeBand_AGE_BAND_UNDER_15    AgeBand = 3
	AgeBand_AGE_BAND_UNDER_16    AgeBand = 4
	AgeBand_AGE_BAND_UNDER_17    AgeBand = 5
	AgeBand_AGE_BAND_UNDER_18    AgeBand = 6
	AgeBand_AGE_BAND_UNDER_19    AgeBand = 7
	AgeBand_AGE_BAND_UNDER_20    AgeBand = 8
	AgeBand_AGE_BAND_UNDER_21    AgeBand = 9
	AgeBand_AGE_BAND_UNDER_22    AgeBand = 10
	AgeBand_AGE_BAND_UNDER_23    AgeBand = 11
	AgeBand_AGE_BAND_UNDER_24    AgeBand = 12
	AgeBand_AGE_BAND_UNDER_25    AgeBand = 13
	AgeBand_AGE_BAND_OVER_13     AgeBand = 14
	AgeBand_AGE_BAND_OVER_14     AgeBand = 15
	AgeBand_AGE_BAND_OVER_15     AgeBand = 16
	AgeBand_AGE_BAND_OVER_16     AgeBand = 17
	AgeBand_AGE_BAND_OVER_17     AgeBand = 18
	AgeBand_AGE_BAND_OVER_18     AgeBand = 19
	AgeBand_AGE_BAND_OVER_19     AgeBand = 20
	AgeBand_AGE_BAND_OVER_20     AgeBand = 21
	AgeBand_AGE_BAND_OVER_21     AgeBand = 22
	AgeBand_AGE_BAND_OVER_22     AgeBand = 23
	AgeBand_AGE_BAND_OVER_23     AgeBand = 24
	AgeBand_AGE_BAND_OVER_24     AgeBand = 25
	AgeBand_AGE_BAND_OVER_25     AgeBand = 26
)

// Enum value maps for AgeBand.
var (
	AgeBand_name = map[int32]string{
		0:  "AGE_BAND_UNSPECIFIED",
		1:  "AGE_BAND_UNDER_13",
		2:  "AGE_BAND_UNDER_14",
		3:  "AGE_BAND_UNDER_15",
		4:  "AGE_BAND_UNDER_16",
		5:  "AGE_BAND_UNDER_17",
		6:  "AGE_BAND_UNDER_18",
		7:  "AGE_BAND_UNDER_19",
		8:  "AGE_BAND_UNDER_20",
		9:  "AGE_BAND_UNDER_21",
		10: "AGE_BAND_UNDER_22",
		11: "AGE_BAND_UNDER_23",
		12: "AGE_BAND_UNDER_24",
		13: "AGE_BAND_UNDER_25",
		14: "AGE_BAND_OVER_13",
		15: "AGE_BAND_OVER_14",
		16: "AGE_BAND_OVER_15",
		17: "AGE_BAND_OVER_16",
		18: "AGE_BAND_OVER_17",
		19: "AGE_BAND_OVER_18",
		20: "AGE_BAND_OVER_19",
		21: "AGE_BAND_OVER_20",
		22: "AGE_BAND_OVER_21",
		23: "AGE_BAND_OVER_22",
		24: "AGE_BAND_OVER_23",
		25: "AGE_BAND_OVER_24",
		26: "AGE_BAND_OVER_25",
	}
	AgeBand_value = map[string]int32{
		"AGE_BAND_UNSPECIFIED": 0,
		"AGE_BAND_UNDER_13":    1,
		"AGE_BAND_UNDER_14":    2,
		"AGE_BAND_UNDER_15":    3,
		"AGE_BAND_UNDER_16":    4,
		"AGE_BAND_UNDER_17":    5,
		"AGE_BAND_UNDER_18":    6,
		"AGE_BAND_UNDER_19":    7,
		"AGE_BAND_UNDER_20":    8,
		"AGE_BAND_UNDER_21":    9,
		"AGE_BAND_UNDER_22":    10,
		"AGE_BAND_UNDER_23":    11,
		"AGE_BAND_UNDER_24":    12,
		"AGE_BAND_UNDER_25":    13,
		"AGE_BAND_OVER_13":     14,
		"AGE_BAND_OVER_14":     15,
		"AGE_BAND_OVER_15":     16,
		"AGE_BAND_OVER_16":     17,
		"AGE_BAND_OVER_17":     18,
		"AGE_BAND_OVER_18":     19,
		"AGE_BAND_OVER_19":     20,
		"AGE_BAND_OVER_20":     21,
		"AGE_BAND_OVER_21":     22,
		"AGE_BAND_OVER_22":     23,
		"AGE_BAND_OVER_23":     24,
		"AGE_BAND_OVER_24":     25,
		"AGE_BAND_OVER_25":     26,
	}
)

func (x AgeBand) Enum() *AgeBand {
	p := new(AgeBand)
	*p = x
	return p
}

func (x AgeBand) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgeBand) Descriptor() protoreflect.EnumDescriptor {
	return file_micashared_common_enums_ageband_v1_age_band_proto_enumTypes[0].Descriptor()
}

func (AgeBand) Type() protoreflect.EnumType {
	return &file_micashared_common_enums_ageband_v1_age_band_proto_enumTypes[0]
}

func (x AgeBand) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgeBand.Descriptor instead.
func (AgeBand) EnumDescriptor() ([]byte, []int) {
	return file_micashared_common_enums_ageband_v1_age_band_proto_rawDescGZIP(), []int{0}
}

var File_micashared_common_enums_ageband_v1_age_band_proto protoreflect.FileDescriptor

var file_micashared_common_enums_ageband_v1_age_band_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x62, 0x61, 0x6e,
	0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x22, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x67, 0x65,
	0x62, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2a, 0xec, 0x04, 0x0a, 0x07, 0x41, 0x67, 0x65, 0x42,
	0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f,
	0x31, 0x33, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44,
	0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x31, 0x34, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x41,
	0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x31, 0x35,
	0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x55,
	0x4e, 0x44, 0x45, 0x52, 0x5f, 0x31, 0x36, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x47, 0x45,
	0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x31, 0x37, 0x10, 0x05,
	0x12, 0x15, 0x0a, 0x11, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x44,
	0x45, 0x52, 0x5f, 0x31, 0x38, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x47, 0x45, 0x5f, 0x42,
	0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x31, 0x39, 0x10, 0x07, 0x12, 0x15,
	0x0a, 0x11, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52,
	0x5f, 0x32, 0x30, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e,
	0x44, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x32, 0x31, 0x10, 0x09, 0x12, 0x15, 0x0a, 0x11,
	0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x32,
	0x32, 0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f,
	0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x32, 0x33, 0x10, 0x0b, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x47,
	0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x32, 0x34, 0x10,
	0x0c, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e,
	0x44, 0x45, 0x52, 0x5f, 0x32, 0x35, 0x10, 0x0d, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x47, 0x45, 0x5f,
	0x42, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x31, 0x33, 0x10, 0x0e, 0x12, 0x14,
	0x0a, 0x10, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f,
	0x31, 0x34, 0x10, 0x0f, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44,
	0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x31, 0x35, 0x10, 0x10, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x47,
	0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x31, 0x36, 0x10, 0x11,
	0x12, 0x14, 0x0a, 0x10, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x56, 0x45,
	0x52, 0x5f, 0x31, 0x37, 0x10, 0x12, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41,
	0x4e, 0x44, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x31, 0x38, 0x10, 0x13, 0x12, 0x14, 0x0a, 0x10,
	0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x31, 0x39,
	0x10, 0x14, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x4f,
	0x56, 0x45, 0x52, 0x5f, 0x32, 0x30, 0x10, 0x15, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x47, 0x45, 0x5f,
	0x42, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x32, 0x31, 0x10, 0x16, 0x12, 0x14,
	0x0a, 0x10, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f,
	0x32, 0x32, 0x10, 0x17, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44,
	0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x32, 0x33, 0x10, 0x18, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x47,
	0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x32, 0x34, 0x10, 0x19,
	0x12, 0x14, 0x0a, 0x10, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x56, 0x45,
	0x52, 0x5f, 0x32, 0x35, 0x10, 0x1a, 0x42, 0x5b, 0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x41, 0x67, 0x65, 0x42, 0x61, 0x6e, 0x64,
	0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x21, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x62, 0x61, 0x6e, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d,
	0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_enums_ageband_v1_age_band_proto_rawDescOnce sync.Once
	file_micashared_common_enums_ageband_v1_age_band_proto_rawDescData = file_micashared_common_enums_ageband_v1_age_band_proto_rawDesc
)

func file_micashared_common_enums_ageband_v1_age_band_proto_rawDescGZIP() []byte {
	file_micashared_common_enums_ageband_v1_age_band_proto_rawDescOnce.Do(func() {
		file_micashared_common_enums_ageband_v1_age_band_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_enums_ageband_v1_age_band_proto_rawDescData)
	})
	return file_micashared_common_enums_ageband_v1_age_band_proto_rawDescData
}

var file_micashared_common_enums_ageband_v1_age_band_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_micashared_common_enums_ageband_v1_age_band_proto_goTypes = []interface{}{
	(AgeBand)(0), // 0: micashared.common.enums.ageband.v1.AgeBand
}
var file_micashared_common_enums_ageband_v1_age_band_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_micashared_common_enums_ageband_v1_age_band_proto_init() }
func file_micashared_common_enums_ageband_v1_age_band_proto_init() {
	if File_micashared_common_enums_ageband_v1_age_band_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_micashared_common_enums_ageband_v1_age_band_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_enums_ageband_v1_age_band_proto_goTypes,
		DependencyIndexes: file_micashared_common_enums_ageband_v1_age_band_proto_depIdxs,
		EnumInfos:         file_micashared_common_enums_ageband_v1_age_band_proto_enumTypes,
	}.Build()
	File_micashared_common_enums_ageband_v1_age_band_proto = out.File
	file_micashared_common_enums_ageband_v1_age_band_proto_rawDesc = nil
	file_micashared_common_enums_ageband_v1_age_band_proto_goTypes = nil
	file_micashared_common_enums_ageband_v1_age_band_proto_depIdxs = nil
}
