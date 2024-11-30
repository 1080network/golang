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
// source: micashared/common/v1/user.proto

package commonv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	agebandv1 "github.com/1080network/golang/fullsdk/proto/micashared/common/enums/agebandv1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserDemographic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The first name of the user.
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"` // @gotags: mask:"pii"
	// The last name of the user.
	LastName string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"` // @gotags: mask:"pii"
	// The email of the user.
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"` // @gotags: mask:"pii"
	// The street address of the user.
	Address *Address `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// The age band of the user.
	AgeBand agebandv1.AgeBand `protobuf:"varint,7,opt,name=age_band,json=ageBand,proto3,enum=micashared.common.enums.ageband.v1.AgeBand" json:"age_band,omitempty"`
	// E.164 format.
	Phone string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"` // @gotags: mask:"pii"
}

func (x *UserDemographic) Reset() {
	*x = UserDemographic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDemographic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDemographic) ProtoMessage() {}

func (x *UserDemographic) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDemographic.ProtoReflect.Descriptor instead.
func (*UserDemographic) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserDemographic) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserDemographic) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserDemographic) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserDemographic) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *UserDemographic) GetAgeBand() agebandv1.AgeBand {
	if x != nil {
		return x.AgeBand
	}
	return agebandv1.AgeBand(0)
}

func (x *UserDemographic) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// version of the user record, used for optimistic locking.
	Version int64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Date that the Store was created at Mica.
	Created *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	// Date that the Store was last updated at Mica.
	Updated *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	// Details of the user.
	UserDemographic *UserDemographic `protobuf:"bytes,4,opt,name=user_demographic,json=userDemographic,proto3" json:"user_demographic,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *User) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *User) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *User) GetUserDemographic() *UserDemographic {
	if x != nil {
		return x.UserDemographic
	}
	return nil
}

var File_micashared_common_v1_user_proto protoreflect.FileDescriptor

var file_micashared_common_v1_user_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x61, 0x67, 0x65, 0x62, 0x61, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x65,
	0x5f, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x02, 0x0a, 0x0f, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x37,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x46, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x5f, 0x62,
	0x61, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x61, 0x67, 0x65, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x67, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x61, 0x67, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x12,
	0x32, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c,
	0xfa, 0x42, 0x19, 0x72, 0x17, 0x32, 0x15, 0x5e, 0x24, 0x7c, 0x5e, 0x5c, 0x2b, 0x5b, 0x31, 0x2d,
	0x39, 0x5d, 0x5c, 0x64, 0x7b, 0x31, 0x30, 0x2c, 0x31, 0x34, 0x7d, 0x24, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x22, 0xde, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x50, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x69, 0x63, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x69, 0x63, 0x42, 0x4a, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_v1_user_proto_rawDescOnce sync.Once
	file_micashared_common_v1_user_proto_rawDescData = file_micashared_common_v1_user_proto_rawDesc
)

func file_micashared_common_v1_user_proto_rawDescGZIP() []byte {
	file_micashared_common_v1_user_proto_rawDescOnce.Do(func() {
		file_micashared_common_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_v1_user_proto_rawDescData)
	})
	return file_micashared_common_v1_user_proto_rawDescData
}

var file_micashared_common_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_micashared_common_v1_user_proto_goTypes = []interface{}{
	(*UserDemographic)(nil),       // 0: micashared.common.v1.UserDemographic
	(*User)(nil),                  // 1: micashared.common.v1.User
	(*Address)(nil),               // 2: micashared.common.v1.Address
	(agebandv1.AgeBand)(0),        // 3: micashared.common.enums.ageband.v1.AgeBand
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_micashared_common_v1_user_proto_depIdxs = []int32{
	2, // 0: micashared.common.v1.UserDemographic.address:type_name -> micashared.common.v1.Address
	3, // 1: micashared.common.v1.UserDemographic.age_band:type_name -> micashared.common.enums.ageband.v1.AgeBand
	4, // 2: micashared.common.v1.User.created:type_name -> google.protobuf.Timestamp
	4, // 3: micashared.common.v1.User.updated:type_name -> google.protobuf.Timestamp
	0, // 4: micashared.common.v1.User.user_demographic:type_name -> micashared.common.v1.UserDemographic
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_micashared_common_v1_user_proto_init() }
func file_micashared_common_v1_user_proto_init() {
	if File_micashared_common_v1_user_proto != nil {
		return
	}
	file_micashared_common_v1_address_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_micashared_common_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDemographic); i {
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
		file_micashared_common_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_micashared_common_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_v1_user_proto_goTypes,
		DependencyIndexes: file_micashared_common_v1_user_proto_depIdxs,
		MessageInfos:      file_micashared_common_v1_user_proto_msgTypes,
	}.Build()
	File_micashared_common_v1_user_proto = out.File
	file_micashared_common_v1_user_proto_rawDesc = nil
	file_micashared_common_v1_user_proto_goTypes = nil
	file_micashared_common_v1_user_proto_depIdxs = nil
}