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
// source: FungusFighterTrainingGallerySettleNotify.proto

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

// CmdId: 23931
// EnetChannelId: 0
// EnetIsReliable: true
type FungusFighterTrainingGallerySettleNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsNewRecord          bool              `protobuf:"varint,14,opt,name=is_new_record,json=isNewRecord,proto3" json:"is_new_record,omitempty"`
	TotalUsedTime        uint32            `protobuf:"varint,4,opt,name=total_used_time,json=totalUsedTime,proto3" json:"total_used_time,omitempty"`
	Reason               GalleryStopReason `protobuf:"varint,5,opt,name=reason,proto3,enum=proto.GalleryStopReason" json:"reason,omitempty"`
	DeadFungusNum        uint32            `protobuf:"varint,1,opt,name=dead_fungus_num,json=deadFungusNum,proto3" json:"dead_fungus_num,omitempty"`
	SettleRound          uint32            `protobuf:"varint,15,opt,name=settle_round,json=settleRound,proto3" json:"settle_round,omitempty"`
	IsFinalSettle        bool              `protobuf:"varint,10,opt,name=is_final_settle,json=isFinalSettle,proto3" json:"is_final_settle,omitempty"`
	GadgetLifePercentage uint32            `protobuf:"varint,11,opt,name=gadget_life_percentage,json=gadgetLifePercentage,proto3" json:"gadget_life_percentage,omitempty"`
	FinalScore           uint32            `protobuf:"varint,9,opt,name=final_score,json=finalScore,proto3" json:"final_score,omitempty"`
}

func (x *FungusFighterTrainingGallerySettleNotify) Reset() {
	*x = FungusFighterTrainingGallerySettleNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FungusFighterTrainingGallerySettleNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FungusFighterTrainingGallerySettleNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FungusFighterTrainingGallerySettleNotify) ProtoMessage() {}

func (x *FungusFighterTrainingGallerySettleNotify) ProtoReflect() protoreflect.Message {
	mi := &file_FungusFighterTrainingGallerySettleNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FungusFighterTrainingGallerySettleNotify.ProtoReflect.Descriptor instead.
func (*FungusFighterTrainingGallerySettleNotify) Descriptor() ([]byte, []int) {
	return file_FungusFighterTrainingGallerySettleNotify_proto_rawDescGZIP(), []int{0}
}

func (x *FungusFighterTrainingGallerySettleNotify) GetIsNewRecord() bool {
	if x != nil {
		return x.IsNewRecord
	}
	return false
}

func (x *FungusFighterTrainingGallerySettleNotify) GetTotalUsedTime() uint32 {
	if x != nil {
		return x.TotalUsedTime
	}
	return 0
}

func (x *FungusFighterTrainingGallerySettleNotify) GetReason() GalleryStopReason {
	if x != nil {
		return x.Reason
	}
	return GalleryStopReason_GALLERY_STOP_REASON_NONE
}

func (x *FungusFighterTrainingGallerySettleNotify) GetDeadFungusNum() uint32 {
	if x != nil {
		return x.DeadFungusNum
	}
	return 0
}

func (x *FungusFighterTrainingGallerySettleNotify) GetSettleRound() uint32 {
	if x != nil {
		return x.SettleRound
	}
	return 0
}

func (x *FungusFighterTrainingGallerySettleNotify) GetIsFinalSettle() bool {
	if x != nil {
		return x.IsFinalSettle
	}
	return false
}

func (x *FungusFighterTrainingGallerySettleNotify) GetGadgetLifePercentage() uint32 {
	if x != nil {
		return x.GadgetLifePercentage
	}
	return 0
}

func (x *FungusFighterTrainingGallerySettleNotify) GetFinalScore() uint32 {
	if x != nil {
		return x.FinalScore
	}
	return 0
}

var File_FungusFighterTrainingGallerySettleNotify_proto protoreflect.FileDescriptor

var file_FungusFighterTrainingGallerySettleNotify_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x46, 0x69, 0x67, 0x68, 0x74, 0x65, 0x72, 0x54,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf2, 0x02, 0x0a, 0x28, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x46, 0x69, 0x67, 0x68, 0x74,
	0x65, 0x72, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x22, 0x0a,
	0x0d, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x55, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x64,
	0x65, 0x61, 0x64, 0x5f, 0x66, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x64, 0x65, 0x61, 0x64, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73,
	0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x12, 0x34,
	0x0a, 0x16, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x5f, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14,
	0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FungusFighterTrainingGallerySettleNotify_proto_rawDescOnce sync.Once
	file_FungusFighterTrainingGallerySettleNotify_proto_rawDescData = file_FungusFighterTrainingGallerySettleNotify_proto_rawDesc
)

func file_FungusFighterTrainingGallerySettleNotify_proto_rawDescGZIP() []byte {
	file_FungusFighterTrainingGallerySettleNotify_proto_rawDescOnce.Do(func() {
		file_FungusFighterTrainingGallerySettleNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_FungusFighterTrainingGallerySettleNotify_proto_rawDescData)
	})
	return file_FungusFighterTrainingGallerySettleNotify_proto_rawDescData
}

var file_FungusFighterTrainingGallerySettleNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FungusFighterTrainingGallerySettleNotify_proto_goTypes = []interface{}{
	(*FungusFighterTrainingGallerySettleNotify)(nil), // 0: proto.FungusFighterTrainingGallerySettleNotify
	(GalleryStopReason)(0),                           // 1: proto.GalleryStopReason
}
var file_FungusFighterTrainingGallerySettleNotify_proto_depIdxs = []int32{
	1, // 0: proto.FungusFighterTrainingGallerySettleNotify.reason:type_name -> proto.GalleryStopReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FungusFighterTrainingGallerySettleNotify_proto_init() }
func file_FungusFighterTrainingGallerySettleNotify_proto_init() {
	if File_FungusFighterTrainingGallerySettleNotify_proto != nil {
		return
	}
	file_GalleryStopReason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FungusFighterTrainingGallerySettleNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FungusFighterTrainingGallerySettleNotify); i {
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
			RawDescriptor: file_FungusFighterTrainingGallerySettleNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FungusFighterTrainingGallerySettleNotify_proto_goTypes,
		DependencyIndexes: file_FungusFighterTrainingGallerySettleNotify_proto_depIdxs,
		MessageInfos:      file_FungusFighterTrainingGallerySettleNotify_proto_msgTypes,
	}.Build()
	File_FungusFighterTrainingGallerySettleNotify_proto = out.File
	file_FungusFighterTrainingGallerySettleNotify_proto_rawDesc = nil
	file_FungusFighterTrainingGallerySettleNotify_proto_goTypes = nil
	file_FungusFighterTrainingGallerySettleNotify_proto_depIdxs = nil
}
