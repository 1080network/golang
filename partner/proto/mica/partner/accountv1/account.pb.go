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
// source: mica/partner/account/v1/account.proto

package accountv1

import (
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

type SearchABAAccountResponse_Status int32

const (
	SearchABAAccountResponse_STATUS_UNSPECIFIED SearchABAAccountResponse_Status = 0
	SearchABAAccountResponse_STATUS_SUCCESS     SearchABAAccountResponse_Status = 1
	SearchABAAccountResponse_STATUS_ERROR       SearchABAAccountResponse_Status = 2
	SearchABAAccountResponse_STATUS_NOT_FOUND   SearchABAAccountResponse_Status = 3
)

// Enum value maps for SearchABAAccountResponse_Status.
var (
	SearchABAAccountResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	SearchABAAccountResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x SearchABAAccountResponse_Status) Enum() *SearchABAAccountResponse_Status {
	p := new(SearchABAAccountResponse_Status)
	*p = x
	return p
}

func (x SearchABAAccountResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchABAAccountResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_partner_account_v1_account_proto_enumTypes[0].Descriptor()
}

func (SearchABAAccountResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_partner_account_v1_account_proto_enumTypes[0]
}

func (x SearchABAAccountResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchABAAccountResponse_Status.Descriptor instead.
func (SearchABAAccountResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_partner_account_v1_account_proto_rawDescGZIP(), []int{1, 0}
}

// request to search ABA accounts available in the partner instance
type SearchABAAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// if set the search will filter by bank name
	BankName string `protobuf:"bytes,1,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
}

func (x *SearchABAAccountRequest) Reset() {
	*x = SearchABAAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_account_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchABAAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchABAAccountRequest) ProtoMessage() {}

func (x *SearchABAAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_account_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchABAAccountRequest.ProtoReflect.Descriptor instead.
func (*SearchABAAccountRequest) Descriptor() ([]byte, []int) {
	return file_mica_partner_account_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *SearchABAAccountRequest) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

type SearchABAAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      SearchABAAccountResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.partner.account.v1.SearchABAAccountResponse_Status" json:"status,omitempty"`
	Error       *v1.Error                       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	AbaAccounts []*v1.ABAAccountEntry           `protobuf:"bytes,3,rep,name=aba_accounts,json=abaAccounts,proto3" json:"aba_accounts,omitempty"`
}

func (x *SearchABAAccountResponse) Reset() {
	*x = SearchABAAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_account_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchABAAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchABAAccountResponse) ProtoMessage() {}

func (x *SearchABAAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_account_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchABAAccountResponse.ProtoReflect.Descriptor instead.
func (*SearchABAAccountResponse) Descriptor() ([]byte, []int) {
	return file_mica_partner_account_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *SearchABAAccountResponse) GetStatus() SearchABAAccountResponse_Status {
	if x != nil {
		return x.Status
	}
	return SearchABAAccountResponse_STATUS_UNSPECIFIED
}

func (x *SearchABAAccountResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SearchABAAccountResponse) GetAbaAccounts() []*v1.ABAAccountEntry {
	if x != nil {
		return x.AbaAccounts
	}
	return nil
}

var File_mica_partner_account_v1_account_proto protoreflect.FileDescriptor

var file_mica_partner_account_v1_account_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x36, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x42, 0x41, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xc7, 0x02, 0x0a, 0x18, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x42, 0x41, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x42, 0x41, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x48, 0x0a, 0x0c,
	0x61, 0x62, 0x61, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x42, 0x41, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x62, 0x61, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x14,
	0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x03, 0x42, 0x49, 0x0a, 0x1a, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x42, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x16, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mica_partner_account_v1_account_proto_rawDescOnce sync.Once
	file_mica_partner_account_v1_account_proto_rawDescData = file_mica_partner_account_v1_account_proto_rawDesc
)

func file_mica_partner_account_v1_account_proto_rawDescGZIP() []byte {
	file_mica_partner_account_v1_account_proto_rawDescOnce.Do(func() {
		file_mica_partner_account_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_mica_partner_account_v1_account_proto_rawDescData)
	})
	return file_mica_partner_account_v1_account_proto_rawDescData
}

var file_mica_partner_account_v1_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mica_partner_account_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mica_partner_account_v1_account_proto_goTypes = []interface{}{
	(SearchABAAccountResponse_Status)(0), // 0: mica.partner.account.v1.SearchABAAccountResponse.Status
	(*SearchABAAccountRequest)(nil),      // 1: mica.partner.account.v1.SearchABAAccountRequest
	(*SearchABAAccountResponse)(nil),     // 2: mica.partner.account.v1.SearchABAAccountResponse
	(*v1.Error)(nil),                     // 3: micashared.common.v1.Error
	(*v1.ABAAccountEntry)(nil),           // 4: micashared.common.v1.ABAAccountEntry
}
var file_mica_partner_account_v1_account_proto_depIdxs = []int32{
	0, // 0: mica.partner.account.v1.SearchABAAccountResponse.status:type_name -> mica.partner.account.v1.SearchABAAccountResponse.Status
	3, // 1: mica.partner.account.v1.SearchABAAccountResponse.error:type_name -> micashared.common.v1.Error
	4, // 2: mica.partner.account.v1.SearchABAAccountResponse.aba_accounts:type_name -> micashared.common.v1.ABAAccountEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mica_partner_account_v1_account_proto_init() }
func file_mica_partner_account_v1_account_proto_init() {
	if File_mica_partner_account_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mica_partner_account_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchABAAccountRequest); i {
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
		file_mica_partner_account_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchABAAccountResponse); i {
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
			RawDescriptor: file_mica_partner_account_v1_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mica_partner_account_v1_account_proto_goTypes,
		DependencyIndexes: file_mica_partner_account_v1_account_proto_depIdxs,
		EnumInfos:         file_mica_partner_account_v1_account_proto_enumTypes,
		MessageInfos:      file_mica_partner_account_v1_account_proto_msgTypes,
	}.Build()
	File_mica_partner_account_v1_account_proto = out.File
	file_mica_partner_account_v1_account_proto_rawDesc = nil
	file_mica_partner_account_v1_account_proto_goTypes = nil
	file_mica_partner_account_v1_account_proto_depIdxs = nil
}
