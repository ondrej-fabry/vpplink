// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/tapv2.api.json

/*
Package tapv2 is a generated VPP binary API for 'tapv2' module.

It consists of:
	  6 messages
	  3 services
*/
package tapv2

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
	ModuleName = "tapv2"
	// APIVersion is the API version of this module.
	APIVersion = "2.1.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x25beb6c0
)

// SwInterfaceTapV2Details represents VPP binary API message 'sw_interface_tap_v2_details'.
type SwInterfaceTapV2Details struct {
	SwIfIndex        uint32
	ID               uint32
	DevName          []byte `struc:"[64]byte"`
	TxRingSz         uint16
	RxRingSz         uint16
	HostMacAddr      []byte `struc:"[6]byte"`
	HostIfName       []byte `struc:"[64]byte"`
	HostNamespace    []byte `struc:"[64]byte"`
	HostBridge       []byte `struc:"[64]byte"`
	HostIP4Addr      []byte `struc:"[4]byte"`
	HostIP4PrefixLen uint8
	HostIP6Addr      []byte `struc:"[16]byte"`
	HostIP6PrefixLen uint8
	HostMtuSize      uint32
	TapFlags         uint32
}

func (*SwInterfaceTapV2Details) GetMessageName() string {
	return "sw_interface_tap_v2_details"
}
func (*SwInterfaceTapV2Details) GetCrcString() string {
	return "5ee87a5f"
}
func (*SwInterfaceTapV2Details) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceTapV2Dump represents VPP binary API message 'sw_interface_tap_v2_dump'.
type SwInterfaceTapV2Dump struct{}

func (*SwInterfaceTapV2Dump) GetMessageName() string {
	return "sw_interface_tap_v2_dump"
}
func (*SwInterfaceTapV2Dump) GetCrcString() string {
	return "51077d14"
}
func (*SwInterfaceTapV2Dump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapCreateV2 represents VPP binary API message 'tap_create_v2'.
type TapCreateV2 struct {
	ID               uint32
	UseRandomMac     uint8
	MacAddress       []byte `struc:"[6]byte"`
	TxRingSz         uint16
	RxRingSz         uint16
	HostNamespaceSet uint8
	HostNamespace    []byte `struc:"[64]byte"`
	HostMacAddrSet   uint8
	HostMacAddr      []byte `struc:"[6]byte"`
	HostIfNameSet    uint8
	HostIfName       []byte `struc:"[64]byte"`
	HostBridgeSet    uint8
	HostBridge       []byte `struc:"[64]byte"`
	HostIP4AddrSet   uint8
	HostIP4Addr      []byte `struc:"[4]byte"`
	HostIP4PrefixLen uint8
	HostIP6AddrSet   uint8
	HostIP6Addr      []byte `struc:"[16]byte"`
	HostIP6PrefixLen uint8
	HostIP4GwSet     uint8
	HostIP4Gw        []byte `struc:"[4]byte"`
	HostIP6GwSet     uint8
	HostIP6Gw        []byte `struc:"[16]byte"`
	HostMtuSet       uint8
	HostMtuSize      uint32
	Tag              []byte `struc:"[64]byte"`
	TapFlags         uint32
}

func (*TapCreateV2) GetMessageName() string {
	return "tap_create_v2"
}
func (*TapCreateV2) GetCrcString() string {
	return "8fa99320"
}
func (*TapCreateV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapCreateV2Reply represents VPP binary API message 'tap_create_v2_reply'.
type TapCreateV2Reply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*TapCreateV2Reply) GetMessageName() string {
	return "tap_create_v2_reply"
}
func (*TapCreateV2Reply) GetCrcString() string {
	return "fda5941f"
}
func (*TapCreateV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// TapDeleteV2 represents VPP binary API message 'tap_delete_v2'.
type TapDeleteV2 struct {
	SwIfIndex uint32
}

func (*TapDeleteV2) GetMessageName() string {
	return "tap_delete_v2"
}
func (*TapDeleteV2) GetCrcString() string {
	return "529cb13f"
}
func (*TapDeleteV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapDeleteV2Reply represents VPP binary API message 'tap_delete_v2_reply'.
type TapDeleteV2Reply struct {
	Retval int32
}

func (*TapDeleteV2Reply) GetMessageName() string {
	return "tap_delete_v2_reply"
}
func (*TapDeleteV2Reply) GetCrcString() string {
	return "e8d4e804"
}
func (*TapDeleteV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*SwInterfaceTapV2Details)(nil), "tapv2.SwInterfaceTapV2Details")
	api.RegisterMessage((*SwInterfaceTapV2Dump)(nil), "tapv2.SwInterfaceTapV2Dump")
	api.RegisterMessage((*TapCreateV2)(nil), "tapv2.TapCreateV2")
	api.RegisterMessage((*TapCreateV2Reply)(nil), "tapv2.TapCreateV2Reply")
	api.RegisterMessage((*TapDeleteV2)(nil), "tapv2.TapDeleteV2")
	api.RegisterMessage((*TapDeleteV2Reply)(nil), "tapv2.TapDeleteV2Reply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SwInterfaceTapV2Details)(nil),
		(*SwInterfaceTapV2Dump)(nil),
		(*TapCreateV2)(nil),
		(*TapCreateV2Reply)(nil),
		(*TapDeleteV2)(nil),
		(*TapDeleteV2Reply)(nil),
	}
}

// RPCService represents RPC service API for tapv2 module.
type RPCService interface {
	DumpSwInterfaceTapV2(ctx context.Context, in *SwInterfaceTapV2Dump) (RPCService_DumpSwInterfaceTapV2Client, error)
	TapCreateV2(ctx context.Context, in *TapCreateV2) (*TapCreateV2Reply, error)
	TapDeleteV2(ctx context.Context, in *TapDeleteV2) (*TapDeleteV2Reply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpSwInterfaceTapV2(ctx context.Context, in *SwInterfaceTapV2Dump) (RPCService_DumpSwInterfaceTapV2Client, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpSwInterfaceTapV2Client{stream}
	return x, nil
}

type RPCService_DumpSwInterfaceTapV2Client interface {
	Recv() (*SwInterfaceTapV2Details, error)
}

type serviceClient_DumpSwInterfaceTapV2Client struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpSwInterfaceTapV2Client) Recv() (*SwInterfaceTapV2Details, error) {
	m := new(SwInterfaceTapV2Details)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) TapCreateV2(ctx context.Context, in *TapCreateV2) (*TapCreateV2Reply, error) {
	out := new(TapCreateV2Reply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) TapDeleteV2(ctx context.Context, in *TapDeleteV2) (*TapDeleteV2Reply, error) {
	out := new(TapDeleteV2Reply)
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
