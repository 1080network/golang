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
// source: common/enums/region/v1/region.proto

package regionv1

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

type Region int32

const (
	Region_REGION_UNSPECIFIED Region = 0
	//ISO 3166-2:US
	Region_REGION_US_AL Region = 1  // Alabama (state)
	Region_REGION_US_AK Region = 2  // Alaska (state)
	Region_REGION_US_AZ Region = 3  // Arizona (state)
	Region_REGION_US_AR Region = 4  // Arkansas (state)
	Region_REGION_US_CA Region = 5  // California (state)
	Region_REGION_US_CO Region = 6  // Colorado (state)
	Region_REGION_US_CT Region = 7  // Connecticut (state)
	Region_REGION_US_DE Region = 8  // Delaware (state)
	Region_REGION_US_FL Region = 9  // Florida (state)
	Region_REGION_US_GA Region = 10 // Georgia (state)
	Region_REGION_US_HI Region = 11 // Hawaii (state)
	Region_REGION_US_ID Region = 12 // Idaho (state)
	Region_REGION_US_IL Region = 13 // Illinois (state)
	Region_REGION_US_IN Region = 14 // Indiana (state)
	Region_REGION_US_IA Region = 15 // Iowa (state)
	Region_REGION_US_KS Region = 16 // Kansas (state)
	Region_REGION_US_KY Region = 17 // Kentucky (state)
	Region_REGION_US_LA Region = 18 // Louisiana (state)
	Region_REGION_US_ME Region = 19 // Maine (state)
	Region_REGION_US_MD Region = 20 // Maryland (state)
	Region_REGION_US_MA Region = 21 // Massachusetts (state)
	Region_REGION_US_MI Region = 22 // Michigan (state)
	Region_REGION_US_MN Region = 23 // Minnesota (state)
	Region_REGION_US_MS Region = 24 // Mississippi (state)
	Region_REGION_US_MO Region = 25 // Missouri (state)
	Region_REGION_US_MT Region = 26 // Montana (state)
	Region_REGION_US_NE Region = 27 // Nebraska (state)
	Region_REGION_US_NV Region = 28 // Nevada (state)
	Region_REGION_US_NH Region = 29 // New Hampshire (state)
	Region_REGION_US_NJ Region = 30 // New Jersey (state)
	Region_REGION_US_NM Region = 31 // New Mexico (state)
	Region_REGION_US_NY Region = 32 // New York (state)
	Region_REGION_US_NC Region = 33 // North Carolina (state)
	Region_REGION_US_ND Region = 34 // North Dakota (state)
	Region_REGION_US_OH Region = 35 // Ohio (state)
	Region_REGION_US_OK Region = 36 // Oklahoma (state)
	Region_REGION_US_OR Region = 37 // Oregon (state)
	Region_REGION_US_PA Region = 38 // Pennsylvania (state)
	Region_REGION_US_RI Region = 39 // Rhode Island (state)
	Region_REGION_US_SC Region = 40 // South Carolina (state)
	Region_REGION_US_SD Region = 41 // South Dakota (state)
	Region_REGION_US_TN Region = 42 // Tennessee (state)
	Region_REGION_US_TX Region = 43 // Texas (state)
	Region_REGION_US_UT Region = 44 // Utah (state)
	Region_REGION_US_VT Region = 45 // Vermont (state)
	Region_REGION_US_VA Region = 46 // Virginia (state)
	Region_REGION_US_WA Region = 47 // Washington (state)
	Region_REGION_US_WV Region = 48 // West Virginia (state)
	Region_REGION_US_WI Region = 49 // Wisconsin (state)
	Region_REGION_US_WY Region = 50 // Wyoming (state)
	Region_REGION_US_DC Region = 51 // District of Columbia (district)
	Region_REGION_US_AS Region = 52 // American Samoa
	Region_REGION_US_GU Region = 53 // Guam
	Region_REGION_US_MP Region = 54 // Northern Mariana Islands
	Region_REGION_US_PR Region = 55 // Puerto Rico
	Region_REGION_US_UM Region = 56 // United States Minor Outlying Islands
	Region_REGION_US_VI Region = 57 // U.S. Virgin Islands
)

