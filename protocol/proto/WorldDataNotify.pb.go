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
// source: WorldDataNotify.proto

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

type WorldDataNotify_DataType int32

const (
	WorldDataNotify_DATA_TYPE_NONE          WorldDataNotify_DataType = 0
	WorldDataNotify_DATA_TYPE_WORLD_LEVEL   WorldDataNotify_DataType = 1
	WorldDataNotify_DATA_TYPE_IS_IN_MP_MODE WorldDataNotify_DataType = 2
)

// Enum value maps for WorldDataNotify_DataType.
var (
	WorldDataNotify_DataType_name = map[int32]string{
		0: "DATA_TYPE_NONE",
		1: "DATA_TYPE_WORLD_LEVEL",
		2: "DATA_TYPE_IS_IN_MP_MODE",
	}
	WorldDataNotify_DataType_value = map[string]int32{
		"DATA_TYPE_NONE":          0,
		"DATA_TYPE_WORLD_LEVEL":   1,
		"DATA_TYPE_IS_IN_MP_MODE": 2,
	}
)

func (x WorldDataNotify_DataType) Enum() *WorldDataNotify_DataType {
	p := new(WorldDataNotify_DataType)
	*p = x
	return p
}

func (x WorldDataNotify_DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorldDataNotify_DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_WorldDataNotify_proto_enumTypes[0].Descriptor()
}

func (WorldDataNotify_DataType) Type() protoreflect.EnumType {
	return &file_WorldDataNotify_proto_enumTypes[0]
}

func (x WorldDataNotify_DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorldDataNotify_DataType.Descriptor instead.
func (WorldDataNotify_DataType) EnumDescriptor() ([]byte, []int) {
	return file_WorldDataNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 3308
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type WorldDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorldPropMap map[uint32]*PropValue `protobuf:"bytes,9,rep,name=world_prop_map,json=worldPropMap,proto3" json:"world_prop_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WorldDataNotify) Reset() {
	*x = WorldDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldDataNotify) ProtoMessage() {}

func (x *WorldDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WorldDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldDataNotify.ProtoReflect.Descriptor instead.
func (*WorldDataNotify) Descriptor() ([]byte, []int) {
	return file_WorldDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WorldDataNotify) GetWorldPropMap() map[uint32]*PropValue {
	if x != nil {
		return x.WorldPropMap
	}
	return nil
}

var File_WorldDataNotify_proto protoreflect.FileDescriptor

var file_WorldDataNotify_proto_rawDesc = []byte{
	0x0a, 0x15, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x50, 0x72, 0x6f, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8c, 0x02, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x4e, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x70, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x70,
	0x4d, 0x61, 0x70, 0x1a, 0x51, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x70,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x56, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10,
	0x01, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49,
	0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x4d, 0x50, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10, 0x02, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_WorldDataNotify_proto_rawDescOnce sync.Once
	file_WorldDataNotify_proto_rawDescData = file_WorldDataNotify_proto_rawDesc
)

func file_WorldDataNotify_proto_rawDescGZIP() []byte {
	file_WorldDataNotify_proto_rawDescOnce.Do(func() {
		file_WorldDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WorldDataNotify_proto_rawDescData)
	})
	return file_WorldDataNotify_proto_rawDescData
}

var file_WorldDataNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_WorldDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_WorldDataNotify_proto_goTypes = []interface{}{
	(WorldDataNotify_DataType)(0), // 0: proto.WorldDataNotify.DataType
	(*WorldDataNotify)(nil),       // 1: proto.WorldDataNotify
	nil,                           // 2: proto.WorldDataNotify.WorldPropMapEntry
	(*PropValue)(nil),             // 3: proto.PropValue
}
var file_WorldDataNotify_proto_depIdxs = []int32{
	2, // 0: proto.WorldDataNotify.world_prop_map:type_name -> proto.WorldDataNotify.WorldPropMapEntry
	3, // 1: proto.WorldDataNotify.WorldPropMapEntry.value:type_name -> proto.PropValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_WorldDataNotify_proto_init() }
func file_WorldDataNotify_proto_init() {
	if File_WorldDataNotify_proto != nil {
		return
	}
	file_PropValue_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WorldDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldDataNotify); i {
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
			RawDescriptor: file_WorldDataNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WorldDataNotify_proto_goTypes,
		DependencyIndexes: file_WorldDataNotify_proto_depIdxs,
		EnumInfos:         file_WorldDataNotify_proto_enumTypes,
		MessageInfos:      file_WorldDataNotify_proto_msgTypes,
	}.Build()
	File_WorldDataNotify_proto = out.File
	file_WorldDataNotify_proto_rawDesc = nil
	file_WorldDataNotify_proto_goTypes = nil
	file_WorldDataNotify_proto_depIdxs = nil
}
