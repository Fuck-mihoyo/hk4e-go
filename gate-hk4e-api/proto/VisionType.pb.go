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
// source: VisionType.proto

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

type VisionType int32

const (
	VisionType_VISION_TYPE_NONE                VisionType = 0
	VisionType_VISION_TYPE_MEET                VisionType = 1
	VisionType_VISION_TYPE_REBORN              VisionType = 2
	VisionType_VISION_TYPE_REPLACE             VisionType = 3
	VisionType_VISION_TYPE_WAYPOINT_REBORN     VisionType = 4
	VisionType_VISION_TYPE_MISS                VisionType = 5
	VisionType_VISION_TYPE_DIE                 VisionType = 6
	VisionType_VISION_TYPE_GATHER_ESCAPE       VisionType = 7
	VisionType_VISION_TYPE_REFRESH             VisionType = 8
	VisionType_VISION_TYPE_TRANSPORT           VisionType = 9
	VisionType_VISION_TYPE_REPLACE_DIE         VisionType = 10
	VisionType_VISION_TYPE_REPLACE_NO_NOTIFY   VisionType = 11
	VisionType_VISION_TYPE_BORN                VisionType = 12
	VisionType_VISION_TYPE_PICKUP              VisionType = 13
	VisionType_VISION_TYPE_REMOVE              VisionType = 14
	VisionType_VISION_TYPE_CHANGE_COSTUME      VisionType = 15
	VisionType_VISION_TYPE_FISH_REFRESH        VisionType = 16
	VisionType_VISION_TYPE_FISH_BIG_SHOCK      VisionType = 17
	VisionType_VISION_TYPE_FISH_QTE_SUCC       VisionType = 18
	VisionType_VISION_TYPE_Unk2700_EPFKMOIPADB VisionType = 19
)

// Enum value maps for VisionType.
var (
	VisionType_name = map[int32]string{
		0:  "VISION_TYPE_NONE",
		1:  "VISION_TYPE_MEET",
		2:  "VISION_TYPE_REBORN",
		3:  "VISION_TYPE_REPLACE",
		4:  "VISION_TYPE_WAYPOINT_REBORN",
		5:  "VISION_TYPE_MISS",
		6:  "VISION_TYPE_DIE",
		7:  "VISION_TYPE_GATHER_ESCAPE",
		8:  "VISION_TYPE_REFRESH",
		9:  "VISION_TYPE_TRANSPORT",
		10: "VISION_TYPE_REPLACE_DIE",
		11: "VISION_TYPE_REPLACE_NO_NOTIFY",
		12: "VISION_TYPE_BORN",
		13: "VISION_TYPE_PICKUP",
		14: "VISION_TYPE_REMOVE",
		15: "VISION_TYPE_CHANGE_COSTUME",
		16: "VISION_TYPE_FISH_REFRESH",
		17: "VISION_TYPE_FISH_BIG_SHOCK",
		18: "VISION_TYPE_FISH_QTE_SUCC",
		19: "VISION_TYPE_Unk2700_EPFKMOIPADB",
	}
	VisionType_value = map[string]int32{
		"VISION_TYPE_NONE":                0,
		"VISION_TYPE_MEET":                1,
		"VISION_TYPE_REBORN":              2,
		"VISION_TYPE_REPLACE":             3,
		"VISION_TYPE_WAYPOINT_REBORN":     4,
		"VISION_TYPE_MISS":                5,
		"VISION_TYPE_DIE":                 6,
		"VISION_TYPE_GATHER_ESCAPE":       7,
		"VISION_TYPE_REFRESH":             8,
		"VISION_TYPE_TRANSPORT":           9,
		"VISION_TYPE_REPLACE_DIE":         10,
		"VISION_TYPE_REPLACE_NO_NOTIFY":   11,
		"VISION_TYPE_BORN":                12,
		"VISION_TYPE_PICKUP":              13,
		"VISION_TYPE_REMOVE":              14,
		"VISION_TYPE_CHANGE_COSTUME":      15,
		"VISION_TYPE_FISH_REFRESH":        16,
		"VISION_TYPE_FISH_BIG_SHOCK":      17,
		"VISION_TYPE_FISH_QTE_SUCC":       18,
		"VISION_TYPE_Unk2700_EPFKMOIPADB": 19,
	}
)

