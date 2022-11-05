// Code generated by protoc-gen-go. DO NOT EDIT.
// source: index.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Status int32

const (
	Status_INDEX_STATUS_UNKNOWN  Status = 0
	Status_INDEX_STATUS_CREATING Status = 5
	Status_INDEX_STATUS_ACTIVE   Status = 10
	Status_INDEX_STATUS_UPDATING Status = 15
	Status_INDEX_STATUS_DELETING Status = 20
)

var Status_name = map[int32]string{
	0:  "INDEX_STATUS_UNKNOWN",
	5:  "INDEX_STATUS_CREATING",
	10: "INDEX_STATUS_ACTIVE",
	15: "INDEX_STATUS_UPDATING",
	20: "INDEX_STATUS_DELETING",
}
var Status_value = map[string]int32{
	"INDEX_STATUS_UNKNOWN":  0,
	"INDEX_STATUS_CREATING": 5,
	"INDEX_STATUS_ACTIVE":   10,
	"INDEX_STATUS_UPDATING": 15,
	"INDEX_STATUS_DELETING": 20,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type IndexDefinition struct {
	Fields []*FieldDefinition `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
}

func (m *IndexDefinition) Reset()                    { *m = IndexDefinition{} }
func (m *IndexDefinition) String() string            { return proto.CompactTextString(m) }
func (*IndexDefinition) ProtoMessage()               {}
func (*IndexDefinition) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *IndexDefinition) GetFields() []*FieldDefinition {
	if m != nil {
		return m.Fields
	}
	return nil
}

type IndexMeta struct {
	PrimaryKey string `protobuf:"bytes,1,opt,name=primary_key,json=primaryKey" json:"primary_key,omitempty"`
	Status     Status `protobuf:"varint,2,opt,name=status,enum=api.Status" json:"status,omitempty"`
	Created    string `protobuf:"bytes,3,opt,name=created" json:"created,omitempty"`
	Updated    string `protobuf:"bytes,4,opt,name=updated" json:"updated,omitempty"`
}

func (m *IndexMeta) Reset()                    { *m = IndexMeta{} }
func (m *IndexMeta) String() string            { return proto.CompactTextString(m) }
func (*IndexMeta) ProtoMessage()               {}
func (*IndexMeta) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *IndexMeta) GetPrimaryKey() string {
	if m != nil {
		return m.PrimaryKey
	}
	return ""
}

func (m *IndexMeta) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_INDEX_STATUS_UNKNOWN
}

func (m *IndexMeta) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *IndexMeta) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

type AddIndexRequest struct {
	Name       string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PrimaryKey string           `protobuf:"bytes,2,opt,name=primary_key,json=primaryKey" json:"primary_key,omitempty"`
	Index      *IndexDefinition `protobuf:"bytes,3,opt,name=index" json:"index,omitempty"`
}

func (m *AddIndexRequest) Reset()                    { *m = AddIndexRequest{} }
func (m *AddIndexRequest) String() string            { return proto.CompactTextString(m) }
func (*AddIndexRequest) ProtoMessage()               {}
func (*AddIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *AddIndexRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddIndexRequest) GetPrimaryKey() string {
	if m != nil {
		return m.PrimaryKey
	}
	return ""
}

func (m *AddIndexRequest) GetIndex() *IndexDefinition {
	if m != nil {
		return m.Index
	}
	return nil
}

type AddIndexResponse struct {
	Name string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Meta *IndexMeta `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
}

func (m *AddIndexResponse) Reset()                    { *m = AddIndexResponse{} }
func (m *AddIndexResponse) String() string            { return proto.CompactTextString(m) }
func (*AddIndexResponse) ProtoMessage()               {}
func (*AddIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *AddIndexResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddIndexResponse) GetMeta() *IndexMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type DescribeIndexRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DescribeIndexRequest) Reset()                    { *m = DescribeIndexRequest{} }
func (m *DescribeIndexRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeIndexRequest) ProtoMessage()               {}
func (*DescribeIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *DescribeIndexRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DescribeIndexResponse struct {
	Name  string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Index *IndexDefinition `protobuf:"bytes,2,opt,name=index" json:"index,omitempty"`
	Meta  *IndexMeta       `protobuf:"bytes,3,opt,name=meta" json:"meta,omitempty"`
}

func (m *DescribeIndexResponse) Reset()                    { *m = DescribeIndexResponse{} }
func (m *DescribeIndexResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeIndexResponse) ProtoMessage()               {}
func (*DescribeIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *DescribeIndexResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DescribeIndexResponse) GetIndex() *IndexDefinition {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *DescribeIndexResponse) GetMeta() *IndexMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type DeleteIndexRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteIndexRequest) Reset()                    { *m = DeleteIndexRequest{} }
func (m *DeleteIndexRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteIndexRequest) ProtoMessage()               {}
func (*DeleteIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *DeleteIndexRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteIndexResponse struct {
	Name string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Meta *IndexMeta `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
}

func (m *DeleteIndexResponse) Reset()                    { *m = DeleteIndexResponse{} }
func (m *DeleteIndexResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteIndexResponse) ProtoMessage()               {}
func (*DeleteIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *DeleteIndexResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteIndexResponse) GetMeta() *IndexMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type ListIndexesRequest struct {
	Pattern       string `protobuf:"bytes,1,opt,name=pattern" json:"pattern,omitempty"`
	CreatedAfter  string `protobuf:"bytes,2,opt,name=created_after,json=createdAfter" json:"created_after,omitempty"`
	CreatedBefore string `protobuf:"bytes,3,opt,name=created_before,json=createdBefore" json:"created_before,omitempty"`
	UpdatedAfter  string `protobuf:"bytes,4,opt,name=updated_after,json=updatedAfter" json:"updated_after,omitempty"`
	UpdatedBefore string `protobuf:"bytes,5,opt,name=updated_before,json=updatedBefore" json:"updated_before,omitempty"`
}

func (m *ListIndexesRequest) Reset()                    { *m = ListIndexesRequest{} }
func (m *ListIndexesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListIndexesRequest) ProtoMessage()               {}
func (*ListIndexesRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *ListIndexesRequest) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *ListIndexesRequest) GetCreatedAfter() string {
	if m != nil {
		return m.CreatedAfter
	}
	return ""
}

func (m *ListIndexesRequest) GetCreatedBefore() string {
	if m != nil {
		return m.CreatedBefore
	}
	return ""
}

func (m *ListIndexesRequest) GetUpdatedAfter() string {
	if m != nil {
		return m.UpdatedAfter
	}
	return ""
}

func (m *ListIndexesRequest) GetUpdatedBefore() string {
	if m != nil {
		return m.UpdatedBefore
	}
	return ""
}

type ListIndexesResponse struct {
	Results []*DescribeIndexResponse `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *ListIndexesResponse) Reset()                    { *m = ListIndexesResponse{} }
func (m *ListIndexesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListIndexesResponse) ProtoMessage()               {}
func (*ListIndexesResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *ListIndexesResponse) GetResults() []*DescribeIndexResponse {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*IndexDefinition)(nil), "api.IndexDefinition")
	proto.RegisterType((*IndexMeta)(nil), "api.IndexMeta")
	proto.RegisterType((*AddIndexRequest)(nil), "api.AddIndexRequest")
	proto.RegisterType((*AddIndexResponse)(nil), "api.AddIndexResponse")
	proto.RegisterType((*DescribeIndexRequest)(nil), "api.DescribeIndexRequest")
	proto.RegisterType((*DescribeIndexResponse)(nil), "api.DescribeIndexResponse")
	proto.RegisterType((*DeleteIndexRequest)(nil), "api.DeleteIndexRequest")
	proto.RegisterType((*DeleteIndexResponse)(nil), "api.DeleteIndexResponse")
	proto.RegisterType((*ListIndexesRequest)(nil), "api.ListIndexesRequest")
	proto.RegisterType((*ListIndexesResponse)(nil), "api.ListIndexesResponse")
	proto.RegisterEnum("api.Status", Status_name, Status_value)
}

