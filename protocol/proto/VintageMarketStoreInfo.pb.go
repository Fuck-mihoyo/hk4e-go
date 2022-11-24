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
// source: VintageMarketStoreInfo.proto

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

type VintageMarketStoreInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurAttrList     []uint32 `protobuf:"varint,5,rep,packed,name=cur_attr_list,json=curAttrList,proto3" json:"cur_attr_list,omitempty"`
	NextAimAttrList []uint32 `protobuf:"varint,14,rep,packed,name=next_aim_attr_list,json=nextAimAttrList,proto3" json:"next_aim_attr_list,omitempty"`
	StrategyList    []uint32 `protobuf:"varint,2,rep,packed,name=strategy_list,json=strategyList,proto3" json:"strategy_list,omitempty"`
	SlotCount       uint32   `protobuf:"varint,3,opt,name=slot_count,json=slotCount,proto3" json:"slot_count,omitempty"`
	StoreId         uint32   `protobuf:"varint,10,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
}

func (x *VintageMarketStoreInfo) Reset() {
	*x = VintageMarketStoreInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VintageMarketStoreInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VintageMarketStoreInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VintageMarketStoreInfo) ProtoMessage() {}

func (x *VintageMarketStoreInfo) ProtoReflect() protoreflect.Message {
	mi := &file_VintageMarketStoreInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VintageMarketStoreInfo.ProtoReflect.Descriptor instead.
func (*VintageMarketStoreInfo) Descriptor() ([]byte, []int) {
	return file_VintageMarketStoreInfo_proto_rawDescGZIP(), []int{0}
}

func (x *VintageMarketStoreInfo) GetCurAttrList() []uint32 {
	if x != nil {
		return x.CurAttrList
	}
	return nil
}

func (x *VintageMarketStoreInfo) GetNextAimAttrList() []uint32 {
	if x != nil {
		return x.NextAimAttrList
	}
	return nil
}

func (x *VintageMarketStoreInfo) GetStrategyList() []uint32 {
	if x != nil {
		return x.StrategyList
	}
	return nil
}

func (x *VintageMarketStoreInfo) GetSlotCount() uint32 {
	if x != nil {
		return x.SlotCount
	}
	return 0
}

func (x *VintageMarketStoreInfo) GetStoreId() uint32 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

var File_VintageMarketStoreInfo_proto protoreflect.FileDescriptor

var file_VintageMarketStoreInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x16, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x22, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x41, 0x74, 0x74, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x12, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x69, 0x6d,
	0x5f, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x41, 0x69, 0x6d, 0x41, 0x74, 0x74, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x6c, 0x6f, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_VintageMarketStoreInfo_proto_rawDescOnce sync.Once
	file_VintageMarketStoreInfo_proto_rawDescData = file_VintageMarketStoreInfo_proto_rawDesc
)

func file_VintageMarketStoreInfo_proto_rawDescGZIP() []byte {
	file_VintageMarketStoreInfo_proto_rawDescOnce.Do(func() {
		file_VintageMarketStoreInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_VintageMarketStoreInfo_proto_rawDescData)
	})
	return file_VintageMarketStoreInfo_proto_rawDescData
}

var file_VintageMarketStoreInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_VintageMarketStoreInfo_proto_goTypes = []interface{}{
	(*VintageMarketStoreInfo)(nil), // 0: proto.VintageMarketStoreInfo
}
var file_VintageMarketStoreInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_VintageMarketStoreInfo_proto_init() }
func file_VintageMarketStoreInfo_proto_init() {
	if File_VintageMarketStoreInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_VintageMarketStoreInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VintageMarketStoreInfo); i {
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
			RawDescriptor: file_VintageMarketStoreInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VintageMarketStoreInfo_proto_goTypes,
		DependencyIndexes: file_VintageMarketStoreInfo_proto_depIdxs,
		MessageInfos:      file_VintageMarketStoreInfo_proto_msgTypes,
	}.Build()
	File_VintageMarketStoreInfo_proto = out.File
	file_VintageMarketStoreInfo_proto_rawDesc = nil
	file_VintageMarketStoreInfo_proto_goTypes = nil
	file_VintageMarketStoreInfo_proto_depIdxs = nil
}
