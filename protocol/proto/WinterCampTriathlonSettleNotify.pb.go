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
// source: WinterCampTriathlonSettleNotify.proto

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

// CmdId: 8342
// EnetChannelId: 0
// EnetIsReliable: true
type WinterCampTriathlonSettleNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LimitedCoin uint32 `protobuf:"varint,9,opt,name=limited_coin,json=limitedCoin,proto3" json:"limited_coin,omitempty"`
	NormalCoin  uint32 `protobuf:"varint,2,opt,name=normal_coin,json=normalCoin,proto3" json:"normal_coin,omitempty"`
	IsNewRecord bool   `protobuf:"varint,7,opt,name=is_new_record,json=isNewRecord,proto3" json:"is_new_record,omitempty"`
	IsSuccess   bool   `protobuf:"varint,3,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	GalleryId   uint32 `protobuf:"varint,13,opt,name=gallery_id,json=galleryId,proto3" json:"gallery_id,omitempty"`
	RemainTime  uint32 `protobuf:"varint,4,opt,name=remain_time,json=remainTime,proto3" json:"remain_time,omitempty"`
	Score       uint32 `protobuf:"varint,11,opt,name=score,proto3" json:"score,omitempty"`
	RaceId      uint32 `protobuf:"varint,15,opt,name=race_id,json=raceId,proto3" json:"race_id,omitempty"`
}

func (x *WinterCampTriathlonSettleNotify) Reset() {
	*x = WinterCampTriathlonSettleNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WinterCampTriathlonSettleNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinterCampTriathlonSettleNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinterCampTriathlonSettleNotify) ProtoMessage() {}

func (x *WinterCampTriathlonSettleNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WinterCampTriathlonSettleNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinterCampTriathlonSettleNotify.ProtoReflect.Descriptor instead.
func (*WinterCampTriathlonSettleNotify) Descriptor() ([]byte, []int) {
	return file_WinterCampTriathlonSettleNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WinterCampTriathlonSettleNotify) GetLimitedCoin() uint32 {
	if x != nil {
		return x.LimitedCoin
	}
	return 0
}

func (x *WinterCampTriathlonSettleNotify) GetNormalCoin() uint32 {
	if x != nil {
		return x.NormalCoin
	}
	return 0
}

func (x *WinterCampTriathlonSettleNotify) GetIsNewRecord() bool {
	if x != nil {
		return x.IsNewRecord
	}
	return false
}

func (x *WinterCampTriathlonSettleNotify) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *WinterCampTriathlonSettleNotify) GetGalleryId() uint32 {
	if x != nil {
		return x.GalleryId
	}
	return 0
}

func (x *WinterCampTriathlonSettleNotify) GetRemainTime() uint32 {
	if x != nil {
		return x.RemainTime
	}
	return 0
}

func (x *WinterCampTriathlonSettleNotify) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *WinterCampTriathlonSettleNotify) GetRaceId() uint32 {
	if x != nil {
		return x.RaceId
	}
	return 0
}

var File_WinterCampTriathlonSettleNotify_proto protoreflect.FileDescriptor

var file_WinterCampTriathlonSettleNotify_proto_rawDesc = []byte{
	0x0a, 0x25, 0x57, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x70, 0x54, 0x72, 0x69, 0x61,
	0x74, 0x68, 0x6c, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97,
	0x02, 0x0a, 0x1f, 0x57, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x70, 0x54, 0x72, 0x69,
	0x61, 0x74, 0x68, 0x6c, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77,
	0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69,
	0x73, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WinterCampTriathlonSettleNotify_proto_rawDescOnce sync.Once
	file_WinterCampTriathlonSettleNotify_proto_rawDescData = file_WinterCampTriathlonSettleNotify_proto_rawDesc
)

func file_WinterCampTriathlonSettleNotify_proto_rawDescGZIP() []byte {
	file_WinterCampTriathlonSettleNotify_proto_rawDescOnce.Do(func() {
		file_WinterCampTriathlonSettleNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WinterCampTriathlonSettleNotify_proto_rawDescData)
	})
	return file_WinterCampTriathlonSettleNotify_proto_rawDescData
}

var file_WinterCampTriathlonSettleNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WinterCampTriathlonSettleNotify_proto_goTypes = []interface{}{
	(*WinterCampTriathlonSettleNotify)(nil), // 0: proto.WinterCampTriathlonSettleNotify
}
var file_WinterCampTriathlonSettleNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_WinterCampTriathlonSettleNotify_proto_init() }
func file_WinterCampTriathlonSettleNotify_proto_init() {
	if File_WinterCampTriathlonSettleNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WinterCampTriathlonSettleNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinterCampTriathlonSettleNotify); i {
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
			RawDescriptor: file_WinterCampTriathlonSettleNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WinterCampTriathlonSettleNotify_proto_goTypes,
		DependencyIndexes: file_WinterCampTriathlonSettleNotify_proto_depIdxs,
		MessageInfos:      file_WinterCampTriathlonSettleNotify_proto_msgTypes,
	}.Build()
	File_WinterCampTriathlonSettleNotify_proto = out.File
	file_WinterCampTriathlonSettleNotify_proto_rawDesc = nil
	file_WinterCampTriathlonSettleNotify_proto_goTypes = nil
	file_WinterCampTriathlonSettleNotify_proto_depIdxs = nil
}
