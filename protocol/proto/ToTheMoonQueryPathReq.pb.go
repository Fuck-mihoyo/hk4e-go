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
// source: ToTheMoonQueryPathReq.proto

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

type ToTheMoonQueryPathReq_OptionType int32

const (
	ToTheMoonQueryPathReq_OPTION_TYPE_NONE   ToTheMoonQueryPathReq_OptionType = 0
	ToTheMoonQueryPathReq_OPTION_TYPE_NORMAL ToTheMoonQueryPathReq_OptionType = 1
)

// Enum value maps for ToTheMoonQueryPathReq_OptionType.
var (
	ToTheMoonQueryPathReq_OptionType_name = map[int32]string{
		0: "OPTION_TYPE_NONE",
		1: "OPTION_TYPE_NORMAL",
	}
	ToTheMoonQueryPathReq_OptionType_value = map[string]int32{
		"OPTION_TYPE_NONE":   0,
		"OPTION_TYPE_NORMAL": 1,
	}
)

func (x ToTheMoonQueryPathReq_OptionType) Enum() *ToTheMoonQueryPathReq_OptionType {
	p := new(ToTheMoonQueryPathReq_OptionType)
	*p = x
	return p
}

func (x ToTheMoonQueryPathReq_OptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ToTheMoonQueryPathReq_OptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_ToTheMoonQueryPathReq_proto_enumTypes[0].Descriptor()
}

func (ToTheMoonQueryPathReq_OptionType) Type() protoreflect.EnumType {
	return &file_ToTheMoonQueryPathReq_proto_enumTypes[0]
}