// Enum value maps for Region.
var (
	Region_name = map[int32]string{
		0:  "REGION_UNSPECIFIED",
		1:  "REGION_US_AL",
		2:  "REGION_US_AK",
		3:  "REGION_US_AZ",
		4:  "REGION_US_AR",
		5:  "REGION_US_CA",
		6:  "REGION_US_CO",
		7:  "REGION_US_CT",
		8:  "REGION_US_DE",
		9:  "REGION_US_FL",
		10: "REGION_US_GA",
		11: "REGION_US_HI",
		12: "REGION_US_ID",
		13: "REGION_US_IL",
		14: "REGION_US_IN",
		15: "REGION_US_IA",
		16: "REGION_US_KS",
		17: "REGION_US_KY",
		18: "REGION_US_LA",
		19: "REGION_US_ME",
		20: "REGION_US_MD",
		21: "REGION_US_MA",
		22: "REGION_US_MI",
		23: "REGION_US_MN",
		24: "REGION_US_MS",
		25: "REGION_US_MO",
		26: "REGION_US_MT",
		27: "REGION_US_NE",
		28: "REGION_US_NV",
		29: "REGION_US_NH",
		30: "REGION_US_NJ",
		31: "REGION_US_NM",
		32: "REGION_US_NY",
		33: "REGION_US_NC",
		34: "REGION_US_ND",
		35: "REGION_US_OH",
		36: "REGION_US_OK",
		37: "REGION_US_OR",
		38: "REGION_US_PA",
		39: "REGION_US_RI",
		40: "REGION_US_SC",
		41: "REGION_US_SD",
		42: "REGION_US_TN",
		43: "REGION_US_TX",
		44: "REGION_US_UT",
		45: "REGION_US_VT",
		46: "REGION_US_VA",
		47: "REGION_US_WA",
		48: "REGION_US_WV",
		49: "REGION_US_WI",
		50: "REGION_US_WY",
		51: "REGION_US_DC",
		52: "REGION_US_AS",
		53: "REGION_US_GU",
		54: "REGION_US_MP",
		55: "REGION_US_PR",
		56: "REGION_US_UM",
		57: "REGION_US_VI",
	}
	Region_value = map[string]int32{
		"REGION_UNSPECIFIED": 0,
		"REGION_US_AL":       1,
		"REGION_US_AK":       2,
		"REGION_US_AZ":       3,
		"REGION_US_AR":       4,
		"REGION_US_CA":       5,
		"REGION_US_CO":       6,
		"REGION_US_CT":       7,
		"REGION_US_DE":       8,
		"REGION_US_FL":       9,
		"REGION_US_GA":       10,
		"REGION_US_HI":       11,
		"REGION_US_ID":       12,
		"REGION_US_IL":       13,
		"REGION_US_IN":       14,
		"REGION_US_IA":       15,
		"REGION_US_KS":       16,
		"REGION_US_KY":       17,
		"REGION_US_LA":       18,
		"REGION_US_ME":       19,
		"REGION_US_MD":       20,
		"REGION_US_MA":       21,
		"REGION_US_MI":       22,
		"REGION_US_MN":       23,
		"REGION_US_MS":       24,
		"REGION_US_MO":       25,
		"REGION_US_MT":       26,
		"REGION_US_NE":       27,
		"REGION_US_NV":       28,
		"REGION_US_NH":       29,
		"REGION_US_NJ":       30,
		"REGION_US_NM":       31,
		"REGION_US_NY":       32,
		"REGION_US_NC":       33,
		"REGION_US_ND":       34,
		"REGION_US_OH":       35,
		"REGION_US_OK":       36,
		"REGION_US_OR":       37,
		"REGION_US_PA":       38,
		"REGION_US_RI":       39,
		"REGION_US_SC":       40,
		"REGION_US_SD":       41,
		"REGION_US_TN":       42,
		"REGION_US_TX":       43,
		"REGION_US_UT":       44,
		"REGION_US_VT":       45,
		"REGION_US_VA":       46,
		"REGION_US_WA":       47,
		"REGION_US_WV":       48,
		"REGION_US_WI":       49,
		"REGION_US_WY":       50,
		"REGION_US_DC":       51,
		"REGION_US_AS":       52,
		"REGION_US_GU":       53,
		"REGION_US_MP":       54,
		"REGION_US_PR":       55,
		"REGION_US_UM":       56,
		"REGION_US_VI":       57,
	}
)

func (x Region) Enum() *Region {
	p := new(Region)
	*p = x
	return p
}

func (x Region) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Region) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enums_region_v1_region_proto_enumTypes[0].Descriptor()
}

func (Region) Type() protoreflect.EnumType {
	return &file_common_enums_region_v1_region_proto_enumTypes[0]
}

func (x Region) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Region.Descriptor instead.
func (Region) EnumDescriptor() ([]byte, []int) {
	return file_common_enums_region_v1_region_proto_rawDescGZIP(), []int{0}
}

var File_common_enums_region_v1_region_proto protoreflect.FileDescriptor

