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
// source: QuickUseWidgetReq.proto

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

// CmdId: 4299
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type QuickUseWidgetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Param:
	//	*QuickUseWidgetReq_LocationInfo
	//	*QuickUseWidgetReq_CameraInfo
	//	*QuickUseWidgetReq_CreatorInfo
	//	*QuickUseWidgetReq_ThunderBirdFeatherInfo
	Param isQuickUseWidgetReq_Param `protobuf_oneof:"param"`
}

func (x *QuickUseWidgetReq) Reset() {
	*x = QuickUseWidgetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QuickUseWidgetReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickUseWidgetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickUseWidgetReq) ProtoMessage() {}

func (x *QuickUseWidgetReq) ProtoReflect() protoreflect.Message {
	mi := &file_QuickUseWidgetReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickUseWidgetReq.ProtoReflect.Descriptor instead.
func (*QuickUseWidgetReq) Descriptor() ([]byte, []int) {
	return file_QuickUseWidgetReq_proto_rawDescGZIP(), []int{0}
}

func (m *QuickUseWidgetReq) GetParam() isQuickUseWidgetReq_Param {
	if m != nil {
		return m.Param
	}
	return nil
}

func (x *QuickUseWidgetReq) GetLocationInfo() *WidgetCreateLocationInfo {
	if x, ok := x.GetParam().(*QuickUseWidgetReq_LocationInfo); ok {
		return x.LocationInfo
	}
	return nil
}

func (x *QuickUseWidgetReq) GetCameraInfo() *WidgetCameraInfo {
	if x, ok := x.GetParam().(*QuickUseWidgetReq_CameraInfo); ok {
		return x.CameraInfo
	}
	return nil
}

func (x *QuickUseWidgetReq) GetCreatorInfo() *WidgetCreatorInfo {
	if x, ok := x.GetParam().(*QuickUseWidgetReq_CreatorInfo); ok {
		return x.CreatorInfo
	}
	return nil
}

func (x *QuickUseWidgetReq) GetThunderBirdFeatherInfo() *WidgetThunderBirdFeatherInfo {
	if x, ok := x.GetParam().(*QuickUseWidgetReq_ThunderBirdFeatherInfo); ok {
		return x.ThunderBirdFeatherInfo
	}
	return nil
}

type isQuickUseWidgetReq_Param interface {
	isQuickUseWidgetReq_Param()
}

type QuickUseWidgetReq_LocationInfo struct {
	LocationInfo *WidgetCreateLocationInfo `protobuf:"bytes,676,opt,name=location_info,json=locationInfo,proto3,oneof"`
}

type QuickUseWidgetReq_CameraInfo struct {
	CameraInfo *WidgetCameraInfo `protobuf:"bytes,478,opt,name=camera_info,json=cameraInfo,proto3,oneof"`
}

type QuickUseWidgetReq_CreatorInfo struct {
	CreatorInfo *WidgetCreatorInfo `protobuf:"bytes,812,opt,name=creator_info,json=creatorInfo,proto3,oneof"`
}

type QuickUseWidgetReq_ThunderBirdFeatherInfo struct {
	ThunderBirdFeatherInfo *WidgetThunderBirdFeatherInfo `protobuf:"bytes,1859,opt,name=thunder_bird_feather_info,json=thunderBirdFeatherInfo,proto3,oneof"`
}

func (*QuickUseWidgetReq_LocationInfo) isQuickUseWidgetReq_Param() {}

func (*QuickUseWidgetReq_CameraInfo) isQuickUseWidgetReq_Param() {}

func (*QuickUseWidgetReq_CreatorInfo) isQuickUseWidgetReq_Param() {}

func (*QuickUseWidgetReq_ThunderBirdFeatherInfo) isQuickUseWidgetReq_Param() {}

var File_QuickUseWidgetReq_proto protoreflect.FileDescriptor

var file_QuickUseWidgetReq_proto_rawDesc = []byte{
	0x0a, 0x17, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x54, 0x68, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x69, 0x72, 0x64, 0x46, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad,
	0x02, 0x0a, 0x11, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x41, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0xa4, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35, 0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x65, 0x72,
	0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0xde, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f,
	0x48, 0x00, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38,
	0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0xac,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5b, 0x0a, 0x19, 0x74, 0x68, 0x75, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x62, 0x69, 0x72, 0x64, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0xc3, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x69, 0x72, 0x64,
	0x46, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x16, 0x74,
	0x68, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x69, 0x72, 0x64, 0x46, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x07, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_QuickUseWidgetReq_proto_rawDescOnce sync.Once
	file_QuickUseWidgetReq_proto_rawDescData = file_QuickUseWidgetReq_proto_rawDesc
)

func file_QuickUseWidgetReq_proto_rawDescGZIP() []byte {
	file_QuickUseWidgetReq_proto_rawDescOnce.Do(func() {
		file_QuickUseWidgetReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_QuickUseWidgetReq_proto_rawDescData)
	})
	return file_QuickUseWidgetReq_proto_rawDescData
}

var file_QuickUseWidgetReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_QuickUseWidgetReq_proto_goTypes = []interface{}{
	(*QuickUseWidgetReq)(nil),            // 0: QuickUseWidgetReq
	(*WidgetCreateLocationInfo)(nil),     // 1: WidgetCreateLocationInfo
	(*WidgetCameraInfo)(nil),             // 2: WidgetCameraInfo
	(*WidgetCreatorInfo)(nil),            // 3: WidgetCreatorInfo
	(*WidgetThunderBirdFeatherInfo)(nil), // 4: WidgetThunderBirdFeatherInfo
}
var file_QuickUseWidgetReq_proto_depIdxs = []int32{
	1, // 0: QuickUseWidgetReq.location_info:type_name -> WidgetCreateLocationInfo
	2, // 1: QuickUseWidgetReq.camera_info:type_name -> WidgetCameraInfo
	3, // 2: QuickUseWidgetReq.creator_info:type_name -> WidgetCreatorInfo
	4, // 3: QuickUseWidgetReq.thunder_bird_feather_info:type_name -> WidgetThunderBirdFeatherInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_QuickUseWidgetReq_proto_init() }
func file_QuickUseWidgetReq_proto_init() {
	if File_QuickUseWidgetReq_proto != nil {
		return
	}
	file_WidgetCameraInfo_proto_init()
	file_WidgetCreateLocationInfo_proto_init()
	file_WidgetCreatorInfo_proto_init()
	file_WidgetThunderBirdFeatherInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_QuickUseWidgetReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuickUseWidgetReq); i {
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
	file_QuickUseWidgetReq_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*QuickUseWidgetReq_LocationInfo)(nil),
		(*QuickUseWidgetReq_CameraInfo)(nil),
		(*QuickUseWidgetReq_CreatorInfo)(nil),
		(*QuickUseWidgetReq_ThunderBirdFeatherInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_QuickUseWidgetReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QuickUseWidgetReq_proto_goTypes,
		DependencyIndexes: file_QuickUseWidgetReq_proto_depIdxs,
		MessageInfos:      file_QuickUseWidgetReq_proto_msgTypes,
	}.Build()
	File_QuickUseWidgetReq_proto = out.File
	file_QuickUseWidgetReq_proto_rawDesc = nil
	file_QuickUseWidgetReq_proto_goTypes = nil
	file_QuickUseWidgetReq_proto_depIdxs = nil
}
