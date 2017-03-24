// Code generated by protoc-gen-go.
// source: orderer/configuration.proto
// DO NOT EDIT!

package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConsensusType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ConsensusType) Reset()                    { *m = ConsensusType{} }
func (m *ConsensusType) String() string            { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()               {}
func (*ConsensusType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type BatchSize struct {
	// Simply specified as number of messages for now, in the future
	// we may want to allow this to be specified by size in bytes
	MaxMessageCount uint32 `protobuf:"varint,1,opt,name=maxMessageCount" json:"maxMessageCount,omitempty"`
	// The byte count of the serialized messages in a batch cannot
	// exceed this value.
	AbsoluteMaxBytes uint32 `protobuf:"varint,2,opt,name=absoluteMaxBytes" json:"absoluteMaxBytes,omitempty"`
	// The byte count of the serialized messages in a batch should not
	// exceed this value.
	PreferredMaxBytes uint32 `protobuf:"varint,3,opt,name=preferredMaxBytes" json:"preferredMaxBytes,omitempty"`
}

func (m *BatchSize) Reset()                    { *m = BatchSize{} }
func (m *BatchSize) String() string            { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()               {}
func (*BatchSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type BatchTimeout struct {
	// Any duration string parseable by ParseDuration():
	// https://golang.org/pkg/time/#ParseDuration
	Timeout string `protobuf:"bytes,1,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *BatchTimeout) Reset()                    { *m = BatchTimeout{} }
func (m *BatchTimeout) String() string            { return proto.CompactTextString(m) }
func (*BatchTimeout) ProtoMessage()               {}
func (*BatchTimeout) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// When submitting a new chain configuration transaction to create a new chain,
// the first configuration item must be of type Orderer with Key CreationPolicy
// and contents of a Marshaled CreationPolicy. The policy should be set to the
// policy which was supplied by the ordering service for the client's chain
// creation. The digest should be the hash of the concatenation of the remaining
// ConfigurationItem bytes. The signatures of the configuration item should
// satisfy the policy for chain creation.
type CreationPolicy struct {
	// The name of the policy which should be used to validate the creation of
	// this chain
	Policy string `protobuf:"bytes,1,opt,name=policy" json:"policy,omitempty"`
}

func (m *CreationPolicy) Reset()                    { *m = CreationPolicy{} }
func (m *CreationPolicy) String() string            { return proto.CompactTextString(m) }
func (*CreationPolicy) ProtoMessage()               {}
func (*CreationPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// IngressPolicyNames is the set of policy names which incoming Broadcast signatures are filtered against
type IngressPolicyNames struct {
	// A list of policies, in evaluation these are 'or'-ed, note this is not a proper policy
	// because implementing referential policies in a general way is difficult, and dangerous
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *IngressPolicyNames) Reset()                    { *m = IngressPolicyNames{} }
func (m *IngressPolicyNames) String() string            { return proto.CompactTextString(m) }
func (*IngressPolicyNames) ProtoMessage()               {}
func (*IngressPolicyNames) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

// EgressPolicyNames is the set of policy names which incoming Deliver signatures are filtered against
type EgressPolicyNames struct {
	// A list of policies, in evaluation these are 'or'-ed, note this is not a proper policy
	// because implementing referential policies in a general way is difficult, and dangerous
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *EgressPolicyNames) Reset()                    { *m = EgressPolicyNames{} }
func (m *EgressPolicyNames) String() string            { return proto.CompactTextString(m) }
func (*EgressPolicyNames) ProtoMessage()               {}
func (*EgressPolicyNames) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

// ChainCreationPolicyNames is the set of policies which may be invoked for chain creation
type ChainCreationPolicyNames struct {
	// A list of policies, in evaluation these are 'or'-ed, note this is not a proper policy
	// because implementing referential policies in a general way is difficult, and dangerous
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *ChainCreationPolicyNames) Reset()                    { *m = ChainCreationPolicyNames{} }
func (m *ChainCreationPolicyNames) String() string            { return proto.CompactTextString(m) }
func (*ChainCreationPolicyNames) ProtoMessage()               {}
func (*ChainCreationPolicyNames) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

// Carries a list of bootstrap brokers, i.e. this is not the exclusive set of
// brokers an ordering service
type KafkaBrokers struct {
	// Each broker here should be identified using the (IP|host):port notation,
	// e.g. 127.0.0.1:7050, or localhost:7050 are valid entries
	Brokers []string `protobuf:"bytes,1,rep,name=brokers" json:"brokers,omitempty"`
}

func (m *KafkaBrokers) Reset()                    { *m = KafkaBrokers{} }
func (m *KafkaBrokers) String() string            { return proto.CompactTextString(m) }
func (*KafkaBrokers) ProtoMessage()               {}
func (*KafkaBrokers) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func init() {
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
	proto.RegisterType((*BatchTimeout)(nil), "orderer.BatchTimeout")
	proto.RegisterType((*CreationPolicy)(nil), "orderer.CreationPolicy")
	proto.RegisterType((*IngressPolicyNames)(nil), "orderer.IngressPolicyNames")
	proto.RegisterType((*EgressPolicyNames)(nil), "orderer.EgressPolicyNames")
	proto.RegisterType((*ChainCreationPolicyNames)(nil), "orderer.ChainCreationPolicyNames")
	proto.RegisterType((*KafkaBrokers)(nil), "orderer.KafkaBrokers")
}

func init() { proto.RegisterFile("orderer/configuration.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0xd5, 0x96, 0x0e, 0xad, 0xda, 0x45, 0x24, 0xe0, 0xa5, 0xc4, 0x4b, 0x2c, 0xa5,
	0x11, 0x7c, 0x83, 0x04, 0x0f, 0x22, 0x15, 0xa9, 0x3d, 0x79, 0xdb, 0xa4, 0x93, 0x64, 0x69, 0xb2,
	0x1b, 0x66, 0x37, 0xd0, 0xf8, 0x12, 0xbe, 0xb2, 0x74, 0x93, 0x0a, 0x5a, 0x10, 0x6f, 0xff, 0xb7,
	0xf3, 0x2d, 0xfc, 0xc3, 0xc0, 0x8d, 0xa2, 0x0d, 0x12, 0x52, 0x90, 0x28, 0x99, 0x8a, 0xac, 0x26,
	0x6e, 0x84, 0x92, 0x8b, 0x8a, 0x94, 0x51, 0x6c, 0xd0, 0x0d, 0xbd, 0x5b, 0x18, 0x47, 0x4a, 0x6a,
	0x94, 0xba, 0xd6, 0xeb, 0xa6, 0x42, 0xc6, 0xe0, 0xd4, 0x34, 0x15, 0xba, 0xce, 0xd4, 0xf1, 0x87,
	0x2b, 0x9b, 0xbd, 0x4f, 0x07, 0x86, 0x21, 0x37, 0x49, 0xfe, 0x26, 0x3e, 0x90, 0xf9, 0x70, 0x51,
	0xf2, 0xdd, 0x12, 0xb5, 0xe6, 0x19, 0x46, 0xaa, 0x96, 0xc6, 0xca, 0xe3, 0xd5, 0xef, 0x67, 0x36,
	0x83, 0x4b, 0x1e, 0x6b, 0x55, 0xd4, 0x06, 0x97, 0x7c, 0x17, 0x36, 0x06, 0xb5, 0x7b, 0x62, 0xd5,
	0xa3, 0x77, 0x36, 0x87, 0x49, 0x45, 0x98, 0x22, 0x11, 0x6e, 0xbe, 0xe5, 0x9e, 0x95, 0x8f, 0x07,
	0x9e, 0x0f, 0x23, 0x5b, 0x68, 0x2d, 0x4a, 0x54, 0xb5, 0x61, 0x2e, 0x0c, 0x4c, 0x1b, 0xbb, 0xe2,
	0x07, 0xf4, 0x7c, 0x38, 0x8f, 0x08, 0xed, 0xee, 0xaf, 0xaa, 0x10, 0x49, 0xc3, 0xae, 0xa1, 0x5f,
	0xd9, 0xd4, 0xa9, 0x1d, 0x79, 0x33, 0x60, 0x4f, 0x32, 0x23, 0xd4, 0xba, 0x15, 0x5f, 0x78, 0x89,
	0x9a, 0x5d, 0xc1, 0x99, 0xdc, 0x07, 0xd7, 0x99, 0xf6, 0xfc, 0xe1, 0xaa, 0x05, 0xef, 0x0e, 0x26,
	0x8f, 0xff, 0x54, 0xef, 0xc1, 0x8d, 0x72, 0x2e, 0xe4, 0xcf, 0x16, 0x7f, 0xfd, 0xf0, 0x61, 0xf4,
	0xcc, 0xd3, 0x2d, 0x0f, 0x49, 0x6d, 0x91, 0xf4, 0x7e, 0xb9, 0xb8, 0x8d, 0x9d, 0x77, 0xc0, 0x70,
	0xf1, 0x3e, 0xcf, 0x84, 0xc9, 0xeb, 0x78, 0x91, 0xa8, 0x32, 0xc8, 0x9b, 0x0a, 0xa9, 0xc0, 0x4d,
	0x86, 0x14, 0xa4, 0x3c, 0x26, 0x91, 0x04, 0xf6, 0xda, 0x3a, 0xe8, 0xae, 0x1d, 0xf7, 0x2d, 0x3f,
	0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x58, 0x13, 0xb8, 0x1c, 0x02, 0x00, 0x00,
}