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
// source: Unk3000_EMMKKLIECLB.proto

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

type Unk3000_EMMKKLIECLB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TreePos  *Vector `protobuf:"bytes,12,opt,name=tree_pos,json=treePos,proto3" json:"tree_pos,omitempty"`
	TreeType uint32  `protobuf:"varint,8,opt,name=tree_type,json=treeType,proto3" json:"tree_type,omitempty"`
}

func (x *Unk3000_EMMKKLIECLB) Reset() {
	*x = Unk3000_EMMKKLIECLB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk3000_EMMKKLIECLB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk3000_EMMKKLIECLB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk3000_EMMKKLIECLB) ProtoMessage() {}

func (x *Unk3000_EMMKKLIECLB) ProtoReflect() protoreflect.Message {
	mi := &file_Unk3000_EMMKKLIECLB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk3000_EMMKKLIECLB.ProtoReflect.Descriptor instead.
func (*Unk3000_EMMKKLIECLB) Descriptor() ([]byte, []int) {
	return file_Unk3000_EMMKKLIECLB_proto_rawDescGZIP(), []int{0}
}

func (x *Unk3000_EMMKKLIECLB) GetTreePos() *Vector {
	if x != nil {
		return x.TreePos
	}
	return nil
}

func (x *Unk3000_EMMKKLIECLB) GetTreeType() uint32 {
	if x != nil {
		return x.TreeType
	}
	return 0
}

var File_Unk3000_EMMKKLIECLB_proto protoreflect.FileDescriptor

var file_Unk3000_EMMKKLIECLB_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x45, 0x4d, 0x4d, 0x4b, 0x4b, 0x4c,
	0x49, 0x45, 0x43, 0x4c, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x13, 0x55, 0x6e, 0x6b,
	0x33, 0x30, 0x30, 0x30, 0x5f, 0x45, 0x4d, 0x4d, 0x4b, 0x4b, 0x4c, 0x49, 0x45, 0x43, 0x4c, 0x42,
	0x12, 0x22, 0x0a, 0x08, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x74, 0x72, 0x65,
	0x65, 0x50, 0x6f, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x72, 0x65, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk3000_EMMKKLIECLB_proto_rawDescOnce sync.Once
	file_Unk3000_EMMKKLIECLB_proto_rawDescData = file_Unk3000_EMMKKLIECLB_proto_rawDesc
)

func file_Unk3000_EMMKKLIECLB_proto_rawDescGZIP() []byte {
	file_Unk3000_EMMKKLIECLB_proto_rawDescOnce.Do(func() {
		file_Unk3000_EMMKKLIECLB_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk3000_EMMKKLIECLB_proto_rawDescData)
	})
	return file_Unk3000_EMMKKLIECLB_proto_rawDescData
}

var file_Unk3000_EMMKKLIECLB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk3000_EMMKKLIECLB_proto_goTypes = []interface{}{
	(*Unk3000_EMMKKLIECLB)(nil), // 0: Unk3000_EMMKKLIECLB
	(*Vector)(nil),              // 1: Vector
}
var file_Unk3000_EMMKKLIECLB_proto_depIdxs = []int32{
	1, // 0: Unk3000_EMMKKLIECLB.tree_pos:type_name -> Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Unk3000_EMMKKLIECLB_proto_init() }
func file_Unk3000_EMMKKLIECLB_proto_init() {
	if File_Unk3000_EMMKKLIECLB_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk3000_EMMKKLIECLB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk3000_EMMKKLIECLB); i {
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
			RawDescriptor: file_Unk3000_EMMKKLIECLB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk3000_EMMKKLIECLB_proto_goTypes,
		DependencyIndexes: file_Unk3000_EMMKKLIECLB_proto_depIdxs,
		MessageInfos:      file_Unk3000_EMMKKLIECLB_proto_msgTypes,
	}.Build()
	File_Unk3000_EMMKKLIECLB_proto = out.File
	file_Unk3000_EMMKKLIECLB_proto_rawDesc = nil
	file_Unk3000_EMMKKLIECLB_proto_goTypes = nil
	file_Unk3000_EMMKKLIECLB_proto_depIdxs = nil
}
