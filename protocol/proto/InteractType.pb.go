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
// source: InteractType.proto

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

type InteractType int32

const (
	InteractType_INTERACT_TYPE_NONE            InteractType = 0
	InteractType_INTERACT_TYPE_PICK_ITEM       InteractType = 1
	InteractType_INTERACT_TYPE_GATHER          InteractType = 2
	InteractType_INTERACT_TYPE_OPEN_CHEST      InteractType = 3
	InteractType_INTERACT_TYPE_OPEN_STATUE     InteractType = 4
	InteractType_INTERACT_TYPE_CONSUM          InteractType = 5
	InteractType_INTERACT_TYPE_MP_PLAY_REWARD  InteractType = 6
	InteractType_INTERACT_TYPE_VIEW            InteractType = 7
	InteractType_INTERACT_TYPE_GENERAL_REWARD  InteractType = 8
	InteractType_INTERACT_TYPE_MIRACLE_RING    InteractType = 9
	InteractType_INTERACT_TYPE_FOUNDATION      InteractType = 10
	InteractType_INTERACT_TYPE_ECHO_SHELL      InteractType = 11
	InteractType_INTERACT_TYPE_HOME_GATHER     InteractType = 12
	InteractType_INTERACT_TYPE_ENV_ANIMAL      InteractType = 13
	InteractType_INTERACT_TYPE_QUEST_GADGET    InteractType = 14
	InteractType_INTERACT_TYPE_UI_INTERACT     InteractType = 15
	InteractType_INTERACT_TYPE_DESHRET_OBELISK InteractType = 16
)

// Enum value maps for InteractType.
var (
	InteractType_name = map[int32]string{
		0:  "INTERACT_TYPE_NONE",
		1:  "INTERACT_TYPE_PICK_ITEM",
		2:  "INTERACT_TYPE_GATHER",
		3:  "INTERACT_TYPE_OPEN_CHEST",
		4:  "INTERACT_TYPE_OPEN_STATUE",
		5:  "INTERACT_TYPE_CONSUM",
		6:  "INTERACT_TYPE_MP_PLAY_REWARD",
		7:  "INTERACT_TYPE_VIEW",
		8:  "INTERACT_TYPE_GENERAL_REWARD",
		9:  "INTERACT_TYPE_MIRACLE_RING",
		10: "INTERACT_TYPE_FOUNDATION",
		11: "INTERACT_TYPE_ECHO_SHELL",
		12: "INTERACT_TYPE_HOME_GATHER",
		13: "INTERACT_TYPE_ENV_ANIMAL",
		14: "INTERACT_TYPE_QUEST_GADGET",
		15: "INTERACT_TYPE_UI_INTERACT",
		16: "INTERACT_TYPE_DESHRET_OBELISK",
	}
	InteractType_value = map[string]int32{
		"INTERACT_TYPE_NONE":            0,
		"INTERACT_TYPE_PICK_ITEM":       1,
		"INTERACT_TYPE_GATHER":          2,
		"INTERACT_TYPE_OPEN_CHEST":      3,
		"INTERACT_TYPE_OPEN_STATUE":     4,
		"INTERACT_TYPE_CONSUM":          5,
		"INTERACT_TYPE_MP_PLAY_REWARD":  6,
		"INTERACT_TYPE_VIEW":            7,
		"INTERACT_TYPE_GENERAL_REWARD":  8,
		"INTERACT_TYPE_MIRACLE_RING":    9,
		"INTERACT_TYPE_FOUNDATION":      10,
		"INTERACT_TYPE_ECHO_SHELL":      11,
		"INTERACT_TYPE_HOME_GATHER":     12,
		"INTERACT_TYPE_ENV_ANIMAL":      13,
		"INTERACT_TYPE_QUEST_GADGET":    14,
		"INTERACT_TYPE_UI_INTERACT":     15,
		"INTERACT_TYPE_DESHRET_OBELISK": 16,
	}
)

func (x InteractType) Enum() *InteractType {
	p := new(InteractType)
	*p = x
	return p
}

func (x InteractType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InteractType) Descriptor() protoreflect.EnumDescriptor {
	return file_InteractType_proto_enumTypes[0].Descriptor()
}

func (InteractType) Type() protoreflect.EnumType {
	return &file_InteractType_proto_enumTypes[0]
}

func (x InteractType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InteractType.Descriptor instead.
func (InteractType) EnumDescriptor() ([]byte, []int) {
	return file_InteractType_proto_rawDescGZIP(), []int{0}
}

var File_InteractType_proto protoreflect.FileDescriptor

var file_InteractType_proto_rawDesc = []byte{
	0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x8b, 0x04, 0x0a, 0x0c,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x43, 0x4b, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x10,
	0x01, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x47, 0x41, 0x54, 0x48, 0x45, 0x52, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x50, 0x45,
	0x4e, 0x5f, 0x43, 0x48, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x45, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d,
	0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4d, 0x50, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x52, 0x45, 0x57, 0x41,
	0x52, 0x44, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x10, 0x07, 0x12, 0x20, 0x0a, 0x1c,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x45,
	0x4e, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52, 0x44, 0x10, 0x08, 0x12, 0x1e,
	0x0a, 0x1a, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4d, 0x49, 0x52, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x1c,
	0x0a, 0x18, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x1c, 0x0a, 0x18,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x43,
	0x48, 0x4f, 0x5f, 0x53, 0x48, 0x45, 0x4c, 0x4c, 0x10, 0x0b, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x4f, 0x4d, 0x45,
	0x5f, 0x47, 0x41, 0x54, 0x48, 0x45, 0x52, 0x10, 0x0c, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x56, 0x5f, 0x41,
	0x4e, 0x49, 0x4d, 0x41, 0x4c, 0x10, 0x0d, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x47,
	0x41, 0x44, 0x47, 0x45, 0x54, 0x10, 0x0e, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x49, 0x5f, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x41, 0x43, 0x54, 0x10, 0x0f, 0x12, 0x21, 0x0a, 0x1d, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x41,
	0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x53, 0x48, 0x52, 0x45, 0x54, 0x5f,
	0x4f, 0x42, 0x45, 0x4c, 0x49, 0x53, 0x4b, 0x10, 0x10, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InteractType_proto_rawDescOnce sync.Once
	file_InteractType_proto_rawDescData = file_InteractType_proto_rawDesc
)

func file_InteractType_proto_rawDescGZIP() []byte {
	file_InteractType_proto_rawDescOnce.Do(func() {
		file_InteractType_proto_rawDescData = protoimpl.X.CompressGZIP(file_InteractType_proto_rawDescData)
	})
	return file_InteractType_proto_rawDescData
}

var file_InteractType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_InteractType_proto_goTypes = []interface{}{
	(InteractType)(0), // 0: proto.InteractType
}
var file_InteractType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_InteractType_proto_init() }
func file_InteractType_proto_init() {
	if File_InteractType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_InteractType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InteractType_proto_goTypes,
		DependencyIndexes: file_InteractType_proto_depIdxs,
		EnumInfos:         file_InteractType_proto_enumTypes,
	}.Build()
	File_InteractType_proto = out.File
	file_InteractType_proto_rawDesc = nil
	file_InteractType_proto_goTypes = nil
	file_InteractType_proto_depIdxs = nil
}