func (x ToTheMoonQueryPathReq_OptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ToTheMoonQueryPathReq_OptionType.Descriptor instead.
func (ToTheMoonQueryPathReq_OptionType) EnumDescriptor() ([]byte, []int) {
	return file_ToTheMoonQueryPathReq_proto_rawDescGZIP(), []int{0, 0}
}

type ToTheMoonQueryPathReq_AStarMethod int32

const (
	ToTheMoonQueryPathReq_A_STAR_METHOD_CLASSIC    ToTheMoonQueryPathReq_AStarMethod = 0
	ToTheMoonQueryPathReq_A_STAR_METHOD_TENDENCY   ToTheMoonQueryPathReq_AStarMethod = 1
	ToTheMoonQueryPathReq_A_STAR_METHOD_ADAPTIVE   ToTheMoonQueryPathReq_AStarMethod = 2
	ToTheMoonQueryPathReq_A_STAR_METHOD_INFLECTION ToTheMoonQueryPathReq_AStarMethod = 3
)

// Enum value maps for ToTheMoonQueryPathReq_AStarMethod.
var (
	ToTheMoonQueryPathReq_AStarMethod_name = map[int32]string{
		0: "A_STAR_METHOD_CLASSIC",
		1: "A_STAR_METHOD_TENDENCY",
		2: "A_STAR_METHOD_ADAPTIVE",
		3: "A_STAR_METHOD_INFLECTION",
	}
	ToTheMoonQueryPathReq_AStarMethod_value = map[string]int32{
		"A_STAR_METHOD_CLASSIC":    0,
		"A_STAR_METHOD_TENDENCY":   1,
		"A_STAR_METHOD_ADAPTIVE":   2,
		"A_STAR_METHOD_INFLECTION": 3,
	}
)

func (x ToTheMoonQueryPathReq_AStarMethod) Enum() *ToTheMoonQueryPathReq_AStarMethod {
	p := new(ToTheMoonQueryPathReq_AStarMethod)
	*p = x
	return p
}

func (x ToTheMoonQueryPathReq_AStarMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ToTheMoonQueryPathReq_AStarMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_ToTheMoonQueryPathReq_proto_enumTypes[1].Descriptor()
}

func (ToTheMoonQueryPathReq_AStarMethod) Type() protoreflect.EnumType {
	return &file_ToTheMoonQueryPathReq_proto_enumTypes[1]
}

func (x ToTheMoonQueryPathReq_AStarMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ToTheMoonQueryPathReq_AStarMethod.Descriptor instead.
func (ToTheMoonQueryPathReq_AStarMethod) EnumDescriptor() ([]byte, []int) {
	return file_ToTheMoonQueryPathReq_proto_rawDescGZIP(), []int{0, 1}
}

type ToTheMoonQueryPathReq_FilterType int32

const (
	ToTheMoonQueryPathReq_FILTER_TYPE_ALL   ToTheMoonQueryPathReq_FilterType = 0
	ToTheMoonQueryPathReq_FILTER_TYPE_AIR   ToTheMoonQueryPathReq_FilterType = 1
	ToTheMoonQueryPathReq_FILTER_TYPE_WATER ToTheMoonQueryPathReq_FilterType = 2
)

// Enum value maps for ToTheMoonQueryPathReq_FilterType.
var (
	ToTheMoonQueryPathReq_FilterType_name = map[int32]string{
		0: "FILTER_TYPE_ALL",
		1: "FILTER_TYPE_AIR",
		2: "FILTER_TYPE_WATER",
	}
	ToTheMoonQueryPathReq_FilterType_value = map[string]int32{
		"FILTER_TYPE_ALL":   0,
		"FILTER_TYPE_AIR":   1,
		"FILTER_TYPE_WATER": 2,
	}
)

func (x ToTheMoonQueryPathReq_FilterType) Enum() *ToTheMoonQueryPathReq_FilterType {
	p := new(ToTheMoonQueryPathReq_FilterType)
	*p = x
	return p
}

func (x ToTheMoonQueryPathReq_FilterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ToTheMoonQueryPathReq_FilterType) Descriptor() protoreflect.EnumDescriptor {
	return file_ToTheMoonQueryPathReq_proto_enumTypes[2].Descriptor()
}

func (ToTheMoonQueryPathReq_FilterType) Type() protoreflect.EnumType {
	return &file_ToTheMoonQueryPathReq_proto_enumTypes[2]
}

func (x ToTheMoonQueryPathReq_FilterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ToTheMoonQueryPathReq_FilterType.Descriptor instead.
func (ToTheMoonQueryPathReq_FilterType) EnumDescriptor() ([]byte, []int) {
	return file_ToTheMoonQueryPathReq_proto_rawDescGZIP(), []int{0, 2}
}

// CmdId: 6172
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type ToTheMoonQueryPathReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationPos  *Vector                           `protobuf:"bytes,9,opt,name=destination_pos,json=destinationPos,proto3" json:"destination_pos,omitempty"`
	FuzzyRange      int32                             `protobuf:"varint,15,opt,name=fuzzy_range,json=fuzzyRange,proto3" json:"fuzzy_range,omitempty"`
	QueryType       ToTheMoonQueryPathReq_OptionType  `protobuf:"varint,8,opt,name=query_type,json=queryType,proto3,enum=proto.ToTheMoonQueryPathReq_OptionType" json:"query_type,omitempty"`
	AstarMethod     ToTheMoonQueryPathReq_AStarMethod `protobuf:"varint,1,opt,name=astar_method,json=astarMethod,proto3,enum=proto.ToTheMoonQueryPathReq_AStarMethod" json:"astar_method,omitempty"`
	SceneId         uint32                            `protobuf:"varint,6,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	QueryId         int32                             `protobuf:"varint,11,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	FilterType      ToTheMoonQueryPathReq_FilterType  `protobuf:"varint,3,opt,name=filter_type,json=filterType,proto3,enum=proto.ToTheMoonQueryPathReq_FilterType" json:"filter_type,omitempty"`
	Refined         bool                              `protobuf:"varint,13,opt,name=refined,proto3" json:"refined,omitempty"`
	UseFullNeighbor bool                              `protobuf:"varint,5,opt,name=use_full_neighbor,json=useFullNeighbor,proto3" json:"use_full_neighbor,omitempty"`
	SourcePos       *Vector                           `protobuf:"bytes,10,opt,name=source_pos,json=sourcePos,proto3" json:"source_pos,omitempty"`
}

func (x *ToTheMoonQueryPathReq) Reset() {
	*x = ToTheMoonQueryPathReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ToTheMoonQueryPathReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToTheMoonQueryPathReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToTheMoonQueryPathReq) ProtoMessage() {}

func (x *ToTheMoonQueryPathReq) ProtoReflect() protoreflect.Message {
	mi := &file_ToTheMoonQueryPathReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToTheMoonQueryPathReq.ProtoReflect.Descriptor instead.
func (*ToTheMoonQueryPathReq) Descriptor() ([]byte, []int) {
	return file_ToTheMoonQueryPathReq_proto_rawDescGZIP(), []int{0}
}

func (x *ToTheMoonQueryPathReq) GetDestinationPos() *Vector {
	if x != nil {
		return x.DestinationPos
	}
	return nil
}

func (x *ToTheMoonQueryPathReq) GetFuzzyRange() int32 {
	if x != nil {
		return x.FuzzyRange
	}
	return 0
}

func (x *ToTheMoonQueryPathReq) GetQueryType() ToTheMoonQueryPathReq_OptionType {
	if x != nil {
		return x.QueryType
	}
	return ToTheMoonQueryPathReq_OPTION_TYPE_NONE
}

func (x *ToTheMoonQueryPathReq) GetAstarMethod() ToTheMoonQueryPathReq_AStarMethod {
	if x != nil {
		return x.AstarMethod
	}
	return ToTheMoonQueryPathReq_A_STAR_METHOD_CLASSIC
}

func (x *ToTheMoonQueryPathReq) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *ToTheMoonQueryPathReq) GetQueryId() int32 {
	if x != nil {
		return x.QueryId
	}
	return 0
}

func (x *ToTheMoonQueryPathReq) GetFilterType() ToTheMoonQueryPathReq_FilterType {
	if x != nil {
		return x.FilterType
	}
	return ToTheMoonQueryPathReq_FILTER_TYPE_ALL
}

func (x *ToTheMoonQueryPathReq) GetRefined() bool {
	if x != nil {
		return x.Refined
	}
	return false
}

func (x *ToTheMoonQueryPathReq) GetUseFullNeighbor() bool {
	if x != nil {
		return x.UseFullNeighbor
	}
	return false
}

func (x *ToTheMoonQueryPathReq) GetSourcePos() *Vector {
	if x != nil {
		return x.SourcePos
	}
	return nil
}

var File_ToTheMoonQueryPathReq_proto protoreflect.FileDescriptor

var file_ToTheMoonQueryPathReq_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x54, 0x6f, 0x54, 0x68, 0x65, 0x4d, 0x6f, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x84, 0x06, 0x0a, 0x15, 0x54, 0x6f, 0x54, 0x68, 0x65, 0x4d, 0x6f, 0x6f, 0x6e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x36, 0x0a, 0x0f,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x75, 0x7a, 0x7a, 0x79, 0x5f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x75, 0x7a, 0x7a, 0x79,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x6f, 0x54, 0x68, 0x65, 0x4d, 0x6f, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a,
	0x0c, 0x61, 0x73, 0x74, 0x61, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x54, 0x68,
	0x65, 0x4d, 0x6f, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x2e, 0x41, 0x53, 0x74, 0x61, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0b, 0x61,
	0x73, 0x74, 0x61, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64,
	0x12, 0x48, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f,
	0x54, 0x68, 0x65, 0x4d, 0x6f, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x5f, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x75, 0x73, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72,
	0x12, 0x2c, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x22, 0x3a,
	0x0a, 0x0a, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10,
	0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x22, 0x7e, 0x0a, 0x0b, 0x41, 0x53,
	0x74, 0x61, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x5f, 0x53,
	0x54, 0x41, 0x52, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53,
	0x49, 0x43, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x5f, 0x4d,
	0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x54, 0x45, 0x4e, 0x44, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x01,
	0x12, 0x1a, 0x0a, 0x16, 0x41, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f,
	0x44, 0x5f, 0x41, 0x44, 0x41, 0x50, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18,
	0x41, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x49, 0x4e,
	0x46, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x22, 0x4d, 0x0a, 0x0a, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x4c, 0x54,
	0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x49, 0x52,
	0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x57, 0x41, 0x54, 0x45, 0x52, 0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ToTheMoonQueryPathReq_proto_rawDescOnce sync.Once
	file_ToTheMoonQueryPathReq_proto_rawDescData = file_ToTheMoonQueryPathReq_proto_rawDesc
)

func file_ToTheMoonQueryPathReq_proto_rawDescGZIP() []byte {
	file_ToTheMoonQueryPathReq_proto_rawDescOnce.Do(func() {
		file_ToTheMoonQueryPathReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ToTheMoonQueryPathReq_proto_rawDescData)
	})
	return file_ToTheMoonQueryPathReq_proto_rawDescData
}

