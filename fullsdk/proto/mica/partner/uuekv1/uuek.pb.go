// Copyright (c) 2022-2024 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: mica/partner/uuek/v1/uuek.proto

package uuekv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "github.com/1080network/golang/fullsdk/proto/micashared/common/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UUEK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// reference to the user at the Partner
	PartnerUserRef string `protobuf:"bytes,3,opt,name=partner_user_ref,json=partnerUserRef,proto3" json:"partner_user_ref,omitempty"`
	// reference to the user's instrument at the Partner
	PartnerInstrumentRef string `protobuf:"bytes,1,opt,name=partner_instrument_ref,json=partnerInstrumentRef,proto3" json:"partner_instrument_ref,omitempty"`
	// depending on how the Partner handles instruments and UUEKs, this will either be their reference to this
	// UUEK or a random UUID value if they don't have this concept
	PartnerUuekRef string         `protobuf:"bytes,4,opt,name=partner_uuek_ref,json=partnerUuekRef,proto3" json:"partner_uuek_ref,omitempty"`
	Uuek           *v1.CommonUUEK `protobuf:"bytes,2,opt,name=uuek,proto3" json:"uuek,omitempty"`
}

func (x *UUEK) Reset() {
	*x = UUEK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUEK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUEK) ProtoMessage() {}

func (x *UUEK) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_uuek_v1_uuek_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUEK.ProtoReflect.Descriptor instead.
func (*UUEK) Descriptor() ([]byte, []int) {
	return file_mica_partner_uuek_v1_uuek_proto_rawDescGZIP(), []int{0}
}

func (x *UUEK) GetPartnerUserRef() string {
	if x != nil {
		return x.PartnerUserRef
	}
	return ""
}

func (x *UUEK) GetPartnerInstrumentRef() string {
	if x != nil {
		return x.PartnerInstrumentRef
	}
	return ""
}

func (x *UUEK) GetPartnerUuekRef() string {
	if x != nil {
		return x.PartnerUuekRef
	}
	return ""
}

func (x *UUEK) GetUuek() *v1.CommonUUEK {
	if x != nil {
		return x.Uuek
	}
	return nil
}

var File_mica_partner_uuek_v1_uuek_proto protoreflect.FileDescriptor

var file_mica_partner_uuek_v1_uuek_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x75,
	0x75, 0x65, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x75, 0x65, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e,
	0x75, 0x75, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x75,
	0x65, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x04, 0x55, 0x55, 0x45, 0x4b, 0x12, 0x33, 0x0a, 0x10, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52,
	0x0e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x66, 0x12,
	0x3f, 0x0a, 0x16, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x14, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66,
	0x12, 0x33, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x65, 0x6b,
	0x5f, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x55, 0x75,
	0x65, 0x6b, 0x52, 0x65, 0x66, 0x12, 0x34, 0x0a, 0x04, 0x75, 0x75, 0x65, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x55, 0x55, 0x45, 0x4b, 0x52, 0x04, 0x75, 0x75, 0x65, 0x6b, 0x42, 0x40, 0x0a, 0x17, 0x69,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x75,
	0x75, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x55, 0x55, 0x45, 0x4b, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x13, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f,
	0x75, 0x75, 0x65, 0x6b, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mica_partner_uuek_v1_uuek_proto_rawDescOnce sync.Once
	file_mica_partner_uuek_v1_uuek_proto_rawDescData = file_mica_partner_uuek_v1_uuek_proto_rawDesc
)

func file_mica_partner_uuek_v1_uuek_proto_rawDescGZIP() []byte {
	file_mica_partner_uuek_v1_uuek_proto_rawDescOnce.Do(func() {
		file_mica_partner_uuek_v1_uuek_proto_rawDescData = protoimpl.X.CompressGZIP(file_mica_partner_uuek_v1_uuek_proto_rawDescData)
	})
	return file_mica_partner_uuek_v1_uuek_proto_rawDescData
}

var file_mica_partner_uuek_v1_uuek_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mica_partner_uuek_v1_uuek_proto_goTypes = []interface{}{
	(*UUEK)(nil),          // 0: mica.partner.uuek.v1.UUEK
	(*v1.CommonUUEK)(nil), // 1: micashared.common.v1.CommonUUEK
}
var file_mica_partner_uuek_v1_uuek_proto_depIdxs = []int32{
	1, // 0: mica.partner.uuek.v1.UUEK.uuek:type_name -> micashared.common.v1.CommonUUEK
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mica_partner_uuek_v1_uuek_proto_init() }
func file_mica_partner_uuek_v1_uuek_proto_init() {
	if File_mica_partner_uuek_v1_uuek_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mica_partner_uuek_v1_uuek_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUEK); i {
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
			RawDescriptor: file_mica_partner_uuek_v1_uuek_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mica_partner_uuek_v1_uuek_proto_goTypes,
		DependencyIndexes: file_mica_partner_uuek_v1_uuek_proto_depIdxs,
		MessageInfos:      file_mica_partner_uuek_v1_uuek_proto_msgTypes,
	}.Build()
	File_mica_partner_uuek_v1_uuek_proto = out.File
	file_mica_partner_uuek_v1_uuek_proto_rawDesc = nil
	file_mica_partner_uuek_v1_uuek_proto_goTypes = nil
	file_mica_partner_uuek_v1_uuek_proto_depIdxs = nil
}
