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
// source: HomeBlockArrangementInfo.proto

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

type HomeBlockArrangementInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsUnlocked              bool                      `protobuf:"varint,1,opt,name=is_unlocked,json=isUnlocked,proto3" json:"is_unlocked,omitempty"`
	ComfortValue            uint32                    `protobuf:"varint,2,opt,name=comfort_value,json=comfortValue,proto3" json:"comfort_value,omitempty"`
	DeployAnimalList        []*HomeAnimalData         `protobuf:"bytes,4,rep,name=deploy_animal_list,json=deployAnimalList,proto3" json:"deploy_animal_list,omitempty"`
	Unk2700_HGIECHILOJL     []*Unk2700_GOHMLAFNBGF    `protobuf:"bytes,5,rep,name=Unk2700_HGIECHILOJL,json=Unk2700HGIECHILOJL,proto3" json:"Unk2700_HGIECHILOJL,omitempty"`
	WeekendDjinnInfoList    []*WeekendDjinnInfo       `protobuf:"bytes,13,rep,name=weekend_djinn_info_list,json=weekendDjinnInfoList,proto3" json:"weekend_djinn_info_list,omitempty"`
	FurnitureSuiteList      []*HomeFurnitureSuiteData `protobuf:"bytes,15,rep,name=furniture_suite_list,json=furnitureSuiteList,proto3" json:"furniture_suite_list,omitempty"`
	FieldList               []*HomeBlockFieldData     `protobuf:"bytes,3,rep,name=field_list,json=fieldList,proto3" json:"field_list,omitempty"`
	DeployNpcList           []*HomeNpcData            `protobuf:"bytes,11,rep,name=deploy_npc_list,json=deployNpcList,proto3" json:"deploy_npc_list,omitempty"`
	DotPatternList          []*HomeBlockDotPattern    `protobuf:"bytes,7,rep,name=dot_pattern_list,json=dotPatternList,proto3" json:"dot_pattern_list,omitempty"`
	PersistentFurnitureList []*HomeFurnitureData      `protobuf:"bytes,9,rep,name=persistent_furniture_list,json=persistentFurnitureList,proto3" json:"persistent_furniture_list,omitempty"`
	DeployFurniureList      []*HomeFurnitureData      `protobuf:"bytes,12,rep,name=deploy_furniure_list,json=deployFurniureList,proto3" json:"deploy_furniure_list,omitempty"`
	BlockId                 uint32                    `protobuf:"varint,6,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Unk2700_KJGLLEEHBPF     []*Unk2700_BIEMCDLIFOD    `protobuf:"bytes,14,rep,name=Unk2700_KJGLLEEHBPF,json=Unk2700KJGLLEEHBPF,proto3" json:"Unk2700_KJGLLEEHBPF,omitempty"`
}

func (x *HomeBlockArrangementInfo) Reset() {
	*x = HomeBlockArrangementInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeBlockArrangementInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeBlockArrangementInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeBlockArrangementInfo) ProtoMessage() {}

func (x *HomeBlockArrangementInfo) ProtoReflect() protoreflect.Message {
	mi := &file_HomeBlockArrangementInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeBlockArrangementInfo.ProtoReflect.Descriptor instead.
func (*HomeBlockArrangementInfo) Descriptor() ([]byte, []int) {
	return file_HomeBlockArrangementInfo_proto_rawDescGZIP(), []int{0}
}

func (x *HomeBlockArrangementInfo) GetIsUnlocked() bool {
	if x != nil {
		return x.IsUnlocked
	}
	return false
}

func (x *HomeBlockArrangementInfo) GetComfortValue() uint32 {
	if x != nil {
		return x.ComfortValue
	}
	return 0
}

func (x *HomeBlockArrangementInfo) GetDeployAnimalList() []*HomeAnimalData {
	if x != nil {
		return x.DeployAnimalList
	}
	return nil
}

func (x *HomeBlockArrangementInfo) GetUnk2700_HGIECHILOJL() []*Unk2700_GOHMLAFNBGF {
	if x != nil {
		return x.Unk2700_HGIECHILOJL
	}
	return nil
}

func (x *HomeBlockArrangementInfo) GetWeekendDjinnInfoList() []*WeekendDjinnInfo {
	if x != nil {
		return x.WeekendDjinnInfoList
	}
	return nil
}

func (x *HomeBlockArrangementInfo) GetFurnitureSuiteList() []*HomeFurnitureSuiteData {
	if x != nil {
		return x.FurnitureSuiteList
	}
	return nil
}

func (x *HomeBlockArrangementInfo) GetFieldList() []*HomeBlockFieldData {
	if x != nil {
		return x.FieldList
	}
	return nil
}

func (x *HomeBlockArrangementInfo) GetDeployNpcList() []*HomeNpcData {
	if x != nil {
		return x.DeployNpcList
	}
	return nil
}

func (x *HomeBlockArrangementInfo) GetDotPatternList() []*HomeBlockDotPattern {
	if x != nil {
		return x.DotPatternList
	}
	return nil
}

func (x *HomeBlockArrangementInfo) GetPersistentFurnitureList() []*HomeFurnitureData {
	if x != nil {
		return x.PersistentFurnitureList
	}
	return nil
}

func (x *HomeBlockArrangementInfo) GetDeployFurniureList() []*HomeFurnitureData {
	if x != nil {
		return x.DeployFurniureList
	}
	return nil
}

func (x *HomeBlockArrangementInfo) GetBlockId() uint32 {
	if x != nil {
		return x.BlockId
	}
	return 0
}

func (x *HomeBlockArrangementInfo) GetUnk2700_KJGLLEEHBPF() []*Unk2700_BIEMCDLIFOD {
	if x != nil {
		return x.Unk2700_KJGLLEEHBPF
	}
	return nil
}

var File_HomeBlockArrangementInfo_proto protoreflect.FileDescriptor

var file_HomeBlockArrangementInfo_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x72, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x44, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x48, 0x6f, 0x6d,
	0x65, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x48, 0x6f, 0x6d, 0x65, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74,
	0x75, 0x72, 0x65, 0x53, 0x75, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x48, 0x6f, 0x6d, 0x65, 0x4e, 0x70, 0x63, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x42,
	0x49, 0x45, 0x4d, 0x43, 0x44, 0x4c, 0x49, 0x46, 0x4f, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x47, 0x4f, 0x48, 0x4d, 0x4c, 0x41,
	0x46, 0x4e, 0x42, 0x47, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x57, 0x65, 0x65,
	0x6b, 0x65, 0x6e, 0x64, 0x44, 0x6a, 0x69, 0x6e, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x06, 0x0a, 0x18, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x66, 0x6f, 0x72, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x66, 0x6f, 0x72,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3d, 0x0a, 0x12, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x6e, 0x69, 0x6d, 0x61,
	0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x48, 0x47, 0x49, 0x45, 0x43, 0x48, 0x49, 0x4c, 0x4f, 0x4a, 0x4c, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x47, 0x4f, 0x48,
	0x4d, 0x4c, 0x41, 0x46, 0x4e, 0x42, 0x47, 0x46, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x48, 0x47, 0x49, 0x45, 0x43, 0x48, 0x49, 0x4c, 0x4f, 0x4a, 0x4c, 0x12, 0x48, 0x0a, 0x17,
	0x77, 0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x6a, 0x69, 0x6e, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x57, 0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x44, 0x6a, 0x69, 0x6e, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x14, 0x77, 0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x44, 0x6a, 0x69, 0x6e, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x14, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x46, 0x75, 0x72, 0x6e, 0x69,
	0x74, 0x75, 0x72, 0x65, 0x53, 0x75, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x12, 0x66,
	0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x53, 0x75, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x32, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f,
	0x6e, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x4e, 0x70, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x4e, 0x70, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x10, 0x64,
	0x6f, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x44, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x0e, 0x64, 0x6f, 0x74,
	0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x19, 0x70,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x17, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x46, 0x75,
	0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x14, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x75, 0x72, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x48, 0x6f, 0x6d, 0x65,
	0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x12, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4b, 0x4a, 0x47, 0x4c, 0x4c, 0x45, 0x45, 0x48,
	0x42, 0x50, 0x46, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x42, 0x49, 0x45, 0x4d, 0x43, 0x44, 0x4c, 0x49, 0x46, 0x4f, 0x44, 0x52,
	0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4b, 0x4a, 0x47, 0x4c, 0x4c, 0x45, 0x45, 0x48,
	0x42, 0x50, 0x46, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeBlockArrangementInfo_proto_rawDescOnce sync.Once
	file_HomeBlockArrangementInfo_proto_rawDescData = file_HomeBlockArrangementInfo_proto_rawDesc
)

func file_HomeBlockArrangementInfo_proto_rawDescGZIP() []byte {
	file_HomeBlockArrangementInfo_proto_rawDescOnce.Do(func() {
		file_HomeBlockArrangementInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeBlockArrangementInfo_proto_rawDescData)
	})
	return file_HomeBlockArrangementInfo_proto_rawDescData
}

var file_HomeBlockArrangementInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeBlockArrangementInfo_proto_goTypes = []interface{}{
	(*HomeBlockArrangementInfo)(nil), // 0: HomeBlockArrangementInfo
	(*HomeAnimalData)(nil),           // 1: HomeAnimalData
	(*Unk2700_GOHMLAFNBGF)(nil),      // 2: Unk2700_GOHMLAFNBGF
	(*WeekendDjinnInfo)(nil),         // 3: WeekendDjinnInfo
	(*HomeFurnitureSuiteData)(nil),   // 4: HomeFurnitureSuiteData
	(*HomeBlockFieldData)(nil),       // 5: HomeBlockFieldData
	(*HomeNpcData)(nil),              // 6: HomeNpcData
	(*HomeBlockDotPattern)(nil),      // 7: HomeBlockDotPattern
	(*HomeFurnitureData)(nil),        // 8: HomeFurnitureData
	(*Unk2700_BIEMCDLIFOD)(nil),      // 9: Unk2700_BIEMCDLIFOD
}
var file_HomeBlockArrangementInfo_proto_depIdxs = []int32{
	1,  // 0: HomeBlockArrangementInfo.deploy_animal_list:type_name -> HomeAnimalData
	2,  // 1: HomeBlockArrangementInfo.Unk2700_HGIECHILOJL:type_name -> Unk2700_GOHMLAFNBGF
	3,  // 2: HomeBlockArrangementInfo.weekend_djinn_info_list:type_name -> WeekendDjinnInfo
	4,  // 3: HomeBlockArrangementInfo.furniture_suite_list:type_name -> HomeFurnitureSuiteData
	5,  // 4: HomeBlockArrangementInfo.field_list:type_name -> HomeBlockFieldData
	6,  // 5: HomeBlockArrangementInfo.deploy_npc_list:type_name -> HomeNpcData
	7,  // 6: HomeBlockArrangementInfo.dot_pattern_list:type_name -> HomeBlockDotPattern
	8,  // 7: HomeBlockArrangementInfo.persistent_furniture_list:type_name -> HomeFurnitureData
	8,  // 8: HomeBlockArrangementInfo.deploy_furniure_list:type_name -> HomeFurnitureData
	9,  // 9: HomeBlockArrangementInfo.Unk2700_KJGLLEEHBPF:type_name -> Unk2700_BIEMCDLIFOD
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_HomeBlockArrangementInfo_proto_init() }
func file_HomeBlockArrangementInfo_proto_init() {
	if File_HomeBlockArrangementInfo_proto != nil {
		return
	}
	file_HomeAnimalData_proto_init()
	file_HomeBlockDotPattern_proto_init()
	file_HomeBlockFieldData_proto_init()
	file_HomeFurnitureData_proto_init()
	file_HomeFurnitureSuiteData_proto_init()
	file_HomeNpcData_proto_init()
	file_Unk2700_BIEMCDLIFOD_proto_init()
	file_Unk2700_GOHMLAFNBGF_proto_init()
	file_WeekendDjinnInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeBlockArrangementInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeBlockArrangementInfo); i {
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
			RawDescriptor: file_HomeBlockArrangementInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeBlockArrangementInfo_proto_goTypes,
		DependencyIndexes: file_HomeBlockArrangementInfo_proto_depIdxs,
		MessageInfos:      file_HomeBlockArrangementInfo_proto_msgTypes,
	}.Build()
	File_HomeBlockArrangementInfo_proto = out.File
	file_HomeBlockArrangementInfo_proto_rawDesc = nil
	file_HomeBlockArrangementInfo_proto_goTypes = nil
	file_HomeBlockArrangementInfo_proto_depIdxs = nil
}
