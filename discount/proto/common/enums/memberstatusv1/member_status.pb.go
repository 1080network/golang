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
// source: micacommon/enums/memberstatus/v1/member_status.proto

package memberstatusv1

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

// State transition diagram:
//   PROVISIONING (initial state) -> PROVISIONED
//   PROVISIONED -> ACTIVE
//   ACTIVE -> SUSPENDED
//   SUSPENDED -> ACTIVE
type MemberStatus int32

const (
	MemberStatus_MEMBER_STATUS_UNSPECIFIED  MemberStatus = 0
	MemberStatus_MEMBER_STATUS_PROVISIONING MemberStatus = 1
	MemberStatus_MEMBER_STATUS_PROVISIONED  MemberStatus = 2
	MemberStatus_MEMBER_STATUS_ACTIVE       MemberStatus = 3
	MemberStatus_MEMBER_STATUS_SUSPENDED    MemberStatus = 4
)

// Enum value maps for MemberStatus.
var (
	MemberStatus_name = map[int32]string{
		0: "MEMBER_STATUS_UNSPECIFIED",
		1: "MEMBER_STATUS_PROVISIONING",
		2: "MEMBER_STATUS_PROVISIONED",
		3: "MEMBER_STATUS_ACTIVE",
		4: "MEMBER_STATUS_SUSPENDED",
	}
	MemberStatus_value = map[string]int32{
		"MEMBER_STATUS_UNSPECIFIED":  0,
		"MEMBER_STATUS_PROVISIONING": 1,
		"MEMBER_STATUS_PROVISIONED":  2,
		"MEMBER_STATUS_ACTIVE":       3,
		"MEMBER_STATUS_SUSPENDED":    4,
	}
)

func (x MemberStatus) Enum() *MemberStatus {
	p := new(MemberStatus)
	*p = x
	return p
}

func (x MemberStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MemberStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_micacommon_enums_memberstatus_v1_member_status_proto_enumTypes[0].Descriptor()
}

func (MemberStatus) Type() protoreflect.EnumType {
	return &file_micacommon_enums_memberstatus_v1_member_status_proto_enumTypes[0]
}

func (x MemberStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MemberStatus.Descriptor instead.
func (MemberStatus) EnumDescriptor() ([]byte, []int) {
	return file_micacommon_enums_memberstatus_v1_member_status_proto_rawDescGZIP(), []int{0}
}

var File_micacommon_enums_memberstatus_v1_member_status_proto protoreflect.FileDescriptor

var file_micacommon_enums_memberstatus_v1_member_status_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6d, 0x69, 0x63, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x6d, 0x69, 0x63, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0xa3, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x45, 0x4d,
	0x42, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x45, 0x4d, 0x42,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53,
	0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x45, 0x4d, 0x42,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53,
	0x49, 0x4f, 0x4e, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x4d, 0x42, 0x45,
	0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x03, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x04, 0x42, 0x5f,
	0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x26, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micacommon_enums_memberstatus_v1_member_status_proto_rawDescOnce sync.Once
	file_micacommon_enums_memberstatus_v1_member_status_proto_rawDescData = file_micacommon_enums_memberstatus_v1_member_status_proto_rawDesc
)

func file_micacommon_enums_memberstatus_v1_member_status_proto_rawDescGZIP() []byte {
	file_micacommon_enums_memberstatus_v1_member_status_proto_rawDescOnce.Do(func() {
		file_micacommon_enums_memberstatus_v1_member_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_micacommon_enums_memberstatus_v1_member_status_proto_rawDescData)
	})
	return file_micacommon_enums_memberstatus_v1_member_status_proto_rawDescData
}

var file_micacommon_enums_memberstatus_v1_member_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_micacommon_enums_memberstatus_v1_member_status_proto_goTypes = []interface{}{
	(MemberStatus)(0), // 0: micacommon.enums.memberstatus.v1.MemberStatus
}
var file_micacommon_enums_memberstatus_v1_member_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_micacommon_enums_memberstatus_v1_member_status_proto_init() }
func file_micacommon_enums_memberstatus_v1_member_status_proto_init() {
	if File_micacommon_enums_memberstatus_v1_member_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_micacommon_enums_memberstatus_v1_member_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micacommon_enums_memberstatus_v1_member_status_proto_goTypes,
		DependencyIndexes: file_micacommon_enums_memberstatus_v1_member_status_proto_depIdxs,
		EnumInfos:         file_micacommon_enums_memberstatus_v1_member_status_proto_enumTypes,
	}.Build()
	File_micacommon_enums_memberstatus_v1_member_status_proto = out.File
	file_micacommon_enums_memberstatus_v1_member_status_proto_rawDesc = nil
	file_micacommon_enums_memberstatus_v1_member_status_proto_goTypes = nil
	file_micacommon_enums_memberstatus_v1_member_status_proto_depIdxs = nil
}
