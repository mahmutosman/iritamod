// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: perm/perm.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
	strconv "strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Role represents a role
type Role int32

const (
	// ROOT_ADMIN defines the root admin role index.
	RoleRootAdmin Role = 0
	// PERM_ADMIN defines the permission admin role index.
	RolePermAdmin Role = 1
	// BLACKLIST_ADMIN defines the blacklist admin role index.
	RoleBlacklistAdmin Role = 2
	// NODE_ADMIN defines the node admin role index.
	RoleNodeAdmin Role = 3
	// PARAM_ADMIN defines the param admin role index.
	RoleParamAdmin Role = 4
	// POWER_USER defines the power user role index.
	RolePowerUser Role = 5
	// RELAYER_USER defines the relayer user role index.
	RoleRelayerUser Role = 6
	// ID_ADMIN defines the identity admin role index.
	RoleIDAdmin Role = 7
	// BASE_M1_ADMIN defines the base M1 admin role index.
	RoleBaseM1Admin Role = 8
	// Chain_Account_Role defines the platform admin role index.
	RolePlatformUser Role = 9
)

var Role_name = map[int32]string{
	0: "ROOT_ADMIN",
	1: "PERM_ADMIN",
	2: "BLACKLIST_ADMIN",
	3: "NODE_ADMIN",
	4: "PARAM_ADMIN",
	5: "POWER_USER",
	6: "RELAYER_USER",
	7: "ID_ADMIN",
	8: "BASE_M1_ADMIN",
	9: "PLATFORM_USER",
}

var Role_value = map[string]int32{
	"ROOT_ADMIN":      0,
	"PERM_ADMIN":      1,
	"BLACKLIST_ADMIN": 2,
	"NODE_ADMIN":      3,
	"PARAM_ADMIN":     4,
	"POWER_USER":      5,
	"RELAYER_USER":    6,
	"ID_ADMIN":        7,
	"BASE_M1_ADMIN":   8,
	"PLATFORM_USER":   9,
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb77ba30a3a45e51, []int{0}
}

func init() {
	proto.RegisterEnum("iritamod.perm.Role", Role_name, Role_value)
	golang_proto.RegisterEnum("iritamod.perm.Role", Role_name, Role_value)
}

func init() { proto.RegisterFile("perm/perm.proto", fileDescriptor_bb77ba30a3a45e51) }
func init() { golang_proto.RegisterFile("perm/perm.proto", fileDescriptor_bb77ba30a3a45e51) }

