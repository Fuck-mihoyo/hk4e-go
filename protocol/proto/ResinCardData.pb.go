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
// source: ResinCardData.proto

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

type ResinCardData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemainRewardDays    uint32 `protobuf:"varint,3,opt,name=remain_reward_days,json=remainRewardDays,proto3" json:"remain_reward_days,omitempty"`
	ExpireTime          uint32 `protobuf:"varint,12,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	LastDailyRewardTime uint32 `protobuf:"varint,2,opt,name=last_daily_reward_time,json=lastDailyRewardTime,proto3" json:"last_daily_reward_time,omitempty"`
	ConfigId            uint32 `protobuf:"varint,7,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
}

func (x *ResinCardData) Reset() {
	*x = ResinCardData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ResinCardData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResinCardData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResinCardData) ProtoMessage() {}

func (x *ResinCardData) ProtoReflect() protoreflect.Message {
	mi := &file_ResinCardData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResinCardData.ProtoReflect.Descriptor instead.
func (*ResinCardData) Descriptor() ([]byte, []int) {
	return file_ResinCardData_proto_rawDescGZIP(), []int{0}
}

func (x *ResinCardData) GetRemainRewardDays() uint32 {
	if x != nil {
		return x.RemainRewardDays
	}
	return 0
}

func (x *ResinCardData) GetExpireTime() uint32 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

func (x *ResinCardData) GetLastDailyRewardTime() uint32 {
	if x != nil {
		return x.LastDailyRewardTime
	}
	return 0
}

func (x *ResinCardData) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

var File_ResinCardData_proto protoreflect.FileDescriptor

var file_ResinCardData_proto_rawDesc = []byte{
	0x0a, 0x13, 0x52, 0x65, 0x73, 0x69, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a,
	0x0d, 0x52, 0x65, 0x73, 0x69, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c,
	0x0a, 0x12, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x64, 0x61, 0x79, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61, 0x79, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a,
	0x16, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6c,
	0x61, 0x73, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ResinCardData_proto_rawDescOnce sync.Once
	file_ResinCardData_proto_rawDescData = file_ResinCardData_proto_rawDesc
)

func file_ResinCardData_proto_rawDescGZIP() []byte {
	file_ResinCardData_proto_rawDescOnce.Do(func() {
		file_ResinCardData_proto_rawDescData = protoimpl.X.CompressGZIP(file_ResinCardData_proto_rawDescData)
	})
	return file_ResinCardData_proto_rawDescData
}

var file_ResinCardData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ResinCardData_proto_goTypes = []interface{}{
	(*ResinCardData)(nil), // 0: proto.ResinCardData
}
var file_ResinCardData_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ResinCardData_proto_init() }
func file_ResinCardData_proto_init() {
	if File_ResinCardData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ResinCardData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResinCardData); i {
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
			RawDescriptor: file_ResinCardData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ResinCardData_proto_goTypes,
		DependencyIndexes: file_ResinCardData_proto_depIdxs,
		MessageInfos:      file_ResinCardData_proto_msgTypes,
	}.Build()
	File_ResinCardData_proto = out.File
	file_ResinCardData_proto_rawDesc = nil
	file_ResinCardData_proto_goTypes = nil
	file_ResinCardData_proto_depIdxs = nil
}
