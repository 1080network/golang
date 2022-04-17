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
// source: common/authenticate/v1/authenticate.proto

package authenticatev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	realmv1 "mica/proto/common/enums/realmv1"
	v1 "mica/proto/common/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthenticateResponse_Status int32

const (
	AuthenticateResponse_STATUS_UNSPECIFIED    AuthenticateResponse_Status = 0
	AuthenticateResponse_STATUS_SUCCESS        AuthenticateResponse_Status = 1
	AuthenticateResponse_STATUS_ERROR          AuthenticateResponse_Status = 2
	AuthenticateResponse_STATUS_ACCOUNT_CLOSED AuthenticateResponse_Status = 3
)

// Enum value maps for AuthenticateResponse_Status.
var (
	AuthenticateResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_ACCOUNT_CLOSED",
	}
	AuthenticateResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED":    0,
		"STATUS_SUCCESS":        1,
		"STATUS_ERROR":          2,
		"STATUS_ACCOUNT_CLOSED": 3,
	}
)

func (x AuthenticateResponse_Status) Enum() *AuthenticateResponse_Status {
	p := new(AuthenticateResponse_Status)
	*p = x
	return p
}

func (x AuthenticateResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthenticateResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_common_authenticate_v1_authenticate_proto_enumTypes[0].Descriptor()
}

func (AuthenticateResponse_Status) Type() protoreflect.EnumType {
	return &file_common_authenticate_v1_authenticate_proto_enumTypes[0]
}

func (x AuthenticateResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthenticateResponse_Status.Descriptor instead.
func (AuthenticateResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_common_authenticate_v1_authenticate_proto_rawDescGZIP(), []int{1, 0}
}

type AuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Realm    realmv1.Realm `protobuf:"varint,1,opt,name=realm,proto3,enum=common.enums.realm.v1.Realm" json:"realm,omitempty"`
	Username string        `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string        `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_authenticate_v1_authenticate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateRequest) ProtoMessage() {}

func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_authenticate_v1_authenticate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_common_authenticate_v1_authenticate_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticateRequest) GetRealm() realmv1.Realm {
	if x != nil {
		return x.Realm
	}
	return realmv1.Realm(0)
}

func (x *AuthenticateRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthenticateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthenticateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status AuthenticateResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=common.authenticate.v1.AuthenticateResponse_Status" json:"status,omitempty"`
	Error  *v1.Error                   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// JWT
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthenticateResponse) Reset() {
	*x = AuthenticateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_authenticate_v1_authenticate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateResponse) ProtoMessage() {}

func (x *AuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_authenticate_v1_authenticate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_common_authenticate_v1_authenticate_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateResponse) GetStatus() AuthenticateResponse_Status {
	if x != nil {
		return x.Status
	}
	return AuthenticateResponse_STATUS_UNSPECIFIED
}

func (x *AuthenticateResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *AuthenticateResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_common_authenticate_v1_authenticate_proto protoreflect.FileDescriptor

var file_common_authenticate_v1_authenticate_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01,
	0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61,
	0x6c, 0x6d, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x84, 0x02, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x61, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x19, 0x0a,
	0x15, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f,
	0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x03, 0x42, 0x61, 0x0a, 0x23, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x11, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x20, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_authenticate_v1_authenticate_proto_rawDescOnce sync.Once
	file_common_authenticate_v1_authenticate_proto_rawDescData = file_common_authenticate_v1_authenticate_proto_rawDesc
)

func file_common_authenticate_v1_authenticate_proto_rawDescGZIP() []byte {
	file_common_authenticate_v1_authenticate_proto_rawDescOnce.Do(func() {
		file_common_authenticate_v1_authenticate_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_authenticate_v1_authenticate_proto_rawDescData)
	})
	return file_common_authenticate_v1_authenticate_proto_rawDescData
}

var file_common_authenticate_v1_authenticate_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_authenticate_v1_authenticate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_authenticate_v1_authenticate_proto_goTypes = []interface{}{
	(AuthenticateResponse_Status)(0), // 0: common.authenticate.v1.AuthenticateResponse.Status
	(*AuthenticateRequest)(nil),      // 1: common.authenticate.v1.AuthenticateRequest
	(*AuthenticateResponse)(nil),     // 2: common.authenticate.v1.AuthenticateResponse
	(realmv1.Realm)(0),               // 3: common.enums.realm.v1.Realm
	(*v1.Error)(nil),                 // 4: common.v1.Error
}
var file_common_authenticate_v1_authenticate_proto_depIdxs = []int32{
	3, // 0: common.authenticate.v1.AuthenticateRequest.realm:type_name -> common.enums.realm.v1.Realm
	0, // 1: common.authenticate.v1.AuthenticateResponse.status:type_name -> common.authenticate.v1.AuthenticateResponse.Status
	4, // 2: common.authenticate.v1.AuthenticateResponse.error:type_name -> common.v1.Error
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_authenticate_v1_authenticate_proto_init() }
func file_common_authenticate_v1_authenticate_proto_init() {
	if File_common_authenticate_v1_authenticate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_authenticate_v1_authenticate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateRequest); i {
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
		file_common_authenticate_v1_authenticate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateResponse); i {
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
			RawDescriptor: file_common_authenticate_v1_authenticate_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_authenticate_v1_authenticate_proto_goTypes,
		DependencyIndexes: file_common_authenticate_v1_authenticate_proto_depIdxs,
		EnumInfos:         file_common_authenticate_v1_authenticate_proto_enumTypes,
		MessageInfos:      file_common_authenticate_v1_authenticate_proto_msgTypes,
	}.Build()
	File_common_authenticate_v1_authenticate_proto = out.File
	file_common_authenticate_v1_authenticate_proto_rawDesc = nil
	file_common_authenticate_v1_authenticate_proto_goTypes = nil
	file_common_authenticate_v1_authenticate_proto_depIdxs = nil
}
