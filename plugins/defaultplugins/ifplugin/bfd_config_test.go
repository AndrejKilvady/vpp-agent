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

package ifplugin_test

import (
	"net"
	"strings"
	"testing"

	"git.fd.io/govpp.git/adapter/mock"
	"git.fd.io/govpp.git/core"
	"github.com/ligato/cn-infra/logging"
	"github.com/ligato/cn-infra/logging/logrus"
	"github.com/ligato/vpp-agent/idxvpp/nametoidx"
	bfd_api "github.com/ligato/vpp-agent/plugins/defaultplugins/common/bin_api/bfd"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/common/bin_api/vpe"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/bfd"
	if_api "github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/interfaces"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/ifplugin"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/ifplugin/ifaceidx"
	"github.com/ligato/vpp-agent/tests/vppcallmock"
	. "github.com/onsi/gomega"
)

/* BFD configurator init and close */

// Test init function
func TestBfdConfiguratorInit(t *testing.T) {
	RegisterTestingT(t)
	connection, _ := core.Connect(&mock.VppAdapter{})
	defer connection.Disconnect()

	plugin := &ifplugin.BFDConfigurator{}
	err := plugin.Init(logging.ForPlugin("test-log", logrus.NewLogRegistry()), connection,
		nil, true)
	Expect(err).To(BeNil())

	err = plugin.Close()
	Expect(err).To(BeNil())
}

/* BFD Sessions */

// Configure BFD session without interface
func TestBfdConfiguratorConfigureSessionNoInterfaceError(t *testing.T) {
	var err error
	// Setup
	_, connection, plugin, _ := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Data
	data := getTestBfdSession("if1", "10.0.0.1")
	// Test configure BFD session without interface
	err = plugin.ConfigureBfdSession(data)
	Expect(err).ToNot(BeNil())
}

// Configure BFD session while interface metadata is missing
func TestBfdConfiguratorConfigureSessionNoInterfaceMetaError(t *testing.T) {
	var err error
	// Setup
	_, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Data
	data := getTestBfdSession("if1", "10.0.0.1")
	// Register
	ifIndexes.RegisterName("if1", 1, nil)
	// Test configure BFD session
	err = plugin.ConfigureBfdSession(data)
	Expect(err).ToNot(BeNil())
	_, _, found := plugin.GetBfdSessionIndexes().LookupIdx(data.Interface)
	Expect(found).To(BeFalse())
}

// Configure BFD session while source IP does not match with interface IP
func TestBfdConfiguratorConfigureSessionSrcDoNotMatch(t *testing.T) {
	var err error
	// Setup
	_, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Data
	data := getTestBfdSession("if1", "10.0.0.2")
	// Register
	ifIndexes.RegisterName("if1", 1, getSimpleTestInterface("if1", []string{"10.0.0.1/24"}))
	// Test configure BFD session
	err = plugin.ConfigureBfdSession(data)
	Expect(err).ToNot(BeNil())
	_, _, found := plugin.GetBfdSessionIndexes().LookupIdx(data.Interface)
	Expect(found).To(BeFalse())
}

// Configure BFD session
func TestBfdConfiguratorConfigureSession(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdUDPAddReply{})
	// Data
	data := getTestBfdSession("if1", "10.0.0.1")
	// Register
	ifIndexes.RegisterName("if1", 1, getSimpleTestInterface("if1", []string{"10.0.0.1/24"}))
	// Test configure BFD session
	err = plugin.ConfigureBfdSession(data)
	Expect(err).To(BeNil())
	_, meta, found := plugin.GetBfdSessionIndexes().LookupIdx(data.Interface)
	Expect(found).To(BeTrue())
	Expect(meta).To(BeNil())
}

