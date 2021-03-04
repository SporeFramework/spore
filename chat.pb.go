// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.proto

package main

import (
	fmt "fmt"
	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request_Type int32

const (
	Request_SEND_TRANSACTION Request_Type = 0
	Request_UPDATE_PEER      Request_Type = 1
)

var Request_Type_name = map[int32]string{
	0: "SEND_TRANSACTION",
	1: "UPDATE_PEER",
}

var Request_Type_value = map[string]int32{
	"SEND_TRANSACTION": 0,
	"UPDATE_PEER":      1,
}

func (x Request_Type) Enum() *Request_Type {
	p := new(Request_Type)
	*p = x
	return p
}

func (x Request_Type) String() string {
	return proto.EnumName(Request_Type_name, int32(x))
}

func (x *Request_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Request_Type_value, data, "Request_Type")
	if err != nil {
		return err
	}
	*x = Request_Type(value)
	return nil
}

func (Request_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0, 0}
}

type Request struct {
	Type                 *Request_Type    `protobuf:"varint,1,req,name=type,enum=main.Request_Type" json:"type,omitempty"`
	SendTransaction      *SendTransaction `protobuf:"bytes,2,opt,name=sendTransaction" json:"sendTransaction,omitempty"`
	UpdatePeer           *UpdatePeer      `protobuf:"bytes,3,opt,name=updatePeer" json:"updatePeer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetType() Request_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Request_SEND_TRANSACTION
}

func (m *Request) GetSendTransaction() *SendTransaction {
	if m != nil {
		return m.SendTransaction
	}
	return nil
}

func (m *Request) GetUpdatePeer() *UpdatePeer {
	if m != nil {
		return m.UpdatePeer
	}
	return nil
}

type SendTransaction struct {
	Data                 []byte   `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	Created              *int64   `protobuf:"varint,2,req,name=created" json:"created,omitempty"`
	Id                   []byte   `protobuf:"bytes,3,req,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendTransaction) Reset()         { *m = SendTransaction{} }
func (m *SendTransaction) String() string { return proto.CompactTextString(m) }
func (*SendTransaction) ProtoMessage()    {}
func (*SendTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}
func (m *SendTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendTransaction.Merge(m, src)
}
func (m *SendTransaction) XXX_Size() int {
	return m.Size()
}
func (m *SendTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_SendTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_SendTransaction proto.InternalMessageInfo

func (m *SendTransaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SendTransaction) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *SendTransaction) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type UpdatePeer struct {
	UserHandle           []byte   `protobuf:"bytes,1,opt,name=userHandle" json:"userHandle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePeer) Reset()         { *m = UpdatePeer{} }
func (m *UpdatePeer) String() string { return proto.CompactTextString(m) }
func (*UpdatePeer) ProtoMessage()    {}
func (*UpdatePeer) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}
func (m *UpdatePeer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdatePeer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdatePeer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdatePeer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePeer.Merge(m, src)
}
func (m *UpdatePeer) XXX_Size() int {
	return m.Size()
}
func (m *UpdatePeer) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePeer.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePeer proto.InternalMessageInfo

func (m *UpdatePeer) GetUserHandle() []byte {
	if m != nil {
		return m.UserHandle
	}
	return nil
}

