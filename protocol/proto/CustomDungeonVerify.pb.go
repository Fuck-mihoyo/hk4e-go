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
// source: CustomDungeonVerify.proto

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

type CustomDungeonVerify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DungeonGuid uint64 `protobuf:"varint,3,opt,name=dungeon_guid,json=dungeonGuid,proto3" json:"dungeon_guid,omitempty"`
	Uid         uint32 `protobuf:"varint,15,opt,name=uid,proto3" json:"uid,omitempty"`
	Timestamp   uint32 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Region      string `protobuf:"bytes,11,opt,name=region,proto3" json:"region,omitempty"`
	Lang        uint32 `protobuf:"varint,13,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *CustomDungeonVerify) Reset() {
	*x = CustomDungeonVerify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CustomDungeonVerify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomDungeonVerify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomDungeonVerify) ProtoMessage() {}

func (x *CustomDungeonVerify) ProtoReflect() protoreflect.Message {
	mi := &file_CustomDungeonVerify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomDungeonVerify.ProtoReflect.Descriptor instead.
func (*CustomDungeonVerify) Descriptor() ([]byte, []int) {
	return file_CustomDungeonVerify_proto_rawDescGZIP(), []int{0}
}

func (x *CustomDungeonVerify) GetDungeonGuid() uint64 {
	if x != nil {
		return x.DungeonGuid
	}
	return 0
}

func (x *CustomDungeonVerify) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *CustomDungeonVerify) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *CustomDungeonVerify) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CustomDungeonVerify) GetLang() uint32 {
	if x != nil {
		return x.Lang
	}
	return 0
}

var File_CustomDungeonVerify_proto protoreflect.FileDescriptor

var file_CustomDungeonVerify_proto_rawDesc = []byte{
	0x0a, 0x19, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x75,
	0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x47, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CustomDungeonVerify_proto_rawDescOnce sync.Once
	file_CustomDungeonVerify_proto_rawDescData = file_CustomDungeonVerify_proto_rawDesc
)

func file_CustomDungeonVerify_proto_rawDescGZIP() []byte {
	file_CustomDungeonVerify_proto_rawDescOnce.Do(func() {
		file_CustomDungeonVerify_proto_rawDescData = protoimpl.X.CompressGZIP(file_CustomDungeonVerify_proto_rawDescData)
	})
	return file_CustomDungeonVerify_proto_rawDescData
}

var file_CustomDungeonVerify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CustomDungeonVerify_proto_goTypes = []interface{}{
	(*CustomDungeonVerify)(nil), // 0: proto.CustomDungeonVerify
}
var file_CustomDungeonVerify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CustomDungeonVerify_proto_init() }
func file_CustomDungeonVerify_proto_init() {
	if File_CustomDungeonVerify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CustomDungeonVerify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomDungeonVerify); i {
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
			RawDescriptor: file_CustomDungeonVerify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CustomDungeonVerify_proto_goTypes,
		DependencyIndexes: file_CustomDungeonVerify_proto_depIdxs,
		MessageInfos:      file_CustomDungeonVerify_proto_msgTypes,
	}.Build()
	File_CustomDungeonVerify_proto = out.File
	file_CustomDungeonVerify_proto_rawDesc = nil
	file_CustomDungeonVerify_proto_goTypes = nil
	file_CustomDungeonVerify_proto_depIdxs = nil
}
