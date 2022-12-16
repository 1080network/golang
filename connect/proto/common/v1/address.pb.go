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
// source: common/v1/address.proto

package commonv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	countryv1 "github.com/1080network/golang/connect/proto/common/enums/countryv1"
	regionv1 "github.com/1080network/golang/connect/proto/common/enums/regionv1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreetLines []string `protobuf:"bytes,1,rep,name=street_lines,json=streetLines,proto3" json:"street_lines,omitempty"`
	// typically city
	Locality string `protobuf:"bytes,2,opt,name=locality,proto3" json:"locality,omitempty"`
	// this would be state in the US (ISO 3166-2:US), prefecture in Japan (ISO 3166-2:JP) or
	// territory in Australia (ISO 3166-2:AU)
	Region regionv1.Region `protobuf:"varint,3,opt,name=region,proto3,enum=common.enums.region.v1.Region" json:"region,omitempty"`
	// postal code, in the US this would be zip code
	PostalCode string `protobuf:"bytes,4,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	// https://en.wikipedia.org/wiki/ISO_3166-2
	Country countryv1.Country `protobuf:"varint,5,opt,name=country,proto3,enum=common.enums.country.v1.Country" json:"country,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_common_v1_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetStreetLines() []string {
	if x != nil {
		return x.StreetLines
	}
	return nil
}

func (x *Address) GetLocality() string {
	if x != nil {
		return x.Locality
	}
	return ""
}

func (x *Address) GetRegion() regionv1.Region {
	if x != nil {
		return x.Region
	}
	return regionv1.Region(0)
}

func (x *Address) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Address) GetCountry() countryv1.Country {
	if x != nil {
		return x.Country
	}
	return countryv1.Country(0)
}

var File_common_v1_address_proto protoreflect.FileDescriptor

var file_common_v1_address_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdd, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x47, 0x0a, 0x11, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_v1_address_proto_rawDescOnce sync.Once
	file_common_v1_address_proto_rawDescData = file_common_v1_address_proto_rawDesc
)

func file_common_v1_address_proto_rawDescGZIP() []byte {
	file_common_v1_address_proto_rawDescOnce.Do(func() {
		file_common_v1_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v1_address_proto_rawDescData)
	})
	return file_common_v1_address_proto_rawDescData
}

var file_common_v1_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_v1_address_proto_goTypes = []interface{}{
	(*Address)(nil),        // 0: common.v1.Address
	(regionv1.Region)(0),   // 1: common.enums.region.v1.Region
	(countryv1.Country)(0), // 2: common.enums.country.v1.Country
}
var file_common_v1_address_proto_depIdxs = []int32{
	1, // 0: common.v1.Address.region:type_name -> common.enums.region.v1.Region
	2, // 1: common.v1.Address.country:type_name -> common.enums.country.v1.Country
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_v1_address_proto_init() }
func file_common_v1_address_proto_init() {
	if File_common_v1_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_v1_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_v1_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_address_proto_goTypes,
		DependencyIndexes: file_common_v1_address_proto_depIdxs,
		MessageInfos:      file_common_v1_address_proto_msgTypes,
	}.Build()
	File_common_v1_address_proto = out.File
	file_common_v1_address_proto_rawDesc = nil
	file_common_v1_address_proto_goTypes = nil
	file_common_v1_address_proto_depIdxs = nil
}
