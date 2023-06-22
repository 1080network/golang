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
// source: mica/partner/widget/v1/widget.proto

package instrumentv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "github.com/1080network/golang/partner/proto/micashared/common/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InitializeWidgetResponse_Status int32

const (
	InitializeWidgetResponse_STATUS_UNSPECIFIED InitializeWidgetResponse_Status = 0
	InitializeWidgetResponse_STATUS_SUCCESS     InitializeWidgetResponse_Status = 1
	// generic error that's not one of the following
	InitializeWidgetResponse_STATUS_ERROR InitializeWidgetResponse_Status = 2
)

// Enum value maps for InitializeWidgetResponse_Status.
var (
	InitializeWidgetResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
	}
	InitializeWidgetResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
	}
)

func (x InitializeWidgetResponse_Status) Enum() *InitializeWidgetResponse_Status {
	p := new(InitializeWidgetResponse_Status)
	*p = x
	return p
}

func (x InitializeWidgetResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InitializeWidgetResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_partner_widget_v1_widget_proto_enumTypes[0].Descriptor()
}

func (InitializeWidgetResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_partner_widget_v1_widget_proto_enumTypes[0]
}

func (x InitializeWidgetResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InitializeWidgetResponse_Status.Descriptor instead.
func (InitializeWidgetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_partner_widget_v1_widget_proto_rawDescGZIP(), []int{1, 0}
}

type InitializeWidgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Details of the user.
	UserDemographic *v1.UserDemographic `protobuf:"bytes,3,opt,name=user_demographic,json=userDemographic,proto3" json:"user_demographic,omitempty"`
}

func (x *InitializeWidgetRequest) Reset() {
	*x = InitializeWidgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_widget_v1_widget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeWidgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeWidgetRequest) ProtoMessage() {}

func (x *InitializeWidgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_widget_v1_widget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeWidgetRequest.ProtoReflect.Descriptor instead.
func (*InitializeWidgetRequest) Descriptor() ([]byte, []int) {
	return file_mica_partner_widget_v1_widget_proto_rawDescGZIP(), []int{0}
}

func (x *InitializeWidgetRequest) GetUserDemographic() *v1.UserDemographic {
	if x != nil {
		return x.UserDemographic
	}
	return nil
}

type InitializeWidgetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     InitializeWidgetResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.partner.widget.v1.InitializeWidgetResponse_Status" json:"status,omitempty"`
	Error      *v1.Error                       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	SessionKey string                          `protobuf:"bytes,3,opt,name=session_key,json=sessionKey,proto3" json:"session_key,omitempty"`
}

func (x *InitializeWidgetResponse) Reset() {
	*x = InitializeWidgetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_widget_v1_widget_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeWidgetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeWidgetResponse) ProtoMessage() {}

func (x *InitializeWidgetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_widget_v1_widget_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeWidgetResponse.ProtoReflect.Descriptor instead.
func (*InitializeWidgetResponse) Descriptor() ([]byte, []int) {
	return file_mica_partner_widget_v1_widget_proto_rawDescGZIP(), []int{1}
}

func (x *InitializeWidgetResponse) GetStatus() InitializeWidgetResponse_Status {
	if x != nil {
		return x.Status
	}
	return InitializeWidgetResponse_STATUS_UNSPECIFIED
}

func (x *InitializeWidgetResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *InitializeWidgetResponse) GetSessionKey() string {
	if x != nil {
		return x.SessionKey
	}
	return ""
}

var File_mica_partner_widget_v1_widget_proto protoreflect.FileDescriptor

var file_mica_partner_widget_v1_widget_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x17, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6d,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6d, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x92, 0x02, 0x0a, 0x18, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4b, 0x65, 0x79, 0x22, 0x46, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x42, 0x4e, 0x0a, 0x1d, 0x69,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x57, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x19, 0x6d, 0x69, 0x63, 0x61, 0x2f,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mica_partner_widget_v1_widget_proto_rawDescOnce sync.Once
	file_mica_partner_widget_v1_widget_proto_rawDescData = file_mica_partner_widget_v1_widget_proto_rawDesc
)

func file_mica_partner_widget_v1_widget_proto_rawDescGZIP() []byte {
	file_mica_partner_widget_v1_widget_proto_rawDescOnce.Do(func() {
		file_mica_partner_widget_v1_widget_proto_rawDescData = protoimpl.X.CompressGZIP(file_mica_partner_widget_v1_widget_proto_rawDescData)
	})
	return file_mica_partner_widget_v1_widget_proto_rawDescData
}

var file_mica_partner_widget_v1_widget_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mica_partner_widget_v1_widget_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mica_partner_widget_v1_widget_proto_goTypes = []interface{}{
	(InitializeWidgetResponse_Status)(0), // 0: mica.partner.widget.v1.InitializeWidgetResponse.Status
	(*InitializeWidgetRequest)(nil),      // 1: mica.partner.widget.v1.InitializeWidgetRequest
	(*InitializeWidgetResponse)(nil),     // 2: mica.partner.widget.v1.InitializeWidgetResponse
	(*v1.UserDemographic)(nil),           // 3: micashared.common.v1.UserDemographic
	(*v1.Error)(nil),                     // 4: micashared.common.v1.Error
}
var file_mica_partner_widget_v1_widget_proto_depIdxs = []int32{
	3, // 0: mica.partner.widget.v1.InitializeWidgetRequest.user_demographic:type_name -> micashared.common.v1.UserDemographic
	0, // 1: mica.partner.widget.v1.InitializeWidgetResponse.status:type_name -> mica.partner.widget.v1.InitializeWidgetResponse.Status
	4, // 2: mica.partner.widget.v1.InitializeWidgetResponse.error:type_name -> micashared.common.v1.Error
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mica_partner_widget_v1_widget_proto_init() }
func file_mica_partner_widget_v1_widget_proto_init() {
	if File_mica_partner_widget_v1_widget_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mica_partner_widget_v1_widget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeWidgetRequest); i {
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
		file_mica_partner_widget_v1_widget_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeWidgetResponse); i {
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
			RawDescriptor: file_mica_partner_widget_v1_widget_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mica_partner_widget_v1_widget_proto_goTypes,
		DependencyIndexes: file_mica_partner_widget_v1_widget_proto_depIdxs,
		EnumInfos:         file_mica_partner_widget_v1_widget_proto_enumTypes,
		MessageInfos:      file_mica_partner_widget_v1_widget_proto_msgTypes,
	}.Build()
	File_mica_partner_widget_v1_widget_proto = out.File
	file_mica_partner_widget_v1_widget_proto_rawDesc = nil
	file_mica_partner_widget_v1_widget_proto_goTypes = nil
	file_mica_partner_widget_v1_widget_proto_depIdxs = nil
}
