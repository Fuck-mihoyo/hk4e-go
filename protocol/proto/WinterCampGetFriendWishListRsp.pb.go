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
// source: WinterCampGetFriendWishListRsp.proto

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

// CmdId: 8937
// EnetChannelId: 0
// EnetIsReliable: true
type WinterCampGetFriendWishListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode      int32                       `protobuf:"varint,12,opt,name=retcode,proto3" json:"retcode,omitempty"`
	WishDataList []*WinterCampFriendWishData `protobuf:"bytes,5,rep,name=wish_data_list,json=wishDataList,proto3" json:"wish_data_list,omitempty"`
}

func (x *WinterCampGetFriendWishListRsp) Reset() {
	*x = WinterCampGetFriendWishListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WinterCampGetFriendWishListRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinterCampGetFriendWishListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinterCampGetFriendWishListRsp) ProtoMessage() {}

func (x *WinterCampGetFriendWishListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_WinterCampGetFriendWishListRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinterCampGetFriendWishListRsp.ProtoReflect.Descriptor instead.
func (*WinterCampGetFriendWishListRsp) Descriptor() ([]byte, []int) {
	return file_WinterCampGetFriendWishListRsp_proto_rawDescGZIP(), []int{0}
}

func (x *WinterCampGetFriendWishListRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *WinterCampGetFriendWishListRsp) GetWishDataList() []*WinterCampFriendWishData {
	if x != nil {
		return x.WishDataList
	}
	return nil
}

var File_WinterCampGetFriendWishListRsp_proto protoreflect.FileDescriptor

var file_WinterCampGetFriendWishListRsp_proto_rawDesc = []byte{
	0x0a, 0x24, 0x57, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x70, 0x47, 0x65, 0x74, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x57,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x70, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x57,
	0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01,
	0x0a, 0x1e, 0x57, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x70, 0x47, 0x65, 0x74, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x77, 0x69,
	0x73, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x43, 0x61, 0x6d, 0x70, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x57, 0x69, 0x73, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x0c, 0x77, 0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WinterCampGetFriendWishListRsp_proto_rawDescOnce sync.Once
	file_WinterCampGetFriendWishListRsp_proto_rawDescData = file_WinterCampGetFriendWishListRsp_proto_rawDesc
)

func file_WinterCampGetFriendWishListRsp_proto_rawDescGZIP() []byte {
	file_WinterCampGetFriendWishListRsp_proto_rawDescOnce.Do(func() {
		file_WinterCampGetFriendWishListRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_WinterCampGetFriendWishListRsp_proto_rawDescData)
	})
	return file_WinterCampGetFriendWishListRsp_proto_rawDescData
}

var file_WinterCampGetFriendWishListRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WinterCampGetFriendWishListRsp_proto_goTypes = []interface{}{
	(*WinterCampGetFriendWishListRsp)(nil), // 0: proto.WinterCampGetFriendWishListRsp
	(*WinterCampFriendWishData)(nil),       // 1: proto.WinterCampFriendWishData
}
var file_WinterCampGetFriendWishListRsp_proto_depIdxs = []int32{
	1, // 0: proto.WinterCampGetFriendWishListRsp.wish_data_list:type_name -> proto.WinterCampFriendWishData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_WinterCampGetFriendWishListRsp_proto_init() }
func file_WinterCampGetFriendWishListRsp_proto_init() {
	if File_WinterCampGetFriendWishListRsp_proto != nil {
		return
	}
	file_WinterCampFriendWishData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WinterCampGetFriendWishListRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinterCampGetFriendWishListRsp); i {
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
			RawDescriptor: file_WinterCampGetFriendWishListRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WinterCampGetFriendWishListRsp_proto_goTypes,
		DependencyIndexes: file_WinterCampGetFriendWishListRsp_proto_depIdxs,
		MessageInfos:      file_WinterCampGetFriendWishListRsp_proto_msgTypes,
	}.Build()
	File_WinterCampGetFriendWishListRsp_proto = out.File
	file_WinterCampGetFriendWishListRsp_proto_rawDesc = nil
	file_WinterCampGetFriendWishListRsp_proto_goTypes = nil
	file_WinterCampGetFriendWishListRsp_proto_depIdxs = nil
}
