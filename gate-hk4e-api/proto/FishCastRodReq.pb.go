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
// source: FishCastRodReq.proto

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

// CmdId: 5802
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type FishCastRodReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaitId      uint32  `protobuf:"varint,14,opt,name=bait_id,json=baitId,proto3" json:"bait_id,omitempty"`
	RodId       uint32  `protobuf:"varint,4,opt,name=rod_id,json=rodId,proto3" json:"rod_id,omitempty"`
	RodEntityId uint32  `protobuf:"varint,7,opt,name=rod_entity_id,json=rodEntityId,proto3" json:"rod_entity_id,omitempty"`
	Pos         *Vector `protobuf:"bytes,12,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *FishCastRodReq) Reset() {
	*x = FishCastRodReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FishCastRodReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FishCastRodReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FishCastRodReq) ProtoMessage() {}

func (x *FishCastRodReq) ProtoReflect() protoreflect.Message {
	mi := &file_FishCastRodReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FishCastRodReq.ProtoReflect.Descriptor instead.
func (*FishCastRodReq) Descriptor() ([]byte, []int) {
	return file_FishCastRodReq_proto_rawDescGZIP(), []int{0}
}

func (x *FishCastRodReq) GetBaitId() uint32 {
	if x != nil {
		return x.BaitId
	}
	return 0
}

func (x *FishCastRodReq) GetRodId() uint32 {
	if x != nil {
		return x.RodId
	}
	return 0
}

func (x *FishCastRodReq) GetRodEntityId() uint32 {
	if x != nil {
		return x.RodEntityId
	}
	return 0
}

func (x *FishCastRodReq) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

var File_FishCastRodReq_proto protoreflect.FileDescriptor

var file_FishCastRodReq_proto_rawDesc = []byte{
	0x0a, 0x14, 0x46, 0x69, 0x73, 0x68, 0x43, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x64, 0x52, 0x65, 0x71,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x0e, 0x46, 0x69, 0x73, 0x68, 0x43, 0x61, 0x73, 0x74,
	0x52, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x61, 0x69, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x62, 0x61, 0x69, 0x74, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x72, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x6f, 0x64, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72,
	0x6f, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x70, 0x6f,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x03, 0x70, 0x6f, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FishCastRodReq_proto_rawDescOnce sync.Once
	file_FishCastRodReq_proto_rawDescData = file_FishCastRodReq_proto_rawDesc
)

func file_FishCastRodReq_proto_rawDescGZIP() []byte {
	file_FishCastRodReq_proto_rawDescOnce.Do(func() {
		file_FishCastRodReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_FishCastRodReq_proto_rawDescData)
	})
	return file_FishCastRodReq_proto_rawDescData
}

var file_FishCastRodReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FishCastRodReq_proto_goTypes = []interface{}{
	(*FishCastRodReq)(nil), // 0: FishCastRodReq
	(*Vector)(nil),         // 1: Vector
}
var file_FishCastRodReq_proto_depIdxs = []int32{
	1, // 0: FishCastRodReq.pos:type_name -> Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FishCastRodReq_proto_init() }
func file_FishCastRodReq_proto_init() {
	if File_FishCastRodReq_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FishCastRodReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FishCastRodReq); i {
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
			RawDescriptor: file_FishCastRodReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FishCastRodReq_proto_goTypes,
		DependencyIndexes: file_FishCastRodReq_proto_depIdxs,
		MessageInfos:      file_FishCastRodReq_proto_msgTypes,
	}.Build()
	File_FishCastRodReq_proto = out.File
	file_FishCastRodReq_proto_rawDesc = nil
	file_FishCastRodReq_proto_goTypes = nil
	file_FishCastRodReq_proto_depIdxs = nil
}
