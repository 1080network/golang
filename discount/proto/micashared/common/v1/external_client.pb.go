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
// 	protoc        v5.28.3
// source: micashared/common/v1/external_client.proto

package commonv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientAuthenticationType int32

const (
	ClientAuthenticationType_CLIENT_AUTHENTICATION_TYPE_UNSPECIFIED        ClientAuthenticationType = 0
	ClientAuthenticationType_CLIENT_AUTHENTICATION_TYPE_CLIENT_CERTIFICATE ClientAuthenticationType = 1
	ClientAuthenticationType_CLIENT_AUTHENTICATION_TYPE_API_TOKEN          ClientAuthenticationType = 2
)

// Enum value maps for ClientAuthenticationType.
var (
	ClientAuthenticationType_name = map[int32]string{
		0: "CLIENT_AUTHENTICATION_TYPE_UNSPECIFIED",
		1: "CLIENT_AUTHENTICATION_TYPE_CLIENT_CERTIFICATE",
		2: "CLIENT_AUTHENTICATION_TYPE_API_TOKEN",
	}
	ClientAuthenticationType_value = map[string]int32{
		"CLIENT_AUTHENTICATION_TYPE_UNSPECIFIED":        0,
		"CLIENT_AUTHENTICATION_TYPE_CLIENT_CERTIFICATE": 1,
		"CLIENT_AUTHENTICATION_TYPE_API_TOKEN":          2,
	}
)

func (x ClientAuthenticationType) Enum() *ClientAuthenticationType {
	p := new(ClientAuthenticationType)
	*p = x
	return p
}

func (x ClientAuthenticationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientAuthenticationType) Descriptor() protoreflect.EnumDescriptor {
	return file_micashared_common_v1_external_client_proto_enumTypes[0].Descriptor()
}

func (ClientAuthenticationType) Type() protoreflect.EnumType {
	return &file_micashared_common_v1_external_client_proto_enumTypes[0]
}

