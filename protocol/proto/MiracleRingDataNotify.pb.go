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
// source: MiracleRingDataNotify.proto

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

// CmdId: 5225
// EnetChannelId: 0
// EnetIsReliable: true
type MiracleRingDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsGadgetCreated     bool   `protobuf:"varint,8,opt,name=is_gadget_created,json=isGadgetCreated,proto3" json:"is_gadget_created,omitempty"`
	LastTakeRewardTime  uint32 `protobuf:"varint,14,opt,name=last_take_reward_time,json=lastTakeRewardTime,proto3" json:"last_take_reward_time,omitempty"`
	GadgetEntityId      uint32 `protobuf:"varint,12,opt,name=gadget_entity_id,json=gadgetEntityId,proto3" json:"gadget_entity_id,omitempty"`
	LastDeliverItemTime uint32 `protobuf:"varint,10,opt,name=last_deliver_item_time,json=lastDeliverItemTime,proto3" json:"last_deliver_item_time,omitempty"`
	MiracleRingCd       uint32 `protobuf:"varint,7,opt,name=miracle_ring_cd,json=miracleRingCd,proto3" json:"miracle_ring_cd,omitempty"`
}

func (x *MiracleRingDataNotify) Reset() {
	*x = MiracleRingDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiracleRingDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiracleRingDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiracleRingDataNotify) ProtoMessage() {}

func (x *MiracleRingDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_MiracleRingDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiracleRingDataNotify.ProtoReflect.Descriptor instead.
func (*MiracleRingDataNotify) Descriptor() ([]byte, []int) {
	return file_MiracleRingDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *MiracleRingDataNotify) GetIsGadgetCreated() bool {
	if x != nil {
		return x.IsGadgetCreated
	}
	return false
}

func (x *MiracleRingDataNotify) GetLastTakeRewardTime() uint32 {
	if x != nil {
		return x.LastTakeRewardTime
	}
	return 0
}

func (x *MiracleRingDataNotify) GetGadgetEntityId() uint32 {
	if x != nil {
		return x.GadgetEntityId
	}
	return 0
}

func (x *MiracleRingDataNotify) GetLastDeliverItemTime() uint32 {
	if x != nil {
		return x.LastDeliverItemTime
	}
	return 0
}

func (x *MiracleRingDataNotify) GetMiracleRingCd() uint32 {
	if x != nil {
		return x.MiracleRingCd
	}
	return 0
}

var File_MiracleRingDataNotify_proto protoreflect.FileDescriptor

var file_MiracleRingDataNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x15, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x52, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2a,
	0x0a, 0x11, 0x69, 0x73, 0x5f, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x47, 0x61, 0x64,
	0x67, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x54,
	0x61, 0x6b, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x69,
	0x6e, 0x67, 0x43, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MiracleRingDataNotify_proto_rawDescOnce sync.Once
	file_MiracleRingDataNotify_proto_rawDescData = file_MiracleRingDataNotify_proto_rawDesc
)

func file_MiracleRingDataNotify_proto_rawDescGZIP() []byte {
	file_MiracleRingDataNotify_proto_rawDescOnce.Do(func() {
		file_MiracleRingDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_MiracleRingDataNotify_proto_rawDescData)
	})
	return file_MiracleRingDataNotify_proto_rawDescData
}

var file_MiracleRingDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MiracleRingDataNotify_proto_goTypes = []interface{}{
	(*MiracleRingDataNotify)(nil), // 0: proto.MiracleRingDataNotify
}
var file_MiracleRingDataNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MiracleRingDataNotify_proto_init() }
func file_MiracleRingDataNotify_proto_init() {
	if File_MiracleRingDataNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MiracleRingDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiracleRingDataNotify); i {
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
			RawDescriptor: file_MiracleRingDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MiracleRingDataNotify_proto_goTypes,
		DependencyIndexes: file_MiracleRingDataNotify_proto_depIdxs,
		MessageInfos:      file_MiracleRingDataNotify_proto_msgTypes,
	}.Build()
	File_MiracleRingDataNotify_proto = out.File
	file_MiracleRingDataNotify_proto_rawDesc = nil
	file_MiracleRingDataNotify_proto_goTypes = nil
	file_MiracleRingDataNotify_proto_depIdxs = nil
}
