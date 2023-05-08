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
// source: micashared/common/enums/organizationcategory/v1/organization_category.proto

package organizationcategoryv1

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

type OrganizationCategory int32

const (
	OrganizationCategory_ORGANIZATION_CATEGORY_UNSPECIFIED                 OrganizationCategory = 0
	OrganizationCategory_ORGANIZATION_CATEGORY_ADULT_SERVICES              OrganizationCategory = 1000
	OrganizationCategory_ORGANIZATION_CATEGORY_AIRLINE                     OrganizationCategory = 1010
	OrganizationCategory_ORGANIZATION_CATEGORY_ANTIQUES_AND_ARTWORK        OrganizationCategory = 1020
	OrganizationCategory_ORGANIZATION_CATEGORY_APPAREL                     OrganizationCategory = 1030
	OrganizationCategory_ORGANIZATION_CATEGORY_AUTOMOTIVE_REPAIR           OrganizationCategory = 1040
	OrganizationCategory_ORGANIZATION_CATEGORY_AUTOMOTIVE_SALES            OrganizationCategory = 1050
	OrganizationCategory_ORGANIZATION_CATEGORY_BOOKSTORE                   OrganizationCategory = 1060
	OrganizationCategory_ORGANIZATION_CATEGORY_CAR_RENTAL                  OrganizationCategory = 1070
	OrganizationCategory_ORGANIZATION_CATEGORY_CONSUMER_ELECTRONICS        OrganizationCategory = 1080
	OrganizationCategory_ORGANIZATION_CATEGORY_CONSUMER_ELECTRONICS_REPAIR OrganizationCategory = 1090
	OrganizationCategory_ORGANIZATION_CATEGORY_CONVENIENCE_STORE           OrganizationCategory = 1100
	OrganizationCategory_ORGANIZATION_CATEGORY_COSMETICS                   OrganizationCategory = 1110
	OrganizationCategory_ORGANIZATION_CATEGORY_CRUISE                      OrganizationCategory = 1120
	OrganizationCategory_ORGANIZATION_CATEGORY_FAST_CASUAL_DINING          OrganizationCategory = 1130
	OrganizationCategory_ORGANIZATION_CATEGORY_FINANCIAL_SERVICES          OrganizationCategory = 1140
	OrganizationCategory_ORGANIZATION_CATEGORY_FUEL                        OrganizationCategory = 1150
	OrganizationCategory_ORGANIZATION_CATEGORY_FURNITURE                   OrganizationCategory = 1160
	OrganizationCategory_ORGANIZATION_CATEGORY_GROCERY                     OrganizationCategory = 1170
	OrganizationCategory_ORGANIZATION_CATEGORY_HEALTH_CARE                 OrganizationCategory = 1180
	OrganizationCategory_ORGANIZATION_CATEGORY_HOME_IMPROVEMENT            OrganizationCategory = 1190
	OrganizationCategory_ORGANIZATION_CATEGORY_HOTEL                       OrganizationCategory = 1200
	OrganizationCategory_ORGANIZATION_CATEGORY_MONEY_TRANSMISSION          OrganizationCategory = 1210
	OrganizationCategory_ORGANIZATION_CATEGORY_PHARMACY                    OrganizationCategory = 1220
	OrganizationCategory_ORGANIZATION_CATEGORY_PROFESSIONAL_SERVICES       OrganizationCategory = 1230
	OrganizationCategory_ORGANIZATION_CATEGORY_QUICK_SERVICE_RESTAURANT    OrganizationCategory = 1240
	OrganizationCategory_ORGANIZATION_CATEGORY_RESTAURANT                  OrganizationCategory = 1250
	OrganizationCategory_ORGANIZATION_CATEGORY_STREAMING_MEDIA             OrganizationCategory = 1260
	OrganizationCategory_ORGANIZATION_CATEGORY_TELEPHONY_AND_INTERNET      OrganizationCategory = 1270
)

