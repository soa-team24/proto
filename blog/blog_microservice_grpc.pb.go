// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0
// source: blog_microservice.proto

package blog

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BlogService_GetAllBlogs_FullMethodName        = "/BlogService/GetAllBlogs"
	BlogService_GetBlogById_FullMethodName        = "/BlogService/GetBlogById"
	BlogService_PostBlog_FullMethodName           = "/BlogService/PostBlog"
	BlogService_UpdateBlog_FullMethodName         = "/BlogService/UpdateBlog"
	BlogService_DeleteBlog_FullMethodName         = "/BlogService/DeleteBlog"
	BlogService_GetBlogsByAuthorId_FullMethodName = "/BlogService/GetBlogsByAuthorId"
	BlogService_GetAllVotes_FullMethodName        = "/BlogService/GetAllVotes"
	BlogService_GetVotesCount_FullMethodName      = "/BlogService/GetVotesCount"
	BlogService_AddVote_FullMethodName            = "/BlogService/AddVote"
	BlogService_ChangeVote_FullMethodName         = "/BlogService/ChangeVote"
	BlogService_AddComment_FullMethodName         = "/BlogService/AddComment"
	BlogService_UpdateComment_FullMethodName      = "/BlogService/UpdateComment"
	BlogService_DeleteComment_FullMethodName      = "/BlogService/DeleteComment"
)

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogServiceClient interface {
	GetAllBlogs(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetBlogsResponse, error)
	GetBlogById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*BlogResponse, error)
	PostBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*BlogResponse, error)
	UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*BlogResponse, error)
	DeleteBlog(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*BlogResponse, error)
	GetBlogsByAuthorId(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetBlogsResponse, error)
	GetAllVotes(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetAllVotesResponse, error)
	GetVotesCount(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetVotesCountResponse, error)
	AddVote(ctx context.Context, in *AddVoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	ChangeVote(ctx context.Context, in *ChangeVoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) GetAllBlogs(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetBlogsResponse, error) {
	out := new(GetBlogsResponse)
	err := c.cc.Invoke(ctx, BlogService_GetAllBlogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GetBlogById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*BlogResponse, error) {
	out := new(BlogResponse)
	err := c.cc.Invoke(ctx, BlogService_GetBlogById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) PostBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*BlogResponse, error) {
	out := new(BlogResponse)
	err := c.cc.Invoke(ctx, BlogService_PostBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*BlogResponse, error) {
	out := new(BlogResponse)
	err := c.cc.Invoke(ctx, BlogService_UpdateBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteBlog(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*BlogResponse, error) {
	out := new(BlogResponse)
	err := c.cc.Invoke(ctx, BlogService_DeleteBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GetBlogsByAuthorId(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetBlogsResponse, error) {
	out := new(GetBlogsResponse)
	err := c.cc.Invoke(ctx, BlogService_GetBlogsByAuthorId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GetAllVotes(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetAllVotesResponse, error) {
	out := new(GetAllVotesResponse)
	err := c.cc.Invoke(ctx, BlogService_GetAllVotes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GetVotesCount(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetVotesCountResponse, error) {
	out := new(GetVotesCountResponse)
	err := c.cc.Invoke(ctx, BlogService_GetVotesCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) AddVote(ctx context.Context, in *AddVoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, BlogService_AddVote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ChangeVote(ctx context.Context, in *ChangeVoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, BlogService_ChangeVote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, BlogService_AddComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, BlogService_UpdateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, BlogService_DeleteComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
// All implementations must embed UnimplementedBlogServiceServer
// for forward compatibility
type BlogServiceServer interface {
	GetAllBlogs(context.Context, *GetAllRequest) (*GetBlogsResponse, error)
	GetBlogById(context.Context, *GetByIdRequest) (*BlogResponse, error)
	PostBlog(context.Context, *CreateBlogRequest) (*BlogResponse, error)
	UpdateBlog(context.Context, *UpdateBlogRequest) (*BlogResponse, error)
	DeleteBlog(context.Context, *GetByIdRequest) (*BlogResponse, error)
	GetBlogsByAuthorId(context.Context, *GetByIdRequest) (*GetBlogsResponse, error)
	GetAllVotes(context.Context, *GetByIdRequest) (*GetAllVotesResponse, error)
	GetVotesCount(context.Context, *GetByIdRequest) (*GetVotesCountResponse, error)
	AddVote(context.Context, *AddVoteRequest) (*VoteResponse, error)
	ChangeVote(context.Context, *ChangeVoteRequest) (*VoteResponse, error)
	AddComment(context.Context, *AddCommentRequest) (*CommentResponse, error)
	UpdateComment(context.Context, *UpdateCommentRequest) (*CommentResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*CommentResponse, error)
	mustEmbedUnimplementedBlogServiceServer()
}

// UnimplementedBlogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (UnimplementedBlogServiceServer) GetAllBlogs(context.Context, *GetAllRequest) (*GetBlogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBlogs not implemented")
}
func (UnimplementedBlogServiceServer) GetBlogById(context.Context, *GetByIdRequest) (*BlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlogById not implemented")
}
func (UnimplementedBlogServiceServer) PostBlog(context.Context, *CreateBlogRequest) (*BlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostBlog not implemented")
}
func (UnimplementedBlogServiceServer) UpdateBlog(context.Context, *UpdateBlogRequest) (*BlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (UnimplementedBlogServiceServer) DeleteBlog(context.Context, *GetByIdRequest) (*BlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
func (UnimplementedBlogServiceServer) GetBlogsByAuthorId(context.Context, *GetByIdRequest) (*GetBlogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlogsByAuthorId not implemented")
}
func (UnimplementedBlogServiceServer) GetAllVotes(context.Context, *GetByIdRequest) (*GetAllVotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllVotes not implemented")
}
func (UnimplementedBlogServiceServer) GetVotesCount(context.Context, *GetByIdRequest) (*GetVotesCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVotesCount not implemented")
}
func (UnimplementedBlogServiceServer) AddVote(context.Context, *AddVoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVote not implemented")
}
func (UnimplementedBlogServiceServer) ChangeVote(context.Context, *ChangeVoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeVote not implemented")
}
func (UnimplementedBlogServiceServer) AddComment(context.Context, *AddCommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedBlogServiceServer) UpdateComment(context.Context, *UpdateCommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedBlogServiceServer) DeleteComment(context.Context, *DeleteCommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedBlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {}

// UnsafeBlogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServiceServer will
// result in compilation errors.
type UnsafeBlogServiceServer interface {
	mustEmbedUnimplementedBlogServiceServer()
}

func RegisterBlogServiceServer(s grpc.ServiceRegistrar, srv BlogServiceServer) {
	s.RegisterService(&BlogService_ServiceDesc, srv)
}

func _BlogService_GetAllBlogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetAllBlogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_GetAllBlogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetAllBlogs(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GetBlogById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetBlogById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_GetBlogById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetBlogById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_PostBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).PostBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_PostBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).PostBlog(ctx, req.(*CreateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_UpdateBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateBlog(ctx, req.(*UpdateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_DeleteBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteBlog(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GetBlogsByAuthorId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetBlogsByAuthorId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_GetBlogsByAuthorId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetBlogsByAuthorId(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GetAllVotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetAllVotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_GetAllVotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetAllVotes(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GetVotesCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetVotesCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_GetVotesCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetVotesCount(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_AddVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).AddVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_AddVote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).AddVote(ctx, req.(*AddVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ChangeVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ChangeVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_ChangeVote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ChangeVote(ctx, req.(*ChangeVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).AddComment(ctx, req.(*AddCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateComment(ctx, req.(*UpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogService_ServiceDesc is the grpc.ServiceDesc for BlogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllBlogs",
			Handler:    _BlogService_GetAllBlogs_Handler,
		},
		{
			MethodName: "GetBlogById",
			Handler:    _BlogService_GetBlogById_Handler,
		},
		{
			MethodName: "PostBlog",
			Handler:    _BlogService_PostBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _BlogService_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _BlogService_DeleteBlog_Handler,
		},
		{
			MethodName: "GetBlogsByAuthorId",
			Handler:    _BlogService_GetBlogsByAuthorId_Handler,
		},
		{
			MethodName: "GetAllVotes",
			Handler:    _BlogService_GetAllVotes_Handler,
		},
		{
			MethodName: "GetVotesCount",
			Handler:    _BlogService_GetVotesCount_Handler,
		},
		{
			MethodName: "AddVote",
			Handler:    _BlogService_AddVote_Handler,
		},
		{
			MethodName: "ChangeVote",
			Handler:    _BlogService_ChangeVote_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _BlogService_AddComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _BlogService_UpdateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _BlogService_DeleteComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog_microservice.proto",
}
