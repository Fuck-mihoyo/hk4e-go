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
// source: BattlePassSchedule.proto

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

type BattlePassSchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level                  uint32                 `protobuf:"varint,14,opt,name=level,proto3" json:"level,omitempty"`
	BeginTime              uint32                 `protobuf:"varint,2,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	EndTime                uint32                 `protobuf:"varint,15,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Point                  uint32                 `protobuf:"varint,1,opt,name=point,proto3" json:"point,omitempty"`
	CurCycle               *BattlePassCycle       `protobuf:"bytes,4,opt,name=cur_cycle,json=curCycle,proto3" json:"cur_cycle,omitempty"`
	UnlockStatus           BattlePassUnlockStatus `protobuf:"varint,7,opt,name=unlock_status,json=unlockStatus,proto3,enum=BattlePassUnlockStatus" json:"unlock_status,omitempty"`
	RewardTakenList        []*BattlePassRewardTag `protobuf:"bytes,11,rep,name=reward_taken_list,json=rewardTakenList,proto3" json:"reward_taken_list,omitempty"`
	CurCyclePoints         uint32                 `protobuf:"varint,10,opt,name=cur_cycle_points,json=curCyclePoints,proto3" json:"cur_cycle_points,omitempty"`
	Unk2700_ODHAAHEPFAG    uint32                 `protobuf:"varint,12,opt,name=Unk2700_ODHAAHEPFAG,json=Unk2700ODHAAHEPFAG,proto3" json:"Unk2700_ODHAAHEPFAG,omitempty"`
	ProductInfo            *BattlePassProduct     `protobuf:"bytes,13,opt,name=product_info,json=productInfo,proto3" json:"product_info,omitempty"`
	IsExtraPaidRewardTaken bool                   `protobuf:"varint,6,opt,name=is_extra_paid_reward_taken,json=isExtraPaidRewardTaken,proto3" json:"is_extra_paid_reward_taken,omitempty"`
	IsViewed               bool                   `protobuf:"varint,3,opt,name=is_viewed,json=isViewed,proto3" json:"is_viewed,omitempty"`
	ScheduleId             uint32                 `protobuf:"varint,9,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
}

func (x *BattlePassSchedule) Reset() {
	*x = BattlePassSchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BattlePassSchedule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattlePassSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattlePassSchedule) ProtoMessage() {}

func (x *BattlePassSchedule) ProtoReflect() protoreflect.Message {
	mi := &file_BattlePassSchedule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattlePassSchedule.ProtoReflect.Descriptor instead.
func (*BattlePassSchedule) Descriptor() ([]byte, []int) {
	return file_BattlePassSchedule_proto_rawDescGZIP(), []int{0}
}

func (x *BattlePassSchedule) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *BattlePassSchedule) GetBeginTime() uint32 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *BattlePassSchedule) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *BattlePassSchedule) GetPoint() uint32 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *BattlePassSchedule) GetCurCycle() *BattlePassCycle {
	if x != nil {
		return x.CurCycle
	}
	return nil
}

func (x *BattlePassSchedule) GetUnlockStatus() BattlePassUnlockStatus {
	if x != nil {
		return x.UnlockStatus
	}
	return BattlePassUnlockStatus_BATTLE_PASS_UNLOCK_STATUS_INVALID
}

func (x *BattlePassSchedule) GetRewardTakenList() []*BattlePassRewardTag {
	if x != nil {
		return x.RewardTakenList
	}
	return nil
}

func (x *BattlePassSchedule) GetCurCyclePoints() uint32 {
	if x != nil {
		return x.CurCyclePoints
	}
	return 0
}

func (x *BattlePassSchedule) GetUnk2700_ODHAAHEPFAG() uint32 {
	if x != nil {
		return x.Unk2700_ODHAAHEPFAG
	}
	return 0
}

func (x *BattlePassSchedule) GetProductInfo() *BattlePassProduct {
	if x != nil {
		return x.ProductInfo
	}
	return nil
}

func (x *BattlePassSchedule) GetIsExtraPaidRewardTaken() bool {
	if x != nil {
		return x.IsExtraPaidRewardTaken
	}
	return false
}

func (x *BattlePassSchedule) GetIsViewed() bool {
	if x != nil {
		return x.IsViewed
	}
	return false
}

func (x *BattlePassSchedule) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

var File_BattlePassSchedule_proto protoreflect.FileDescriptor

var file_BattlePassSchedule_proto_rawDesc = []byte{
	0x0a, 0x18, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x61, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x04, 0x0a, 0x12, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x2d, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x43, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x08, 0x63, 0x75, 0x72, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12,
	0x3c, 0x0a, 0x0d, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x0c, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a,
	0x11, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x61, 0x67, 0x52, 0x0f,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x43, 0x79,
	0x63, 0x6c, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x4f, 0x44, 0x48, 0x41, 0x41, 0x48, 0x45, 0x50, 0x46, 0x41, 0x47,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4f,
	0x44, 0x48, 0x41, 0x41, 0x48, 0x45, 0x50, 0x46, 0x41, 0x47, 0x12, 0x35, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x3a, 0x0a, 0x1a, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x70, 0x61,
	0x69, 0x64, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x69, 0x73, 0x45, 0x78, 0x74, 0x72, 0x61, 0x50, 0x61,
	0x69, 0x64, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BattlePassSchedule_proto_rawDescOnce sync.Once
	file_BattlePassSchedule_proto_rawDescData = file_BattlePassSchedule_proto_rawDesc
)

func file_BattlePassSchedule_proto_rawDescGZIP() []byte {
	file_BattlePassSchedule_proto_rawDescOnce.Do(func() {
		file_BattlePassSchedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_BattlePassSchedule_proto_rawDescData)
	})
	return file_BattlePassSchedule_proto_rawDescData
}

var file_BattlePassSchedule_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BattlePassSchedule_proto_goTypes = []interface{}{
	(*BattlePassSchedule)(nil),  // 0: BattlePassSchedule
	(*BattlePassCycle)(nil),     // 1: BattlePassCycle
	(BattlePassUnlockStatus)(0), // 2: BattlePassUnlockStatus
	(*BattlePassRewardTag)(nil), // 3: BattlePassRewardTag
	(*BattlePassProduct)(nil),   // 4: BattlePassProduct
}
var file_BattlePassSchedule_proto_depIdxs = []int32{
	1, // 0: BattlePassSchedule.cur_cycle:type_name -> BattlePassCycle
	2, // 1: BattlePassSchedule.unlock_status:type_name -> BattlePassUnlockStatus
	3, // 2: BattlePassSchedule.reward_taken_list:type_name -> BattlePassRewardTag
	4, // 3: BattlePassSchedule.product_info:type_name -> BattlePassProduct
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_BattlePassSchedule_proto_init() }
func file_BattlePassSchedule_proto_init() {
	if File_BattlePassSchedule_proto != nil {
		return
	}
	file_BattlePassCycle_proto_init()
	file_BattlePassProduct_proto_init()
	file_BattlePassRewardTag_proto_init()
	file_BattlePassUnlockStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BattlePassSchedule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattlePassSchedule); i {
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
			RawDescriptor: file_BattlePassSchedule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BattlePassSchedule_proto_goTypes,
		DependencyIndexes: file_BattlePassSchedule_proto_depIdxs,
		MessageInfos:      file_BattlePassSchedule_proto_msgTypes,
	}.Build()
	File_BattlePassSchedule_proto = out.File
	file_BattlePassSchedule_proto_rawDesc = nil
	file_BattlePassSchedule_proto_goTypes = nil
	file_BattlePassSchedule_proto_depIdxs = nil
}
