// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: book.proto

package book_service

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

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	Create(ctx context.Context, in *CreateBook, opts ...grpc.CallOption) (*BookId, error)
	GetAll(ctx context.Context, in *GetAllBookRequest, opts ...grpc.CallOption) (*GetAllBookResponse, error)
	GetById(ctx context.Context, in *BookId, opts ...grpc.CallOption) (*GetBookByIdResponse, error)
	Update(ctx context.Context, in *UpdateBook, opts ...grpc.CallOption) (*MsgRespons, error)
	Delete(ctx context.Context, in *BookId, opts ...grpc.CallOption) (*MsgRespons, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) Create(ctx context.Context, in *CreateBook, opts ...grpc.CallOption) (*BookId, error) {
	out := new(BookId)
	err := c.cc.Invoke(ctx, "/genproto.BookService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetAll(ctx context.Context, in *GetAllBookRequest, opts ...grpc.CallOption) (*GetAllBookResponse, error) {
	out := new(GetAllBookResponse)
	err := c.cc.Invoke(ctx, "/genproto.BookService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetById(ctx context.Context, in *BookId, opts ...grpc.CallOption) (*GetBookByIdResponse, error) {
	out := new(GetBookByIdResponse)
	err := c.cc.Invoke(ctx, "/genproto.BookService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Update(ctx context.Context, in *UpdateBook, opts ...grpc.CallOption) (*MsgRespons, error) {
	out := new(MsgRespons)
	err := c.cc.Invoke(ctx, "/genproto.BookService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Delete(ctx context.Context, in *BookId, opts ...grpc.CallOption) (*MsgRespons, error) {
	out := new(MsgRespons)
	err := c.cc.Invoke(ctx, "/genproto.BookService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility
type BookServiceServer interface {
	Create(context.Context, *CreateBook) (*BookId, error)
	GetAll(context.Context, *GetAllBookRequest) (*GetAllBookResponse, error)
	GetById(context.Context, *BookId) (*GetBookByIdResponse, error)
	Update(context.Context, *UpdateBook) (*MsgRespons, error)
	Delete(context.Context, *BookId) (*MsgRespons, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (UnimplementedBookServiceServer) Create(context.Context, *CreateBook) (*BookId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBookServiceServer) GetAll(context.Context, *GetAllBookRequest) (*GetAllBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedBookServiceServer) GetById(context.Context, *BookId) (*GetBookByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedBookServiceServer) Update(context.Context, *UpdateBook) (*MsgRespons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBookServiceServer) Delete(context.Context, *BookId) (*MsgRespons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.BookService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Create(ctx, req.(*CreateBook))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.BookService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetAll(ctx, req.(*GetAllBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.BookService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetById(ctx, req.(*BookId))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.BookService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Update(ctx, req.(*UpdateBook))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.BookService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Delete(ctx, req.(*BookId))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BookService_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _BookService_GetAll_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _BookService_GetById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BookService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BookService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}
