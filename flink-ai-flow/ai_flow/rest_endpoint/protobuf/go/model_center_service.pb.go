//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model_center_service.proto

package ai_flow

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	//_ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateRegisteredModelRequest struct {
	RegisteredModel      *RegisteredModelParam `protobuf:"bytes,1,opt,name=registered_model,json=registeredModel,proto3" json:"registered_model,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateRegisteredModelRequest) Reset()         { *m = CreateRegisteredModelRequest{} }
func (m *CreateRegisteredModelRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRegisteredModelRequest) ProtoMessage()    {}
func (*CreateRegisteredModelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2c1b045eb52880, []int{0}
}

func (m *CreateRegisteredModelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRegisteredModelRequest.Unmarshal(m, b)
}
func (m *CreateRegisteredModelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRegisteredModelRequest.Marshal(b, m, deterministic)
}
func (m *CreateRegisteredModelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRegisteredModelRequest.Merge(m, src)
}
func (m *CreateRegisteredModelRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRegisteredModelRequest.Size(m)
}
func (m *CreateRegisteredModelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRegisteredModelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRegisteredModelRequest proto.InternalMessageInfo

func (m *CreateRegisteredModelRequest) GetRegisteredModel() *RegisteredModelParam {
	if m != nil {
		return m.RegisteredModel
	}
	return nil
}

type UpdateRegisteredModelRequest struct {
	ModelMeta            *ModelMetaParam       `protobuf:"bytes,1,opt,name=model_meta,json=modelMeta,proto3" json:"model_meta,omitempty"`
	RegisteredModel      *RegisteredModelParam `protobuf:"bytes,2,opt,name=registered_model,json=registeredModel,proto3" json:"registered_model,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateRegisteredModelRequest) Reset()         { *m = UpdateRegisteredModelRequest{} }
func (m *UpdateRegisteredModelRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRegisteredModelRequest) ProtoMessage()    {}
func (*UpdateRegisteredModelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2c1b045eb52880, []int{1}
}

