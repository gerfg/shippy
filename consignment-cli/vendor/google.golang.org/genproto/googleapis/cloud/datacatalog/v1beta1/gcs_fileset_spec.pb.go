// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datacatalog/v1beta1/gcs_fileset_spec.proto

package datacatalog

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Describes a Cloud Storage fileset entry.
type GcsFilesetSpec struct {
	// Required. Patterns to identify a set of files in Google Cloud Storage.
	//
	// Examples of valid file_patterns:
	//
	//  * `gs://bucket_name/*`: matches all files in `bucket_name`
	//  * `gs://bucket_name/file*`: matches files prefixed by `file` in
	//                              `bucket_name`
	//  * `gs://bucket_name/a/*/b`: matches all files in `bucket_name` that match
	//                              `a/*/b` pattern, such as `a/c/b`, `a/d/b`
	//  * `gs://another_bucket/a.txt`: matches `gs://another_bucket/a.txt`
	FilePatterns []string `protobuf:"bytes,1,rep,name=file_patterns,json=filePatterns,proto3" json:"file_patterns,omitempty"`
	// Output only. Sample files contained in this fileset, not all files contained in this
	// fileset are represented here.
	SampleGcsFileSpecs   []*GcsFileSpec `protobuf:"bytes,2,rep,name=sample_gcs_file_specs,json=sampleGcsFileSpecs,proto3" json:"sample_gcs_file_specs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GcsFilesetSpec) Reset()         { *m = GcsFilesetSpec{} }
func (m *GcsFilesetSpec) String() string { return proto.CompactTextString(m) }
func (*GcsFilesetSpec) ProtoMessage()    {}
func (*GcsFilesetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6379beb7306c38, []int{0}
}

func (m *GcsFilesetSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GcsFilesetSpec.Unmarshal(m, b)
}
func (m *GcsFilesetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GcsFilesetSpec.Marshal(b, m, deterministic)
}
func (m *GcsFilesetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GcsFilesetSpec.Merge(m, src)
}
func (m *GcsFilesetSpec) XXX_Size() int {
	return xxx_messageInfo_GcsFilesetSpec.Size(m)
}
func (m *GcsFilesetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GcsFilesetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_GcsFilesetSpec proto.InternalMessageInfo

func (m *GcsFilesetSpec) GetFilePatterns() []string {
	if m != nil {
		return m.FilePatterns
	}
	return nil
}

func (m *GcsFilesetSpec) GetSampleGcsFileSpecs() []*GcsFileSpec {
	if m != nil {
		return m.SampleGcsFileSpecs
	}
	return nil
}

// Specifications of a single file in GCS.
type GcsFileSpec struct {
	// Required. The full file path. Example: `gs://bucket_name/a/b.txt`.
	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	// Output only. Timestamps about the GCS file.
	GcsTimestamps *SystemTimestamps `protobuf:"bytes,2,opt,name=gcs_timestamps,json=gcsTimestamps,proto3" json:"gcs_timestamps,omitempty"`
	// Output only. The size of the file, in bytes.
	SizeBytes            int64    `protobuf:"varint,4,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GcsFileSpec) Reset()         { *m = GcsFileSpec{} }
func (m *GcsFileSpec) String() string { return proto.CompactTextString(m) }
func (*GcsFileSpec) ProtoMessage()    {}
func (*GcsFileSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6379beb7306c38, []int{1}
}

func (m *GcsFileSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GcsFileSpec.Unmarshal(m, b)
}
func (m *GcsFileSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GcsFileSpec.Marshal(b, m, deterministic)
}
func (m *GcsFileSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GcsFileSpec.Merge(m, src)
}
func (m *GcsFileSpec) XXX_Size() int {
	return xxx_messageInfo_GcsFileSpec.Size(m)
}
func (m *GcsFileSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GcsFileSpec.DiscardUnknown(m)
}

var xxx_messageInfo_GcsFileSpec proto.InternalMessageInfo

func (m *GcsFileSpec) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *GcsFileSpec) GetGcsTimestamps() *SystemTimestamps {
	if m != nil {
		return m.GcsTimestamps
	}
	return nil
}

func (m *GcsFileSpec) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*GcsFilesetSpec)(nil), "google.cloud.datacatalog.v1beta1.GcsFilesetSpec")
	proto.RegisterType((*GcsFileSpec)(nil), "google.cloud.datacatalog.v1beta1.GcsFileSpec")
}

func init() {
	proto.RegisterFile("google/cloud/datacatalog/v1beta1/gcs_fileset_spec.proto", fileDescriptor_4d6379beb7306c38)
}

var fileDescriptor_4d6379beb7306c38 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0xe5, 0x06, 0x21, 0xea, 0xd2, 0x0e, 0x91, 0x90, 0xa2, 0x0a, 0x89, 0xa8, 0x53, 0x16,
	0x12, 0xb5, 0x0c, 0x0c, 0x4c, 0x74, 0x80, 0x81, 0xa5, 0x6a, 0x99, 0x18, 0x88, 0x1c, 0xf7, 0xaf,
	0x6b, 0xc9, 0xa9, 0xad, 0xfc, 0xa6, 0x52, 0x79, 0x1b, 0x9e, 0x81, 0x17, 0xe2, 0x31, 0x18, 0x91,
	0xe3, 0x94, 0x66, 0xa9, 0xba, 0xde, 0xf9, 0xfe, 0xfb, 0x4e, 0xa6, 0xf7, 0x42, 0x6b, 0xa1, 0x20,
	0xe3, 0x4a, 0x7f, 0x2c, 0xb3, 0x25, 0xb3, 0x8c, 0x33, 0xcb, 0x94, 0x16, 0xd9, 0x76, 0x5c, 0x80,
	0x65, 0xe3, 0x4c, 0x70, 0xcc, 0x57, 0x52, 0x01, 0x82, 0xcd, 0xd1, 0x00, 0x4f, 0x4d, 0xa5, 0xad,
	0x0e, 0x63, 0x1f, 0x4c, 0xeb, 0x60, 0xda, 0x0a, 0xa6, 0x4d, 0x70, 0x78, 0xd3, 0x9c, 0x66, 0x46,
	0x66, 0x2b, 0x09, 0x6a, 0x99, 0x17, 0xb0, 0x66, 0x5b, 0xa9, 0x2b, 0x7f, 0x62, 0x38, 0x3e, 0xd9,
	0x6d, 0x65, 0x09, 0x68, 0x59, 0x69, 0xd0, 0x47, 0x46, 0x5f, 0x84, 0x0e, 0x9e, 0x39, 0x3e, 0x79,
	0x9e, 0x85, 0x01, 0x1e, 0x26, 0xb4, 0xef, 0xf0, 0x72, 0xc3, 0xac, 0x85, 0x6a, 0x83, 0x11, 0x89,
	0x83, 0xa4, 0x3b, 0x0d, 0x7e, 0x1e, 0x3b, 0xf3, 0x4b, 0xe7, 0xcc, 0x1a, 0x23, 0x04, 0x7a, 0x85,
	0xac, 0x34, 0x0a, 0xf2, 0xfd, 0xa6, 0x7a, 0x10, 0x46, 0x9d, 0x38, 0x48, 0x7a, 0x93, 0xdb, 0xf4,
	0xd4, 0xa4, 0xb4, 0xa9, 0x76, 0xbd, 0xae, 0x20, 0x98, 0x87, 0xfe, 0x60, 0x4b, 0xc7, 0xd1, 0x37,
	0xa1, 0xbd, 0x96, 0x10, 0xc6, 0xb4, 0xbb, 0x07, 0x5c, 0x47, 0x24, 0x26, 0x7b, 0xb8, 0x8b, 0x06,
	0x6e, 0x1d, 0xbe, 0xd3, 0x81, 0x23, 0x3a, 0xac, 0x8d, 0x3a, 0x31, 0x49, 0x7a, 0x93, 0xc9, 0x69,
	0xa2, 0xc5, 0x0e, 0x2d, 0x94, 0xaf, 0xff, 0x49, 0x8f, 0xd5, 0x17, 0x1c, 0x0f, 0x5a, 0x38, 0xa2,
	0x14, 0xe5, 0x27, 0xe4, 0xc5, 0xce, 0x02, 0x46, 0x67, 0x31, 0x49, 0x02, 0xff, 0xae, 0xeb, 0xe4,
	0xa9, 0x53, 0xa7, 0x86, 0x5e, 0x73, 0x5d, 0x1e, 0x2d, 0x9c, 0x91, 0xb7, 0x97, 0xc6, 0x13, 0x5a,
	0xb1, 0x8d, 0x48, 0x75, 0x25, 0x32, 0x01, 0x9b, 0xfa, 0x5f, 0x32, 0x6f, 0x31, 0x23, 0xf1, 0xf8,
	0x6f, 0x3e, 0xb4, 0xb4, 0x5f, 0x42, 0x8a, 0xf3, 0x3a, 0x7a, 0xf7, 0x17, 0x00, 0x00, 0xff, 0xff,
	0xee, 0xb5, 0x12, 0xd0, 0x83, 0x02, 0x00, 0x00,
}
