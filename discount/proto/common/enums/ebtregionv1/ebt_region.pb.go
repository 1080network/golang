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
// 	protoc        v3.19.4
// source: common/enums/ebtregion/v1/ebt_region.proto

package ebtregionv1

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

type EbtRegion int32

const (
	EbtRegion_EBT_REGION_UNSPECIFIED EbtRegion = 0
	//ISO 3166-2:US
	EbtRegion_EBT_REGION_US_AL EbtRegion = 1  // Alabama (state)
	EbtRegion_EBT_REGION_US_AK EbtRegion = 2  // Alaska (state)
	EbtRegion_EBT_REGION_US_AZ EbtRegion = 3  // Arizona (state)
	EbtRegion_EBT_REGION_US_AR EbtRegion = 4  // Arkansas (state)
	EbtRegion_EBT_REGION_US_CA EbtRegion = 5  // California (state)
	EbtRegion_EBT_REGION_US_CO EbtRegion = 6  // Colorado (state)
	EbtRegion_EBT_REGION_US_CT EbtRegion = 7  // Connecticut (state)
	EbtRegion_EBT_REGION_US_DE EbtRegion = 8  // Delaware (state)
	EbtRegion_EBT_REGION_US_FL EbtRegion = 9  // Florida (state)
	EbtRegion_EBT_REGION_US_GA EbtRegion = 10 // Georgia (state)
	EbtRegion_EBT_REGION_US_HI EbtRegion = 11 // Hawaii (state)
	EbtRegion_EBT_REGION_US_ID EbtRegion = 12 // Idaho (state)
	EbtRegion_EBT_REGION_US_IL EbtRegion = 13 // Illinois (state)
	EbtRegion_EBT_REGION_US_IN EbtRegion = 14 // Indiana (state)
	EbtRegion_EBT_REGION_US_IA EbtRegion = 15 // Iowa (state)
	EbtRegion_EBT_REGION_US_KS EbtRegion = 16 // Kansas (state)
	EbtRegion_EBT_REGION_US_KY EbtRegion = 17 // Kentucky (state)
	EbtRegion_EBT_REGION_US_LA EbtRegion = 18 // Louisiana (state)
	EbtRegion_EBT_REGION_US_ME EbtRegion = 19 // Maine (state)
	EbtRegion_EBT_REGION_US_MD EbtRegion = 20 // Maryland (state)
	EbtRegion_EBT_REGION_US_MA EbtRegion = 21 // Massachusetts (state)
	EbtRegion_EBT_REGION_US_MI EbtRegion = 22 // Michigan (state)
	EbtRegion_EBT_REGION_US_MN EbtRegion = 23 // Minnesota (state)
	EbtRegion_EBT_REGION_US_MS EbtRegion = 24 // Mississippi (state)
	EbtRegion_EBT_REGION_US_MO EbtRegion = 25 // Missouri (state)
	EbtRegion_EBT_REGION_US_MT EbtRegion = 26 // Montana (state)
	EbtRegion_EBT_REGION_US_NE EbtRegion = 27 // Nebraska (state)
	EbtRegion_EBT_REGION_US_NV EbtRegion = 28 // Nevada (state)
	EbtRegion_EBT_REGION_US_NH EbtRegion = 29 // New Hampshire (state)
	EbtRegion_EBT_REGION_US_NJ EbtRegion = 30 // New Jersey (state)
	EbtRegion_EBT_REGION_US_NM EbtRegion = 31 // New Mexico (state)
	EbtRegion_EBT_REGION_US_NY EbtRegion = 32 // New York (state)
	EbtRegion_EBT_REGION_US_NC EbtRegion = 33 // North Carolina (state)
	EbtRegion_EBT_REGION_US_ND EbtRegion = 34 // North Dakota (state)
	EbtRegion_EBT_REGION_US_OH EbtRegion = 35 // Ohio (state)
	EbtRegion_EBT_REGION_US_OK EbtRegion = 36 // Oklahoma (state)
	EbtRegion_EBT_REGION_US_OR EbtRegion = 37 // Oregon (state)
	EbtRegion_EBT_REGION_US_PA EbtRegion = 38 // Pennsylvania (state)
	EbtRegion_EBT_REGION_US_RI EbtRegion = 39 // Rhode Island (state)
	EbtRegion_EBT_REGION_US_SC EbtRegion = 40 // South Carolina (state)
	EbtRegion_EBT_REGION_US_SD EbtRegion = 41 // South Dakota (state)
	EbtRegion_EBT_REGION_US_TN EbtRegion = 42 // Tennessee (state)
	EbtRegion_EBT_REGION_US_TX EbtRegion = 43 // Texas (state)
	EbtRegion_EBT_REGION_US_UT EbtRegion = 44 // Utah (state)
	EbtRegion_EBT_REGION_US_VT EbtRegion = 45 // Vermont (state)
	EbtRegion_EBT_REGION_US_VA EbtRegion = 46 // Virginia (state)
	EbtRegion_EBT_REGION_US_WA EbtRegion = 47 // Washington (state)
	EbtRegion_EBT_REGION_US_WV EbtRegion = 48 // West Virginia (state)
	EbtRegion_EBT_REGION_US_WI EbtRegion = 49 // Wisconsin (state)
	EbtRegion_EBT_REGION_US_WY EbtRegion = 50 // Wyoming (state)
	EbtRegion_EBT_REGION_US_DC EbtRegion = 51 // District of Columbia (district)
	EbtRegion_EBT_REGION_US_AS EbtRegion = 52 // American Samoa
	EbtRegion_EBT_REGION_US_GU EbtRegion = 53 // Guam
	EbtRegion_EBT_REGION_US_MP EbtRegion = 54 // Northern Mariana Islands
	EbtRegion_EBT_REGION_US_PR EbtRegion = 55 // Puerto Rico
	EbtRegion_EBT_REGION_US_UM EbtRegion = 56 // United States Minor Outlying Islands
	EbtRegion_EBT_REGION_US_VI EbtRegion = 57 // U.S. Virgin Islands
)

