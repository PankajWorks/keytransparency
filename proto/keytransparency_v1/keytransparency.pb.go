// Code generated by protoc-gen-go.
// source: proto/keytransparency_v1/keytransparency.proto
// DO NOT EDIT!

/*
Package keytransparency_v1 is a generated protocol buffer package.

It is generated from these files:
	proto/keytransparency_v1/keytransparency.proto

It has these top-level messages:
	HkpLookupRequest
	HttpResponse
	GetEntryResponse
	Committed
	Profile
	Entry
	PublicKey
	KeyValue
	SignedKV
	GetEntryRequest
	ListEntryHistoryRequest
	ListEntryHistoryResponse
	EntryUpdate
	UpdateEntryRequest
	UpdateEntryResponse
*/
package keytransparency_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ctmap "github.com/google/key-transparency/proto/ctmap"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

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

// HkpLookupRequest contains query parameters for retrieving PGP keys.
type HkpLookupRequest struct {
	// Op specifies the operation to be performed on the keyserver.
	// - "get" returns the pgp key specified in the search parameter.
	// - "index" returns 501 (not implemented).
	// - "vindex" returns 501 (not implemented).
	Op string `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	// Search specifies the email address or key id being queried.
	Search string `protobuf:"bytes,2,opt,name=search" json:"search,omitempty"`
	// Options specifies what output format to use.
	// - "mr" machine readable will set the content type to "application/pgp-keys"
	// - other options will be ignored.
	Options string `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
	// Exact specifies an exact match on search. Always on. If specified in the
	// URL, its value will be ignored.
	Exact string `protobuf:"bytes,4,opt,name=exact" json:"exact,omitempty"`
	// fingerprint is ignored.
	Fingerprint string `protobuf:"bytes,5,opt,name=fingerprint" json:"fingerprint,omitempty"`
}

