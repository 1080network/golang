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
// source: micashared/common/enums/country/v1/country.proto

package countryv1

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

// ISO 3166-2:US
type Country int32

const (
	Country_COUNTRY_UNSPECIFIED Country = 0
	Country_COUNTRY_US          Country = 1 // United States
	Country_COUNTRY_JM          Country = 2 // Jamacia
)

// Enum value maps for Country.
var (
	Country_name = map[int32]string{
		0: "COUNTRY_UNSPECIFIED",
		1: "COUNTRY_US",
		2: "COUNTRY_JM",
	}
	Country_value = map[string]int32{
		"COUNTRY_UNSPECIFIED": 0,
		"COUNTRY_US":          1,
		"COUNTRY_JM":          2,
	}
)

func (x Country) Enum() *Country {
	p := new(Country)
	*p = x
	return p
}

func (x Country) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Country) Descriptor() protoreflect.EnumDescriptor {
	return file_micashared_common_enums_country_v1_country_proto_enumTypes[0].Descriptor()
}

func (Country) Type() protoreflect.EnumType {
	return &file_micashared_common_enums_country_v1_country_proto_enumTypes[0]
}

func (x Country) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Country.Descriptor instead.
func (Country) EnumDescriptor() ([]byte, []int) {
	return file_micashared_common_enums_country_v1_country_proto_rawDescGZIP(), []int{0}
}

var File_micashared_common_enums_country_v1_country_proto protoreflect.FileDescriptor

var file_micashared_common_enums_country_v1_country_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2a, 0x42, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x55, 0x53, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x4a, 0x4d, 0x10, 0x02, 0x42, 0x5b, 0x0a, 0x1d, 0x69, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x21, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x76, 0x31,
	0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_enums_country_v1_country_proto_rawDescOnce sync.Once
	file_micashared_common_enums_country_v1_country_proto_rawDescData = file_micashared_common_enums_country_v1_country_proto_rawDesc
)

func file_micashared_common_enums_country_v1_country_proto_rawDescGZIP() []byte {
	file_micashared_common_enums_country_v1_country_proto_rawDescOnce.Do(func() {
		file_micashared_common_enums_country_v1_country_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_enums_country_v1_country_proto_rawDescData)
	})
	return file_micashared_common_enums_country_v1_country_proto_rawDescData
}

var file_micashared_common_enums_country_v1_country_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_micashared_common_enums_country_v1_country_proto_goTypes = []interface{}{
	(Country)(0), // 0: micashared.common.enums.country.v1.Country
}
var file_micashared_common_enums_country_v1_country_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_micashared_common_enums_country_v1_country_proto_init() }
func file_micashared_common_enums_country_v1_country_proto_init() {
	if File_micashared_common_enums_country_v1_country_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_micashared_common_enums_country_v1_country_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_enums_country_v1_country_proto_goTypes,
		DependencyIndexes: file_micashared_common_enums_country_v1_country_proto_depIdxs,
		EnumInfos:         file_micashared_common_enums_country_v1_country_proto_enumTypes,
	}.Build()
	File_micashared_common_enums_country_v1_country_proto = out.File
	file_micashared_common_enums_country_v1_country_proto_rawDesc = nil
	file_micashared_common_enums_country_v1_country_proto_goTypes = nil
	file_micashared_common_enums_country_v1_country_proto_depIdxs = nil
}