func (m *UpdateRegisteredModelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRegisteredModelRequest.Unmarshal(m, b)
}
func (m *UpdateRegisteredModelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRegisteredModelRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRegisteredModelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRegisteredModelRequest.Merge(m, src)
}
func (m *UpdateRegisteredModelRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRegisteredModelRequest.Size(m)
}
func (m *UpdateRegisteredModelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRegisteredModelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRegisteredModelRequest proto.InternalMessageInfo

func (m *UpdateRegisteredModelRequest) GetModelMeta() *ModelMetaParam {
	if m != nil {
		return m.ModelMeta
	}
	return nil
}

func (m *UpdateRegisteredModelRequest) GetRegisteredModel() *RegisteredModelParam {
	if m != nil {
		return m.RegisteredModel
	}
	return nil
}

type DeleteRegisteredModelRequest struct {
	ModelMeta            *ModelMetaParam `protobuf:"bytes,1,opt,name=model_meta,json=modelMeta,proto3" json:"model_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DeleteRegisteredModelRequest) Reset()         { *m = DeleteRegisteredModelRequest{} }
func (m *DeleteRegisteredModelRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRegisteredModelRequest) ProtoMessage()    {}
func (*DeleteRegisteredModelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2c1b045eb52880, []int{2}
}

func (m *DeleteRegisteredModelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRegisteredModelRequest.Unmarshal(m, b)
}
func (m *DeleteRegisteredModelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRegisteredModelRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRegisteredModelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRegisteredModelRequest.Merge(m, src)
}
func (m *DeleteRegisteredModelRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRegisteredModelRequest.Size(m)
}
func (m *DeleteRegisteredModelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRegisteredModelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRegisteredModelRequest proto.InternalMessageInfo

func (m *DeleteRegisteredModelRequest) GetModelMeta() *ModelMetaParam {
	if m != nil {
		return m.ModelMeta
	}
	return nil
}

type ListRegisteredModelsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRegisteredModelsRequest) Reset()         { *m = ListRegisteredModelsRequest{} }
func (m *ListRegisteredModelsRequest) String() string { return proto.CompactTextString(m) }
func (*ListRegisteredModelsRequest) ProtoMessage()    {}
func (*ListRegisteredModelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2c1b045eb52880, []int{3}
}

func (m *ListRegisteredModelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRegisteredModelsRequest.Unmarshal(m, b)
}
func (m *ListRegisteredModelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRegisteredModelsRequest.Marshal(b, m, deterministic)
}
func (m *ListRegisteredModelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRegisteredModelsRequest.Merge(m, src)
}
func (m *ListRegisteredModelsRequest) XXX_Size() int {
	return xxx_messageInfo_ListRegisteredModelsRequest.Size(m)
}
func (m *ListRegisteredModelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRegisteredModelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRegisteredModelsRequest proto.InternalMessageInfo

type GetRegisteredModelDetailRequest struct {
	ModelMeta            *ModelMetaParam `protobuf:"bytes,1,opt,name=model_meta,json=modelMeta,proto3" json:"model_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetRegisteredModelDetailRequest) Reset()         { *m = GetRegisteredModelDetailRequest{} }
func (m *GetRegisteredModelDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetRegisteredModelDetailRequest) ProtoMessage()    {}
func (*GetRegisteredModelDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2c1b045eb52880, []int{4}
}

func (m *GetRegisteredModelDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRegisteredModelDetailRequest.Unmarshal(m, b)
}
func (m *GetRegisteredModelDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRegisteredModelDetailRequest.Marshal(b, m, deterministic)
}
func (m *GetRegisteredModelDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRegisteredModelDetailRequest.Merge(m, src)
}
func (m *GetRegisteredModelDetailRequest) XXX_Size() int {
	return xxx_messageInfo_GetRegisteredModelDetailRequest.Size(m)
}
func (m *GetRegisteredModelDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRegisteredModelDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRegisteredModelDetailRequest proto.InternalMessageInfo

func (m *GetRegisteredModelDetailRequest) GetModelMeta() *ModelMetaParam {
	if m != nil {
		return m.ModelMeta
	}
	return nil
}

type CreateModelVersionRequest struct {
	ModelMeta            *ModelMetaParam    `protobuf:"bytes,1,opt,name=model_meta,json=modelMeta,proto3" json:"model_meta,omitempty"`
	ModelVersion         *ModelVersionParam `protobuf:"bytes,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateModelVersionRequest) Reset()         { *m = CreateModelVersionRequest{} }
func (m *CreateModelVersionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateModelVersionRequest) ProtoMessage()    {}
func (*CreateModelVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2c1b045eb52880, []int{5}
}

func (m *CreateModelVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateModelVersionRequest.Unmarshal(m, b)
}
func (m *CreateModelVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateModelVersionRequest.Marshal(b, m, deterministic)
}
func (m *CreateModelVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateModelVersionRequest.Merge(m, src)
}
func (m *CreateModelVersionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateModelVersionRequest.Size(m)
}
func (m *CreateModelVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateModelVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateModelVersionRequest proto.InternalMessageInfo

func (m *CreateModelVersionRequest) GetModelMeta() *ModelMetaParam {
	if m != nil {
		return m.ModelMeta
	}
	return nil
}

func (m *CreateModelVersionRequest) GetModelVersion() *ModelVersionParam {
	if m != nil {
		return m.ModelVersion
	}
	return nil
}

type UpdateModelVersionRequest struct {
	ModelMeta            *ModelMetaParam    `protobuf:"bytes,1,opt,name=model_meta,json=modelMeta,proto3" json:"model_meta,omitempty"`
	ModelVersion         *ModelVersionParam `protobuf:"bytes,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateModelVersionRequest) Reset()         { *m = UpdateModelVersionRequest{} }
func (m *UpdateModelVersionRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateModelVersionRequest) ProtoMessage()    {}
func (*UpdateModelVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2c1b045eb52880, []int{6}
}

func (m *UpdateModelVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateModelVersionRequest.Unmarshal(m, b)
}
func (m *UpdateModelVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateModelVersionRequest.Marshal(b, m, deterministic)
}
func (m *UpdateModelVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateModelVersionRequest.Merge(m, src)
}
func (m *UpdateModelVersionRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateModelVersionRequest.Size(m)
}
func (m *UpdateModelVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateModelVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateModelVersionRequest proto.InternalMessageInfo

func (m *UpdateModelVersionRequest) GetModelMeta() *ModelMetaParam {
	if m != nil {
		return m.ModelMeta
	}
	return nil
}

func (m *UpdateModelVersionRequest) GetModelVersion() *ModelVersionParam {
	if m != nil {
		return m.ModelVersion
	}
	return nil
}

type DeleteModelVersionRequest struct {
	ModelMeta            *ModelMetaParam `protobuf:"bytes,1,opt,name=model_meta,json=modelMeta,proto3" json:"model_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DeleteModelVersionRequest) Reset()         { *m = DeleteModelVersionRequest{} }
func (m *DeleteModelVersionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteModelVersionRequest) ProtoMessage()    {}
func (*DeleteModelVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2c1b045eb52880, []int{7}
}

func (m *DeleteModelVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteModelVersionRequest.Unmarshal(m, b)
}
func (m *DeleteModelVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteModelVersionRequest.Marshal(b, m, deterministic)
}
func (m *DeleteModelVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteModelVersionRequest.Merge(m, src)
}
func (m *DeleteModelVersionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteModelVersionRequest.Size(m)
}
func (m *DeleteModelVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteModelVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteModelVersionRequest proto.InternalMessageInfo

func (m *DeleteModelVersionRequest) GetModelMeta() *ModelMetaParam {
	if m != nil {
		return m.ModelMeta
	}
	return nil
}

type GetModelVersionDetailRequest struct {
	ModelMeta            *ModelMetaParam `protobuf:"bytes,1,opt,name=model_meta,json=modelMeta,proto3" json:"model_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetModelVersionDetailRequest) Reset()         { *m = GetModelVersionDetailRequest{} }
func (m *GetModelVersionDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetModelVersionDetailRequest) ProtoMessage()    {}
func (*GetModelVersionDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2c1b045eb52880, []int{8}
}

func (m *GetModelVersionDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModelVersionDetailRequest.Unmarshal(m, b)
}
func (m *GetModelVersionDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModelVersionDetailRequest.Marshal(b, m, deterministic)
}
func (m *GetModelVersionDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModelVersionDetailRequest.Merge(m, src)
}
func (m *GetModelVersionDetailRequest) XXX_Size() int {
	return xxx_messageInfo_GetModelVersionDetailRequest.Size(m)
}
func (m *GetModelVersionDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModelVersionDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetModelVersionDetailRequest proto.InternalMessageInfo

func (m *GetModelVersionDetailRequest) GetModelMeta() *ModelMetaParam {
	if m != nil {
		return m.ModelMeta
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRegisteredModelRequest)(nil), "ai_flow.CreateRegisteredModelRequest")
	proto.RegisterType((*UpdateRegisteredModelRequest)(nil), "ai_flow.UpdateRegisteredModelRequest")
	proto.RegisterType((*DeleteRegisteredModelRequest)(nil), "ai_flow.DeleteRegisteredModelRequest")
	proto.RegisterType((*ListRegisteredModelsRequest)(nil), "ai_flow.ListRegisteredModelsRequest")
	proto.RegisterType((*GetRegisteredModelDetailRequest)(nil), "ai_flow.GetRegisteredModelDetailRequest")
	proto.RegisterType((*CreateModelVersionRequest)(nil), "ai_flow.CreateModelVersionRequest")
	proto.RegisterType((*UpdateModelVersionRequest)(nil), "ai_flow.UpdateModelVersionRequest")
	proto.RegisterType((*DeleteModelVersionRequest)(nil), "ai_flow.DeleteModelVersionRequest")
	proto.RegisterType((*GetModelVersionDetailRequest)(nil), "ai_flow.GetModelVersionDetailRequest")
}

func init() { proto.RegisterFile("model_center_service.proto", fileDescriptor_ee2c1b045eb52880) }

var fileDescriptor_ee2c1b045eb52880 = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xe5, 0x1d, 0x40, 0x33, 0x4c, 0x0c, 0x8b, 0x89, 0x2e, 0x74, 0x02, 0x45, 0xfc, 0xe9,
	0x2a, 0x2d, 0x41, 0xe5, 0xcf, 0x81, 0x0b, 0x12, 0x9b, 0x34, 0x0e, 0x4c, 0x42, 0x9d, 0x98, 0xc4,
	0xa9, 0x32, 0xc9, 0x4b, 0x88, 0x94, 0xc4, 0xc5, 0x76, 0xc7, 0x19, 0x4e, 0x3b, 0x22, 0xc1, 0x07,
	0xe0, 0x3b, 0xc1, 0x47, 0xe0, 0x83, 0xa0, 0xf8, 0x4d, 0x23, 0xb7, 0x75, 0x43, 0x44, 0x7b, 0xe0,
	0x18, 0xbf, 0xf6, 0xfb, 0xfc, 0x6a, 0xbf, 0xcf, 0xa3, 0x52, 0x2f, 0x17, 0x31, 0x64, 0xa3, 0x08,
	0x0a, 0x0d, 0x72, 0xa4, 0x40, 0x9e, 0xa7, 0x11, 0x04, 0x63, 0x29, 0xb4, 0x60, 0x97, 0x79, 0x3a,
	0x7a, 0x9f, 0x89, 0x4f, 0xde, 0x56, 0x0e, 0x4a, 0xf1, 0xa4, 0x5a, 0xf7, 0xba, 0x89, 0x10, 0x49,
	0x06, 0x21, 0x1f, 0xa7, 0x21, 0x2f, 0x0a, 0xa1, 0xb9, 0x4e, 0x45, 0xa1, 0xb0, 0xea, 0x7f, 0xa0,
	0xdd, 0x43, 0x09, 0x5c, 0xc3, 0x10, 0x92, 0x54, 0x69, 0x90, 0x10, 0x9f, 0x94, 0x1a, 0x43, 0xf8,
	0x38, 0x01, 0xa5, 0xd9, 0x4b, 0xba, 0x2d, 0xeb, 0xca, 0xc8, 0xc8, 0x77, 0xc8, 0x1d, 0xd2, 0xbb,
	0x32, 0xd8, 0x0b, 0x2a, 0xc1, 0x60, 0xee, 0xe8, 0x6b, 0x2e, 0x79, 0x3e, 0xbc, 0x26, 0x67, 0x57,
	0xfd, 0x1f, 0x84, 0x76, 0xdf, 0x8c, 0xe3, 0xe5, 0x52, 0x4f, 0x29, 0xc5, 0x9f, 0x97, 0x83, 0xe6,
	0x95, 0xc8, 0xcd, 0x5a, 0xc4, 0x6c, 0x3d, 0x01, 0xcd, 0xb1, 0xfd, 0x66, 0x3e, 0xfd, 0x76, 0x22,
	0x6e, 0xfc, 0x13, 0xe2, 0x19, 0xed, 0x1e, 0x41, 0x06, 0xeb, 0x26, 0xf4, 0xf7, 0xe8, 0xad, 0x57,
	0xa9, 0xd2, 0x73, 0x5d, 0x55, 0xd5, 0xd6, 0x7f, 0x4b, 0x6f, 0x1f, 0xc3, 0x7c, 0xf5, 0x08, 0x34,
	0x4f, 0x57, 0x56, 0xfe, 0x4e, 0xe8, 0x2e, 0xbe, 0xaf, 0xd9, 0x73, 0x06, 0x52, 0xa5, 0xa2, 0x58,
	0xf5, 0xc6, 0x9f, 0xd3, 0x2d, 0x3c, 0x77, 0x8e, 0xfd, 0xaa, 0xeb, 0xf6, 0x66, 0x8f, 0x56, 0x62,
	0x78, 0xfa, 0x6a, 0x6e, 0x2d, 0x19, 0x2c, 0x9c, 0x85, 0xff, 0x0a, 0xeb, 0x94, 0xee, 0xe2, 0xfb,
	0xaf, 0x91, 0xaa, 0x1c, 0xaa, 0x63, 0xd0, 0x76, 0xc7, 0xb5, 0x3c, 0xed, 0xe0, 0xd7, 0x26, 0x65,
	0xa6, 0x7a, 0x68, 0xd2, 0xe0, 0x14, 0xc3, 0x80, 0x5d, 0x10, 0xba, 0x13, 0xb9, 0x1c, 0xcd, 0xee,
	0xd5, 0x4d, 0x9b, 0x1c, 0xef, 0x5d, 0xb7, 0x4c, 0xa3, 0xc6, 0xa2, 0x50, 0xe0, 0x3f, 0xf9, 0xf2,
	0xf3, 0xf7, 0xb7, 0x8d, 0xd0, 0xef, 0x87, 0x3c, 0x2d, 0x2b, 0xa1, 0xa1, 0x38, 0xc0, 0x18, 0x0a,
	0xa7, 0x2e, 0x3a, 0x30, 0xab, 0x21, 0x8a, 0x3f, 0x23, 0x7d, 0x83, 0x32, 0x71, 0x39, 0xde, 0x42,
	0x69, 0x4a, 0x84, 0x06, 0x94, 0x41, 0x2b, 0x14, 0x14, 0x9f, 0xa2, 0xc4, 0x2e, 0x6b, 0x5b, 0x28,
	0x4d, 0xd6, 0x5f, 0x23, 0xca, 0x67, 0x42, 0x6f, 0x64, 0x8e, 0x34, 0x60, 0x77, 0x6b, 0x89, 0x86,
	0xb0, 0x70, 0x81, 0x3c, 0x34, 0x20, 0x7d, 0xd6, 0x6b, 0x03, 0x52, 0x4a, 0x97, 0xd7, 0xd1, 0x49,
	0x96, 0x44, 0x0e, 0xeb, 0xd5, 0x0a, 0x7f, 0x49, 0x25, 0x17, 0x4b, 0x68, 0x58, 0xf6, 0xd9, 0x83,
	0x36, 0x2c, 0x09, 0xe8, 0xf2, 0x3a, 0x58, 0xb4, 0x90, 0x50, 0xcc, 0x9f, 0x1b, 0x56, 0x87, 0x23,
	0x5d, 0xf2, 0x8f, 0x8d, 0x7c, 0xe0, 0xef, 0x3b, 0xe5, 0xf1, 0xa3, 0x4a, 0x07, 0x6b, 0x50, 0x4b,
	0x86, 0xc9, 0x42, 0x1c, 0x59, 0x0c, 0x4b, 0xb3, 0xaa, 0x81, 0x61, 0xd0, 0x86, 0x61, 0x66, 0x2c,
	0x58, 0xbc, 0x10, 0x3e, 0x16, 0xc3, 0xd2, 0x64, 0x5a, 0x95, 0x01, 0x55, 0x2b, 0x86, 0x9d, 0xc4,
	0x95, 0x55, 0x96, 0x4b, 0x9a, 0xb2, 0xcc, 0x45, 0x12, 0x18, 0x92, 0x1e, 0xbb, 0xdf, 0x82, 0x24,
	0x01, 0xfd, 0xa2, 0x43, 0xb7, 0x23, 0x91, 0x07, 0xb8, 0x19, 0xff, 0xa4, 0x5c, 0x10, 0xf2, 0x95,
	0x90, 0x77, 0x97, 0xcc, 0xc7, 0xa3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x8d, 0x6b, 0x5a,
	0x05, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ModelCenterServiceClient is the client API for ModelCenterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModelCenterServiceClient interface {
	// Create registered model with metadata of RegisteredModel.
	CreateRegisteredModel(ctx context.Context, in *CreateRegisteredModelRequest, opts ...grpc.CallOption) (*Response, error)
	// Update registered model with metadata of RegisteredModel.
	UpdateRegisteredModel(ctx context.Context, in *UpdateRegisteredModelRequest, opts ...grpc.CallOption) (*Response, error)
	// Delete registered model with metadata of RegisteredModel.
	DeleteRegisteredModel(ctx context.Context, in *DeleteRegisteredModelRequest, opts ...grpc.CallOption) (*Response, error)
	// List registered models about metadata of RegisteredModel.
	ListRegisteredModels(ctx context.Context, in *ListRegisteredModelsRequest, opts ...grpc.CallOption) (*Response, error)
	// Get registered model detail including metadata of RegisteredModel.
	GetRegisteredModelDetail(ctx context.Context, in *GetRegisteredModelDetailRequest, opts ...grpc.CallOption) (*Response, error)
	// Create model version with metadata of ModelVersion.
	CreateModelVersion(ctx context.Context, in *CreateModelVersionRequest, opts ...grpc.CallOption) (*Response, error)
	// Update model version with metadata of ModelVersion.
	UpdateModelVersion(ctx context.Context, in *UpdateModelVersionRequest, opts ...grpc.CallOption) (*Response, error)
	// Delete model version with metadata of ModelVersion.
	DeleteModelVersion(ctx context.Context, in *DeleteModelVersionRequest, opts ...grpc.CallOption) (*Response, error)
	// Get model version detail with metadata of ModelVersion.
	GetModelVersionDetail(ctx context.Context, in *GetModelVersionDetailRequest, opts ...grpc.CallOption) (*Response, error)
}

type modelCenterServiceClient struct {
	cc *grpc.ClientConn
}

func NewModelCenterServiceClient(cc *grpc.ClientConn) ModelCenterServiceClient {
	return &modelCenterServiceClient{cc}
}

func (c *modelCenterServiceClient) CreateRegisteredModel(ctx context.Context, in *CreateRegisteredModelRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ai_flow.ModelCenterService/createRegisteredModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelCenterServiceClient) UpdateRegisteredModel(ctx context.Context, in *UpdateRegisteredModelRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ai_flow.ModelCenterService/updateRegisteredModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelCenterServiceClient) DeleteRegisteredModel(ctx context.Context, in *DeleteRegisteredModelRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ai_flow.ModelCenterService/deleteRegisteredModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelCenterServiceClient) ListRegisteredModels(ctx context.Context, in *ListRegisteredModelsRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ai_flow.ModelCenterService/listRegisteredModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelCenterServiceClient) GetRegisteredModelDetail(ctx context.Context, in *GetRegisteredModelDetailRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ai_flow.ModelCenterService/getRegisteredModelDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelCenterServiceClient) CreateModelVersion(ctx context.Context, in *CreateModelVersionRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ai_flow.ModelCenterService/createModelVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelCenterServiceClient) UpdateModelVersion(ctx context.Context, in *UpdateModelVersionRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ai_flow.ModelCenterService/updateModelVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelCenterServiceClient) DeleteModelVersion(ctx context.Context, in *DeleteModelVersionRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ai_flow.ModelCenterService/deleteModelVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelCenterServiceClient) GetModelVersionDetail(ctx context.Context, in *GetModelVersionDetailRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ai_flow.ModelCenterService/getModelVersionDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelCenterServiceServer is the server API for ModelCenterService service.
type ModelCenterServiceServer interface {
	// Create registered model with metadata of RegisteredModel.
	CreateRegisteredModel(context.Context, *CreateRegisteredModelRequest) (*Response, error)
	// Update registered model with metadata of RegisteredModel.
	UpdateRegisteredModel(context.Context, *UpdateRegisteredModelRequest) (*Response, error)
	// Delete registered model with metadata of RegisteredModel.
	DeleteRegisteredModel(context.Context, *DeleteRegisteredModelRequest) (*Response, error)
	// List registered models about metadata of RegisteredModel.
	ListRegisteredModels(context.Context, *ListRegisteredModelsRequest) (*Response, error)
	// Get registered model detail including metadata of RegisteredModel.
	GetRegisteredModelDetail(context.Context, *GetRegisteredModelDetailRequest) (*Response, error)
	// Create model version with metadata of ModelVersion.
	CreateModelVersion(context.Context, *CreateModelVersionRequest) (*Response, error)
	// Update model version with metadata of ModelVersion.
	UpdateModelVersion(context.Context, *UpdateModelVersionRequest) (*Response, error)
	// Delete model version with metadata of ModelVersion.
	DeleteModelVersion(context.Context, *DeleteModelVersionRequest) (*Response, error)
	// Get model version detail with metadata of ModelVersion.
	GetModelVersionDetail(context.Context, *GetModelVersionDetailRequest) (*Response, error)
}

// UnimplementedModelCenterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedModelCenterServiceServer struct {
}

func (*UnimplementedModelCenterServiceServer) CreateRegisteredModel(ctx context.Context, req *CreateRegisteredModelRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegisteredModel not implemented")
}
func (*UnimplementedModelCenterServiceServer) UpdateRegisteredModel(ctx context.Context, req *UpdateRegisteredModelRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRegisteredModel not implemented")
}
func (*UnimplementedModelCenterServiceServer) DeleteRegisteredModel(ctx context.Context, req *DeleteRegisteredModelRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegisteredModel not implemented")
}
func (*UnimplementedModelCenterServiceServer) ListRegisteredModels(ctx context.Context, req *ListRegisteredModelsRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRegisteredModels not implemented")
}
func (*UnimplementedModelCenterServiceServer) GetRegisteredModelDetail(ctx context.Context, req *GetRegisteredModelDetailRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisteredModelDetail not implemented")
}
func (*UnimplementedModelCenterServiceServer) CreateModelVersion(ctx context.Context, req *CreateModelVersionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModelVersion not implemented")
}
func (*UnimplementedModelCenterServiceServer) UpdateModelVersion(ctx context.Context, req *UpdateModelVersionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateModelVersion not implemented")
}
func (*UnimplementedModelCenterServiceServer) DeleteModelVersion(ctx context.Context, req *DeleteModelVersionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteModelVersion not implemented")
}
func (*UnimplementedModelCenterServiceServer) GetModelVersionDetail(ctx context.Context, req *GetModelVersionDetailRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelVersionDetail not implemented")
}

func RegisterModelCenterServiceServer(s *grpc.Server, srv ModelCenterServiceServer) {
	s.RegisterService(&_ModelCenterService_serviceDesc, srv)
}

func _ModelCenterService_CreateRegisteredModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegisteredModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelCenterServiceServer).CreateRegisteredModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai_flow.ModelCenterService/CreateRegisteredModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelCenterServiceServer).CreateRegisteredModel(ctx, req.(*CreateRegisteredModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelCenterService_UpdateRegisteredModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegisteredModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelCenterServiceServer).UpdateRegisteredModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai_flow.ModelCenterService/UpdateRegisteredModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelCenterServiceServer).UpdateRegisteredModel(ctx, req.(*UpdateRegisteredModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelCenterService_DeleteRegisteredModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegisteredModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelCenterServiceServer).DeleteRegisteredModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai_flow.ModelCenterService/DeleteRegisteredModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelCenterServiceServer).DeleteRegisteredModel(ctx, req.(*DeleteRegisteredModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelCenterService_ListRegisteredModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegisteredModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelCenterServiceServer).ListRegisteredModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai_flow.ModelCenterService/ListRegisteredModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelCenterServiceServer).ListRegisteredModels(ctx, req.(*ListRegisteredModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelCenterService_GetRegisteredModelDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegisteredModelDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelCenterServiceServer).GetRegisteredModelDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai_flow.ModelCenterService/GetRegisteredModelDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelCenterServiceServer).GetRegisteredModelDetail(ctx, req.(*GetRegisteredModelDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelCenterService_CreateModelVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModelVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelCenterServiceServer).CreateModelVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai_flow.ModelCenterService/CreateModelVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelCenterServiceServer).CreateModelVersion(ctx, req.(*CreateModelVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelCenterService_UpdateModelVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateModelVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelCenterServiceServer).UpdateModelVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai_flow.ModelCenterService/UpdateModelVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelCenterServiceServer).UpdateModelVersion(ctx, req.(*UpdateModelVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelCenterService_DeleteModelVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteModelVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelCenterServiceServer).DeleteModelVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai_flow.ModelCenterService/DeleteModelVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelCenterServiceServer).DeleteModelVersion(ctx, req.(*DeleteModelVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelCenterService_GetModelVersionDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelVersionDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelCenterServiceServer).GetModelVersionDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai_flow.ModelCenterService/GetModelVersionDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelCenterServiceServer).GetModelVersionDetail(ctx, req.(*GetModelVersionDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ModelCenterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai_flow.ModelCenterService",
	HandlerType: (*ModelCenterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createRegisteredModel",
			Handler:    _ModelCenterService_CreateRegisteredModel_Handler,
		},
		{
			MethodName: "updateRegisteredModel",
			Handler:    _ModelCenterService_UpdateRegisteredModel_Handler,
		},
		{
			MethodName: "deleteRegisteredModel",
			Handler:    _ModelCenterService_DeleteRegisteredModel_Handler,
		},
		{
			MethodName: "listRegisteredModels",
			Handler:    _ModelCenterService_ListRegisteredModels_Handler,
		},
		{
			MethodName: "getRegisteredModelDetail",
			Handler:    _ModelCenterService_GetRegisteredModelDetail_Handler,
		},
		{
			MethodName: "createModelVersion",
			Handler:    _ModelCenterService_CreateModelVersion_Handler,
		},
		{
			MethodName: "updateModelVersion",
			Handler:    _ModelCenterService_UpdateModelVersion_Handler,
		},
		{
			MethodName: "deleteModelVersion",
			Handler:    _ModelCenterService_DeleteModelVersion_Handler,
		},
		{
			MethodName: "getModelVersionDetail",
			Handler:    _ModelCenterService_GetModelVersionDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model_center_service.proto",
}
