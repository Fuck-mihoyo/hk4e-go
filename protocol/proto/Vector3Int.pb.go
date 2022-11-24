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
// source: Vector3Int.proto

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

type Vector3Int struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z int32 `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Vector3Int) Reset() {
	*x = Vector3Int{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Vector3Int_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector3Int) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector3Int) ProtoMessage() {}

func (x *Vector3Int) ProtoReflect() protoreflect.Message {
	mi := &file_Vector3Int_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector3Int.ProtoReflect.Descriptor instead.
func (*Vector3Int) Descriptor() ([]byte, []int) {
	return file_Vector3Int_proto_rawDescGZIP(), []int{0}
}

func (x *Vector3Int) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector3Int) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vector3Int) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

var File_Vector3Int_proto protoreflect.FileDescriptor

var file_Vector3Int_proto_rawDesc = []byte{
	0x0a, 0x10, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0a, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x7a, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Vector3Int_proto_rawDescOnce sync.Once
	file_Vector3Int_proto_rawDescData = file_Vector3Int_proto_rawDesc
)

func file_Vector3Int_proto_rawDescGZIP() []byte {
	file_Vector3Int_proto_rawDescOnce.Do(func() {
		file_Vector3Int_proto_rawDescData = protoimpl.X.CompressGZIP(file_Vector3Int_proto_rawDescData)
	})
	return file_Vector3Int_proto_rawDescData
}

var file_Vector3Int_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Vector3Int_proto_goTypes = []interface{}{
	(*Vector3Int)(nil), // 0: proto.Vector3Int
}
var file_Vector3Int_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Vector3Int_proto_init() }
func file_Vector3Int_proto_init() {
	if File_Vector3Int_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Vector3Int_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vector3Int); i {
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
			RawDescriptor: file_Vector3Int_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Vector3Int_proto_goTypes,
		DependencyIndexes: file_Vector3Int_proto_depIdxs,
		MessageInfos:      file_Vector3Int_proto_msgTypes,
	}.Build()
	File_Vector3Int_proto = out.File
	file_Vector3Int_proto_rawDesc = nil
	file_Vector3Int_proto_goTypes = nil
	file_Vector3Int_proto_depIdxs = nil
}
