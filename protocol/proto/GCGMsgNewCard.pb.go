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
// source: GCGMsgNewCard.proto

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

type GCGMsgNewCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Card *GCGCard `protobuf:"bytes,15,opt,name=card,proto3" json:"card,omitempty"`
}

func (x *GCGMsgNewCard) Reset() {
	*x = GCGMsgNewCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGMsgNewCard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGMsgNewCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGMsgNewCard) ProtoMessage() {}

func (x *GCGMsgNewCard) ProtoReflect() protoreflect.Message {
	mi := &file_GCGMsgNewCard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGMsgNewCard.ProtoReflect.Descriptor instead.
func (*GCGMsgNewCard) Descriptor() ([]byte, []int) {
	return file_GCGMsgNewCard_proto_rawDescGZIP(), []int{0}
}

func (x *GCGMsgNewCard) GetCard() *GCGCard {
	if x != nil {
		return x.Card
	}
	return nil
}

var File_GCGMsgNewCard_proto protoreflect.FileDescriptor

var file_GCGMsgNewCard_proto_rawDesc = []byte{
	0x0a, 0x13, 0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x4e, 0x65, 0x77, 0x43, 0x61, 0x72, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x47, 0x43,
	0x47, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x0d, 0x47,
	0x43, 0x47, 0x4d, 0x73, 0x67, 0x4e, 0x65, 0x77, 0x43, 0x61, 0x72, 0x64, 0x12, 0x22, 0x0a, 0x04,
	0x63, 0x61, 0x72, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGMsgNewCard_proto_rawDescOnce sync.Once
	file_GCGMsgNewCard_proto_rawDescData = file_GCGMsgNewCard_proto_rawDesc
)

func file_GCGMsgNewCard_proto_rawDescGZIP() []byte {
	file_GCGMsgNewCard_proto_rawDescOnce.Do(func() {
		file_GCGMsgNewCard_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGMsgNewCard_proto_rawDescData)
	})
	return file_GCGMsgNewCard_proto_rawDescData
}

var file_GCGMsgNewCard_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGMsgNewCard_proto_goTypes = []interface{}{
	(*GCGMsgNewCard)(nil), // 0: proto.GCGMsgNewCard
	(*GCGCard)(nil),       // 1: proto.GCGCard
}
var file_GCGMsgNewCard_proto_depIdxs = []int32{
	1, // 0: proto.GCGMsgNewCard.card:type_name -> proto.GCGCard
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GCGMsgNewCard_proto_init() }
func file_GCGMsgNewCard_proto_init() {
	if File_GCGMsgNewCard_proto != nil {
		return
	}
	file_GCGCard_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGMsgNewCard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGMsgNewCard); i {
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
			RawDescriptor: file_GCGMsgNewCard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGMsgNewCard_proto_goTypes,
		DependencyIndexes: file_GCGMsgNewCard_proto_depIdxs,
		MessageInfos:      file_GCGMsgNewCard_proto_msgTypes,
	}.Build()
	File_GCGMsgNewCard_proto = out.File
	file_GCGMsgNewCard_proto_rawDesc = nil
	file_GCGMsgNewCard_proto_goTypes = nil
	file_GCGMsgNewCard_proto_depIdxs = nil
}
