// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Person.proto

//包名，通过protoc生成时go文件时

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

//手机类型
//枚举类型第一个字段必须为0
type PhoneType int32

const (
	PhoneType_HOME PhoneType = 0
	PhoneType_WORK PhoneType = 1
)

var PhoneType_name = map[int32]string{
	0: "HOME",
	1: "WORK",
}

var PhoneType_value = map[string]int32{
	"HOME": 0,
	"WORK": 1,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}

func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{0}
}

//手机
type Phone struct {
	Type                 PhoneType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.PhoneType" json:"type,omitempty"`
	Number               string    `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Phone) Reset()         { *m = Phone{} }
func (m *Phone) String() string { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()    {}
func (*Phone) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{0}
}

func (m *Phone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phone.Unmarshal(m, b)
}
func (m *Phone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phone.Marshal(b, m, deterministic)
}
func (m *Phone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phone.Merge(m, src)
}
func (m *Phone) XXX_Size() int {
	return xxx_messageInfo_Phone.Size(m)
}
func (m *Phone) XXX_DiscardUnknown() {
	xxx_messageInfo_Phone.DiscardUnknown(m)
}

var xxx_messageInfo_Phone proto.InternalMessageInfo

func (m *Phone) GetType() PhoneType {
	if m != nil {
		return m.Type
	}
	return PhoneType_HOME
}

func (m *Phone) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

//人
type Person struct {
	//后面的数字表示标识号
	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//repeated表示可重复
	//可以有多个手机
	Phones               []*Phone `protobuf:"bytes,3,rep,name=phones,proto3" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{1}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetPhones() []*Phone {
	if m != nil {
		return m.Phones
	}
	return nil
}

//联系簿
type ContactBook struct {
	Persons              []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ContactBook) Reset()         { *m = ContactBook{} }
func (m *ContactBook) String() string { return proto.CompactTextString(m) }
func (*ContactBook) ProtoMessage()    {}
func (*ContactBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{2}
}

func (m *ContactBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactBook.Unmarshal(m, b)
}
func (m *ContactBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactBook.Marshal(b, m, deterministic)
}
func (m *ContactBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactBook.Merge(m, src)
}
func (m *ContactBook) XXX_Size() int {
	return xxx_messageInfo_ContactBook.Size(m)
}
func (m *ContactBook) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactBook.DiscardUnknown(m)
}

var xxx_messageInfo_ContactBook proto.InternalMessageInfo

func (m *ContactBook) GetPersons() []*Person {
	if m != nil {
		return m.Persons
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.PhoneType", PhoneType_name, PhoneType_value)
	proto.RegisterType((*Phone)(nil), "pb.Phone")
	proto.RegisterType((*Person)(nil), "pb.Person")
	proto.RegisterType((*ContactBook)(nil), "pb.ContactBook")
}

func init() { proto.RegisterFile("Person.proto", fileDescriptor_841ab6396175eaf3) }

var fileDescriptor_841ab6396175eaf3 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x4d, 0xb6, 0x1b, 0xed, 0xac, 0x2e, 0xcb, 0x1c, 0x24, 0x37, 0xbb, 0xc5, 0x43, 0xf1,
	0xd0, 0x43, 0xfb, 0x06, 0x15, 0x41, 0x10, 0x69, 0x09, 0x82, 0xe7, 0xc6, 0x06, 0x2c, 0xd2, 0x24,
	0xb4, 0xf1, 0xd0, 0xb7, 0x97, 0x8e, 0xdd, 0xde, 0xfe, 0xe4, 0xcb, 0xff, 0x65, 0x06, 0x6e, 0x1b,
	0x33, 0x4e, 0xce, 0xe6, 0x7e, 0x74, 0xc1, 0x21, 0xf7, 0x3a, 0xad, 0x60, 0xdf, 0x7c, 0x3b, 0x6b,
	0xf0, 0x0c, 0x51, 0x98, 0xbd, 0x91, 0x2c, 0x61, 0xd9, 0xb1, 0xb8, 0xcb, 0xbd, 0xce, 0x09, 0x7c,
	0xcc, 0xde, 0x28, 0x42, 0x78, 0x0f, 0xc2, 0xfe, 0x0e, 0xda, 0x8c, 0x92, 0x27, 0x2c, 0x8b, 0xd5,
	0x7a, 0x4a, 0x6b, 0x10, 0xff, 0x5e, 0x3c, 0x02, 0xef, 0x3b, 0x52, 0xec, 0x15, 0xef, 0x3b, 0x44,
	0x88, 0x6c, 0x3b, 0x98, 0xf5, 0x3d, 0x65, 0x3c, 0x83, 0xf0, 0x8b, 0x78, 0x92, 0xbb, 0x64, 0x97,
	0x1d, 0x8a, 0x78, 0xfb, 0x4a, 0xad, 0x20, 0x2d, 0xe1, 0xf0, 0xec, 0x6c, 0x68, 0xbf, 0x42, 0xe5,
	0xdc, 0x0f, 0x3e, 0xc2, 0xb5, 0x27, 0xff, 0x24, 0x19, 0x55, 0x80, 0x2a, 0x74, 0xa5, 0x2e, 0xe8,
	0xe9, 0x01, 0xe2, 0x6d, 0x60, 0xbc, 0x81, 0xe8, 0xb5, 0x7e, 0x7f, 0x39, 0x5d, 0x2d, 0xe9, 0xb3,
	0x56, 0x6f, 0x27, 0xa6, 0x05, 0x6d, 0x5d, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x6d, 0xb8,
	0xf4, 0x05, 0x01, 0x00, 0x00,
}
