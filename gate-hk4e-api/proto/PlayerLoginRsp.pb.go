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
// source: PlayerLoginRsp.proto

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

// CmdId: 135
// EnetChannelId: 0
// EnetIsReliable: true
type PlayerLoginRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbilityHashMap             map[string]int32        `protobuf:"bytes,11,rep,name=ability_hash_map,json=abilityHashMap,proto3" json:"ability_hash_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	IsAudit                    bool                    `protobuf:"varint,1685,opt,name=is_audit,json=isAudit,proto3" json:"is_audit,omitempty"`
	IsNewPlayer                bool                    `protobuf:"varint,8,opt,name=is_new_player,json=isNewPlayer,proto3" json:"is_new_player,omitempty"`
	ResVersionConfig           *ResVersionConfig       `protobuf:"bytes,1969,opt,name=res_version_config,json=resVersionConfig,proto3" json:"res_version_config,omitempty"`
	IsEnableClientHashDebug    bool                    `protobuf:"varint,932,opt,name=is_enable_client_hash_debug,json=isEnableClientHashDebug,proto3" json:"is_enable_client_hash_debug,omitempty"`
	ClientMd5                  string                  `protobuf:"bytes,1830,opt,name=client_md5,json=clientMd5,proto3" json:"client_md5,omitempty"`
	ClientDataVersion          uint32                  `protobuf:"varint,1,opt,name=client_data_version,json=clientDataVersion,proto3" json:"client_data_version,omitempty"`
	CountryCode                string                  `protobuf:"bytes,1900,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	IsRelogin                  bool                    `protobuf:"varint,10,opt,name=is_relogin,json=isRelogin,proto3" json:"is_relogin,omitempty"`
	PlayerData                 []byte                  `protobuf:"bytes,13,opt,name=player_data,json=playerData,proto3" json:"player_data,omitempty"`
	GameBiz                    string                  `protobuf:"bytes,5,opt,name=game_biz,json=gameBiz,proto3" json:"game_biz,omitempty"`
	BlockInfoMap               map[uint32]*BlockInfo   `protobuf:"bytes,571,rep,name=block_info_map,json=blockInfoMap,proto3" json:"block_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RegisterCps                string                  `protobuf:"bytes,2040,opt,name=register_cps,json=registerCps,proto3" json:"register_cps,omitempty"`
	NextResVersionConfig       *ResVersionConfig       `protobuf:"bytes,1573,opt,name=next_res_version_config,json=nextResVersionConfig,proto3" json:"next_res_version_config,omitempty"`
	IsTransfer                 bool                    `protobuf:"varint,2018,opt,name=is_transfer,json=isTransfer,proto3" json:"is_transfer,omitempty"`
	TargetHomeOwnerUid         uint32                  `protobuf:"varint,553,opt,name=target_home_owner_uid,json=targetHomeOwnerUid,proto3" json:"target_home_owner_uid,omitempty"`
	ShortAbilityHashMap        []*ShortAbilityHashPair `protobuf:"bytes,250,rep,name=short_ability_hash_map,json=shortAbilityHashMap,proto3" json:"short_ability_hash_map,omitempty"`
	AbilityHashCode            int32                   `protobuf:"varint,12,opt,name=ability_hash_code,json=abilityHashCode,proto3" json:"ability_hash_code,omitempty"`
	IsScOpen                   bool                    `protobuf:"varint,1429,opt,name=is_sc_open,json=isScOpen,proto3" json:"is_sc_open,omitempty"`
	ClientSilenceDataVersion   uint32                  `protobuf:"varint,6,opt,name=client_silence_data_version,json=clientSilenceDataVersion,proto3" json:"client_silence_data_version,omitempty"`
	Birthday                   string                  `protobuf:"bytes,624,opt,name=birthday,proto3" json:"birthday,omitempty"`
	IsUseAbilityHash           bool                    `protobuf:"varint,2,opt,name=is_use_ability_hash,json=isUseAbilityHash,proto3" json:"is_use_ability_hash,omitempty"`
	ClientSilenceVersionSuffix string                  `protobuf:"bytes,1299,opt,name=client_silence_version_suffix,json=clientSilenceVersionSuffix,proto3" json:"client_silence_version_suffix,omitempty"`
	PlayerDataVersion          uint32                  `protobuf:"varint,7,opt,name=player_data_version,json=playerDataVersion,proto3" json:"player_data_version,omitempty"`
	IsDataNeedRelogin          bool                    `protobuf:"varint,951,opt,name=is_data_need_relogin,json=isDataNeedRelogin,proto3" json:"is_data_need_relogin,omitempty"`
	FeatureBlockInfoList       []*FeatureBlockInfo     `protobuf:"bytes,1352,rep,name=feature_block_info_list,json=featureBlockInfoList,proto3" json:"feature_block_info_list,omitempty"`
	ClientSilenceMd5           string                  `protobuf:"bytes,1746,opt,name=client_silence_md5,json=clientSilenceMd5,proto3" json:"client_silence_md5,omitempty"`
	TargetUid                  uint32                  `protobuf:"varint,14,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
	TotalTickTime              float64                 `protobuf:"fixed64,125,opt,name=total_tick_time,json=totalTickTime,proto3" json:"total_tick_time,omitempty"`
	LoginRand                  uint64                  `protobuf:"varint,4,opt,name=login_rand,json=loginRand,proto3" json:"login_rand,omitempty"`
	ScInfo                     []byte                  `protobuf:"bytes,2024,opt,name=sc_info,json=scInfo,proto3" json:"sc_info,omitempty"`
	ClientVersionSuffix        string                  `protobuf:"bytes,1047,opt,name=client_version_suffix,json=clientVersionSuffix,proto3" json:"client_version_suffix,omitempty"`
	NextResourceUrl            string                  `protobuf:"bytes,621,opt,name=next_resource_url,json=nextResourceUrl,proto3" json:"next_resource_url,omitempty"`
	Retcode                    int32                   `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *PlayerLoginRsp) Reset() {
	*x = PlayerLoginRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerLoginRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLoginRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLoginRsp) ProtoMessage() {}

func (x *PlayerLoginRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerLoginRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLoginRsp.ProtoReflect.Descriptor instead.
func (*PlayerLoginRsp) Descriptor() ([]byte, []int) {
	return file_PlayerLoginRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerLoginRsp) GetAbilityHashMap() map[string]int32 {
	if x != nil {
		return x.AbilityHashMap
	}
	return nil
}

func (x *PlayerLoginRsp) GetIsAudit() bool {
	if x != nil {
		return x.IsAudit
	}
	return false
}

func (x *PlayerLoginRsp) GetIsNewPlayer() bool {
	if x != nil {
		return x.IsNewPlayer
	}
	return false
}

func (x *PlayerLoginRsp) GetResVersionConfig() *ResVersionConfig {
	if x != nil {
		return x.ResVersionConfig
	}
	return nil
}

func (x *PlayerLoginRsp) GetIsEnableClientHashDebug() bool {
	if x != nil {
		return x.IsEnableClientHashDebug
	}
	return false
}

func (x *PlayerLoginRsp) GetClientMd5() string {
	if x != nil {
		return x.ClientMd5
	}
	return ""
}

func (x *PlayerLoginRsp) GetClientDataVersion() uint32 {
	if x != nil {
		return x.ClientDataVersion
	}
	return 0
}

func (x *PlayerLoginRsp) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *PlayerLoginRsp) GetIsRelogin() bool {
	if x != nil {
		return x.IsRelogin
	}
	return false
}

func (x *PlayerLoginRsp) GetPlayerData() []byte {
	if x != nil {
		return x.PlayerData
	}
	return nil
}

func (x *PlayerLoginRsp) GetGameBiz() string {
	if x != nil {
		return x.GameBiz
	}
	return ""
}

func (x *PlayerLoginRsp) GetBlockInfoMap() map[uint32]*BlockInfo {
	if x != nil {
		return x.BlockInfoMap
	}
	return nil
}

func (x *PlayerLoginRsp) GetRegisterCps() string {
	if x != nil {
		return x.RegisterCps
	}
	return ""
}

func (x *PlayerLoginRsp) GetNextResVersionConfig() *ResVersionConfig {
	if x != nil {
		return x.NextResVersionConfig
	}
	return nil
}

func (x *PlayerLoginRsp) GetIsTransfer() bool {
	if x != nil {
		return x.IsTransfer
	}
	return false
}

func (x *PlayerLoginRsp) GetTargetHomeOwnerUid() uint32 {
	if x != nil {
		return x.TargetHomeOwnerUid
	}
	return 0
}

func (x *PlayerLoginRsp) GetShortAbilityHashMap() []*ShortAbilityHashPair {
	if x != nil {
		return x.ShortAbilityHashMap
	}
	return nil
}

func (x *PlayerLoginRsp) GetAbilityHashCode() int32 {
	if x != nil {
		return x.AbilityHashCode
	}
	return 0
}

func (x *PlayerLoginRsp) GetIsScOpen() bool {
	if x != nil {
		return x.IsScOpen
	}
	return false
}

func (x *PlayerLoginRsp) GetClientSilenceDataVersion() uint32 {
	if x != nil {
		return x.ClientSilenceDataVersion
	}
	return 0
}

func (x *PlayerLoginRsp) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *PlayerLoginRsp) GetIsUseAbilityHash() bool {
	if x != nil {
		return x.IsUseAbilityHash
	}
	return false
}

func (x *PlayerLoginRsp) GetClientSilenceVersionSuffix() string {
	if x != nil {
		return x.ClientSilenceVersionSuffix
	}
	return ""
}

func (x *PlayerLoginRsp) GetPlayerDataVersion() uint32 {
	if x != nil {
		return x.PlayerDataVersion
	}
	return 0
}

func (x *PlayerLoginRsp) GetIsDataNeedRelogin() bool {
	if x != nil {
		return x.IsDataNeedRelogin
	}
	return false
}

func (x *PlayerLoginRsp) GetFeatureBlockInfoList() []*FeatureBlockInfo {
	if x != nil {
		return x.FeatureBlockInfoList
	}
	return nil
}

func (x *PlayerLoginRsp) GetClientSilenceMd5() string {
	if x != nil {
		return x.ClientSilenceMd5
	}
	return ""
}

func (x *PlayerLoginRsp) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

func (x *PlayerLoginRsp) GetTotalTickTime() float64 {
	if x != nil {
		return x.TotalTickTime
	}
	return 0
}

func (x *PlayerLoginRsp) GetLoginRand() uint64 {
	if x != nil {
		return x.LoginRand
	}
	return 0
}

func (x *PlayerLoginRsp) GetScInfo() []byte {
	if x != nil {
		return x.ScInfo
	}
	return nil
}

func (x *PlayerLoginRsp) GetClientVersionSuffix() string {
	if x != nil {
		return x.ClientVersionSuffix
	}
	return ""
}

func (x *PlayerLoginRsp) GetNextResourceUrl() string {
	if x != nil {
		return x.NextResourceUrl
	}
	return ""
}

func (x *PlayerLoginRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_PlayerLoginRsp_proto protoreflect.FileDescriptor

var file_PlayerLoginRsp_proto_rawDesc = []byte{
	0x0a, 0x14, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x41, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x48, 0x61, 0x73, 0x68, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x0d, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x4d, 0x0a, 0x10, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73,
	0x70, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x48, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x48, 0x61,
	0x73, 0x68, 0x4d, 0x61, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x18, 0x95, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0xb1, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x72, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d, 0x0a, 0x1b, 0x69, 0x73, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x5f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0xa4, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x69,
	0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73,
	0x68, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x64, 0x35, 0x18, 0xa6, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4d, 0x64, 0x35, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0xec, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x72, 0x65, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x52, 0x65, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x62, 0x69, 0x7a, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x61,
	0x6d, 0x65, 0x42, 0x69, 0x7a, 0x12, 0x48, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0xbb, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x70, 0x73, 0x18,
	0xf8, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x43, 0x70, 0x73, 0x12, 0x49, 0x0a, 0x17, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0xa5,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x14, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20,
	0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0xe2, 0x0f,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x12, 0x32, 0x0a, 0x15, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x5f,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0xa9, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x12, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x55, 0x69, 0x64, 0x12, 0x4b, 0x0a, 0x16, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0xfa,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x48, 0x61, 0x73, 0x68, 0x50, 0x61, 0x69, 0x72, 0x52, 0x13, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x48, 0x61, 0x73, 0x68, 0x4d, 0x61,
	0x70, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x73, 0x63, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x95, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x63, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x3d, 0x0a, 0x1b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x08, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0xf0, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x2d, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x75,
	0x73, 0x65, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x55, 0x73, 0x65, 0x41, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x42, 0x0a, 0x1d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x93, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x1a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x2e, 0x0a, 0x13, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x69,
	0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x18, 0xb7, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x4e, 0x65, 0x65, 0x64, 0x52, 0x65, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x49, 0x0a,
	0x17, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xc8, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x14, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0xd2,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6c,
	0x65, 0x6e, 0x63, 0x65, 0x4d, 0x64, 0x35, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x74, 0x69, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x7d, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0xe8, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x73, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x15, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78,
	0x18, 0x97, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x2b, 0x0a, 0x11,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0xed, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x1a, 0x41, 0x0a, 0x13, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x48, 0x61,
	0x73, 0x68, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerLoginRsp_proto_rawDescOnce sync.Once
	file_PlayerLoginRsp_proto_rawDescData = file_PlayerLoginRsp_proto_rawDesc
)

func file_PlayerLoginRsp_proto_rawDescGZIP() []byte {
	file_PlayerLoginRsp_proto_rawDescOnce.Do(func() {
		file_PlayerLoginRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerLoginRsp_proto_rawDescData)
	})
	return file_PlayerLoginRsp_proto_rawDescData
}

var file_PlayerLoginRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_PlayerLoginRsp_proto_goTypes = []interface{}{
	(*PlayerLoginRsp)(nil),       // 0: PlayerLoginRsp
	nil,                          // 1: PlayerLoginRsp.AbilityHashMapEntry
	nil,                          // 2: PlayerLoginRsp.BlockInfoMapEntry
	(*ResVersionConfig)(nil),     // 3: ResVersionConfig
	(*ShortAbilityHashPair)(nil), // 4: ShortAbilityHashPair
	(*FeatureBlockInfo)(nil),     // 5: FeatureBlockInfo
	(*BlockInfo)(nil),            // 6: BlockInfo
}
var file_PlayerLoginRsp_proto_depIdxs = []int32{
	1, // 0: PlayerLoginRsp.ability_hash_map:type_name -> PlayerLoginRsp.AbilityHashMapEntry
	3, // 1: PlayerLoginRsp.res_version_config:type_name -> ResVersionConfig
	2, // 2: PlayerLoginRsp.block_info_map:type_name -> PlayerLoginRsp.BlockInfoMapEntry
	3, // 3: PlayerLoginRsp.next_res_version_config:type_name -> ResVersionConfig
	4, // 4: PlayerLoginRsp.short_ability_hash_map:type_name -> ShortAbilityHashPair
	5, // 5: PlayerLoginRsp.feature_block_info_list:type_name -> FeatureBlockInfo
	6, // 6: PlayerLoginRsp.BlockInfoMapEntry.value:type_name -> BlockInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_PlayerLoginRsp_proto_init() }
func file_PlayerLoginRsp_proto_init() {
	if File_PlayerLoginRsp_proto != nil {
		return
	}
	file_BlockInfo_proto_init()
	file_FeatureBlockInfo_proto_init()
	file_ResVersionConfig_proto_init()
	file_ShortAbilityHashPair_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerLoginRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLoginRsp); i {
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
			RawDescriptor: file_PlayerLoginRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerLoginRsp_proto_goTypes,
		DependencyIndexes: file_PlayerLoginRsp_proto_depIdxs,
		MessageInfos:      file_PlayerLoginRsp_proto_msgTypes,
	}.Build()
	File_PlayerLoginRsp_proto = out.File
	file_PlayerLoginRsp_proto_rawDesc = nil
	file_PlayerLoginRsp_proto_goTypes = nil
	file_PlayerLoginRsp_proto_depIdxs = nil
}
