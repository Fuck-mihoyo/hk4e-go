// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: FoundationStatus.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FoundationStatus int32

const (
	FoundationStatus_FOUNDATION_STATUS_NONE     FoundationStatus = 0
	FoundationStatus_FOUNDATION_STATUS_INIT     FoundationStatus = 1
	FoundationStatus_FOUNDATION_STATUS_BUILDING FoundationStatus = 2
	FoundationStatus_FOUNDATION_STATUS_BUILT    FoundationStatus = 3
)

// Enum value maps for FoundationStatus.
var (
	FoundationStatus_name = map[int32]string{
		0: "FOUNDATION_STATUS_NONE",
		1: "FOUNDATION_STATUS_INIT",
		2: "FOUNDATION_STATUS_BUILDING",
		3: "FOUNDATION_STATUS_BUILT",
	}
	FoundationStatus_value = map[string]int32{
		"FOUNDATION_STATUS_NONE":     0,
		"FOUNDATION_STATUS_INIT":     1,
		"FOUNDATION_STATUS_BUILDING": 2,
		"FOUNDATION_STATUS_BUILT":    3,
	}
)

func (x FoundationStatus) Enum() *FoundationStatus {
	p := new(FoundationStatus)
	*p = x
	return p
}

func (x FoundationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FoundationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_FoundationStatus_proto_enumTypes[0].Descriptor()
}

func (FoundationStatus) Type() protoreflect.EnumType {
	return &file_FoundationStatus_proto_enumTypes[0]
}

func (x FoundationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FoundationStatus.Descriptor instead.
func (FoundationStatus) EnumDescriptor() ([]byte, []int) {
	return file_FoundationStatus_proto_rawDescGZIP(), []int{0}
}

var File_FoundationStatus_proto protoreflect.FileDescriptor

var file_FoundationStatus_proto_rawDesc = []byte{
	0x0a, 0x16, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x87, 0x01, 0x0a, 0x10, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a,
	0x16, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49,
	0x4e, 0x49, 0x54, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x42, 0x55, 0x49, 0x4c, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x42, 0x55, 0x49, 0x4c, 0x54,
	0x10, 0x03, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FoundationStatus_proto_rawDescOnce sync.Once
	file_FoundationStatus_proto_rawDescData = file_FoundationStatus_proto_rawDesc
)

func file_FoundationStatus_proto_rawDescGZIP() []byte {
	file_FoundationStatus_proto_rawDescOnce.Do(func() {
		file_FoundationStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_FoundationStatus_proto_rawDescData)
	})
	return file_FoundationStatus_proto_rawDescData
}

var file_FoundationStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_FoundationStatus_proto_goTypes = []interface{}{
	(FoundationStatus)(0), // 0: FoundationStatus
}
var file_FoundationStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FoundationStatus_proto_init() }
func file_FoundationStatus_proto_init() {
	if File_FoundationStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FoundationStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FoundationStatus_proto_goTypes,
		DependencyIndexes: file_FoundationStatus_proto_depIdxs,
		EnumInfos:         file_FoundationStatus_proto_enumTypes,
	}.Build()
	File_FoundationStatus_proto = out.File
	file_FoundationStatus_proto_rawDesc = nil
	file_FoundationStatus_proto_goTypes = nil
	file_FoundationStatus_proto_depIdxs = nil
}