// Enum value maps for OrganizationCategory.
var (
	OrganizationCategory_name = map[int32]string{
		0:    "ORGANIZATION_CATEGORY_UNSPECIFIED",
		1000: "ORGANIZATION_CATEGORY_ADULT_SERVICES",
		1010: "ORGANIZATION_CATEGORY_AIRLINE",
		1020: "ORGANIZATION_CATEGORY_ANTIQUES_AND_ARTWORK",
		1030: "ORGANIZATION_CATEGORY_APPAREL",
		1040: "ORGANIZATION_CATEGORY_AUTOMOTIVE_REPAIR",
		1050: "ORGANIZATION_CATEGORY_AUTOMOTIVE_SALES",
		1060: "ORGANIZATION_CATEGORY_BOOKSTORE",
		1070: "ORGANIZATION_CATEGORY_CAR_RENTAL",
		1080: "ORGANIZATION_CATEGORY_CONSUMER_ELECTRONICS",
		1090: "ORGANIZATION_CATEGORY_CONSUMER_ELECTRONICS_REPAIR",
		1100: "ORGANIZATION_CATEGORY_CONVENIENCE_STORE",
		1110: "ORGANIZATION_CATEGORY_COSMETICS",
		1120: "ORGANIZATION_CATEGORY_CRUISE",
		1130: "ORGANIZATION_CATEGORY_FAST_CASUAL_DINING",
		1140: "ORGANIZATION_CATEGORY_FINANCIAL_SERVICES",
		1150: "ORGANIZATION_CATEGORY_FUEL",
		1160: "ORGANIZATION_CATEGORY_FURNITURE",
		1170: "ORGANIZATION_CATEGORY_GROCERY",
		1180: "ORGANIZATION_CATEGORY_HEALTH_CARE",
		1190: "ORGANIZATION_CATEGORY_HOME_IMPROVEMENT",
		1200: "ORGANIZATION_CATEGORY_HOTEL",
		1210: "ORGANIZATION_CATEGORY_MONEY_TRANSMISSION",
		1220: "ORGANIZATION_CATEGORY_PHARMACY",
		1230: "ORGANIZATION_CATEGORY_PROFESSIONAL_SERVICES",
		1240: "ORGANIZATION_CATEGORY_QUICK_SERVICE_RESTAURANT",
		1250: "ORGANIZATION_CATEGORY_RESTAURANT",
		1260: "ORGANIZATION_CATEGORY_STREAMING_MEDIA",
		1270: "ORGANIZATION_CATEGORY_TELEPHONY_AND_INTERNET",
	}
	OrganizationCategory_value = map[string]int32{
		"ORGANIZATION_CATEGORY_UNSPECIFIED":                 0,
		"ORGANIZATION_CATEGORY_ADULT_SERVICES":              1000,
		"ORGANIZATION_CATEGORY_AIRLINE":                     1010,
		"ORGANIZATION_CATEGORY_ANTIQUES_AND_ARTWORK":        1020,
		"ORGANIZATION_CATEGORY_APPAREL":                     1030,
		"ORGANIZATION_CATEGORY_AUTOMOTIVE_REPAIR":           1040,
		"ORGANIZATION_CATEGORY_AUTOMOTIVE_SALES":            1050,
		"ORGANIZATION_CATEGORY_BOOKSTORE":                   1060,
		"ORGANIZATION_CATEGORY_CAR_RENTAL":                  1070,
		"ORGANIZATION_CATEGORY_CONSUMER_ELECTRONICS":        1080,
		"ORGANIZATION_CATEGORY_CONSUMER_ELECTRONICS_REPAIR": 1090,
		"ORGANIZATION_CATEGORY_CONVENIENCE_STORE":           1100,
		"ORGANIZATION_CATEGORY_COSMETICS":                   1110,
		"ORGANIZATION_CATEGORY_CRUISE":                      1120,
		"ORGANIZATION_CATEGORY_FAST_CASUAL_DINING":          1130,
		"ORGANIZATION_CATEGORY_FINANCIAL_SERVICES":          1140,
		"ORGANIZATION_CATEGORY_FUEL":                        1150,
		"ORGANIZATION_CATEGORY_FURNITURE":                   1160,
		"ORGANIZATION_CATEGORY_GROCERY":                     1170,
		"ORGANIZATION_CATEGORY_HEALTH_CARE":                 1180,
		"ORGANIZATION_CATEGORY_HOME_IMPROVEMENT":            1190,
		"ORGANIZATION_CATEGORY_HOTEL":                       1200,
		"ORGANIZATION_CATEGORY_MONEY_TRANSMISSION":          1210,
		"ORGANIZATION_CATEGORY_PHARMACY":                    1220,
		"ORGANIZATION_CATEGORY_PROFESSIONAL_SERVICES":       1230,
		"ORGANIZATION_CATEGORY_QUICK_SERVICE_RESTAURANT":    1240,
		"ORGANIZATION_CATEGORY_RESTAURANT":                  1250,
		"ORGANIZATION_CATEGORY_STREAMING_MEDIA":             1260,
		"ORGANIZATION_CATEGORY_TELEPHONY_AND_INTERNET":      1270,
	}
)

func (x OrganizationCategory) Enum() *OrganizationCategory {
	p := new(OrganizationCategory)
	*p = x
	return p
}

func (x OrganizationCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrganizationCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_micashared_common_enums_organizationcategory_v1_organization_category_proto_enumTypes[0].Descriptor()
}