// Enum value maps for EbtRegion.
var (
	EbtRegion_name = map[int32]string{
		0:  "EBT_REGION_UNSPECIFIED",
		1:  "EBT_REGION_US_AL",
		2:  "EBT_REGION_US_AK",
		3:  "EBT_REGION_US_AZ",
		4:  "EBT_REGION_US_AR",
		5:  "EBT_REGION_US_CA",
		6:  "EBT_REGION_US_CO",
		7:  "EBT_REGION_US_CT",
		8:  "EBT_REGION_US_DE",
		9:  "EBT_REGION_US_FL",
		10: "EBT_REGION_US_GA",
		11: "EBT_REGION_US_HI",
		12: "EBT_REGION_US_ID",
		13: "EBT_REGION_US_IL",
		14: "EBT_REGION_US_IN",
		15: "EBT_REGION_US_IA",
		16: "EBT_REGION_US_KS",
		17: "EBT_REGION_US_KY",
		18: "EBT_REGION_US_LA",
		19: "EBT_REGION_US_ME",
		20: "EBT_REGION_US_MD",
		21: "EBT_REGION_US_MA",
		22: "EBT_REGION_US_MI",
		23: "EBT_REGION_US_MN",
		24: "EBT_REGION_US_MS",
		25: "EBT_REGION_US_MO",
		26: "EBT_REGION_US_MT",
		27: "EBT_REGION_US_NE",
		28: "EBT_REGION_US_NV",
		29: "EBT_REGION_US_NH",
		30: "EBT_REGION_US_NJ",
		31: "EBT_REGION_US_NM",
		32: "EBT_REGION_US_NY",
		33: "EBT_REGION_US_NC",
		34: "EBT_REGION_US_ND",
		35: "EBT_REGION_US_OH",
		36: "EBT_REGION_US_OK",
		37: "EBT_REGION_US_OR",
		38: "EBT_REGION_US_PA",
		39: "EBT_REGION_US_RI",
		40: "EBT_REGION_US_SC",
		41: "EBT_REGION_US_SD",
		42: "EBT_REGION_US_TN",
		43: "EBT_REGION_US_TX",
		44: "EBT_REGION_US_UT",
		45: "EBT_REGION_US_VT",
		46: "EBT_REGION_US_VA",
		47: "EBT_REGION_US_WA",
		48: "EBT_REGION_US_WV",
		49: "EBT_REGION_US_WI",
		50: "EBT_REGION_US_WY",
		51: "EBT_REGION_US_DC",
		52: "EBT_REGION_US_AS",
		53: "EBT_REGION_US_GU",
		54: "EBT_REGION_US_MP",
		55: "EBT_REGION_US_PR",
		56: "EBT_REGION_US_UM",
		57: "EBT_REGION_US_VI",
	}
	EbtRegion_value = map[string]int32{
		"EBT_REGION_UNSPECIFIED": 0,
		"EBT_REGION_US_AL":       1,
		"EBT_REGION_US_AK":       2,
		"EBT_REGION_US_AZ":       3,
		"EBT_REGION_US_AR":       4,
		"EBT_REGION_US_CA":       5,
		"EBT_REGION_US_CO":       6,
		"EBT_REGION_US_CT":       7,
		"EBT_REGION_US_DE":       8,
		"EBT_REGION_US_FL":       9,
		"EBT_REGION_US_GA":       10,
		"EBT_REGION_US_HI":       11,
		"EBT_REGION_US_ID":       12,
		"EBT_REGION_US_IL":       13,
		"EBT_REGION_US_IN":       14,
		"EBT_REGION_US_IA":       15,
		"EBT_REGION_US_KS":       16,
		"EBT_REGION_US_KY":       17,
		"EBT_REGION_US_LA":       18,
		"EBT_REGION_US_ME":       19,
		"EBT_REGION_US_MD":       20,
		"EBT_REGION_US_MA":       21,
		"EBT_REGION_US_MI":       22,
		"EBT_REGION_US_MN":       23,
		"EBT_REGION_US_MS":       24,
		"EBT_REGION_US_MO":       25,
		"EBT_REGION_US_MT":       26,
		"EBT_REGION_US_NE":       27,
		"EBT_REGION_US_NV":       28,
		"EBT_REGION_US_NH":       29,
		"EBT_REGION_US_NJ":       30,
		"EBT_REGION_US_NM":       31,
		"EBT_REGION_US_NY":       32,
		"EBT_REGION_US_NC":       33,
		"EBT_REGION_US_ND":       34,
		"EBT_REGION_US_OH":       35,
		"EBT_REGION_US_OK":       36,
		"EBT_REGION_US_OR":       37,
		"EBT_REGION_US_PA":       38,
		"EBT_REGION_US_RI":       39,
		"EBT_REGION_US_SC":       40,
		"EBT_REGION_US_SD":       41,
		"EBT_REGION_US_TN":       42,
		"EBT_REGION_US_TX":       43,
		"EBT_REGION_US_UT":       44,
		"EBT_REGION_US_VT":       45,
		"EBT_REGION_US_VA":       46,
		"EBT_REGION_US_WA":       47,
		"EBT_REGION_US_WV":       48,
		"EBT_REGION_US_WI":       49,
		"EBT_REGION_US_WY":       50,
		"EBT_REGION_US_DC":       51,
		"EBT_REGION_US_AS":       52,
		"EBT_REGION_US_GU":       53,
		"EBT_REGION_US_MP":       54,
		"EBT_REGION_US_PR":       55,
		"EBT_REGION_US_UM":       56,
		"EBT_REGION_US_VI":       57,
	}
)

