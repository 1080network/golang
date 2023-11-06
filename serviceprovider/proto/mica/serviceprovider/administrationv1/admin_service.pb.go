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
// source: mica/serviceprovider/administration/v1/admin_service.proto

package administrationv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	pingv1 "github.com/1080network/golang/serviceprovider/proto/micashared/common/pingv1"
	v1 "github.com/1080network/golang/serviceprovider/proto/micashared/common/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PingExternalWithCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertificateRefKey string `protobuf:"bytes,1,opt,name=certificate_ref_key,json=certificateRefKey,proto3" json:"certificate_ref_key,omitempty"`
}

func (x *PingExternalWithCertificateRequest) Reset() {
	*x = PingExternalWithCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_serviceprovider_administration_v1_admin_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingExternalWithCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingExternalWithCertificateRequest) ProtoMessage() {}

func (x *PingExternalWithCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_serviceprovider_administration_v1_admin_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingExternalWithCertificateRequest.ProtoReflect.Descriptor instead.
func (*PingExternalWithCertificateRequest) Descriptor() ([]byte, []int) {
	return file_mica_serviceprovider_administration_v1_admin_service_proto_rawDescGZIP(), []int{0}
}

func (x *PingExternalWithCertificateRequest) GetCertificateRefKey() string {
	if x != nil {
		return x.CertificateRefKey
	}
	return ""
}

type PingExternalWithCertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertificateRefKey string               `protobuf:"bytes,1,opt,name=certificate_ref_key,json=certificateRefKey,proto3" json:"certificate_ref_key,omitempty"`
	Response          *pingv1.PingResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *PingExternalWithCertificateResponse) Reset() {
	*x = PingExternalWithCertificateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_serviceprovider_administration_v1_admin_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingExternalWithCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingExternalWithCertificateResponse) ProtoMessage() {}

func (x *PingExternalWithCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_serviceprovider_administration_v1_admin_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingExternalWithCertificateResponse.ProtoReflect.Descriptor instead.
func (*PingExternalWithCertificateResponse) Descriptor() ([]byte, []int) {
	return file_mica_serviceprovider_administration_v1_admin_service_proto_rawDescGZIP(), []int{1}
}

func (x *PingExternalWithCertificateResponse) GetCertificateRefKey() string {
	if x != nil {
		return x.CertificateRefKey
	}
	return ""
}

func (x *PingExternalWithCertificateResponse) GetResponse() *pingv1.PingResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_mica_serviceprovider_administration_v1_admin_service_proto protoreflect.FileDescriptor

var file_mica_serviceprovider_administration_v1_admin_service_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x6d, 0x69,
	0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2b, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5f, 0x0a, 0x22, 0x50, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x57, 0x69, 0x74, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x13, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52,
	0x11, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x66, 0x4b,
	0x65, 0x79, 0x22, 0x9a, 0x01, 0x0a, 0x23, 0x50, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x57, 0x69, 0x74, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x66, 0x4b, 0x65, 0x79, 0x12, 0x43, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x9a, 0x14, 0x0a, 0x24, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x54,
	0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x32, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x54, 0x4c, 0x53, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x33, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x12, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x54,
	0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9a, 0x01, 0x0a, 0x1d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3a, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e,
	0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9a, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f,
	0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x93, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3a, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x91, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x37, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8e, 0x01, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x36, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xac, 0x01,
	0x0a, 0x23, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x40, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xb2, 0x01, 0x0a,
	0x25, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x42, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0xac, 0x01, 0x0a, 0x23, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x40, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0xac, 0x01, 0x0a, 0x23, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x40, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x89, 0x01, 0x0a, 0x1e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x12, 0x31, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x80, 0x01, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x2e, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61,
	0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x26,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0xb8, 0x01, 0x0a, 0x1b, 0x50, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x57, 0x69, 0x74, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x4a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x57, 0x69, 0x74, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4b, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x57, 0x69, 0x74, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x70, 0x0a, 0x29,
	0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5a, 0x25, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mica_serviceprovider_administration_v1_admin_service_proto_rawDescOnce sync.Once
	file_mica_serviceprovider_administration_v1_admin_service_proto_rawDescData = file_mica_serviceprovider_administration_v1_admin_service_proto_rawDesc
)