// Configure BFD session with non-zero reply
func TestBfdConfiguratorConfigureSessionError(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdUDPAddReply{
		Retval: 1,
	})
	// Data
	data := getTestBfdSession("if1", "10.0.0.1")
	// Register
	ifIndexes.RegisterName("if1", 1, getSimpleTestInterface("if1", []string{"10.0.0.1/24"}))
	// Test configure BFD session
	err = plugin.ConfigureBfdSession(data)
	Expect(err).ToNot(BeNil())
	_, _, found := plugin.GetBfdSessionIndexes().LookupIdx(data.Interface)
	Expect(found).To(BeFalse())
}

// Modify BFD session without interface
func TestBfdConfiguratorModifySessionNoInterfaceError(t *testing.T) {
	var err error
	// Setup
	_, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Data
	oldData := getTestBfdSession("if1", "10.0.0.1")
	newData := getTestBfdSession("if2", "10.0.0.2")
	// Register
	ifIndexes.RegisterName("if1", 1, getSimpleTestInterface("if1", []string{"10.0.0.1/24"}))
	// Test modify BFD session
	err = plugin.ModifyBfdSession(oldData, newData)
	Expect(err).ToNot(BeNil())
}

// Modify BFD session without interface metadata
func TestBfdConfiguratorModifySessionNoInterfaceMeta(t *testing.T) {
	var err error
	// Setup
	_, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Data
	oldData := getTestBfdSession("if1", "10.0.0.1")
	newData := getTestBfdSession("if1", "10.0.0.2")
	// Register
	ifIndexes.RegisterName("if1", 1, nil)
	// Test modify BFD session
	err = plugin.ModifyBfdSession(oldData, newData)
	Expect(err).ToNot(BeNil())
}

// Modify BFD session where source IP does not match
func TestBfdConfiguratorModifySessionSrcDoNotMatchError(t *testing.T) {
	var err error
	// Setup
	_, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Data
	oldData := getTestBfdSession("if1", "10.0.0.1")
	newData := getTestBfdSession("if1", "10.0.0.2")
	// Register
	plugin.GetBfdSessionIndexes().RegisterName(oldData.Interface, 1, nil)
	ifIndexes.RegisterName("if1", 1, getSimpleTestInterface("if1", []string{"10.0.0.3/24"}))
	// Test modify BFD session
	err = plugin.ModifyBfdSession(oldData, newData)
	Expect(err).ToNot(BeNil())
}

// Modify BFD session without previous data
func TestBfdConfiguratorModifySessionNoPrevious(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdUDPAddReply{})
	// Data
	oldData := getTestBfdSession("if1", "10.0.0.1")
	newData := getTestBfdSession("if2", "10.0.0.2")
	// Register
	ifIndexes.RegisterName("if2", 1, getSimpleTestInterface("if1", []string{"10.0.0.2/24"}))
	// Test modify BFD session
	err = plugin.ModifyBfdSession(oldData, newData)
	Expect(err).To(BeNil())
	_, meta, found := plugin.GetBfdSessionIndexes().LookupIdx(newData.Interface)
	Expect(found).To(BeTrue())
	Expect(meta).To(BeNil())
}

// Modify BFD session different source addresses
func TestBfdConfiguratorModifySessionSrcAddrDiffError(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdUDPModReply{})
	// Data
	oldData := getTestBfdSession("if1", "10.0.0.1")
	newData := getTestBfdSession("if1", "10.0.0.2")
	// Register
	plugin.GetBfdSessionIndexes().RegisterName(oldData.Interface, 1, nil)
	ifIndexes.RegisterName("if1", 1, getSimpleTestInterface("if1", []string{"10.0.0.1/24", "10.0.0.2/24"}))
	// Test modify BFD session
	err = plugin.ModifyBfdSession(oldData, newData)
	Expect(err).ToNot(BeNil())
}

// Modify BFD session
func TestBfdConfiguratorModifySession(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdUDPModReply{})
	// Data
	oldData := getTestBfdSession("if1", "10.0.0.1")
	newData := getTestBfdSession("if1", "10.0.0.1")
	// Register
	plugin.GetBfdSessionIndexes().RegisterName(oldData.Interface, 1, nil)
	ifIndexes.RegisterName("if1", 1, getSimpleTestInterface("if1", []string{"10.0.0.1/24"}))
	// Test modify BFD session
	err = plugin.ModifyBfdSession(oldData, newData)
	Expect(err).To(BeNil())
	_, meta, found := plugin.GetBfdSessionIndexes().LookupIdx(newData.Interface)
	Expect(found).To(BeTrue())
	Expect(meta).To(BeNil())
	err = plugin.Close()
	Expect(err).To(BeNil())
}

