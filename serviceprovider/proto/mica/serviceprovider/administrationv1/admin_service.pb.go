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

	CertificateId string `protobuf:"bytes,1,opt,name=certificate_id,json=certificateId,proto3" json:"certificate_id,omitempty"`
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

func (x *PingExternalWithCertificateRequest) GetCertificateId() string {
	if x != nil {
		return x.CertificateId
	}
	return ""
}

type PingExternalWithCertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertificateId string               `protobuf:"bytes,1,opt,name=certificate_id,json=certificateId,proto3" json:"certificate_id,omitempty"`
	Response      *pingv1.PingResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
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

func (x *PingExternalWithCertificateResponse) GetCertificateId() string {
	if x != nil {
		return x.CertificateId
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
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x22, 0x50, 0x69, 0x6e, 0x67,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x57, 0x69, 0x74, 0x68, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x0e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18,
	0x32, 0x52, 0x0d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x22, 0x91, 0x01, 0x0a, 0x23, 0x50, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x57, 0x69, 0x74, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x43, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa5, 0x0c, 0x0a, 0x24, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa6, 0x01,
	0x0a, 0x21, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69,
	0x63, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x3e, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa0, 0x01, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa0, 0x01, 0x0a, 0x1f, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d,
	0x69, 0x63, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63,
	0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa0, 0x01, 0x0a,
	0x1f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x3c, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x72,
	0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x72, 0x6f, 0x6d,
	0x4d, 0x69, 0x63, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x8e, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x36, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0xac, 0x01, 0x0a, 0x23, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d,
	0x69, 0x63, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x40, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
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
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x75, 0x0a, 0x29,
	0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x1a, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x25, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d,
	0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*PingExternalWithCertificateRequest)(nil),             // 0: mica.serviceprovider.administration.v1.PingExternalWithCertificateRequest
	(*PingExternalWithCertificateResponse)(nil),            // 1: mica.serviceprovider.administration.v1.PingExternalWithCertificateResponse
	(*pingv1.PingResponse)(nil),                            // 2: micashared.common.ping.v1.PingResponse
	(*v1.GenerateFromMicaClientCertificateRequest)(nil),    // 3: micashared.common.v1.GenerateFromMicaClientCertificateRequest
	(*v1.UpdateFromMicaClientCertificateRequest)(nil),      // 4: micashared.common.v1.UpdateFromMicaClientCertificateRequest
	(*v1.EnableFromMicaClientCertificateRequest)(nil),      // 5: micashared.common.v1.EnableFromMicaClientCertificateRequest
	(*v1.SearchFromMicaClientCertificateRequest)(nil),      // 6: micashared.common.v1.SearchFromMicaClientCertificateRequest
	(*v1.GetFromMicaClientSettingsRequest)(nil),            // 7: micashared.common.v1.GetFromMicaClientSettingsRequest
	(*v1.UpdateFromMicaClientCallBackAddressRequest)(nil),  // 8: micashared.common.v1.UpdateFromMicaClientCallBackAddressRequest
	(*v1.SearchDataExtractionRequest)(nil),                 // 9: micashared.common.v1.SearchDataExtractionRequest
	(*v1.GetDataExtractionRequest)(nil),                    // 10: micashared.common.v1.GetDataExtractionRequest
	(*pingv1.PingRequest)(nil),                             // 11: micashared.common.ping.v1.PingRequest
	(*v1.GenerateFromMicaClientCertificateResponse)(nil),   // 12: micashared.common.v1.GenerateFromMicaClientCertificateResponse
	(*v1.UpdateFromMicaClientCertificateResponse)(nil),     // 13: micashared.common.v1.UpdateFromMicaClientCertificateResponse
	(*v1.EnableFromMicaClientCertificateResponse)(nil),     // 14: micashared.common.v1.EnableFromMicaClientCertificateResponse
	(*v1.SearchFromMicaClientCertificateResponse)(nil),     // 15: micashared.common.v1.SearchFromMicaClientCertificateResponse
	(*v1.GetFromMicaClientSettingsResponse)(nil),           // 16: micashared.common.v1.GetFromMicaClientSettingsResponse
	(*v1.UpdateFromMicaClientCallBackAddressResponse)(nil), // 17: micashared.common.v1.UpdateFromMicaClientCallBackAddressResponse
	(*v1.SearchDataExtractionResponse)(nil),                // 18: micashared.common.v1.SearchDataExtractionResponse
	(*v1.GetDataExtractionResponse)(nil),                   // 19: micashared.common.v1.GetDataExtractionResponse
}
var file_mica_serviceprovider_administration_v1_admin_service_proto_depIdxs = []int32{
	2,  // 0: mica.serviceprovider.administration.v1.PingExternalWithCertificateResponse.response:type_name -> micashared.common.ping.v1.PingResponse
	3,  // 1: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GenerateFromMicaClientCertificate:input_type -> micashared.common.v1.GenerateFromMicaClientCertificateRequest
	4,  // 2: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.UpdateFromMicaClientCertificate:input_type -> micashared.common.v1.UpdateFromMicaClientCertificateRequest
	5,  // 3: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.EnableFromMicaClientCertificate:input_type -> micashared.common.v1.EnableFromMicaClientCertificateRequest
	6,  // 4: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.SearchFromMicaClientCertificate:input_type -> micashared.common.v1.SearchFromMicaClientCertificateRequest
	7,  // 5: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GetFromMicaClientSettings:input_type -> micashared.common.v1.GetFromMicaClientSettingsRequest
	8,  // 6: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.UpdateFromMicaClientCallbackAddress:input_type -> micashared.common.v1.UpdateFromMicaClientCallBackAddressRequest
	9,  // 7: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.SearchDataExtractionStatistics:input_type -> micashared.common.v1.SearchDataExtractionRequest
	10, // 8: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GetDataExtractionStatistics:input_type -> micashared.common.v1.GetDataExtractionRequest
	11, // 9: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.PingExternal:input_type -> micashared.common.ping.v1.PingRequest
	0,  // 10: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.PingExternalWithCertificate:input_type -> mica.serviceprovider.administration.v1.PingExternalWithCertificateRequest
	12, // 11: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GenerateFromMicaClientCertificate:output_type -> micashared.common.v1.GenerateFromMicaClientCertificateResponse
	13, // 12: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.UpdateFromMicaClientCertificate:output_type -> micashared.common.v1.UpdateFromMicaClientCertificateResponse
	14, // 13: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.EnableFromMicaClientCertificate:output_type -> micashared.common.v1.EnableFromMicaClientCertificateResponse
	15, // 14: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.SearchFromMicaClientCertificate:output_type -> micashared.common.v1.SearchFromMicaClientCertificateResponse
	16, // 15: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GetFromMicaClientSettings:output_type -> micashared.common.v1.GetFromMicaClientSettingsResponse
	17, // 16: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.UpdateFromMicaClientCallbackAddress:output_type -> micashared.common.v1.UpdateFromMicaClientCallBackAddressResponse
	18, // 17: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.SearchDataExtractionStatistics:output_type -> micashared.common.v1.SearchDataExtractionResponse
	19, // 18: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.GetDataExtractionStatistics:output_type -> micashared.common.v1.GetDataExtractionResponse
	2,  // 19: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.PingExternal:output_type -> micashared.common.ping.v1.PingResponse
	1,  // 20: mica.serviceprovider.administration.v1.ServiceProviderAdministrationService.PingExternalWithCertificate:output_type -> mica.serviceprovider.administration.v1.PingExternalWithCertificateResponse
	11, // [11:21] is the sub-list for method output_type
	1,  // [1:11] is the sub-list for method input_type
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
