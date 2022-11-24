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
// source: HomeMarkPointFurnitureData.proto

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

type HomeMarkPointFurnitureData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid          uint32  `protobuf:"varint,1,opt,name=guid,proto3" json:"guid,omitempty"`
	FurnitureId   uint32  `protobuf:"varint,2,opt,name=furniture_id,json=furnitureId,proto3" json:"furniture_id,omitempty"`
	FurnitureType uint32  `protobuf:"varint,3,opt,name=furniture_type,json=furnitureType,proto3" json:"furniture_type,omitempty"`
	Pos           *Vector `protobuf:"bytes,4,opt,name=pos,proto3" json:"pos,omitempty"`
	// Types that are assignable to Extra:
	//
	//	*HomeMarkPointFurnitureData_NpcData
	//	*HomeMarkPointFurnitureData_SuiteData
	Extra isHomeMarkPointFurnitureData_Extra `protobuf_oneof:"extra"`
}

func (x *HomeMarkPointFurnitureData) Reset() {
	*x = HomeMarkPointFurnitureData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeMarkPointFurnitureData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeMarkPointFurnitureData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeMarkPointFurnitureData) ProtoMessage() {}

func (x *HomeMarkPointFurnitureData) ProtoReflect() protoreflect.Message {
	mi := &file_HomeMarkPointFurnitureData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeMarkPointFurnitureData.ProtoReflect.Descriptor instead.
func (*HomeMarkPointFurnitureData) Descriptor() ([]byte, []int) {
	return file_HomeMarkPointFurnitureData_proto_rawDescGZIP(), []int{0}
}

func (x *HomeMarkPointFurnitureData) GetGuid() uint32 {
	if x != nil {
		return x.Guid
	}
	return 0
}

func (x *HomeMarkPointFurnitureData) GetFurnitureId() uint32 {
	if x != nil {
		return x.FurnitureId
	}
	return 0
}

func (x *HomeMarkPointFurnitureData) GetFurnitureType() uint32 {
	if x != nil {
		return x.FurnitureType
	}
	return 0
}

func (x *HomeMarkPointFurnitureData) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (m *HomeMarkPointFurnitureData) GetExtra() isHomeMarkPointFurnitureData_Extra {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (x *HomeMarkPointFurnitureData) GetNpcData() *HomeMarkPointNPCData {
	if x, ok := x.GetExtra().(*HomeMarkPointFurnitureData_NpcData); ok {
		return x.NpcData
	}
	return nil
}

func (x *HomeMarkPointFurnitureData) GetSuiteData() *HomeMarkPointSuiteData {
	if x, ok := x.GetExtra().(*HomeMarkPointFurnitureData_SuiteData); ok {
		return x.SuiteData
	}
	return nil
}

type isHomeMarkPointFurnitureData_Extra interface {
	isHomeMarkPointFurnitureData_Extra()
}

type HomeMarkPointFurnitureData_NpcData struct {
	NpcData *HomeMarkPointNPCData `protobuf:"bytes,6,opt,name=npc_data,json=npcData,proto3,oneof"`
}

type HomeMarkPointFurnitureData_SuiteData struct {
	SuiteData *HomeMarkPointSuiteData `protobuf:"bytes,7,opt,name=suite_data,json=suiteData,proto3,oneof"`
}

func (*HomeMarkPointFurnitureData_NpcData) isHomeMarkPointFurnitureData_Extra() {}

func (*HomeMarkPointFurnitureData_SuiteData) isHomeMarkPointFurnitureData_Extra() {}

var File_HomeMarkPointFurnitureData_proto protoreflect.FileDescriptor

var file_HomeMarkPointFurnitureData_proto_rawDesc = []byte{
	0x0a, 0x20, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x46,
	0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x48, 0x6f, 0x6d, 0x65, 0x4d,
	0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4e, 0x50, 0x43, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9e, 0x02, 0x0a, 0x1a, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x67, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x75, 0x72, 0x6e,
	0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x75, 0x72, 0x6e, 0x69,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f,
	0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12,
	0x38, 0x0a, 0x08, 0x6e, 0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x61,
	0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4e, 0x50, 0x43, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00,
	0x52, 0x07, 0x6e, 0x70, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x75, 0x69,
	0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x09,
	0x73, 0x75, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeMarkPointFurnitureData_proto_rawDescOnce sync.Once
	file_HomeMarkPointFurnitureData_proto_rawDescData = file_HomeMarkPointFurnitureData_proto_rawDesc
)

func file_HomeMarkPointFurnitureData_proto_rawDescGZIP() []byte {
	file_HomeMarkPointFurnitureData_proto_rawDescOnce.Do(func() {
		file_HomeMarkPointFurnitureData_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeMarkPointFurnitureData_proto_rawDescData)
	})
	return file_HomeMarkPointFurnitureData_proto_rawDescData
}

var file_HomeMarkPointFurnitureData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeMarkPointFurnitureData_proto_goTypes = []interface{}{
	(*HomeMarkPointFurnitureData)(nil), // 0: proto.HomeMarkPointFurnitureData
	(*Vector)(nil),                     // 1: proto.Vector
	(*HomeMarkPointNPCData)(nil),       // 2: proto.HomeMarkPointNPCData
	(*HomeMarkPointSuiteData)(nil),     // 3: proto.HomeMarkPointSuiteData
}
var file_HomeMarkPointFurnitureData_proto_depIdxs = []int32{
	1, // 0: proto.HomeMarkPointFurnitureData.pos:type_name -> proto.Vector
	2, // 1: proto.HomeMarkPointFurnitureData.npc_data:type_name -> proto.HomeMarkPointNPCData
	3, // 2: proto.HomeMarkPointFurnitureData.suite_data:type_name -> proto.HomeMarkPointSuiteData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_HomeMarkPointFurnitureData_proto_init() }
func file_HomeMarkPointFurnitureData_proto_init() {
	if File_HomeMarkPointFurnitureData_proto != nil {
		return
	}
	file_HomeMarkPointNPCData_proto_init()
	file_HomeMarkPointSuiteData_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeMarkPointFurnitureData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeMarkPointFurnitureData); i {
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
	file_HomeMarkPointFurnitureData_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HomeMarkPointFurnitureData_NpcData)(nil),
		(*HomeMarkPointFurnitureData_SuiteData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HomeMarkPointFurnitureData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeMarkPointFurnitureData_proto_goTypes,
		DependencyIndexes: file_HomeMarkPointFurnitureData_proto_depIdxs,
		MessageInfos:      file_HomeMarkPointFurnitureData_proto_msgTypes,
	}.Build()
	File_HomeMarkPointFurnitureData_proto = out.File
	file_HomeMarkPointFurnitureData_proto_rawDesc = nil
	file_HomeMarkPointFurnitureData_proto_goTypes = nil
	file_HomeMarkPointFurnitureData_proto_depIdxs = nil
}