var fileDescriptor_bb77ba30a3a45e51 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd2, 0xcf, 0x6a, 0xd4, 0x40,
	0x1c, 0x07, 0xf0, 0xc4, 0xd6, 0x5a, 0xa7, 0x5d, 0x77, 0x8d, 0xc5, 0xc3, 0x80, 0x43, 0x44, 0x54,
	0x50, 0xd8, 0x58, 0x7c, 0x82, 0x89, 0x89, 0x10, 0x4c, 0x36, 0x61, 0xb6, 0x45, 0xf4, 0xb2, 0xcc,
	0x36, 0xe3, 0x3a, 0x9a, 0xe9, 0x2c, 0x93, 0x14, 0xe9, 0x1b, 0x68, 0x4e, 0xbe, 0x40, 0x4e, 0xf6,
	0xe0, 0xd1, 0x47, 0xf0, 0xd8, 0x63, 0x8f, 0x1e, 0x75, 0xf3, 0x22, 0x32, 0xf9, 0x43, 0x7a, 0x09,
	0x21, 0xf3, 0xe1, 0xfb, 0xfd, 0x85, 0xdf, 0x80, 0xf1, 0x9a, 0x29, 0xe1, 0xe8, 0xc7, 0x74, 0xad,
	0x64, 0x21, 0xad, 0x11, 0x57, 0xbc, 0xa0, 0x42, 0xa6, 0x53, 0xfd, 0x11, 0x1e, 0xac, 0xe4, 0x4a,
	0x36, 0x27, 0x8e, 0x7e, 0x6b, 0xd1, 0xb3, 0x6f, 0x5b, 0x60, 0x9b, 0xc8, 0x8c, 0x59, 0x0f, 0x01,
	0x20, 0x71, 0x7c, 0xb4, 0xc0, 0x5e, 0x14, 0xcc, 0x26, 0x06, 0xbc, 0x5b, 0x56, 0xf6, 0x48, 0x9f,
	0x10, 0x29, 0x0b, 0x9c, 0x0a, 0x7e, 0xaa, 0x49, 0xe2, 0x93, 0xa8, 0x23, 0xe6, 0x40, 0x12, 0xa6,
	0x44, 0x4b, 0x9e, 0x83, 0xb1, 0x1b, 0xe2, 0x57, 0x6f, 0xc2, 0x60, 0xde, 0x47, 0xdd, 0x80, 0xf7,
	0xcb, 0xca, 0xb6, 0xb4, 0x73, 0x33, 0x7a, 0xf2, 0x39, 0xe3, 0xf9, 0x90, 0x37, 0x8b, 0x3d, 0xbf,
	0x73, 0x5b, 0x43, 0xde, 0x4c, 0xa6, 0xac, 0x25, 0x8f, 0xc0, 0x5e, 0x82, 0x09, 0xee, 0x3b, 0xb7,
	0xa1, 0x55, 0x56, 0xf6, 0x9d, 0xa6, 0x93, 0x2a, 0x2a, 0x86, 0xb9, 0xe2, 0xb7, 0x3e, 0x59, 0x1c,
	0xcf, 0x7d, 0x32, 0xb9, 0x79, 0x6d, 0x2e, 0xf9, 0x85, 0xa9, 0xe3, 0x9c, 0x29, 0xeb, 0x31, 0xd8,
	0x27, 0x7e, 0x88, 0xdf, 0xf5, 0x68, 0x07, 0xde, 0x2b, 0x2b, 0x7b, 0xdc, 0xfc, 0x1f, 0xcb, 0xe8,
	0x79, 0xc7, 0x1e, 0x80, 0xdd, 0xc0, 0xeb, 0xba, 0x6e, 0xc1, 0x71, 0x59, 0xd9, 0x7b, 0x9a, 0x04,
	0x5e, 0x5b, 0xf4, 0x04, 0x8c, 0x5c, 0x3c, 0xf7, 0x17, 0xd1, 0x61, 0x67, 0x76, 0x87, 0x18, 0x97,
	0xe6, 0x2c, 0x3a, 0x6c, 0xdd, 0x53, 0x30, 0x4a, 0x42, 0x7c, 0xf4, 0x3a, 0x26, 0x51, 0x5b, 0x77,
	0x1b, 0x1e, 0x94, 0x95, 0x3d, 0x69, 0x66, 0xca, 0x68, 0xf1, 0x41, 0x2a, 0xa1, 0xfb, 0xe0, 0xfe,
	0xd7, 0x1f, 0xc8, 0xf8, 0x79, 0x81, 0x8c, 0x5f, 0x17, 0xc8, 0x74, 0xc9, 0xe5, 0x3f, 0x64, 0x5c,
	0x6e, 0x90, 0x79, 0xb5, 0x41, 0xe6, 0xdf, 0x0d, 0x32, 0xbf, 0xd7, 0xc8, 0xf8, 0x5d, 0x23, 0xf3,
	0xaa, 0x46, 0xc6, 0x9f, 0x1a, 0x19, 0xef, 0x5f, 0xac, 0x78, 0xf1, 0xf1, 0x6c, 0x39, 0x3d, 0x91,
	0xc2, 0x59, 0x72, 0x7a, 0xfa, 0x89, 0x33, 0xca, 0x9d, 0x7e, 0xcf, 0x8e, 0x90, 0xe9, 0x59, 0xc6,
	0xf2, 0xe6, 0x12, 0x38, 0xc5, 0xf9, 0x9a, 0xe5, 0xcb, 0x9d, 0x66, 0xcd, 0x2f, 0xff, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x2e, 0xe4, 0xe2, 0x0f, 0x1e, 0x02, 0x00, 0x00,
}

func (x Role) String() string {
	s, ok := Role_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
