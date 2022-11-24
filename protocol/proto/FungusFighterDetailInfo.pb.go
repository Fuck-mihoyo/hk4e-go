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
// source: FungusFighterDetailInfo.proto

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

type FungusFighterDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlotStageDetailList               []*FungusPlotStageDetail        `protobuf:"bytes,6,rep,name=plot_stage_detail_list,json=plotStageDetailList,proto3" json:"plot_stage_detail_list,omitempty"`
	FungusDetailList                  []*FungusDetail                 `protobuf:"bytes,4,rep,name=fungus_detail_list,json=fungusDetailList,proto3" json:"fungus_detail_list,omitempty"`
	UnlockCampIdList                  []uint32                        `protobuf:"varint,12,rep,packed,name=unlock_camp_id_list,json=unlockCampIdList,proto3" json:"unlock_camp_id_list,omitempty"`
	TrainingDungeonProgressDetailList []*FungusTrainingProgressDetail `protobuf:"bytes,3,rep,name=training_dungeon_progress_detail_list,json=trainingDungeonProgressDetailList,proto3" json:"training_dungeon_progress_detail_list,omitempty"`
	TrainingDungeonDetailList         []*FungusTrainingDungeonDetail  `protobuf:"bytes,15,rep,name=training_dungeon_detail_list,json=trainingDungeonDetailList,proto3" json:"training_dungeon_detail_list,omitempty"`
	FinishCampIdList                  []uint32                        `protobuf:"varint,1,rep,packed,name=finish_camp_id_list,json=finishCampIdList,proto3" json:"finish_camp_id_list,omitempty"`
	UnlockCultivateIdList             []uint32                        `protobuf:"varint,8,rep,packed,name=unlock_cultivate_id_list,json=unlockCultivateIdList,proto3" json:"unlock_cultivate_id_list,omitempty"`
}

func (x *FungusFighterDetailInfo) Reset() {
	*x = FungusFighterDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FungusFighterDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FungusFighterDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FungusFighterDetailInfo) ProtoMessage() {}

func (x *FungusFighterDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_FungusFighterDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FungusFighterDetailInfo.ProtoReflect.Descriptor instead.
func (*FungusFighterDetailInfo) Descriptor() ([]byte, []int) {
	return file_FungusFighterDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *FungusFighterDetailInfo) GetPlotStageDetailList() []*FungusPlotStageDetail {
	if x != nil {
		return x.PlotStageDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetFungusDetailList() []*FungusDetail {
	if x != nil {
		return x.FungusDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetUnlockCampIdList() []uint32 {
	if x != nil {
		return x.UnlockCampIdList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetTrainingDungeonProgressDetailList() []*FungusTrainingProgressDetail {
	if x != nil {
		return x.TrainingDungeonProgressDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetTrainingDungeonDetailList() []*FungusTrainingDungeonDetail {
	if x != nil {
		return x.TrainingDungeonDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetFinishCampIdList() []uint32 {
	if x != nil {
		return x.FinishCampIdList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetUnlockCultivateIdList() []uint32 {
	if x != nil {
		return x.UnlockCultivateIdList
	}
	return nil
}

var File_FungusFighterDetailInfo_proto protoreflect.FileDescriptor

var file_FungusFighterDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x46, 0x69, 0x67, 0x68, 0x74, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x46, 0x75, 0x6e, 0x67,
	0x75, 0x73, 0x50, 0x6c, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x54,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x46, 0x75, 0x6e, 0x67,
	0x75, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2,
	0x04, 0x0a, 0x17, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x46, 0x69, 0x67, 0x68, 0x74, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x51, 0x0a, 0x16, 0x70, 0x6c,
	0x6f, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x50, 0x6c, 0x6f, 0x74, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x13, 0x70, 0x6c, 0x6f, 0x74, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x41, 0x0a,
	0x12, 0x66, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x10,
	0x66, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x13, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x5f,
	0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x75,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x6d, 0x70, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x75, 0x0a, 0x25, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x21, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x63, 0x0a, 0x1c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x19, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x13, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x43, 0x61, 0x6d, 0x70, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x18, 0x75, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x15, 0x75, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FungusFighterDetailInfo_proto_rawDescOnce sync.Once
	file_FungusFighterDetailInfo_proto_rawDescData = file_FungusFighterDetailInfo_proto_rawDesc
)

func file_FungusFighterDetailInfo_proto_rawDescGZIP() []byte {
	file_FungusFighterDetailInfo_proto_rawDescOnce.Do(func() {
		file_FungusFighterDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_FungusFighterDetailInfo_proto_rawDescData)
	})
	return file_FungusFighterDetailInfo_proto_rawDescData
}

var file_FungusFighterDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FungusFighterDetailInfo_proto_goTypes = []interface{}{
	(*FungusFighterDetailInfo)(nil),      // 0: proto.FungusFighterDetailInfo
	(*FungusPlotStageDetail)(nil),        // 1: proto.FungusPlotStageDetail
	(*FungusDetail)(nil),                 // 2: proto.FungusDetail
	(*FungusTrainingProgressDetail)(nil), // 3: proto.FungusTrainingProgressDetail
	(*FungusTrainingDungeonDetail)(nil),  // 4: proto.FungusTrainingDungeonDetail
}
var file_FungusFighterDetailInfo_proto_depIdxs = []int32{
	1, // 0: proto.FungusFighterDetailInfo.plot_stage_detail_list:type_name -> proto.FungusPlotStageDetail
	2, // 1: proto.FungusFighterDetailInfo.fungus_detail_list:type_name -> proto.FungusDetail
	3, // 2: proto.FungusFighterDetailInfo.training_dungeon_progress_detail_list:type_name -> proto.FungusTrainingProgressDetail
	4, // 3: proto.FungusFighterDetailInfo.training_dungeon_detail_list:type_name -> proto.FungusTrainingDungeonDetail
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_FungusFighterDetailInfo_proto_init() }
func file_FungusFighterDetailInfo_proto_init() {
	if File_FungusFighterDetailInfo_proto != nil {
		return
	}
	file_FungusDetail_proto_init()
	file_FungusPlotStageDetail_proto_init()
	file_FungusTrainingDungeonDetail_proto_init()
	file_FungusTrainingProgressDetail_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FungusFighterDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FungusFighterDetailInfo); i {
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
			RawDescriptor: file_FungusFighterDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FungusFighterDetailInfo_proto_goTypes,
		DependencyIndexes: file_FungusFighterDetailInfo_proto_depIdxs,
		MessageInfos:      file_FungusFighterDetailInfo_proto_msgTypes,
	}.Build()
	File_FungusFighterDetailInfo_proto = out.File
	file_FungusFighterDetailInfo_proto_rawDesc = nil
	file_FungusFighterDetailInfo_proto_goTypes = nil
	file_FungusFighterDetailInfo_proto_depIdxs = nil
}
