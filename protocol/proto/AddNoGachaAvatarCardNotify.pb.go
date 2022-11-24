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
// source: AddNoGachaAvatarCardNotify.proto

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

// CmdId: 1655
// EnetChannelId: 0
// EnetIsReliable: true
type AddNoGachaAvatarCardNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferItemList    []*AddNoGachaAvatarCardTransferItem `protobuf:"bytes,4,rep,name=transfer_item_list,json=transferItemList,proto3" json:"transfer_item_list,omitempty"`
	InitialPromoteLevel uint32                              `protobuf:"varint,2,opt,name=initial_promote_level,json=initialPromoteLevel,proto3" json:"initial_promote_level,omitempty"`
	AvatarId            uint32                              `protobuf:"varint,8,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	IsTransferToItem    bool                                `protobuf:"varint,6,opt,name=is_transfer_to_item,json=isTransferToItem,proto3" json:"is_transfer_to_item,omitempty"`
	Reason              uint32                              `protobuf:"varint,9,opt,name=reason,proto3" json:"reason,omitempty"`
	InitialLevel        uint32                              `protobuf:"varint,10,opt,name=initial_level,json=initialLevel,proto3" json:"initial_level,omitempty"`
	ItemId              uint32                              `protobuf:"varint,14,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *AddNoGachaAvatarCardNotify) Reset() {
	*x = AddNoGachaAvatarCardNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AddNoGachaAvatarCardNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNoGachaAvatarCardNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoGachaAvatarCardNotify) ProtoMessage() {}

func (x *AddNoGachaAvatarCardNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AddNoGachaAvatarCardNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNoGachaAvatarCardNotify.ProtoReflect.Descriptor instead.
func (*AddNoGachaAvatarCardNotify) Descriptor() ([]byte, []int) {
	return file_AddNoGachaAvatarCardNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AddNoGachaAvatarCardNotify) GetTransferItemList() []*AddNoGachaAvatarCardTransferItem {
	if x != nil {
		return x.TransferItemList
	}
	return nil
}

func (x *AddNoGachaAvatarCardNotify) GetInitialPromoteLevel() uint32 {
	if x != nil {
		return x.InitialPromoteLevel
	}
	return 0
}

func (x *AddNoGachaAvatarCardNotify) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *AddNoGachaAvatarCardNotify) GetIsTransferToItem() bool {
	if x != nil {
		return x.IsTransferToItem
	}
	return false
}

func (x *AddNoGachaAvatarCardNotify) GetReason() uint32 {
	if x != nil {
		return x.Reason
	}
	return 0
}

func (x *AddNoGachaAvatarCardNotify) GetInitialLevel() uint32 {
	if x != nil {
		return x.InitialLevel
	}
	return 0
}

func (x *AddNoGachaAvatarCardNotify) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

var File_AddNoGachaAvatarCardNotify_proto protoreflect.FileDescriptor

var file_AddNoGachaAvatarCardNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x41, 0x64, 0x64, 0x4e, 0x6f,
	0x47, 0x61, 0x63, 0x68, 0x61, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x61, 0x72, 0x64, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc9, 0x02, 0x0a, 0x1a, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x47, 0x61, 0x63, 0x68, 0x61,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x55, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x61, 0x72, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x54, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_AddNoGachaAvatarCardNotify_proto_rawDescOnce sync.Once
	file_AddNoGachaAvatarCardNotify_proto_rawDescData = file_AddNoGachaAvatarCardNotify_proto_rawDesc
)

func file_AddNoGachaAvatarCardNotify_proto_rawDescGZIP() []byte {
	file_AddNoGachaAvatarCardNotify_proto_rawDescOnce.Do(func() {
		file_AddNoGachaAvatarCardNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AddNoGachaAvatarCardNotify_proto_rawDescData)
	})
	return file_AddNoGachaAvatarCardNotify_proto_rawDescData
}

var file_AddNoGachaAvatarCardNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AddNoGachaAvatarCardNotify_proto_goTypes = []interface{}{
	(*AddNoGachaAvatarCardNotify)(nil),       // 0: proto.AddNoGachaAvatarCardNotify
	(*AddNoGachaAvatarCardTransferItem)(nil), // 1: proto.AddNoGachaAvatarCardTransferItem
}
var file_AddNoGachaAvatarCardNotify_proto_depIdxs = []int32{
	1, // 0: proto.AddNoGachaAvatarCardNotify.transfer_item_list:type_name -> proto.AddNoGachaAvatarCardTransferItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AddNoGachaAvatarCardNotify_proto_init() }
func file_AddNoGachaAvatarCardNotify_proto_init() {
	if File_AddNoGachaAvatarCardNotify_proto != nil {
		return
	}
	file_AddNoGachaAvatarCardTransferItem_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AddNoGachaAvatarCardNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNoGachaAvatarCardNotify); i {
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
			RawDescriptor: file_AddNoGachaAvatarCardNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AddNoGachaAvatarCardNotify_proto_goTypes,
		DependencyIndexes: file_AddNoGachaAvatarCardNotify_proto_depIdxs,
		MessageInfos:      file_AddNoGachaAvatarCardNotify_proto_msgTypes,
	}.Build()
	File_AddNoGachaAvatarCardNotify_proto = out.File
	file_AddNoGachaAvatarCardNotify_proto_rawDesc = nil
	file_AddNoGachaAvatarCardNotify_proto_goTypes = nil
	file_AddNoGachaAvatarCardNotify_proto_depIdxs = nil
}
