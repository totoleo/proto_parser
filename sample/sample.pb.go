// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sample.proto

package sample

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

type RecallRequest struct {
	Uid                  string            `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Dtu                  int32             `protobuf:"varint,3,opt,name=Dtu,proto3" json:"Dtu,omitempty"`
	Cate                 int32             `protobuf:"varint,4,opt,name=Cate,proto3" json:"Cate,omitempty"`
	RecType              int32             `protobuf:"varint,5,opt,name=RecType,proto3" json:"RecType,omitempty"`
	DeviceCode           string            `protobuf:"bytes,6,opt,name=DeviceCode,proto3" json:"DeviceCode,omitempty"`
	ContentType          []int32           `protobuf:"varint,7,rep,packed,name=ContentType,proto3" json:"ContentType,omitempty"`
	ExcludeDocs          []int32           `protobuf:"varint,8,rep,packed,name=ExcludeDocs,proto3" json:"ExcludeDocs,omitempty"`
	SimilarDocId         int32             `protobuf:"varint,9,opt,name=SimilarDocId,proto3" json:"SimilarDocId,omitempty"`
	IsWifi               bool              `protobuf:"varint,10,opt,name=IsWifi,proto3" json:"IsWifi,omitempty"`
	IsVideoPage          bool              `protobuf:"varint,11,opt,name=IsVideoPage,proto3" json:"IsVideoPage,omitempty"`
	RegTime              int64             `protobuf:"varint,12,opt,name=RegTime,proto3" json:"RegTime,omitempty"`
	ActTime              int64             `protobuf:"varint,13,opt,name=ActTime,proto3" json:"ActTime,omitempty"`
	TkId                 string            `protobuf:"bytes,14,opt,name=TkId,proto3" json:"TkId,omitempty"`
	Op                   int32             `protobuf:"varint,15,opt,name=Op,proto3" json:"Op,omitempty"`
	Page                 int32             `protobuf:"varint,16,opt,name=Page,proto3" json:"Page,omitempty"`
	Region               int32             `protobuf:"varint,17,opt,name=Region,proto3" json:"Region,omitempty"`
	ReqFrom              string            `protobuf:"bytes,19,opt,name=ReqFrom,proto3" json:"ReqFrom,omitempty"`
	ClientVersion        int32             `protobuf:"varint,20,opt,name=ClientVersion,proto3" json:"ClientVersion,omitempty"`
	AutoplayFilter       int32             `protobuf:"varint,21,opt,name=AutoplayFilter,proto3" json:"AutoplayFilter,omitempty"`
	RegionFrequency      int32             `protobuf:"varint,22,opt,name=RegionFrequency,proto3" json:"RegionFrequency,omitempty"`
	Lon                  float32           `protobuf:"fixed32,23,opt,name=Lon,proto3" json:"Lon,omitempty"`
	Lat                  float32           `protobuf:"fixed32,24,opt,name=Lat,proto3" json:"Lat,omitempty"`
	ExcludeDocsV2        []int32           `protobuf:"varint,25,rep,packed,name=ExcludeDocsV2,proto3" json:"ExcludeDocsV2,omitempty"`
	Extend               map[string]string `protobuf:"bytes,26,rep,name=Extend,proto3" json:"Extend,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TUid                 string            `protobuf:"bytes,27,opt,name=TUid,proto3" json:"TUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RecallRequest) Reset()         { *m = RecallRequest{} }
func (m *RecallRequest) String() string { return proto.CompactTextString(m) }
func (*RecallRequest) ProtoMessage()    {}
func (*RecallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{0}
}

