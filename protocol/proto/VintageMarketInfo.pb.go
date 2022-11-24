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
// source: VintageMarketInfo.proto

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

type VintageMarketInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsHelpModuleOpen        bool                      `protobuf:"varint,1485,opt,name=is_help_module_open,json=isHelpModuleOpen,proto3" json:"is_help_module_open,omitempty"`
	IsStoreContentInterrupt bool                      `protobuf:"varint,15,opt,name=is_store_content_interrupt,json=isStoreContentInterrupt,proto3" json:"is_store_content_interrupt,omitempty"`
	DealInfo                *VintageMarketDealInfo    `protobuf:"bytes,11,opt,name=deal_info,json=dealInfo,proto3" json:"deal_info,omitempty"`
	StoreRound              uint32                    `protobuf:"varint,7,opt,name=store_round,json=storeRound,proto3" json:"store_round,omitempty"`
	StoreRoundIncomeList    []uint32                  `protobuf:"varint,207,rep,packed,name=store_round_income_list,json=storeRoundIncomeList,proto3" json:"store_round_income_list,omitempty"`
	IsStoreContentFinish    bool                      `protobuf:"varint,5,opt,name=is_store_content_finish,json=isStoreContentFinish,proto3" json:"is_store_content_finish,omitempty"`
	CurEnvEventList         []uint32                  `protobuf:"varint,8,rep,packed,name=cur_env_event_list,json=curEnvEventList,proto3" json:"cur_env_event_list,omitempty"`
	IsMarketContentOpen     bool                      `protobuf:"varint,10,opt,name=is_market_content_open,json=isMarketContentOpen,proto3" json:"is_market_content_open,omitempty"`
	NextCanUseHelpRound     uint32                    `protobuf:"varint,1800,opt,name=next_can_use_help_round,json=nextCanUseHelpRound,proto3" json:"next_can_use_help_round,omitempty"`
	IsMarketContentFinish   bool                      `protobuf:"varint,2,opt,name=is_market_content_finish,json=isMarketContentFinish,proto3" json:"is_market_content_finish,omitempty"`
	ViewedStrategyList      []uint32                  `protobuf:"varint,14,rep,packed,name=viewed_strategy_list,json=viewedStrategyList,proto3" json:"viewed_strategy_list,omitempty"`
	PrevCoinCNum            uint32                    `protobuf:"varint,3,opt,name=prev_coin_c_num,json=prevCoinCNum,proto3" json:"prev_coin_c_num,omitempty"`
	BargainInfoMap          map[uint32]bool           `protobuf:"bytes,6,rep,name=bargain_info_map,json=bargainInfoMap,proto3" json:"bargain_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	DividendRewardCount     uint32                    `protobuf:"varint,1798,opt,name=dividend_reward_count,json=dividendRewardCount,proto3" json:"dividend_reward_count,omitempty"`
	CurNpcEventList         []uint32                  `protobuf:"varint,4,rep,packed,name=cur_npc_event_list,json=curNpcEventList,proto3" json:"cur_npc_event_list,omitempty"`
	IsHelpInCd              bool                      `protobuf:"varint,366,opt,name=is_help_in_cd,json=isHelpInCd,proto3" json:"is_help_in_cd,omitempty"`
	PrevCoinBNum            uint32                    `protobuf:"varint,1,opt,name=prev_coin_b_num,json=prevCoinBNum,proto3" json:"prev_coin_b_num,omitempty"`
	OpenStoreList           []*VintageMarketStoreInfo `protobuf:"bytes,9,rep,name=open_store_list,json=openStoreList,proto3" json:"open_store_list,omitempty"`
	HelpSkillId             uint32                    `protobuf:"varint,760,opt,name=help_skill_id,json=helpSkillId,proto3" json:"help_skill_id,omitempty"`
	IsRoundTipsView         bool                      `protobuf:"varint,12,opt,name=is_round_tips_view,json=isRoundTipsView,proto3" json:"is_round_tips_view,omitempty"`
	IsStrategyModuleOpen    bool                      `protobuf:"varint,876,opt,name=is_strategy_module_open,json=isStrategyModuleOpen,proto3" json:"is_strategy_module_open,omitempty"`
	UnlockStrategyList      []uint32                  `protobuf:"varint,13,rep,packed,name=unlock_strategy_list,json=unlockStrategyList,proto3" json:"unlock_strategy_list,omitempty"`
}

func (x *VintageMarketInfo) Reset() {
	*x = VintageMarketInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VintageMarketInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VintageMarketInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VintageMarketInfo) ProtoMessage() {}

func (x *VintageMarketInfo) ProtoReflect() protoreflect.Message {
	mi := &file_VintageMarketInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VintageMarketInfo.ProtoReflect.Descriptor instead.
func (*VintageMarketInfo) Descriptor() ([]byte, []int) {
	return file_VintageMarketInfo_proto_rawDescGZIP(), []int{0}
}

func (x *VintageMarketInfo) GetIsHelpModuleOpen() bool {
	if x != nil {
		return x.IsHelpModuleOpen
	}
	return false
}

func (x *VintageMarketInfo) GetIsStoreContentInterrupt() bool {
	if x != nil {
		return x.IsStoreContentInterrupt
	}
	return false
}

func (x *VintageMarketInfo) GetDealInfo() *VintageMarketDealInfo {
	if x != nil {
		return x.DealInfo
	}
	return nil
}

func (x *VintageMarketInfo) GetStoreRound() uint32 {
	if x != nil {
		return x.StoreRound
	}
	return 0
}

func (x *VintageMarketInfo) GetStoreRoundIncomeList() []uint32 {
	if x != nil {
		return x.StoreRoundIncomeList
	}
	return nil
}

func (x *VintageMarketInfo) GetIsStoreContentFinish() bool {
	if x != nil {
		return x.IsStoreContentFinish
	}
	return false
}

func (x *VintageMarketInfo) GetCurEnvEventList() []uint32 {
	if x != nil {
		return x.CurEnvEventList
	}
	return nil
}

func (x *VintageMarketInfo) GetIsMarketContentOpen() bool {
	if x != nil {
		return x.IsMarketContentOpen
	}
	return false
}

func (x *VintageMarketInfo) GetNextCanUseHelpRound() uint32 {
	if x != nil {
		return x.NextCanUseHelpRound
	}
	return 0
}

func (x *VintageMarketInfo) GetIsMarketContentFinish() bool {
	if x != nil {
		return x.IsMarketContentFinish
	}
	return false
}

func (x *VintageMarketInfo) GetViewedStrategyList() []uint32 {
	if x != nil {
		return x.ViewedStrategyList
	}
	return nil
}

func (x *VintageMarketInfo) GetPrevCoinCNum() uint32 {
	if x != nil {
		return x.PrevCoinCNum
	}
	return 0
}

func (x *VintageMarketInfo) GetBargainInfoMap() map[uint32]bool {
	if x != nil {
		return x.BargainInfoMap
	}
	return nil
}

func (x *VintageMarketInfo) GetDividendRewardCount() uint32 {
	if x != nil {
		return x.DividendRewardCount
	}
	return 0
}

func (x *VintageMarketInfo) GetCurNpcEventList() []uint32 {
	if x != nil {
		return x.CurNpcEventList
	}
	return nil
}

func (x *VintageMarketInfo) GetIsHelpInCd() bool {
	if x != nil {
		return x.IsHelpInCd
	}
	return false
}

func (x *VintageMarketInfo) GetPrevCoinBNum() uint32 {
	if x != nil {
		return x.PrevCoinBNum
	}
	return 0
}

func (x *VintageMarketInfo) GetOpenStoreList() []*VintageMarketStoreInfo {
	if x != nil {
		return x.OpenStoreList
	}
	return nil
}

func (x *VintageMarketInfo) GetHelpSkillId() uint32 {
	if x != nil {
		return x.HelpSkillId
	}
	return 0
}

func (x *VintageMarketInfo) GetIsRoundTipsView() bool {
	if x != nil {
		return x.IsRoundTipsView
	}
	return false
}

func (x *VintageMarketInfo) GetIsStrategyModuleOpen() bool {
	if x != nil {
		return x.IsStrategyModuleOpen
	}
	return false
}

func (x *VintageMarketInfo) GetUnlockStrategyList() []uint32 {
	if x != nil {
		return x.UnlockStrategyList
	}
	return nil
}

var File_VintageMarketInfo_proto protoreflect.FileDescriptor

var file_VintageMarketInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44,
	0x65, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x56,
	0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x09, 0x0a, 0x11,
	0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x5f, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0xcd, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x10, 0x69, 0x73, 0x48, 0x65, 0x6c, 0x70, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4f, 0x70, 0x65,
	0x6e, 0x12, 0x3b, 0x0a, 0x1a, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x69, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x12, 0x39,
	0x0a, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x64, 0x65, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x36, 0x0a, 0x17, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xcf, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x14, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x35, 0x0a, 0x17, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x2b, 0x0a, 0x12, 0x63, 0x75, 0x72,
	0x5f, 0x65, 0x6e, 0x76, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x45, 0x6e, 0x76, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x16, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x69, 0x73, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x35, 0x0a, 0x17, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x63, 0x61, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x6c, 0x70,
	0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x88, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6e,
	0x65, 0x78, 0x74, 0x43, 0x61, 0x6e, 0x55, 0x73, 0x65, 0x48, 0x65, 0x6c, 0x70, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x69, 0x73, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x30, 0x0a, 0x14, 0x76,
	0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x76, 0x69, 0x65, 0x77, 0x65,
	0x64, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a,
	0x0f, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x63, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x43, 0x6f, 0x69, 0x6e,
	0x43, 0x4e, 0x75, 0x6d, 0x12, 0x56, 0x0a, 0x10, 0x62, 0x61, 0x72, 0x67, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x42, 0x61, 0x72, 0x67, 0x61, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x62, 0x61,
	0x72, 0x67, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x12, 0x33, 0x0a, 0x15,
	0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x86, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x64, 0x69,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2b, 0x0a, 0x12, 0x63, 0x75, 0x72, 0x5f, 0x6e, 0x70, 0x63, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x63,
	0x75, 0x72, 0x4e, 0x70, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x64, 0x18,
	0xee, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x48, 0x65, 0x6c, 0x70, 0x49, 0x6e,
	0x43, 0x64, 0x12, 0x25, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x5f,
	0x62, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x65,
	0x76, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x4e, 0x75, 0x6d, 0x12, 0x45, 0x0a, 0x0f, 0x6f, 0x70, 0x65,
	0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x68, 0x65, 0x6c, 0x70, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0xf8, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x68, 0x65, 0x6c, 0x70, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x69, 0x73, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x69, 0x70, 0x73, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x36, 0x0a, 0x17, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0xec, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x75, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x41, 0x0a, 0x13,
	0x42, 0x61, 0x72, 0x67, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_VintageMarketInfo_proto_rawDescOnce sync.Once
	file_VintageMarketInfo_proto_rawDescData = file_VintageMarketInfo_proto_rawDesc
)

func file_VintageMarketInfo_proto_rawDescGZIP() []byte {
	file_VintageMarketInfo_proto_rawDescOnce.Do(func() {
		file_VintageMarketInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_VintageMarketInfo_proto_rawDescData)
	})
	return file_VintageMarketInfo_proto_rawDescData
}

var file_VintageMarketInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_VintageMarketInfo_proto_goTypes = []interface{}{
	(*VintageMarketInfo)(nil),      // 0: proto.VintageMarketInfo
	nil,                            // 1: proto.VintageMarketInfo.BargainInfoMapEntry
	(*VintageMarketDealInfo)(nil),  // 2: proto.VintageMarketDealInfo
	(*VintageMarketStoreInfo)(nil), // 3: proto.VintageMarketStoreInfo
}
var file_VintageMarketInfo_proto_depIdxs = []int32{
	2, // 0: proto.VintageMarketInfo.deal_info:type_name -> proto.VintageMarketDealInfo
	1, // 1: proto.VintageMarketInfo.bargain_info_map:type_name -> proto.VintageMarketInfo.BargainInfoMapEntry
	3, // 2: proto.VintageMarketInfo.open_store_list:type_name -> proto.VintageMarketStoreInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_VintageMarketInfo_proto_init() }
func file_VintageMarketInfo_proto_init() {
	if File_VintageMarketInfo_proto != nil {
		return
	}
	file_VintageMarketDealInfo_proto_init()
	file_VintageMarketStoreInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_VintageMarketInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VintageMarketInfo); i {
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
			RawDescriptor: file_VintageMarketInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VintageMarketInfo_proto_goTypes,
		DependencyIndexes: file_VintageMarketInfo_proto_depIdxs,
		MessageInfos:      file_VintageMarketInfo_proto_msgTypes,
	}.Build()
	File_VintageMarketInfo_proto = out.File
	file_VintageMarketInfo_proto_rawDesc = nil
	file_VintageMarketInfo_proto_goTypes = nil
	file_VintageMarketInfo_proto_depIdxs = nil
}
