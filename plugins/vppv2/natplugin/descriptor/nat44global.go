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
	"github.com/gogo/protobuf/proto"
	"github.com/go-errors/errors"

	"github.com/ligato/cn-infra/logging"

	scheduler "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/plugins/vppv2/natplugin/descriptor/adapter"
	vpp_ifdescriptor "github.com/ligato/vpp-agent/plugins/vppv2/ifplugin/descriptor"
	"github.com/ligato/vpp-agent/plugins/vppv2/natplugin/vppcalls"
	"github.com/ligato/vpp-agent/plugins/vppv2/model/nat"
)

const (
	// NAT44GlobalDescriptorName is the name of the descriptor for VPP NAT44 global
	// configuration.
	NAT44GlobalDescriptorName = "vpp-nat44-global"

	// default virtual reassembly configuration
	natReassTimeoutDefault = 2 // seconds
	natMaxReassDefault = 1024
	natMaxFragDefault = 5
	natDropFragDefault = false
)

// A list of non-retriable errors:
var (
	// ErrNATInterfaceFeatureCollision is returned when VPP NAT features assigned
	// to a single interface collide.
	ErrNATInterfaceFeatureCollision = errors.New("VPP NAT interface feature collision")
)

// defaultGlobalCfg is the default NAT44 global configuration.
var defaultGlobalCfg = &nat.Nat44Global{
	VirtualReassemblyIpv4: &nat.Nat44Global_VirtualReassembly{
		Timeout:  natReassTimeoutDefault,
		MaxReass: natMaxReassDefault,
		MaxFrag:  natMaxFragDefault,
		DropFrag: natDropFragDefault,
	},
	VirtualReassemblyIpv6: &nat.Nat44Global_VirtualReassembly{
		Timeout:  natReassTimeoutDefault,
		MaxReass: natMaxReassDefault,
		MaxFrag:  natMaxFragDefault,
		DropFrag: natDropFragDefault,
	},
}

// NAT44GlobalDescriptor teaches KVScheduler how to configure global options for
// VPP NAT44.
type NAT44GlobalDescriptor struct {
	log        logging.Logger
	natHandler vppcalls.NatVppAPI
}

// NewNAT44GlobalDescriptor creates a new instance of the NAT44Global descriptor.
func NewNAT44GlobalDescriptor(natHandler vppcalls.NatVppAPI, log logging.PluginLogger) *NAT44GlobalDescriptor {

	return &NAT44GlobalDescriptor{
		natHandler: natHandler,
		log:        log.NewLogger("nat-descriptor"),
	}
}

// GetDescriptor returns descriptor suitable for registration (via adapter) with
// the KVScheduler.
func (d *NAT44GlobalDescriptor) GetDescriptor() *adapter.NAT44GlobalDescriptor {
	return &adapter.NAT44GlobalDescriptor{
		Name:               NAT44GlobalDescriptorName,
		KeySelector:        d.IsNAT44GlobalKey,
		ValueTypeName:      proto.MessageName(&nat.Nat44Global{}),
		NBKeyPrefix:        nat.Prefix,
		Add:                d.Add,
		Delete:             d.Delete,
		Modify:             d.Modify,
		IsRetriableFailure: d.IsRetriableFailure,
		DerivedValues:      d.DerivedValues,
		Dump:               d.Dump,
		DumpDependencies:   []string{vpp_ifdescriptor.InterfaceDescriptorName},
	}
}

// IsNAT44GlobalKey returns true if the key is identifying global VPP NAT44 options.
func (d *NAT44GlobalDescriptor) IsNAT44GlobalKey(key string) bool {
	return key == nat.GlobalKey
}

// Add applies NAT44 global options.
func (d *NAT44GlobalDescriptor) Add(key string, globalCfg *nat.Nat44Global) (metadata interface{}, err error) {
	return d.Modify(key, defaultGlobalCfg, globalCfg, nil)
}

// Delete sets NAT44 global options back to the defaults.
func (d *NAT44GlobalDescriptor) Delete(key string, globalCfg *nat.Nat44Global, metadata interface{}) error {
	_, err := d.Modify(key, globalCfg, defaultGlobalCfg, metadata)
	return err
}

