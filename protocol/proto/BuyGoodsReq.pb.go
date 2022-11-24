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
// source: BuyGoodsReq.proto

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

// CmdId: 712
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type BuyGoodsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuyCount uint32     `protobuf:"varint,14,opt,name=buy_count,json=buyCount,proto3" json:"buy_count,omitempty"`
	Goods    *ShopGoods `protobuf:"bytes,15,opt,name=goods,proto3" json:"goods,omitempty"`
	ShopType uint32     `protobuf:"varint,7,opt,name=shop_type,json=shopType,proto3" json:"shop_type,omitempty"`
}

func (x *BuyGoodsReq) Reset() {
	*x = BuyGoodsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BuyGoodsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyGoodsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyGoodsReq) ProtoMessage() {}

func (x *BuyGoodsReq) ProtoReflect() protoreflect.Message {
	mi := &file_BuyGoodsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyGoodsReq.ProtoReflect.Descriptor instead.
func (*BuyGoodsReq) Descriptor() ([]byte, []int) {
	return file_BuyGoodsReq_proto_rawDescGZIP(), []int{0}
}

func (x *BuyGoodsReq) GetBuyCount() uint32 {
	if x != nil {
		return x.BuyCount
	}
	return 0
}

func (x *BuyGoodsReq) GetGoods() *ShopGoods {
	if x != nil {
		return x.Goods
	}
	return nil
}

func (x *BuyGoodsReq) GetShopType() uint32 {
	if x != nil {
		return x.ShopType
	}
	return 0
}

var File_BuyGoodsReq_proto protoreflect.FileDescriptor

var file_BuyGoodsReq_proto_rawDesc = []byte{
	0x0a, 0x11, 0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x68, 0x6f, 0x70,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x0b, 0x42,
	0x75, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75,
	0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62,
	0x75, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x68, 0x6f, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BuyGoodsReq_proto_rawDescOnce sync.Once
	file_BuyGoodsReq_proto_rawDescData = file_BuyGoodsReq_proto_rawDesc
)

func file_BuyGoodsReq_proto_rawDescGZIP() []byte {
	file_BuyGoodsReq_proto_rawDescOnce.Do(func() {
		file_BuyGoodsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_BuyGoodsReq_proto_rawDescData)
	})
	return file_BuyGoodsReq_proto_rawDescData
}

var file_BuyGoodsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BuyGoodsReq_proto_goTypes = []interface{}{
	(*BuyGoodsReq)(nil), // 0: proto.BuyGoodsReq
	(*ShopGoods)(nil),   // 1: proto.ShopGoods
}
var file_BuyGoodsReq_proto_depIdxs = []int32{
	1, // 0: proto.BuyGoodsReq.goods:type_name -> proto.ShopGoods
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BuyGoodsReq_proto_init() }
func file_BuyGoodsReq_proto_init() {
	if File_BuyGoodsReq_proto != nil {
		return
	}
	file_ShopGoods_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BuyGoodsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyGoodsReq); i {
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
			RawDescriptor: file_BuyGoodsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BuyGoodsReq_proto_goTypes,
		DependencyIndexes: file_BuyGoodsReq_proto_depIdxs,
		MessageInfos:      file_BuyGoodsReq_proto_msgTypes,
	}.Build()
	File_BuyGoodsReq_proto = out.File
	file_BuyGoodsReq_proto_rawDesc = nil
	file_BuyGoodsReq_proto_goTypes = nil
	file_BuyGoodsReq_proto_depIdxs = nil
}
