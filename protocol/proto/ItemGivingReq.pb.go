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
// source: ItemGivingReq.proto

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

type ItemGivingReq_ItemGivingRsp int32

const (
	ItemGivingReq_ITEM_GIVING_RSP_QUEST  ItemGivingReq_ItemGivingRsp = 0
	ItemGivingReq_ITEM_GIVING_RSP_GADGET ItemGivingReq_ItemGivingRsp = 1
)

// Enum value maps for ItemGivingReq_ItemGivingRsp.
var (
	ItemGivingReq_ItemGivingRsp_name = map[int32]string{
		0: "ITEM_GIVING_RSP_QUEST",
		1: "ITEM_GIVING_RSP_GADGET",
	}
	ItemGivingReq_ItemGivingRsp_value = map[string]int32{
		"ITEM_GIVING_RSP_QUEST":  0,
		"ITEM_GIVING_RSP_GADGET": 1,
	}
)

func (x ItemGivingReq_ItemGivingRsp) Enum() *ItemGivingReq_ItemGivingRsp {
	p := new(ItemGivingReq_ItemGivingRsp)
	*p = x
	return p
}

func (x ItemGivingReq_ItemGivingRsp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ItemGivingReq_ItemGivingRsp) Descriptor() protoreflect.EnumDescriptor {
	return file_ItemGivingReq_proto_enumTypes[0].Descriptor()
}

func (ItemGivingReq_ItemGivingRsp) Type() protoreflect.EnumType {
	return &file_ItemGivingReq_proto_enumTypes[0]
}

func (x ItemGivingReq_ItemGivingRsp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ItemGivingReq_ItemGivingRsp.Descriptor instead.
func (ItemGivingReq_ItemGivingRsp) EnumDescriptor() ([]byte, []int) {
	return file_ItemGivingReq_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 140
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type ItemGivingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemGuidCountMap map[uint64]uint32           `protobuf:"bytes,15,rep,name=item_guid_count_map,json=itemGuidCountMap,proto3" json:"item_guid_count_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	GivingId         uint32                      `protobuf:"varint,13,opt,name=giving_id,json=givingId,proto3" json:"giving_id,omitempty"`
	ItemParamList    []*ItemParam                `protobuf:"bytes,4,rep,name=item_param_list,json=itemParamList,proto3" json:"item_param_list,omitempty"`
	ItemGivingType   ItemGivingReq_ItemGivingRsp `protobuf:"varint,2,opt,name=item_giving_type,json=itemGivingType,proto3,enum=proto.ItemGivingReq_ItemGivingRsp" json:"item_giving_type,omitempty"`
}

func (x *ItemGivingReq) Reset() {
	*x = ItemGivingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ItemGivingReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemGivingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemGivingReq) ProtoMessage() {}

func (x *ItemGivingReq) ProtoReflect() protoreflect.Message {
	mi := &file_ItemGivingReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemGivingReq.ProtoReflect.Descriptor instead.
func (*ItemGivingReq) Descriptor() ([]byte, []int) {
	return file_ItemGivingReq_proto_rawDescGZIP(), []int{0}
}

func (x *ItemGivingReq) GetItemGuidCountMap() map[uint64]uint32 {
	if x != nil {
		return x.ItemGuidCountMap
	}
	return nil
}

func (x *ItemGivingReq) GetGivingId() uint32 {
	if x != nil {
		return x.GivingId
	}
	return 0
}

func (x *ItemGivingReq) GetItemParamList() []*ItemParam {
	if x != nil {
		return x.ItemParamList
	}
	return nil
}

func (x *ItemGivingReq) GetItemGivingType() ItemGivingReq_ItemGivingRsp {
	if x != nil {
		return x.ItemGivingType
	}
	return ItemGivingReq_ITEM_GIVING_RSP_QUEST
}

var File_ItemGivingReq_proto protoreflect.FileDescriptor

var file_ItemGivingReq_proto_rawDesc = []byte{
	0x0a, 0x13, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x49, 0x74,
	0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x03,
	0x0a, 0x0d, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x59, 0x0a, 0x13, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x75, 0x69, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x69, 0x74, 0x65, 0x6d, 0x47, 0x75,
	0x69, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x69,
	0x76, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67,
	0x69, 0x76, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0f, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x52, 0x0d, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x4c, 0x0a, 0x10, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x67, 0x69, 0x76, 0x69, 0x6e, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x52,
	0x0e, 0x69, 0x74, 0x65, 0x6d, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x1a,
	0x43, 0x0a, 0x15, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x75, 0x69, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x46, 0x0a, 0x0d, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x52, 0x73, 0x70, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x47, 0x49,
	0x56, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x53, 0x50, 0x5f, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x00,
	0x12, 0x1a, 0x0a, 0x16, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x47, 0x49, 0x56, 0x49, 0x4e, 0x47, 0x5f,
	0x52, 0x53, 0x50, 0x5f, 0x47, 0x41, 0x44, 0x47, 0x45, 0x54, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ItemGivingReq_proto_rawDescOnce sync.Once
	file_ItemGivingReq_proto_rawDescData = file_ItemGivingReq_proto_rawDesc
)

func file_ItemGivingReq_proto_rawDescGZIP() []byte {
	file_ItemGivingReq_proto_rawDescOnce.Do(func() {
		file_ItemGivingReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ItemGivingReq_proto_rawDescData)
	})
	return file_ItemGivingReq_proto_rawDescData
}

var file_ItemGivingReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ItemGivingReq_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ItemGivingReq_proto_goTypes = []interface{}{
	(ItemGivingReq_ItemGivingRsp)(0), // 0: proto.ItemGivingReq.ItemGivingRsp
	(*ItemGivingReq)(nil),            // 1: proto.ItemGivingReq
	nil,                              // 2: proto.ItemGivingReq.ItemGuidCountMapEntry
	(*ItemParam)(nil),                // 3: proto.ItemParam
}
var file_ItemGivingReq_proto_depIdxs = []int32{
	2, // 0: proto.ItemGivingReq.item_guid_count_map:type_name -> proto.ItemGivingReq.ItemGuidCountMapEntry
	3, // 1: proto.ItemGivingReq.item_param_list:type_name -> proto.ItemParam
	0, // 2: proto.ItemGivingReq.item_giving_type:type_name -> proto.ItemGivingReq.ItemGivingRsp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ItemGivingReq_proto_init() }
func file_ItemGivingReq_proto_init() {
	if File_ItemGivingReq_proto != nil {
		return
	}
	file_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ItemGivingReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemGivingReq); i {
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
			RawDescriptor: file_ItemGivingReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ItemGivingReq_proto_goTypes,
		DependencyIndexes: file_ItemGivingReq_proto_depIdxs,
		EnumInfos:         file_ItemGivingReq_proto_enumTypes,
		MessageInfos:      file_ItemGivingReq_proto_msgTypes,
	}.Build()
	File_ItemGivingReq_proto = out.File
	file_ItemGivingReq_proto_rawDesc = nil
	file_ItemGivingReq_proto_goTypes = nil
	file_ItemGivingReq_proto_depIdxs = nil
}
