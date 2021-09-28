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
// source: google/cloud/securitycenter/v1/resource.proto

package securitycenter

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Information related to the Google Cloud resource.
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The full resource name of the resource. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The full resource name of project that the resource belongs to.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// The human readable name of project that the resource belongs to.
	ProjectDisplayName string `protobuf:"bytes,3,opt,name=project_display_name,json=projectDisplayName,proto3" json:"project_display_name,omitempty"`
	// The full resource name of resource's parent.
	Parent string `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	// The human readable name of resource's parent.
	ParentDisplayName string `protobuf:"bytes,5,opt,name=parent_display_name,json=parentDisplayName,proto3" json:"parent_display_name,omitempty"`
	// Output only. Contains a Folder message for each folder in the assets ancestry.
	// The first folder is the deepest nested folder, and the last folder is the
	// folder directly under the Organization.
	Folders []*Folder `protobuf:"bytes,7,rep,name=folders,proto3" json:"folders,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Resource) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Resource) GetProjectDisplayName() string {
	if x != nil {
		return x.ProjectDisplayName
	}
	return ""
}

func (x *Resource) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *Resource) GetParentDisplayName() string {
	if x != nil {
		return x.ParentDisplayName
	}
	return ""
}

func (x *Resource) GetFolders() []*Folder {
	if x != nil {
		return x.Folders
	}
	return nil
}

var File_google_cloud_securitycenter_v1_resource_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_resource_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x01, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x2e, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x45, 0x0a, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x42, 0xe9, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0d,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0xaa, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_resource_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_resource_proto_rawDescData = file_google_cloud_securitycenter_v1_resource_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_resource_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_resource_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_resource_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_resource_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_v1_resource_proto_goTypes = []interface{}{
	(*Resource)(nil), // 0: google.cloud.securitycenter.v1.Resource
	(*Folder)(nil),   // 1: google.cloud.securitycenter.v1.Folder
}
var file_google_cloud_securitycenter_v1_resource_proto_depIdxs = []int32{
	1, // 0: google.cloud.securitycenter.v1.Resource.folders:type_name -> google.cloud.securitycenter.v1.Folder
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_resource_proto_init() }
func file_google_cloud_securitycenter_v1_resource_proto_init() {
	if File_google_cloud_securitycenter_v1_resource_proto != nil {
		return
	}
	file_google_cloud_securitycenter_v1_folder_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
			RawDescriptor: file_google_cloud_securitycenter_v1_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_resource_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_resource_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v1_resource_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_resource_proto = out.File
	file_google_cloud_securitycenter_v1_resource_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_resource_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_resource_proto_depIdxs = nil
}