// Modify updates NAT44 global options.
func (d *NAT44GlobalDescriptor) Modify(key string, oldGlobalCfg, newGlobalCfg *nat.Nat44Global, oldMetadata interface{}) (newMetadata interface{}, err error) {
	// validate configuration first
	err = d.validateNAT44GlobalConfig(newGlobalCfg)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	// update forwarding
	if oldGlobalCfg.Forwarding != newGlobalCfg.Forwarding {
		if err = d.natHandler.SetNat44Forwarding(newGlobalCfg.Forwarding); err != nil {
			err = errors.Errorf("failed to set NAT44 forwarding to %t: %v", newGlobalCfg.Forwarding, err)
			d.log.Error(err)
			return nil, err
		}
	}

	// update virtual reassembly for IPv4
	if !proto.Equal(getVirtualReassembly(oldGlobalCfg, false), getVirtualReassembly(newGlobalCfg, false)) {
		if err = d.natHandler.SetVirtualReassemblyIPv4(getVirtualReassembly(newGlobalCfg, false)); err != nil {
			err = errors.Errorf("failed to set NAT virtual reassembly for IPv4: %v", err)
			d.log.Error(err)
			return nil, err
		}
	}

	// update virtual reassembly for IPv6
	if !proto.Equal(getVirtualReassembly(oldGlobalCfg, true), getVirtualReassembly(newGlobalCfg, true)) {
		if err = d.natHandler.SetVirtualReassemblyIPv6(getVirtualReassembly(newGlobalCfg, false)); err != nil {
			err = errors.Errorf("failed to set NAT virtual reassembly for IPv6: %v", err)
			d.log.Error(err)
			return nil, err
		}
	}

	// remove obsolete addresses from the pool
	for _, oldAddr := range oldGlobalCfg.AddressPool {
		found := false
		for _, newAddr := range newGlobalCfg.AddressPool {
			if proto.Equal(oldAddr, newAddr) {
				found = true
				break
			}
		}
		if !found {
			if err = d.natHandler.DelNat44Address(oldAddr.Address, oldAddr.VrfId, oldAddr.TwiceNat); err != nil {
				err = errors.Errorf("failed to remove address %s from the NAT pool: %v", oldAddr.Address, err)
				d.log.Error(err)
				return nil, err
			}
		}
	}

	// add new addresses into the pool
	for _, newAddr := range newGlobalCfg.AddressPool {
		found := false
		for _, oldAddr := range oldGlobalCfg.AddressPool {
			if proto.Equal(oldAddr, newAddr) {
				found = true
				break
			}
		}
		if !found {
			if err = d.natHandler.AddNat44Address(newAddr.Address, newAddr.VrfId, newAddr.TwiceNat); err != nil {
				err = errors.Errorf("failed to add address %s into the NAT pool: %v", newAddr.Address, err)
				d.log.Error(err)
				return nil, err
			}
		}
	}

	return nil, nil
}

// IsRetriableFailure returns <false> for errors related to invalid configuration.
func (d *NAT44GlobalDescriptor) IsRetriableFailure(err error) bool {
	return err != ErrNATInterfaceFeatureCollision
}

// DerivedValues derives nat.NatInterface for every interface with assigned NAT configuration.
func (d *NAT44GlobalDescriptor) DerivedValues(key string, globalCfg *nat.Nat44Global) (derValues []scheduler.KeyValuePair) {
	// NAT interfaces
	for _, natIface := range globalCfg.NatInterfaces {
		derValues = append(derValues, scheduler.KeyValuePair{
			Key:   nat.NATInterfaceKey(natIface.Name, natIface.IsInside),
			Value: natIface,
		})
	}
	return derValues
}

// Dump returns the current NAT44 global configuration.
func (d *NAT44GlobalDescriptor) Dump(correlate []adapter.NAT44GlobalKVWithMetadata) ([]adapter.NAT44GlobalKVWithMetadata, error) {
	globalCfg, err := d.natHandler.Nat44GlobalConfigDump()
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	origin := scheduler.FromNB
	if proto.Equal(globalCfg, defaultGlobalCfg) {
		origin = scheduler.FromSB
	}

	dump := []adapter.NAT44GlobalKVWithMetadata{
		{
			Key:    nat.GlobalKey,
			Value:  globalCfg,
			Origin: origin,
		},
	}

	d.log.Debugf("Dumping NAT44 global configuration: %v", globalCfg)
	return dump, nil
}

// natIface accumulates NAT interface configuration for validation purposes.
type natIface struct {
	// feature assignment counters
	in     int
	out    int
	output int
}

// validateNAT44GlobalConfig validates VPP NAT44 global configuration.
func (d *NAT44GlobalDescriptor) validateNAT44GlobalConfig(globalCfg *nat.Nat44Global) error {
	// check NAT interface features for collisions
	natIfaceMap := make(map[string]*natIface)
	for _, natIface := range globalCfg.NatInterfaces {
		if _, hasEntry := natIfaceMap[natIface.Name]; !hasEntry {
			natIfaceMap[natIface.Name] = &natIface{}
		}
		ifaceCfg := natIfaceMap[natIface.Name]
		if natIface.IsInside {
			ifaceCfg.in += 1
		} else {
			ifaceCfg.out += 1
		}
		if natIface.OutputFeature {
			ifaceCfg.output += 1
		}
	}
	for _, ifaceCfg := range natIfaceMap {
		if ifaceCfg.in > 1 {
			// duplicate IN
			return ErrNATInterfaceFeatureCollision
		}
		if ifaceCfg.out > 1 {
			// duplicate OUT
			return ErrNATInterfaceFeatureCollision
		}
		if ifaceCfg.output == 1 && (ifaceCfg.in + ifaceCfg.out > 1) {
			// OUTPUT interface cannot be both IN and OUT
			return ErrNATInterfaceFeatureCollision
		}
	}
	return nil
}

func getVirtualReassembly(globalCfg *nat.Nat44Global, ipv6 bool) *nat.Nat44Global_VirtualReassembly {
	if ipv6 {
		if globalCfg.VirtualReassemblyIpv6 == nil {
			return defaultGlobalCfg.VirtualReassemblyIpv6
		}
		return globalCfg.VirtualReassemblyIpv6
	}
	if globalCfg.VirtualReassemblyIpv4 == nil {
		return defaultGlobalCfg.VirtualReassemblyIpv4
	}
	return globalCfg.VirtualReassemblyIpv4
}