func init() { proto.RegisterFile("index.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xcf, 0x8f, 0xd2, 0x40,
	0x14, 0xb6, 0xfc, 0xda, 0xec, 0xeb, 0x0a, 0x64, 0x60, 0x23, 0xee, 0xc5, 0x4d, 0x8d, 0x09, 0x21,
	0x86, 0x03, 0x7a, 0x37, 0x75, 0x5b, 0x0d, 0xb2, 0x5b, 0x4d, 0x01, 0xf5, 0x46, 0xca, 0xf6, 0x91,
	0x4c, 0x84, 0xb6, 0xce, 0x0c, 0x89, 0xc4, 0xb3, 0x37, 0xff, 0x2e, 0xff, 0x2e, 0xdb, 0xe9, 0x0c,
	0xa4, 0xb0, 0x81, 0xc3, 0xde, 0x3a, 0xdf, 0xf7, 0xbd, 0xef, 0xfd, 0x98, 0x37, 0x05, 0x93, 0x46,
	0x21, 0xfe, 0xea, 0x27, 0x2c, 0x16, 0x31, 0x29, 0x07, 0x09, 0xbd, 0x32, 0x17, 0x14, 0x97, 0x61,
	0x8e, 0x58, 0xef, 0xa0, 0x31, 0xcc, 0x04, 0x0e, 0x2e, 0x68, 0x44, 0x05, 0x8d, 0x23, 0xf2, 0x1a,
	0x6a, 0x52, 0xc1, 0x3b, 0xa5, 0xeb, 0x72, 0xd7, 0x1c, 0xb4, 0xfb, 0x69, 0x54, 0xff, 0x43, 0x06,
	0xed, 0x54, 0xbe, 0xd2, 0x58, 0x7f, 0x0c, 0x38, 0x97, 0x0e, 0x77, 0x28, 0x02, 0xf2, 0x02, 0xcc,
	0x84, 0xd1, 0x55, 0xc0, 0x36, 0xb3, 0x1f, 0xb8, 0xe9, 0x18, 0xd7, 0x46, 0xf7, 0xdc, 0x07, 0x05,
	0x8d, 0x70, 0x43, 0x5e, 0x42, 0x8d, 0x8b, 0x40, 0xac, 0x33, 0x73, 0xa3, 0x5b, 0x1f, 0x98, 0xd2,
	0x7c, 0x2c, 0x21, 0x5f, 0x51, 0xa4, 0x03, 0x67, 0xf7, 0x0c, 0x03, 0x81, 0x61, 0xa7, 0x2c, 0x1d,
	0xf4, 0x31, 0x63, 0xd6, 0x49, 0x28, 0x99, 0x4a, 0xce, 0xa8, 0xa3, 0xc5, 0xa0, 0x61, 0x87, 0xa1,
	0xac, 0xc4, 0xc7, 0x9f, 0x6b, 0xe4, 0x82, 0x10, 0xa8, 0x44, 0xc1, 0x0a, 0x55, 0x15, 0xf2, 0x7b,
	0xbf, 0xc0, 0xd2, 0x41, 0x81, 0x3d, 0xa8, 0xca, 0x89, 0xc9, 0xcc, 0xba, 0xf9, 0xbd, 0x11, 0xf9,
	0xb9, 0xc4, 0xfa, 0x04, 0xcd, 0x5d, 0x4e, 0x9e, 0xc4, 0x11, 0xc7, 0x07, 0x93, 0x5a, 0x50, 0x59,
	0xa5, 0xd3, 0x91, 0xd9, 0xcc, 0x41, 0x7d, 0x67, 0x99, 0xcd, 0xcc, 0x97, 0x9c, 0xd5, 0x83, 0xb6,
	0x83, 0xfc, 0x9e, 0xd1, 0x39, 0x9e, 0x6a, 0xc2, 0xfa, 0x0d, 0x97, 0x7b, 0xda, 0x23, 0xc9, 0xb7,
	0x0d, 0x95, 0x4e, 0x36, 0xb4, 0x2d, 0xb4, 0x7c, 0xa4, 0xd0, 0x2e, 0x10, 0x07, 0x97, 0x28, 0x4e,
	0x97, 0x79, 0x07, 0xad, 0x82, 0xf2, 0x91, 0x13, 0xfa, 0x67, 0x00, 0xb9, 0xa5, 0x5c, 0x48, 0x1c,
	0xb9, 0xce, 0x9c, 0xae, 0x44, 0x12, 0x08, 0x81, 0x2c, 0x52, 0x8e, 0xfa, 0x98, 0xee, 0xda, 0x53,
	0xb5, 0x37, 0xb3, 0x60, 0x91, 0x22, 0xea, 0xb6, 0x2f, 0x14, 0x68, 0x67, 0x18, 0x79, 0x05, 0x75,
	0x2d, 0x9a, 0xe3, 0x22, 0x66, 0xa8, 0x56, 0x4e, 0x87, 0xbe, 0x97, 0x60, 0xe6, 0xa5, 0x36, 0x4d,
	0x79, 0xe5, 0xeb, 0x77, 0xa1, 0xc0, 0xad, 0x97, 0x16, 0x29, 0xaf, 0x6a, 0xee, 0xa5, 0xd0, 0xdc,
	0xcb, 0x1a, 0x41, 0xab, 0xd0, 0x87, 0x9a, 0xcb, 0x5b, 0x38, 0x63, 0xc8, 0xd7, 0x4b, 0xc1, 0xd3,
	0x46, 0xb2, 0x87, 0x77, 0x25, 0xc7, 0xf0, 0xe0, 0x4d, 0xfb, 0x5a, 0xda, 0xfb, 0x6b, 0x40, 0x6d,
	0xac, 0x9f, 0x4d, 0x7b, 0xe8, 0x39, 0xee, 0xf7, 0xd9, 0x78, 0x62, 0x4f, 0xa6, 0xe3, 0xd9, 0xd4,
	0x1b, 0x79, 0x9f, 0xbf, 0x79, 0xcd, 0x27, 0xe4, 0x39, 0x5c, 0x16, 0x98, 0x1b, 0xdf, 0xb5, 0x27,
	0x43, 0xef, 0x63, 0xb3, 0x4a, 0x9e, 0x41, 0xab, 0x40, 0xd9, 0x37, 0x93, 0xe1, 0x57, 0xb7, 0x09,
	0x07, 0x31, 0xd3, 0x2f, 0x4e, 0x1e, 0xd3, 0x38, 0xa0, 0x1c, 0xf7, 0xd6, 0x95, 0x54, 0x7b, 0x5e,
	0x93, 0xbf, 0x95, 0x37, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x58, 0xd2, 0x25, 0xf9, 0x77, 0x04,
	0x00, 0x00,
}