func (m *HkpLookupRequest) Reset()                    { *m = HkpLookupRequest{} }
func (m *HkpLookupRequest) String() string            { return proto.CompactTextString(m) }
func (*HkpLookupRequest) ProtoMessage()               {}
func (*HkpLookupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// HttpBody represents an http body.
type HttpResponse struct {
	// Header content type.
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	// The http body itself.
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *HttpResponse) Reset()                    { *m = HttpResponse{} }
func (m *HttpResponse) String() string            { return proto.CompactTextString(m) }
func (*HttpResponse) ProtoMessage()               {}
func (*HttpResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// GetEntryResponse
type GetEntryResponse struct {
	// vrf is the output of VRF on user_id.
	Vrf []byte `protobuf:"bytes,1,opt,name=vrf,proto3" json:"vrf,omitempty"`
	// vrf_proof is the proof for VRF on user_id.
	VrfProof []byte `protobuf:"bytes,2,opt,name=vrf_proof,json=vrfProof,proto3" json:"vrf_proof,omitempty"`
	// committed contains the profile for this account and connects the data
	// in profile to the commitment in leaf_proof.
	Committed *Committed `protobuf:"bytes,3,opt,name=committed" json:"committed,omitempty"`
	// leaf_proof contains an Entry and an inclusion proof in the sparse merkle tree at end_epoch.
	LeafProof *ctmap.GetLeafResponse `protobuf:"bytes,5,opt,name=leaf_proof,json=leafProof" json:"leaf_proof,omitempty"`
	// smh contains the signed map head for the sparse merkle tree.
	// smh is also stored in the append only log.
	Smh *ctmap.SignedMapHead `protobuf:"bytes,6,opt,name=smh" json:"smh,omitempty"`
	// smh_sct is the SCT showing that smh was submitted to CT logs.
	// TODO: Support storing smh in multiple logs.
	SmhSct []byte `protobuf:"bytes,7,opt,name=smh_sct,json=smhSct,proto3" json:"smh_sct,omitempty"`
}

func (m *GetEntryResponse) Reset()                    { *m = GetEntryResponse{} }
func (m *GetEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEntryResponse) ProtoMessage()               {}
func (*GetEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetEntryResponse) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

func (m *GetEntryResponse) GetLeafProof() *ctmap.GetLeafResponse {
	if m != nil {
		return m.LeafProof
	}
	return nil
}

func (m *GetEntryResponse) GetSmh() *ctmap.SignedMapHead {
	if m != nil {
		return m.Smh
	}
	return nil
}

// Committed represents the data committed to in a cryptographic commitment.
// commitment = HMAC_SHA512_256(key, data)
type Committed struct {
	// key is the 16 byte random commitment key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// data is the data being committed to.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Committed) Reset()                    { *m = Committed{} }
func (m *Committed) String() string            { return proto.CompactTextString(m) }
func (*Committed) ProtoMessage()               {}
func (*Committed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Profile contains data hidden behind the crypto commitement.
type Profile struct {
	// Keys is a map of appIds to keys.
	Keys map[string][]byte `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Profile) GetKeys() map[string][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

// Entry contains a commitment to profile and a set of authorized update keys.
// Entry is placed in the verifiable map as leaf data.
type Entry struct {
	// commitment is a cryptographic commitment to arbitrary data.
	Commitment []byte `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	// authorized_keys is the set of keys allowed to sign updates for this entry.
	AuthorizedKeys []*PublicKey `protobuf:"bytes,2,rep,name=authorized_keys,json=authorizedKeys" json:"authorized_keys,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Entry) GetAuthorizedKeys() []*PublicKey {
	if m != nil {
		return m.AuthorizedKeys
	}
	return nil
}

// PublicKey defines a key this domain uses to sign MapHeads with.
type PublicKey struct {
	// KeyFormats from Keyczar.
	//
	// Types that are valid to be assigned to KeyType:
	//	*PublicKey_Ed25519
	//	*PublicKey_RsaVerifyingSha256_2048
	//	*PublicKey_EcdsaVerifyingP256
	KeyType isPublicKey_KeyType `protobuf_oneof:"key_type"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isPublicKey_KeyType interface {
	isPublicKey_KeyType()
}

type PublicKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}
type PublicKey_RsaVerifyingSha256_2048 struct {
	RsaVerifyingSha256_2048 []byte `protobuf:"bytes,2,opt,name=rsa_verifying_sha256_2048,json=rsaVerifyingSha2562048,proto3,oneof"`
}
type PublicKey_EcdsaVerifyingP256 struct {
	EcdsaVerifyingP256 []byte `protobuf:"bytes,3,opt,name=ecdsa_verifying_p256,json=ecdsaVerifyingP256,proto3,oneof"`
}

func (*PublicKey_Ed25519) isPublicKey_KeyType()                 {}
func (*PublicKey_RsaVerifyingSha256_2048) isPublicKey_KeyType() {}
func (*PublicKey_EcdsaVerifyingP256) isPublicKey_KeyType()      {}

func (m *PublicKey) GetKeyType() isPublicKey_KeyType {
	if m != nil {
		return m.KeyType
	}
	return nil
}

func (m *PublicKey) GetEd25519() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PublicKey) GetRsaVerifyingSha256_2048() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_RsaVerifyingSha256_2048); ok {
		return x.RsaVerifyingSha256_2048
	}
	return nil
}

func (m *PublicKey) GetEcdsaVerifyingP256() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_EcdsaVerifyingP256); ok {
		return x.EcdsaVerifyingP256
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PublicKey) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PublicKey_OneofMarshaler, _PublicKey_OneofUnmarshaler, _PublicKey_OneofSizer, []interface{}{
		(*PublicKey_Ed25519)(nil),
		(*PublicKey_RsaVerifyingSha256_2048)(nil),
		(*PublicKey_EcdsaVerifyingP256)(nil),
	}
}

func _PublicKey_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_2048:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RsaVerifyingSha256_2048)
	case *PublicKey_EcdsaVerifyingP256:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.EcdsaVerifyingP256)
	case nil:
	default:
		return fmt.Errorf("PublicKey.KeyType has unexpected type %T", x)
	}
	return nil
}

func _PublicKey_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PublicKey)
	switch tag {
	case 1: // key_type.ed25519
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_Ed25519{x}
		return true, err
	case 2: // key_type.rsa_verifying_sha256_2048
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_RsaVerifyingSha256_2048{x}
		return true, err
	case 3: // key_type.ecdsa_verifying_p256
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_EcdsaVerifyingP256{x}
		return true, err
	default:
		return false, nil
	}
}

