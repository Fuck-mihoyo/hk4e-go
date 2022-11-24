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
// source: SetPlayerHeadImageRsp.proto

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

// CmdId: 4047
// EnetChannelId: 0
// EnetIsReliable: true
type SetPlayerHeadImageRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfilePicture *ProfilePicture `protobuf:"bytes,6,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	AvatarId       uint32          `protobuf:"varint,5,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	Retcode        int32           `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *SetPlayerHeadImageRsp) Reset() {
	*x = SetPlayerHeadImageRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetPlayerHeadImageRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPlayerHeadImageRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPlayerHeadImageRsp) ProtoMessage() {}

func (x *SetPlayerHeadImageRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SetPlayerHeadImageRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPlayerHeadImageRsp.ProtoReflect.Descriptor instead.
func (*SetPlayerHeadImageRsp) Descriptor() ([]byte, []int) {
	return file_SetPlayerHeadImageRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SetPlayerHeadImageRsp) GetProfilePicture() *ProfilePicture {
	if x != nil {
		return x.ProfilePicture
	}
	return nil
}

func (x *SetPlayerHeadImageRsp) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *SetPlayerHeadImageRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_SetPlayerHeadImageRsp_proto protoreflect.FileDescriptor

var file_SetPlayerHeadImageRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x65, 0x61, 0x64, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x15, 0x53,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x73, 0x70, 0x12, 0x3e, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetPlayerHeadImageRsp_proto_rawDescOnce sync.Once
	file_SetPlayerHeadImageRsp_proto_rawDescData = file_SetPlayerHeadImageRsp_proto_rawDesc
)

func file_SetPlayerHeadImageRsp_proto_rawDescGZIP() []byte {
	file_SetPlayerHeadImageRsp_proto_rawDescOnce.Do(func() {
		file_SetPlayerHeadImageRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetPlayerHeadImageRsp_proto_rawDescData)
	})
	return file_SetPlayerHeadImageRsp_proto_rawDescData
}

var file_SetPlayerHeadImageRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetPlayerHeadImageRsp_proto_goTypes = []interface{}{
	(*SetPlayerHeadImageRsp)(nil), // 0: proto.SetPlayerHeadImageRsp
	(*ProfilePicture)(nil),        // 1: proto.ProfilePicture
}
var file_SetPlayerHeadImageRsp_proto_depIdxs = []int32{
	1, // 0: proto.SetPlayerHeadImageRsp.profile_picture:type_name -> proto.ProfilePicture
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SetPlayerHeadImageRsp_proto_init() }
func file_SetPlayerHeadImageRsp_proto_init() {
	if File_SetPlayerHeadImageRsp_proto != nil {
		return
	}
	file_ProfilePicture_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SetPlayerHeadImageRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPlayerHeadImageRsp); i {
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
			RawDescriptor: file_SetPlayerHeadImageRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetPlayerHeadImageRsp_proto_goTypes,
		DependencyIndexes: file_SetPlayerHeadImageRsp_proto_depIdxs,
		MessageInfos:      file_SetPlayerHeadImageRsp_proto_msgTypes,
	}.Build()
	File_SetPlayerHeadImageRsp_proto = out.File
	file_SetPlayerHeadImageRsp_proto_rawDesc = nil
	file_SetPlayerHeadImageRsp_proto_goTypes = nil
	file_SetPlayerHeadImageRsp_proto_depIdxs = nil
}