func (x VisionType) Enum() *VisionType {
	p := new(VisionType)
	*p = x
	return p
}

func (x VisionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VisionType) Descriptor() protoreflect.EnumDescriptor {
	return file_VisionType_proto_enumTypes[0].Descriptor()
}

func (VisionType) Type() protoreflect.EnumType {
	return &file_VisionType_proto_enumTypes[0]
}

func (x VisionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VisionType.Descriptor instead.
func (VisionType) EnumDescriptor() ([]byte, []int) {
	return file_VisionType_proto_rawDescGZIP(), []int{0}
}

var File_VisionType_proto protoreflect.FileDescriptor

var file_VisionType_proto_rawDesc = []byte{
	0x0a, 0x10, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0xb0, 0x04, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x49, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x45, 0x54, 0x10, 0x01, 0x12, 0x16, 0x0a,
	0x12, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x42,
	0x4f, 0x52, 0x4e, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x03, 0x12, 0x1f,
	0x0a, 0x1b, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x41,
	0x59, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x42, 0x4f, 0x52, 0x4e, 0x10, 0x04, 0x12,
	0x14, 0x0a, 0x10, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d,
	0x49, 0x53, 0x53, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x49, 0x45, 0x10, 0x06, 0x12, 0x1d, 0x0a, 0x19, 0x56, 0x49,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x41, 0x54, 0x48, 0x45, 0x52,
	0x5f, 0x45, 0x53, 0x43, 0x41, 0x50, 0x45, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x56, 0x49, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48,
	0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x09, 0x12, 0x1b, 0x0a,
	0x17, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x5f, 0x44, 0x49, 0x45, 0x10, 0x0a, 0x12, 0x21, 0x0a, 0x1d, 0x56, 0x49,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43,
	0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x10, 0x0b, 0x12, 0x14, 0x0a,
	0x10, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x52,
	0x4e, 0x10, 0x0c, 0x12, 0x16, 0x0a, 0x12, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x50, 0x49, 0x43, 0x4b, 0x55, 0x50, 0x10, 0x0d, 0x12, 0x16, 0x0a, 0x12, 0x56,
	0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56,
	0x45, 0x10, 0x0e, 0x12, 0x1e, 0x0a, 0x1a, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x53, 0x54, 0x55, 0x4d,
	0x45, 0x10, 0x0f, 0x12, 0x1c, 0x0a, 0x18, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x46, 0x49, 0x53, 0x48, 0x5f, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x10,
	0x10, 0x12, 0x1e, 0x0a, 0x1a, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x46, 0x49, 0x53, 0x48, 0x5f, 0x42, 0x49, 0x47, 0x5f, 0x53, 0x48, 0x4f, 0x43, 0x4b, 0x10,
	0x11, 0x12, 0x1d, 0x0a, 0x19, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x46, 0x49, 0x53, 0x48, 0x5f, 0x51, 0x54, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x10, 0x12,
	0x12, 0x23, 0x0a, 0x1f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x45, 0x50, 0x46, 0x4b, 0x4d, 0x4f, 0x49, 0x50,
	0x41, 0x44, 0x42, 0x10, 0x13, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_VisionType_proto_rawDescOnce sync.Once
	file_VisionType_proto_rawDescData = file_VisionType_proto_rawDesc
)

func file_VisionType_proto_rawDescGZIP() []byte {
	file_VisionType_proto_rawDescOnce.Do(func() {
		file_VisionType_proto_rawDescData = protoimpl.X.CompressGZIP(file_VisionType_proto_rawDescData)
	})
	return file_VisionType_proto_rawDescData
}

var file_VisionType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_VisionType_proto_goTypes = []interface{}{
	(VisionType)(0), // 0: VisionType
}
var file_VisionType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_VisionType_proto_init() }
func file_VisionType_proto_init() {
	if File_VisionType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_VisionType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VisionType_proto_goTypes,
		DependencyIndexes: file_VisionType_proto_depIdxs,
		EnumInfos:         file_VisionType_proto_enumTypes,
	}.Build()
	File_VisionType_proto = out.File
	file_VisionType_proto_rawDesc = nil
	file_VisionType_proto_goTypes = nil
	file_VisionType_proto_depIdxs = nil
}
