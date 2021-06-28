// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: a.proto

package proto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age      uint32 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Birthday string `protobuf:"bytes,3,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Address  string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a_4e5d658bc99589e9, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
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

func (m *User) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Work struct {
	Company string `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Code    string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	IpaTime string `protobuf:"bytes,5,opt,name=ipa_time,json=ipaTime,proto3" json:"ipa_time,omitempty"`
}

func (m *Work) Reset()         { *m = Work{} }
func (m *Work) String() string { return proto.CompactTextString(m) }
func (*Work) ProtoMessage()    {}
func (*Work) Descriptor() ([]byte, []int) {
	return fileDescriptor_a_4e5d658bc99589e9, []int{1}
}
func (m *Work) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Work) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Work.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Work) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Work.Merge(dst, src)
}
func (m *Work) XXX_Size() int {
	return m.Size()
}
func (m *Work) XXX_DiscardUnknown() {
	xxx_messageInfo_Work.DiscardUnknown(m)
}

var xxx_messageInfo_Work proto.InternalMessageInfo

func (m *Work) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *Work) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Work) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Work) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Work) GetIpaTime() string {
	if m != nil {
		return m.IpaTime
	}
	return ""
}

type User_Work struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age      uint32 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Birthday string `protobuf:"bytes,3,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Address  string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Work     *Work  `protobuf:"bytes,6,opt,name=work" json:"work,omitempty"`
}

func (m *User_Work) Reset()         { *m = User_Work{} }
func (m *User_Work) String() string { return proto.CompactTextString(m) }
func (*User_Work) ProtoMessage()    {}
func (*User_Work) Descriptor() ([]byte, []int) {
	return fileDescriptor_a_4e5d658bc99589e9, []int{2}
}
func (m *User_Work) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User_Work) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User_Work.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *User_Work) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User_Work.Merge(dst, src)
}
func (m *User_Work) XXX_Size() int {
	return m.Size()
}
func (m *User_Work) XXX_DiscardUnknown() {
	xxx_messageInfo_User_Work.DiscardUnknown(m)
}

var xxx_messageInfo_User_Work proto.InternalMessageInfo

func (m *User_Work) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User_Work) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User_Work) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *User_Work) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User_Work) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User_Work) GetWork() *Work {
	if m != nil {
		return m.Work
	}
	return nil
}

type PutOne struct {
	Users []*User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
	Works []*Work `protobuf:"bytes,2,rep,name=works" json:"works,omitempty"`
}

func (m *PutOne) Reset()         { *m = PutOne{} }
func (m *PutOne) String() string { return proto.CompactTextString(m) }
func (*PutOne) ProtoMessage()    {}
func (*PutOne) Descriptor() ([]byte, []int) {
	return fileDescriptor_a_4e5d658bc99589e9, []int{3}
}
func (m *PutOne) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PutOne) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PutOne.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PutOne) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutOne.Merge(dst, src)
}
func (m *PutOne) XXX_Size() int {
	return m.Size()
}
func (m *PutOne) XXX_DiscardUnknown() {
	xxx_messageInfo_PutOne.DiscardUnknown(m)
}

var xxx_messageInfo_PutOne proto.InternalMessageInfo

func (m *PutOne) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *PutOne) GetWorks() []*Work {
	if m != nil {
		return m.Works
	}
	return nil
}

type PutTwo struct {
	Uws []*User_Work `protobuf:"bytes,1,rep,name=uws" json:"uws,omitempty"`
}

func (m *PutTwo) Reset()         { *m = PutTwo{} }
func (m *PutTwo) String() string { return proto.CompactTextString(m) }
func (*PutTwo) ProtoMessage()    {}
func (*PutTwo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a_4e5d658bc99589e9, []int{4}
}
func (m *PutTwo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PutTwo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PutTwo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PutTwo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutTwo.Merge(dst, src)
}
func (m *PutTwo) XXX_Size() int {
	return m.Size()
}
func (m *PutTwo) XXX_DiscardUnknown() {
	xxx_messageInfo_PutTwo.DiscardUnknown(m)
}

var xxx_messageInfo_PutTwo proto.InternalMessageInfo

func (m *PutTwo) GetUws() []*User_Work {
	if m != nil {
		return m.Uws
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*Work)(nil), "proto.Work")
	proto.RegisterType((*User_Work)(nil), "proto.User_Work")
	proto.RegisterType((*PutOne)(nil), "proto.PutOne")
	proto.RegisterType((*PutTwo)(nil), "proto.PutTwo")
}
func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintA(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Age != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintA(dAtA, i, uint64(m.Age))
	}
	if len(m.Birthday) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintA(dAtA, i, uint64(len(m.Birthday)))
		i += copy(dAtA[i:], m.Birthday)
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintA(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.Email) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintA(dAtA, i, uint64(len(m.Email)))
		i += copy(dAtA[i:], m.Email)
	}
	return i, nil
}

func (m *Work) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Work) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Company) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintA(dAtA, i, uint64(len(m.Company)))
		i += copy(dAtA[i:], m.Company)
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintA(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.Email) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintA(dAtA, i, uint64(len(m.Email)))
		i += copy(dAtA[i:], m.Email)
	}
	if len(m.Code) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintA(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	if len(m.IpaTime) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintA(dAtA, i, uint64(len(m.IpaTime)))
		i += copy(dAtA[i:], m.IpaTime)
	}
	return i, nil
}

func (m *User_Work) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User_Work) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintA(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Age != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintA(dAtA, i, uint64(m.Age))
	}
	if len(m.Birthday) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintA(dAtA, i, uint64(len(m.Birthday)))
		i += copy(dAtA[i:], m.Birthday)
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintA(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.Email) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintA(dAtA, i, uint64(len(m.Email)))
		i += copy(dAtA[i:], m.Email)
	}
	if m.Work != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintA(dAtA, i, uint64(m.Work.Size()))
		n1, err := m.Work.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *PutOne) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PutOne) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Users) > 0 {
		for _, msg := range m.Users {
			dAtA[i] = 0xa
			i++
			i = encodeVarintA(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Works) > 0 {
		for _, msg := range m.Works {
			dAtA[i] = 0x12
			i++
			i = encodeVarintA(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *PutTwo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PutTwo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uws) > 0 {
		for _, msg := range m.Uws {
			dAtA[i] = 0xa
			i++
			i = encodeVarintA(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintA(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovA(uint64(m.Age))
	}
	l = len(m.Birthday)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	return n
}

func (m *Work) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Company)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	l = len(m.IpaTime)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	return n
}

func (m *User_Work) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovA(uint64(m.Age))
	}
	l = len(m.Birthday)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	if m.Work != nil {
		l = m.Work.Size()
		n += 1 + l + sovA(uint64(l))
	}
	return n
}

func (m *PutOne) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Users) > 0 {
		for _, e := range m.Users {
			l = e.Size()
			n += 1 + l + sovA(uint64(l))
		}
	}
	if len(m.Works) > 0 {
		for _, e := range m.Works {
			l = e.Size()
			n += 1 + l + sovA(uint64(l))
		}
	}
	return n
}

func (m *PutTwo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Uws) > 0 {
		for _, e := range m.Uws {
			l = e.Size()
			n += 1 + l + sovA(uint64(l))
		}
	}
	return n
}

func sovA(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozA(x uint64) (n int) {
	return sovA(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowA
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Birthday", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Birthday = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipA(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthA
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
func (m *Work) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowA
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Work: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Work: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Company", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Company = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpaTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpaTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipA(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthA
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
func (m *User_Work) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowA
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: User_Work: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User_Work: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Birthday", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Birthday = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Work", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Work == nil {
				m.Work = &Work{}
			}
			if err := m.Work.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipA(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthA
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
func (m *PutOne) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowA
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PutOne: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PutOne: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, &User{})
			if err := m.Users[len(m.Users)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Works", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Works = append(m.Works, &Work{})
			if err := m.Works[len(m.Works)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipA(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthA
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
func (m *PutTwo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowA
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PutTwo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PutTwo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uws", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uws = append(m.Uws, &User_Work{})
			if err := m.Uws[len(m.Uws)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipA(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthA
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
func skipA(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowA
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
					return 0, ErrIntOverflowA
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowA
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthA
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowA
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipA(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthA = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowA   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("a.proto", fileDescriptor_a_4e5d658bc99589e9) }

var fileDescriptor_a_4e5d658bc99589e9 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x91, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0xeb, 0x3a, 0xe9, 0x9f, 0x57, 0x21, 0x55, 0x16, 0x83, 0x61, 0x30, 0xc1, 0x53, 0x07,
	0xd4, 0x01, 0x6e, 0xc0, 0x01, 0x00, 0x45, 0x45, 0x8c, 0x95, 0xdb, 0x58, 0x60, 0x95, 0xc4, 0x91,
	0x93, 0x28, 0x74, 0xe2, 0x0a, 0x9c, 0x80, 0xf3, 0x30, 0x76, 0x64, 0x44, 0xc9, 0x45, 0x90, 0x9d,
	0x04, 0x45, 0x48, 0xac, 0x4c, 0x79, 0x5f, 0xbe, 0xcf, 0xbf, 0xf7, 0xfc, 0x0c, 0x63, 0xb1, 0x4c,
	0x8d, 0xce, 0x35, 0xf1, 0xdd, 0x87, 0xbf, 0x80, 0x77, 0x9f, 0x49, 0x43, 0x08, 0x78, 0x89, 0x88,
	0x25, 0x45, 0x01, 0x5a, 0x4c, 0x43, 0x57, 0x93, 0x39, 0x60, 0xf1, 0x28, 0xe9, 0x30, 0x40, 0x8b,
	0xa3, 0xd0, 0x96, 0xe4, 0x14, 0x26, 0x1b, 0x65, 0xf2, 0xa7, 0x48, 0xec, 0x29, 0x76, 0xc9, 0x1f,
	0x4d, 0x28, 0x8c, 0x45, 0x14, 0x19, 0x99, 0x65, 0xd4, 0x73, 0x56, 0x27, 0xc9, 0x31, 0xf8, 0x32,
	0x16, 0xea, 0x99, 0xfa, 0xee, 0x7f, 0x23, 0xf8, 0x2b, 0x78, 0x0f, 0xda, 0xec, 0xec, 0xb9, 0xad,
	0x8e, 0x53, 0x91, 0xec, 0xdb, 0xe6, 0x9d, 0xec, 0x13, 0x87, 0x7f, 0x10, 0x71, 0x8f, 0x68, 0xef,
	0xb0, 0xd5, 0x91, 0x6c, 0xdb, 0xbb, 0x9a, 0x9c, 0xc0, 0x44, 0xa5, 0x62, 0x9d, 0xab, 0x58, 0xb6,
	0xed, 0xc7, 0x2a, 0x15, 0x2b, 0x15, 0x4b, 0xfe, 0x8e, 0x60, 0x6a, 0xef, 0xbe, 0x76, 0x63, 0xfc,
	0xfb, 0x02, 0xc8, 0x19, 0x78, 0xa5, 0x36, 0x3b, 0x3a, 0x0a, 0xd0, 0x62, 0x76, 0x39, 0x6b, 0xde,
	0x65, 0x69, 0x87, 0x09, 0x9d, 0xc1, 0x6f, 0x60, 0x74, 0x57, 0xe4, 0xb7, 0x89, 0x24, 0xe7, 0xe0,
	0x17, 0x99, 0x34, 0x19, 0x45, 0x01, 0xee, 0x65, 0xed, 0xf4, 0x61, 0xe3, 0xd8, 0x88, 0x3d, 0x64,
	0x57, 0x85, 0x7f, 0xe3, 0x1a, 0x87, 0x5f, 0x38, 0xde, 0xaa, 0xd4, 0x84, 0x03, 0x2e, 0xca, 0x8e,
	0x36, 0xef, 0xd1, 0xdc, 0x2e, 0x42, 0x6b, 0x5e, 0xd3, 0x8f, 0x8a, 0xa1, 0x43, 0xc5, 0xd0, 0x57,
	0xc5, 0xd0, 0x5b, 0xcd, 0x06, 0x87, 0x9a, 0x0d, 0x3e, 0x6b, 0x36, 0xd8, 0x8c, 0x5c, 0xfe, 0xea,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xbc, 0x4d, 0x27, 0x4c, 0x02, 0x00, 0x00,
}
