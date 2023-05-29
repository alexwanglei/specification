// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config_file.proto

package config_manage // import "github.com/polarismesh/specification/source/go/api/v1/config_manage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConfigFileGroup struct {
	Id                   *wrapperspb.UInt64Value   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrapperspb.StringValue   `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Comment              *wrapperspb.StringValue   `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	CreateTime           *wrapperspb.StringValue   `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrapperspb.StringValue   `protobuf:"bytes,6,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrapperspb.StringValue   `protobuf:"bytes,7,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrapperspb.StringValue   `protobuf:"bytes,8,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	FileCount            *wrapperspb.UInt64Value   `protobuf:"bytes,9,opt,name=fileCount,proto3" json:"fileCount,omitempty"`
	UserIds              []*wrapperspb.StringValue `protobuf:"bytes,10,rep,name=user_ids,proto3" json:"user_ids,omitempty"`
	GroupIds             []*wrapperspb.StringValue `protobuf:"bytes,11,rep,name=group_ids,proto3" json:"group_ids,omitempty"`
	RemoveUserIds        []*wrapperspb.StringValue `protobuf:"bytes,13,rep,name=remove_user_ids,proto3" json:"remove_user_ids,omitempty"`
	RemoveGroupIds       []*wrapperspb.StringValue `protobuf:"bytes,14,rep,name=remove_group_ids,proto3" json:"remove_group_ids,omitempty"`
	Editable             *wrapperspb.BoolValue     `protobuf:"bytes,15,opt,name=editable,proto3" json:"editable,omitempty"`
	Owner                *wrapperspb.StringValue   `protobuf:"bytes,16,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ConfigFileGroup) Reset()         { *m = ConfigFileGroup{} }
func (m *ConfigFileGroup) String() string { return proto.CompactTextString(m) }
func (*ConfigFileGroup) ProtoMessage()    {}
func (*ConfigFileGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b9a411930c31ddf, []int{0}
}
func (m *ConfigFileGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileGroup.Unmarshal(m, b)
}
func (m *ConfigFileGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileGroup.Marshal(b, m, deterministic)
}
func (dst *ConfigFileGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileGroup.Merge(dst, src)
}
func (m *ConfigFileGroup) XXX_Size() int {
	return xxx_messageInfo_ConfigFileGroup.Size(m)
}
func (m *ConfigFileGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileGroup proto.InternalMessageInfo

func (m *ConfigFileGroup) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileGroup) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileGroup) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileGroup) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileGroup) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileGroup) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileGroup) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileGroup) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

func (m *ConfigFileGroup) GetFileCount() *wrapperspb.UInt64Value {
	if m != nil {
		return m.FileCount
	}
	return nil
}

func (m *ConfigFileGroup) GetUserIds() []*wrapperspb.StringValue {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *ConfigFileGroup) GetGroupIds() []*wrapperspb.StringValue {
	if m != nil {
		return m.GroupIds
	}
	return nil
}

func (m *ConfigFileGroup) GetRemoveUserIds() []*wrapperspb.StringValue {
	if m != nil {
		return m.RemoveUserIds
	}
	return nil
}

func (m *ConfigFileGroup) GetRemoveGroupIds() []*wrapperspb.StringValue {
	if m != nil {
		return m.RemoveGroupIds
	}
	return nil
}

func (m *ConfigFileGroup) GetEditable() *wrapperspb.BoolValue {
	if m != nil {
		return m.Editable
	}
	return nil
}

func (m *ConfigFileGroup) GetOwner() *wrapperspb.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