func (OrganizationCategory) Type() protoreflect.EnumType {
	return &file_micashared_common_enums_organizationcategory_v1_organization_category_proto_enumTypes[0]
}

func (x OrganizationCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrganizationCategory.Descriptor instead.
func (OrganizationCategory) EnumDescriptor() ([]byte, []int) {
	return file_micashared_common_enums_organizationcategory_v1_organization_category_proto_rawDescGZIP(), []int{0}
}

var File_micashared_common_enums_organizationcategory_v1_organization_category_proto protoreflect.FileDescriptor

var file_micashared_common_enums_organizationcategory_v1_organization_category_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2a, 0xea,
	0x09, 0x0a, 0x14, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x21, 0x4f, 0x52, 0x47, 0x41, 0x4e,
	0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x29,
	0x0a, 0x24, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x41, 0x44, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x53, 0x10, 0xe8, 0x07, 0x12, 0x22, 0x0a, 0x1d, 0x4f, 0x52, 0x47,
	0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x59, 0x5f, 0x41, 0x49, 0x52, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0xf2, 0x07, 0x12, 0x2f, 0x0a,
	0x2a, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x41, 0x4e, 0x54, 0x49, 0x51, 0x55, 0x45, 0x53, 0x5f,
	0x41, 0x4e, 0x44, 0x5f, 0x41, 0x52, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0xfc, 0x07, 0x12, 0x22,
	0x0a, 0x1d, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x41, 0x50, 0x50, 0x41, 0x52, 0x45, 0x4c, 0x10,
	0x86, 0x08, 0x12, 0x2c, 0x0a, 0x27, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x41, 0x55, 0x54, 0x4f,
	0x4d, 0x4f, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x41, 0x49, 0x52, 0x10, 0x90, 0x08,
	0x12, 0x2b, 0x0a, 0x26, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x4f,
	0x54, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x10, 0x9a, 0x08, 0x12, 0x24, 0x0a,
	0x1f, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x53, 0x54, 0x4f, 0x52, 0x45,
	0x10, 0xa4, 0x08, 0x12, 0x25, 0x0a, 0x20, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x41, 0x52,
	0x5f, 0x52, 0x45, 0x4e, 0x54, 0x41, 0x4c, 0x10, 0xae, 0x08, 0x12, 0x2f, 0x0a, 0x2a, 0x4f, 0x52,
	0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47,
	0x4f, 0x52, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x5f, 0x45, 0x4c, 0x45,
	0x43, 0x54, 0x52, 0x4f, 0x4e, 0x49, 0x43, 0x53, 0x10, 0xb8, 0x08, 0x12, 0x36, 0x0a, 0x31, 0x4f,
	0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x5f, 0x45, 0x4c,
	0x45, 0x43, 0x54, 0x52, 0x4f, 0x4e, 0x49, 0x43, 0x53, 0x5f, 0x52, 0x45, 0x50, 0x41, 0x49, 0x52,
	0x10, 0xc2, 0x08, 0x12, 0x2c, 0x0a, 0x27, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x4e, 0x49, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x10, 0xcc,
	0x08, 0x12, 0x24, 0x0a, 0x1f, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x4f, 0x53, 0x4d, 0x45,
	0x54, 0x49, 0x43, 0x53, 0x10, 0xd6, 0x08, 0x12, 0x21, 0x0a, 0x1c, 0x4f, 0x52, 0x47, 0x41, 0x4e,
	0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59,
	0x5f, 0x43, 0x52, 0x55, 0x49, 0x53, 0x45, 0x10, 0xe0, 0x08, 0x12, 0x2d, 0x0a, 0x28, 0x4f, 0x52,
	0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47,
	0x4f, 0x52, 0x59, 0x5f, 0x46, 0x41, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x53, 0x55, 0x41, 0x4c, 0x5f,
	0x44, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0xea, 0x08, 0x12, 0x2d, 0x0a, 0x28, 0x4f, 0x52, 0x47,
	0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x59, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4e, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x53, 0x10, 0xf4, 0x08, 0x12, 0x1f, 0x0a, 0x1a, 0x4f, 0x52, 0x47, 0x41,
	0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52,
	0x59, 0x5f, 0x46, 0x55, 0x45, 0x4c, 0x10, 0xfe, 0x08, 0x12, 0x24, 0x0a, 0x1f, 0x4f, 0x52, 0x47,
	0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x59, 0x5f, 0x46, 0x55, 0x52, 0x4e, 0x49, 0x54, 0x55, 0x52, 0x45, 0x10, 0x88, 0x09, 0x12,
	0x22, 0x0a, 0x1d, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x47, 0x52, 0x4f, 0x43, 0x45, 0x52, 0x59,
	0x10, 0x92, 0x09, 0x12, 0x26, 0x0a, 0x21, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x48, 0x45, 0x41,
	0x4c, 0x54, 0x48, 0x5f, 0x43, 0x41, 0x52, 0x45, 0x10, 0x9c, 0x09, 0x12, 0x2b, 0x0a, 0x26, 0x4f,
	0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x5f, 0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x49, 0x4d, 0x50, 0x52, 0x4f, 0x56,
	0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0xa6, 0x09, 0x12, 0x20, 0x0a, 0x1b, 0x4f, 0x52, 0x47, 0x41,
	0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52,
	0x59, 0x5f, 0x48, 0x4f, 0x54, 0x45, 0x4c, 0x10, 0xb0, 0x09, 0x12, 0x2d, 0x0a, 0x28, 0x4f, 0x52,
	0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47,
	0x4f, 0x52, 0x59, 0x5f, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0xba, 0x09, 0x12, 0x23, 0x0a, 0x1e, 0x4f, 0x52, 0x47,
	0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x59, 0x5f, 0x50, 0x48, 0x41, 0x52, 0x4d, 0x41, 0x43, 0x59, 0x10, 0xc4, 0x09, 0x12, 0x30,
	0x0a, 0x2b, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x45, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x53, 0x10, 0xce, 0x09,
	0x12, 0x33, 0x0a, 0x2e, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x51, 0x55, 0x49, 0x43, 0x4b, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x41, 0x55, 0x52, 0x41,
	0x4e, 0x54, 0x10, 0xd8, 0x09, 0x12, 0x25, 0x0a, 0x20, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x52,
	0x45, 0x53, 0x54, 0x41, 0x55, 0x52, 0x41, 0x4e, 0x54, 0x10, 0xe2, 0x09, 0x12, 0x2a, 0x0a, 0x25,
	0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54,
	0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x49, 0x4e, 0x47, 0x5f,
	0x4d, 0x45, 0x44, 0x49, 0x41, 0x10, 0xec, 0x09, 0x12, 0x31, 0x0a, 0x2c, 0x4f, 0x52, 0x47, 0x41,
	0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52,
	0x59, 0x5f, 0x54, 0x45, 0x4c, 0x45, 0x50, 0x48, 0x4f, 0x4e, 0x59, 0x5f, 0x41, 0x4e, 0x44, 0x5f,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x10, 0xf6, 0x09, 0x42, 0x75, 0x0a, 0x1d, 0x69,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x1d, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49,
	0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_enums_organizationcategory_v1_organization_category_proto_rawDescOnce sync.Once
	file_micashared_common_enums_organizationcategory_v1_organization_category_proto_rawDescData = file_micashared_common_enums_organizationcategory_v1_organization_category_proto_rawDesc
)