// Test delete BFD session no interface
func TestBfdConfiguratorDeleteSessionNoInterfaceError(t *testing.T) {
	var err error
	// Setup
	_, connection, plugin, _ := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Data
	data := getTestBfdSession("if1", "10.0.0.1")
	// Register
	plugin.GetBfdSessionIndexes().RegisterName(data.Interface, 1, nil)
	// Modify BFD session
	err = plugin.DeleteBfdSession(data)
	Expect(err).ToNot(BeNil())
}

// Test delete BFD session
func TestBfdConfiguratorDeleteSession(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdUDPDelReply{})
	// Data
	data := getTestBfdSession("if1", "10.0.0.1")
	// Register
	plugin.GetBfdSessionIndexes().RegisterName(data.Interface, 1, nil)
	ifIndexes.RegisterName("if1", 1, getSimpleTestInterface("if1", []string{"10.0.0.1/24"}))
	// Modify BFD session
	err = plugin.DeleteBfdSession(data)
	Expect(err).To(BeNil())
	_, _, found := plugin.GetBfdSessionIndexes().LookupIdx(data.Interface)
	Expect(found).To(BeFalse())
}

// Test delete BFD session with retval error
func TestBfdConfiguratorDeleteSessionError(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdUDPDelReply{
		Retval: 1,
	})
	// Data
	data := getTestBfdSession("if1", "10.0.0.1")
	// Register
	plugin.GetBfdSessionIndexes().RegisterName(data.Interface, 1, nil)
	ifIndexes.RegisterName("if1", 1, getSimpleTestInterface("if1", []string{"10.0.0.1/24"}))
	// Modify BFD session
	err = plugin.DeleteBfdSession(data)
	Expect(err).ToNot(BeNil())
}

// BFD session dump
func TestBfdConfiguratorDumpBfdSessions(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdUDPSessionDetails{
		SwIfIndex: 1,
		LocalAddr: net.ParseIP("10.0.0.1").To4(),
		PeerAddr:  net.ParseIP("10.0.0.2").To4(),
	})
	ctx.MockVpp.MockReply(&bfd_api.BfdUDPSessionDetails{
		SwIfIndex: 2,
		LocalAddr: net.ParseIP("10.0.0.3").To4(),
		PeerAddr:  net.ParseIP("10.0.0.4").To4(),
	})
	ctx.MockVpp.MockReply(&vpe.ControlPingReply{})
	// Register (only first interface)
	ifIndexes.RegisterName("if1", 1, nil)
	// Test bfd session dump
	sessions, err := plugin.DumpBfdSessions()
	Expect(err).To(BeNil())
	Expect(sessions).To(HaveLen(2))
}

// Configure BFD authentication key
func TestBfdConfiguratorSetAuthKey(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, _ := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdAuthSetKeyReply{})
	// Data
	data := getTestBfdAuthKey("key1", "secret", 1, 1, bfd.SingleHopBFD_Key_KEYED_SHA1)
	// Test key configuration
	err = plugin.ConfigureBfdAuthKey(data)
	Expect(err).To(BeNil())
	_, _, found := plugin.GetBfdKeyIndexes().LookupIdx(ifplugin.AuthKeyIdentifier(data.Id))
	Expect(found).To(BeTrue())
}