type ConfigFile struct {
	Id          *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace   *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group       *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	Content     *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Format      *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=format,proto3" json:"format,omitempty"`
	Comment     *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	Status      *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Tags        []*ConfigFileTag        `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	CreateTime  *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy    *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime  *wrapperspb.StringValue `protobuf:"bytes,12,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy    *wrapperspb.StringValue `protobuf:"bytes,13,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	ReleaseTime *wrapperspb.StringValue `protobuf:"bytes,14,opt,name=release_time,json=releaseTime,proto3" json:"release_time,omitempty"`
	ReleaseBy   *wrapperspb.StringValue `protobuf:"bytes,15,opt,name=release_by,json=releaseBy,proto3" json:"release_by,omitempty"`
	// 是否为加密配置文件
	Encrypted *wrapperspb.BoolValue `protobuf:"bytes,16,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	// 加密算法
	EncryptAlgo          *wrapperspb.StringValue `protobuf:"bytes,17,opt,name=encrypt_algo,json=encryptAlgo,proto3" json:"encrypt_algo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFile) Reset()         { *m = ConfigFile{} }
func (m *ConfigFile) String() string { return proto.CompactTextString(m) }
func (*ConfigFile) ProtoMessage()    {}
func (*ConfigFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b9a411930c31ddf, []int{1}
}
func (m *ConfigFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFile.Unmarshal(m, b)
}
func (m *ConfigFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFile.Marshal(b, m, deterministic)
}
func (dst *ConfigFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFile.Merge(dst, src)
}
func (m *ConfigFile) XXX_Size() int {
	return xxx_messageInfo_ConfigFile.Size(m)
}
func (m *ConfigFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFile.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFile proto.InternalMessageInfo

func (m *ConfigFile) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFile) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFile) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFile) GetGroup() *wrapperspb.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ConfigFile) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFile) GetFormat() *wrapperspb.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigFile) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFile) GetStatus() *wrapperspb.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ConfigFile) GetTags() []*ConfigFileTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ConfigFile) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFile) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFile) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFile) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

func (m *ConfigFile) GetReleaseTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ReleaseTime
	}
	return nil
}

func (m *ConfigFile) GetReleaseBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ReleaseBy
	}
	return nil
}

func (m *ConfigFile) GetEncrypted() *wrapperspb.BoolValue {
	if m != nil {
		return m.Encrypted
	}
	return nil
}

func (m *ConfigFile) GetEncryptAlgo() *wrapperspb.StringValue {
	if m != nil {
		return m.EncryptAlgo
	}
	return nil
}

type ConfigFileTag struct {
	Key                  *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFileTag) Reset()         { *m = ConfigFileTag{} }
func (m *ConfigFileTag) String() string { return proto.CompactTextString(m) }
func (*ConfigFileTag) ProtoMessage()    {}
func (*ConfigFileTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b9a411930c31ddf, []int{2}
}
func (m *ConfigFileTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileTag.Unmarshal(m, b)
}
func (m *ConfigFileTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileTag.Marshal(b, m, deterministic)
}
func (dst *ConfigFileTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileTag.Merge(dst, src)
}
func (m *ConfigFileTag) XXX_Size() int {
	return xxx_messageInfo_ConfigFileTag.Size(m)
}
func (m *ConfigFileTag) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileTag.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileTag proto.InternalMessageInfo

func (m *ConfigFileTag) GetKey() *wrapperspb.StringValue {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ConfigFileTag) GetValue() *wrapperspb.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type ConfigFileRelease struct {
	Id                   *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group                *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	FileName             *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content              *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Comment              *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	Md5                  *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=md5,proto3" json:"md5,omitempty"`
	Version              *wrapperspb.UInt64Value `protobuf:"bytes,9,opt,name=version,proto3" json:"version,omitempty"`
	CreateTime           *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrapperspb.StringValue `protobuf:"bytes,12,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrapperspb.StringValue `protobuf:"bytes,13,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFileRelease) Reset()         { *m = ConfigFileRelease{} }
func (m *ConfigFileRelease) String() string { return proto.CompactTextString(m) }
func (*ConfigFileRelease) ProtoMessage()    {}
func (*ConfigFileRelease) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b9a411930c31ddf, []int{3}
}
func (m *ConfigFileRelease) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileRelease.Unmarshal(m, b)
}
func (m *ConfigFileRelease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileRelease.Marshal(b, m, deterministic)
}
func (dst *ConfigFileRelease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileRelease.Merge(dst, src)
}
func (m *ConfigFileRelease) XXX_Size() int {
	return xxx_messageInfo_ConfigFileRelease.Size(m)
}
func (m *ConfigFileRelease) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileRelease.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileRelease proto.InternalMessageInfo

