// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/ibcratelimit/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// ParamsRequest is the request type for the Query/Params RPC method.
type ParamsRequest struct {
}

func (m *ParamsRequest) Reset()         { *m = ParamsRequest{} }
func (m *ParamsRequest) String() string { return proto.CompactTextString(m) }
func (*ParamsRequest) ProtoMessage()    {}
func (*ParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8427b621352b6b3f, []int{0}
}
func (m *ParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsRequest.Merge(m, src)
}
func (m *ParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsRequest proto.InternalMessageInfo

// aramsResponse is the response type for the Query/Params RPC method.
type ParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *ParamsResponse) Reset()         { *m = ParamsResponse{} }
func (m *ParamsResponse) String() string { return proto.CompactTextString(m) }
func (*ParamsResponse) ProtoMessage()    {}
func (*ParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8427b621352b6b3f, []int{1}
}
func (m *ParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsResponse.Merge(m, src)
}
func (m *ParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsResponse proto.InternalMessageInfo

func (m *ParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*ParamsRequest)(nil), "comdex.ibcratelimit.v1beta1.ParamsRequest")
	proto.RegisterType((*ParamsResponse)(nil), "comdex.ibcratelimit.v1beta1.ParamsResponse")
}

func init() {
	proto.RegisterFile("comdex/ibcratelimit/v1beta1/query.proto", fileDescriptor_8427b621352b6b3f)
}

var fileDescriptor_8427b621352b6b3f = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0x9b, 0x1f, 0x3f, 0x77, 0x88, 0xa8, 0x50, 0x3c, 0xc8, 0x94, 0x28, 0x1b, 0xe2, 0xd8,
	0xb0, 0x61, 0x13, 0xbc, 0xbb, 0x57, 0x30, 0xe7, 0xcd, 0x5b, 0x5a, 0xb3, 0x18, 0x68, 0xfb, 0x64,
	0x4d, 0x2a, 0xee, 0xea, 0x0b, 0x10, 0xc1, 0x8b, 0x2f, 0x69, 0xc7, 0x81, 0x17, 0x4f, 0x22, 0xad,
	0x2f, 0x44, 0xda, 0x64, 0x22, 0x1e, 0xe6, 0x6e, 0xf9, 0xf3, 0xf9, 0xfe, 0x79, 0x1e, 0x7c, 0x12,
	0x41, 0x72, 0xc3, 0xef, 0xa9, 0x0c, 0xa3, 0x8c, 0x19, 0x1e, 0xcb, 0x44, 0x1a, 0x7a, 0xd7, 0x0f,
	0xb9, 0x61, 0x7d, 0x3a, 0xcd, 0x79, 0x36, 0x0b, 0x54, 0x06, 0x06, 0xfc, 0x7d, 0x0b, 0x06, 0x3f,
	0xc1, 0xc0, 0x81, 0xcd, 0x5d, 0x01, 0x02, 0x6a, 0x8e, 0x56, 0x27, 0x2b, 0x69, 0x1e, 0x08, 0x00,
	0x11, 0x73, 0xca, 0x94, 0xa4, 0x2c, 0x4d, 0xc1, 0x30, 0x23, 0x21, 0xd5, 0xee, 0xb7, 0x1b, 0x81,
	0x4e, 0x40, 0xd3, 0x90, 0x69, 0x6e, 0x93, 0xbe, 0x73, 0x15, 0x13, 0x32, 0xad, 0x61, 0xc7, 0x76,
	0x56, 0xb5, 0x54, 0x2c, 0x63, 0x89, 0x73, 0x6d, 0xed, 0xe0, 0xad, 0x51, 0x7d, 0x1f, 0xf3, 0x69,
	0xce, 0xb5, 0x69, 0x5d, 0xe1, 0xed, 0xe5, 0x83, 0x56, 0x90, 0x6a, 0xee, 0x5f, 0xe0, 0x86, 0x95,
	0xec, 0xa1, 0x23, 0xd4, 0xd9, 0x1c, 0xb4, 0x83, 0x15, 0xa3, 0x05, 0x56, 0x3c, 0xfc, 0x3f, 0x7f,
	0x3f, 0xf4, 0xc6, 0x4e, 0x38, 0x78, 0x41, 0x78, 0xe3, 0xb2, 0xaa, 0xec, 0x3f, 0x22, 0xdc, 0xb0,
	0x88, 0xdf, 0x5d, 0xc3, 0xc7, 0xb5, 0x6a, 0xf6, 0xd6, 0x62, 0x6d, 0xe1, 0x56, 0xef, 0xe1, 0xf5,
	0xf3, 0xf9, 0xdf, 0xb1, 0xdf, 0xa6, 0x7f, 0xaf, 0x61, 0x38, 0x9a, 0x17, 0x04, 0x2d, 0x0a, 0x82,
	0x3e, 0x0a, 0x82, 0x9e, 0x4a, 0xe2, 0x2d, 0x4a, 0xe2, 0xbd, 0x95, 0xc4, 0xbb, 0x3e, 0x17, 0xd2,
	0xdc, 0xe6, 0x61, 0x95, 0xec, 0x8c, 0x4e, 0x61, 0x32, 0x91, 0x91, 0x64, 0xf1, 0xd2, 0xf8, 0x97,
	0xb5, 0x99, 0x29, 0xae, 0xc3, 0x46, 0xbd, 0xd9, 0xb3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83,
	0xa6, 0x44, 0x57, 0x2b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params defines a gRPC query method that returns the ibcratelimit module's
	// parameters.
	Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error) {
	out := new(ParamsResponse)
	err := c.cc.Invoke(ctx, "/comdex.ibcratelimit.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params defines a gRPC query method that returns the ibcratelimit module's
	// parameters.
	Params(context.Context, *ParamsRequest) (*ParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *ParamsRequest) (*ParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.ibcratelimit.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*ParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comdex.ibcratelimit.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comdex/ibcratelimit/v1beta1/query.proto",
}

func (m *ParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *ParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