func _PublicKey_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ed25519)))
		n += len(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_2048:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RsaVerifyingSha256_2048)))
		n += len(x.RsaVerifyingSha256_2048)
	case *PublicKey_EcdsaVerifyingP256:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.EcdsaVerifyingP256)))
		n += len(x.EcdsaVerifyingP256)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// KeyValue is a map entry.
type KeyValue struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// SignedKV is a signed change to a map entry.
type SignedKV struct {
	// key_value is a serialized KeyValue.
	KeyValue []byte `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
	// signatures on keyvalue. Must be signed by keys from both previous and
	// current epochs. The first proves ownership of new epoch key, and the
	// second proves that the correct owner is making this change.
	Signatures map[uint64][]byte `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"fixed64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// previous contains the hash of the previous entry that this mutation is
	// modifying creating a hash chain of all mutations. The hash used is
	// CommonJSON in "github.com/benlaurie/objecthash/go/objecthash".
	Previous []byte `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
}

func (m *SignedKV) Reset()                    { *m = SignedKV{} }
func (m *SignedKV) String() string            { return proto.CompactTextString(m) }
func (*SignedKV) ProtoMessage()               {}
func (*SignedKV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SignedKV) GetSignatures() map[uint64][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// Get request for a user object.
type GetEntryRequest struct {
	// Last trusted epoch by the client.
	// int64 epoch_start = 3;
	// Absence of the epoch_end field indicates a request for the current value.
	EpochEnd int64 `protobuf:"varint,1,opt,name=epoch_end,json=epochEnd" json:"epoch_end,omitempty"`
	// User identifier. Most commonly an email address.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *GetEntryRequest) Reset()                    { *m = GetEntryRequest{} }
func (m *GetEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntryRequest) ProtoMessage()               {}
func (*GetEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// Get a list of historical values for a user.
type ListEntryHistoryRequest struct {
	// The user identifier.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// from_epoch is the starting epcoh.
	StartEpoch int64 `protobuf:"varint,2,opt,name=start_epoch,json=startEpoch" json:"start_epoch,omitempty"`
	// The maximum number of entries to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListEntryHistoryRequest) Reset()                    { *m = ListEntryHistoryRequest{} }
func (m *ListEntryHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryRequest) ProtoMessage()               {}
func (*ListEntryHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// A paginated history of values for a user.
type ListEntryHistoryResponse struct {
	// The list of values this user_id has contained over time.
	Values []*GetEntryResponse `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	// The next time to query for pagination.
	NextEpoch int64 `protobuf:"varint,2,opt,name=next_epoch,json=nextEpoch" json:"next_epoch,omitempty"`
}

func (m *ListEntryHistoryResponse) Reset()                    { *m = ListEntryHistoryResponse{} }
func (m *ListEntryHistoryResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryResponse) ProtoMessage()               {}
func (*ListEntryHistoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ListEntryHistoryResponse) GetValues() []*GetEntryResponse {
	if m != nil {
		return m.Values
	}
	return nil
}

type EntryUpdate struct {
	// update authorizes the change to profile.
	Update *SignedKV `protobuf:"bytes,2,opt,name=update" json:"update,omitempty"`
	// commitment contains the serialized Profile protobuf.
	Committed *Committed `protobuf:"bytes,3,opt,name=committed" json:"committed,omitempty"`
}

func (m *EntryUpdate) Reset()                    { *m = EntryUpdate{} }
func (m *EntryUpdate) String() string            { return proto.CompactTextString(m) }
func (*EntryUpdate) ProtoMessage()               {}
func (*EntryUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *EntryUpdate) GetUpdate() *SignedKV {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *EntryUpdate) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

