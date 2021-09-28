// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/osconfig/v1/osconfig_zonal_service.proto

package osconfig

import (
	context "context"
	reflect "reflect"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_google_cloud_osconfig_v1_osconfig_zonal_service_proto protoreflect.FileDescriptor

var file_google_cloud_osconfig_v1_osconfig_zonal_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x73, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x7a, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x73, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x97, 0x07, 0x0a, 0x14,
	0x4f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5a, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xaa, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x46, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x39, 0x12, 0x37, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x2f,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0xc4, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x3d, 0x12, 0x3b, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0xda,
	0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0xd2, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x50, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x43, 0x12, 0x41, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f,
	0x2a, 0x2f, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xe8, 0x01,
	0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x55, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x46, 0x12, 0x44, 0x2f, 0x76, 0x31, 0x2f, 0x7b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x76, 0x75, 0x6c, 0x6e, 0x65,
	0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0xda,
	0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x1a, 0x4b, 0xca, 0x41, 0x17, 0x6f, 0x73, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77,
	0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xd1, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x19, 0x4f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5a, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x73, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0xaa, 0x02, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x4f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1b, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4f, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_google_cloud_osconfig_v1_osconfig_zonal_service_proto_goTypes = []interface{}{
	(*GetInventoryRequest)(nil),              // 0: google.cloud.osconfig.v1.GetInventoryRequest
	(*ListInventoriesRequest)(nil),           // 1: google.cloud.osconfig.v1.ListInventoriesRequest
	(*GetVulnerabilityReportRequest)(nil),    // 2: google.cloud.osconfig.v1.GetVulnerabilityReportRequest
	(*ListVulnerabilityReportsRequest)(nil),  // 3: google.cloud.osconfig.v1.ListVulnerabilityReportsRequest
	(*Inventory)(nil),                        // 4: google.cloud.osconfig.v1.Inventory
	(*ListInventoriesResponse)(nil),          // 5: google.cloud.osconfig.v1.ListInventoriesResponse
	(*VulnerabilityReport)(nil),              // 6: google.cloud.osconfig.v1.VulnerabilityReport
	(*ListVulnerabilityReportsResponse)(nil), // 7: google.cloud.osconfig.v1.ListVulnerabilityReportsResponse
}
var file_google_cloud_osconfig_v1_osconfig_zonal_service_proto_depIdxs = []int32{
	0, // 0: google.cloud.osconfig.v1.OsConfigZonalService.GetInventory:input_type -> google.cloud.osconfig.v1.GetInventoryRequest
	1, // 1: google.cloud.osconfig.v1.OsConfigZonalService.ListInventories:input_type -> google.cloud.osconfig.v1.ListInventoriesRequest
	2, // 2: google.cloud.osconfig.v1.OsConfigZonalService.GetVulnerabilityReport:input_type -> google.cloud.osconfig.v1.GetVulnerabilityReportRequest
	3, // 3: google.cloud.osconfig.v1.OsConfigZonalService.ListVulnerabilityReports:input_type -> google.cloud.osconfig.v1.ListVulnerabilityReportsRequest
	4, // 4: google.cloud.osconfig.v1.OsConfigZonalService.GetInventory:output_type -> google.cloud.osconfig.v1.Inventory
	5, // 5: google.cloud.osconfig.v1.OsConfigZonalService.ListInventories:output_type -> google.cloud.osconfig.v1.ListInventoriesResponse
	6, // 6: google.cloud.osconfig.v1.OsConfigZonalService.GetVulnerabilityReport:output_type -> google.cloud.osconfig.v1.VulnerabilityReport
	7, // 7: google.cloud.osconfig.v1.OsConfigZonalService.ListVulnerabilityReports:output_type -> google.cloud.osconfig.v1.ListVulnerabilityReportsResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_osconfig_v1_osconfig_zonal_service_proto_init() }
func file_google_cloud_osconfig_v1_osconfig_zonal_service_proto_init() {
	if File_google_cloud_osconfig_v1_osconfig_zonal_service_proto != nil {
		return
	}
	file_google_cloud_osconfig_v1_inventory_proto_init()
	file_google_cloud_osconfig_v1_vulnerability_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_osconfig_v1_osconfig_zonal_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_osconfig_v1_osconfig_zonal_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_osconfig_v1_osconfig_zonal_service_proto_depIdxs,
	}.Build()
	File_google_cloud_osconfig_v1_osconfig_zonal_service_proto = out.File
	file_google_cloud_osconfig_v1_osconfig_zonal_service_proto_rawDesc = nil
	file_google_cloud_osconfig_v1_osconfig_zonal_service_proto_goTypes = nil
	file_google_cloud_osconfig_v1_osconfig_zonal_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OsConfigZonalServiceClient is the client API for OsConfigZonalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OsConfigZonalServiceClient interface {
	// Get inventory data for the specified VM instance. If the VM has no
	// associated inventory, the message `NOT_FOUND` is returned.
	GetInventory(ctx context.Context, in *GetInventoryRequest, opts ...grpc.CallOption) (*Inventory, error)
	// List inventory data for all VM instances in the specified zone.
	ListInventories(ctx context.Context, in *ListInventoriesRequest, opts ...grpc.CallOption) (*ListInventoriesResponse, error)
	// Gets the vulnerability report for the specified VM instance. Only VMs with
	// inventory data have vulnerability reports associated with them.
	GetVulnerabilityReport(ctx context.Context, in *GetVulnerabilityReportRequest, opts ...grpc.CallOption) (*VulnerabilityReport, error)
	// List vulnerability reports for all VM instances in the specified zone.
	ListVulnerabilityReports(ctx context.Context, in *ListVulnerabilityReportsRequest, opts ...grpc.CallOption) (*ListVulnerabilityReportsResponse, error)
}

type osConfigZonalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOsConfigZonalServiceClient(cc grpc.ClientConnInterface) OsConfigZonalServiceClient {
	return &osConfigZonalServiceClient{cc}
}

func (c *osConfigZonalServiceClient) GetInventory(ctx context.Context, in *GetInventoryRequest, opts ...grpc.CallOption) (*Inventory, error) {
	out := new(Inventory)
	err := c.cc.Invoke(ctx, "/google.cloud.osconfig.v1.OsConfigZonalService/GetInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osConfigZonalServiceClient) ListInventories(ctx context.Context, in *ListInventoriesRequest, opts ...grpc.CallOption) (*ListInventoriesResponse, error) {
	out := new(ListInventoriesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.osconfig.v1.OsConfigZonalService/ListInventories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osConfigZonalServiceClient) GetVulnerabilityReport(ctx context.Context, in *GetVulnerabilityReportRequest, opts ...grpc.CallOption) (*VulnerabilityReport, error) {
	out := new(VulnerabilityReport)
	err := c.cc.Invoke(ctx, "/google.cloud.osconfig.v1.OsConfigZonalService/GetVulnerabilityReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osConfigZonalServiceClient) ListVulnerabilityReports(ctx context.Context, in *ListVulnerabilityReportsRequest, opts ...grpc.CallOption) (*ListVulnerabilityReportsResponse, error) {
	out := new(ListVulnerabilityReportsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.osconfig.v1.OsConfigZonalService/ListVulnerabilityReports", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OsConfigZonalServiceServer is the server API for OsConfigZonalService service.
type OsConfigZonalServiceServer interface {
	// Get inventory data for the specified VM instance. If the VM has no
	// associated inventory, the message `NOT_FOUND` is returned.
	GetInventory(context.Context, *GetInventoryRequest) (*Inventory, error)
	// List inventory data for all VM instances in the specified zone.
	ListInventories(context.Context, *ListInventoriesRequest) (*ListInventoriesResponse, error)
	// Gets the vulnerability report for the specified VM instance. Only VMs with
	// inventory data have vulnerability reports associated with them.
	GetVulnerabilityReport(context.Context, *GetVulnerabilityReportRequest) (*VulnerabilityReport, error)
	// List vulnerability reports for all VM instances in the specified zone.
	ListVulnerabilityReports(context.Context, *ListVulnerabilityReportsRequest) (*ListVulnerabilityReportsResponse, error)
}

// UnimplementedOsConfigZonalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOsConfigZonalServiceServer struct {
}

func (*UnimplementedOsConfigZonalServiceServer) GetInventory(context.Context, *GetInventoryRequest) (*Inventory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInventory not implemented")
}
func (*UnimplementedOsConfigZonalServiceServer) ListInventories(context.Context, *ListInventoriesRequest) (*ListInventoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInventories not implemented")
}
func (*UnimplementedOsConfigZonalServiceServer) GetVulnerabilityReport(context.Context, *GetVulnerabilityReportRequest) (*VulnerabilityReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVulnerabilityReport not implemented")
}
func (*UnimplementedOsConfigZonalServiceServer) ListVulnerabilityReports(context.Context, *ListVulnerabilityReportsRequest) (*ListVulnerabilityReportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVulnerabilityReports not implemented")
}

func RegisterOsConfigZonalServiceServer(s *grpc.Server, srv OsConfigZonalServiceServer) {
	s.RegisterService(&_OsConfigZonalService_serviceDesc, srv)
}

func _OsConfigZonalService_GetInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsConfigZonalServiceServer).GetInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.osconfig.v1.OsConfigZonalService/GetInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsConfigZonalServiceServer).GetInventory(ctx, req.(*GetInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsConfigZonalService_ListInventories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInventoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsConfigZonalServiceServer).ListInventories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.osconfig.v1.OsConfigZonalService/ListInventories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsConfigZonalServiceServer).ListInventories(ctx, req.(*ListInventoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsConfigZonalService_GetVulnerabilityReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVulnerabilityReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsConfigZonalServiceServer).GetVulnerabilityReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.osconfig.v1.OsConfigZonalService/GetVulnerabilityReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsConfigZonalServiceServer).GetVulnerabilityReport(ctx, req.(*GetVulnerabilityReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsConfigZonalService_ListVulnerabilityReports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVulnerabilityReportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsConfigZonalServiceServer).ListVulnerabilityReports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.osconfig.v1.OsConfigZonalService/ListVulnerabilityReports",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsConfigZonalServiceServer).ListVulnerabilityReports(ctx, req.(*ListVulnerabilityReportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OsConfigZonalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.osconfig.v1.OsConfigZonalService",
	HandlerType: (*OsConfigZonalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInventory",
			Handler:    _OsConfigZonalService_GetInventory_Handler,
		},
		{
			MethodName: "ListInventories",
			Handler:    _OsConfigZonalService_ListInventories_Handler,
		},
		{
			MethodName: "GetVulnerabilityReport",
			Handler:    _OsConfigZonalService_GetVulnerabilityReport_Handler,
		},
		{
			MethodName: "ListVulnerabilityReports",
			Handler:    _OsConfigZonalService_ListVulnerabilityReports_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/osconfig/v1/osconfig_zonal_service.proto",
}
