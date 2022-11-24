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
// 	protoc        v3.14.0
// source: FireworksLaunchParamType.proto

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

type FireworksLaunchParamType int32

const (
	FireworksLaunchParamType_FIREWORKS_LAUNCH_PARAM_TYPE_NONE          FireworksLaunchParamType = 0
	FireworksLaunchParamType_FIREWORKS_LAUNCH_PARAM_TYPE_REPEAT        FireworksLaunchParamType = 1
	FireworksLaunchParamType_FIREWORKS_LAUNCH_PARAM_TYPE_INTEVAL       FireworksLaunchParamType = 2
	FireworksLaunchParamType_FIREWORKS_LAUNCH_PARAM_TYPE_DELAY         FireworksLaunchParamType = 3
	FireworksLaunchParamType_FIREWORKS_LAUNCH_PARAM_TYPE_ROUND_INTEVAL FireworksLaunchParamType = 4
	FireworksLaunchParamType_FIREWORKS_LAUNCH_PARAM_TYPE_MAX           FireworksLaunchParamType = 5
)

// Enum value maps for FireworksLaunchParamType.
var (
	FireworksLaunchParamType_name = map[int32]string{
		0: "FIREWORKS_LAUNCH_PARAM_TYPE_NONE",
		1: "FIREWORKS_LAUNCH_PARAM_TYPE_REPEAT",
		2: "FIREWORKS_LAUNCH_PARAM_TYPE_INTEVAL",
		3: "FIREWORKS_LAUNCH_PARAM_TYPE_DELAY",
		4: "FIREWORKS_LAUNCH_PARAM_TYPE_ROUND_INTEVAL",
		5: "FIREWORKS_LAUNCH_PARAM_TYPE_MAX",
	}
	FireworksLaunchParamType_value = map[string]int32{
		"FIREWORKS_LAUNCH_PARAM_TYPE_NONE":          0,
		"FIREWORKS_LAUNCH_PARAM_TYPE_REPEAT":        1,
		"FIREWORKS_LAUNCH_PARAM_TYPE_INTEVAL":       2,
		"FIREWORKS_LAUNCH_PARAM_TYPE_DELAY":         3,
		"FIREWORKS_LAUNCH_PARAM_TYPE_ROUND_INTEVAL": 4,
		"FIREWORKS_LAUNCH_PARAM_TYPE_MAX":           5,
	}
)

func (x FireworksLaunchParamType) Enum() *FireworksLaunchParamType {
	p := new(FireworksLaunchParamType)
	*p = x
	return p
}

func (x FireworksLaunchParamType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FireworksLaunchParamType) Descriptor() protoreflect.EnumDescriptor {
	return file_FireworksLaunchParamType_proto_enumTypes[0].Descriptor()
}

func (FireworksLaunchParamType) Type() protoreflect.EnumType {
	return &file_FireworksLaunchParamType_proto_enumTypes[0]
}

func (x FireworksLaunchParamType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FireworksLaunchParamType.Descriptor instead.
func (FireworksLaunchParamType) EnumDescriptor() ([]byte, []int) {
	return file_FireworksLaunchParamType_proto_rawDescGZIP(), []int{0}
}

var File_FireworksLaunchParamType_proto protoreflect.FileDescriptor

var file_FireworksLaunchParamType_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x4c, 0x61, 0x75, 0x6e, 0x63,
	0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x8c, 0x02, 0x0a, 0x18, 0x46, 0x69, 0x72, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x46, 0x49, 0x52, 0x45, 0x57, 0x4f, 0x52, 0x4b,
	0x53, 0x5f, 0x4c, 0x41, 0x55, 0x4e, 0x43, 0x48, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x46, 0x49,
	0x52, 0x45, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x5f, 0x4c, 0x41, 0x55, 0x4e, 0x43, 0x48, 0x5f, 0x50,
	0x41, 0x52, 0x41, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x45, 0x41, 0x54,
	0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x46, 0x49, 0x52, 0x45, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x5f,
	0x4c, 0x41, 0x55, 0x4e, 0x43, 0x48, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x56, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x25, 0x0a, 0x21, 0x46,
	0x49, 0x52, 0x45, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x5f, 0x4c, 0x41, 0x55, 0x4e, 0x43, 0x48, 0x5f,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x41, 0x59,
	0x10, 0x03, 0x12, 0x2d, 0x0a, 0x29, 0x46, 0x49, 0x52, 0x45, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x5f,
	0x4c, 0x41, 0x55, 0x4e, 0x43, 0x48, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x56, 0x41, 0x4c, 0x10,
	0x04, 0x12, 0x23, 0x0a, 0x1f, 0x46, 0x49, 0x52, 0x45, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x5f, 0x4c,
	0x41, 0x55, 0x4e, 0x43, 0x48, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4d, 0x41, 0x58, 0x10, 0x05, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FireworksLaunchParamType_proto_rawDescOnce sync.Once
	file_FireworksLaunchParamType_proto_rawDescData = file_FireworksLaunchParamType_proto_rawDesc
)

func file_FireworksLaunchParamType_proto_rawDescGZIP() []byte {
	file_FireworksLaunchParamType_proto_rawDescOnce.Do(func() {
		file_FireworksLaunchParamType_proto_rawDescData = protoimpl.X.CompressGZIP(file_FireworksLaunchParamType_proto_rawDescData)
	})
	return file_FireworksLaunchParamType_proto_rawDescData
}

var file_FireworksLaunchParamType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_FireworksLaunchParamType_proto_goTypes = []interface{}{
	(FireworksLaunchParamType)(0), // 0: proto.FireworksLaunchParamType
}
var file_FireworksLaunchParamType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FireworksLaunchParamType_proto_init() }
func file_FireworksLaunchParamType_proto_init() {
	if File_FireworksLaunchParamType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FireworksLaunchParamType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FireworksLaunchParamType_proto_goTypes,
		DependencyIndexes: file_FireworksLaunchParamType_proto_depIdxs,
		EnumInfos:         file_FireworksLaunchParamType_proto_enumTypes,
	}.Build()
	File_FireworksLaunchParamType_proto = out.File
	file_FireworksLaunchParamType_proto_rawDesc = nil
	file_FireworksLaunchParamType_proto_goTypes = nil
	file_FireworksLaunchParamType_proto_depIdxs = nil
}
