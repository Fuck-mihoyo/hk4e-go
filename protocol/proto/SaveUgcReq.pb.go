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
// source: SaveUgcReq.proto

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

// CmdId: 6329
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type SaveUgcReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UgcType UgcType `protobuf:"varint,11,opt,name=ugc_type,json=ugcType,proto3,enum=proto.UgcType" json:"ugc_type,omitempty"`
	// Types that are assignable to Record:
	//
	//	*SaveUgcReq_MusicRecord
	Record isSaveUgcReq_Record `protobuf_oneof:"record"`
	// Types that are assignable to Brief:
	//
	//	*SaveUgcReq_MusicBriefInfo
	Brief isSaveUgcReq_Brief `protobuf_oneof:"brief"`
}

func (x *SaveUgcReq) Reset() {
	*x = SaveUgcReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SaveUgcReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveUgcReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveUgcReq) ProtoMessage() {}

func (x *SaveUgcReq) ProtoReflect() protoreflect.Message {
	mi := &file_SaveUgcReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveUgcReq.ProtoReflect.Descriptor instead.
func (*SaveUgcReq) Descriptor() ([]byte, []int) {
	return file_SaveUgcReq_proto_rawDescGZIP(), []int{0}
}

func (x *SaveUgcReq) GetUgcType() UgcType {
	if x != nil {
		return x.UgcType
	}
	return UgcType_UGC_TYPE_NONE
}

func (m *SaveUgcReq) GetRecord() isSaveUgcReq_Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (x *SaveUgcReq) GetMusicRecord() *UgcMusicRecord {
	if x, ok := x.GetRecord().(*SaveUgcReq_MusicRecord); ok {
		return x.MusicRecord
	}
	return nil
}

func (m *SaveUgcReq) GetBrief() isSaveUgcReq_Brief {
	if m != nil {
		return m.Brief
	}
	return nil
}

func (x *SaveUgcReq) GetMusicBriefInfo() *UgcMusicBriefInfo {
	if x, ok := x.GetBrief().(*SaveUgcReq_MusicBriefInfo); ok {
		return x.MusicBriefInfo
	}
	return nil
}

type isSaveUgcReq_Record interface {
	isSaveUgcReq_Record()
}

type SaveUgcReq_MusicRecord struct {
	MusicRecord *UgcMusicRecord `protobuf:"bytes,2,opt,name=music_record,json=musicRecord,proto3,oneof"`
}

func (*SaveUgcReq_MusicRecord) isSaveUgcReq_Record() {}

type isSaveUgcReq_Brief interface {
	isSaveUgcReq_Brief()
}

type SaveUgcReq_MusicBriefInfo struct {
	MusicBriefInfo *UgcMusicBriefInfo `protobuf:"bytes,1488,opt,name=music_brief_info,json=musicBriefInfo,proto3,oneof"`
}

func (*SaveUgcReq_MusicBriefInfo) isSaveUgcReq_Brief() {}

var File_SaveUgcReq_proto protoreflect.FileDescriptor

var file_SaveUgcReq_proto_rawDesc = []byte{
	0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x55, 0x67, 0x63, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x55, 0x67, 0x63, 0x4d, 0x75,
	0x73, 0x69, 0x63, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x55, 0x67, 0x63, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x55, 0x67, 0x63, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65,
	0x55, 0x67, 0x63, 0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x08, 0x75, 0x67, 0x63, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x67, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x75, 0x67, 0x63, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x67, 0x63, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x48, 0x00,
	0x52, 0x0b, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x45, 0x0a,
	0x10, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0xd0, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x67, 0x63, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e,
	0x66, 0x6f, 0x48, 0x01, 0x52, 0x0e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x42, 0x72, 0x69, 0x65, 0x66,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x07,
	0x0a, 0x05, 0x62, 0x72, 0x69, 0x65, 0x66, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SaveUgcReq_proto_rawDescOnce sync.Once
	file_SaveUgcReq_proto_rawDescData = file_SaveUgcReq_proto_rawDesc
)

func file_SaveUgcReq_proto_rawDescGZIP() []byte {
	file_SaveUgcReq_proto_rawDescOnce.Do(func() {
		file_SaveUgcReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SaveUgcReq_proto_rawDescData)
	})
	return file_SaveUgcReq_proto_rawDescData
}

var file_SaveUgcReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SaveUgcReq_proto_goTypes = []interface{}{
	(*SaveUgcReq)(nil),        // 0: proto.SaveUgcReq
	(UgcType)(0),              // 1: proto.UgcType
	(*UgcMusicRecord)(nil),    // 2: proto.UgcMusicRecord
	(*UgcMusicBriefInfo)(nil), // 3: proto.UgcMusicBriefInfo
}
var file_SaveUgcReq_proto_depIdxs = []int32{
	1, // 0: proto.SaveUgcReq.ugc_type:type_name -> proto.UgcType
	2, // 1: proto.SaveUgcReq.music_record:type_name -> proto.UgcMusicRecord
	3, // 2: proto.SaveUgcReq.music_brief_info:type_name -> proto.UgcMusicBriefInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_SaveUgcReq_proto_init() }
func file_SaveUgcReq_proto_init() {
	if File_SaveUgcReq_proto != nil {
		return
	}
	file_UgcMusicBriefInfo_proto_init()
	file_UgcMusicRecord_proto_init()
	file_UgcType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SaveUgcReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveUgcReq); i {
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
	file_SaveUgcReq_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SaveUgcReq_MusicRecord)(nil),
		(*SaveUgcReq_MusicBriefInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SaveUgcReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SaveUgcReq_proto_goTypes,
		DependencyIndexes: file_SaveUgcReq_proto_depIdxs,
		MessageInfos:      file_SaveUgcReq_proto_msgTypes,
	}.Build()
	File_SaveUgcReq_proto = out.File
	file_SaveUgcReq_proto_rawDesc = nil
	file_SaveUgcReq_proto_goTypes = nil
	file_SaveUgcReq_proto_depIdxs = nil
}
