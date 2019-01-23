*** Settings ***
Library      OperatingSystem
#Library      RequestsLibrary
#Library      SSHLibrary      timeout=60s
#Library      String

Resource     ../../../variables/${VARIABLES}_variables.robot
Resource     ../../../libraries/all_libs.robot
Resource     ../../../libraries/pretty_keywords.robot

Force Tags        crud     IPv4
Suite Setup       Testsuite Setup
Suite Teardown    Testsuite Teardown
Test Setup        TestSetup
Test Teardown     TestTeardown

*** Variables ***
${VARIABLES}=          common
${ENV}=                common
${WAIT_TIMEOUT}=         20s
${SYNC_SLEEP}=           3s
${EXT_IP_1}=             20.0.1.1
${IP_UNUSED}=            21.0.1.1
${EXT_PORT_1}=           3001
${INTERFACE_NAME_1}=     vpp1_tap1
${INTERFACE_NAME_2}=     vpp1_memif1
${MEMIF11_MAC}=          1a:00:00:11:11:11
${MEMIF12_MAC}=          3a:00:00:33:33:33
${error_message_1}=      Evaluating expression 'json.loads('''None''')' failed: ValueError: No JSON object could be decoded

${NAME_VPP1_TAP1}=          vpp1_tap1
${NAME_VPP2_TAP1}=          vpp2_tap1
${MAC_VPP1_TAP1}=           12:21:21:11:11:11
${MAC_VPP2_TAP1}=           22:21:21:22:22:22
${IP_VPP1_TAP1}=            10.10.1.1
${IP_VPP2_TAP1}=            20.20.1.1
${IP_LINUX_VPP1_TAP1}=      10.10.1.2
${IP_LINUX_VPP2_TAP1}=      20.20.1.2
${IP_VPP1_TAP1_NETWORK}=    10.10.1.0
${IP_VPP2_TAP1_NETWORK}=    20.20.1.0
${NAME_VPP1_MEMIF1}=        vpp1_memif1
${NAME_VPP2_MEMIF1}=        vpp2_memif1
${MAC_VPP1_MEMIF1}=         13:21:21:11:11:11
${MAC_VPP2_MEMIF1}=         23:21:21:22:22:22
${IP_VPP1_MEMIF1}=          192.168.1.1
${IP_VPP2_MEMIF1}=          192.168.1.2
${PREFIX}=                  24
${UP_STATE}=                up

${UDP_PORT}=                3001


*** Test Cases ***
Configure Environment
    [Tags]    setup
    Configure Environment 8

Add VPP1_TAP1 Interface
    vpp_term: Interface Not Exists  node=agent_vpp_1    mac=${MAC_VPP1_TAP1}
    Put TAP Interface With IP    node=agent_vpp_1    name=${NAME_VPP1_TAP1}    mac=${MAC_VPP1_TAP1}    ip=${IP_VPP1_TAP1}    prefix=${PREFIX}    host_if_name=linux_${NAME_VPP1_TAP1}
    linux: Set Host TAP Interface    node=agent_vpp_1    host_if_name=linux_${NAME_VPP1_TAP1}    ip=${IP_LINUX_VPP1_TAP1}    prefix=${PREFIX}
#
#Check Ping Between VPP1 and linux_VPP1_TAP1 Interface
#    linux: Check Ping    node=agent_vpp_1    ip=${IP_VPP1_TAP1}
#    vpp_term: Check Ping    node=agent_vpp_1    ip=${IP_LINUX_VPP1_TAP1}

Add VPP1_memif1 Interface
    vpp_term: Interface Not Exists    node=agent_vpp_1    mac=${MAC_VPP1_MEMIF1}
    Put Memif Interface With IP    node=agent_vpp_1    name=${NAME_VPP1_MEMIF1}    mac=${MAC_VPP1_MEMIF1}    master=true    id=1    ip=${IP_VPP1_MEMIF1}    prefix=24    socket=memif.sock
    Wait Until Keyword Succeeds   ${WAIT_TIMEOUT}   ${SYNC_SLEEP}    vpp_term: Interface Is Created    node=agent_vpp_1    mac=${MAC_VPP1_MEMIF1}

Add VPP2_TAP1 Interface
    vpp_term: Interface Not Exists  node=agent_vpp_2    mac=${MAC_VPP2_TAP1}
    Put TAP Interface With IP    node=agent_vpp_2    name=${NAME_VPP2_TAP1}    mac=${MAC_VPP2_TAP1}    ip=${IP_VPP2_TAP1}    prefix=${PREFIX}    host_if_name=linux_${NAME_VPP2_TAP1}
    linux: Set Host TAP Interface    node=agent_vpp_2    host_if_name=linux_${NAME_VPP2_TAP1}    ip=${IP_LINUX_VPP2_TAP1}    prefix=${PREFIX}

#Check Ping Between VPP2 And linux_VPP2_TAP1 Interface
#    linux: Check Ping    node=agent_vpp_2    ip=${IP_VPP2_TAP1}
#    vpp_term: Check Ping    node=agent_vpp_2    ip=${IP_LINUX_VPP2_TAP1}

Add VPP2_memif1 Interface
    vpp_term: Interface Not Exists    node=agent_vpp_2    mac=${MAC_VPP2_MEMIF1}
    Put Memif Interface With IP    node=agent_vpp_2    name=${NAME_VPP2_MEMIF1}    mac=${MAC_VPP2_MEMIF1}    master=false    id=1    ip=${IP_VPP2_MEMIF1}    prefix=24    socket=memif.sock
    Wait Until Keyword Succeeds   ${WAIT_TIMEOUT}   ${SYNC_SLEEP}    vpp_term: Interface Is Created    node=agent_vpp_1    mac=${MAC_VPP1_MEMIF1}
#
#Check Ping From VPP1 To VPP2_memif1
#    vpp_term: Check Ping    node=agent_vpp_1    ip=${IP_VPP2_MEMIF1}
#
#Check Ping From VPP2 To VPP1_memif1
#    vpp_term: Check Ping    node=agent_vpp_2    ip=${IP_VPP1_MEMIF1}

Ping From VPP1 Linux To VPP2_TAP1 And LINUX_VPP2_TAP1 Should Not Pass
    ${status1}=    Run Keyword And Return Status    linux: Check Ping    node=agent_vpp_1    ip=${IP_VPP2_TAP1}
    ${status2}=    Run Keyword And Return Status    linux: Check Ping    node=agent_vpp_1    ip=${IP_LINUX_VPP2_TAP1}
    Should Be Equal As Strings    ${status1}    False
    Should Be Equal As Strings    ${status2}    False

Ping From VPP2 Linux To VPP1_TAP1 And LINUX_VPP1_TAP1 Should Not Pass
    ${status1}=    Run Keyword And Return Status    linux: Check Ping    node=agent_vpp_2    ip=${IP_VPP1_TAP1}
    ${status2}=    Run Keyword And Return Status    linux: Check Ping    node=agent_vpp_2    ip=${IP_LINUX_VPP1_TAP1}
    Should Be Equal As Strings    ${status1}    False
    Should Be Equal As Strings    ${status2}    False

Add Static Route From VPP1 Linux To VPP2
    linux: Add Route    node=agent_vpp_1    destination_ip=${IP_VPP2_TAP1_NETWORK}    prefix=${PREFIX}    next_hop_ip=${IP_VPP1_TAP1}

Add Static Route From VPP1 To VPP2
    Create Route On agent_vpp_1 With IP ${IP_VPP2_TAP1_NETWORK}/${PREFIX} With Next Hop ${IP_VPP2_MEMIF1} And Vrf Id 0

Add Static Route From VPP2 Linux To VPP1
    linux: Add Route    node=agent_vpp_2    destination_ip=${IP_VPP1_TAP1_NETWORK}    prefix=${PREFIX}    next_hop_ip=${IP_VPP2_TAP1}

Add Static Route From VPP2 To VPP1
    Create Route On agent_vpp_2 With IP ${IP_VPP1_TAP1_NETWORK}/${PREFIX} With Next Hop ${IP_VPP1_MEMIF1} And Vrf Id 0
     Sleep     ${SYNC_SLEEP}

Check Ping From VPP1 Linux To VPP2_TAP1 And LINUX_VPP2_TAP1
    linux: Check Ping    node=agent_vpp_1    ip=${IP_VPP2_TAP1}
    linux: Check Ping    node=agent_vpp_1    ip=${IP_LINUX_VPP2_TAP1}

Check Ping From VPP2 Linux To VPP1_TAP1 And LINUX_VPP1_TAP1
    linux: Check Ping    node=agent_vpp_2    ip=${IP_VPP1_TAP1}
    linux: Check Ping    node=agent_vpp_2    ip=${IP_LINUX_VPP1_TAP1}

Start TCP And UDP Listeners
    #Sleep  295s
    Start UDP and TCP Ping Servers
    Sleep  10s

Check UDP Ping Agent1 -> Agent2
    linux: UDPPing Port Open  agent_vpp_1     ${IP_VPP2_TAP1}   ${UDP_PORT}

Check UDP Ping Agent1 -> Agent22
    linux: UDPPing Port Open  agent_vpp_1     ${IP_LINUX_VPP2_TAP1}   ${UDP_PORT}


Add NAT1, NAT2, Nat Global And Check Are Created
    Create DNat On agent_vpp_1 With Name dnat1 Local IP ${IP_VPP2_TAP1} Local Port ${UDP_PORT} External IP ${EXT_IP_1} External Port ${EXT_PORT_1} Vrf Id 0
    Create Interface GlobalNat On agent_vpp_1 With First IP ${EXT_IP_1} On Inteface ${IP_VPP1_TAP1} And Second IP ${IP_UNUSED} On Interface ${IP_VPP1_TAP1} Vrf Id 0 Config File nat-global-reduced.json
    Wait Until Keyword Succeeds   ${WAIT_TIMEOUT}   ${SYNC_SLEEP}    Get VPP NAT44 Config As Json    agent_vpp_1    dnat1
    Wait Until Keyword Succeeds   ${WAIT_TIMEOUT}   ${SYNC_SLEEP}    Get VPP NAT44 Global Config As Json    agent_vpp_1
    Wait Until Keyword Succeeds   ${WAIT_TIMEOUT}   ${SYNC_SLEEP}    Get VPP NAT44 Config As Json    agent_vpp_1    dnat1
    vpp_term: Check DNAT exists    agent_vpp_1    dnat_all_output_match.txt
    vpp_term: Check DNAT Global exists    agent_vpp_1    dnat_global_output_match.txt


Check UDP Ping Agent1 -> Agent2
    #Sleep  210s
    linux: UDPPing Port Open  agent_vpp_1     ${EXT_IP_1}   ${UDP_PORT}

Check UDP Ping Agent1 -> Agent22
    linux: UDPPing Port Open  agent_vpp_1     ${IP_VPP1_MEMIF1}   ${UDP_PORT}


*** Keywords ***
Start UDP and TCP Ping Servers
    #linux: Run TCP Ping Server On Node      node_1     ${TCP_PORT}
    linux: Run UDP Ping Server On Node      agent_vpp_1     ${UDP_PORT}
    #linux: Run TCP Ping Server On Node      node_2     ${TCP_PORT}
    linux: Run UDP Ping Server On Node      agent_vpp_2     ${UDP_PORT}
    linux: Check Processes on Node      agent_vpp_1
    linux: Check Processes on Node      agent_vpp_2
    Sleep    ${SYNC_SLEEP}
    ${bds_dump1}=    Execute On Machine    agent_vpp_1    netstat -nlt
    ${bds_dump2}=    Execute On Machine    agent_vpp_2    netstat -nlt
    ${bds_dump3}=    Execute On Machine    docker    netstat -nlt

TestSetup
    Make Datastore Snapshots    ${TEST_NAME}_test_setup

TestTeardown
    Make Datastore Snapshots    ${TEST_NAME}_test_teardown