func file_mica_serviceprovider_administration_v1_admin_service_proto_rawDescGZIP() []byte {
	file_mica_serviceprovider_administration_v1_admin_service_proto_rawDescOnce.Do(func() {
		file_mica_serviceprovider_administration_v1_admin_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_mica_serviceprovider_administration_v1_admin_service_proto_rawDescData)
	})
	return file_mica_serviceprovider_administration_v1_admin_service_proto_rawDescData
}

var file_mica_serviceprovider_administration_v1_admin_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mica_serviceprovider_administration_v1_admin_service_proto_goTypes = []interface{}{
	(*PingExternalWithCertificateRequest)(nil),               // 0: mica.serviceprovider.administration.v1.PingExternalWithCertificateRequest
	(*PingExternalWithCertificateResponse)(nil),              // 1: mica.serviceprovider.administration.v1.PingExternalWithCertificateResponse
	(*pingv1.PingResponse)(nil),                              // 2: micashared.common.ping.v1.PingResponse
	(*v1.GenerateMTLSCertificateRequest)(nil),                // 3: micashared.common.v1.GenerateMTLSCertificateRequest
	(*v1.UpdateMTLSCertificateRequest)(nil),                  // 4: micashared.common.v1.UpdateMTLSCertificateRequest
	(*v1.SearchMTLSCertificateRequest)(nil),                  // 5: micashared.common.v1.SearchMTLSCertificateRequest
	(*v1.GetMTLSCertificateRequest)(nil),                     // 6: micashared.common.v1.GetMTLSCertificateRequest
	(*v1.CreateSingleSignOnConsoleUserRequest)(nil),          // 7: micashared.common.v1.CreateSingleSignOnConsoleUserRequest
	(*v1.UpdateSingleSignOnConsoleUserRequest)(nil),          // 8: micashared.common.v1.UpdateSingleSignOnConsoleUserRequest
	(*v1.SearchSingleSignOnConsoleUserRequest)(nil),          // 9: micashared.common.v1.SearchSingleSignOnConsoleUserRequest
	(*v1.GetSingleSignOnConsoleUserRequest)(nil),             // 10: micashared.common.v1.GetSingleSignOnConsoleUserRequest
	(*v1.GetExternalClientSettingsRequest)(nil),              // 11: micashared.common.v1.GetExternalClientSettingsRequest
	(*v1.UpdateExternalClientCallBackAddressRequest)(nil),    // 12: micashared.common.v1.UpdateExternalClientCallBackAddressRequest
	(*v1.GenerateExternalClientMTLSCertificateRequest)(nil),  // 13: micashared.common.v1.GenerateExternalClientMTLSCertificateRequest
	(*v1.UpdateExternalClientMTLSCertificateRequest)(nil),    // 14: micashared.common.v1.UpdateExternalClientMTLSCertificateRequest
	(*v1.SearchExternalClientMTLSCertificateRequest)(nil),    // 15: micashared.common.v1.SearchExternalClientMTLSCertificateRequest
	(*v1.SearchDataExtractionRequest)(nil),                   // 16: micashared.common.v1.SearchDataExtractionRequest
	(*v1.GetDataExtractionRequest)(nil),                      // 17: micashared.common.v1.GetDataExtractionRequest
	(*pingv1.PingRequest)(nil),                               // 18: micashared.common.ping.v1.PingRequest
	(*v1.GenerateMTLSCertificateResponse)(nil),               // 19: micashared.common.v1.GenerateMTLSCertificateResponse
	(*v1.UpdateMTLSCertificateResponse)(nil),                 // 20: micashared.common.v1.UpdateMTLSCertificateResponse
	(*v1.SearchMTLSCertificateResponse)(nil),                 // 21: micashared.common.v1.SearchMTLSCertificateResponse
	(*v1.GetMTLSCertificateResponse)(nil),                    // 22: micashared.common.v1.GetMTLSCertificateResponse
	(*v1.CreateSingleSignOnConsoleUserResponse)(nil),         // 23: micashared.common.v1.CreateSingleSignOnConsoleUserResponse
	(*v1.UpdateSingleSignOnConsoleUserResponse)(nil),         // 24: micashared.common.v1.UpdateSingleSignOnConsoleUserResponse
	(*v1.SearchSingleSignOnConsoleUserResponse)(nil),         // 25: micashared.common.v1.SearchSingleSignOnConsoleUserResponse
	(*v1.GetSingleSignOnConsoleUserResponse)(nil),            // 26: micashared.common.v1.GetSingleSignOnConsoleUserResponse
	(*v1.GetExternalClientSettingsResponse)(nil),             // 27: micashared.common.v1.GetExternalClientSettingsResponse
	(*v1.UpdateExternalClientCallBackAddressResponse)(nil),   // 28: micashared.common.v1.UpdateExternalClientCallBackAddressResponse
	(*v1.GenerateExternalClientMTLSCertificateResponse)(nil), // 29: micashared.common.v1.GenerateExternalClientMTLSCertificateResponse
	(*v1.UpdateExternalClientMTLSCertificateResponse)(nil),   // 30: micashared.common.v1.UpdateExternalClientMTLSCertificateResponse
	(*v1.SearchExternalClientMTLSCertificateResponse)(nil),   // 31: micashared.common.v1.SearchExternalClientMTLSCertificateResponse
	(*v1.SearchDataExtractionResponse)(nil),                  // 32: micashared.common.v1.SearchDataExtractionResponse
	(*v1.GetDataExtractionResponse)(nil),                     // 33: micashared.common.v1.GetDataExtractionResponse
}
var file_mica_serviceprovider_administration_v1_admin_service_proto_depIdxs = []int32{
	2,  // 0: mica.serviceprovider.administration.v1.PingExternalWithCertificateResponse.response:type_name -> micashared.common.ping.v1.PingResponse
	3,  // 1: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GenerateMTLSCertificate:input_type -> micashared.common.v1.GenerateMTLSCertificateRequest
	4,  // 2: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.UpdateMTLSCertificate:input_type -> micashared.common.v1.UpdateMTLSCertificateRequest
	5,  // 3: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.SearchMTLSCertificate:input_type -> micashared.common.v1.SearchMTLSCertificateRequest
	6,  // 4: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GetMTLSCertificate:input_type -> micashared.common.v1.GetMTLSCertificateRequest
	7,  // 5: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.CreateSingleSignOnConsoleUser:input_type -> micashared.common.v1.CreateSingleSignOnConsoleUserRequest
	8,  // 6: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.UpdateSingleSignOnConsoleUser:input_type -> micashared.common.v1.UpdateSingleSignOnConsoleUserRequest
	9,  // 7: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.SearchSingleSignOnUser:input_type -> micashared.common.v1.SearchSingleSignOnConsoleUserRequest
	10, // 8: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GetSingleSignOnConsoleUser:input_type -> micashared.common.v1.GetSingleSignOnConsoleUserRequest
	11, // 9: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GetExternalClientSettings:input_type -> micashared.common.v1.GetExternalClientSettingsRequest
	12, // 10: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.UpdateExternalClientCallbackAddress:input_type -> micashared.common.v1.UpdateExternalClientCallBackAddressRequest
	13, // 11: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GenerateExternalClientMTLSCertificate:input_type -> micashared.common.v1.GenerateExternalClientMTLSCertificateRequest
	14, // 12: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.UpdateExternalClientMTLSCertificate:input_type -> micashared.common.v1.UpdateExternalClientMTLSCertificateRequest
	15, // 13: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.SearchExternalClientMTLSCertificate:input_type -> micashared.common.v1.SearchExternalClientMTLSCertificateRequest
	16, // 14: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.SearchDataExtractionStatistics:input_type -> micashared.common.v1.SearchDataExtractionRequest
	17, // 15: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GetDataExtractionStatistics:input_type -> micashared.common.v1.GetDataExtractionRequest
	18, // 16: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.PingExternal:input_type -> micashared.common.ping.v1.PingRequest
	0,  // 17: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.PingExternalWithCertificate:input_type -> mica.serviceprovider.administration.v1.PingExternalWithCertificateRequest
	19, // 18: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GenerateMTLSCertificate:output_type -> micashared.common.v1.GenerateMTLSCertificateResponse
	20, // 19: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.UpdateMTLSCertificate:output_type -> micashared.common.v1.UpdateMTLSCertificateResponse
	21, // 20: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.SearchMTLSCertificate:output_type -> micashared.common.v1.SearchMTLSCertificateResponse
	22, // 21: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GetMTLSCertificate:output_type -> micashared.common.v1.GetMTLSCertificateResponse
	23, // 22: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.CreateSingleSignOnConsoleUser:output_type -> micashared.common.v1.CreateSingleSignOnConsoleUserResponse
	24, // 23: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.UpdateSingleSignOnConsoleUser:output_type -> micashared.common.v1.UpdateSingleSignOnConsoleUserResponse
	25, // 24: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.SearchSingleSignOnUser:output_type -> micashared.common.v1.SearchSingleSignOnConsoleUserResponse
	26, // 25: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GetSingleSignOnConsoleUser:output_type -> micashared.common.v1.GetSingleSignOnConsoleUserResponse
	27, // 26: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GetExternalClientSettings:output_type -> micashared.common.v1.GetExternalClientSettingsResponse
	28, // 27: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.UpdateExternalClientCallbackAddress:output_type -> micashared.common.v1.UpdateExternalClientCallBackAddressResponse
	29, // 28: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GenerateExternalClientMTLSCertificate:output_type -> micashared.common.v1.GenerateExternalClientMTLSCertificateResponse
	30, // 29: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.UpdateExternalClientMTLSCertificate:output_type -> micashared.common.v1.UpdateExternalClientMTLSCertificateResponse
	31, // 30: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.SearchExternalClientMTLSCertificate:output_type -> micashared.common.v1.SearchExternalClientMTLSCertificateResponse
	32, // 31: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.SearchDataExtractionStatistics:output_type -> micashared.common.v1.SearchDataExtractionResponse
	33, // 32: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GetDataExtractionStatistics:output_type -> micashared.common.v1.GetDataExtractionResponse
	2,  // 33: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.PingExternal:output_type -> micashared.common.ping.v1.PingResponse
	1,  // 34: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.PingExternalWithCertificate:output_type -> mica.serviceprovider.administration.v1.PingExternalWithCertificateResponse
	18, // [18:35] is the sub-list for method output_type
	1,  // [1:18] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_mica_serviceprovider_administration_v1_admin_service_proto_init() }
func file_mica_serviceprovider_administration_v1_admin_service_proto_init() {
	if File_mica_serviceprovider_administration_v1_admin_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mica_serviceprovider_administration_v1_admin_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingExternalWithCertificateRequest); i {
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
		file_mica_serviceprovider_administration_v1_admin_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingExternalWithCertificateResponse); i {
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
			RawDescriptor: file_mica_serviceprovider_administration_v1_admin_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mica_serviceprovider_administration_v1_admin_service_proto_goTypes,
		DependencyIndexes: file_mica_serviceprovider_administration_v1_admin_service_proto_depIdxs,
		MessageInfos:      file_mica_serviceprovider_administration_v1_admin_service_proto_msgTypes,
	}.Build()
	File_mica_serviceprovider_administration_v1_admin_service_proto = out.File
	file_mica_serviceprovider_administration_v1_admin_service_proto_rawDesc = nil
	file_mica_serviceprovider_administration_v1_admin_service_proto_goTypes = nil
	file_mica_serviceprovider_administration_v1_admin_service_proto_depIdxs = nil
}
