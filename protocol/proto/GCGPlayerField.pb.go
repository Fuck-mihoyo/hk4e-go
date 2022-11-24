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
// source: GCGPlayerField.proto

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

type GCGPlayerField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModifyZoneMap        map[uint32]*GCGZone    `protobuf:"bytes,2,rep,name=modify_zone_map,json=modifyZoneMap,proto3" json:"modify_zone_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CurWaitingIndex      uint32                 `protobuf:"varint,383,opt,name=cur_waiting_index,json=curWaitingIndex,proto3" json:"cur_waiting_index,omitempty"`
	SummonZone           *GCGZone               `protobuf:"bytes,1,opt,name=summon_zone,json=summonZone,proto3" json:"summon_zone,omitempty"`
	FieldShowId          uint32                 `protobuf:"varint,8,opt,name=field_show_id,json=fieldShowId,proto3" json:"field_show_id,omitempty"`
	CardBackShowId       uint32                 `protobuf:"varint,12,opt,name=card_back_show_id,json=cardBackShowId,proto3" json:"card_back_show_id,omitempty"`
	DiceCount            uint32                 `protobuf:"varint,3,opt,name=dice_count,json=diceCount,proto3" json:"dice_count,omitempty"`
	ControllerId         uint32                 `protobuf:"varint,10,opt,name=controller_id,json=controllerId,proto3" json:"controller_id,omitempty"`
	OnStageZone          *GCGZone               `protobuf:"bytes,14,opt,name=on_stage_zone,json=onStageZone,proto3" json:"on_stage_zone,omitempty"`
	IsPassed             bool                   `protobuf:"varint,7,opt,name=is_passed,json=isPassed,proto3" json:"is_passed,omitempty"`
	CharacterZone        *GCGZone               `protobuf:"bytes,5,opt,name=character_zone,json=characterZone,proto3" json:"character_zone,omitempty"`
	OnStageCharacterGuid uint32                 `protobuf:"varint,6,opt,name=on_stage_character_guid,json=onStageCharacterGuid,proto3" json:"on_stage_character_guid,omitempty"`
	AssistZone           *GCGZone               `protobuf:"bytes,15,opt,name=assist_zone,json=assistZone,proto3" json:"assist_zone,omitempty"`
	DeckCardNum          uint32                 `protobuf:"varint,13,opt,name=deck_card_num,json=deckCardNum,proto3" json:"deck_card_num,omitempty"`
	DiceSideList         []GCGDiceSideType      `protobuf:"varint,11,rep,packed,name=dice_side_list,json=diceSideList,proto3,enum=proto.GCGDiceSideType" json:"dice_side_list,omitempty"`
	HandZone             *GCGZone               `protobuf:"bytes,9,opt,name=hand_zone,json=handZone,proto3" json:"hand_zone,omitempty"`
	IntentionList        []*GCGPVEIntention     `protobuf:"bytes,1192,rep,name=intention_list,json=intentionList,proto3" json:"intention_list,omitempty"`
	WaitingList          []*GCGWaitingCharacter `protobuf:"bytes,4,rep,name=waiting_list,json=waitingList,proto3" json:"waiting_list,omitempty"`
}

func (x *GCGPlayerField) Reset() {
	*x = GCGPlayerField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGPlayerField_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGPlayerField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGPlayerField) ProtoMessage() {}

func (x *GCGPlayerField) ProtoReflect() protoreflect.Message {
	mi := &file_GCGPlayerField_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGPlayerField.ProtoReflect.Descriptor instead.
func (*GCGPlayerField) Descriptor() ([]byte, []int) {
	return file_GCGPlayerField_proto_rawDescGZIP(), []int{0}
}

func (x *GCGPlayerField) GetModifyZoneMap() map[uint32]*GCGZone {
	if x != nil {
		return x.ModifyZoneMap
	}
	return nil
}

func (x *GCGPlayerField) GetCurWaitingIndex() uint32 {
	if x != nil {
		return x.CurWaitingIndex
	}
	return 0
}

func (x *GCGPlayerField) GetSummonZone() *GCGZone {
	if x != nil {
		return x.SummonZone
	}
	return nil
}

func (x *GCGPlayerField) GetFieldShowId() uint32 {
	if x != nil {
		return x.FieldShowId
	}
	return 0
}

func (x *GCGPlayerField) GetCardBackShowId() uint32 {
	if x != nil {
		return x.CardBackShowId
	}
	return 0
}

func (x *GCGPlayerField) GetDiceCount() uint32 {
	if x != nil {
		return x.DiceCount
	}
	return 0
}

func (x *GCGPlayerField) GetControllerId() uint32 {
	if x != nil {
		return x.ControllerId
	}
	return 0
}

func (x *GCGPlayerField) GetOnStageZone() *GCGZone {
	if x != nil {
		return x.OnStageZone
	}
	return nil
}

func (x *GCGPlayerField) GetIsPassed() bool {
	if x != nil {
		return x.IsPassed
	}
	return false
}

func (x *GCGPlayerField) GetCharacterZone() *GCGZone {
	if x != nil {
		return x.CharacterZone
	}
	return nil
}

func (x *GCGPlayerField) GetOnStageCharacterGuid() uint32 {
	if x != nil {
		return x.OnStageCharacterGuid
	}
	return 0
}

func (x *GCGPlayerField) GetAssistZone() *GCGZone {
	if x != nil {
		return x.AssistZone
	}
	return nil
}

func (x *GCGPlayerField) GetDeckCardNum() uint32 {
	if x != nil {
		return x.DeckCardNum
	}
	return 0
}

func (x *GCGPlayerField) GetDiceSideList() []GCGDiceSideType {
	if x != nil {
		return x.DiceSideList
	}
	return nil
}

func (x *GCGPlayerField) GetHandZone() *GCGZone {
	if x != nil {
		return x.HandZone
	}
	return nil
}

func (x *GCGPlayerField) GetIntentionList() []*GCGPVEIntention {
	if x != nil {
		return x.IntentionList
	}
	return nil
}

func (x *GCGPlayerField) GetWaitingList() []*GCGWaitingCharacter {
	if x != nil {
		return x.WaitingList
	}
	return nil
}

var File_GCGPlayerField_proto protoreflect.FileDescriptor

var file_GCGPlayerField_proto_rawDesc = []byte{
	0x0a, 0x14, 0x47, 0x43, 0x47, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x47,
	0x43, 0x47, 0x44, 0x69, 0x63, 0x65, 0x53, 0x69, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x47, 0x43, 0x47, 0x50, 0x56, 0x45, 0x49, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x47, 0x43, 0x47,
	0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x47, 0x43, 0x47, 0x5a, 0x6f, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x07, 0x0a, 0x0e, 0x47, 0x43, 0x47, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x50, 0x0a, 0x0f, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5a,
	0x6f, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x75,
	0x72, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0xff, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x57, 0x61, 0x69, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2f, 0x0a, 0x0b, 0x73, 0x75, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x0a, 0x73, 0x75,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x68, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x11,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x69,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x61, 0x72, 0x64, 0x42, 0x61, 0x63,
	0x6b, 0x53, 0x68, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0d, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47, 0x5a, 0x6f,
	0x6e, 0x65, 0x52, 0x0b, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x61, 0x73, 0x73, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x0e,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47,
	0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5a,
	0x6f, 0x6e, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x47, 0x75, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x0b, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47, 0x5a, 0x6f, 0x6e, 0x65, 0x52,
	0x0a, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64,
	0x65, 0x63, 0x6b, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x64, 0x65, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x12,
	0x3c, 0x0a, 0x0e, 0x64, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x43, 0x47, 0x44, 0x69, 0x63, 0x65, 0x53, 0x69, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0c, 0x64, 0x69, 0x63, 0x65, 0x53, 0x69, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x09, 0x68, 0x61, 0x6e, 0x64, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47, 0x5a, 0x6f, 0x6e, 0x65,
	0x52, 0x08, 0x68, 0x61, 0x6e, 0x64, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x3e, 0x0a, 0x0e, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xa8, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47, 0x50,
	0x56, 0x45, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x69, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x77, 0x61,
	0x69, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47, 0x57, 0x61, 0x69, 0x74,
	0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x77, 0x61,
	0x69, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x50, 0x0a, 0x12, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47, 0x5a, 0x6f, 0x6e, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGPlayerField_proto_rawDescOnce sync.Once
	file_GCGPlayerField_proto_rawDescData = file_GCGPlayerField_proto_rawDesc
)

func file_GCGPlayerField_proto_rawDescGZIP() []byte {
	file_GCGPlayerField_proto_rawDescOnce.Do(func() {
		file_GCGPlayerField_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGPlayerField_proto_rawDescData)
	})
	return file_GCGPlayerField_proto_rawDescData
}

var file_GCGPlayerField_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GCGPlayerField_proto_goTypes = []interface{}{
	(*GCGPlayerField)(nil),      // 0: proto.GCGPlayerField
	nil,                         // 1: proto.GCGPlayerField.ModifyZoneMapEntry
	(*GCGZone)(nil),             // 2: proto.GCGZone
	(GCGDiceSideType)(0),        // 3: proto.GCGDiceSideType
	(*GCGPVEIntention)(nil),     // 4: proto.GCGPVEIntention
	(*GCGWaitingCharacter)(nil), // 5: proto.GCGWaitingCharacter
}
var file_GCGPlayerField_proto_depIdxs = []int32{
	1,  // 0: proto.GCGPlayerField.modify_zone_map:type_name -> proto.GCGPlayerField.ModifyZoneMapEntry
	2,  // 1: proto.GCGPlayerField.summon_zone:type_name -> proto.GCGZone
	2,  // 2: proto.GCGPlayerField.on_stage_zone:type_name -> proto.GCGZone
	2,  // 3: proto.GCGPlayerField.character_zone:type_name -> proto.GCGZone
	2,  // 4: proto.GCGPlayerField.assist_zone:type_name -> proto.GCGZone
	3,  // 5: proto.GCGPlayerField.dice_side_list:type_name -> proto.GCGDiceSideType
	2,  // 6: proto.GCGPlayerField.hand_zone:type_name -> proto.GCGZone
	4,  // 7: proto.GCGPlayerField.intention_list:type_name -> proto.GCGPVEIntention
	5,  // 8: proto.GCGPlayerField.waiting_list:type_name -> proto.GCGWaitingCharacter
	2,  // 9: proto.GCGPlayerField.ModifyZoneMapEntry.value:type_name -> proto.GCGZone
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_GCGPlayerField_proto_init() }
func file_GCGPlayerField_proto_init() {
	if File_GCGPlayerField_proto != nil {
		return
	}
	file_GCGDiceSideType_proto_init()
	file_GCGPVEIntention_proto_init()
	file_GCGWaitingCharacter_proto_init()
	file_GCGZone_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGPlayerField_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGPlayerField); i {
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
			RawDescriptor: file_GCGPlayerField_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGPlayerField_proto_goTypes,
		DependencyIndexes: file_GCGPlayerField_proto_depIdxs,
		MessageInfos:      file_GCGPlayerField_proto_msgTypes,
	}.Build()
	File_GCGPlayerField_proto = out.File
	file_GCGPlayerField_proto_rawDesc = nil
	file_GCGPlayerField_proto_goTypes = nil
	file_GCGPlayerField_proto_depIdxs = nil
}
