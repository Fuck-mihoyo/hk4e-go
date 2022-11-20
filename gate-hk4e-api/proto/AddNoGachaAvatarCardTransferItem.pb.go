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
// source: AddNoGachaAvatarCardTransferItem.proto

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

type AddNoGachaAvatarCardTransferItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  uint32 `protobuf:"varint,9,opt,name=count,proto3" json:"count,omitempty"`
	ItemId uint32 `protobuf:"varint,6,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	IsNew  bool   `protobuf:"varint,15,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
}

func (x *AddNoGachaAvatarCardTransferItem) Reset() {
	*x = AddNoGachaAvatarCardTransferItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AddNoGachaAvatarCardTransferItem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNoGachaAvatarCardTransferItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoGachaAvatarCardTransferItem) ProtoMessage() {}

func (x *AddNoGachaAvatarCardTransferItem) ProtoReflect() protoreflect.Message {
	mi := &file_AddNoGachaAvatarCardTransferItem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNoGachaAvatarCardTransferItem.ProtoReflect.Descriptor instead.
func (*AddNoGachaAvatarCardTransferItem) Descriptor() ([]byte, []int) {
	return file_AddNoGachaAvatarCardTransferItem_proto_rawDescGZIP(), []int{0}
}

func (x *AddNoGachaAvatarCardTransferItem) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *AddNoGachaAvatarCardTransferItem) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *AddNoGachaAvatarCardTransferItem) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

var File_AddNoGachaAvatarCardTransferItem_proto protoreflect.FileDescriptor

var file_AddNoGachaAvatarCardTransferItem_proto_rawDesc = []byte{
	0x0a, 0x26, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x43, 0x61, 0x72, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x20, 0x41, 0x64, 0x64, 0x4e,
	0x6f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x61, 0x72, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x69,
	0x73, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x4e,
	0x65, 0x77, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AddNoGachaAvatarCardTransferItem_proto_rawDescOnce sync.Once
	file_AddNoGachaAvatarCardTransferItem_proto_rawDescData = file_AddNoGachaAvatarCardTransferItem_proto_rawDesc
)

func file_AddNoGachaAvatarCardTransferItem_proto_rawDescGZIP() []byte {
	file_AddNoGachaAvatarCardTransferItem_proto_rawDescOnce.Do(func() {
		file_AddNoGachaAvatarCardTransferItem_proto_rawDescData = protoimpl.X.CompressGZIP(file_AddNoGachaAvatarCardTransferItem_proto_rawDescData)
	})
	return file_AddNoGachaAvatarCardTransferItem_proto_rawDescData
}

var file_AddNoGachaAvatarCardTransferItem_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AddNoGachaAvatarCardTransferItem_proto_goTypes = []interface{}{
	(*AddNoGachaAvatarCardTransferItem)(nil), // 0: AddNoGachaAvatarCardTransferItem
}
var file_AddNoGachaAvatarCardTransferItem_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AddNoGachaAvatarCardTransferItem_proto_init() }
func file_AddNoGachaAvatarCardTransferItem_proto_init() {
	if File_AddNoGachaAvatarCardTransferItem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AddNoGachaAvatarCardTransferItem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNoGachaAvatarCardTransferItem); i {
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
			RawDescriptor: file_AddNoGachaAvatarCardTransferItem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AddNoGachaAvatarCardTransferItem_proto_goTypes,
		DependencyIndexes: file_AddNoGachaAvatarCardTransferItem_proto_depIdxs,
		MessageInfos:      file_AddNoGachaAvatarCardTransferItem_proto_msgTypes,
	}.Build()
	File_AddNoGachaAvatarCardTransferItem_proto = out.File
	file_AddNoGachaAvatarCardTransferItem_proto_rawDesc = nil
	file_AddNoGachaAvatarCardTransferItem_proto_goTypes = nil
	file_AddNoGachaAvatarCardTransferItem_proto_depIdxs = nil
}
