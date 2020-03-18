// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/plugins/nsim.api.json

/*
Package nsim is a generated VPP binary API for 'nsim' module.

It consists of:
	  6 messages
	  3 services
*/
package nsim

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
	ModuleName = "nsim"
	// APIVersion is the API version of this module.
	APIVersion = "2.1.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x760acaa
)

// NsimConfigure represents VPP binary API message 'nsim_configure'.
type NsimConfigure struct {
	DelayInUsec              uint32
	AveragePacketSize        uint32
	BandwidthInBitsPerSecond uint64
	PacketsPerDrop           uint32
}

func (*NsimConfigure) GetMessageName() string {
	return "nsim_configure"
}
func (*NsimConfigure) GetCrcString() string {
	return "16ed400f"
}
func (*NsimConfigure) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// NsimConfigureReply represents VPP binary API message 'nsim_configure_reply'.
type NsimConfigureReply struct {
	Retval int32
}

func (*NsimConfigureReply) GetMessageName() string {
	return "nsim_configure_reply"
}
func (*NsimConfigureReply) GetCrcString() string {
	return "e8d4e804"
}
func (*NsimConfigureReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// NsimCrossConnectEnableDisable represents VPP binary API message 'nsim_cross_connect_enable_disable'.
type NsimCrossConnectEnableDisable struct {
	EnableDisable uint8
	SwIfIndex0    uint32
	SwIfIndex1    uint32
}

func (*NsimCrossConnectEnableDisable) GetMessageName() string {
	return "nsim_cross_connect_enable_disable"
}
func (*NsimCrossConnectEnableDisable) GetCrcString() string {
	return "df4e7ba9"
}
func (*NsimCrossConnectEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// NsimCrossConnectEnableDisableReply represents VPP binary API message 'nsim_cross_connect_enable_disable_reply'.
type NsimCrossConnectEnableDisableReply struct {
	Retval int32
}

func (*NsimCrossConnectEnableDisableReply) GetMessageName() string {
	return "nsim_cross_connect_enable_disable_reply"
}
func (*NsimCrossConnectEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*NsimCrossConnectEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// NsimOutputFeatureEnableDisable represents VPP binary API message 'nsim_output_feature_enable_disable'.
type NsimOutputFeatureEnableDisable struct {
	EnableDisable uint8
	SwIfIndex     uint32
}

func (*NsimOutputFeatureEnableDisable) GetMessageName() string {
	return "nsim_output_feature_enable_disable"
}
func (*NsimOutputFeatureEnableDisable) GetCrcString() string {
	return "57298519"
}
func (*NsimOutputFeatureEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// NsimOutputFeatureEnableDisableReply represents VPP binary API message 'nsim_output_feature_enable_disable_reply'.
type NsimOutputFeatureEnableDisableReply struct {
	Retval int32
}

func (*NsimOutputFeatureEnableDisableReply) GetMessageName() string {
	return "nsim_output_feature_enable_disable_reply"
}
func (*NsimOutputFeatureEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*NsimOutputFeatureEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*NsimConfigure)(nil), "nsim.NsimConfigure")
	api.RegisterMessage((*NsimConfigureReply)(nil), "nsim.NsimConfigureReply")
	api.RegisterMessage((*NsimCrossConnectEnableDisable)(nil), "nsim.NsimCrossConnectEnableDisable")
	api.RegisterMessage((*NsimCrossConnectEnableDisableReply)(nil), "nsim.NsimCrossConnectEnableDisableReply")
	api.RegisterMessage((*NsimOutputFeatureEnableDisable)(nil), "nsim.NsimOutputFeatureEnableDisable")
	api.RegisterMessage((*NsimOutputFeatureEnableDisableReply)(nil), "nsim.NsimOutputFeatureEnableDisableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*NsimConfigure)(nil),
		(*NsimConfigureReply)(nil),
		(*NsimCrossConnectEnableDisable)(nil),
		(*NsimCrossConnectEnableDisableReply)(nil),
		(*NsimOutputFeatureEnableDisable)(nil),
		(*NsimOutputFeatureEnableDisableReply)(nil),
	}
}

// RPCService represents RPC service API for nsim module.
type RPCService interface {
	NsimConfigure(ctx context.Context, in *NsimConfigure) (*NsimConfigureReply, error)
	NsimCrossConnectEnableDisable(ctx context.Context, in *NsimCrossConnectEnableDisable) (*NsimCrossConnectEnableDisableReply, error)
	NsimOutputFeatureEnableDisable(ctx context.Context, in *NsimOutputFeatureEnableDisable) (*NsimOutputFeatureEnableDisableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) NsimConfigure(ctx context.Context, in *NsimConfigure) (*NsimConfigureReply, error) {
	out := new(NsimConfigureReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) NsimCrossConnectEnableDisable(ctx context.Context, in *NsimCrossConnectEnableDisable) (*NsimCrossConnectEnableDisableReply, error) {
	out := new(NsimCrossConnectEnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) NsimOutputFeatureEnableDisable(ctx context.Context, in *NsimOutputFeatureEnableDisable) (*NsimOutputFeatureEnableDisableReply, error) {
	out := new(NsimOutputFeatureEnableDisableReply)
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
