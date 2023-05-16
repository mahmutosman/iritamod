// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer2/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GenesisState struct {
	StartingSpaceId   uint64             `protobuf:"varint,1,opt,name=starting_space_id,json=startingSpaceId,proto3" json:"starting_space_id,omitempty"`
	Spaces            []Space            `protobuf:"bytes,2,rep,name=spaces,proto3" json:"spaces"`
	L2BlockHeaders    []L2BlockHeader    `protobuf:"bytes,3,rep,name=l2_block_headers,json=l2BlockHeaders,proto3" json:"l2_block_headers"`
	ClassesForNft     []ClassForNFT      `protobuf:"bytes,4,rep,name=classes_for_nft,json=classesForNft,proto3" json:"classes_for_nft"`
	CollectionsForNft []CollectionForNFT `protobuf:"bytes,5,rep,name=collections_for_nft,json=collectionsForNft,proto3" json:"collections_for_nft"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_841d40cba3971155, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetStartingSpaceId() uint64 {
	if m != nil {
		return m.StartingSpaceId
	}
	return 0
}

func (m *GenesisState) GetSpaces() []Space {
	if m != nil {
		return m.Spaces
	}
	return nil
}

func (m *GenesisState) GetL2BlockHeaders() []L2BlockHeader {
	if m != nil {
		return m.L2BlockHeaders
	}
	return nil
}

func (m *GenesisState) GetClassesForNft() []ClassForNFT {
	if m != nil {
		return m.ClassesForNft
	}
	return nil
}

func (m *GenesisState) GetCollectionsForNft() []CollectionForNFT {
	if m != nil {
		return m.CollectionsForNft
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "iritamod.layer2.v1.GenesisState")
}

func init() { proto.RegisterFile("layer2/v1/genesis.proto", fileDescriptor_841d40cba3971155) }

var fileDescriptor_841d40cba3971155 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd1, 0xcf, 0x4e, 0xc2, 0x30,
	0x00, 0xc7, 0xf1, 0x0d, 0x90, 0x43, 0xfd, 0x83, 0x54, 0xa3, 0x93, 0xc3, 0x40, 0xe3, 0x81, 0x78,
	0x58, 0xc3, 0x3c, 0x78, 0xc7, 0x04, 0x35, 0x41, 0x13, 0xc1, 0x13, 0x97, 0xa5, 0xdb, 0xca, 0xa8,
	0x96, 0x95, 0xb4, 0x85, 0x84, 0xb7, 0xf0, 0xb1, 0x38, 0x72, 0xf4, 0x64, 0x0c, 0x3c, 0x81, 0x6f,
	0x60, 0xd6, 0x0d, 0xc1, 0xc8, 0x6d, 0xf9, 0xf5, 0xbb, 0xcf, 0x96, 0x14, 0x9c, 0x32, 0x3c, 0x25,
	0xc2, 0x45, 0x93, 0x06, 0x8a, 0x48, 0x4c, 0x24, 0x95, 0xce, 0x48, 0x70, 0xc5, 0x21, 0xa4, 0x82,
	0x2a, 0x3c, 0xe4, 0xa1, 0x93, 0x16, 0xce, 0xa4, 0x51, 0x39, 0x8e, 0x78, 0xc4, 0xf5, 0x31, 0x4a,
	0x9e, 0xd2, 0xb2, 0x72, 0xb2, 0x26, 0xb2, 0x54, 0xef, 0x17, 0xdf, 0x39, 0xb0, 0x77, 0x97, 0x9a,
	0x5d, 0x85, 0x15, 0x81, 0x57, 0xa0, 0x2c, 0x15, 0x16, 0x8a, 0xc6, 0x91, 0x27, 0x47, 0x38, 0x20,
	0x1e, 0x0d, 0x2d, 0xb3, 0x66, 0xd6, 0x0b, 0x9d, 0xd2, 0xea, 0xa0, 0x9b, 0xec, 0x0f, 0x21, 0xbc,
	0x01, 0x45, 0x9d, 0x48, 0x2b, 0x57, 0xcb, 0xd7, 0x77, 0xdd, 0x33, 0xe7, 0xff, 0xff, 0x38, 0x3a,
	0x6e, 0x16, 0x66, 0x9f, 0x55, 0xa3, 0x93, 0xe5, 0xf0, 0x19, 0x1c, 0x32, 0xd7, 0xf3, 0x19, 0x0f,
	0xde, 0xbc, 0x01, 0xc1, 0x21, 0x11, 0xd2, 0xca, 0x6b, 0xe2, 0x7c, 0x1b, 0xd1, 0x76, 0x9b, 0x49,
	0x7a, 0xaf, 0xcb, 0x8c, 0x3a, 0x60, 0x9b, 0xa3, 0x84, 0x8f, 0xa0, 0x14, 0x30, 0x2c, 0x25, 0x91,
	0x5e, 0x9f, 0x0b, 0x2f, 0xee, 0x2b, 0xab, 0xa0, 0xc5, 0xea, 0x36, 0xf1, 0x36, 0x49, 0x5b, 0x5c,
	0x3c, 0xb5, 0x5e, 0x32, 0x6f, 0x3f, 0x7b, 0x3b, 0x19, 0xfb, 0x0a, 0xf6, 0xc0, 0x51, 0xc0, 0x19,
	0x23, 0x81, 0xa2, 0x3c, 0x5e, 0x93, 0x3b, 0x9a, 0xbc, 0xdc, 0x4a, 0xfe, 0xe6, 0x7f, 0xdc, 0xf2,
	0x06, 0x93, 0xda, 0xcd, 0xf6, 0x6c, 0x61, 0x9b, 0xf3, 0x85, 0x6d, 0x7e, 0x2d, 0x6c, 0xf3, 0x7d,
	0x69, 0x1b, 0xf3, 0xa5, 0x6d, 0x7c, 0x2c, 0x6d, 0xa3, 0xe7, 0x46, 0x54, 0x0d, 0xc6, 0xbe, 0x13,
	0xf0, 0x21, 0xf2, 0x29, 0x8e, 0x5f, 0x29, 0xc1, 0x14, 0xad, 0x3e, 0x86, 0x86, 0x3c, 0x1c, 0x33,
	0x22, 0xb3, 0x1b, 0x44, 0x6a, 0x3a, 0x22, 0xd2, 0x2f, 0xea, 0x8b, 0xbc, 0xfe, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x3d, 0x11, 0x4c, 0xfa, 0x25, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CollectionsForNft) > 0 {
		for iNdEx := len(m.CollectionsForNft) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CollectionsForNft[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ClassesForNft) > 0 {
		for iNdEx := len(m.ClassesForNft) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClassesForNft[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.L2BlockHeaders) > 0 {
		for iNdEx := len(m.L2BlockHeaders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.L2BlockHeaders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Spaces) > 0 {
		for iNdEx := len(m.Spaces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Spaces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.StartingSpaceId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartingSpaceId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartingSpaceId != 0 {
		n += 1 + sovGenesis(uint64(m.StartingSpaceId))
	}
	if len(m.Spaces) > 0 {
		for _, e := range m.Spaces {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.L2BlockHeaders) > 0 {
		for _, e := range m.L2BlockHeaders {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClassesForNft) > 0 {
		for _, e := range m.ClassesForNft {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CollectionsForNft) > 0 {
		for _, e := range m.CollectionsForNft {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingSpaceId", wireType)
			}
			m.StartingSpaceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingSpaceId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spaces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spaces = append(m.Spaces, Space{})
			if err := m.Spaces[len(m.Spaces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field L2BlockHeaders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.L2BlockHeaders = append(m.L2BlockHeaders, L2BlockHeader{})
			if err := m.L2BlockHeaders[len(m.L2BlockHeaders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassesForNft", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassesForNft = append(m.ClassesForNft, ClassForNFT{})
			if err := m.ClassesForNft[len(m.ClassesForNft)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionsForNft", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectionsForNft = append(m.CollectionsForNft, CollectionForNFT{})
			if err := m.CollectionsForNft[len(m.CollectionsForNft)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