func (m *RecallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecallRequest.Unmarshal(m, b)
}
func (m *RecallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecallRequest.Marshal(b, m, deterministic)
}
func (m *RecallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecallRequest.Merge(m, src)
}
func (m *RecallRequest) XXX_Size() int {
	return xxx_messageInfo_RecallRequest.Size(m)
}
func (m *RecallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecallRequest proto.InternalMessageInfo

func (m *RecallRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RecallRequest) GetDtu() int32 {
	if m != nil {
		return m.Dtu
	}
	return 0
}

func (m *RecallRequest) GetCate() int32 {
	if m != nil {
		return m.Cate
	}
	return 0
}

func (m *RecallRequest) GetRecType() int32 {
	if m != nil {
		return m.RecType
	}
	return 0
}

func (m *RecallRequest) GetDeviceCode() string {
	if m != nil {
		return m.DeviceCode
	}
	return ""
}

func (m *RecallRequest) GetContentType() []int32 {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *RecallRequest) GetExcludeDocs() []int32 {
	if m != nil {
		return m.ExcludeDocs
	}
	return nil
}

func (m *RecallRequest) GetSimilarDocId() int32 {
	if m != nil {
		return m.SimilarDocId
	}
	return 0
}

func (m *RecallRequest) GetIsWifi() bool {
	if m != nil {
		return m.IsWifi
	}
	return false
}

func (m *RecallRequest) GetIsVideoPage() bool {
	if m != nil {
		return m.IsVideoPage
	}
	return false
}

func (m *RecallRequest) GetRegTime() int64 {
	if m != nil {
		return m.RegTime
	}
	return 0
}

func (m *RecallRequest) GetActTime() int64 {
	if m != nil {
		return m.ActTime
	}
	return 0
}

func (m *RecallRequest) GetTkId() string {
	if m != nil {
		return m.TkId
	}
	return ""
}

func (m *RecallRequest) GetOp() int32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *RecallRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *RecallRequest) GetRegion() int32 {
	if m != nil {
		return m.Region
	}
	return 0
}

func (m *RecallRequest) GetReqFrom() string {
	if m != nil {
		return m.ReqFrom
	}
	return ""
}

func (m *RecallRequest) GetClientVersion() int32 {
	if m != nil {
		return m.ClientVersion
	}
	return 0
}

func (m *RecallRequest) GetAutoplayFilter() int32 {
	if m != nil {
		return m.AutoplayFilter
	}
	return 0
}

func (m *RecallRequest) GetRegionFrequency() int32 {
	if m != nil {
		return m.RegionFrequency
	}
	return 0
}

func (m *RecallRequest) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

func (m *RecallRequest) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *RecallRequest) GetExcludeDocsV2() []int32 {
	if m != nil {
		return m.ExcludeDocsV2
	}
	return nil
}

func (m *RecallRequest) GetExtend() map[string]string {
	if m != nil {
		return m.Extend
	}
	return nil
}

func (m *RecallRequest) GetTUid() string {
	if m != nil {
		return m.TUid
	}
	return ""
}

func init() {
	proto.RegisterType((*RecallRequest)(nil), "sample.RecallRequest")
	proto.RegisterMapType((map[string]string)(nil), "sample.RecallRequest.ExtendEntry")
}

func init() { proto.RegisterFile("sample.proto", fileDescriptor_2141552de9bf11d0) }

