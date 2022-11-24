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
// source: AISnapshotEntityData.proto

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

type AISnapshotEntityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TickTime            float32                       `protobuf:"fixed32,5,opt,name=tick_time,json=tickTime,proto3" json:"tick_time,omitempty"`
	Tactic              uint32                        `protobuf:"varint,2,opt,name=tactic,proto3" json:"tactic,omitempty"`
	FinishedSkillCycles []*AISnapshotEntitySkillCycle `protobuf:"bytes,9,rep,name=finished_skill_cycles,json=finishedSkillCycles,proto3" json:"finished_skill_cycles,omitempty"`
	MovedDistance       float32                       `protobuf:"fixed32,4,opt,name=moved_distance,json=movedDistance,proto3" json:"moved_distance,omitempty"`
	AiTargetId          uint32                        `protobuf:"varint,13,opt,name=ai_target_id,json=aiTargetId,proto3" json:"ai_target_id,omitempty"`
	ThreatTargetId      uint32                        `protobuf:"varint,3,opt,name=threat_target_id,json=threatTargetId,proto3" json:"threat_target_id,omitempty"`
	ThreatListSize      uint32                        `protobuf:"varint,1,opt,name=threat_list_size,json=threatListSize,proto3" json:"threat_list_size,omitempty"`
	EntityId            uint32                        `protobuf:"varint,15,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	HittingAvatars      map[uint32]uint32             `protobuf:"bytes,7,rep,name=hitting_avatars,json=hittingAvatars,proto3" json:"hitting_avatars,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	DistanceToPlayer    float32                       `protobuf:"fixed32,11,opt,name=distance_to_player,json=distanceToPlayer,proto3" json:"distance_to_player,omitempty"`
	AttackTargetId      uint32                        `protobuf:"varint,10,opt,name=attack_target_id,json=attackTargetId,proto3" json:"attack_target_id,omitempty"`
	RealTime            float32                       `protobuf:"fixed32,14,opt,name=real_time,json=realTime,proto3" json:"real_time,omitempty"`
}

func (x *AISnapshotEntityData) Reset() {
	*x = AISnapshotEntityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AISnapshotEntityData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AISnapshotEntityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AISnapshotEntityData) ProtoMessage() {}

func (x *AISnapshotEntityData) ProtoReflect() protoreflect.Message {
	mi := &file_AISnapshotEntityData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AISnapshotEntityData.ProtoReflect.Descriptor instead.
func (*AISnapshotEntityData) Descriptor() ([]byte, []int) {
	return file_AISnapshotEntityData_proto_rawDescGZIP(), []int{0}
}

func (x *AISnapshotEntityData) GetTickTime() float32 {
	if x != nil {
		return x.TickTime
	}
	return 0
}

func (x *AISnapshotEntityData) GetTactic() uint32 {
	if x != nil {
		return x.Tactic
	}
	return 0
}

func (x *AISnapshotEntityData) GetFinishedSkillCycles() []*AISnapshotEntitySkillCycle {
	if x != nil {
		return x.FinishedSkillCycles
	}
	return nil
}

func (x *AISnapshotEntityData) GetMovedDistance() float32 {
	if x != nil {
		return x.MovedDistance
	}
	return 0
}

func (x *AISnapshotEntityData) GetAiTargetId() uint32 {
	if x != nil {
		return x.AiTargetId
	}
	return 0
}

func (x *AISnapshotEntityData) GetThreatTargetId() uint32 {
	if x != nil {
		return x.ThreatTargetId
	}
	return 0
}

func (x *AISnapshotEntityData) GetThreatListSize() uint32 {
	if x != nil {
		return x.ThreatListSize
	}
	return 0
}

func (x *AISnapshotEntityData) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *AISnapshotEntityData) GetHittingAvatars() map[uint32]uint32 {
	if x != nil {
		return x.HittingAvatars
	}
	return nil
}

func (x *AISnapshotEntityData) GetDistanceToPlayer() float32 {
	if x != nil {
		return x.DistanceToPlayer
	}
	return 0
}

func (x *AISnapshotEntityData) GetAttackTargetId() uint32 {
	if x != nil {
		return x.AttackTargetId
	}
	return 0
}

func (x *AISnapshotEntityData) GetRealTime() float32 {
	if x != nil {
		return x.RealTime
	}
	return 0
}

var File_AISnapshotEntityData_proto protoreflect.FileDescriptor

var file_AISnapshotEntityData_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x41, 0x49, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x41, 0x49, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x04, 0x0a, 0x14, 0x41, 0x49, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x61, 0x63, 0x74, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x63,
	0x74, 0x69, 0x63, 0x12, 0x55, 0x0a, 0x15, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x49, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x43, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x13, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0d, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x69, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x69, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74,
	0x68, 0x72, 0x65, 0x61, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x10, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x58, 0x0a, 0x0f, 0x68, 0x69, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x49, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x48, 0x69, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e,
	0x68, 0x69, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73, 0x12, 0x2c,
	0x0a, 0x12, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x64, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x54, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x1a, 0x41, 0x0a, 0x13, 0x48, 0x69, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AISnapshotEntityData_proto_rawDescOnce sync.Once
	file_AISnapshotEntityData_proto_rawDescData = file_AISnapshotEntityData_proto_rawDesc
)

func file_AISnapshotEntityData_proto_rawDescGZIP() []byte {
	file_AISnapshotEntityData_proto_rawDescOnce.Do(func() {
		file_AISnapshotEntityData_proto_rawDescData = protoimpl.X.CompressGZIP(file_AISnapshotEntityData_proto_rawDescData)
	})
	return file_AISnapshotEntityData_proto_rawDescData
}

var file_AISnapshotEntityData_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AISnapshotEntityData_proto_goTypes = []interface{}{
	(*AISnapshotEntityData)(nil),       // 0: proto.AISnapshotEntityData
	nil,                                // 1: proto.AISnapshotEntityData.HittingAvatarsEntry
	(*AISnapshotEntitySkillCycle)(nil), // 2: proto.AISnapshotEntitySkillCycle
}
var file_AISnapshotEntityData_proto_depIdxs = []int32{
	2, // 0: proto.AISnapshotEntityData.finished_skill_cycles:type_name -> proto.AISnapshotEntitySkillCycle
	1, // 1: proto.AISnapshotEntityData.hitting_avatars:type_name -> proto.AISnapshotEntityData.HittingAvatarsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AISnapshotEntityData_proto_init() }
func file_AISnapshotEntityData_proto_init() {
	if File_AISnapshotEntityData_proto != nil {
		return
	}
	file_AISnapshotEntitySkillCycle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AISnapshotEntityData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AISnapshotEntityData); i {
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
			RawDescriptor: file_AISnapshotEntityData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AISnapshotEntityData_proto_goTypes,
		DependencyIndexes: file_AISnapshotEntityData_proto_depIdxs,
		MessageInfos:      file_AISnapshotEntityData_proto_msgTypes,
	}.Build()
	File_AISnapshotEntityData_proto = out.File
	file_AISnapshotEntityData_proto_rawDesc = nil
	file_AISnapshotEntityData_proto_goTypes = nil
	file_AISnapshotEntityData_proto_depIdxs = nil
}
