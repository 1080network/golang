// Copyright (c) 2022 Mica, Inc. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: common/enums/serviceprovidertype/v1/service_provider_type.proto

package serviceprovidertypev1

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

type ServiceProviderType int32

const (
	ServiceProviderType_SERVICE_PROVIDER_TYPE_UNSPECIFIED ServiceProviderType = 0
	ServiceProviderType_SERVICE_PROVIDER_TYPE_MICA        ServiceProviderType = 1
	ServiceProviderType_SERVICE_PROVIDER_TYPE_RIBBIT      ServiceProviderType = 2
)

// Enum value maps for ServiceProviderType.
var (
	ServiceProviderType_name = map[int32]string{
		0: "SERVICE_PROVIDER_TYPE_UNSPECIFIED",
		1: "SERVICE_PROVIDER_TYPE_MICA",
		2: "SERVICE_PROVIDER_TYPE_RIBBIT",
	}
	ServiceProviderType_value = map[string]int32{
		"SERVICE_PROVIDER_TYPE_UNSPECIFIED": 0,
		"SERVICE_PROVIDER_TYPE_MICA":        1,
		"SERVICE_PROVIDER_TYPE_RIBBIT":      2,
	}
)

func (x ServiceProviderType) Enum() *ServiceProviderType {
	p := new(ServiceProviderType)
	*p = x
	return p
}

func (x ServiceProviderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceProviderType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enums_serviceprovidertype_v1_service_provider_type_proto_enumTypes[0].Descriptor()
}

func (ServiceProviderType) Type() protoreflect.EnumType {
	return &file_common_enums_serviceprovidertype_v1_service_provider_type_proto_enumTypes[0]
}

func (x ServiceProviderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceProviderType.Descriptor instead.
func (ServiceProviderType) EnumDescriptor() ([]byte, []int) {
	return file_common_enums_serviceprovidertype_v1_service_provider_type_proto_rawDescGZIP(), []int{0}
}

var File_common_enums_serviceprovidertype_v1_service_provider_type_proto protoreflect.FileDescriptor

var file_common_enums_serviceprovidertype_v1_service_provider_type_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2a, 0x7e, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x21, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x49,
	0x43, 0x41, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x49,
	0x42, 0x42, 0x49, 0x54, 0x10, 0x02, 0x42, 0x72, 0x0a, 0x1c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x1c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2d, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x74, 0x79, 0x70,
	0x65, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_common_enums_serviceprovidertype_v1_service_provider_type_proto_rawDescOnce sync.Once
	file_common_enums_serviceprovidertype_v1_service_provider_type_proto_rawDescData = file_common_enums_serviceprovidertype_v1_service_provider_type_proto_rawDesc
)

func file_common_enums_serviceprovidertype_v1_service_provider_type_proto_rawDescGZIP() []byte {
	file_common_enums_serviceprovidertype_v1_service_provider_type_proto_rawDescOnce.Do(func() {
		file_common_enums_serviceprovidertype_v1_service_provider_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_enums_serviceprovidertype_v1_service_provider_type_proto_rawDescData)
	})
	return file_common_enums_serviceprovidertype_v1_service_provider_type_proto_rawDescData
}

var file_common_enums_serviceprovidertype_v1_service_provider_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_enums_serviceprovidertype_v1_service_provider_type_proto_goTypes = []interface{}{
	(ServiceProviderType)(0), // 0: common.enums.serviceprovidertype.v1.ServiceProviderType
}
var file_common_enums_serviceprovidertype_v1_service_provider_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_enums_serviceprovidertype_v1_service_provider_type_proto_init() }
func file_common_enums_serviceprovidertype_v1_service_provider_type_proto_init() {
	if File_common_enums_serviceprovidertype_v1_service_provider_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_enums_serviceprovidertype_v1_service_provider_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_enums_serviceprovidertype_v1_service_provider_type_proto_goTypes,
		DependencyIndexes: file_common_enums_serviceprovidertype_v1_service_provider_type_proto_depIdxs,
		EnumInfos:         file_common_enums_serviceprovidertype_v1_service_provider_type_proto_enumTypes,
	}.Build()
	File_common_enums_serviceprovidertype_v1_service_provider_type_proto = out.File
	file_common_enums_serviceprovidertype_v1_service_provider_type_proto_rawDesc = nil
	file_common_enums_serviceprovidertype_v1_service_provider_type_proto_goTypes = nil
	file_common_enums_serviceprovidertype_v1_service_provider_type_proto_depIdxs = nil
}