// Configure BFD authentication key with error return value
func TestBfdConfiguratorSetAuthKeyError(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, _ := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdAuthSetKeyReply{
		Retval: 1,
	})
	// Data
	data := getTestBfdAuthKey("key1", "secret", 1, 1, bfd.SingleHopBFD_Key_KEYED_SHA1)
	// Test key configuration
	err = plugin.ConfigureBfdAuthKey(data)
	Expect(err).ToNot(BeNil())
	_, _, found := plugin.GetBfdKeyIndexes().LookupIdx(ifplugin.AuthKeyIdentifier(data.Id))
	Expect(found).To(BeFalse())
}

// Modify BFD authentication key which is not used in any session
func TestBfdConfiguratorModifyUnusedAuthKey(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, _ := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&vpe.ControlPingReply{})       // Session dump
	ctx.MockVpp.MockReply(&bfd_api.BfdAuthDelKeyReply{}) // Authentication key delete/create
	ctx.MockVpp.MockReply(&bfd_api.BfdAuthSetKeyReply{})
	// Data
	oldData := getTestBfdAuthKey("key1", "secret", 1, 1, bfd.SingleHopBFD_Key_KEYED_SHA1)
	newData := getTestBfdAuthKey("key2", "secret", 1, 1, bfd.SingleHopBFD_Key_METICULOUS_KEYED_SHA1)
	// Register
	plugin.GetBfdKeyIndexes().RegisterName(ifplugin.AuthKeyIdentifier(oldData.Id), 1, nil)
	// Test key modification
	err = plugin.ModifyBfdAuthKey(oldData, newData)
	Expect(err).To(BeNil())
}

// Modify BFD authentication key which is used in session
func TestBfdConfiguratorModifyUsedAuthKey(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, _ := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply handler
	ctx.MockVpp.RegisterBinAPITypes(bfd_api.Types)
	ctx.MockVpp.RegisterBinAPITypes(vpe.Types)
	ctx.MockVpp.MockReplyHandler(bfdVppMockHandler(ctx.MockVpp))
	// Data
	oldData := getTestBfdAuthKey("key1", "secret", 1, 1, bfd.SingleHopBFD_Key_KEYED_SHA1)
	newData := getTestBfdAuthKey("key1", "secret", 1, 1, bfd.SingleHopBFD_Key_METICULOUS_KEYED_SHA1)
	// Register
	plugin.GetBfdKeyIndexes().RegisterName(ifplugin.AuthKeyIdentifier(oldData.Id), 1, nil)
	// Test key modification
	err = plugin.ModifyBfdAuthKey(oldData, newData)
	Expect(err).To(BeNil())
}

// Delete BFD authentication key which is not used in any session
func TestBfdConfiguratorDeleteUnusedAuthKey(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, _ := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&vpe.ControlPingReply{})       // Session dump
	ctx.MockVpp.MockReply(&bfd_api.BfdAuthDelKeyReply{}) // Authentication key delete
	// Data
	data := getTestBfdAuthKey("key1", "secret", 1, 1, bfd.SingleHopBFD_Key_KEYED_SHA1)
	// Register
	plugin.GetBfdKeyIndexes().RegisterName(ifplugin.AuthKeyIdentifier(data.Id), 1, nil)
	// Test key modification
	err = plugin.DeleteBfdAuthKey(data)
	Expect(err).To(BeNil())
}

// Delete BFD authentication key which is used in session
func TestBfdConfiguratorDeleteUsedAuthKey(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, _ := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply handler
	ctx.MockVpp.RegisterBinAPITypes(bfd_api.Types)
	ctx.MockVpp.RegisterBinAPITypes(vpe.Types)
	ctx.MockVpp.MockReplyHandler(bfdVppMockHandler(ctx.MockVpp))
	// Data
	data := getTestBfdAuthKey("key1", "secret", 1, 1, bfd.SingleHopBFD_Key_KEYED_SHA1)
	// Register
	plugin.GetBfdKeyIndexes().RegisterName(ifplugin.AuthKeyIdentifier(data.Id), 1, nil)
	// Test key modification
	err = plugin.DeleteBfdAuthKey(data)
	Expect(err).To(BeNil())
}