var file_ToTheMoonQueryPathReq_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_ToTheMoonQueryPathReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ToTheMoonQueryPathReq_proto_goTypes = []interface{}{
	(ToTheMoonQueryPathReq_OptionType)(0),  // 0: proto.ToTheMoonQueryPathReq.OptionType
	(ToTheMoonQueryPathReq_AStarMethod)(0), // 1: proto.ToTheMoonQueryPathReq.AStarMethod
	(ToTheMoonQueryPathReq_FilterType)(0),  // 2: proto.ToTheMoonQueryPathReq.FilterType
	(*ToTheMoonQueryPathReq)(nil),          // 3: proto.ToTheMoonQueryPathReq
	(*Vector)(nil),                         // 4: proto.Vector
}
var file_ToTheMoonQueryPathReq_proto_depIdxs = []int32{
	4, // 0: proto.ToTheMoonQueryPathReq.destination_pos:type_name -> proto.Vector
	0, // 1: proto.ToTheMoonQueryPathReq.query_type:type_name -> proto.ToTheMoonQueryPathReq.OptionType
	1, // 2: proto.ToTheMoonQueryPathReq.astar_method:type_name -> proto.ToTheMoonQueryPathReq.AStarMethod
	2, // 3: proto.ToTheMoonQueryPathReq.filter_type:type_name -> proto.ToTheMoonQueryPathReq.FilterType
	4, // 4: proto.ToTheMoonQueryPathReq.source_pos:type_name -> proto.Vector
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ToTheMoonQueryPathReq_proto_init() }
func file_ToTheMoonQueryPathReq_proto_init() {
	if File_ToTheMoonQueryPathReq_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ToTheMoonQueryPathReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToTheMoonQueryPathReq); i {
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
			RawDescriptor: file_ToTheMoonQueryPathReq_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ToTheMoonQueryPathReq_proto_goTypes,
		DependencyIndexes: file_ToTheMoonQueryPathReq_proto_depIdxs,
		EnumInfos:         file_ToTheMoonQueryPathReq_proto_enumTypes,
		MessageInfos:      file_ToTheMoonQueryPathReq_proto_msgTypes,
	}.Build()
	File_ToTheMoonQueryPathReq_proto = out.File
	file_ToTheMoonQueryPathReq_proto_rawDesc = nil
	file_ToTheMoonQueryPathReq_proto_goTypes = nil
	file_ToTheMoonQueryPathReq_proto_depIdxs = nil
}
