// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accesssecurity.proto

// Package access_security provides basic definition of user and permission group

package access_security

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PasswordHash         string   `protobuf:"bytes,2,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	Permissions          []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bb83253976cb388, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPasswordHash() string {
	if m != nil {
		return m.PasswordHash
	}
	return ""
}

func (m *User) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type PermissionGroup struct {
	Name                 string                         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Permissions          []*PermissionGroup_Permissions `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *PermissionGroup) Reset()         { *m = PermissionGroup{} }
func (m *PermissionGroup) String() string { return proto.CompactTextString(m) }
func (*PermissionGroup) ProtoMessage()    {}
func (*PermissionGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bb83253976cb388, []int{1}
}

func (m *PermissionGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionGroup.Unmarshal(m, b)
}
func (m *PermissionGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionGroup.Marshal(b, m, deterministic)
}
func (m *PermissionGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionGroup.Merge(m, src)
}
func (m *PermissionGroup) XXX_Size() int {
	return xxx_messageInfo_PermissionGroup.Size(m)
}
func (m *PermissionGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionGroup.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionGroup proto.InternalMessageInfo

func (m *PermissionGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PermissionGroup) GetPermissions() []*PermissionGroup_Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type PermissionGroup_Permissions struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	AllowedMethods       []string `protobuf:"bytes,2,rep,name=allowed_methods,json=allowedMethods,proto3" json:"allowed_methods,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PermissionGroup_Permissions) Reset()         { *m = PermissionGroup_Permissions{} }
func (m *PermissionGroup_Permissions) String() string { return proto.CompactTextString(m) }
func (*PermissionGroup_Permissions) ProtoMessage()    {}
func (*PermissionGroup_Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bb83253976cb388, []int{1, 0}
}

func (m *PermissionGroup_Permissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionGroup_Permissions.Unmarshal(m, b)
}
func (m *PermissionGroup_Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionGroup_Permissions.Marshal(b, m, deterministic)
}
func (m *PermissionGroup_Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionGroup_Permissions.Merge(m, src)
}
func (m *PermissionGroup_Permissions) XXX_Size() int {
	return xxx_messageInfo_PermissionGroup_Permissions.Size(m)
}
func (m *PermissionGroup_Permissions) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionGroup_Permissions.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionGroup_Permissions proto.InternalMessageInfo

func (m *PermissionGroup_Permissions) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PermissionGroup_Permissions) GetAllowedMethods() []string {
	if m != nil {
		return m.AllowedMethods
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "access_security.User")
	proto.RegisterType((*PermissionGroup)(nil), "access_security.PermissionGroup")
	proto.RegisterType((*PermissionGroup_Permissions)(nil), "access_security.PermissionGroup.Permissions")
}

func init() { proto.RegisterFile("accesssecurity.proto", fileDescriptor_6bb83253976cb388) }

var fileDescriptor_6bb83253976cb388 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x6a, 0x83, 0x40,
	0x14, 0x85, 0xd1, 0x91, 0x82, 0xd7, 0xb6, 0x96, 0xa1, 0x0b, 0xe9, 0x6a, 0xb0, 0x8b, 0xba, 0x28,
	0x2e, 0xda, 0x87, 0x88, 0x9b, 0x84, 0x20, 0x64, 0x2d, 0x13, 0x1d, 0x18, 0x41, 0x9d, 0x61, 0xae,
	0x22, 0x79, 0xba, 0xbc, 0x5a, 0x88, 0x3f, 0x41, 0x25, 0xbb, 0x7b, 0xbe, 0x03, 0xe7, 0x1b, 0x06,
	0x3e, 0x79, 0x9e, 0x0b, 0x44, 0x14, 0x79, 0x67, 0xca, 0xf6, 0x12, 0x6b, 0xa3, 0x5a, 0x45, 0xfd,
	0x91, 0x66, 0x33, 0x0e, 0x39, 0x38, 0x27, 0x14, 0x86, 0x52, 0x70, 0x1a, 0x5e, 0x8b, 0xc0, 0x62,
	0x56, 0xe4, 0xa6, 0xc3, 0x4d, 0xbf, 0xe1, 0x4d, 0x73, 0xc4, 0x5e, 0x99, 0x22, 0x93, 0x1c, 0x65,
	0x60, 0x0f, 0xe5, 0xeb, 0x0c, 0x13, 0x8e, 0x92, 0x32, 0xf0, 0xb4, 0x30, 0x75, 0x89, 0x58, 0xaa,
	0x06, 0x03, 0xc2, 0x48, 0xe4, 0xa6, 0x4b, 0x14, 0x5e, 0x2d, 0xf0, 0x8f, 0x8f, 0xbc, 0x33, 0xaa,
	0xd3, 0x4f, 0x75, 0x87, 0xf5, 0x92, 0xcd, 0x48, 0xe4, 0xfd, 0xfd, 0xc6, 0x9b, 0x17, 0xc7, 0x9b,
	0xa9, 0x45, 0xc6, 0x95, 0xf7, 0x2b, 0x01, 0x6f, 0xd1, 0xd1, 0x0f, 0x20, 0x9d, 0xa9, 0x26, 0xe3,
	0xfd, 0xa4, 0x3f, 0xe0, 0xf3, 0xaa, 0x52, 0xbd, 0x28, 0xb2, 0x5a, 0xb4, 0x52, 0x15, 0xa3, 0xd4,
	0x4d, 0xdf, 0x27, 0xbc, 0x1f, 0xe9, 0xf9, 0x65, 0xf8, 0xbc, 0xff, 0x5b, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8d, 0x8c, 0xed, 0x5a, 0x54, 0x01, 0x00, 0x00,
}
