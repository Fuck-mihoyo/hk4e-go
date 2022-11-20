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
// source: ClientLoadingCostumeVerificationNotify.proto

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

// CmdId: 3487
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type ClientLoadingCostumeVerificationNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostumeId  uint32 `protobuf:"varint,9,opt,name=costume_id,json=costumeId,proto3" json:"costume_id,omitempty"`
	PrefabHash uint64 `protobuf:"varint,2,opt,name=prefab_hash,json=prefabHash,proto3" json:"prefab_hash,omitempty"`
	Guid       uint64 `protobuf:"varint,14,opt,name=guid,proto3" json:"guid,omitempty"`
}

func (x *ClientLoadingCostumeVerificationNotify) Reset() {
	*x = ClientLoadingCostumeVerificationNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientLoadingCostumeVerificationNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientLoadingCostumeVerificationNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientLoadingCostumeVerificationNotify) ProtoMessage() {}

func (x *ClientLoadingCostumeVerificationNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ClientLoadingCostumeVerificationNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientLoadingCostumeVerificationNotify.ProtoReflect.Descriptor instead.
func (*ClientLoadingCostumeVerificationNotify) Descriptor() ([]byte, []int) {
	return file_ClientLoadingCostumeVerificationNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ClientLoadingCostumeVerificationNotify) GetCostumeId() uint32 {
	if x != nil {
		return x.CostumeId
	}
	return 0
}

func (x *ClientLoadingCostumeVerificationNotify) GetPrefabHash() uint64 {
	if x != nil {
		return x.PrefabHash
	}
	return 0
}

func (x *ClientLoadingCostumeVerificationNotify) GetGuid() uint64 {
	if x != nil {
		return x.Guid
	}
	return 0
}

var File_ClientLoadingCostumeVerificationNotify_proto protoreflect.FileDescriptor

var file_ClientLoadingCostumeVerificationNotify_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c,
	0x0a, 0x26, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74,
	0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f,
	0x73, 0x74, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x66, 0x61,
	0x62, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72,
	0x65, 0x66, 0x61, 0x62, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClientLoadingCostumeVerificationNotify_proto_rawDescOnce sync.Once
	file_ClientLoadingCostumeVerificationNotify_proto_rawDescData = file_ClientLoadingCostumeVerificationNotify_proto_rawDesc
)

func file_ClientLoadingCostumeVerificationNotify_proto_rawDescGZIP() []byte {
	file_ClientLoadingCostumeVerificationNotify_proto_rawDescOnce.Do(func() {
		file_ClientLoadingCostumeVerificationNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClientLoadingCostumeVerificationNotify_proto_rawDescData)
	})
	return file_ClientLoadingCostumeVerificationNotify_proto_rawDescData
}

var file_ClientLoadingCostumeVerificationNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ClientLoadingCostumeVerificationNotify_proto_goTypes = []interface{}{
	(*ClientLoadingCostumeVerificationNotify)(nil), // 0: ClientLoadingCostumeVerificationNotify
}
var file_ClientLoadingCostumeVerificationNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ClientLoadingCostumeVerificationNotify_proto_init() }
func file_ClientLoadingCostumeVerificationNotify_proto_init() {
	if File_ClientLoadingCostumeVerificationNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ClientLoadingCostumeVerificationNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientLoadingCostumeVerificationNotify); i {
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
			RawDescriptor: file_ClientLoadingCostumeVerificationNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClientLoadingCostumeVerificationNotify_proto_goTypes,
		DependencyIndexes: file_ClientLoadingCostumeVerificationNotify_proto_depIdxs,
		MessageInfos:      file_ClientLoadingCostumeVerificationNotify_proto_msgTypes,
	}.Build()
	File_ClientLoadingCostumeVerificationNotify_proto = out.File
	file_ClientLoadingCostumeVerificationNotify_proto_rawDesc = nil
	file_ClientLoadingCostumeVerificationNotify_proto_goTypes = nil
	file_ClientLoadingCostumeVerificationNotify_proto_depIdxs = nil
}
