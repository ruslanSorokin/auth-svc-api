// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: auth_ext_service.proto

package proto

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_auth_ext_service_proto protoreflect.FileDescriptor

var file_auth_ext_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x76, 0x31, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xab, 0x09, 0x0a, 0x1d, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xaa, 0x02, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x2e,
	0x2e, 0x76, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x76, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xbf, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x96, 0x01, 0x0a, 0x08, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x2b, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x76,
	0x69, 0x61, 0x20, 0x69, 0x74, 0x73, 0x20, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x1a, 0x5d, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x20, 0x61, 0x6e, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x76, 0x69, 0x61, 0x20,
	0x69, 0x74, 0x73, 0x20, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65,
	0x77, 0x20, 0x70, 0x61, 0x69, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x20, 0x26, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x12, 0xec, 0x02, 0x0a, 0x10, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x50, 0x61, 0x69, 0x72, 0x12, 0x39, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3a, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe0, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x22, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2d, 0x70, 0x61, 0x69, 0x72, 0x3a,
	0x01, 0x2a, 0x92, 0x41, 0xaa, 0x01, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x12, 0x35, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x27, 0x73, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x76, 0x69, 0x61, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x67, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x20, 0x61, 0x6e, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x27, 0x73, 0x20, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x76, 0x69, 0x61,
	0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x61,
	0x6e, 0x64, 0x20, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77,
	0x20, 0x70, 0x61, 0x69, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x20, 0x26, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x12, 0xdc, 0x01, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x2f, 0x2e, 0x76, 0x31,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x6c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x46, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x12, 0x16, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x20, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x20, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x22, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x27, 0x73,
	0x20, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x20, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x8e, 0x02, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x32, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x33, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x97, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22,
	0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x2d, 0x61, 0x6c, 0x6c,
	0x3a, 0x01, 0x2a, 0x92, 0x41, 0x6a, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x12, 0x2b, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x31, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x20, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x60, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x75, 0x73, 0x6c, 0x61, 0x6e, 0x53, 0x6f, 0x72, 0x6f, 0x6b, 0x69, 0x6e, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x92, 0x41, 0x0a, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a,
	0x01, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_auth_ext_service_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),             // 0: v1.authentication.external.proto.LoginRequest
	(*RefreshTokenPairRequest)(nil),  // 1: v1.authentication.external.proto.RefreshTokenPairRequest
	(*LogoutRequest)(nil),            // 2: v1.authentication.external.proto.LogoutRequest
	(*LogoutAllRequest)(nil),         // 3: v1.authentication.external.proto.LogoutAllRequest
	(*LoginResponse)(nil),            // 4: v1.authentication.external.proto.LoginResponse
	(*RefreshTokenPairResponse)(nil), // 5: v1.authentication.external.proto.RefreshTokenPairResponse
	(*LogoutResponse)(nil),           // 6: v1.authentication.external.proto.LogoutResponse
	(*LogoutAllResponse)(nil),        // 7: v1.authentication.external.proto.LogoutAllResponse
}
var file_auth_ext_service_proto_depIdxs = []int32{
	0, // 0: v1.authentication.external.proto.AuthenticationExternalService.Login:input_type -> v1.authentication.external.proto.LoginRequest
	1, // 1: v1.authentication.external.proto.AuthenticationExternalService.RefreshTokenPair:input_type -> v1.authentication.external.proto.RefreshTokenPairRequest
	2, // 2: v1.authentication.external.proto.AuthenticationExternalService.Logout:input_type -> v1.authentication.external.proto.LogoutRequest
	3, // 3: v1.authentication.external.proto.AuthenticationExternalService.LogoutAll:input_type -> v1.authentication.external.proto.LogoutAllRequest
	4, // 4: v1.authentication.external.proto.AuthenticationExternalService.Login:output_type -> v1.authentication.external.proto.LoginResponse
	5, // 5: v1.authentication.external.proto.AuthenticationExternalService.RefreshTokenPair:output_type -> v1.authentication.external.proto.RefreshTokenPairResponse
	6, // 6: v1.authentication.external.proto.AuthenticationExternalService.Logout:output_type -> v1.authentication.external.proto.LogoutResponse
	7, // 7: v1.authentication.external.proto.AuthenticationExternalService.LogoutAll:output_type -> v1.authentication.external.proto.LogoutAllResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_ext_service_proto_init() }
func file_auth_ext_service_proto_init() {
	if File_auth_ext_service_proto != nil {
		return
	}
	file_auth_ext_transport_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_ext_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_ext_service_proto_goTypes,
		DependencyIndexes: file_auth_ext_service_proto_depIdxs,
	}.Build()
	File_auth_ext_service_proto = out.File
	file_auth_ext_service_proto_rawDesc = nil
	file_auth_ext_service_proto_goTypes = nil
	file_auth_ext_service_proto_depIdxs = nil
}
