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
// source: SceneGalleryVintageHuntingSettleNotify.proto

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

// CmdId: 20324
// EnetChannelId: 0
// EnetIsReliable: true
type SceneGalleryVintageHuntingSettleNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasNewWatcher      bool   `protobuf:"varint,11,opt,name=has_new_watcher,json=hasNewWatcher,proto3" json:"has_new_watcher,omitempty"`
	StageId            uint32 `protobuf:"varint,9,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	TotalWatcherNum    uint32 `protobuf:"varint,12,opt,name=total_watcher_num,json=totalWatcherNum,proto3" json:"total_watcher_num,omitempty"`
	FinishedWatcherNum uint32 `protobuf:"varint,6,opt,name=finished_watcher_num,json=finishedWatcherNum,proto3" json:"finished_watcher_num,omitempty"`
	IsNewRecord        bool   `protobuf:"varint,1,opt,name=is_new_record,json=isNewRecord,proto3" json:"is_new_record,omitempty"`
	// Types that are assignable to Info:
	//
	//	*SceneGalleryVintageHuntingSettleNotify_FirstStageInfo
	//	*SceneGalleryVintageHuntingSettleNotify_SecondStageInfo
	//	*SceneGalleryVintageHuntingSettleNotify_ThirdStageInfo
	Info isSceneGalleryVintageHuntingSettleNotify_Info `protobuf_oneof:"info"`
}

func (x *SceneGalleryVintageHuntingSettleNotify) Reset() {
	*x = SceneGalleryVintageHuntingSettleNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneGalleryVintageHuntingSettleNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneGalleryVintageHuntingSettleNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneGalleryVintageHuntingSettleNotify) ProtoMessage() {}

func (x *SceneGalleryVintageHuntingSettleNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SceneGalleryVintageHuntingSettleNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneGalleryVintageHuntingSettleNotify.ProtoReflect.Descriptor instead.
func (*SceneGalleryVintageHuntingSettleNotify) Descriptor() ([]byte, []int) {
	return file_SceneGalleryVintageHuntingSettleNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SceneGalleryVintageHuntingSettleNotify) GetHasNewWatcher() bool {
	if x != nil {
		return x.HasNewWatcher
	}
	return false
}

func (x *SceneGalleryVintageHuntingSettleNotify) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *SceneGalleryVintageHuntingSettleNotify) GetTotalWatcherNum() uint32 {
	if x != nil {
		return x.TotalWatcherNum
	}
	return 0
}

func (x *SceneGalleryVintageHuntingSettleNotify) GetFinishedWatcherNum() uint32 {
	if x != nil {
		return x.FinishedWatcherNum
	}
	return 0
}

func (x *SceneGalleryVintageHuntingSettleNotify) GetIsNewRecord() bool {
	if x != nil {
		return x.IsNewRecord
	}
	return false
}

func (m *SceneGalleryVintageHuntingSettleNotify) GetInfo() isSceneGalleryVintageHuntingSettleNotify_Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (x *SceneGalleryVintageHuntingSettleNotify) GetFirstStageInfo() *VintageHuntingFirstStageSettleInfo {
	if x, ok := x.GetInfo().(*SceneGalleryVintageHuntingSettleNotify_FirstStageInfo); ok {
		return x.FirstStageInfo
	}
	return nil
}

func (x *SceneGalleryVintageHuntingSettleNotify) GetSecondStageInfo() *VintageHuntingSecondStageSettleInfo {
	if x, ok := x.GetInfo().(*SceneGalleryVintageHuntingSettleNotify_SecondStageInfo); ok {
		return x.SecondStageInfo
	}
	return nil
}

func (x *SceneGalleryVintageHuntingSettleNotify) GetThirdStageInfo() *VintageHuntingThirdStageSettleInfo {
	if x, ok := x.GetInfo().(*SceneGalleryVintageHuntingSettleNotify_ThirdStageInfo); ok {
		return x.ThirdStageInfo
	}
	return nil
}

type isSceneGalleryVintageHuntingSettleNotify_Info interface {
	isSceneGalleryVintageHuntingSettleNotify_Info()
}

type SceneGalleryVintageHuntingSettleNotify_FirstStageInfo struct {
	FirstStageInfo *VintageHuntingFirstStageSettleInfo `protobuf:"bytes,4,opt,name=first_stage_info,json=firstStageInfo,proto3,oneof"`
}

type SceneGalleryVintageHuntingSettleNotify_SecondStageInfo struct {
	SecondStageInfo *VintageHuntingSecondStageSettleInfo `protobuf:"bytes,10,opt,name=second_stage_info,json=secondStageInfo,proto3,oneof"`
}

type SceneGalleryVintageHuntingSettleNotify_ThirdStageInfo struct {
	ThirdStageInfo *VintageHuntingThirdStageSettleInfo `protobuf:"bytes,8,opt,name=third_stage_info,json=thirdStageInfo,proto3,oneof"`
}

func (*SceneGalleryVintageHuntingSettleNotify_FirstStageInfo) isSceneGalleryVintageHuntingSettleNotify_Info() {
}

func (*SceneGalleryVintageHuntingSettleNotify_SecondStageInfo) isSceneGalleryVintageHuntingSettleNotify_Info() {
}

func (*SceneGalleryVintageHuntingSettleNotify_ThirdStageInfo) isSceneGalleryVintageHuntingSettleNotify_Info() {
}

var File_SceneGalleryVintageHuntingSettleNotify_proto protoreflect.FileDescriptor

var file_SceneGalleryVintageHuntingSettleNotify_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x56, 0x69,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x48, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x48, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x72, 0x73, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x29, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x48, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x56, 0x69, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x48, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x68, 0x69, 0x72, 0x64, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x03, 0x0a, 0x26, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x79, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x48, 0x75, 0x6e, 0x74,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x26, 0x0a, 0x0f, 0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x77,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x30,
	0x0a, 0x14, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x4e, 0x75, 0x6d,
	0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x55, 0x0a, 0x10, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x48, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x72, 0x73, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0e, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x58, 0x0a, 0x11, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x48, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x55, 0x0a, 0x10, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x48,
	0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x68, 0x69, 0x72, 0x64, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x68,
	0x69, 0x72, 0x64, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneGalleryVintageHuntingSettleNotify_proto_rawDescOnce sync.Once
	file_SceneGalleryVintageHuntingSettleNotify_proto_rawDescData = file_SceneGalleryVintageHuntingSettleNotify_proto_rawDesc
)

func file_SceneGalleryVintageHuntingSettleNotify_proto_rawDescGZIP() []byte {
	file_SceneGalleryVintageHuntingSettleNotify_proto_rawDescOnce.Do(func() {
		file_SceneGalleryVintageHuntingSettleNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneGalleryVintageHuntingSettleNotify_proto_rawDescData)
	})
	return file_SceneGalleryVintageHuntingSettleNotify_proto_rawDescData
}

var file_SceneGalleryVintageHuntingSettleNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneGalleryVintageHuntingSettleNotify_proto_goTypes = []interface{}{
	(*SceneGalleryVintageHuntingSettleNotify)(nil), // 0: proto.SceneGalleryVintageHuntingSettleNotify
	(*VintageHuntingFirstStageSettleInfo)(nil),     // 1: proto.VintageHuntingFirstStageSettleInfo
	(*VintageHuntingSecondStageSettleInfo)(nil),    // 2: proto.VintageHuntingSecondStageSettleInfo
	(*VintageHuntingThirdStageSettleInfo)(nil),     // 3: proto.VintageHuntingThirdStageSettleInfo
}
var file_SceneGalleryVintageHuntingSettleNotify_proto_depIdxs = []int32{
	1, // 0: proto.SceneGalleryVintageHuntingSettleNotify.first_stage_info:type_name -> proto.VintageHuntingFirstStageSettleInfo
	2, // 1: proto.SceneGalleryVintageHuntingSettleNotify.second_stage_info:type_name -> proto.VintageHuntingSecondStageSettleInfo
	3, // 2: proto.SceneGalleryVintageHuntingSettleNotify.third_stage_info:type_name -> proto.VintageHuntingThirdStageSettleInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_SceneGalleryVintageHuntingSettleNotify_proto_init() }
func file_SceneGalleryVintageHuntingSettleNotify_proto_init() {
	if File_SceneGalleryVintageHuntingSettleNotify_proto != nil {
		return
	}
	file_VintageHuntingFirstStageSettleInfo_proto_init()
	file_VintageHuntingSecondStageSettleInfo_proto_init()
	file_VintageHuntingThirdStageSettleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneGalleryVintageHuntingSettleNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneGalleryVintageHuntingSettleNotify); i {
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
	file_SceneGalleryVintageHuntingSettleNotify_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SceneGalleryVintageHuntingSettleNotify_FirstStageInfo)(nil),
		(*SceneGalleryVintageHuntingSettleNotify_SecondStageInfo)(nil),
		(*SceneGalleryVintageHuntingSettleNotify_ThirdStageInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SceneGalleryVintageHuntingSettleNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneGalleryVintageHuntingSettleNotify_proto_goTypes,
		DependencyIndexes: file_SceneGalleryVintageHuntingSettleNotify_proto_depIdxs,
		MessageInfos:      file_SceneGalleryVintageHuntingSettleNotify_proto_msgTypes,
	}.Build()
	File_SceneGalleryVintageHuntingSettleNotify_proto = out.File
	file_SceneGalleryVintageHuntingSettleNotify_proto_rawDesc = nil
	file_SceneGalleryVintageHuntingSettleNotify_proto_goTypes = nil
	file_SceneGalleryVintageHuntingSettleNotify_proto_depIdxs = nil
}