func (m *ConfigFileRelease) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileRelease) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileRelease) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileRelease) GetGroup() *wrapperspb.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ConfigFileRelease) GetFileName() *wrapperspb.StringValue {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *ConfigFileRelease) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFileRelease) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileRelease) GetMd5() *wrapperspb.StringValue {
	if m != nil {
		return m.Md5
	}
	return nil
}

func (m *ConfigFileRelease) GetVersion() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ConfigFileRelease) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileRelease) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileRelease) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileRelease) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

type ConfigFileReleaseHistory struct {
	Id                   *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group                *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	FileName             *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content              *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Format               *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=format,proto3" json:"format,omitempty"`
	Comment              *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
	Md5                  *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=md5,proto3" json:"md5,omitempty"`
	Type                 *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Status               *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 []*ConfigFileTag        `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	CreateTime           *wrapperspb.StringValue `protobuf:"bytes,13,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrapperspb.StringValue `protobuf:"bytes,14,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrapperspb.StringValue `protobuf:"bytes,15,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrapperspb.StringValue `protobuf:"bytes,16,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFileReleaseHistory) Reset()         { *m = ConfigFileReleaseHistory{} }
func (m *ConfigFileReleaseHistory) String() string { return proto.CompactTextString(m) }
func (*ConfigFileReleaseHistory) ProtoMessage()    {}
func (*ConfigFileReleaseHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b9a411930c31ddf, []int{4}
}
func (m *ConfigFileReleaseHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileReleaseHistory.Unmarshal(m, b)
}
func (m *ConfigFileReleaseHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileReleaseHistory.Marshal(b, m, deterministic)
}
func (dst *ConfigFileReleaseHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileReleaseHistory.Merge(dst, src)
}
func (m *ConfigFileReleaseHistory) XXX_Size() int {
	return xxx_messageInfo_ConfigFileReleaseHistory.Size(m)
}
func (m *ConfigFileReleaseHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileReleaseHistory.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileReleaseHistory proto.InternalMessageInfo

func (m *ConfigFileReleaseHistory) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetGroup() *wrapperspb.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetFileName() *wrapperspb.StringValue {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetFormat() *wrapperspb.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetMd5() *wrapperspb.StringValue {
	if m != nil {
		return m.Md5
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetType() *wrapperspb.StringValue {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetStatus() *wrapperspb.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetTags() []*ConfigFileTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

type ConfigFileTemplate struct {
	Id                   *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Content              *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Format               *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=format,proto3" json:"format,omitempty"`
	Comment              *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	CreateTime           *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFileTemplate) Reset()         { *m = ConfigFileTemplate{} }
func (m *ConfigFileTemplate) String() string { return proto.CompactTextString(m) }
func (*ConfigFileTemplate) ProtoMessage()    {}
func (*ConfigFileTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b9a411930c31ddf, []int{5}
}
func (m *ConfigFileTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileTemplate.Unmarshal(m, b)
}
func (m *ConfigFileTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileTemplate.Marshal(b, m, deterministic)
}
func (dst *ConfigFileTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileTemplate.Merge(dst, src)
}
func (m *ConfigFileTemplate) XXX_Size() int {
	return xxx_messageInfo_ConfigFileTemplate.Size(m)
}
func (m *ConfigFileTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileTemplate proto.InternalMessageInfo

func (m *ConfigFileTemplate) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileTemplate) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileTemplate) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFileTemplate) GetFormat() *wrapperspb.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigFileTemplate) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileTemplate) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileTemplate) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileTemplate) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileTemplate) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