func file_micashared_common_enums_organizationcategory_v1_organization_category_proto_rawDescGZIP() []byte {
	file_micashared_common_enums_organizationcategory_v1_organization_category_proto_rawDescOnce.Do(func() {
		file_micashared_common_enums_organizationcategory_v1_organization_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_enums_organizationcategory_v1_organization_category_proto_rawDescData)
	})
	return file_micashared_common_enums_organizationcategory_v1_organization_category_proto_rawDescData
}

var file_micashared_common_enums_organizationcategory_v1_organization_category_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_micashared_common_enums_organizationcategory_v1_organization_category_proto_goTypes = []interface{}{
	(OrganizationCategory)(0), // 0: micashared.common.enums.organizationcategory.v1.OrganizationCategory
}
var file_micashared_common_enums_organizationcategory_v1_organization_category_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_micashared_common_enums_organizationcategory_v1_organization_category_proto_init() }
func file_micashared_common_enums_organizationcategory_v1_organization_category_proto_init() {
	if File_micashared_common_enums_organizationcategory_v1_organization_category_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_micashared_common_enums_organizationcategory_v1_organization_category_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_enums_organizationcategory_v1_organization_category_proto_goTypes,
		DependencyIndexes: file_micashared_common_enums_organizationcategory_v1_organization_category_proto_depIdxs,
		EnumInfos:         file_micashared_common_enums_organizationcategory_v1_organization_category_proto_enumTypes,
	}.Build()
	File_micashared_common_enums_organizationcategory_v1_organization_category_proto = out.File
	file_micashared_common_enums_organizationcategory_v1_organization_category_proto_rawDesc = nil
	file_micashared_common_enums_organizationcategory_v1_organization_category_proto_goTypes = nil
	file_micashared_common_enums_organizationcategory_v1_organization_category_proto_depIdxs = nil
}
