// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/policer.api.json

/*
Package policer is a generated VPP binary API for 'policer' module.

It consists of:
	  4 messages
	  2 services
*/
package policer

import (
	bytes "bytes"
	context "context"
	io "io"
	strconv "strconv"

	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "policer"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x7807cc8d
)

// PolicerAddDel represents VPP binary API message 'policer_add_del'.
type PolicerAddDel struct {
	IsAdd             uint8
	Name              []byte `struc:"[64]byte"`
	Cir               uint32
	Eir               uint32
	Cb                uint64
	Eb                uint64
	RateType          uint8
	RoundType         uint8
	Type              uint8
	ColorAware        uint8
	ConformActionType uint8
	ConformDscp       uint8
	ExceedActionType  uint8
	ExceedDscp        uint8
	ViolateActionType uint8
	ViolateDscp       uint8
}

func (*PolicerAddDel) GetMessageName() string {
	return "policer_add_del"
}
func (*PolicerAddDel) GetCrcString() string {
	return "dfea2be8"
}
func (*PolicerAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PolicerAddDelReply represents VPP binary API message 'policer_add_del_reply'.
type PolicerAddDelReply struct {
	Retval       int32
	PolicerIndex uint32
}

func (*PolicerAddDelReply) GetMessageName() string {
	return "policer_add_del_reply"
}
func (*PolicerAddDelReply) GetCrcString() string {
	return "a177cef2"
}
func (*PolicerAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// PolicerDetails represents VPP binary API message 'policer_details'.
type PolicerDetails struct {
	Name               []byte `struc:"[64]byte"`
	Cir                uint32
	Eir                uint32
	Cb                 uint64
	Eb                 uint64
	RateType           uint8
	RoundType          uint8
	Type               uint8
	ConformActionType  uint8
	ConformDscp        uint8
	ExceedActionType   uint8
	ExceedDscp         uint8
	ViolateActionType  uint8
	ViolateDscp        uint8
	SingleRate         uint8
	ColorAware         uint8
	Scale              uint32
	CirTokensPerPeriod uint32
	PirTokensPerPeriod uint32
	CurrentLimit       uint32
	CurrentBucket      uint32
	ExtendedLimit      uint32
	ExtendedBucket     uint32
	LastUpdateTime     uint64
}

func (*PolicerDetails) GetMessageName() string {
	return "policer_details"
}
func (*PolicerDetails) GetCrcString() string {
	return "ff2765f0"
}
func (*PolicerDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// PolicerDump represents VPP binary API message 'policer_dump'.
type PolicerDump struct {
	MatchNameValid uint8
	MatchName      []byte `struc:"[64]byte"`
}

func (*PolicerDump) GetMessageName() string {
	return "policer_dump"
}
func (*PolicerDump) GetCrcString() string {
	return "8be04d34"
}
func (*PolicerDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func init() {
	api.RegisterMessage((*PolicerAddDel)(nil), "policer.PolicerAddDel")
	api.RegisterMessage((*PolicerAddDelReply)(nil), "policer.PolicerAddDelReply")
	api.RegisterMessage((*PolicerDetails)(nil), "policer.PolicerDetails")
	api.RegisterMessage((*PolicerDump)(nil), "policer.PolicerDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PolicerAddDel)(nil),
		(*PolicerAddDelReply)(nil),
		(*PolicerDetails)(nil),
		(*PolicerDump)(nil),
	}
}

// RPCService represents RPC service API for policer module.
type RPCService interface {
	DumpPolicer(ctx context.Context, in *PolicerDump) (RPCService_DumpPolicerClient, error)
	PolicerAddDel(ctx context.Context, in *PolicerAddDel) (*PolicerAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpPolicer(ctx context.Context, in *PolicerDump) (RPCService_DumpPolicerClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpPolicerClient{stream}
	return x, nil
}

type RPCService_DumpPolicerClient interface {
	Recv() (*PolicerDetails, error)
}

type serviceClient_DumpPolicerClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpPolicerClient) Recv() (*PolicerDetails, error) {
	m := new(PolicerDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) PolicerAddDel(ctx context.Context, in *PolicerAddDel) (*PolicerAddDelReply, error) {
	out := new(PolicerAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
