// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package descriptor

import (
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/ligato/cn-infra/logging"
	scheduler "github.com/ligato/vpp-agent/plugins/kvscheduler/api"

	"github.com/go-errors/errors"
	"github.com/ligato/vpp-agent/plugins/vppv2/ifplugin/descriptor/adapter"
	"github.com/ligato/vpp-agent/plugins/vppv2/ifplugin/ifaceidx"
	"github.com/ligato/vpp-agent/plugins/vppv2/ifplugin/vppcalls"
	"github.com/ligato/vpp-agent/plugins/vppv2/model/interfaces"
)

const (
	// UnnumberedIfDescriptorName is the name of the descriptor for the unnumbered
	// config-subsection of VPP interfaces.
	UnnumberedIfDescriptorName = "vpp-unnumbered-interface"

	// dependency labels
	unnumberedInterfaceWithIPDep = "unnumbered-interface-with-IP"
)

// UnnumberedIfDescriptor sets/unsets VPP interfaces as unnumbered.
// Values = Interface_Unnumbered{} derived from interfaces where IsUnnumbered==true
type UnnumberedIfDescriptor struct {
	log       logging.Logger
	ifHandler vppcalls.IfVppAPI
	intfIndex ifaceidx.IfaceMetadataIndex
}

// NewUnnumberedIfDescriptor creates a new instance of UnnumberedIfDescriptor.
func NewUnnumberedIfDescriptor(ifHandler vppcalls.IfVppAPI, log logging.PluginLogger) *UnnumberedIfDescriptor {
	return &UnnumberedIfDescriptor{
		ifHandler: ifHandler,
		log:       log.NewLogger("unif-descriptor"),
	}
}

// GetDescriptor returns descriptor suitable for registration (via adapter)
// with the KVScheduler.
func (d *UnnumberedIfDescriptor) GetDescriptor() *adapter.UnnumberedDescriptor {
	return &adapter.UnnumberedDescriptor{
		Name:               UnnumberedIfDescriptorName,
		KeySelector:        d.IsUnnumberedInterfaceKey,
		ValueTypeName:      proto.MessageName(&interfaces.Interface_Unnumbered{}),
		Add:                d.Add,
		Delete:             d.Delete,
		ModifyWithRecreate: d.ModifyWithRecreate,
		Dependencies:       d.Dependencies,
	}
}

// SetInterfaceIndex should be used to provide interface index immediately after
// the descriptor registration.
func (d *UnnumberedIfDescriptor) SetInterfaceIndex(intfIndex ifaceidx.IfaceMetadataIndex) {
	d.intfIndex = intfIndex
}

// IsUnnumberedInterfaceKey returns true if the key is identifying unnumbered
// VPP interface.
func (d *UnnumberedIfDescriptor) IsUnnumberedInterfaceKey(key string) bool {
	return strings.HasPrefix(key, interfaces.UnnumberedKeyPrefix)
}

// Add sets interface as unnumbered.
func (d *UnnumberedIfDescriptor) Add(key string, unIntf *interfaces.Interface_Unnumbered) (metadata interface{}, err error) {
	ifName := strings.TrimPrefix(key, interfaces.UnnumberedKeyPrefix)
	ifMeta, found := d.intfIndex.LookupByName(ifName)
	if !found {
		err = errors.Errorf("failed to find unnumbered interface %s", ifName)
		d.log.Error(err)
		return nil, err
	}

	ifWithIPMeta, found := d.intfIndex.LookupByName(unIntf.InterfaceWithIp)
	if !found {
		err = errors.Errorf("failed to find interface %s referenced by unnumbered interface %s",
			unIntf.InterfaceWithIp, ifName)
		d.log.Error(err)
		return nil, err
	}

	err = d.ifHandler.SetUnnumberedIP(ifMeta.SwIfIndex, ifWithIPMeta.SwIfIndex)
	if err != nil {
		d.log.Error(err)
	}
	return nil, err
}

// Delete un-sets interface as unnumbered.
func (d *UnnumberedIfDescriptor) Delete(key string, unIntf *interfaces.Interface_Unnumbered, metadata interface{}) error {
	ifName := strings.TrimPrefix(key, interfaces.UnnumberedKeyPrefix)
	ifMeta, found := d.intfIndex.LookupByName(ifName)
	if !found {
		err := errors.Errorf("failed to find unnumbered interface %s", ifName)
		d.log.Error(err)
		return err
	}

	err := d.ifHandler.UnsetUnnumberedIP(ifMeta.SwIfIndex)
	if err != nil {
		d.log.Error(err)
	}
	return err
}

// ModifyWithRecreate returns always true so that the link to interface with IP
// address is reconfigured with UnsetUnnumberedIP followed by SetUnnumberedIP for the new interface.
func (d *UnnumberedIfDescriptor) ModifyWithRecreate(key string, oldUnIntf, newUnIntf *interfaces.Interface_Unnumbered, oldMetadata interface{}) bool {
	return true
}

// Dependencies lists dependencies for an unnumbered VPP interface.
func (d *UnnumberedIfDescriptor) Dependencies(key string, unIntf *interfaces.Interface_Unnumbered) []scheduler.Dependency {
	// link between unnumbered interface and the referenced interface with IP address
	// - satisfied as along as the referenced interface is configured and has at least
	//   one IP address assigned
	return []scheduler.Dependency{
		{
			Label: unnumberedInterfaceWithIPDep,
			AnyOf: func(key string) bool {
				ifName, _, _, isIfaceAddrKey := interfaces.ParseInterfaceAddressKey(key)
				return isIfaceAddrKey && ifName == unIntf.InterfaceWithIp
			},
		},
	}
}