var file_common_enums_region_v1_region_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2a, 0xa2, 0x08,
	0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x47, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x41, 0x4c,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f,
	0x41, 0x4b, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x53, 0x5f, 0x41, 0x5a, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x53, 0x5f, 0x41, 0x52, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45,
	0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c,
	0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x43, 0x54, 0x10, 0x07, 0x12, 0x10,
	0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x10, 0x08,
	0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x46, 0x4c,
	0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f,
	0x47, 0x41, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x53, 0x5f, 0x48, 0x49, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x53, 0x5f, 0x49, 0x44, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x49, 0x4c, 0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45,
	0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x10, 0x0e, 0x12, 0x10, 0x0a, 0x0c,
	0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x49, 0x41, 0x10, 0x0f, 0x12, 0x10,
	0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4b, 0x53, 0x10, 0x10,
	0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4b, 0x59,
	0x10, 0x11, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f,
	0x4c, 0x41, 0x10, 0x12, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x53, 0x5f, 0x4d, 0x45, 0x10, 0x13, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x53, 0x5f, 0x4d, 0x44, 0x10, 0x14, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4d, 0x41, 0x10, 0x15, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45,
	0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4d, 0x49, 0x10, 0x16, 0x12, 0x10, 0x0a, 0x0c,
	0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4d, 0x4e, 0x10, 0x17, 0x12, 0x10,
	0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4d, 0x53, 0x10, 0x18,
	0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4d, 0x4f,
	0x10, 0x19, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f,
	0x4d, 0x54, 0x10, 0x1a, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x53, 0x5f, 0x4e, 0x45, 0x10, 0x1b, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x53, 0x5f, 0x4e, 0x56, 0x10, 0x1c, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4e, 0x48, 0x10, 0x1d, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45,
	0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4e, 0x4a, 0x10, 0x1e, 0x12, 0x10, 0x0a, 0x0c,
	0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4e, 0x4d, 0x10, 0x1f, 0x12, 0x10,
	0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4e, 0x59, 0x10, 0x20,
	0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4e, 0x43,
	0x10, 0x21, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f,
	0x4e, 0x44, 0x10, 0x22, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x53, 0x5f, 0x4f, 0x48, 0x10, 0x23, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x24, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4f, 0x52, 0x10, 0x25, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45,
	0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x10, 0x26, 0x12, 0x10, 0x0a, 0x0c,
	0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x52, 0x49, 0x10, 0x27, 0x12, 0x10,
	0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x53, 0x43, 0x10, 0x28,
	0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x53, 0x44,
	0x10, 0x29, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f,
	0x54, 0x4e, 0x10, 0x2a, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x53, 0x5f, 0x54, 0x58, 0x10, 0x2b, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x53, 0x5f, 0x55, 0x54, 0x10, 0x2c, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x56, 0x54, 0x10, 0x2d, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45,
	0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x56, 0x41, 0x10, 0x2e, 0x12, 0x10, 0x0a, 0x0c,
	0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x57, 0x41, 0x10, 0x2f, 0x12, 0x10,
	0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x57, 0x56, 0x10, 0x30,
	0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x57, 0x49,
	0x10, 0x31, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f,
	0x57, 0x59, 0x10, 0x32, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x53, 0x5f, 0x44, 0x43, 0x10, 0x33, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x53, 0x5f, 0x41, 0x53, 0x10, 0x34, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x47, 0x55, 0x10, 0x35, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45,
	0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x4d, 0x50, 0x10, 0x36, 0x12, 0x10, 0x0a, 0x0c,
	0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x50, 0x52, 0x10, 0x37, 0x12, 0x10,
	0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x55, 0x4d, 0x10, 0x38,
	0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x5f, 0x56, 0x49,
	0x10, 0x39, 0x42, 0x53, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x20,
	0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x76, 0x31,
	0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_enums_region_v1_region_proto_rawDescOnce sync.Once
	file_common_enums_region_v1_region_proto_rawDescData = file_common_enums_region_v1_region_proto_rawDesc
)

func file_common_enums_region_v1_region_proto_rawDescGZIP() []byte {
	file_common_enums_region_v1_region_proto_rawDescOnce.Do(func() {
		file_common_enums_region_v1_region_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_enums_region_v1_region_proto_rawDescData)
	})
	return file_common_enums_region_v1_region_proto_rawDescData
}

var file_common_enums_region_v1_region_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_enums_region_v1_region_proto_goTypes = []interface{}{
	(Region)(0), // 0: common.enums.region.v1.Region
}
var file_common_enums_region_v1_region_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_enums_region_v1_region_proto_init() }
func file_common_enums_region_v1_region_proto_init() {
	if File_common_enums_region_v1_region_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_enums_region_v1_region_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_enums_region_v1_region_proto_goTypes,
		DependencyIndexes: file_common_enums_region_v1_region_proto_depIdxs,
		EnumInfos:         file_common_enums_region_v1_region_proto_enumTypes,
	}.Build()
	File_common_enums_region_v1_region_proto = out.File
	file_common_enums_region_v1_region_proto_rawDesc = nil
	file_common_enums_region_v1_region_proto_goTypes = nil
	file_common_enums_region_v1_region_proto_depIdxs = nil
}
