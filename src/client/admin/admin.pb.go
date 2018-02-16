// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client/admin/admin.proto

/*
	Package admin is a generated protocol buffer package.

	It is generated from these files:
		client/admin/admin.proto

	It has these top-level messages:
		Op17
		Op
		ExtractRequest
		RestoreRequest
*/
package admin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"
import pps "github.com/pachyderm/pachyderm/src/client/pps"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Op17 struct {
	Object   *pfs.PutObjectRequest      `protobuf:"bytes,2,opt,name=object" json:"object,omitempty"`
	Tag      *pfs.TagObjectRequest      `protobuf:"bytes,3,opt,name=tag" json:"tag,omitempty"`
	Repo     *pfs.CreateRepoRequest     `protobuf:"bytes,4,opt,name=repo" json:"repo,omitempty"`
	Commit   *pfs.BuildCommitRequest    `protobuf:"bytes,5,opt,name=commit" json:"commit,omitempty"`
	Branch   *pfs.SetBranchRequest      `protobuf:"bytes,6,opt,name=branch" json:"branch,omitempty"`
	Pipeline *pps.CreatePipelineRequest `protobuf:"bytes,7,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *Op17) Reset()                    { *m = Op17{} }
func (m *Op17) String() string            { return proto.CompactTextString(m) }
func (*Op17) ProtoMessage()               {}
func (*Op17) Descriptor() ([]byte, []int) { return fileDescriptorAdmin, []int{0} }

func (m *Op17) GetObject() *pfs.PutObjectRequest {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *Op17) GetTag() *pfs.TagObjectRequest {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *Op17) GetRepo() *pfs.CreateRepoRequest {
	if m != nil {
		return m.Repo
	}
	return nil
}

func (m *Op17) GetCommit() *pfs.BuildCommitRequest {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *Op17) GetBranch() *pfs.SetBranchRequest {
	if m != nil {
		return m.Branch
	}
	return nil
}

func (m *Op17) GetPipeline() *pps.CreatePipelineRequest {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type Op struct {
	Op17 *Op17 `protobuf:"bytes,1,opt,name=op17" json:"op17,omitempty"`
}

func (m *Op) Reset()                    { *m = Op{} }
func (m *Op) String() string            { return proto.CompactTextString(m) }
func (*Op) ProtoMessage()               {}
func (*Op) Descriptor() ([]byte, []int) { return fileDescriptorAdmin, []int{1} }

func (m *Op) GetOp17() *Op17 {
	if m != nil {
		return m.Op17
	}
	return nil
}

type ExtractRequest struct {
	// URL is an object storage URL, if it's not "" data will be extracted to
	// this URL rather than returned.
	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
	// NoObjects, if true, will cause extract to omit objects (and tags)
	NoObjects bool `protobuf:"varint,2,opt,name=no_objects,json=noObjects,proto3" json:"no_objects,omitempty"`
	// NoRepos, if true, will cause extract to omit repos, commits and branches.
	NoRepos bool `protobuf:"varint,3,opt,name=no_repos,json=noRepos,proto3" json:"no_repos,omitempty"`
	// NoPipelines, if true, will cause extract to omit pipelines.
	NoPipelines bool `protobuf:"varint,4,opt,name=no_pipelines,json=noPipelines,proto3" json:"no_pipelines,omitempty"`
}

func (m *ExtractRequest) Reset()                    { *m = ExtractRequest{} }
func (m *ExtractRequest) String() string            { return proto.CompactTextString(m) }
func (*ExtractRequest) ProtoMessage()               {}
func (*ExtractRequest) Descriptor() ([]byte, []int) { return fileDescriptorAdmin, []int{2} }

func (m *ExtractRequest) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *ExtractRequest) GetNoObjects() bool {
	if m != nil {
		return m.NoObjects
	}
	return false
}

func (m *ExtractRequest) GetNoRepos() bool {
	if m != nil {
		return m.NoRepos
	}
	return false
}

func (m *ExtractRequest) GetNoPipelines() bool {
	if m != nil {
		return m.NoPipelines
	}
	return false
}

type RestoreRequest struct {
	Op *Op `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	// URL is an object storage URL, if it's not "" data will be restored from
	// this URL.
	URL string `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (m *RestoreRequest) Reset()                    { *m = RestoreRequest{} }
func (m *RestoreRequest) String() string            { return proto.CompactTextString(m) }
func (*RestoreRequest) ProtoMessage()               {}
func (*RestoreRequest) Descriptor() ([]byte, []int) { return fileDescriptorAdmin, []int{3} }

func (m *RestoreRequest) GetOp() *Op {
	if m != nil {
		return m.Op
	}
	return nil
}

func (m *RestoreRequest) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func init() {
	proto.RegisterType((*Op17)(nil), "admin.Op17")
	proto.RegisterType((*Op)(nil), "admin.Op")
	proto.RegisterType((*ExtractRequest)(nil), "admin.ExtractRequest")
	proto.RegisterType((*RestoreRequest)(nil), "admin.RestoreRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	Extract(ctx context.Context, in *ExtractRequest, opts ...grpc.CallOption) (API_ExtractClient, error)
	Restore(ctx context.Context, opts ...grpc.CallOption) (API_RestoreClient, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Extract(ctx context.Context, in *ExtractRequest, opts ...grpc.CallOption) (API_ExtractClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_API_serviceDesc.Streams[0], c.cc, "/admin.API/Extract", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIExtractClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_ExtractClient interface {
	Recv() (*Op, error)
	grpc.ClientStream
}

type aPIExtractClient struct {
	grpc.ClientStream
}

func (x *aPIExtractClient) Recv() (*Op, error) {
	m := new(Op)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) Restore(ctx context.Context, opts ...grpc.CallOption) (API_RestoreClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_API_serviceDesc.Streams[1], c.cc, "/admin.API/Restore", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIRestoreClient{stream}
	return x, nil
}

type API_RestoreClient interface {
	Send(*RestoreRequest) error
	CloseAndRecv() (*google_protobuf.Empty, error)
	grpc.ClientStream
}

type aPIRestoreClient struct {
	grpc.ClientStream
}

func (x *aPIRestoreClient) Send(m *RestoreRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aPIRestoreClient) CloseAndRecv() (*google_protobuf.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(google_protobuf.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for API service

type APIServer interface {
	Extract(*ExtractRequest, API_ExtractServer) error
	Restore(API_RestoreServer) error
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Extract_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExtractRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).Extract(m, &aPIExtractServer{stream})
}

type API_ExtractServer interface {
	Send(*Op) error
	grpc.ServerStream
}

type aPIExtractServer struct {
	grpc.ServerStream
}

func (x *aPIExtractServer) Send(m *Op) error {
	return x.ServerStream.SendMsg(m)
}

func _API_Restore_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(APIServer).Restore(&aPIRestoreServer{stream})
}

type API_RestoreServer interface {
	SendAndClose(*google_protobuf.Empty) error
	Recv() (*RestoreRequest, error)
	grpc.ServerStream
}

type aPIRestoreServer struct {
	grpc.ServerStream
}

func (x *aPIRestoreServer) SendAndClose(m *google_protobuf.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aPIRestoreServer) Recv() (*RestoreRequest, error) {
	m := new(RestoreRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admin.API",
	HandlerType: (*APIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Extract",
			Handler:       _API_Extract_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Restore",
			Handler:       _API_Restore_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "client/admin/admin.proto",
}

func (m *Op17) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Op17) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Object != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAdmin(dAtA, i, uint64(m.Object.Size()))
		n1, err := m.Object.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Tag != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAdmin(dAtA, i, uint64(m.Tag.Size()))
		n2, err := m.Tag.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Repo != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAdmin(dAtA, i, uint64(m.Repo.Size()))
		n3, err := m.Repo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Commit != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintAdmin(dAtA, i, uint64(m.Commit.Size()))
		n4, err := m.Commit.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Branch != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintAdmin(dAtA, i, uint64(m.Branch.Size()))
		n5, err := m.Branch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Pipeline != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintAdmin(dAtA, i, uint64(m.Pipeline.Size()))
		n6, err := m.Pipeline.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func (m *Op) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Op) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Op17 != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAdmin(dAtA, i, uint64(m.Op17.Size()))
		n7, err := m.Op17.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func (m *ExtractRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtractRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.URL) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.URL)))
		i += copy(dAtA[i:], m.URL)
	}
	if m.NoObjects {
		dAtA[i] = 0x10
		i++
		if m.NoObjects {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.NoRepos {
		dAtA[i] = 0x18
		i++
		if m.NoRepos {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.NoPipelines {
		dAtA[i] = 0x20
		i++
		if m.NoPipelines {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *RestoreRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RestoreRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Op != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAdmin(dAtA, i, uint64(m.Op.Size()))
		n8, err := m.Op.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if len(m.URL) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.URL)))
		i += copy(dAtA[i:], m.URL)
	}
	return i, nil
}

func encodeVarintAdmin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Op17) Size() (n int) {
	var l int
	_ = l
	if m.Object != nil {
		l = m.Object.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	if m.Tag != nil {
		l = m.Tag.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	if m.Repo != nil {
		l = m.Repo.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	if m.Commit != nil {
		l = m.Commit.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	if m.Branch != nil {
		l = m.Branch.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	if m.Pipeline != nil {
		l = m.Pipeline.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	return n
}

func (m *Op) Size() (n int) {
	var l int
	_ = l
	if m.Op17 != nil {
		l = m.Op17.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	return n
}

func (m *ExtractRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.URL)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	if m.NoObjects {
		n += 2
	}
	if m.NoRepos {
		n += 2
	}
	if m.NoPipelines {
		n += 2
	}
	return n
}

func (m *RestoreRequest) Size() (n int) {
	var l int
	_ = l
	if m.Op != nil {
		l = m.Op.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	l = len(m.URL)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	return n
}

func sovAdmin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAdmin(x uint64) (n int) {
	return sovAdmin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Op17) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: Op17: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Op17: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Object", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Object == nil {
				m.Object = &pfs.PutObjectRequest{}
			}
			if err := m.Object.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tag == nil {
				m.Tag = &pfs.TagObjectRequest{}
			}
			if err := m.Tag.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Repo == nil {
				m.Repo = &pfs.CreateRepoRequest{}
			}
			if err := m.Repo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Commit == nil {
				m.Commit = &pfs.BuildCommitRequest{}
			}
			if err := m.Commit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Branch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Branch == nil {
				m.Branch = &pfs.SetBranchRequest{}
			}
			if err := m.Branch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pipeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pipeline == nil {
				m.Pipeline = &pps.CreatePipelineRequest{}
			}
			if err := m.Pipeline.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
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
func (m *Op) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: Op: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Op: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op17", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Op17 == nil {
				m.Op17 = &Op17{}
			}
			if err := m.Op17.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
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
func (m *ExtractRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: ExtractRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtractRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoObjects", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NoObjects = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoRepos", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NoRepos = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoPipelines", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NoPipelines = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
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
func (m *RestoreRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: RestoreRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RestoreRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Op == nil {
				m.Op = &Op{}
			}
			if err := m.Op.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
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
func skipAdmin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
				return 0, ErrInvalidLengthAdmin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAdmin
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
				next, err := skipAdmin(dAtA[start:])
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
	ErrInvalidLengthAdmin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdmin   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("client/admin/admin.proto", fileDescriptorAdmin) }

var fileDescriptorAdmin = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x86, 0x9b, 0xb4, 0x6b, 0xd3, 0xaf, 0x68, 0x9a, 0x2c, 0x36, 0xbc, 0x20, 0x0a, 0x44, 0x42,
	0x4c, 0x48, 0x24, 0x6c, 0x48, 0xec, 0x02, 0x07, 0x3a, 0xed, 0x80, 0x84, 0xd4, 0xca, 0xc0, 0xb9,
	0x4a, 0x3b, 0xaf, 0x04, 0xb5, 0xfe, 0x4c, 0xec, 0x48, 0x70, 0xe2, 0x6f, 0x70, 0xe6, 0xd7, 0x70,
	0xe4, 0x27, 0xa0, 0xf2, 0x47, 0x90, 0x3f, 0x27, 0x61, 0xe5, 0xe0, 0x28, 0x79, 0xfd, 0xbc, 0xc9,
	0xeb, 0x37, 0x1f, 0xf0, 0xe5, 0xba, 0x90, 0xca, 0x66, 0xf9, 0xd5, 0xa6, 0x50, 0xfe, 0x9a, 0xea,
	0x12, 0x2d, 0xb2, 0x3d, 0x7a, 0x88, 0xef, 0xae, 0x10, 0x57, 0x6b, 0x99, 0x91, 0xb8, 0xa8, 0xae,
	0x33, 0xb9, 0xd1, 0xf6, 0xab, 0x67, 0xe2, 0xdb, 0xb5, 0x5b, 0x5f, 0x1b, 0xb7, 0xfe, 0x57, 0xb5,
	0x71, 0xcb, 0xab, 0xc9, 0x8f, 0x10, 0x7a, 0x53, 0x7d, 0x7a, 0xce, 0x9e, 0x42, 0x1f, 0x17, 0x9f,
	0xe4, 0xd2, 0xf2, 0xf0, 0x41, 0x70, 0x32, 0x3a, 0x3b, 0x4c, 0x9d, 0x75, 0x56, 0xd9, 0x29, 0xa9,
	0x42, 0x7e, 0xae, 0xa4, 0xb1, 0xa2, 0x86, 0xd8, 0x63, 0xe8, 0xda, 0x7c, 0xc5, 0xbb, 0x37, 0xd8,
	0xf7, 0xf9, 0x6a, 0x97, 0x75, 0x04, 0x7b, 0x02, 0xbd, 0x52, 0x6a, 0xe4, 0x3d, 0x22, 0x8f, 0x88,
	0xbc, 0x28, 0x65, 0x6e, 0xa5, 0x90, 0x1a, 0x1b, 0x94, 0x18, 0x96, 0x41, 0x7f, 0x89, 0x9b, 0x4d,
	0x61, 0xf9, 0x1e, 0xd1, 0x77, 0x88, 0x9e, 0x54, 0xc5, 0xfa, 0xea, 0x82, 0xf4, 0x36, 0x85, 0xc7,
	0x5c, 0xe8, 0x45, 0x99, 0xab, 0xe5, 0x47, 0xde, 0xbf, 0x11, 0xe4, 0x9d, 0xb4, 0x13, 0x52, 0x5b,
	0xdc, 0x43, 0xec, 0x05, 0x44, 0xba, 0xd0, 0x72, 0x5d, 0x28, 0xc9, 0x07, 0x64, 0x88, 0x53, 0x57,
	0x85, 0xcf, 0x33, 0xab, 0xb7, 0x1a, 0x57, 0xcb, 0x26, 0x8f, 0x20, 0x9c, 0x6a, 0x76, 0x1f, 0x7a,
	0xa8, 0x4f, 0xcf, 0x79, 0x40, 0xce, 0x51, 0xea, 0x7f, 0x8b, 0x2b, 0x4f, 0xd0, 0x46, 0xf2, 0x0d,
	0xf6, 0x2f, 0xbf, 0xd8, 0x32, 0x6f, 0x1b, 0x60, 0x07, 0xd0, 0xfd, 0x20, 0xde, 0x92, 0x63, 0x28,
	0xdc, 0x2d, 0xbb, 0x07, 0xa0, 0x70, 0xee, 0x4b, 0x34, 0x54, 0x75, 0x24, 0x86, 0x0a, 0x7d, 0x71,
	0x86, 0x1d, 0x43, 0xa4, 0x70, 0xee, 0xca, 0x30, 0xd4, 0x6d, 0x24, 0x06, 0x0a, 0x5d, 0x51, 0x86,
	0x3d, 0x84, 0x5b, 0x0a, 0xe7, 0x4d, 0x26, 0x43, 0x85, 0x46, 0x62, 0xa4, 0xb0, 0xc9, 0x6d, 0x92,
	0x57, 0xb0, 0x2f, 0xa4, 0xb1, 0x58, 0x36, 0x67, 0x60, 0xc7, 0x10, 0xa2, 0xae, 0x13, 0x0f, 0xdb,
	0xc4, 0x22, 0x44, 0xdd, 0x64, 0x0b, 0xdb, 0x6c, 0x67, 0x16, 0xba, 0xaf, 0x67, 0x6f, 0x58, 0x06,
	0x83, 0xfa, 0x18, 0xec, 0xb0, 0xb6, 0xec, 0x1e, 0x2b, 0xfe, 0xf7, 0xa6, 0xa4, 0xf3, 0x2c, 0x60,
	0x2f, 0x61, 0x50, 0x7f, 0xb6, 0x35, 0xec, 0xc6, 0x88, 0x8f, 0x52, 0x3f, 0xaf, 0x69, 0x33, 0xaf,
	0xe9, 0xa5, 0x9b, 0xd7, 0xa4, 0x73, 0x12, 0x4c, 0x0e, 0x7e, 0x6e, 0xc7, 0xc1, 0xaf, 0xed, 0x38,
	0xf8, 0xbd, 0x1d, 0x07, 0xdf, 0xff, 0x8c, 0x3b, 0x8b, 0x3e, 0x51, 0xcf, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x75, 0xd4, 0xfb, 0x32, 0x06, 0x03, 0x00, 0x00,
}