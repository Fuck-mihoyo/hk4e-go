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
// source: MapMarkPointType.proto

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

type MapMarkPointType int32

const (
	MapMarkPointType_MAP_MARK_POINT_TYPE_NPC        MapMarkPointType = 0
	MapMarkPointType_MAP_MARK_POINT_TYPE_QUEST      MapMarkPointType = 1
	MapMarkPointType_MAP_MARK_POINT_TYPE_SPECIAL    MapMarkPointType = 2
	MapMarkPointType_MAP_MARK_POINT_TYPE_MINE       MapMarkPointType = 3
	MapMarkPointType_MAP_MARK_POINT_TYPE_COLLECTION MapMarkPointType = 4
	MapMarkPointType_MAP_MARK_POINT_TYPE_MONSTER    MapMarkPointType = 5
	MapMarkPointType_MAP_MARK_POINT_TYPE_FISH_POOL  MapMarkPointType = 6
)

// Enum value maps for MapMarkPointType.
var (
	MapMarkPointType_name = map[int32]string{
		0: "MAP_MARK_POINT_TYPE_NPC",
		1: "MAP_MARK_POINT_TYPE_QUEST",
		2: "MAP_MARK_POINT_TYPE_SPECIAL",
		3: "MAP_MARK_POINT_TYPE_MINE",
		4: "MAP_MARK_POINT_TYPE_COLLECTION",
		5: "MAP_MARK_POINT_TYPE_MONSTER",
		6: "MAP_MARK_POINT_TYPE_FISH_POOL",
	}
	MapMarkPointType_value = map[string]int32{
		"MAP_MARK_POINT_TYPE_NPC":        0,
		"MAP_MARK_POINT_TYPE_QUEST":      1,
		"MAP_MARK_POINT_TYPE_SPECIAL":    2,
		"MAP_MARK_POINT_TYPE_MINE":       3,
		"MAP_MARK_POINT_TYPE_COLLECTION": 4,
		"MAP_MARK_POINT_TYPE_MONSTER":    5,
		"MAP_MARK_POINT_TYPE_FISH_POOL":  6,
	}
)

func (x MapMarkPointType) Enum() *MapMarkPointType {
	p := new(MapMarkPointType)
	*p = x
	return p
}

func (x MapMarkPointType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MapMarkPointType) Descriptor() protoreflect.EnumDescriptor {
	return file_MapMarkPointType_proto_enumTypes[0].Descriptor()
}

func (MapMarkPointType) Type() protoreflect.EnumType {
	return &file_MapMarkPointType_proto_enumTypes[0]
}

func (x MapMarkPointType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MapMarkPointType.Descriptor instead.
func (MapMarkPointType) EnumDescriptor() ([]byte, []int) {
	return file_MapMarkPointType_proto_rawDescGZIP(), []int{0}
}

var File_MapMarkPointType_proto protoreflect.FileDescriptor

var file_MapMarkPointType_proto_rawDesc = []byte{
	0x0a, 0x16, 0x4d, 0x61, 0x70, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0xf5, 0x01, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x41, 0x50, 0x5f, 0x4d, 0x41, 0x52, 0x4b,
	0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x50, 0x43, 0x10,
	0x00, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x41, 0x50, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x5f, 0x50, 0x4f,
	0x49, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01,
	0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x41, 0x50, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x5f, 0x50, 0x4f, 0x49,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c, 0x10,
	0x02, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x41, 0x50, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x5f, 0x50, 0x4f,
	0x49, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x45, 0x10, 0x03, 0x12,
	0x22, 0x0a, 0x1e, 0x4d, 0x41, 0x50, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x5f, 0x50, 0x4f, 0x49, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x41, 0x50, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x5f,
	0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x53, 0x54,
	0x45, 0x52, 0x10, 0x05, 0x12, 0x21, 0x0a, 0x1d, 0x4d, 0x41, 0x50, 0x5f, 0x4d, 0x41, 0x52, 0x4b,
	0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x53, 0x48,
	0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x10, 0x06, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MapMarkPointType_proto_rawDescOnce sync.Once
	file_MapMarkPointType_proto_rawDescData = file_MapMarkPointType_proto_rawDesc
)

func file_MapMarkPointType_proto_rawDescGZIP() []byte {
	file_MapMarkPointType_proto_rawDescOnce.Do(func() {
		file_MapMarkPointType_proto_rawDescData = protoimpl.X.CompressGZIP(file_MapMarkPointType_proto_rawDescData)
	})
	return file_MapMarkPointType_proto_rawDescData
}

var file_MapMarkPointType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MapMarkPointType_proto_goTypes = []interface{}{
	(MapMarkPointType)(0), // 0: proto.MapMarkPointType
}
var file_MapMarkPointType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MapMarkPointType_proto_init() }
func file_MapMarkPointType_proto_init() {
	if File_MapMarkPointType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MapMarkPointType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MapMarkPointType_proto_goTypes,
		DependencyIndexes: file_MapMarkPointType_proto_depIdxs,
		EnumInfos:         file_MapMarkPointType_proto_enumTypes,
	}.Build()
	File_MapMarkPointType_proto = out.File
	file_MapMarkPointType_proto_rawDesc = nil
	file_MapMarkPointType_proto_goTypes = nil
	file_MapMarkPointType_proto_depIdxs = nil
}
