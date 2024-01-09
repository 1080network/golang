// Copyright (c) 2022 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: micashared/common/v1/date.proto

package commonv1

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

// Represents a whole or partial calendar date, e.g. a birthday. The time of day
// and time zone are either specified elsewhere or are not significant. The date
// is relative to the Proleptic Gregorian Calendar. This can represent:
//
// * A full date, with non-zero year, month and day values
// * A month and day value, with a zero year, e.g. an anniversary
// * A year on its own, with zero month and day values
// * A year and month value, with a zero day, e.g. a credit card expiration date
//
// Related types are [google.type.TimeOfDay][google.type.TimeOfDay] and `google.protobuf.Timestamp`.
type Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Year of date. Must be from 1 to 9999, or 0 if specifying a date without
	// a year.
	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// Month of year. Must be from 1 to 12, or 0 if specifying a year without a
	// month and day.
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	// Day of month. Must be from 1 to 31 and valid for the year and month, or 0
	// if specifying a year by itself or a year and month where the day is not
	// significant.
	Day int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *Date) Reset() {
	*x = Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_date_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_date_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Date.ProtoReflect.Descriptor instead.
func (*Date) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_date_proto_rawDescGZIP(), []int{0}
}

func (x *Date) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Date) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Date) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

var File_micashared_common_v1_date_proto protoreflect.FileDescriptor

var file_micashared_common_v1_date_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x42, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x42, 0x4a, 0x0a, 0x17, 0x69,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x44, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x1d, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31,
	0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_v1_date_proto_rawDescOnce sync.Once
	file_micashared_common_v1_date_proto_rawDescData = file_micashared_common_v1_date_proto_rawDesc
)

func file_micashared_common_v1_date_proto_rawDescGZIP() []byte {
	file_micashared_common_v1_date_proto_rawDescOnce.Do(func() {
		file_micashared_common_v1_date_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_v1_date_proto_rawDescData)
	})
	return file_micashared_common_v1_date_proto_rawDescData
}

var file_micashared_common_v1_date_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_micashared_common_v1_date_proto_goTypes = []interface{}{
	(*Date)(nil), // 0: micashared.common.v1.Date
}
var file_micashared_common_v1_date_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_micashared_common_v1_date_proto_init() }
func file_micashared_common_v1_date_proto_init() {
	if File_micashared_common_v1_date_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_micashared_common_v1_date_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Date); i {
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
			RawDescriptor: file_micashared_common_v1_date_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_v1_date_proto_goTypes,
		DependencyIndexes: file_micashared_common_v1_date_proto_depIdxs,
		MessageInfos:      file_micashared_common_v1_date_proto_msgTypes,
	}.Build()
	File_micashared_common_v1_date_proto = out.File
	file_micashared_common_v1_date_proto_rawDesc = nil
	file_micashared_common_v1_date_proto_goTypes = nil
	file_micashared_common_v1_date_proto_depIdxs = nil
}
