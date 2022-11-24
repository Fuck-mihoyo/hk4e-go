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
// source: BlossomChestInfo.proto

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

type BlossomChestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resin              uint32   `protobuf:"varint,1,opt,name=resin,proto3" json:"resin,omitempty"`
	QualifyUidList     []uint32 `protobuf:"varint,2,rep,packed,name=qualify_uid_list,json=qualifyUidList,proto3" json:"qualify_uid_list,omitempty"`
	RemainUidList      []uint32 `protobuf:"varint,3,rep,packed,name=remain_uid_list,json=remainUidList,proto3" json:"remain_uid_list,omitempty"`
	DeadTime           uint32   `protobuf:"varint,4,opt,name=dead_time,json=deadTime,proto3" json:"dead_time,omitempty"`
	BlossomRefreshType uint32   `protobuf:"varint,5,opt,name=blossom_refresh_type,json=blossomRefreshType,proto3" json:"blossom_refresh_type,omitempty"`
	RefreshId          uint32   `protobuf:"varint,6,opt,name=refresh_id,json=refreshId,proto3" json:"refresh_id,omitempty"`
}

func (x *BlossomChestInfo) Reset() {
	*x = BlossomChestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlossomChestInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlossomChestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlossomChestInfo) ProtoMessage() {}

func (x *BlossomChestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_BlossomChestInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlossomChestInfo.ProtoReflect.Descriptor instead.
func (*BlossomChestInfo) Descriptor() ([]byte, []int) {
	return file_BlossomChestInfo_proto_rawDescGZIP(), []int{0}
}

func (x *BlossomChestInfo) GetResin() uint32 {
	if x != nil {
		return x.Resin
	}
	return 0
}

func (x *BlossomChestInfo) GetQualifyUidList() []uint32 {
	if x != nil {
		return x.QualifyUidList
	}
	return nil
}

func (x *BlossomChestInfo) GetRemainUidList() []uint32 {
	if x != nil {
		return x.RemainUidList
	}
	return nil
}

func (x *BlossomChestInfo) GetDeadTime() uint32 {
	if x != nil {
		return x.DeadTime
	}
	return 0
}

func (x *BlossomChestInfo) GetBlossomRefreshType() uint32 {
	if x != nil {
		return x.BlossomRefreshType
	}
	return 0
}

func (x *BlossomChestInfo) GetRefreshId() uint32 {
	if x != nil {
		return x.RefreshId
	}
	return 0
}

var File_BlossomChestInfo_proto protoreflect.FileDescriptor

var file_BlossomChestInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x43, 0x68, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe8, 0x01, 0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x43, 0x68, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x73, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x65, 0x73, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x71, 0x75,
	0x61, 0x6c, 0x69, 0x66, 0x79, 0x5f, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x79, 0x55, 0x69, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x75,
	0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x55, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x64, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x62, 0x6c, 0x6f,
	0x73, 0x73, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x62, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BlossomChestInfo_proto_rawDescOnce sync.Once
	file_BlossomChestInfo_proto_rawDescData = file_BlossomChestInfo_proto_rawDesc
)

func file_BlossomChestInfo_proto_rawDescGZIP() []byte {
	file_BlossomChestInfo_proto_rawDescOnce.Do(func() {
		file_BlossomChestInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_BlossomChestInfo_proto_rawDescData)
	})
	return file_BlossomChestInfo_proto_rawDescData
}

var file_BlossomChestInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BlossomChestInfo_proto_goTypes = []interface{}{
	(*BlossomChestInfo)(nil), // 0: proto.BlossomChestInfo
}
var file_BlossomChestInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BlossomChestInfo_proto_init() }
func file_BlossomChestInfo_proto_init() {
	if File_BlossomChestInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BlossomChestInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlossomChestInfo); i {
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
			RawDescriptor: file_BlossomChestInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BlossomChestInfo_proto_goTypes,
		DependencyIndexes: file_BlossomChestInfo_proto_depIdxs,
		MessageInfos:      file_BlossomChestInfo_proto_msgTypes,
	}.Build()
	File_BlossomChestInfo_proto = out.File
	file_BlossomChestInfo_proto_rawDesc = nil
	file_BlossomChestInfo_proto_goTypes = nil
	file_BlossomChestInfo_proto_depIdxs = nil
}
