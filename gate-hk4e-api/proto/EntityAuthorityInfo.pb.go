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
// source: EntityAuthorityInfo.proto

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

type EntityAuthorityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbilityInfo         *AbilitySyncStateInfo             `protobuf:"bytes,1,opt,name=ability_info,json=abilityInfo,proto3" json:"ability_info,omitempty"`
	RendererChangedInfo *EntityRendererChangedInfo        `protobuf:"bytes,2,opt,name=renderer_changed_info,json=rendererChangedInfo,proto3" json:"renderer_changed_info,omitempty"`
	AiInfo              *SceneEntityAiInfo                `protobuf:"bytes,3,opt,name=ai_info,json=aiInfo,proto3" json:"ai_info,omitempty"`
	BornPos             *Vector                           `protobuf:"bytes,4,opt,name=born_pos,json=bornPos,proto3" json:"born_pos,omitempty"`
	PoseParaList        []*AnimatorParameterValueInfoPair `protobuf:"bytes,5,rep,name=pose_para_list,json=poseParaList,proto3" json:"pose_para_list,omitempty"`
	Unk2700_KDGMOPELHNE *Unk2700_HFMDKDHCJCM              `protobuf:"bytes,6,opt,name=Unk2700_KDGMOPELHNE,json=Unk2700KDGMOPELHNE,proto3" json:"Unk2700_KDGMOPELHNE,omitempty"`
}

func (x *EntityAuthorityInfo) Reset() {
	*x = EntityAuthorityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityAuthorityInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityAuthorityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityAuthorityInfo) ProtoMessage() {}

func (x *EntityAuthorityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_EntityAuthorityInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityAuthorityInfo.ProtoReflect.Descriptor instead.
func (*EntityAuthorityInfo) Descriptor() ([]byte, []int) {
	return file_EntityAuthorityInfo_proto_rawDescGZIP(), []int{0}
}

func (x *EntityAuthorityInfo) GetAbilityInfo() *AbilitySyncStateInfo {
	if x != nil {
		return x.AbilityInfo
	}
	return nil
}

func (x *EntityAuthorityInfo) GetRendererChangedInfo() *EntityRendererChangedInfo {
	if x != nil {
		return x.RendererChangedInfo
	}
	return nil
}

func (x *EntityAuthorityInfo) GetAiInfo() *SceneEntityAiInfo {
	if x != nil {
		return x.AiInfo
	}
	return nil
}

func (x *EntityAuthorityInfo) GetBornPos() *Vector {
	if x != nil {
		return x.BornPos
	}
	return nil
}

func (x *EntityAuthorityInfo) GetPoseParaList() []*AnimatorParameterValueInfoPair {
	if x != nil {
		return x.PoseParaList
	}
	return nil
}

func (x *EntityAuthorityInfo) GetUnk2700_KDGMOPELHNE() *Unk2700_HFMDKDHCJCM {
	if x != nil {
		return x.Unk2700_KDGMOPELHNE
	}
	return nil
}

var File_EntityAuthorityInfo_proto protoreflect.FileDescriptor

var file_EntityAuthorityInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x6f,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x69, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x48, 0x46, 0x4d, 0x44, 0x4b, 0x44, 0x48, 0x43, 0x4a, 0x43, 0x4d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfe, 0x02, 0x0a, 0x13, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0c, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x4e, 0x0a, 0x15, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x5f, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x13, 0x72,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x69, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x41, 0x69, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x61, 0x69, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x22, 0x0a, 0x08, 0x62, 0x6f, 0x72, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x62, 0x6f, 0x72, 0x6e,
	0x50, 0x6f, 0x73, 0x12, 0x45, 0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x41, 0x6e,
	0x69, 0x6d, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0c, 0x70, 0x6f,
	0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4b, 0x44, 0x47, 0x4d, 0x4f, 0x50, 0x45, 0x4c, 0x48, 0x4e,
	0x45, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x5f, 0x48, 0x46, 0x4d, 0x44, 0x4b, 0x44, 0x48, 0x43, 0x4a, 0x43, 0x4d, 0x52, 0x12, 0x55,
	0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4b, 0x44, 0x47, 0x4d, 0x4f, 0x50, 0x45, 0x4c, 0x48, 0x4e,
	0x45, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EntityAuthorityInfo_proto_rawDescOnce sync.Once
	file_EntityAuthorityInfo_proto_rawDescData = file_EntityAuthorityInfo_proto_rawDesc
)

func file_EntityAuthorityInfo_proto_rawDescGZIP() []byte {
	file_EntityAuthorityInfo_proto_rawDescOnce.Do(func() {
		file_EntityAuthorityInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityAuthorityInfo_proto_rawDescData)
	})
	return file_EntityAuthorityInfo_proto_rawDescData
}

var file_EntityAuthorityInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EntityAuthorityInfo_proto_goTypes = []interface{}{
	(*EntityAuthorityInfo)(nil),            // 0: EntityAuthorityInfo
	(*AbilitySyncStateInfo)(nil),           // 1: AbilitySyncStateInfo
	(*EntityRendererChangedInfo)(nil),      // 2: EntityRendererChangedInfo
	(*SceneEntityAiInfo)(nil),              // 3: SceneEntityAiInfo
	(*Vector)(nil),                         // 4: Vector
	(*AnimatorParameterValueInfoPair)(nil), // 5: AnimatorParameterValueInfoPair
	(*Unk2700_HFMDKDHCJCM)(nil),            // 6: Unk2700_HFMDKDHCJCM
}
var file_EntityAuthorityInfo_proto_depIdxs = []int32{
	1, // 0: EntityAuthorityInfo.ability_info:type_name -> AbilitySyncStateInfo
	2, // 1: EntityAuthorityInfo.renderer_changed_info:type_name -> EntityRendererChangedInfo
	3, // 2: EntityAuthorityInfo.ai_info:type_name -> SceneEntityAiInfo
	4, // 3: EntityAuthorityInfo.born_pos:type_name -> Vector
	5, // 4: EntityAuthorityInfo.pose_para_list:type_name -> AnimatorParameterValueInfoPair
	6, // 5: EntityAuthorityInfo.Unk2700_KDGMOPELHNE:type_name -> Unk2700_HFMDKDHCJCM
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_EntityAuthorityInfo_proto_init() }
func file_EntityAuthorityInfo_proto_init() {
	if File_EntityAuthorityInfo_proto != nil {
		return
	}
	file_AbilitySyncStateInfo_proto_init()
	file_AnimatorParameterValueInfoPair_proto_init()
	file_EntityRendererChangedInfo_proto_init()
	file_SceneEntityAiInfo_proto_init()
	file_Unk2700_HFMDKDHCJCM_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EntityAuthorityInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityAuthorityInfo); i {
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
			RawDescriptor: file_EntityAuthorityInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EntityAuthorityInfo_proto_goTypes,
		DependencyIndexes: file_EntityAuthorityInfo_proto_depIdxs,
		MessageInfos:      file_EntityAuthorityInfo_proto_msgTypes,
	}.Build()
	File_EntityAuthorityInfo_proto = out.File
	file_EntityAuthorityInfo_proto_rawDesc = nil
	file_EntityAuthorityInfo_proto_goTypes = nil
	file_EntityAuthorityInfo_proto_depIdxs = nil
}
