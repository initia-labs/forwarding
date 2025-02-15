// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: noble/forwarding/v1/packet.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type RegisterAccountData struct {
	Recipient string `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Channel   string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (m *RegisterAccountData) Reset()         { *m = RegisterAccountData{} }
func (m *RegisterAccountData) String() string { return proto.CompactTextString(m) }
func (*RegisterAccountData) ProtoMessage()    {}
func (*RegisterAccountData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a0a2a88e68b1d25, []int{0}
}
func (m *RegisterAccountData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterAccountData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterAccountData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterAccountData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAccountData.Merge(m, src)
}
func (m *RegisterAccountData) XXX_Size() int {
	return m.Size()
}
func (m *RegisterAccountData) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAccountData.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAccountData proto.InternalMessageInfo

func (m *RegisterAccountData) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *RegisterAccountData) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

type RegisterAccountMemo struct {
	Noble *RegisterAccountMemo_RegisterAccountDataWrapper `protobuf:"bytes,1,opt,name=noble,proto3" json:"noble,omitempty"`
}

func (m *RegisterAccountMemo) Reset()         { *m = RegisterAccountMemo{} }
func (m *RegisterAccountMemo) String() string { return proto.CompactTextString(m) }
func (*RegisterAccountMemo) ProtoMessage()    {}
func (*RegisterAccountMemo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a0a2a88e68b1d25, []int{1}
}
func (m *RegisterAccountMemo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterAccountMemo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterAccountMemo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterAccountMemo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAccountMemo.Merge(m, src)
}
func (m *RegisterAccountMemo) XXX_Size() int {
	return m.Size()
}
func (m *RegisterAccountMemo) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAccountMemo.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAccountMemo proto.InternalMessageInfo

func (m *RegisterAccountMemo) GetNoble() *RegisterAccountMemo_RegisterAccountDataWrapper {
	if m != nil {
		return m.Noble
	}
	return nil
}

type RegisterAccountMemo_RegisterAccountDataWrapper struct {
	Forwarding *RegisterAccountData `protobuf:"bytes,1,opt,name=forwarding,proto3" json:"forwarding,omitempty"`
}

func (m *RegisterAccountMemo_RegisterAccountDataWrapper) Reset() {
	*m = RegisterAccountMemo_RegisterAccountDataWrapper{}
}
func (m *RegisterAccountMemo_RegisterAccountDataWrapper) String() string {
	return proto.CompactTextString(m)
}
func (*RegisterAccountMemo_RegisterAccountDataWrapper) ProtoMessage() {}
func (*RegisterAccountMemo_RegisterAccountDataWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a0a2a88e68b1d25, []int{1, 0}
}
func (m *RegisterAccountMemo_RegisterAccountDataWrapper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterAccountMemo_RegisterAccountDataWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterAccountMemo_RegisterAccountDataWrapper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterAccountMemo_RegisterAccountDataWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAccountMemo_RegisterAccountDataWrapper.Merge(m, src)
}
func (m *RegisterAccountMemo_RegisterAccountDataWrapper) XXX_Size() int {
	return m.Size()
}
func (m *RegisterAccountMemo_RegisterAccountDataWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAccountMemo_RegisterAccountDataWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAccountMemo_RegisterAccountDataWrapper proto.InternalMessageInfo

func (m *RegisterAccountMemo_RegisterAccountDataWrapper) GetForwarding() *RegisterAccountData {
	if m != nil {
		return m.Forwarding
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterAccountData)(nil), "noble.forwarding.v1.RegisterAccountData")
	proto.RegisterType((*RegisterAccountMemo)(nil), "noble.forwarding.v1.RegisterAccountMemo")
	proto.RegisterType((*RegisterAccountMemo_RegisterAccountDataWrapper)(nil), "noble.forwarding.v1.RegisterAccountMemo.RegisterAccountDataWrapper")
}

func init() { proto.RegisterFile("noble/forwarding/v1/packet.proto", fileDescriptor_9a0a2a88e68b1d25) }

var fileDescriptor_9a0a2a88e68b1d25 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0xcb, 0x4f, 0xca,
	0x49, 0xd5, 0x4f, 0xcb, 0x2f, 0x2a, 0x4f, 0x2c, 0x4a, 0xc9, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0xd4,
	0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x06, 0xab,
	0xd0, 0x43, 0xa8, 0xd0, 0x2b, 0x33, 0x54, 0xf2, 0xe5, 0x12, 0x0e, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e,
	0x49, 0x2d, 0x72, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x71, 0x49, 0x2c, 0x49, 0x14, 0x92, 0xe1,
	0xe2, 0x2c, 0x4a, 0x4d, 0xce, 0x2c, 0xc8, 0x4c, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x42, 0x08, 0x08, 0x49, 0x70, 0xb1, 0x27, 0x67, 0x24, 0xe6, 0xe5, 0xa5, 0xe6, 0x48, 0x30,
	0x81, 0xe5, 0x60, 0x5c, 0xa5, 0x1b, 0x8c, 0x18, 0xe6, 0xf9, 0xa6, 0xe6, 0xe6, 0x0b, 0x45, 0x72,
	0xb1, 0x82, 0x6d, 0x07, 0x9b, 0xc5, 0x6d, 0xe4, 0xac, 0x87, 0xc5, 0x2d, 0x7a, 0x58, 0x34, 0xea,
	0x61, 0x71, 0x5c, 0x78, 0x51, 0x62, 0x41, 0x41, 0x6a, 0x51, 0x10, 0xc4, 0x44, 0xa9, 0x34, 0x2e,
	0x29, 0xdc, 0x8a, 0x84, 0x3c, 0xb8, 0xb8, 0x10, 0x96, 0x40, 0x6d, 0xd7, 0x20, 0xc6, 0x76, 0x90,
	0x21, 0x41, 0x48, 0x7a, 0x9d, 0xfc, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1,
	0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21,
	0xca, 0x34, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x6c, 0xb2, 0x6e,
	0x62, 0x71, 0x71, 0x6a, 0x49, 0x31, 0x72, 0x64, 0x54, 0x20, 0x73, 0x4a, 0x2a, 0x0b, 0x52, 0x8b,
	0x93, 0xd8, 0xc0, 0xd1, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x6f, 0x72, 0x9b, 0xba,
	0x01, 0x00, 0x00,
}

func (m *RegisterAccountData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterAccountData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterAccountData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterAccountMemo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterAccountMemo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterAccountMemo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Noble != nil {
		{
			size, err := m.Noble.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterAccountMemo_RegisterAccountDataWrapper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterAccountMemo_RegisterAccountDataWrapper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterAccountMemo_RegisterAccountDataWrapper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Forwarding != nil {
		{
			size, err := m.Forwarding.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisterAccountData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *RegisterAccountMemo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Noble != nil {
		l = m.Noble.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *RegisterAccountMemo_RegisterAccountDataWrapper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Forwarding != nil {
		l = m.Forwarding.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterAccountData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: RegisterAccountData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterAccountData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *RegisterAccountMemo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: RegisterAccountMemo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterAccountMemo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Noble", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Noble == nil {
				m.Noble = &RegisterAccountMemo_RegisterAccountDataWrapper{}
			}
			if err := m.Noble.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *RegisterAccountMemo_RegisterAccountDataWrapper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: RegisterAccountDataWrapper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterAccountDataWrapper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Forwarding", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Forwarding == nil {
				m.Forwarding = &RegisterAccountData{}
			}
			if err := m.Forwarding.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