var fileDescriptor_2141552de9bf11d0 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0x65, 0xbb, 0x71, 0xdb, 0xc9, 0x47, 0xcb, 0x52, 0xc2, 0x50, 0x24, 0x64, 0x2a, 0x84,
	0x7c, 0xca, 0xa1, 0x5c, 0x28, 0xb7, 0x2a, 0x1f, 0x52, 0x24, 0xa4, 0xa2, 0x25, 0x0d, 0x67, 0x63,
	0x0f, 0xd1, 0x2a, 0x8e, 0xd7, 0xb5, 0xd7, 0x55, 0xfd, 0x23, 0xf8, 0xcf, 0x68, 0xc7, 0x0e, 0x72,
	0x72, 0x9b, 0xf7, 0x99, 0x99, 0xdd, 0xf1, 0xec, 0x6b, 0x18, 0x94, 0xd1, 0x2e, 0x4f, 0x69, 0x92,
	0x17, 0xda, 0x68, 0xe1, 0x37, 0xea, 0xe6, 0xaf, 0x0f, 0x43, 0x49, 0x71, 0x94, 0xa6, 0x92, 0x9e,
	0x2a, 0x2a, 0x8d, 0xb8, 0x04, 0xef, 0x51, 0x25, 0xe8, 0x04, 0x4e, 0x78, 0x2e, 0x6d, 0x68, 0xc9,
	0xcc, 0x54, 0xe8, 0x05, 0x4e, 0xd8, 0x93, 0x36, 0x14, 0x02, 0x4e, 0xa6, 0x91, 0x21, 0x3c, 0x61,
	0xc4, 0xb1, 0x40, 0x38, 0x95, 0x14, 0xaf, 0xea, 0x9c, 0xb0, 0xc7, 0x78, 0x2f, 0xc5, 0x07, 0x80,
	0x19, 0x3d, 0xab, 0x98, 0xa6, 0x3a, 0x21, 0xf4, 0xf9, 0xe0, 0x0e, 0x11, 0x01, 0xf4, 0xa7, 0x3a,
	0x33, 0x94, 0x19, 0xee, 0x3e, 0x0d, 0xbc, 0xb0, 0x27, 0xbb, 0xc8, 0x56, 0xcc, 0x5f, 0xe2, 0xb4,
	0x4a, 0x68, 0xa6, 0xe3, 0x12, 0xcf, 0x9a, 0x8a, 0x0e, 0x12, 0x37, 0x30, 0xf8, 0xa9, 0x76, 0x2a,
	0x8d, 0x8a, 0x99, 0x8e, 0x97, 0x09, 0x9e, 0xf3, 0x08, 0x07, 0x4c, 0x8c, 0xc1, 0x5f, 0x96, 0xbf,
	0xd4, 0x1f, 0x85, 0x10, 0x38, 0xe1, 0x99, 0x6c, 0x95, 0x3d, 0x7d, 0x59, 0xae, 0x55, 0x42, 0xfa,
	0x47, 0xb4, 0x21, 0xec, 0x73, 0xb2, 0x8b, 0x9a, 0x6f, 0xdb, 0xac, 0xd4, 0x8e, 0x70, 0x10, 0x38,
	0xa1, 0x27, 0xf7, 0xd2, 0x66, 0xee, 0x63, 0xc3, 0x99, 0x61, 0x93, 0x69, 0xa5, 0xdd, 0xd1, 0x6a,
	0xbb, 0x4c, 0x70, 0xc4, 0xdf, 0xcb, 0xb1, 0x18, 0x81, 0xfb, 0x90, 0xe3, 0x05, 0xcf, 0xe6, 0x3e,
	0xe4, 0xb6, 0x86, 0xaf, 0xbc, 0x6c, 0xf6, 0xc8, 0x77, 0x8d, 0xc1, 0x97, 0xb4, 0x51, 0x3a, 0xc3,
	0x57, 0x4c, 0x5b, 0xd5, 0xcc, 0xf0, 0xb4, 0x28, 0xf4, 0x0e, 0x5f, 0xf3, 0x91, 0x7b, 0x29, 0x3e,
	0xc1, 0x70, 0x9a, 0x2a, 0xca, 0xcc, 0x9a, 0x8a, 0xd2, 0x36, 0x5e, 0x71, 0xe3, 0x21, 0x14, 0x9f,
	0x61, 0x74, 0x5f, 0x19, 0x9d, 0xa7, 0x51, 0xbd, 0x50, 0xa9, 0xa1, 0x02, 0xdf, 0x70, 0xd9, 0x11,
	0x15, 0x21, 0x5c, 0x34, 0x37, 0x2e, 0x0a, 0xeb, 0x88, 0x2c, 0xae, 0x71, 0xcc, 0x85, 0xc7, 0xd8,
	0xfa, 0xe2, 0xbb, 0xce, 0xf0, 0x6d, 0xe0, 0x84, 0xae, 0xb4, 0x21, 0x93, 0xc8, 0x20, 0xb6, 0x24,
	0x32, 0x76, 0xb6, 0xce, 0x33, 0xad, 0x6f, 0xf1, 0x1d, 0xbf, 0xdd, 0x21, 0x14, 0x77, 0xe0, 0xcf,
	0x5f, 0x0c, 0x65, 0x09, 0x5e, 0x07, 0x5e, 0xd8, 0xbf, 0xfd, 0x38, 0x69, 0xcd, 0x7a, 0x60, 0xcd,
	0x49, 0x53, 0x33, 0xcf, 0x4c, 0x51, 0xcb, 0xb6, 0x81, 0xd7, 0x6c, 0xfd, 0xfa, 0xbe, 0x5d, 0xf3,
	0xa3, 0x4a, 0xae, 0xef, 0xac, 0x5d, 0xfe, 0x97, 0xda, 0xa9, 0xb6, 0x54, 0xef, 0x1d, 0xbd, 0xa5,
	0x5a, 0x5c, 0x41, 0xef, 0x39, 0x4a, 0x2b, 0x42, 0x97, 0x59, 0x23, 0xbe, 0xb9, 0x5f, 0x9d, 0xdf,
	0x3e, 0xff, 0x1e, 0x5f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x51, 0xa7, 0xe4, 0x2e, 0x03,
	0x00, 0x00,
}
