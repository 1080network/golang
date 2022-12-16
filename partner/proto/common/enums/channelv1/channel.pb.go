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
// source: common/enums/channel/v1/channel.proto

package channelv1

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

type Channel int32

const (
	Channel_CHANNEL_UNSPECIFIED        Channel = 0
	Channel_CHANNEL_DOMESTIC_ACH       Channel = 1
	Channel_CHANNEL_INTERNATIONAL_ACH  Channel = 2
	Channel_CHANNEL_WIRE_TRANSFER      Channel = 3
	Channel_CHANNEL_SWIFT_TRANSFER     Channel = 4
	Channel_CHANNEL_REAL_TIME_PAYMENTS Channel = 5
)

// Enum value maps for Channel.
var (
	Channel_name = map[int32]string{
		0: "CHANNEL_UNSPECIFIED",
		1: "CHANNEL_DOMESTIC_ACH",
		2: "CHANNEL_INTERNATIONAL_ACH",
		3: "CHANNEL_WIRE_TRANSFER",
		4: "CHANNEL_SWIFT_TRANSFER",
		5: "CHANNEL_REAL_TIME_PAYMENTS",
	}
	Channel_value = map[string]int32{
		"CHANNEL_UNSPECIFIED":        0,
		"CHANNEL_DOMESTIC_ACH":       1,
		"CHANNEL_INTERNATIONAL_ACH":  2,
		"CHANNEL_WIRE_TRANSFER":      3,
		"CHANNEL_SWIFT_TRANSFER":     4,
		"CHANNEL_REAL_TIME_PAYMENTS": 5,
	}
)

func (x Channel) Enum() *Channel {
	p := new(Channel)
	*p = x
	return p
}

func (x Channel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Channel) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enums_channel_v1_channel_proto_enumTypes[0].Descriptor()
}

func (Channel) Type() protoreflect.EnumType {
	return &file_common_enums_channel_v1_channel_proto_enumTypes[0]
}

func (x Channel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Channel.Descriptor instead.
func (Channel) EnumDescriptor() ([]byte, []int) {
	return file_common_enums_channel_v1_channel_proto_rawDescGZIP(), []int{0}
}

var File_common_enums_channel_v1_channel_proto protoreflect.FileDescriptor

var file_common_enums_channel_v1_channel_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2a, 0xb2, 0x01, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x13,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x44, 0x4f, 0x4d, 0x45, 0x53, 0x54, 0x49, 0x43, 0x5f, 0x41, 0x43, 0x48, 0x10, 0x01, 0x12,
	0x1d, 0x0a, 0x19, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x41, 0x43, 0x48, 0x10, 0x02, 0x12, 0x19,
	0x0a, 0x15, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x57, 0x49, 0x52, 0x45, 0x5f, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x53, 0x57, 0x49, 0x46, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x46, 0x45, 0x52, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x52, 0x45, 0x41, 0x4c, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45,
	0x4e, 0x54, 0x53, 0x10, 0x05, 0x42, 0x55, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x21, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_enums_channel_v1_channel_proto_rawDescOnce sync.Once
	file_common_enums_channel_v1_channel_proto_rawDescData = file_common_enums_channel_v1_channel_proto_rawDesc
)

func file_common_enums_channel_v1_channel_proto_rawDescGZIP() []byte {
	file_common_enums_channel_v1_channel_proto_rawDescOnce.Do(func() {
		file_common_enums_channel_v1_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_enums_channel_v1_channel_proto_rawDescData)
	})
	return file_common_enums_channel_v1_channel_proto_rawDescData
}

var file_common_enums_channel_v1_channel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_enums_channel_v1_channel_proto_goTypes = []interface{}{
	(Channel)(0), // 0: common.enums.channel.v1.Channel
}
var file_common_enums_channel_v1_channel_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_enums_channel_v1_channel_proto_init() }
func file_common_enums_channel_v1_channel_proto_init() {
	if File_common_enums_channel_v1_channel_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_enums_channel_v1_channel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_enums_channel_v1_channel_proto_goTypes,
		DependencyIndexes: file_common_enums_channel_v1_channel_proto_depIdxs,
		EnumInfos:         file_common_enums_channel_v1_channel_proto_enumTypes,
	}.Build()
	File_common_enums_channel_v1_channel_proto = out.File
	file_common_enums_channel_v1_channel_proto_rawDesc = nil
	file_common_enums_channel_v1_channel_proto_goTypes = nil
	file_common_enums_channel_v1_channel_proto_depIdxs = nil
}
