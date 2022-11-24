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
// source: LuaSetOptionNotify.proto

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

type LuaSetOptionNotify_LuaOptionType int32

const (
	LuaSetOptionNotify_LUA_OPTION_TYPE_NONE         LuaSetOptionNotify_LuaOptionType = 0
	LuaSetOptionNotify_LUA_OPTION_TYPE_PLAYER_INPUT LuaSetOptionNotify_LuaOptionType = 1
)

// Enum value maps for LuaSetOptionNotify_LuaOptionType.
var (
	LuaSetOptionNotify_LuaOptionType_name = map[int32]string{
		0: "LUA_OPTION_TYPE_NONE",
		1: "LUA_OPTION_TYPE_PLAYER_INPUT",
	}
	LuaSetOptionNotify_LuaOptionType_value = map[string]int32{
		"LUA_OPTION_TYPE_NONE":         0,
		"LUA_OPTION_TYPE_PLAYER_INPUT": 1,
	}
)

func (x LuaSetOptionNotify_LuaOptionType) Enum() *LuaSetOptionNotify_LuaOptionType {
	p := new(LuaSetOptionNotify_LuaOptionType)
	*p = x
	return p
}

func (x LuaSetOptionNotify_LuaOptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LuaSetOptionNotify_LuaOptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_LuaSetOptionNotify_proto_enumTypes[0].Descriptor()
}

func (LuaSetOptionNotify_LuaOptionType) Type() protoreflect.EnumType {
	return &file_LuaSetOptionNotify_proto_enumTypes[0]
}

func (x LuaSetOptionNotify_LuaOptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LuaSetOptionNotify_LuaOptionType.Descriptor instead.
func (LuaSetOptionNotify_LuaOptionType) EnumDescriptor() ([]byte, []int) {
	return file_LuaSetOptionNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 316
// EnetChannelId: 0
// EnetIsReliable: true
type LuaSetOptionNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LuaSetParam string                           `protobuf:"bytes,8,opt,name=lua_set_param,json=luaSetParam,proto3" json:"lua_set_param,omitempty"`
	OptionType  LuaSetOptionNotify_LuaOptionType `protobuf:"varint,10,opt,name=option_type,json=optionType,proto3,enum=proto.LuaSetOptionNotify_LuaOptionType" json:"option_type,omitempty"`
}

func (x *LuaSetOptionNotify) Reset() {
	*x = LuaSetOptionNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LuaSetOptionNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LuaSetOptionNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LuaSetOptionNotify) ProtoMessage() {}

func (x *LuaSetOptionNotify) ProtoReflect() protoreflect.Message {
	mi := &file_LuaSetOptionNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LuaSetOptionNotify.ProtoReflect.Descriptor instead.
func (*LuaSetOptionNotify) Descriptor() ([]byte, []int) {
	return file_LuaSetOptionNotify_proto_rawDescGZIP(), []int{0}
}

func (x *LuaSetOptionNotify) GetLuaSetParam() string {
	if x != nil {
		return x.LuaSetParam
	}
	return ""
}

func (x *LuaSetOptionNotify) GetOptionType() LuaSetOptionNotify_LuaOptionType {
	if x != nil {
		return x.OptionType
	}
	return LuaSetOptionNotify_LUA_OPTION_TYPE_NONE
}

var File_LuaSetOptionNotify_proto protoreflect.FileDescriptor

var file_LuaSetOptionNotify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4c, 0x75, 0x61, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x12, 0x4c, 0x75, 0x61, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x75, 0x61, 0x5f,
	0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6c, 0x75, 0x61, 0x53, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x48, 0x0a, 0x0b,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x75, 0x61, 0x53, 0x65, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x4c, 0x75, 0x61,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4b, 0x0a, 0x0d, 0x4c, 0x75, 0x61, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x55, 0x41, 0x5f, 0x4f,
	0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x20, 0x0a, 0x1c, 0x4c, 0x55, 0x41, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x50, 0x55,
	0x54, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LuaSetOptionNotify_proto_rawDescOnce sync.Once
	file_LuaSetOptionNotify_proto_rawDescData = file_LuaSetOptionNotify_proto_rawDesc
)

func file_LuaSetOptionNotify_proto_rawDescGZIP() []byte {
	file_LuaSetOptionNotify_proto_rawDescOnce.Do(func() {
		file_LuaSetOptionNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_LuaSetOptionNotify_proto_rawDescData)
	})
	return file_LuaSetOptionNotify_proto_rawDescData
}

var file_LuaSetOptionNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_LuaSetOptionNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LuaSetOptionNotify_proto_goTypes = []interface{}{
	(LuaSetOptionNotify_LuaOptionType)(0), // 0: proto.LuaSetOptionNotify.LuaOptionType
	(*LuaSetOptionNotify)(nil),            // 1: proto.LuaSetOptionNotify
}
var file_LuaSetOptionNotify_proto_depIdxs = []int32{
	0, // 0: proto.LuaSetOptionNotify.option_type:type_name -> proto.LuaSetOptionNotify.LuaOptionType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_LuaSetOptionNotify_proto_init() }
func file_LuaSetOptionNotify_proto_init() {
	if File_LuaSetOptionNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_LuaSetOptionNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LuaSetOptionNotify); i {
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
			RawDescriptor: file_LuaSetOptionNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LuaSetOptionNotify_proto_goTypes,
		DependencyIndexes: file_LuaSetOptionNotify_proto_depIdxs,
		EnumInfos:         file_LuaSetOptionNotify_proto_enumTypes,
		MessageInfos:      file_LuaSetOptionNotify_proto_msgTypes,
	}.Build()
	File_LuaSetOptionNotify_proto = out.File
	file_LuaSetOptionNotify_proto_rawDesc = nil
	file_LuaSetOptionNotify_proto_goTypes = nil
	file_LuaSetOptionNotify_proto_depIdxs = nil
}
