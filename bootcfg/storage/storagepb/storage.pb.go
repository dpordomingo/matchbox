// Code generated by protoc-gen-go.
// source: storage.proto
// DO NOT EDIT!

/*
Package storagepb is a generated protocol buffer package.

It is generated from these files:
	storage.proto

It has these top-level messages:
	Group
	Profile
	NetBoot
*/
package storagepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Group struct {
	// machine readable Id
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// human readable name
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// profile id
	Profile string `protobuf:"bytes,3,opt,name=profile" json:"profile,omitempty"`
	// tags required to match the group
	Requirements map[string]string `protobuf:"bytes,4,rep,name=requirements" json:"requirements,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// custom metadata
	Metadata map[string]string `protobuf:"bytes,5,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Group) Reset()                    { *m = Group{} }
func (m *Group) String() string            { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()               {}
func (*Group) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Group) GetRequirements() map[string]string {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *Group) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Profile struct {
	// profile id
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// human readable name
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// ignition id
	IgnitionId string `protobuf:"bytes,3,opt,name=ignition_id,json=ignitionId" json:"ignition_id,omitempty"`
	// cloud config id
	CloudId string `protobuf:"bytes,4,opt,name=cloud_id,json=cloudId" json:"cloud_id,omitempty"`
	// support network boot / PXE
	Boot *NetBoot `protobuf:"bytes,5,opt,name=boot" json:"boot,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Profile) GetBoot() *NetBoot {
	if m != nil {
		return m.Boot
	}
	return nil
}

type NetBoot struct {
	// the URL of the kernel image
	Kernel string `protobuf:"bytes,1,opt,name=kernel" json:"kernel,omitempty"`
	// the init RAM filesystem URLs
	Initrd []string `protobuf:"bytes,2,rep,name=initrd" json:"initrd,omitempty"`
	// kernel parameters (command line)
	Cmdline map[string]string `protobuf:"bytes,3,rep,name=cmdline" json:"cmdline,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NetBoot) Reset()                    { *m = NetBoot{} }
func (m *NetBoot) String() string            { return proto.CompactTextString(m) }
func (*NetBoot) ProtoMessage()               {}
func (*NetBoot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NetBoot) GetCmdline() map[string]string {
	if m != nil {
		return m.Cmdline
	}
	return nil
}

func init() {
	proto.RegisterType((*Group)(nil), "storagepb.Group")
	proto.RegisterType((*Profile)(nil), "storagepb.Profile")
	proto.RegisterType((*NetBoot)(nil), "storagepb.NetBoot")
}

var fileDescriptor0 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x4b, 0x4e, 0xc3, 0x30,
	0x10, 0x55, 0x3e, 0xfd, 0x4d, 0x5b, 0x04, 0x23, 0x84, 0x42, 0x17, 0xb4, 0xca, 0x02, 0x75, 0x95,
	0x05, 0x6c, 0xa0, 0x2c, 0x90, 0x40, 0x80, 0x58, 0x80, 0x50, 0x2e, 0x80, 0xd2, 0xc6, 0x54, 0x56,
	0x13, 0xbb, 0xb8, 0x0e, 0x52, 0x8f, 0xc1, 0x4d, 0xb8, 0x05, 0xd7, 0xc2, 0xb1, 0x9d, 0xaa, 0x55,
	0x36, 0x74, 0x37, 0xef, 0xcd, 0x9b, 0x8f, 0xdf, 0x18, 0xfa, 0x2b, 0xc9, 0x45, 0x32, 0x27, 0xd1,
	0x52, 0x70, 0xc9, 0xb1, 0x63, 0xe1, 0x72, 0x1a, 0xfe, 0xba, 0xd0, 0x78, 0x12, 0xbc, 0x58, 0xe2,
	0x01, 0xb8, 0x34, 0x0d, 0x9c, 0x91, 0x33, 0xee, 0xc4, 0x2a, 0x42, 0x04, 0x9f, 0x25, 0x39, 0x09,
	0x5c, 0xcd, 0xe8, 0x18, 0x03, 0x68, 0xa9, 0x0e, 0x1f, 0x34, 0x23, 0x81, 0xa7, 0xe9, 0x0a, 0xe2,
	0x23, 0xf4, 0x04, 0xf9, 0x2c, 0xa8, 0x20, 0x39, 0x61, 0x72, 0x15, 0xf8, 0x23, 0x6f, 0xdc, 0xbd,
	0x08, 0xa3, 0xcd, 0xa4, 0x48, 0x4f, 0x89, 0xe2, 0x2d, 0xd1, 0x03, 0x93, 0x62, 0x1d, 0xef, 0xd4,
	0xe1, 0x04, 0xda, 0x39, 0x91, 0x49, 0x9a, 0xc8, 0x24, 0x68, 0xe8, 0x1e, 0x67, 0xb5, 0x1e, 0x2f,
	0x56, 0x60, 0xea, 0x37, 0xfa, 0xc1, 0x2d, 0x1c, 0xd5, 0xda, 0xe3, 0x21, 0x78, 0x0b, 0xb2, 0xb6,
	0xef, 0x2a, 0x43, 0x3c, 0x86, 0xc6, 0x57, 0x92, 0x15, 0xd5, 0xcb, 0x0c, 0x98, 0xb8, 0x57, 0xce,
	0xe0, 0x06, 0xfa, 0x3b, 0xbd, 0xf7, 0x29, 0x0e, 0xbf, 0x1d, 0x68, 0xbd, 0x59, 0x37, 0xfe, 0xe3,
	0xe5, 0x10, 0xba, 0x74, 0xce, 0xa8, 0xa4, 0x9c, 0xbd, 0x2b, 0xb1, 0xf1, 0x13, 0x2a, 0xea, 0x39,
	0xc5, 0x53, 0x68, 0xcf, 0x32, 0x5e, 0xa4, 0x65, 0xd6, 0x37, 0x6e, 0x6b, 0xac, 0x52, 0xe7, 0xe0,
	0x4f, 0x39, 0x97, 0xca, 0x21, 0x47, 0x39, 0x84, 0x5b, 0x0e, 0xbd, 0x12, 0x79, 0xa7, 0x32, 0xb1,
	0xce, 0x87, 0x3f, 0x6a, 0x27, 0xcb, 0xe0, 0x09, 0x34, 0x17, 0x44, 0x30, 0x92, 0xd9, 0xbd, 0x2c,
	0x2a, 0x79, 0xaa, 0x66, 0x8a, 0x54, 0x6d, 0xe7, 0x95, 0xbc, 0x41, 0x78, 0x0d, 0xad, 0x59, 0x9e,
	0x66, 0x94, 0x95, 0xb7, 0x2e, 0x0f, 0x31, 0xac, 0x8f, 0x89, 0xee, 0x8d, 0xc2, 0x5c, 0xa2, 0xd2,
	0x0f, 0x26, 0xd0, 0xdb, 0x4e, 0xec, 0x63, 0xe3, 0xb4, 0xa9, 0xbf, 0xe8, 0xe5, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x97, 0x23, 0xfd, 0x6e, 0xb3, 0x02, 0x00, 0x00,
}