func init() {
	proto.RegisterEnum("main.Request_Type", Request_Type_name, Request_Type_value)
	proto.RegisterType((*Request)(nil), "main.Request")
	proto.RegisterType((*SendTransaction)(nil), "main.SendTransaction")
	proto.RegisterType((*UpdatePeer)(nil), "main.UpdatePeer")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x6a, 0x83, 0x40,
	0x10, 0xc6, 0xb3, 0xab, 0x10, 0x98, 0x84, 0x28, 0x43, 0x0b, 0x7b, 0x12, 0xf1, 0x50, 0x3c, 0xb4,
	0x52, 0xf2, 0x02, 0xc5, 0x36, 0x42, 0x7b, 0x31, 0xb2, 0x9a, 0x73, 0x58, 0xdc, 0x85, 0x0a, 0xad,
	0x5a, 0x5d, 0x0f, 0x79, 0x93, 0x3e, 0x52, 0x8f, 0x85, 0xbe, 0x40, 0xb1, 0x2f, 0x52, 0x34, 0xfd,
	0x93, 0xe4, 0x36, 0xf3, 0x7d, 0xbf, 0x0f, 0xbe, 0x19, 0x80, 0xfc, 0x51, 0xe8, 0xa0, 0x6e, 0x2a,
	0x5d, 0xa1, 0xf9, 0x2c, 0x8a, 0xd2, 0xfb, 0x20, 0x30, 0xe5, 0xea, 0xa5, 0x53, 0xad, 0xc6, 0x0b,
	0x30, 0xf5, 0xae, 0x56, 0x8c, 0xb8, 0xd4, 0x5f, 0x2c, 0x31, 0x18, 0x80, 0xe0, 0xc7, 0x0c, 0xb2,
	0x5d, 0xad, 0xf8, 0xe8, 0xe3, 0x0d, 0x58, 0xad, 0x2a, 0x65, 0xd6, 0x88, 0xb2, 0x15, 0xb9, 0x2e,
	0xaa, 0x92, 0x51, 0x97, 0xf8, 0xb3, 0xe5, 0xf9, 0x3e, 0x92, 0x1e, 0x9b, 0xfc, 0x94, 0xc6, 0x6b,
	0x80, 0xae, 0x96, 0x42, 0xab, 0x44, 0xa9, 0x86, 0x19, 0x63, 0xd6, 0xde, 0x67, 0x37, 0x7f, 0x3a,
	0x3f, 0x60, 0xbc, 0x2b, 0x30, 0x87, 0x02, 0x78, 0x06, 0x76, 0x1a, 0xc5, 0xab, 0x6d, 0xc6, 0xc3,
	0x38, 0x0d, 0xef, 0xb2, 0x87, 0x75, 0x6c, 0x4f, 0xd0, 0x82, 0xd9, 0x26, 0x59, 0x85, 0x59, 0xb4,
	0x4d, 0xa2, 0x88, 0xdb, 0xc4, 0x5b, 0x83, 0x75, 0x52, 0x02, 0x11, 0x4c, 0x29, 0xb4, 0x18, 0x8f,
	0x9b, 0xf3, 0x71, 0x46, 0x06, 0xd3, 0xbc, 0x51, 0x42, 0x2b, 0xc9, 0xa8, 0x4b, 0x7d, 0x83, 0xff,
	0xae, 0xb8, 0x00, 0x5a, 0x48, 0x66, 0x8c, 0x2c, 0x2d, 0xa4, 0x77, 0x09, 0xf0, 0xdf, 0x0c, 0x1d,
	0x80, 0xae, 0x55, 0xcd, 0xbd, 0x28, 0xe5, 0xd3, 0xf0, 0x2e, 0xe2, 0xcf, 0xf9, 0x81, 0x72, 0x6b,
	0xbf, 0xf5, 0x0e, 0x79, 0xef, 0x1d, 0xf2, 0xd9, 0x3b, 0xe4, 0xf5, 0xcb, 0x99, 0x7c, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xfc, 0x1a, 0xe4, 0xfd, 0x79, 0x01, 0x00, 0x00,
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UpdatePeer != nil {
		{
			size, err := m.UpdatePeer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintChat(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.SendTransaction != nil {
		{
			size, err := m.SendTransaction.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintChat(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Type == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i = encodeVarintChat(dAtA, i, uint64(*m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SendTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintChat(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Created == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i = encodeVarintChat(dAtA, i, uint64(*m.Created))
		i--
		dAtA[i] = 0x10
	}
	if m.Data == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintChat(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdatePeer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdatePeer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdatePeer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UserHandle != nil {
		i -= len(m.UserHandle)
		copy(dAtA[i:], m.UserHandle)
		i = encodeVarintChat(dAtA, i, uint64(len(m.UserHandle)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintChat(dAtA []byte, offset int, v uint64) int {
	offset -= sovChat(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != nil {
		n += 1 + sovChat(uint64(*m.Type))
	}
	if m.SendTransaction != nil {
		l = m.SendTransaction.Size()
		n += 1 + l + sovChat(uint64(l))
	}
	if m.UpdatePeer != nil {
		l = m.UpdatePeer.Size()
		n += 1 + l + sovChat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SendTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovChat(uint64(l))
	}
	if m.Created != nil {
		n += 1 + sovChat(uint64(*m.Created))
	}
	if m.Id != nil {
		l = len(m.Id)
		n += 1 + l + sovChat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UpdatePeer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserHandle != nil {
		l = len(m.UserHandle)
		n += 1 + l + sovChat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovChat(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChat(x uint64) (n int) {
	return sovChat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var v Request_Type
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= Request_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Type = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendTransaction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SendTransaction == nil {
				m.SendTransaction = &SendTransaction{}
			}
			if err := m.SendTransaction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatePeer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpdatePeer == nil {
				m.UpdatePeer = &UpdatePeer{}
			}
			if err := m.UpdatePeer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SendTransaction) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
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
			return fmt.Errorf("proto: SendTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Created = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpdatePeer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
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
			return fmt.Errorf("proto: UpdatePeer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdatePeer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserHandle", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserHandle = append(m.UserHandle[:0], dAtA[iNdEx:postIndex]...)
			if m.UserHandle == nil {
				m.UserHandle = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipChat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChat
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
					return 0, ErrIntOverflowChat
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
					return 0, ErrIntOverflowChat
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
				return 0, ErrInvalidLengthChat
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChat
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChat
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChat        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChat          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChat = fmt.Errorf("proto: unexpected end of group")
)
