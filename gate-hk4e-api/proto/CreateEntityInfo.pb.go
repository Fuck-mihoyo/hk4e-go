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
// source: CreateEntityInfo.proto

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

type CreateEntityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level          uint32  `protobuf:"varint,5,opt,name=level,proto3" json:"level,omitempty"`
	Pos            *Vector `protobuf:"bytes,6,opt,name=pos,proto3" json:"pos,omitempty"`
	Rot            *Vector `protobuf:"bytes,7,opt,name=rot,proto3" json:"rot,omitempty"`
	SceneId        uint32  `protobuf:"varint,10,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	RoomId         uint32  `protobuf:"varint,11,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	ClientUniqueId uint32  `protobuf:"varint,12,opt,name=client_unique_id,json=clientUniqueId,proto3" json:"client_unique_id,omitempty"`
	// Types that are assignable to Entity:
	//	*CreateEntityInfo_MonsterId
	//	*CreateEntityInfo_NpcId
	//	*CreateEntityInfo_GadgetId
	//	*CreateEntityInfo_ItemId
	Entity isCreateEntityInfo_Entity `protobuf_oneof:"entity"`
	// Types that are assignable to EntityCreateInfo:
	//	*CreateEntityInfo_Gadget
	EntityCreateInfo isCreateEntityInfo_EntityCreateInfo `protobuf_oneof:"entity_create_info"`
}

func (x *CreateEntityInfo) Reset() {
	*x = CreateEntityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CreateEntityInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntityInfo) ProtoMessage() {}

func (x *CreateEntityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_CreateEntityInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntityInfo.ProtoReflect.Descriptor instead.
func (*CreateEntityInfo) Descriptor() ([]byte, []int) {
	return file_CreateEntityInfo_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEntityInfo) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *CreateEntityInfo) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *CreateEntityInfo) GetRot() *Vector {
	if x != nil {
		return x.Rot
	}
	return nil
}

func (x *CreateEntityInfo) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *CreateEntityInfo) GetRoomId() uint32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *CreateEntityInfo) GetClientUniqueId() uint32 {
	if x != nil {
		return x.ClientUniqueId
	}
	return 0
}

func (m *CreateEntityInfo) GetEntity() isCreateEntityInfo_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (x *CreateEntityInfo) GetMonsterId() uint32 {
	if x, ok := x.GetEntity().(*CreateEntityInfo_MonsterId); ok {
		return x.MonsterId
	}
	return 0
}

func (x *CreateEntityInfo) GetNpcId() uint32 {
	if x, ok := x.GetEntity().(*CreateEntityInfo_NpcId); ok {
		return x.NpcId
	}
	return 0
}

func (x *CreateEntityInfo) GetGadgetId() uint32 {
	if x, ok := x.GetEntity().(*CreateEntityInfo_GadgetId); ok {
		return x.GadgetId
	}
	return 0
}

func (x *CreateEntityInfo) GetItemId() uint32 {
	if x, ok := x.GetEntity().(*CreateEntityInfo_ItemId); ok {
		return x.ItemId
	}
	return 0
}

func (m *CreateEntityInfo) GetEntityCreateInfo() isCreateEntityInfo_EntityCreateInfo {
	if m != nil {
		return m.EntityCreateInfo
	}
	return nil
}

func (x *CreateEntityInfo) GetGadget() *CreateGadgetInfo {
	if x, ok := x.GetEntityCreateInfo().(*CreateEntityInfo_Gadget); ok {
		return x.Gadget
	}
	return nil
}

type isCreateEntityInfo_Entity interface {
	isCreateEntityInfo_Entity()
}

type CreateEntityInfo_MonsterId struct {
	MonsterId uint32 `protobuf:"varint,1,opt,name=monster_id,json=monsterId,proto3,oneof"`
}

type CreateEntityInfo_NpcId struct {
	NpcId uint32 `protobuf:"varint,2,opt,name=npc_id,json=npcId,proto3,oneof"`
}

type CreateEntityInfo_GadgetId struct {
	GadgetId uint32 `protobuf:"varint,3,opt,name=gadget_id,json=gadgetId,proto3,oneof"`
}

type CreateEntityInfo_ItemId struct {
	ItemId uint32 `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3,oneof"`
}

func (*CreateEntityInfo_MonsterId) isCreateEntityInfo_Entity() {}

func (*CreateEntityInfo_NpcId) isCreateEntityInfo_Entity() {}

func (*CreateEntityInfo_GadgetId) isCreateEntityInfo_Entity() {}

func (*CreateEntityInfo_ItemId) isCreateEntityInfo_Entity() {}

type isCreateEntityInfo_EntityCreateInfo interface {
	isCreateEntityInfo_EntityCreateInfo()
}

type CreateEntityInfo_Gadget struct {
	Gadget *CreateGadgetInfo `protobuf:"bytes,13,opt,name=gadget,proto3,oneof"`
}

func (*CreateEntityInfo_Gadget) isCreateEntityInfo_EntityCreateInfo() {}

var File_CreateEntityInfo_proto protoreflect.FileDescriptor

var file_CreateEntityInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd,
	0x02, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x03, 0x70, 0x6f, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x03, 0x70, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x03, 0x72, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x72, 0x6f, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0a, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x06, 0x6e, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00,
	0x52, 0x05, 0x6e, 0x70, 0x63, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x67, 0x61, 0x64, 0x67, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x08, 0x67, 0x61,
	0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x06, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x48, 0x01, 0x52, 0x06, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x42, 0x08,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x14, 0x0a, 0x12, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_CreateEntityInfo_proto_rawDescOnce sync.Once
	file_CreateEntityInfo_proto_rawDescData = file_CreateEntityInfo_proto_rawDesc
)

func file_CreateEntityInfo_proto_rawDescGZIP() []byte {
	file_CreateEntityInfo_proto_rawDescOnce.Do(func() {
		file_CreateEntityInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_CreateEntityInfo_proto_rawDescData)
	})
	return file_CreateEntityInfo_proto_rawDescData
}

var file_CreateEntityInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CreateEntityInfo_proto_goTypes = []interface{}{
	(*CreateEntityInfo)(nil), // 0: CreateEntityInfo
	(*Vector)(nil),           // 1: Vector
	(*CreateGadgetInfo)(nil), // 2: CreateGadgetInfo
}
var file_CreateEntityInfo_proto_depIdxs = []int32{
	1, // 0: CreateEntityInfo.pos:type_name -> Vector
	1, // 1: CreateEntityInfo.rot:type_name -> Vector
	2, // 2: CreateEntityInfo.gadget:type_name -> CreateGadgetInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_CreateEntityInfo_proto_init() }
func file_CreateEntityInfo_proto_init() {
	if File_CreateEntityInfo_proto != nil {
		return
	}
	file_CreateGadgetInfo_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CreateEntityInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntityInfo); i {
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
	file_CreateEntityInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CreateEntityInfo_MonsterId)(nil),
		(*CreateEntityInfo_NpcId)(nil),
		(*CreateEntityInfo_GadgetId)(nil),
		(*CreateEntityInfo_ItemId)(nil),
		(*CreateEntityInfo_Gadget)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CreateEntityInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CreateEntityInfo_proto_goTypes,
		DependencyIndexes: file_CreateEntityInfo_proto_depIdxs,
		MessageInfos:      file_CreateEntityInfo_proto_msgTypes,
	}.Build()
	File_CreateEntityInfo_proto = out.File
	file_CreateEntityInfo_proto_rawDesc = nil
	file_CreateEntityInfo_proto_goTypes = nil
	file_CreateEntityInfo_proto_depIdxs = nil
}