type ClientConfigFileInfo struct {
	Namespace *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group     *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	FileName  *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content   *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Version   *wrapperspb.UInt64Value `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Md5       *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=md5,proto3" json:"md5,omitempty"`
	Tags      []*ConfigFileTag        `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	// 是否为加密配置文件
	Encrypted *wrapperspb.BoolValue `protobuf:"bytes,8,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	// 公钥，用于加密数据密钥
	PublicKey            *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClientConfigFileInfo) Reset()         { *m = ClientConfigFileInfo{} }
func (m *ClientConfigFileInfo) String() string { return proto.CompactTextString(m) }
func (*ClientConfigFileInfo) ProtoMessage()    {}
func (*ClientConfigFileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b9a411930c31ddf, []int{6}
}
func (m *ClientConfigFileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientConfigFileInfo.Unmarshal(m, b)
}
func (m *ClientConfigFileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientConfigFileInfo.Marshal(b, m, deterministic)
}
func (dst *ClientConfigFileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConfigFileInfo.Merge(dst, src)
}
func (m *ClientConfigFileInfo) XXX_Size() int {
	return xxx_messageInfo_ClientConfigFileInfo.Size(m)
}
func (m *ClientConfigFileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConfigFileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConfigFileInfo proto.InternalMessageInfo

func (m *ClientConfigFileInfo) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ClientConfigFileInfo) GetGroup() *wrapperspb.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ClientConfigFileInfo) GetFileName() *wrapperspb.StringValue {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *ClientConfigFileInfo) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ClientConfigFileInfo) GetVersion() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ClientConfigFileInfo) GetMd5() *wrapperspb.StringValue {
	if m != nil {
		return m.Md5
	}
	return nil
}

func (m *ClientConfigFileInfo) GetTags() []*ConfigFileTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ClientConfigFileInfo) GetEncrypted() *wrapperspb.BoolValue {
	if m != nil {
		return m.Encrypted
	}
	return nil
}

