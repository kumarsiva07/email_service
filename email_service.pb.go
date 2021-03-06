// Code generated by protoc-gen-go.
// source: email_service.proto
// DO NOT EDIT!

/*
Package email_service is a generated protocol buffer package.

It is generated from these files:
	email_service.proto

It has these top-level messages:
	Email
	Attachment
	EmailResponse
*/
package email_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Email struct {
	From          string            `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To            []string          `protobuf:"bytes,2,rep,name=to" json:"to,omitempty"`
	Subject       string            `protobuf:"bytes,3,opt,name=subject" json:"subject,omitempty"`
	PlainText     string            `protobuf:"bytes,4,opt,name=plain_text,json=plainText" json:"plain_text,omitempty"`
	HtmlAlternate string            `protobuf:"bytes,5,opt,name=html_alternate,json=htmlAlternate" json:"html_alternate,omitempty"`
	Attachments   []*Attachment     `protobuf:"bytes,6,rep,name=attachments" json:"attachments,omitempty"`
	Headers       map[string]string `protobuf:"bytes,7,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Email) Reset()                    { *m = Email{} }
func (m *Email) String() string            { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()               {}
func (*Email) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Email) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Email) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Email) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Email) GetPlainText() string {
	if m != nil {
		return m.PlainText
	}
	return ""
}

func (m *Email) GetHtmlAlternate() string {
	if m != nil {
		return m.HtmlAlternate
	}
	return ""
}

func (m *Email) GetAttachments() []*Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Email) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

type Attachment struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Body     []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Attachment) Reset()                    { *m = Attachment{} }
func (m *Attachment) String() string            { return proto.CompactTextString(m) }
func (*Attachment) ProtoMessage()               {}
func (*Attachment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Attachment) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Attachment) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type EmailResponse struct {
}

func (m *EmailResponse) Reset()                    { *m = EmailResponse{} }
func (m *EmailResponse) String() string            { return proto.CompactTextString(m) }
func (*EmailResponse) ProtoMessage()               {}
func (*EmailResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Email)(nil), "email_service.Email")
	proto.RegisterType((*Attachment)(nil), "email_service.Attachment")
	proto.RegisterType((*EmailResponse)(nil), "email_service.EmailResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EmailService service

type EmailServiceClient interface {
	Send(ctx context.Context, in *Email, opts ...grpc.CallOption) (*EmailResponse, error)
}

type emailServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmailServiceClient(cc *grpc.ClientConn) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) Send(ctx context.Context, in *Email, opts ...grpc.CallOption) (*EmailResponse, error) {
	out := new(EmailResponse)
	err := grpc.Invoke(ctx, "/email_service.EmailService/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EmailService service

type EmailServiceServer interface {
	Send(context.Context, *Email) (*EmailResponse, error)
}

func RegisterEmailServiceServer(s *grpc.Server, srv EmailServiceServer) {
	s.RegisterService(&_EmailService_serviceDesc, srv)
}

func _EmailService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email_service.EmailService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Send(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "email_service.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _EmailService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email_service.proto",
}

func init() { proto.RegisterFile("email_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6b, 0xf2, 0x40,
	0x10, 0xc6, 0x5f, 0x13, 0xff, 0xbc, 0x8e, 0x7f, 0xde, 0x97, 0xad, 0x87, 0xad, 0xb4, 0x54, 0x03,
	0x05, 0x4f, 0x11, 0xec, 0xa5, 0xa8, 0x17, 0x0b, 0x42, 0x0f, 0x3d, 0xc5, 0x9e, 0x7a, 0x91, 0x4d,
	0x1c, 0x9b, 0xb4, 0x9b, 0x5d, 0x49, 0x46, 0xd1, 0x8f, 0xd7, 0x6f, 0x56, 0xb2, 0x1a, 0xab, 0xc5,
	0xdb, 0xcc, 0x6f, 0x9f, 0x67, 0x98, 0x67, 0x58, 0xb8, 0xc2, 0x58, 0x44, 0x72, 0x9e, 0x62, 0xb2,
	0x89, 0x02, 0x74, 0x57, 0x89, 0x26, 0xcd, 0x1a, 0x67, 0xd0, 0xf9, 0xb2, 0xa0, 0x34, 0xcd, 0x08,
	0x63, 0x50, 0x5c, 0x26, 0x3a, 0xe6, 0x85, 0x4e, 0xa1, 0x57, 0xf5, 0x4c, 0xcd, 0x9a, 0x60, 0x91,
	0xe6, 0x56, 0xc7, 0xee, 0x55, 0x3d, 0x8b, 0x34, 0xe3, 0x50, 0x49, 0xd7, 0xfe, 0x07, 0x06, 0xc4,
	0x6d, 0x23, 0xcb, 0x5b, 0x76, 0x0b, 0xb0, 0x92, 0x22, 0x52, 0x73, 0xc2, 0x2d, 0xf1, 0xa2, 0x79,
	0xac, 0x1a, 0xf2, 0x8a, 0x5b, 0x62, 0xf7, 0xd0, 0x0c, 0x29, 0x96, 0x73, 0x21, 0x09, 0x13, 0x25,
	0x08, 0x79, 0xc9, 0x48, 0x1a, 0x19, 0x9d, 0xe4, 0x90, 0x8d, 0xa0, 0x26, 0x88, 0x44, 0x10, 0xc6,
	0xa8, 0x28, 0xe5, 0xe5, 0x8e, 0xdd, 0xab, 0x0d, 0xae, 0xdd, 0xf3, 0x1c, 0x93, 0xa3, 0xc2, 0x3b,
	0x55, 0xb3, 0x11, 0x54, 0x42, 0x14, 0x0b, 0x4c, 0x52, 0x5e, 0x31, 0xc6, 0xee, 0x2f, 0xa3, 0xc9,
	0xe9, 0x3e, 0xef, 0x35, 0x53, 0x45, 0xc9, 0xce, 0xcb, 0x1d, 0xed, 0x21, 0xd4, 0x4f, 0x1f, 0xd8,
	0x7f, 0xb0, 0x3f, 0x71, 0x77, 0x38, 0x46, 0x56, 0xb2, 0x16, 0x94, 0x36, 0x42, 0xae, 0x91, 0x5b,
	0x86, 0xed, 0x9b, 0xa1, 0xf5, 0x58, 0x70, 0xc6, 0x00, 0x3f, 0x3b, 0xb1, 0x36, 0xfc, 0x5d, 0x46,
	0x12, 0x95, 0x88, 0xf1, 0x60, 0x3f, 0xf6, 0xd9, 0x8d, 0x7d, 0xbd, 0xd8, 0x99, 0x11, 0x75, 0xcf,
	0xd4, 0xce, 0x3f, 0x68, 0x98, 0xc5, 0x3c, 0x4c, 0x57, 0x5a, 0xa5, 0x38, 0x78, 0x81, 0xba, 0x01,
	0xb3, 0xfd, 0xda, 0x6c, 0x0c, 0xc5, 0x19, 0xaa, 0x05, 0x6b, 0x5d, 0x8a, 0xd3, 0xbe, 0xb9, 0x44,
	0xf3, 0x59, 0xce, 0x9f, 0xa7, 0xee, 0xdb, 0xdd, 0x7b, 0x44, 0xe1, 0xda, 0x77, 0x03, 0x1d, 0xf7,
	0x65, 0x24, 0x31, 0xd2, 0xfd, 0x33, 0x8b, 0x5f, 0x36, 0x3f, 0xe3, 0xe1, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x69, 0x42, 0x94, 0x4f, 0x30, 0x02, 0x00, 0x00,
}