func (x EbtRegion) Enum() *EbtRegion {
	p := new(EbtRegion)
	*p = x
	return p
}

func (x EbtRegion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EbtRegion) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enums_ebtregion_v1_ebt_region_proto_enumTypes[0].Descriptor()
}

func (EbtRegion) Type() protoreflect.EnumType {
	return &file_common_enums_ebtregion_v1_ebt_region_proto_enumTypes[0]
}

func (x EbtRegion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EbtRegion.Descriptor instead.
func (EbtRegion) EnumDescriptor() ([]byte, []int) {
	return file_common_enums_ebtregion_v1_ebt_region_proto_rawDescGZIP(), []int{0}
}

var File_common_enums_ebtregion_v1_ebt_region_proto protoreflect.FileDescriptor

var file_common_enums_ebtregion_v1_ebt_region_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65,
	0x62, 0x74, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x62, 0x74, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x65, 0x62, 0x74, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2a, 0x8d, 0x0a, 0x0a, 0x09, 0x45, 0x62, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x53, 0x5f, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52,
	0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x41, 0x4b, 0x10, 0x02, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x41,
	0x5a, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x41, 0x52, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54,
	0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x10, 0x05, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53,
	0x5f, 0x43, 0x4f, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x43, 0x54, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x10,
	0x08, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x53, 0x5f, 0x46, 0x4c, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52,
	0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x47, 0x41, 0x10, 0x0a, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x48,
	0x49, 0x10, 0x0b, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x49, 0x44, 0x10, 0x0c, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54,
	0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x49, 0x4c, 0x10, 0x0d, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53,
	0x5f, 0x49, 0x4e, 0x10, 0x0e, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x49, 0x41, 0x10, 0x0f, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4b, 0x53, 0x10,
	0x10, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x53, 0x5f, 0x4b, 0x59, 0x10, 0x11, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52,
	0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4c, 0x41, 0x10, 0x12, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4d,
	0x45, 0x10, 0x13, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4d, 0x44, 0x10, 0x14, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54,
	0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4d, 0x41, 0x10, 0x15, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53,
	0x5f, 0x4d, 0x49, 0x10, 0x16, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4d, 0x4e, 0x10, 0x17, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4d, 0x53, 0x10,
	0x18, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x53, 0x5f, 0x4d, 0x4f, 0x10, 0x19, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52,
	0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4d, 0x54, 0x10, 0x1a, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4e,
	0x45, 0x10, 0x1b, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4e, 0x56, 0x10, 0x1c, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54,
	0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4e, 0x48, 0x10, 0x1d, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53,
	0x5f, 0x4e, 0x4a, 0x10, 0x1e, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4e, 0x4d, 0x10, 0x1f, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4e, 0x59, 0x10,
	0x20, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x53, 0x5f, 0x4e, 0x43, 0x10, 0x21, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52,
	0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4e, 0x44, 0x10, 0x22, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4f,
	0x48, 0x10, 0x23, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x24, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54,
	0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4f, 0x52, 0x10, 0x25, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53,
	0x5f, 0x50, 0x41, 0x10, 0x26, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x52, 0x49, 0x10, 0x27, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x53, 0x43, 0x10,
	0x28, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x53, 0x5f, 0x53, 0x44, 0x10, 0x29, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52,
	0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x54, 0x4e, 0x10, 0x2a, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x54,
	0x58, 0x10, 0x2b, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x55, 0x54, 0x10, 0x2c, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54,
	0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x56, 0x54, 0x10, 0x2d, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53,
	0x5f, 0x56, 0x41, 0x10, 0x2e, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x57, 0x41, 0x10, 0x2f, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x57, 0x56, 0x10,
	0x30, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x53, 0x5f, 0x57, 0x49, 0x10, 0x31, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52,
	0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x57, 0x59, 0x10, 0x32, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x44,
	0x43, 0x10, 0x33, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x41, 0x53, 0x10, 0x34, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54,
	0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x47, 0x55, 0x10, 0x35, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53,
	0x5f, 0x4d, 0x50, 0x10, 0x36, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x50, 0x52, 0x10, 0x37, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x55, 0x4d, 0x10,
	0x38, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x53, 0x5f, 0x56, 0x49, 0x10, 0x39, 0x42, 0x59, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x12, 0x45, 0x62, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75,
	0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x23, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x65, 0x62, 0x74, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49,
	0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_enums_ebtregion_v1_ebt_region_proto_rawDescOnce sync.Once
	file_common_enums_ebtregion_v1_ebt_region_proto_rawDescData = file_common_enums_ebtregion_v1_ebt_region_proto_rawDesc
)

