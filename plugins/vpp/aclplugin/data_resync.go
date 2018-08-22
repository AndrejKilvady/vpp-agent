// Copyright (c) 2017 Cisco and/or its affiliates.
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

package aclplugin

import (
	"github.com/ligato/vpp-agent/plugins/vpp/model/acl"
)

// Resync writes ACLs to the empty VPP.
func (plugin *ACLConfigurator) Resync(nbACLs []*acl.AccessLists_Acl) error {
	plugin.log.Debug("Resync ACLs started")
	// Calculate and log acl resync.
	defer func() {
		if plugin.stopwatch != nil {
			plugin.stopwatch.PrintLog()
		}
	}()

	// Re-initialize cache
	plugin.clearMapping()

	// Retrieve existing IpACL config
	vppIPACLs, err := plugin.aclHandler.DumpIPACL(plugin.ifIndexes)
	if err != nil {
		return err
	}
	// Retrieve existing MacIpACL config
	vppMacIPACLs, err := plugin.aclHandler.DumpMACIPACL(plugin.ifIndexes)
	if err != nil {
		return err
	}

	// Remove all configured VPP ACLs
	// Note: due to inability to dump ACL interfaces, it is not currently possible to correctly
	// calculate difference between configs
	for _, vppIPACL := range vppIPACLs {

		// ACL with IP-type rules uses different binary call to create/remove than MACIP-type.
		// Check what type of rules is in the ACL
		ipRulesExist := len(vppIPACL.ACL.Rules) > 0 && vppIPACL.ACL.Rules[0].GetMatch().GetIpRule() != nil

		if ipRulesExist {
			if err := plugin.aclHandler.DeleteIPAcl(vppIPACL.Meta.Index); err != nil {
				plugin.log.Error(err)
				return err
			}
			// Unregister.
			plugin.l3l4AclIndexes.UnregisterName(vppIPACL.ACL.AclName)
			continue
		}
	}
	for _, vppMacIPACL := range vppMacIPACLs {
		ipRulesExist := len(vppMacIPACL.ACL.Rules) > 0 && vppMacIPACL.ACL.Rules[0].GetMatch().GetMacipRule() != nil

		if ipRulesExist {
			if err := plugin.aclHandler.DeleteMacIPAcl(vppMacIPACL.Meta.Index); err != nil {
				plugin.log.Error(err)
				return err
			}
			// Unregister.
			plugin.l2AclIndexes.UnregisterName(vppMacIPACL.ACL.AclName)
			continue
		}
	}

	// Configure new ACLs
	for _, nbACL := range nbACLs {
		if err := plugin.ConfigureACL(nbACL); err != nil {
			plugin.log.Error(err)
			return err
		}
	}

	return nil
}
