// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hello.proto

package proto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

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

type UserContext struct {
}

func (m *UserContext) Reset()         { *m = UserContext{} }
func (m *UserContext) String() string { return proto.CompactTextString(m) }
func (*UserContext) ProtoMessage()    {}
func (*UserContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_438cae56c5a34398, []int{0}
}
func (m *UserContext) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserContext.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserContext.Merge(dst, src)
}
func (m *UserContext) XXX_Size() int {
	return m.Size()
}
func (m *UserContext) XXX_DiscardUnknown() {
	xxx_messageInfo_UserContext.DiscardUnknown(m)
}

var xxx_messageInfo_UserContext proto.InternalMessageInfo

type UserContext_Req struct {
	Id uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *UserContext_Req) Reset()         { *m = UserContext_Req{} }
func (m *UserContext_Req) String() string { return proto.CompactTextString(m) }
func (*UserContext_Req) ProtoMessage()    {}
func (*UserContext_Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_438cae56c5a34398, []int{0, 0}
}
func (m *UserContext_Req) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserContext_Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserContext_Req.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserContext_Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserContext_Req.Merge(dst, src)
}
func (m *UserContext_Req) XXX_Size() int {
	return m.Size()
}
func (m *UserContext_Req) XXX_DiscardUnknown() {
	xxx_messageInfo_UserContext_Req.DiscardUnknown(m)
}

var xxx_messageInfo_UserContext_Req proto.InternalMessageInfo

func (m *UserContext_Req) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserContext_Resp struct {
	Res string `protobuf:"bytes,2,opt,name=Res,proto3" json:"Res,omitempty"`
}

func (m *UserContext_Resp) Reset()         { *m = UserContext_Resp{} }
func (m *UserContext_Resp) String() string { return proto.CompactTextString(m) }
func (*UserContext_Resp) ProtoMessage()    {}
func (*UserContext_Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_438cae56c5a34398, []int{0, 1}
}
func (m *UserContext_Resp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserContext_Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserContext_Resp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserContext_Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserContext_Resp.Merge(dst, src)
}
func (m *UserContext_Resp) XXX_Size() int {
	return m.Size()
}
func (m *UserContext_Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserContext_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_UserContext_Resp proto.InternalMessageInfo

func (m *UserContext_Resp) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

func init() {
	proto.RegisterType((*UserContext)(nil), "proto.UserContext")
	proto.RegisterType((*UserContext_Req)(nil), "proto.UserContext.Req")
	proto.RegisterType((*UserContext_Resp)(nil), "proto.UserContext.Resp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Get(ctx context.Context, in *UserContext_Req, opts ...grpc.CallOption) (*UserContext_Resp, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Get(ctx context.Context, in *UserContext_Req, opts ...grpc.CallOption) (*UserContext_Resp, error) {
	out := new(UserContext_Resp)
	err := c.cc.Invoke(ctx, "/proto.User/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Get(context.Context, *UserContext_Req) (*UserContext_Resp, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserContext_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.User/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Get(ctx, req.(*UserContext_Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _User_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func (m *UserContext) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserContext) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *UserContext_Req) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserContext_Req) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHello(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *UserContext_Resp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserContext_Resp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Res) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHello(dAtA, i, uint64(len(m.Res)))
		i += copy(dAtA[i:], m.Res)
	}
	return i, nil
}

func encodeVarintHello(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UserContext) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UserContext_Req) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovHello(uint64(m.Id))
	}
	return n
}

func (m *UserContext_Resp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Res)
	if l > 0 {
		n += 1 + l + sovHello(uint64(l))
	}
	return n
}

func sovHello(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHello(x uint64) (n int) {
	return sovHello(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserContext) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHello
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
			return fmt.Errorf("proto: UserContext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserContext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHello(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHello
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
func (m *UserContext_Req) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHello
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
			return fmt.Errorf("proto: Req: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Req: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHello
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHello(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHello
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
func (m *UserContext_Resp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHello
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
			return fmt.Errorf("proto: Resp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Res", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHello
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
				return ErrInvalidLengthHello
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Res = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHello(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHello
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
func skipHello(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHello
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
					return 0, ErrIntOverflowHello
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
					return 0, ErrIntOverflowHello
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
				return 0, ErrInvalidLengthHello
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHello
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
				next, err := skipHello(dAtA[start:])
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
	ErrInvalidLengthHello = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHello   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("hello.proto", fileDescriptor_hello_438cae56c5a34398) }

var fileDescriptor_hello_438cae56c5a34398 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x76, 0x5c, 0xdc, 0xa1,
	0xc5, 0xa9, 0x45, 0xce, 0xf9, 0x79, 0x25, 0xa9, 0x15, 0x25, 0x52, 0xa2, 0x5c, 0xcc, 0x41, 0xa9,
	0x85, 0x42, 0x7c, 0x5c, 0x4c, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x4c, 0x9e,
	0x29, 0x52, 0x12, 0x5c, 0x2c, 0x41, 0xa9, 0xc5, 0x05, 0x42, 0x02, 0x20, 0xe9, 0x62, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0xd3, 0xc8, 0x81, 0x8b, 0x05, 0xa4, 0x5f, 0xc8, 0x82, 0x8b,
	0xd9, 0x3d, 0xb5, 0x44, 0x48, 0x0c, 0x62, 0xba, 0x1e, 0x92, 0x99, 0x7a, 0x41, 0xa9, 0x85, 0x52,
	0xe2, 0x58, 0xc5, 0x8b, 0x0b, 0x94, 0x18, 0x9c, 0x24, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48,
	0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1,
	0x58, 0x8e, 0x21, 0x89, 0x0d, 0xac, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xad, 0x25, 0xf7,
	0x2d, 0xb8, 0x00, 0x00, 0x00,
}