// Dump BFD authentication key
func TestBfdConfiguratorDumpAuthKey(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, _ := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdAuthKeysDetails{
		ConfKeyID: 1,
		AuthType:  4, // Means KEYED SHA1
	})
	ctx.MockVpp.MockReply(&bfd_api.BfdAuthKeysDetails{
		ConfKeyID: 2,
		AuthType:  1, // Any other number is METICULOUS KEYED SHA1
	})
	ctx.MockVpp.MockReply(&vpe.ControlPingReply{})
	// Test authentication key dump
	keys, err := plugin.DumpBFDAuthKeys()
	Expect(err).To(BeNil())
	Expect(keys).To(HaveLen(2))
	Expect(keys[0].AuthenticationType).To(BeEquivalentTo(bfd.SingleHopBFD_Key_KEYED_SHA1))
	Expect(keys[1].AuthenticationType).To(BeEquivalentTo(bfd.SingleHopBFD_Key_METICULOUS_KEYED_SHA1))
}

// Configure BFD echo function create/modify/delete
func TestBfdConfiguratorEchoFunction(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdUDPSetEchoSourceReply{})
	ctx.MockVpp.MockReply(&bfd_api.BfdUDPDelEchoSourceReply{})
	// Data
	data := getTestBfdEchoFunction("if1")
	//Registration
	ifIndexes.RegisterName("if1", 1, nil)
	// Test Echo function create
	err = plugin.ConfigureBfdEchoFunction(data)
	Expect(err).To(BeNil())
	_, _, found := plugin.GetBfdEchoFunctionIndexes().LookupIdx(data.EchoSourceInterface)
	Expect(found).To(BeTrue())
	// Test Echo function modify
	err = plugin.ModifyBfdEchoFunction(data, data)
	Expect(err).To(BeNil())
	// Test echo function delete
	err = plugin.DeleteBfdEchoFunction(data)
	Expect(err).To(BeNil())
	_, _, found = plugin.GetBfdEchoFunctionIndexes().LookupIdx(data.EchoSourceInterface)
	Expect(found).To(BeFalse())
}

// Configure BFD echo function with return value error
func TestBfdConfiguratorEchoFunctionConfigError(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, ifIndexes := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdUDPSetEchoSourceReply{
		Retval: 1,
	})
	// Data
	data := getTestBfdEchoFunction("if1")
	// Register
	ifIndexes.RegisterName("if1", 1, nil)
	// Test Echo function create
	err = plugin.ConfigureBfdEchoFunction(data)
	Expect(err).ToNot(BeNil())
	_, _, found := plugin.GetBfdEchoFunctionIndexes().LookupIdx(data.EchoSourceInterface)
	Expect(found).To(BeFalse())
}

// Configure BFD echo function create with non-existing interface
func TestBfdConfiguratorEchoFunctionNoInterfaceError(t *testing.T) {
	var err error
	// Setup
	_, connection, plugin, _ := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Data
	data := getTestBfdEchoFunction("if1")
	// Test Echo function create
	err = plugin.ConfigureBfdEchoFunction(data)
	Expect(err).ToNot(BeNil())
	_, _, found := plugin.GetBfdEchoFunctionIndexes().LookupIdx(data.EchoSourceInterface)
	Expect(found).To(BeFalse())
}

// Delete BFD echo function with return value error
func TestBfdConfiguratorEchoFunctionDeleteError(t *testing.T) {
	var err error
	// Setup
	ctx, connection, plugin, _ := bfdTestSetup(t)
	defer bfdTestTeardown(connection, plugin)
	// Reply set
	ctx.MockVpp.MockReply(&bfd_api.BfdUDPSetEchoSourceReply{
		Retval: 1,
	})
	// Data
	data := getTestBfdEchoFunction("if1")
	// Test Echo function create
	err = plugin.DeleteBfdEchoFunction(data)
	Expect(err).ToNot(BeNil())
	_, _, found := plugin.GetBfdEchoFunctionIndexes().LookupIdx(data.EchoSourceInterface)
	Expect(found).To(BeFalse())
}

/* BFD Test Setup */

