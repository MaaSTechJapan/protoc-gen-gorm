// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/contacts/contacts.proto

package contacts

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "github.com/infobloxopen/protoc-gen-gorm/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Contact struct {
	Id           uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FirstName    string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	MiddleName   string `protobuf:"bytes,3,opt,name=middle_name,json=middleName" json:"middle_name,omitempty"`
	LastName     string `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	EmailAddress string `protobuf:"bytes,5,opt,name=email_address,json=emailAddress" json:"email_address,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Contact) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Contact) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Contact) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *Contact) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Contact) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

// eventually this will be replaced with a Page that uses the paging info
type ContactPage struct {
	Results []*Contact `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *ContactPage) Reset()                    { *m = ContactPage{} }
func (m *ContactPage) String() string            { return proto.CompactTextString(m) }
func (*ContactPage) ProtoMessage()               {}
func (*ContactPage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ContactPage) GetResults() []*Contact {
	if m != nil {
		return m.Results
	}
	return nil
}

// eventually this will be replaced with a standard search request
type SearchRequest struct {
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *SearchRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

type GetRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SMSRequest struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *SMSRequest) Reset()                    { *m = SMSRequest{} }
func (m *SMSRequest) String() string            { return proto.CompactTextString(m) }
func (*SMSRequest) ProtoMessage()               {}
func (*SMSRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *SMSRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SMSRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Contact)(nil), "api.contacts.Contact")
	proto.RegisterType((*ContactPage)(nil), "api.contacts.ContactPage")
	proto.RegisterType((*SearchRequest)(nil), "api.contacts.SearchRequest")
	proto.RegisterType((*GetRequest)(nil), "api.contacts.GetRequest")
	proto.RegisterType((*SMSRequest)(nil), "api.contacts.SMSRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Contacts service

type ContactsClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*ContactPage, error)
	Create(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*Contact, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Contact, error)
	Update(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*Contact, error)
	Delete(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	SendSMS(ctx context.Context, in *SMSRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type contactsClient struct {
	cc *grpc.ClientConn
}

func NewContactsClient(cc *grpc.ClientConn) ContactsClient {
	return &contactsClient{cc}
}

func (c *contactsClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*ContactPage, error) {
	out := new(ContactPage)
	err := grpc.Invoke(ctx, "/api.contacts.Contacts/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) Create(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*Contact, error) {
	out := new(Contact)
	err := grpc.Invoke(ctx, "/api.contacts.Contacts/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Contact, error) {
	out := new(Contact)
	err := grpc.Invoke(ctx, "/api.contacts.Contacts/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) Update(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*Contact, error) {
	out := new(Contact)
	err := grpc.Invoke(ctx, "/api.contacts.Contacts/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) Delete(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/api.contacts.Contacts/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) SendSMS(ctx context.Context, in *SMSRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/api.contacts.Contacts/SendSMS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Contacts service

type ContactsServer interface {
	Search(context.Context, *SearchRequest) (*ContactPage, error)
	Create(context.Context, *Contact) (*Contact, error)
	Get(context.Context, *GetRequest) (*Contact, error)
	Update(context.Context, *Contact) (*Contact, error)
	Delete(context.Context, *GetRequest) (*google_protobuf.Empty, error)
	SendSMS(context.Context, *SMSRequest) (*google_protobuf.Empty, error)
}

func RegisterContactsServer(s *grpc.Server, srv ContactsServer) {
	s.RegisterService(&_Contacts_serviceDesc, srv)
}

func _Contacts_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.contacts.Contacts/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.contacts.Contacts/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).Create(ctx, req.(*Contact))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.contacts.Contacts/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.contacts.Contacts/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).Update(ctx, req.(*Contact))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.contacts.Contacts/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).Delete(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_SendSMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).SendSMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.contacts.Contacts/SendSMS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).SendSMS(ctx, req.(*SMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contacts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.contacts.Contacts",
	HandlerType: (*ContactsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Contacts_Search_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Contacts_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Contacts_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Contacts_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Contacts_Delete_Handler,
		},
		{
			MethodName: "SendSMS",
			Handler:    _Contacts_SendSMS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/contacts/contacts.proto",
}

func init() { proto.RegisterFile("example/contacts/contacts.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x96, 0x9d, 0xfc, 0x9c, 0x66, 0xd2, 0x56, 0x3f, 0x16, 0x28, 0xae, 0xdb, 0xaa, 0x95, 0x4f,
	0x51, 0x45, 0x6d, 0xa9, 0x48, 0x1c, 0x52, 0x09, 0xa9, 0x0d, 0xa8, 0x17, 0x5a, 0x81, 0x2d, 0x2e,
	0xbd, 0x94, 0x4d, 0x3c, 0x71, 0x57, 0xb2, 0xbd, 0xc6, 0xbb, 0x81, 0x56, 0x08, 0x09, 0xf1, 0x0a,
	0xbc, 0x07, 0x77, 0x7a, 0xe2, 0x1d, 0x78, 0x05, 0x2e, 0xbc, 0x05, 0xf2, 0xda, 0xce, 0x1f, 0x93,
	0x20, 0xb8, 0xed, 0xce, 0xf7, 0xcd, 0x37, 0xdf, 0xcc, 0x7a, 0x0c, 0xbb, 0x78, 0x4d, 0xe3, 0x34,
	0x42, 0x77, 0xc8, 0x13, 0x49, 0x87, 0x52, 0x4c, 0x0e, 0x4e, 0x9a, 0x71, 0xc9, 0xc9, 0x2a, 0x4d,
	0x99, 0x53, 0xc5, 0xac, 0xad, 0x90, 0xf3, 0x30, 0x42, 0x57, 0x61, 0x83, 0xf1, 0xc8, 0xc5, 0x38,
	0x95, 0x37, 0x05, 0xd5, 0xda, 0x2e, 0x41, 0x9a, 0x32, 0x97, 0x26, 0x09, 0x97, 0x54, 0x32, 0x9e,
	0x94, 0x42, 0xd6, 0x51, 0xc8, 0xe4, 0xd5, 0x78, 0xe0, 0x0c, 0x79, 0xec, 0x46, 0x37, 0x23, 0x59,
	0x68, 0x0c, 0x0f, 0x42, 0x4c, 0x0e, 0xde, 0xd2, 0x88, 0x05, 0x54, 0xa2, 0xfb, 0xdb, 0xa1, 0x4c,
	0x7e, 0x38, 0x43, 0x16, 0xef, 0x68, 0x18, 0x62, 0xe6, 0xf2, 0x54, 0xc9, 0x2f, 0x28, 0xd5, 0x9b,
	0x29, 0xc5, 0x92, 0x11, 0x1f, 0x44, 0xfc, 0x9a, 0xa7, 0x98, 0xcc, 0x96, 0x0c, 0x79, 0x16, 0x4f,
	0x24, 0xf2, 0x4b, 0x91, 0x6b, 0x7f, 0xd1, 0xa0, 0xd5, 0x2f, 0xda, 0x25, 0xeb, 0xa0, 0xb3, 0xc0,
	0xd4, 0xf6, 0xb4, 0x6e, 0xd3, 0xd3, 0x59, 0x40, 0x76, 0x00, 0x46, 0x2c, 0x13, 0xf2, 0x32, 0xa1,
	0x31, 0x9a, 0xfa, 0x9e, 0xd6, 0x6d, 0x7b, 0x6d, 0x15, 0x39, 0xa7, 0x31, 0x92, 0x5d, 0xe8, 0xc4,
	0x2c, 0x08, 0x22, 0x2c, 0xf0, 0x86, 0xc2, 0xa1, 0x08, 0x29, 0xc2, 0x16, 0xb4, 0x23, 0x5a, 0xa5,
	0x37, 0x15, 0xbc, 0x92, 0x07, 0x14, 0xe8, 0xc0, 0x1a, 0xc6, 0x94, 0x45, 0x97, 0x34, 0x08, 0x32,
	0x14, 0xc2, 0xfc, 0x2f, 0x27, 0x9c, 0xb4, 0x6f, 0x7f, 0x7e, 0x6b, 0x34, 0x33, 0xfd, 0xb5, 0xe6,
	0xad, 0x2a, 0xfc, 0xb8, 0x80, 0x7b, 0xc6, 0xed, 0xd7, 0x4d, 0x7d, 0x45, 0xb3, 0x9f, 0x40, 0xa7,
	0xf4, 0xfb, 0x82, 0x86, 0x48, 0x5c, 0x68, 0x65, 0x28, 0xc6, 0x91, 0x14, 0xa6, 0xb6, 0xd7, 0xe8,
	0x76, 0x0e, 0xef, 0x3b, 0xb3, 0x2f, 0xe8, 0x94, 0x5c, 0xaf, 0x62, 0xd9, 0x0e, 0xac, 0xf9, 0x48,
	0xb3, 0xe1, 0x95, 0x87, 0x6f, 0xc6, 0x28, 0x64, 0xad, 0x4b, 0xad, 0xd6, 0xa5, 0xbd, 0x0d, 0x70,
	0x8a, 0xb2, 0x22, 0xd7, 0x46, 0x64, 0x3f, 0x06, 0xf0, 0xcf, 0xfc, 0x25, 0x28, 0x31, 0xa1, 0x15,
	0xa3, 0x10, 0x34, 0xac, 0xa6, 0x57, 0x5d, 0x0f, 0x3f, 0x36, 0x61, 0xa5, 0xb4, 0x26, 0x88, 0x0f,
	0x46, 0x61, 0x89, 0x6c, 0xcd, 0x9b, 0x9f, 0x33, 0x6a, 0x6d, 0x2e, 0xec, 0x2c, 0x9f, 0x82, 0x7d,
	0xe7, 0xd3, 0xf7, 0x1f, 0x9f, 0xf5, 0x0e, 0x69, 0x4f, 0x3e, 0x67, 0xf2, 0x1c, 0x8c, 0x7e, 0x86,
	0x54, 0x22, 0x59, 0x3c, 0x11, 0x6b, 0x71, 0xd8, 0xbe, 0xa7, 0xa4, 0xd6, 0xed, 0xa9, 0x54, 0x4f,
	0xdb, 0x27, 0xe7, 0xd0, 0x38, 0x45, 0x49, 0xcc, 0xf9, 0x9c, 0xe9, 0x60, 0x96, 0xa9, 0x6d, 0x28,
	0xb5, 0xff, 0xc9, 0xfa, 0x74, 0xe1, 0xde, 0xb3, 0xe0, 0x03, 0x79, 0x09, 0xc6, 0xab, 0x34, 0xf8,
	0x77, 0x77, 0x9b, 0x4a, 0xef, 0xae, 0x55, 0xd3, 0xcb, 0x2d, 0x7a, 0x60, 0x3c, 0xc5, 0x08, 0x25,
	0xfe, 0xc1, 0xe5, 0x86, 0x53, 0xec, 0xac, 0x53, 0x2d, 0xb4, 0xf3, 0x2c, 0x5f, 0xe8, 0xca, 0xe6,
	0x7e, 0xdd, 0xe6, 0x05, 0xb4, 0x7c, 0x4c, 0x02, 0xff, 0xcc, 0xaf, 0x8b, 0x4e, 0x5f, 0x7d, 0xa9,
	0xe8, 0x8e, 0x12, 0x7d, 0x60, 0x93, 0x79, 0x51, 0x57, 0xc4, 0xf9, 0x48, 0x4f, 0xfa, 0x17, 0xc7,
	0x7f, 0xbb, 0xb7, 0xf5, 0x9f, 0xd6, 0x51, 0x75, 0x18, 0x18, 0x8a, 0xfa, 0xe8, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x1b, 0x0c, 0xf9, 0xea, 0xd8, 0x04, 0x00, 0x00,
}