func file_common_enums_ebtregion_v1_ebt_region_proto_rawDescGZIP() []byte {
	file_common_enums_ebtregion_v1_ebt_region_proto_rawDescOnce.Do(func() {
		file_common_enums_ebtregion_v1_ebt_region_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_enums_ebtregion_v1_ebt_region_proto_rawDescData)
	})
	return file_common_enums_ebtregion_v1_ebt_region_proto_rawDescData
}

var file_common_enums_ebtregion_v1_ebt_region_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_enums_ebtregion_v1_ebt_region_proto_goTypes = []interface{}{
	(EbtRegion)(0), // 0: common.enums.ebtregion.v1.EbtRegion
}
var file_common_enums_ebtregion_v1_ebt_region_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_enums_ebtregion_v1_ebt_region_proto_init() }
func file_common_enums_ebtregion_v1_ebt_region_proto_init() {
	if File_common_enums_ebtregion_v1_ebt_region_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_enums_ebtregion_v1_ebt_region_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_enums_ebtregion_v1_ebt_region_proto_goTypes,
		DependencyIndexes: file_common_enums_ebtregion_v1_ebt_region_proto_depIdxs,
		EnumInfos:         file_common_enums_ebtregion_v1_ebt_region_proto_enumTypes,
	}.Build()
	File_common_enums_ebtregion_v1_ebt_region_proto = out.File
	file_common_enums_ebtregion_v1_ebt_region_proto_rawDesc = nil
	file_common_enums_ebtregion_v1_ebt_region_proto_goTypes = nil
	file_common_enums_ebtregion_v1_ebt_region_proto_depIdxs = nil
}
