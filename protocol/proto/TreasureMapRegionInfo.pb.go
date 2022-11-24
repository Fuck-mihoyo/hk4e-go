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
// source: TreasureMapRegionInfo.proto

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

type TreasureMapRegionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime       uint32  `protobuf:"varint,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	CurrentProgress uint32  `protobuf:"varint,11,opt,name=current_progress,json=currentProgress,proto3" json:"current_progress,omitempty"`
	IsDoneMpSpot    bool    `protobuf:"varint,3,opt,name=is_done_mp_spot,json=isDoneMpSpot,proto3" json:"is_done_mp_spot,omitempty"`
	SceneId         uint32  `protobuf:"varint,2,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	GoalPoints      uint32  `protobuf:"varint,12,opt,name=goal_points,json=goalPoints,proto3" json:"goal_points,omitempty"`
	IsFindMpSpot    bool    `protobuf:"varint,4,opt,name=is_find_mp_spot,json=isFindMpSpot,proto3" json:"is_find_mp_spot,omitempty"`
	RegionRadius    uint32  `protobuf:"varint,1,opt,name=region_radius,json=regionRadius,proto3" json:"region_radius,omitempty"`
	RegionCenterPos *Vector `protobuf:"bytes,9,opt,name=region_center_pos,json=regionCenterPos,proto3" json:"region_center_pos,omitempty"`
	RegionId        uint32  `protobuf:"varint,5,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
}

func (x *TreasureMapRegionInfo) Reset() {
	*x = TreasureMapRegionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TreasureMapRegionInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasureMapRegionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasureMapRegionInfo) ProtoMessage() {}

func (x *TreasureMapRegionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_TreasureMapRegionInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasureMapRegionInfo.ProtoReflect.Descriptor instead.
func (*TreasureMapRegionInfo) Descriptor() ([]byte, []int) {
	return file_TreasureMapRegionInfo_proto_rawDescGZIP(), []int{0}
}

func (x *TreasureMapRegionInfo) GetStartTime() uint32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *TreasureMapRegionInfo) GetCurrentProgress() uint32 {
	if x != nil {
		return x.CurrentProgress
	}
	return 0
}

func (x *TreasureMapRegionInfo) GetIsDoneMpSpot() bool {
	if x != nil {
		return x.IsDoneMpSpot
	}
	return false
}

func (x *TreasureMapRegionInfo) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *TreasureMapRegionInfo) GetGoalPoints() uint32 {
	if x != nil {
		return x.GoalPoints
	}
	return 0
}

func (x *TreasureMapRegionInfo) GetIsFindMpSpot() bool {
	if x != nil {
		return x.IsFindMpSpot
	}
	return false
}

func (x *TreasureMapRegionInfo) GetRegionRadius() uint32 {
	if x != nil {
		return x.RegionRadius
	}
	return 0
}

func (x *TreasureMapRegionInfo) GetRegionCenterPos() *Vector {
	if x != nil {
		return x.RegionCenterPos
	}
	return nil
}

func (x *TreasureMapRegionInfo) GetRegionId() uint32 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

var File_TreasureMapRegionInfo_proto protoreflect.FileDescriptor

var file_TreasureMapRegionInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe8, 0x02, 0x0a, 0x15, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4d,
	0x61, 0x70, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x64, 0x6f, 0x6e,
	0x65, 0x5f, 0x6d, 0x70, 0x5f, 0x73, 0x70, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x69, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x4d, 0x70, 0x53, 0x70, 0x6f, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x6f, 0x61, 0x6c,
	0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x67,
	0x6f, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0f, 0x69, 0x73, 0x5f,
	0x66, 0x69, 0x6e, 0x64, 0x5f, 0x6d, 0x70, 0x5f, 0x73, 0x70, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x70, 0x53, 0x70, 0x6f, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52,
	0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_TreasureMapRegionInfo_proto_rawDescOnce sync.Once
	file_TreasureMapRegionInfo_proto_rawDescData = file_TreasureMapRegionInfo_proto_rawDesc
)

func file_TreasureMapRegionInfo_proto_rawDescGZIP() []byte {
	file_TreasureMapRegionInfo_proto_rawDescOnce.Do(func() {
		file_TreasureMapRegionInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_TreasureMapRegionInfo_proto_rawDescData)
	})
	return file_TreasureMapRegionInfo_proto_rawDescData
}

var file_TreasureMapRegionInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TreasureMapRegionInfo_proto_goTypes = []interface{}{
	(*TreasureMapRegionInfo)(nil), // 0: proto.TreasureMapRegionInfo
	(*Vector)(nil),                // 1: proto.Vector
}
var file_TreasureMapRegionInfo_proto_depIdxs = []int32{
	1, // 0: proto.TreasureMapRegionInfo.region_center_pos:type_name -> proto.Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TreasureMapRegionInfo_proto_init() }
func file_TreasureMapRegionInfo_proto_init() {
	if File_TreasureMapRegionInfo_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TreasureMapRegionInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasureMapRegionInfo); i {
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
			RawDescriptor: file_TreasureMapRegionInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TreasureMapRegionInfo_proto_goTypes,
		DependencyIndexes: file_TreasureMapRegionInfo_proto_depIdxs,
		MessageInfos:      file_TreasureMapRegionInfo_proto_msgTypes,
	}.Build()
	File_TreasureMapRegionInfo_proto = out.File
	file_TreasureMapRegionInfo_proto_rawDesc = nil
	file_TreasureMapRegionInfo_proto_goTypes = nil
	file_TreasureMapRegionInfo_proto_depIdxs = nil
}
