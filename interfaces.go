// Copyright (C) 2019 Cisco Systems Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vpplink

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
	"github.com/calico-vpp/vpplink/binapi/19_08/interfaces"
	vppip "github.com/calico-vpp/vpplink/binapi/19_08/ip"
	"github.com/calico-vpp/vpplink/binapi/19_08/tapv2"
)

const (
	INVALID_INDEX = ^uint32(0)
)

func (v *VppLink) CreateTap(
	ContNS string,
	ContIfName string,
	Tag string,
	EnableIp6 bool,
	macAddress [6]byte,
	hostMacAddress [6]byte,
) (SwIfIndex uint32, vppIPAddress []byte, err error) {
	response := &tapv2.TapCreateV2Reply{}
	request := &tapv2.TapCreateV2{
		// TODO check namespace len < 64?
		// TODO set MTU?
		ID:               ^uint32(0),
		HostNamespace:    []byte(ContNS),
		HostNamespaceSet: 1,
		HostIfName:       []byte(ContIfName),
		HostIfNameSet:    1,
		Tag:              []byte(Tag),
		MacAddress:       macAddress[:],
		HostMacAddr:      hostMacAddress[:],
		HostMacAddrSet:   1,
	}

	v.lock.Lock()
	v.log.Debugf("Tap creation request: %+v", request)
	err = v.ch.SendRequest(request).ReceiveReply(response)
	v.log.Infof("Tap creation: err %v retval %d sw_if_index = %d", err, response.Retval, response.SwIfIndex)
	v.lock.Unlock()
	if err != nil {
		return INVALID_INDEX, vppIPAddress, errors.Wrap(err, "Tap creation request failed")
	} else if response.Retval != 0 {
		return INVALID_INDEX, vppIPAddress, fmt.Errorf("Tap creation failed (retval %d). Request: %+v", response.Retval, request)
	}

	// Add VPP side fake address
	// TODO: Only if v4 is enabled
	// There is currently a hard limit in VPP to 1024 taps - so this should be safe
	vppIPAddress = []byte{169, 254, byte(response.SwIfIndex >> 8), byte(response.SwIfIndex)}
	err = v.AddInterfaceAddress(response.SwIfIndex, vppIPAddress, 32)
	if err != nil {
		return INVALID_INDEX, vppIPAddress, errors.Wrap(err, "error adding address to new tap")
	}

	// Set interface up
	err = v.InterfaceAdminUp(response.SwIfIndex)
	if err != nil {
		return INVALID_INDEX, vppIPAddress, errors.Wrap(err, "error setting new tap up")
	}

	// Add IPv6 neighbor entry if v6 is enabled
	if EnableIp6 {
		err = v.EnableInterfaceIP6(response.SwIfIndex)
		if err != nil {
			return INVALID_INDEX, vppIPAddress, errors.Wrap(err, "error enabling IPv6 on new tap")
		}
		// Compute a link local address from mac address, and set it
	}
	return response.SwIfIndex, vppIPAddress, err
}

func (v *VppLink) addDelInterfaceAddress(swIfIndex uint32, addr []byte, addrLen uint8, isAdd uint8) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &interfaces.SwInterfaceAddDelAddress{
		SwIfIndex:     swIfIndex,
		IsAdd:         isAdd,
		AddressLength: addrLen,
		Address:       addr,
	}
	response := &interfaces.SwInterfaceAddDelAddressReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrapf(err, "Adding IP address failed: req %+v reply %+v", request, response)
	}
	return nil
}

func (v *VppLink) DelInterfaceAddress(swIfIndex uint32, addr []byte, addrLen uint8) error {
	return v.addDelInterfaceAddress(swIfIndex, addr, addrLen, 0)
}

func (v *VppLink) AddInterfaceAddress(swIfIndex uint32, addr []byte, addrLen uint8) error {
	return v.addDelInterfaceAddress(swIfIndex, addr, addrLen, 1)
}

func (v *VppLink) enableDisableInterfaceIP6(swIfIndex uint32, enable uint8) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &vppip.SwInterfaceIP6EnableDisable{
		SwIfIndex: swIfIndex,
		Enable:    enable,
	}
	response := &vppip.SwInterfaceIP6EnableDisableReply{}
	return v.ch.SendRequest(request).ReceiveReply(response)
}

func (v *VppLink) DisableInterfaceIP6(swIfIndex uint32) error {
	return v.enableDisableInterfaceIP6(swIfIndex, 0)
}

func (v *VppLink) EnableInterfaceIP6(swIfIndex uint32) error {
	return v.enableDisableInterfaceIP6(swIfIndex, 1)
}