func (x ClientAuthenticationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientAuthenticationType.Descriptor instead.
func (ClientAuthenticationType) EnumDescriptor() ([]byte, []int) {
	return file_micashared_common_v1_external_client_proto_rawDescGZIP(), []int{0}
}

type ExternalClientSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version                      int64                         `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Created                      *timestamppb.Timestamp        `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated                      *timestamppb.Timestamp        `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	CallbackAddress              string                        `protobuf:"bytes,4,opt,name=callback_address,json=callbackAddress,proto3" json:"callback_address,omitempty"`
	ClientAuthenticationSettings *ClientAuthenticationSettings `protobuf:"bytes,5,opt,name=client_authentication_settings,json=clientAuthenticationSettings,proto3" json:"client_authentication_settings,omitempty"`
}

func (x *ExternalClientSettings) Reset() {
	*x = ExternalClientSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_external_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalClientSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalClientSettings) ProtoMessage() {}

func (x *ExternalClientSettings) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_external_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalClientSettings.ProtoReflect.Descriptor instead.
func (*ExternalClientSettings) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_external_client_proto_rawDescGZIP(), []int{0}
}

func (x *ExternalClientSettings) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ExternalClientSettings) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ExternalClientSettings) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *ExternalClientSettings) GetCallbackAddress() string {
	if x != nil {
		return x.CallbackAddress
	}
	return ""
}

func (x *ExternalClientSettings) GetClientAuthenticationSettings() *ClientAuthenticationSettings {
	if x != nil {
		return x.ClientAuthenticationSettings
	}
	return nil
}

type ClientAuthenticationSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuportedClientAuthenticationTypes []ClientAuthenticationType `protobuf:"varint,1,rep,packed,name=suported_client_authentication_types,json=suportedClientAuthenticationTypes,proto3,enum=micashared.common.v1.ClientAuthenticationType" json:"suported_client_authentication_types,omitempty"`
	CurrentAuthenticationType         ClientAuthenticationType   `protobuf:"varint,2,opt,name=current_authentication_type,json=currentAuthenticationType,proto3,enum=micashared.common.v1.ClientAuthenticationType" json:"current_authentication_type,omitempty"`
}

func (x *ClientAuthenticationSettings) Reset() {
	*x = ClientAuthenticationSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_external_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAuthenticationSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAuthenticationSettings) ProtoMessage() {}

func (x *ClientAuthenticationSettings) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_external_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAuthenticationSettings.ProtoReflect.Descriptor instead.
func (*ClientAuthenticationSettings) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_external_client_proto_rawDescGZIP(), []int{1}
}

func (x *ClientAuthenticationSettings) GetSuportedClientAuthenticationTypes() []ClientAuthenticationType {
	if x != nil {
		return x.SuportedClientAuthenticationTypes
	}
	return nil
}

func (x *ClientAuthenticationSettings) GetCurrentAuthenticationType() ClientAuthenticationType {
	if x != nil {
		return x.CurrentAuthenticationType
	}
	return ClientAuthenticationType_CLIENT_AUTHENTICATION_TYPE_UNSPECIFIED
}

var File_micashared_common_v1_external_client_proto protoreflect.FileDescriptor

var file_micashared_common_v1_external_client_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a,
	0x16, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28,
	0x00, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x10, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x03, 0x18, 0xc8, 0x01, 0x52, 0x0f, 0x63, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x78, 0x0a,
	0x1e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x1c, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x8f, 0x02, 0x0a, 0x1c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x7f, 0x0a, 0x24, 0x73, 0x75, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x21, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x6e, 0x0a, 0x1b, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x19,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x2a, 0xa3, 0x01, 0x0a, 0x18, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x26, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54,
	0x5f, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x31, 0x0a, 0x2d, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x55, 0x54,
	0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43,
	0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f,
	0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x49, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x02, 0x42,
	0x54, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a,
	0x1d, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02,
	0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_v1_external_client_proto_rawDescOnce sync.Once
	file_micashared_common_v1_external_client_proto_rawDescData = file_micashared_common_v1_external_client_proto_rawDesc
)

func file_micashared_common_v1_external_client_proto_rawDescGZIP() []byte {
	file_micashared_common_v1_external_client_proto_rawDescOnce.Do(func() {
		file_micashared_common_v1_external_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_v1_external_client_proto_rawDescData)
	})
	return file_micashared_common_v1_external_client_proto_rawDescData
}

var file_micashared_common_v1_external_client_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_micashared_common_v1_external_client_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_micashared_common_v1_external_client_proto_goTypes = []interface{}{
	(ClientAuthenticationType)(0),        // 0: micashared.common.v1.ClientAuthenticationType
	(*ExternalClientSettings)(nil),       // 1: micashared.common.v1.ExternalClientSettings
	(*ClientAuthenticationSettings)(nil), // 2: micashared.common.v1.ClientAuthenticationSettings
	(*timestamppb.Timestamp)(nil),        // 3: google.protobuf.Timestamp
}
var file_micashared_common_v1_external_client_proto_depIdxs = []int32{
	3, // 0: micashared.common.v1.ExternalClientSettings.created:type_name -> google.protobuf.Timestamp
	3, // 1: micashared.common.v1.ExternalClientSettings.updated:type_name -> google.protobuf.Timestamp
	2, // 2: micashared.common.v1.ExternalClientSettings.client_authentication_settings:type_name -> micashared.common.v1.ClientAuthenticationSettings
	0, // 3: micashared.common.v1.ClientAuthenticationSettings.suported_client_authentication_types:type_name -> micashared.common.v1.ClientAuthenticationType
	0, // 4: micashared.common.v1.ClientAuthenticationSettings.current_authentication_type:type_name -> micashared.common.v1.ClientAuthenticationType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_micashared_common_v1_external_client_proto_init() }
func file_micashared_common_v1_external_client_proto_init() {
	if File_micashared_common_v1_external_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_micashared_common_v1_external_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalClientSettings); i {
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
		file_micashared_common_v1_external_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAuthenticationSettings); i {
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
			RawDescriptor: file_micashared_common_v1_external_client_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_v1_external_client_proto_goTypes,
		DependencyIndexes: file_micashared_common_v1_external_client_proto_depIdxs,
		EnumInfos:         file_micashared_common_v1_external_client_proto_enumTypes,
		MessageInfos:      file_micashared_common_v1_external_client_proto_msgTypes,
	}.Build()
	File_micashared_common_v1_external_client_proto = out.File
	file_micashared_common_v1_external_client_proto_rawDesc = nil
	file_micashared_common_v1_external_client_proto_goTypes = nil
	file_micashared_common_v1_external_client_proto_depIdxs = nil
}
