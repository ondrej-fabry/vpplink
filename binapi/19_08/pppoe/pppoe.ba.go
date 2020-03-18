// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/plugins/pppoe.api.json

/*
Package pppoe is a generated VPP binary API for 'pppoe' module.

It consists of:
	  4 messages
	  2 services
*/
package pppoe

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
	ModuleName = "pppoe"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x4def67c4
)

// PppoeAddDelSession represents VPP binary API message 'pppoe_add_del_session'.
type PppoeAddDelSession struct {
	IsAdd      uint8
	IsIPv6     uint8
	SessionID  uint16
	ClientIP   []byte `struc:"[16]byte"`
	DecapVrfID uint32
	ClientMac  []byte `struc:"[6]byte"`
}

func (*PppoeAddDelSession) GetMessageName() string {
	return "pppoe_add_del_session"
}
func (*PppoeAddDelSession) GetCrcString() string {
	return "766cbfeb"
}
func (*PppoeAddDelSession) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PppoeAddDelSessionReply represents VPP binary API message 'pppoe_add_del_session_reply'.
type PppoeAddDelSessionReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*PppoeAddDelSessionReply) GetMessageName() string {
	return "pppoe_add_del_session_reply"
}
func (*PppoeAddDelSessionReply) GetCrcString() string {
	return "fda5941f"
}
func (*PppoeAddDelSessionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// PppoeSessionDetails represents VPP binary API message 'pppoe_session_details'.
type PppoeSessionDetails struct {
	SwIfIndex    uint32
	IsIPv6       uint8
	SessionID    uint16
	ClientIP     []byte `struc:"[16]byte"`
	EncapIfIndex uint32
	DecapVrfID   uint32
	LocalMac     []byte `struc:"[6]byte"`
	ClientMac    []byte `struc:"[6]byte"`
}

func (*PppoeSessionDetails) GetMessageName() string {
	return "pppoe_session_details"
}
func (*PppoeSessionDetails) GetCrcString() string {
	return "358fc7a8"
}
func (*PppoeSessionDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// PppoeSessionDump represents VPP binary API message 'pppoe_session_dump'.
type PppoeSessionDump struct {
	SwIfIndex uint32
}

func (*PppoeSessionDump) GetMessageName() string {
	return "pppoe_session_dump"
}
func (*PppoeSessionDump) GetCrcString() string {
	return "529cb13f"
}
func (*PppoeSessionDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func init() {
	api.RegisterMessage((*PppoeAddDelSession)(nil), "pppoe.PppoeAddDelSession")
	api.RegisterMessage((*PppoeAddDelSessionReply)(nil), "pppoe.PppoeAddDelSessionReply")
	api.RegisterMessage((*PppoeSessionDetails)(nil), "pppoe.PppoeSessionDetails")
	api.RegisterMessage((*PppoeSessionDump)(nil), "pppoe.PppoeSessionDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PppoeAddDelSession)(nil),
		(*PppoeAddDelSessionReply)(nil),
		(*PppoeSessionDetails)(nil),
		(*PppoeSessionDump)(nil),
	}
}

// RPCService represents RPC service API for pppoe module.
type RPCService interface {
	DumpPppoeSession(ctx context.Context, in *PppoeSessionDump) (RPCService_DumpPppoeSessionClient, error)
	PppoeAddDelSession(ctx context.Context, in *PppoeAddDelSession) (*PppoeAddDelSessionReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpPppoeSession(ctx context.Context, in *PppoeSessionDump) (RPCService_DumpPppoeSessionClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpPppoeSessionClient{stream}
	return x, nil
}

type RPCService_DumpPppoeSessionClient interface {
	Recv() (*PppoeSessionDetails, error)
}

type serviceClient_DumpPppoeSessionClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpPppoeSessionClient) Recv() (*PppoeSessionDetails, error) {
	m := new(PppoeSessionDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) PppoeAddDelSession(ctx context.Context, in *PppoeAddDelSession) (*PppoeAddDelSessionReply, error) {
	out := new(PppoeAddDelSessionReply)
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