func bfdTestSetup(t *testing.T) (*vppcallmock.TestCtx, *core.Connection, *ifplugin.BFDConfigurator, ifaceidx.SwIfIndexRW) {
	RegisterTestingT(t)
	ctx := &vppcallmock.TestCtx{
		MockVpp: &mock.VppAdapter{},
	}
	connection, err := core.Connect(ctx.MockVpp)
	Expect(err).ShouldNot(HaveOccurred())
	// Logger
	log := logging.ForPlugin("test-log", logrus.NewLogRegistry())
	log.SetLevel(logging.DebugLevel)
	// Interface indices
	swIfIndices := ifaceidx.NewSwIfIndex(nametoidx.NewNameToIdx(log, "stn", nil))
	// Configurator
	plugin := &ifplugin.BFDConfigurator{}
	err = plugin.Init(log, connection, swIfIndices, false)
	Expect(err).To(BeNil())

	return ctx, connection, plugin, swIfIndices
}

func bfdTestTeardown(connection *core.Connection, plugin *ifplugin.BFDConfigurator) {
	connection.Disconnect()
	err := plugin.Close()
	Expect(err).To(BeNil())
}

func bfdVppMockHandler(vppMock *mock.VppAdapter) mock.ReplyHandler {
	var sendControlPing bool
	return func(request mock.MessageDTO) (reply []byte, msgID uint16, prepared bool) {
		logrus.DefaultLogger().Errorf("recived request %v", request.MsgName)
		if sendControlPing {
			sendControlPing = false
			data := &vpe.ControlPingReply{}
			reply, err := vppMock.ReplyBytes(request, data)
			Expect(err).To(BeNil())
			msgID, err := vppMock.GetMsgID(data.GetMessageName(), data.GetCrcString())
			Expect(err).To(BeNil())
			return reply, msgID, true
		}
		if strings.HasSuffix(request.MsgName, "_dump") {
			// Send control ping after first iteration
			sendControlPing = true
			data := &bfd_api.BfdUDPSessionDetails{
				SwIfIndex:       1,
				LocalAddr:       net.ParseIP("10.0.0.1").To4(),
				PeerAddr:        net.ParseIP("10.0.0.2").To4(),
				IsAuthenticated: 1,
				BfdKeyID:        1,
			}
			reply, err := vppMock.ReplyBytes(request, data)
			Expect(err).To(BeNil())
			msgID, err := vppMock.GetMsgID(data.GetMessageName(), data.GetCrcString())
			Expect(err).To(BeNil())
			return reply, msgID, true
		} else {
			if replyMsg, msgID, ok := vppMock.ReplyFor(request.MsgName); ok {
				reply, err := vppMock.ReplyBytes(request, replyMsg)
				Expect(err).To(BeNil())
				return reply, msgID, true
			}
		}

		return reply, 0, false
	}
}

/* BFD Test Data */

func getTestBfdSession(ifName, srcAddr string) *bfd.SingleHopBFD_Session {
	return &bfd.SingleHopBFD_Session{
		Interface:          ifName,
		SourceAddress:      srcAddr,
		DestinationAddress: "10.0.0.5",
	}
}

func getTestBfdAuthKey(name, secret string, keyIdx, id uint32, keyType bfd.SingleHopBFD_Key_AuthenticationType) *bfd.SingleHopBFD_Key {
	return &bfd.SingleHopBFD_Key{
		Name:               name,
		AuthKeyIndex:       keyIdx,
		Id:                 id,
		AuthenticationType: keyType,
		Secret:             secret,
	}
}

func getTestBfdEchoFunction(ifName string) *bfd.SingleHopBFD_EchoFunction {
	return &bfd.SingleHopBFD_EchoFunction{
		Name:                "echo",
		EchoSourceInterface: ifName,
	}
}

func getSimpleTestInterface(name string, ip []string) *if_api.Interfaces_Interface {
	return &if_api.Interfaces_Interface{
		Name:        name,
		IpAddresses: ip,
	}
}