func (m *ClientConfigFileInfo) GetPublicKey() *wrapperspb.StringValue {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type ClientWatchConfigFileRequest struct {
	ClientIp             *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	ServiceName          *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	WatchFiles           []*ClientConfigFileInfo `protobuf:"bytes,3,rep,name=watch_files,json=watchFiles,proto3" json:"watch_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClientWatchConfigFileRequest) Reset()         { *m = ClientWatchConfigFileRequest{} }
func (m *ClientWatchConfigFileRequest) String() string { return proto.CompactTextString(m) }
func (*ClientWatchConfigFileRequest) ProtoMessage()    {}
func (*ClientWatchConfigFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b9a411930c31ddf, []int{7}
}
func (m *ClientWatchConfigFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientWatchConfigFileRequest.Unmarshal(m, b)
}
func (m *ClientWatchConfigFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientWatchConfigFileRequest.Marshal(b, m, deterministic)
}
func (dst *ClientWatchConfigFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientWatchConfigFileRequest.Merge(dst, src)
}
func (m *ClientWatchConfigFileRequest) XXX_Size() int {
	return xxx_messageInfo_ClientWatchConfigFileRequest.Size(m)
}
func (m *ClientWatchConfigFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientWatchConfigFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientWatchConfigFileRequest proto.InternalMessageInfo

func (m *ClientWatchConfigFileRequest) GetClientIp() *wrapperspb.StringValue {
	if m != nil {
		return m.ClientIp
	}
	return nil
}

func (m *ClientWatchConfigFileRequest) GetServiceName() *wrapperspb.StringValue {
	if m != nil {
		return m.ServiceName
	}
	return nil
}

func (m *ClientWatchConfigFileRequest) GetWatchFiles() []*ClientConfigFileInfo {
	if m != nil {
		return m.WatchFiles
	}
	return nil
}

type ConfigFileExportRequest struct {
	Namespace            *wrapperspb.StringValue   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Groups               []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	Names                []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ConfigFileExportRequest) Reset()         { *m = ConfigFileExportRequest{} }
func (m *ConfigFileExportRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigFileExportRequest) ProtoMessage()    {}
func (*ConfigFileExportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b9a411930c31ddf, []int{8}
}
func (m *ConfigFileExportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileExportRequest.Unmarshal(m, b)
}
func (m *ConfigFileExportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileExportRequest.Marshal(b, m, deterministic)
}
func (dst *ConfigFileExportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileExportRequest.Merge(dst, src)
}
func (m *ConfigFileExportRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigFileExportRequest.Size(m)
}
func (m *ConfigFileExportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileExportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileExportRequest proto.InternalMessageInfo

func (m *ConfigFileExportRequest) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileExportRequest) GetGroups() []*wrapperspb.StringValue {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *ConfigFileExportRequest) GetNames() []*wrapperspb.StringValue {
	if m != nil {
		return m.Names
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigFileGroup)(nil), "v1.ConfigFileGroup")
	proto.RegisterType((*ConfigFile)(nil), "v1.ConfigFile")
	proto.RegisterType((*ConfigFileTag)(nil), "v1.ConfigFileTag")
	proto.RegisterType((*ConfigFileRelease)(nil), "v1.ConfigFileRelease")
	proto.RegisterType((*ConfigFileReleaseHistory)(nil), "v1.ConfigFileReleaseHistory")
	proto.RegisterType((*ConfigFileTemplate)(nil), "v1.ConfigFileTemplate")
	proto.RegisterType((*ClientConfigFileInfo)(nil), "v1.ClientConfigFileInfo")
	proto.RegisterType((*ClientWatchConfigFileRequest)(nil), "v1.ClientWatchConfigFileRequest")
	proto.RegisterType((*ConfigFileExportRequest)(nil), "v1.ConfigFileExportRequest")
}

func init() { proto.RegisterFile("config_file.proto", fileDescriptor_config_file_7b9a411930c31ddf) }

var fileDescriptor_config_file_7b9a411930c31ddf = []byte{
	// 1001 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x58, 0xdd, 0x8a, 0xe4, 0x44,
	0x14, 0xa6, 0xff, 0xd2, 0xe9, 0xd3, 0xf3, 0x1b, 0x04, 0xc3, 0xb2, 0xc8, 0xd2, 0x20, 0x78, 0x21,
	0xc9, 0xce, 0x38, 0x0e, 0x3b, 0x2e, 0x22, 0xf6, 0xe0, 0xba, 0x83, 0x20, 0x32, 0xae, 0x0a, 0xde,
	0x34, 0xd5, 0xe9, 0xea, 0x4c, 0x61, 0x92, 0x8a, 0x55, 0x95, 0x1e, 0xf3, 0x08, 0xde, 0xec, 0x2b,
	0xf9, 0x02, 0xde, 0xfa, 0x12, 0x3e, 0x84, 0x4a, 0xfd, 0xc4, 0xf4, 0xec, 0xb8, 0x33, 0x95, 0x04,
	0x41, 0xd4, 0xab, 0x21, 0x9d, 0xef, 0x9c, 0x53, 0x39, 0xf5, 0x7d, 0xdf, 0xa9, 0x29, 0x38, 0x8c,
	0x68, 0xb6, 0x26, 0xf1, 0x62, 0x4d, 0x12, 0x1c, 0xe4, 0x8c, 0x0a, 0xea, 0xf5, 0x37, 0x47, 0x0f,
	0xde, 0x8a, 0x29, 0x8d, 0x13, 0x1c, 0xaa, 0x5f, 0x96, 0xc5, 0x3a, 0xbc, 0x66, 0x28, 0xcf, 0x31,
	0xe3, 0x1a, 0x33, 0xfb, 0x71, 0x0c, 0xfb, 0xe7, 0x2a, 0xf2, 0x19, 0x49, 0xf0, 0xa7, 0x8c, 0x16,
	0xb9, 0xf7, 0x2e, 0xf4, 0xc9, 0xca, 0xef, 0x3d, 0xea, 0xbd, 0x33, 0x3d, 0x7e, 0x18, 0xe8, 0x04,
	0x41, 0x95, 0x20, 0xf8, 0xea, 0x22, 0x13, 0xa7, 0x27, 0x5f, 0xa3, 0xa4, 0xc0, 0x97, 0x7d, 0xb2,
	0xf2, 0x1e, 0xc3, 0x30, 0x43, 0x29, 0xf6, 0xfb, 0xaf, 0xc1, 0x7f, 0x29, 0x18, 0xc9, 0x62, 0x8d,
	0x57, 0x48, 0xef, 0x03, 0x98, 0xc8, 0xbf, 0x3c, 0x47, 0x11, 0xf6, 0x07, 0x16, 0x61, 0x35, 0xdc,
	0x3b, 0x85, 0x71, 0x44, 0xd3, 0x14, 0x67, 0xc2, 0x1f, 0x5a, 0x44, 0x56, 0x60, 0xef, 0x43, 0x98,
	0x46, 0x0c, 0x23, 0x81, 0x17, 0x82, 0xa4, 0xd8, 0x1f, 0x59, 0xc4, 0x82, 0x0e, 0x78, 0x41, 0x52,
	0xec, 0x9d, 0xc1, 0xc4, 0x84, 0x2f, 0x4b, 0xdf, 0xb1, 0x08, 0x76, 0x35, 0x7c, 0x5e, 0xca, 0xca,
	0x29, 0x5d, 0x91, 0x75, 0xa9, 0x2b, 0x8f, 0x6d, 0x2a, 0xeb, 0x80, 0xaa, 0xb2, 0x09, 0x5f, 0x96,
	0xbe, 0x6b, 0x53, 0x59, 0xc3, 0xe7, 0xa5, 0xec, 0xb3, 0x64, 0xc3, 0x39, 0x2d, 0x32, 0xe1, 0x4f,
	0x2c, 0xb6, 0xb3, 0x86, 0x7b, 0x4f, 0xc0, 0x2d, 0x38, 0x66, 0x0b, 0xb2, 0xe2, 0x3e, 0x3c, 0x1a,
	0xdc, 0x5f, 0xb5, 0x42, 0xcb, 0xaa, 0xb1, 0xa4, 0x91, 0x0a, 0x9d, 0x5a, 0x84, 0xd6, 0x70, 0xef,
	0x19, 0xec, 0x33, 0x9c, 0xd2, 0x0d, 0x5e, 0xfc, 0x59, 0x7c, 0xd7, 0x22, 0xc3, 0xab, 0x41, 0xde,
	0x73, 0x38, 0x30, 0x3f, 0xd5, 0x4b, 0xd9, 0xb3, 0x48, 0x74, 0x2b, 0xca, 0x3b, 0x05, 0x17, 0xaf,
	0x88, 0x40, 0xcb, 0x04, 0xfb, 0xfb, 0xaa, 0x85, 0x0f, 0x6e, 0x65, 0x98, 0x53, 0x9a, 0x98, 0x2e,
	0x54, 0x58, 0xef, 0x18, 0x46, 0xf4, 0x3a, 0xc3, 0xcc, 0x3f, 0xb0, 0xd8, 0x32, 0x0d, 0x9d, 0xfd,
	0x3c, 0x06, 0xa8, 0xb5, 0xf8, 0x8f, 0x96, 0xe1, 0x31, 0x8c, 0x54, 0x8f, 0xac, 0x44, 0xa8, 0xa1,
	0x5a, 0xba, 0x99, 0x90, 0xd2, 0x1d, 0xd9, 0x49, 0x57, 0x81, 0xbd, 0x13, 0x70, 0xd6, 0x94, 0xa5,
	0x48, 0x58, 0x09, 0xcf, 0x60, 0xb7, 0x8d, 0x62, 0xdc, 0xc4, 0x28, 0x4e, 0xc0, 0xe1, 0x02, 0x89,
	0x82, 0x5b, 0x89, 0xcd, 0x60, 0xbd, 0xb7, 0x61, 0x28, 0x50, 0xcc, 0xfd, 0x89, 0x22, 0xd9, 0x61,
	0xb0, 0x39, 0x0a, 0xea, 0x9d, 0x7c, 0x81, 0xe2, 0x4b, 0xf5, 0xfa, 0x55, 0x17, 0x82, 0x2e, 0x2e,
	0x34, 0xed, 0xe2, 0x42, 0x3b, 0x5d, 0x5c, 0x68, 0xb7, 0x91, 0x0b, 0x7d, 0x04, 0x3b, 0x0c, 0x27,
	0x18, 0x71, 0xf3, 0xd1, 0x7b, 0x16, 0xd1, 0x53, 0x13, 0xa1, 0x6a, 0x3f, 0x05, 0xa8, 0x12, 0x2c,
	0x4b, 0x23, 0xc2, 0x7b, 0x88, 0x6a, 0xf0, 0xf3, 0xd2, 0x7b, 0x02, 0x13, 0x9c, 0x45, 0xac, 0xcc,
	0x05, 0x5e, 0x19, 0x2d, 0xde, 0x25, 0xe0, 0x1a, 0x2c, 0xd7, 0x6d, 0x1e, 0x16, 0x28, 0x89, 0xa9,
	0x7f, 0x68, 0xb3, 0x6e, 0x13, 0xf1, 0x71, 0x12, 0xd3, 0x19, 0x87, 0xdd, 0x1b, 0x1c, 0xf0, 0x02,
	0x18, 0x7c, 0x87, 0xcb, 0xd7, 0x2a, 0x7a, 0x3b, 0x91, 0x04, 0x4a, 0x91, 0x6d, 0xe4, 0x93, 0x95,
	0xa6, 0x35, 0x74, 0xf6, 0xeb, 0x08, 0x0e, 0xeb, 0xaa, 0x97, 0xba, 0x0f, 0xff, 0x3a, 0x2b, 0x39,
	0xd3, 0x93, 0x6d, 0xa1, 0x96, 0x69, 0x63, 0x26, 0xae, 0x84, 0x7f, 0x2e, 0x97, 0xba, 0xe5, 0x42,
	0x4e, 0x13, 0x17, 0x6a, 0xeb, 0x27, 0x01, 0x0c, 0xd2, 0xd5, 0xfb, 0x56, 0x66, 0x22, 0x81, 0xb2,
	0xce, 0x06, 0x33, 0x4e, 0x68, 0x66, 0x35, 0xb2, 0x2b, 0xf0, 0x7f, 0xd1, 0x5a, 0x66, 0xbf, 0x39,
	0xe0, 0xdf, 0x22, 0xfb, 0x73, 0xc2, 0x05, 0x65, 0xe5, 0xff, 0x9c, 0xef, 0xce, 0xf9, 0x7a, 0xf2,
	0x8e, 0xdb, 0x4d, 0x5e, 0xb7, 0x85, 0x52, 0x26, 0xb6, 0x4a, 0x79, 0x0c, 0x43, 0x51, 0xe6, 0x76,
	0x54, 0x57, 0xc8, 0xad, 0xd9, 0x3e, 0x6d, 0x31, 0xdb, 0x77, 0x1a, 0xcd, 0xf6, 0xdd, 0x2e, 0x02,
	0xdc, 0xeb, 0x22, 0xc0, 0xfd, 0x2e, 0x02, 0x3c, 0x68, 0x24, 0xc0, 0x97, 0x43, 0xf0, 0xb6, 0x7a,
	0x81, 0xd3, 0x3c, 0x41, 0xe2, 0xef, 0x1f, 0x37, 0x5b, 0x7c, 0x1e, 0xb4, 0xe3, 0xf3, 0xb0, 0x1d,
	0x9f, 0x47, 0x1d, 0xfe, 0xe5, 0x74, 0xba, 0x10, 0x62, 0xdc, 0x85, 0x10, 0x6e, 0x17, 0x42, 0x4c,
	0x1a, 0x11, 0xe2, 0xf7, 0x01, 0xbc, 0x71, 0x9e, 0x10, 0x9c, 0x89, 0x9a, 0x16, 0x17, 0xd9, 0x9a,
	0xde, 0x74, 0xcb, 0x5e, 0x4b, 0xb7, 0xec, 0xb7, 0x74, 0xcb, 0x41, 0x5b, 0xb7, 0x1c, 0x36, 0x3c,
	0x21, 0x54, 0x93, 0x7b, 0xd4, 0x64, 0x72, 0x1b, 0xdf, 0x73, 0x6c, 0x7d, 0xaf, 0xf2, 0xa3, 0xf1,
	0xdd, 0x7e, 0x74, 0xe3, 0xe4, 0xeb, 0x36, 0x39, 0xf9, 0x3e, 0x05, 0xc8, 0x8b, 0x65, 0x42, 0xa2,
	0x85, 0x3c, 0xae, 0xda, 0x10, 0x60, 0xa2, 0xf1, 0x9f, 0xe1, 0x72, 0xf6, 0x4b, 0x0f, 0x1e, 0x6a,
	0x06, 0x7c, 0x83, 0x44, 0x74, 0xb5, 0x3d, 0x9e, 0xbf, 0x2f, 0x30, 0x17, 0x8a, 0xd7, 0xea, 0xfd,
	0x82, 0xe4, 0x56, 0x4c, 0x70, 0x35, 0xfc, 0x22, 0x97, 0x47, 0x72, 0x8e, 0xd9, 0x86, 0x44, 0x66,
	0x5f, 0x6d, 0xf8, 0x30, 0x35, 0x11, 0x6a, 0x6b, 0xcf, 0x60, 0x7a, 0x2d, 0x57, 0xa5, 0x6e, 0xc9,
	0xb8, 0x3f, 0x50, 0x1d, 0xf4, 0x55, 0x07, 0xff, 0x82, 0xb4, 0x97, 0xa0, 0xc0, 0xf2, 0x91, 0xcf,
	0x7e, 0xea, 0xc1, 0x9b, 0xf5, 0xeb, 0x4f, 0x7e, 0xc8, 0x29, 0x13, 0xd5, 0x27, 0x75, 0x21, 0xf7,
	0x09, 0x38, 0x8a, 0xb1, 0xdc, 0xef, 0x5b, 0x5c, 0x50, 0x18, 0xac, 0x94, 0x84, 0x4a, 0x61, 0x3e,
	0xe1, 0x1e, 0x49, 0x28, 0xe8, 0xfc, 0x65, 0x0f, 0x4e, 0x23, 0x9a, 0x06, 0x02, 0x67, 0x11, 0xce,
	0x44, 0x90, 0xd3, 0x04, 0x31, 0xc2, 0x03, 0x9e, 0xe3, 0x88, 0xac, 0x49, 0x84, 0x04, 0xa1, 0x59,
	0x80, 0x72, 0x22, 0xfb, 0xa1, 0x6f, 0x13, 0x83, 0x14, 0x65, 0x28, 0xc6, 0xf3, 0xad, 0x2b, 0xc2,
	0x2f, 0x64, 0x85, 0x6f, 0xcf, 0x63, 0x22, 0xae, 0x8a, 0x65, 0x10, 0xd1, 0x34, 0x34, 0x79, 0x52,
	0xcc, 0xaf, 0xc2, 0x1b, 0xb9, 0x42, 0x4e, 0x0b, 0x16, 0xe1, 0x30, 0xa6, 0x21, 0xca, 0x49, 0xb8,
	0x39, 0x0a, 0xcd, 0x1d, 0xa5, 0xce, 0xba, 0x74, 0xd4, 0x6a, 0xdf, 0xfb, 0x23, 0x00, 0x00, 0xff,
	0xff, 0x6e, 0xe1, 0x74, 0xde, 0xbb, 0x14, 0x00, 0x00,
}