func (v *VppLink) SearchInterfaceWithTag(tag string) (err error, swIfIndex uint32) {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &interfaces.SwInterfaceDump{
		//NameFilterValid: true,
		//NameFilter:      "tap",
	}
	response := &interfaces.SwInterfaceDetails{}
	stream := v.ch.SendMultiRequest(request)
	for {
		stop, err := stream.ReceiveReply(response)
		if err != nil {
			v.log.Errorf("error listing VPP interfaces: %v", err)
			return err, ^uint32(0)
		}
		if stop {
			v.log.Errorf("error: interface to delete not found")
			return fmt.Errorf("VPP Error: interface to delete not found"), ^uint32(0)
		}
		intfTag := string(bytes.Trim([]byte(response.Tag), "\x00"))
		v.log.Debugf("found interface %d, tag: %s (len %d)", response.SwIfIndex, intfTag, len(intfTag))
		if intfTag == tag {
			return nil, response.SwIfIndex
		}
	}
}

func (v *VppLink) SearchInterfaceWithName(name string) (err error, swIfIndex uint32) {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &interfaces.SwInterfaceDump{
		SwIfIndex: interfaces.InterfaceIndex(^uint32(0)),
		// TODO: filter by name with NameFilter
	}
	reqCtx := v.ch.SendMultiRequest(request)
	for {
		response := &interfaces.SwInterfaceDetails{}
		stop, err := reqCtx.ReceiveReply(response)
		if err != nil {
			v.log.Errorf("SwInterfaceDump failed: %v", err)
			return err, 0
		}
		if stop {
			break
		}
		interfaceName := string(bytes.Trim([]byte(response.InterfaceName), "\x00"))
		v.log.Debugf("Found interface: %s", interfaceName)
		if interfaceName == name {
			return nil, response.SwIfIndex
		}

	}
	v.log.Errorf("Interface %s not found", name)
	return errors.New("Interface not found"), 0
}

func (v *VppLink) interfaceAdminUpDown(swIfIndex uint32, updown uint8) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	// Set interface down
	request := &interfaces.SwInterfaceSetFlags{
		SwIfIndex:   swIfIndex,
		AdminUpDown: updown,
	}
	response := &interfaces.SwInterfaceSetFlagsReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrapf(err, "setting interface down failed")
	}
	return nil
}

func (v *VppLink) InterfaceAdminDown(swIfIndex uint32) error {
	return v.interfaceAdminUpDown(swIfIndex, 0)
}

func (v *VppLink) InterfaceAdminUp(swIfIndex uint32) error {
	return v.interfaceAdminUpDown(swIfIndex, 1)
}

func (v *VppLink) GetInterfaceNeighbors(swIfIndex uint32, isIPv6 uint8) (err error, neighbors []vppip.IPNeighbor) {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &vppip.IPNeighborDump{
		SwIfIndex: swIfIndex,
		IsIPv6:    isIPv6,
	}
	response := &vppip.IPNeighborDetails{}
	stream := v.ch.SendMultiRequest(request)
	for {
		stop, err := stream.ReceiveReply(response)
		if err != nil {
			v.log.Errorf("error listing VPP neighbors: %v", err)
			return err, nil
		}
		if stop {
			return nil, neighbors
		}
		neighbors = append(neighbors, response.Neighbor)
	}
}

func (v *VppLink) DelTap(swIfIndex uint32) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &tapv2.TapDeleteV2{
		SwIfIndex: swIfIndex,
	}
	response := &tapv2.TapDeleteV2Reply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "failed to delete tap from VPP")
	}
	return nil
}

func (v *VppLink) interfaceSetUnnumbered(unnumberedSwIfIndex uint32, swIfIndex uint32, isAdd uint8) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	// Set interface down
	request := &interfaces.SwInterfaceSetUnnumbered{
		SwIfIndex:   swIfIndex,
		UnnumberedSwIfIndex:   unnumberedSwIfIndex,
		IsAdd: isAdd,
	}
	response := &interfaces.SwInterfaceSetUnnumberedReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrapf(err, "setting interface unnumbered failed %d -> %d", unnumberedSwIfIndex, swIfIndex)
	}
	return nil
}

func (v *VppLink) InterfaceSetUnnumbered(unnumberedSwIfIndex uint32, swIfIndex uint32) error {
	return v.interfaceSetUnnumbered(unnumberedSwIfIndex, swIfIndex, 1)
}

func (v *VppLink) InterfaceUnsetUnnumbered(unnumberedSwIfIndex uint32, swIfIndex uint32) error {
	return v.interfaceSetUnnumbered(unnumberedSwIfIndex, swIfIndex, 0)
}