// Update a user's profile.
type UpdateEntryRequest struct {
	// user_id specifies the id for the new account to be registered.
	UserId      string       `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	EntryUpdate *EntryUpdate `protobuf:"bytes,2,opt,name=entry_update,json=entryUpdate" json:"entry_update,omitempty"`
}

func (m *UpdateEntryRequest) Reset()                    { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()               {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *UpdateEntryRequest) GetEntryUpdate() *EntryUpdate {
	if m != nil {
		return m.EntryUpdate
	}
	return nil
}

// UpdateEntryResponse contains a proof once the update has been included in
// the Merkel Tree.
type UpdateEntryResponse struct {
	Proof *GetEntryResponse `protobuf:"bytes,1,opt,name=proof" json:"proof,omitempty"`
}

func (m *UpdateEntryResponse) Reset()                    { *m = UpdateEntryResponse{} }
func (m *UpdateEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryResponse) ProtoMessage()               {}
func (*UpdateEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *UpdateEntryResponse) GetProof() *GetEntryResponse {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterType((*HkpLookupRequest)(nil), "keytransparency.v1.HkpLookupRequest")
	proto.RegisterType((*HttpResponse)(nil), "keytransparency.v1.HttpResponse")
	proto.RegisterType((*GetEntryResponse)(nil), "keytransparency.v1.GetEntryResponse")
	proto.RegisterType((*Committed)(nil), "keytransparency.v1.Committed")
	proto.RegisterType((*Profile)(nil), "keytransparency.v1.Profile")
	proto.RegisterType((*Entry)(nil), "keytransparency.v1.Entry")
	proto.RegisterType((*PublicKey)(nil), "keytransparency.v1.PublicKey")
	proto.RegisterType((*KeyValue)(nil), "keytransparency.v1.KeyValue")
	proto.RegisterType((*SignedKV)(nil), "keytransparency.v1.SignedKV")
	proto.RegisterType((*GetEntryRequest)(nil), "keytransparency.v1.GetEntryRequest")
	proto.RegisterType((*ListEntryHistoryRequest)(nil), "keytransparency.v1.ListEntryHistoryRequest")
	proto.RegisterType((*ListEntryHistoryResponse)(nil), "keytransparency.v1.ListEntryHistoryResponse")
	proto.RegisterType((*EntryUpdate)(nil), "keytransparency.v1.EntryUpdate")
	proto.RegisterType((*UpdateEntryRequest)(nil), "keytransparency.v1.UpdateEntryRequest")
	proto.RegisterType((*UpdateEntryResponse)(nil), "keytransparency.v1.UpdateEntryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for KeyTransparencyService service

type KeyTransparencyServiceClient interface {
	// GetEntry returns a user's entry in the Merkle Tree. Entries contain
	// signed commitments to a profile, which is also returned.
	GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error)
	// ListEntryHistory returns a list of GetEntryRespons covering several epochs.
	ListEntryHistory(ctx context.Context, in *ListEntryHistoryRequest, opts ...grpc.CallOption) (*ListEntryHistoryResponse, error)
	// UpdateEntry submits a SignedEntryUpdate.  Returns empty until this update
	// has been included in an epoch.  Clients must retry until this function
	// returns a proof.
	UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error)
	//
	// // List the SignedMapHeads, from epoch to epoch.
	// rpc ListSMH(ListSMHRequest) returns (ListSMHResponse);
	//
	// // List the EntryUpdates by update number.
	// rpc ListUpdate(ListUpdateRequest) returns (ListUpdateResponse);
	//
	// // ListSteps combines SMH and EntryUpdates into single list.
	// rpc ListSteps(ListStepsRequest) returns (ListStepsResponse);
	HkpLookup(ctx context.Context, in *HkpLookupRequest, opts ...grpc.CallOption) (*HttpResponse, error)
}

type keyTransparencyServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeyTransparencyServiceClient(cc *grpc.ClientConn) KeyTransparencyServiceClient {
	return &keyTransparencyServiceClient{cc}
}

func (c *keyTransparencyServiceClient) GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error) {
	out := new(GetEntryResponse)
	err := grpc.Invoke(ctx, "/keytransparency.v1.KeyTransparencyService/GetEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyServiceClient) ListEntryHistory(ctx context.Context, in *ListEntryHistoryRequest, opts ...grpc.CallOption) (*ListEntryHistoryResponse, error) {
	out := new(ListEntryHistoryResponse)
	err := grpc.Invoke(ctx, "/keytransparency.v1.KeyTransparencyService/ListEntryHistory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyServiceClient) UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error) {
	out := new(UpdateEntryResponse)
	err := grpc.Invoke(ctx, "/keytransparency.v1.KeyTransparencyService/UpdateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyServiceClient) HkpLookup(ctx context.Context, in *HkpLookupRequest, opts ...grpc.CallOption) (*HttpResponse, error) {
	out := new(HttpResponse)
	err := grpc.Invoke(ctx, "/keytransparency.v1.KeyTransparencyService/HkpLookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KeyTransparencyService service

type KeyTransparencyServiceServer interface {
	// GetEntry returns a user's entry in the Merkle Tree. Entries contain
	// signed commitments to a profile, which is also returned.
	GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error)
	// ListEntryHistory returns a list of GetEntryRespons covering several epochs.
	ListEntryHistory(context.Context, *ListEntryHistoryRequest) (*ListEntryHistoryResponse, error)
	// UpdateEntry submits a SignedEntryUpdate.  Returns empty until this update
	// has been included in an epoch.  Clients must retry until this function
	// returns a proof.
	UpdateEntry(context.Context, *UpdateEntryRequest) (*UpdateEntryResponse, error)
	//
	// // List the SignedMapHeads, from epoch to epoch.
	// rpc ListSMH(ListSMHRequest) returns (ListSMHResponse);
	//
	// // List the EntryUpdates by update number.
	// rpc ListUpdate(ListUpdateRequest) returns (ListUpdateResponse);
	//
	// // ListSteps combines SMH and EntryUpdates into single list.
	// rpc ListSteps(ListStepsRequest) returns (ListStepsResponse);
	HkpLookup(context.Context, *HkpLookupRequest) (*HttpResponse, error)
}

func RegisterKeyTransparencyServiceServer(s *grpc.Server, srv KeyTransparencyServiceServer) {
	s.RegisterService(&_KeyTransparencyService_serviceDesc, srv)
}

func _KeyTransparencyService_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyServiceServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keytransparency.v1.KeyTransparencyService/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyServiceServer).GetEntry(ctx, req.(*GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyService_ListEntryHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntryHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyServiceServer).ListEntryHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keytransparency.v1.KeyTransparencyService/ListEntryHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyServiceServer).ListEntryHistory(ctx, req.(*ListEntryHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyService_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyServiceServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keytransparency.v1.KeyTransparencyService/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyServiceServer).UpdateEntry(ctx, req.(*UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyService_HkpLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HkpLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyServiceServer).HkpLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keytransparency.v1.KeyTransparencyService/HkpLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyServiceServer).HkpLookup(ctx, req.(*HkpLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyTransparencyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "keytransparency.v1.KeyTransparencyService",
	HandlerType: (*KeyTransparencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntry",
			Handler:    _KeyTransparencyService_GetEntry_Handler,
		},
		{
			MethodName: "ListEntryHistory",
			Handler:    _KeyTransparencyService_ListEntryHistory_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _KeyTransparencyService_UpdateEntry_Handler,
		},
		{
			MethodName: "HkpLookup",
			Handler:    _KeyTransparencyService_HkpLookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("proto/keytransparency_v1/keytransparency.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1059 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0xdf, 0x73, 0xdb, 0xc4,
	0x13, 0xff, 0xda, 0x8e, 0x7f, 0x68, 0xed, 0x49, 0x32, 0xd7, 0x7c, 0x5d, 0xa3, 0x34, 0x34, 0x28,
	0x50, 0x60, 0x28, 0x36, 0x31, 0x0d, 0xb4, 0xa1, 0xbc, 0x94, 0x09, 0x35, 0xd3, 0x30, 0x13, 0xe4,
	0x92, 0x57, 0xcd, 0x59, 0x3a, 0xdb, 0x1a, 0xdb, 0x92, 0x2a, 0x9d, 0x4c, 0x0d, 0xd3, 0x19, 0x86,
	0x37, 0x66, 0x78, 0x82, 0x37, 0x9e, 0xf9, 0x83, 0x98, 0xe1, 0x5f, 0xe0, 0x6f, 0xe0, 0x99, 0xbd,
	0x3b, 0xc9, 0x96, 0x1d, 0xb5, 0xcd, 0x0c, 0x2f, 0xf1, 0xed, 0xde, 0x67, 0x77, 0x3f, 0xfb, 0xb9,
	0xd5, 0xe5, 0xa0, 0x1d, 0x84, 0x3e, 0xf7, 0x3b, 0x13, 0xb6, 0xe0, 0x21, 0xf5, 0xa2, 0x80, 0x86,
	0xcc, 0xb3, 0x17, 0xd6, 0xfc, 0x78, 0xd3, 0xa5, 0x80, 0x84, 0x6c, 0xba, 0xe7, 0xc7, 0xfa, 0xc3,
	0x91, 0xcb, 0xc7, 0xf1, 0xa0, 0x6d, 0xfb, 0xb3, 0xce, 0xc8, 0x19, 0xb0, 0xe9, 0xdc, 0xf5, 0x44,
	0xf8, 0x87, 0x59, 0x60, 0x47, 0x15, 0xb2, 0xf9, 0x8c, 0x06, 0xea, 0xaf, 0xca, 0xa8, 0xdf, 0x1a,
	0xf9, 0xfe, 0x68, 0xca, 0x3a, 0x34, 0x70, 0x3b, 0xd4, 0xf3, 0x7c, 0x4e, 0xb9, 0xeb, 0x7b, 0x91,
	0xda, 0x35, 0x7e, 0x29, 0xc0, 0x6e, 0x6f, 0x12, 0x9c, 0xfb, 0xfe, 0x24, 0x0e, 0x4c, 0xf6, 0x2c,
	0x66, 0x11, 0x27, 0xdb, 0x50, 0xf4, 0x83, 0x56, 0xe1, 0xb0, 0xf0, 0x9e, 0x66, 0xe2, 0x8a, 0x34,
	0xa1, 0x12, 0x31, 0x1a, 0xda, 0xe3, 0x56, 0x51, 0xfa, 0x12, 0x8b, 0xb4, 0xa0, 0xea, 0x07, 0x32,
	0x5b, 0xab, 0x24, 0x37, 0x52, 0x93, 0xec, 0x41, 0x99, 0x3d, 0xa7, 0x36, 0x6f, 0x6d, 0x49, 0xbf,
	0x32, 0xc8, 0x21, 0xd4, 0x87, 0xae, 0x37, 0x62, 0x61, 0x10, 0xba, 0x1e, 0x6f, 0x95, 0xe5, 0x5e,
	0xd6, 0x65, 0x9c, 0x41, 0xa3, 0xc7, 0x39, 0x12, 0x89, 0x02, 0x4c, 0xc3, 0xc8, 0x5b, 0xd0, 0xb0,
	0x7d, 0x8f, 0x33, 0x8f, 0x5b, 0x7c, 0x11, 0xb0, 0x84, 0x53, 0x3d, 0xf1, 0x3d, 0x45, 0x17, 0x21,
	0xb0, 0x35, 0xf0, 0x9d, 0x85, 0xa4, 0xd6, 0x30, 0xe5, 0xda, 0xf8, 0x07, 0xbb, 0x7a, 0xcc, 0xf8,
	0x99, 0xc7, 0xc3, 0xc5, 0x32, 0xd7, 0x2e, 0x94, 0xe6, 0xe1, 0x50, 0xa6, 0x68, 0x98, 0x62, 0x49,
	0xf6, 0x41, 0xc3, 0x1f, 0x0b, 0x95, 0xf0, 0x87, 0x49, 0x7c, 0x0d, 0x1d, 0x17, 0xc2, 0x26, 0x9f,
	0x81, 0x86, 0x82, 0xcf, 0x5c, 0xce, 0x99, 0x23, 0xdb, 0xab, 0x77, 0x0f, 0xda, 0x57, 0x4f, 0xa7,
	0xfd, 0x45, 0x0a, 0x32, 0x57, 0x78, 0x72, 0x02, 0x30, 0x65, 0x34, 0x4d, 0x5d, 0x96, 0xd1, 0xcd,
	0xb6, 0x3a, 0x16, 0x24, 0x76, 0x8e, 0x7b, 0x29, 0x2f, 0x53, 0x13, 0x48, 0x55, 0xf3, 0x0e, 0x94,
	0xa2, 0xd9, 0xb8, 0x55, 0x91, 0xf8, 0xbd, 0x04, 0xdf, 0x77, 0x47, 0x1e, 0x73, 0xbe, 0xa6, 0x41,
	0x8f, 0x51, 0xc7, 0x14, 0x00, 0x72, 0x13, 0xaa, 0xf8, 0x63, 0x45, 0x28, 0x70, 0x55, 0xd2, 0xae,
	0xa0, 0xd9, 0xb7, 0xb9, 0x71, 0x0c, 0xda, 0x92, 0x8f, 0x68, 0x18, 0xf9, 0xa6, 0x0d, 0xe3, 0x52,
	0x68, 0xe5, 0x50, 0x4e, 0x53, 0xad, 0xc4, 0xda, 0x78, 0x01, 0x55, 0x2c, 0x3e, 0x74, 0xa7, 0x8c,
	0x3c, 0x80, 0x2d, 0x44, 0x45, 0x18, 0x51, 0xc2, 0xfa, 0xef, 0xe4, 0x75, 0x9b, 0x40, 0xdb, 0x4f,
	0x10, 0xa7, 0xe4, 0x95, 0x21, 0xfa, 0xa7, 0xa0, 0x2d, 0x5d, 0xd9, 0xc2, 0x9a, 0x2a, 0x8c, 0xf3,
	0x30, 0xa7, 0xd3, 0x98, 0x25, 0x95, 0x95, 0x71, 0x5a, 0xbc, 0x5f, 0x30, 0x7c, 0x28, 0xab, 0xa0,
	0x37, 0x01, 0x94, 0x7e, 0x33, 0x3c, 0xd9, 0x84, 0x74, 0xc6, 0x43, 0xbe, 0x84, 0x1d, 0x1a, 0xf3,
	0xb1, 0x1f, 0xba, 0xdf, 0x33, 0xc7, 0x92, 0x3c, 0x8b, 0x92, 0x67, 0xee, 0xa9, 0x5c, 0xc4, 0x83,
	0xa9, 0x6b, 0x23, 0x25, 0x73, 0x7b, 0x15, 0x25, 0x18, 0x1a, 0x7f, 0x14, 0x40, 0x5b, 0xee, 0x12,
	0x1d, 0xaa, 0xcc, 0xe9, 0x9e, 0x9c, 0x1c, 0x3f, 0x50, 0x25, 0x7b, 0xff, 0x33, 0x53, 0x07, 0x4e,
	0xc0, 0x1b, 0x61, 0x44, 0xad, 0x39, 0x0b, 0xdd, 0xe1, 0x02, 0x87, 0xd4, 0x8a, 0xc6, 0xb4, 0x7b,
	0xf2, 0x89, 0xd5, 0xfd, 0xe8, 0xde, 0x7d, 0xd5, 0x08, 0xa2, 0x9b, 0x08, 0xb9, 0x4c, 0x11, 0x7d,
	0x09, 0x10, 0xfb, 0xa4, 0x0b, 0x7b, 0xcc, 0x76, 0xd6, 0xc2, 0x03, 0xdc, 0x93, 0x93, 0x24, 0xe2,
	0x88, 0xdc, 0x5d, 0x46, 0x5e, 0xe0, 0xde, 0x23, 0x80, 0x1a, 0xb6, 0x22, 0x27, 0xdd, 0xe8, 0x42,
	0x0d, 0xf9, 0x5d, 0x0a, 0x9d, 0x72, 0x0e, 0x32, 0x57, 0x4f, 0xe3, 0xcf, 0x02, 0xd4, 0xd4, 0xb4,
	0x3c, 0xb9, 0x14, 0xc3, 0x2d, 0x92, 0x29, 0x98, 0x0a, 0x15, 0xd9, 0x55, 0xc6, 0x73, 0x80, 0x08,
	0x81, 0x94, 0xc7, 0x21, 0x4b, 0x75, 0xbc, 0x9b, 0xa7, 0x63, 0x9a, 0x4e, 0x2e, 0x14, 0x5c, 0x1d,
	0x7b, 0x26, 0x1e, 0x45, 0xac, 0x05, 0x21, 0x9b, 0xbb, 0x7e, 0xac, 0x2e, 0x02, 0xac, 0x94, 0xda,
	0xfa, 0xe7, 0xb0, 0xb3, 0x11, 0x9a, 0x6d, 0xa7, 0xf2, 0xba, 0xf1, 0x78, 0x0c, 0x3b, 0xab, 0x0f,
	0x59, 0xdd, 0x4e, 0xd8, 0x18, 0x0b, 0x7c, 0x7b, 0x6c, 0x31, 0xcf, 0x91, 0x49, 0x4a, 0x66, 0x4d,
	0x3a, 0xce, 0x3c, 0x47, 0x7c, 0x19, 0x71, 0xc4, 0x42, 0xcb, 0x75, 0xd2, 0xbb, 0x4a, 0x98, 0x5f,
	0x39, 0x46, 0x00, 0x37, 0xcf, 0xdd, 0x48, 0x65, 0xea, 0xe1, 0xc2, 0x5f, 0x25, 0xcc, 0xc4, 0x14,
	0xb2, 0x31, 0xe4, 0x36, 0xd4, 0x23, 0x4e, 0x43, 0x6e, 0xc9, 0xf4, 0x32, 0x61, 0x09, 0x1b, 0x17,
	0xae, 0x33, 0xe1, 0x11, 0x54, 0x02, 0x3a, 0x62, 0x56, 0x84, 0xd3, 0x25, 0x3b, 0x2f, 0x63, 0xe7,
	0xe8, 0xe8, 0xa3, 0x6d, 0x7c, 0x07, 0xad, 0xab, 0x15, 0x93, 0xbb, 0xe8, 0x21, 0x54, 0x64, 0x8f,
	0xe9, 0xb7, 0xf6, 0x76, 0x9e, 0xf6, 0x9b, 0x37, 0x98, 0x99, 0xc4, 0x90, 0x03, 0x00, 0x8f, 0x3d,
	0x5f, 0xa7, 0xa5, 0x09, 0x8f, 0x64, 0x65, 0xfc, 0x58, 0x80, 0xba, 0x0c, 0xfc, 0x36, 0xc0, 0x2f,
	0x9c, 0x91, 0x7b, 0x50, 0x89, 0xe5, 0x4a, 0x42, 0xeb, 0xdd, 0x5b, 0xaf, 0x3a, 0x68, 0x33, 0xc1,
	0xfe, 0xa7, 0xfb, 0xcf, 0x78, 0x06, 0x44, 0x15, 0x5f, 0x3b, 0xb9, 0x97, 0x0a, 0xfd, 0x08, 0x1a,
	0x4c, 0x00, 0xad, 0x35, 0x9e, 0xb7, 0xf3, 0xca, 0x65, 0x1a, 0x33, 0xeb, 0x6c, 0x65, 0x18, 0xdf,
	0xc0, 0x8d, 0xb5, 0x92, 0x89, 0xd2, 0xa7, 0x50, 0x56, 0x97, 0x70, 0x41, 0xe6, 0xbc, 0x9e, 0xd0,
	0x2a, 0xa4, 0xfb, 0xfb, 0x16, 0x34, 0xf1, 0x23, 0x7c, 0x9a, 0x81, 0xf7, 0x59, 0x38, 0x77, 0x6d,
	0x46, 0x42, 0xa8, 0xa5, 0x51, 0xe4, 0xe8, 0xd5, 0x39, 0x65, 0xef, 0xfa, 0xb5, 0x0a, 0x1b, 0xfb,
	0x3f, 0xfd, 0xf5, 0xf7, 0x6f, 0xc5, 0xff, 0x93, 0x1b, 0x1d, 0x7c, 0x21, 0x08, 0x71, 0xa2, 0xce,
	0x0f, 0x89, 0x64, 0x2f, 0xc8, 0xaf, 0xf8, 0x5f, 0x6d, 0x73, 0xa2, 0xc8, 0x07, 0x79, 0x79, 0x5f,
	0x32, 0xe9, 0xfa, 0xdd, 0xeb, 0x81, 0x13, 0x32, 0x47, 0x92, 0xcc, 0x01, 0xd9, 0xcf, 0x21, 0xd3,
	0x19, 0x27, 0xf5, 0x7f, 0xc6, 0x61, 0xcb, 0xe8, 0x4e, 0xee, 0xe4, 0x95, 0xb8, 0x3a, 0x0b, 0xfa,
	0xbb, 0xaf, 0xc5, 0x25, 0x2c, 0xde, 0x97, 0x2c, 0x8e, 0xf4, 0x3c, 0x49, 0x4e, 0xd7, 0xc6, 0x86,
	0x4c, 0x40, 0x5b, 0xbe, 0x65, 0x48, 0xae, 0xe0, 0x9b, 0x4f, 0x1d, 0xfd, 0x30, 0x17, 0x95, 0x79,
	0x82, 0x18, 0x4d, 0x59, 0x7f, 0x97, 0x6c, 0x8b, 0xfa, 0xe3, 0x49, 0xd0, 0x99, 0xca, 0x04, 0x83,
	0x8a, 0x7c, 0x40, 0x7d, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x42, 0x8b, 0xbc, 0xe2,
	0x09, 0x00, 0x00,
}
