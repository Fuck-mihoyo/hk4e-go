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
// source: Equip.proto

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

type Equip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsLocked bool `protobuf:"varint,3,opt,name=is_locked,json=isLocked,proto3" json:"is_locked,omitempty"`
	// Types that are assignable to Detail:
	//
	//	*Equip_Reliquary
	//	*Equip_Weapon
	Detail isEquip_Detail `protobuf_oneof:"detail"`
}

func (x *Equip) Reset() {
	*x = Equip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Equip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Equip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Equip) ProtoMessage() {}

func (x *Equip) ProtoReflect() protoreflect.Message {
	mi := &file_Equip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Equip.ProtoReflect.Descriptor instead.
func (*Equip) Descriptor() ([]byte, []int) {
	return file_Equip_proto_rawDescGZIP(), []int{0}
}

func (x *Equip) GetIsLocked() bool {
	if x != nil {
		return x.IsLocked
	}
	return false
}

func (m *Equip) GetDetail() isEquip_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (x *Equip) GetReliquary() *Reliquary {
	if x, ok := x.GetDetail().(*Equip_Reliquary); ok {
		return x.Reliquary
	}
	return nil
}

func (x *Equip) GetWeapon() *Weapon {
	if x, ok := x.GetDetail().(*Equip_Weapon); ok {
		return x.Weapon
	}
	return nil
}

type isEquip_Detail interface {
	isEquip_Detail()
}

type Equip_Reliquary struct {
	Reliquary *Reliquary `protobuf:"bytes,1,opt,name=reliquary,proto3,oneof"`
}

type Equip_Weapon struct {
	Weapon *Weapon `protobuf:"bytes,2,opt,name=weapon,proto3,oneof"`
}

func (*Equip_Reliquary) isEquip_Detail() {}

func (*Equip_Weapon) isEquip_Detail() {}

var File_Equip_proto protoreflect.FileDescriptor

var file_Equip_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x45, 0x71, 0x75, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x05, 0x45, 0x71, 0x75, 0x69, 0x70, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x09, 0x72, 0x65,
	0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x48,
	0x00, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x06,
	0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x77,
	0x65, 0x61, 0x70, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_Equip_proto_rawDescOnce sync.Once
	file_Equip_proto_rawDescData = file_Equip_proto_rawDesc
)

func file_Equip_proto_rawDescGZIP() []byte {
	file_Equip_proto_rawDescOnce.Do(func() {
		file_Equip_proto_rawDescData = protoimpl.X.CompressGZIP(file_Equip_proto_rawDescData)
	})
	return file_Equip_proto_rawDescData
}

var file_Equip_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Equip_proto_goTypes = []interface{}{
	(*Equip)(nil),     // 0: proto.Equip
	(*Reliquary)(nil), // 1: proto.Reliquary
	(*Weapon)(nil),    // 2: proto.Weapon
}
var file_Equip_proto_depIdxs = []int32{
	1, // 0: proto.Equip.reliquary:type_name -> proto.Reliquary
	2, // 1: proto.Equip.weapon:type_name -> proto.Weapon
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Equip_proto_init() }
func file_Equip_proto_init() {
	if File_Equip_proto != nil {
		return
	}
	file_Reliquary_proto_init()
	file_Weapon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Equip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Equip); i {
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
	file_Equip_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Equip_Reliquary)(nil),
		(*Equip_Weapon)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Equip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Equip_proto_goTypes,
		DependencyIndexes: file_Equip_proto_depIdxs,
		MessageInfos:      file_Equip_proto_msgTypes,
	}.Build()
	File_Equip_proto = out.File
	file_Equip_proto_rawDesc = nil
	file_Equip_proto_goTypes = nil
	file_Equip_proto_depIdxs = nil
}
