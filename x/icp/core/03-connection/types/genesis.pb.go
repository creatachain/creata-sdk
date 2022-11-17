// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: icp/core/connection/v1/genesis.proto

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

// GenesisState defines the icp connection submodule's genesis state.
type GenesisState struct {
	Connections           []IdentifiedConnection `protobuf:"bytes,1,rep,name=connections,proto3" json:"connections"`
	ClientConnectionPaths []ConnectionPaths      `protobuf:"bytes,2,rep,name=client_connection_paths,json=clientConnectionPaths,proto3" json:"client_connection_paths" yaml:"client_connection_paths"`
	// the sequence for the next generated connection identifier
	NextConnectionSequence uint64 `protobuf:"varint,3,opt,name=next_connection_sequence,json=nextConnectionSequence,proto3" json:"next_connection_sequence,omitempty" yaml:"next_connection_sequence"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9387f3125d9c6025, []int{0}
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

func (m *GenesisState) GetConnections() []IdentifiedConnection {
	if m != nil {
		return m.Connections
	}
	return nil
}

func (m *GenesisState) GetClientConnectionPaths() []ConnectionPaths {
	if m != nil {
		return m.ClientConnectionPaths
	}
	return nil
}

func (m *GenesisState) GetNextConnectionSequence() uint64 {
	if m != nil {
		return m.NextConnectionSequence
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "icp.core.connection.v1.GenesisState")
}

func init() {
	proto.RegisterFile("icp/core/connection/v1/genesis.proto", fileDescriptor_9387f3125d9c6025)
}

var fileDescriptor_9387f3125d9c6025 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x31, 0x4f, 0xf2, 0x40,
	0x18, 0xc7, 0x5b, 0x20, 0xef, 0x50, 0xde, 0xa9, 0x51, 0x6c, 0x18, 0xae, 0xa4, 0x1a, 0x61, 0x90,
	0x3b, 0x91, 0xcd, 0xc9, 0xd4, 0xc1, 0xb8, 0x19, 0x70, 0x22, 0x31, 0xe4, 0x38, 0x1e, 0xcb, 0x45,
	0xb8, 0xab, 0xdc, 0x41, 0xe0, 0x13, 0xb8, 0xfa, 0xb1, 0x58, 0x4c, 0x18, 0x9d, 0x88, 0x81, 0x6f,
	0xc0, 0x27, 0x30, 0x6d, 0x09, 0x45, 0x63, 0xb7, 0x27, 0xcf, 0xf3, 0x7b, 0x7e, 0xff, 0xe1, 0x6f,
	0x9d, 0x71, 0x16, 0x12, 0x26, 0xc7, 0x40, 0x98, 0x14, 0x02, 0x98, 0xe6, 0x52, 0x90, 0x69, 0x83,
	0x04, 0x20, 0x40, 0x71, 0x85, 0xc3, 0xb1, 0xd4, 0xd2, 0x2e, 0x71, 0x16, 0xe2, 0x88, 0xc2, 0x29,
	0x85, 0xa7, 0x8d, 0xf2, 0x51, 0x20, 0x03, 0x19, 0x23, 0x24, 0x9a, 0x12, 0xba, 0x5c, 0xcd, 0x70,
	0x1e, 0xfc, 0xc6, 0xa0, 0xf7, 0x91, 0xb3, 0xfe, 0xdf, 0x25, 0x41, 0x6d, 0x4d, 0x35, 0xd8, 0x8f,
	0x56, 0x31, 0x85, 0x94, 0x63, 0x56, 0xf2, 0xb5, 0xe2, 0xd5, 0x05, 0xfe, 0x3b, 0x1d, 0xdf, 0xf7,
	0x41, 0x68, 0xfe, 0xcc, 0xa1, 0x7f, 0xbb, 0xdf, 0xfb, 0x85, 0xc5, 0xca, 0x35, 0x5a, 0x87, 0x1a,
	0xfb, 0xcd, 0xb4, 0x4e, 0xd8, 0x90, 0x83, 0xd0, 0xdd, 0x74, 0xdd, 0x0d, 0xa9, 0x1e, 0x28, 0x27,
	0x17, 0x47, 0x54, 0xb3, 0x22, 0x52, 0xf1, 0x43, 0x84, 0xfb, 0xe7, 0x91, 0x7d, 0xbb, 0x72, 0xd1,
	0x9c, 0x8e, 0x86, 0xd7, 0x5e, 0x86, 0xd5, 0x6b, 0x1d, 0x27, 0x97, 0x5f, 0xef, 0xf6, 0x93, 0xe5,
	0x08, 0x98, 0xfd, 0x78, 0x50, 0xf0, 0x3a, 0x01, 0xc1, 0xc0, 0xc9, 0x57, 0xcc, 0x5a, 0xc1, 0x3f,
	0xdd, 0xae, 0x5c, 0x37, 0x91, 0x67, 0x91, 0x5e, 0xab, 0x14, 0x9d, 0x52, 0x77, 0x7b, 0x77, 0xf0,
	0x3b, 0x8b, 0x35, 0x32, 0x97, 0x6b, 0x64, 0x7e, 0xad, 0x91, 0xf9, 0xbe, 0x41, 0xc6, 0x72, 0x83,
	0x8c, 0xcf, 0x0d, 0x32, 0x3a, 0x37, 0x01, 0xd7, 0x83, 0x49, 0x0f, 0x33, 0x39, 0x22, 0x6c, 0x0c,
	0x54, 0x53, 0x36, 0xa0, 0x5c, 0xec, 0xe6, 0xba, 0xea, 0xbf, 0x90, 0x19, 0xd9, 0xd7, 0x76, 0xd9,
	0xac, 0x1f, 0x34, 0xa7, 0xe7, 0x21, 0xa8, 0xde, 0xbf, 0xb8, 0xb2, 0xe6, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x81, 0x51, 0xe4, 0x39, 0x31, 0x02, 0x00, 0x00,
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
	if m.NextConnectionSequence != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextConnectionSequence))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ClientConnectionPaths) > 0 {
		for iNdEx := len(m.ClientConnectionPaths) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientConnectionPaths[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Connections) > 0 {
		for iNdEx := len(m.Connections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Connections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
	if len(m.Connections) > 0 {
		for _, e := range m.Connections {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClientConnectionPaths) > 0 {
		for _, e := range m.ClientConnectionPaths {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NextConnectionSequence != 0 {
		n += 1 + sovGenesis(uint64(m.NextConnectionSequence))
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Connections", wireType)
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
			m.Connections = append(m.Connections, IdentifiedConnection{})
			if err := m.Connections[len(m.Connections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientConnectionPaths", wireType)
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
			m.ClientConnectionPaths = append(m.ClientConnectionPaths, ConnectionPaths{})
			if err := m.ClientConnectionPaths[len(m.ClientConnectionPaths)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextConnectionSequence", wireType)
			}
			m.NextConnectionSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextConnectionSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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