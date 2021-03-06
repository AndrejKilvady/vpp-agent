*** Settings ***
Library      OperatingSystem
#Library      RequestsLibrary
#Library      SSHLibrary      timeout=60s
#Library      String

Resource     ../../../../variables/${VARIABLES}_variables.robot

Resource     ../../../../libraries/all_libs.robot

Suite Setup       Testsuite Setup
Suite Teardown    Suite Cleanup

*** Variables ***
${VARIABLES}=          common
${ENV}=                common
${FINAL_SLEEP}=        3s
${SYNC_SLEEP}=         10s
${IP_1}=               fd30::1:b:0:0:1
${IP_2}=               fd30::1:b:0:0:10


*** Test Cases ***
Configure Environment
    [Tags]    setup
    Add Agent VPP Node    agent_vpp_1    vswitch=${TRUE}
    Add Agent VPP Node    agent_vpp_2
    Add Agent VPP Node    agent_vpp_3
    Start SFC Controller Container With Own Config    basicIPv6.conf
    Sleep    ${SYNC_SLEEP}

Check Memifs On Vswitch
    vat_term: Check Memif Interface State     agent_vpp_1  IF_MEMIF_VSWITCH_agent_vpp_2_vpp2_memif1  role=master  connected=1  enabled=1
    vat_term: Check Memif Interface State     agent_vpp_1  IF_MEMIF_VSWITCH_agent_vpp_3_vpp3_memif1  role=master  connected=1  enabled=1

Check Memif Interface On VPP2
    vat_term: Check Memif Interface State     agent_vpp_2  vpp2_memif1  mac=02:02:02:02:02:02  role=slave  ipv6=${IP_1}/64  connected=1  enabled=1

Check Memif Interface On VPP3
    vat_term: Check Memif Interface State     agent_vpp_3  vpp3_memif1  role=slave  ipv6=${IP_2}/64  connected=1  enabled=1

Show Interfaces And Other Objects
    vpp_term: Show Interfaces    agent_vpp_1
    vpp_term: Show Interfaces    agent_vpp_2
    vpp_term: Show Interfaces    agent_vpp_3
    Write To Machine    agent_vpp_1_term    show int addr
    Write To Machine    agent_vpp_2_term    show int addr
    Write To Machine    agent_vpp_3_term    show int addr
    Write To Machine    agent_vpp_1_term    show h
    Write To Machine    agent_vpp_2_term    show h
    Write To Machine    agent_vpp_3_term    show h
    Write To Machine    agent_vpp_1_term    show br
    Write To Machine    agent_vpp_2_term    show br
    Write To Machine    agent_vpp_3_term    show br
    Write To Machine    agent_vpp_1_term    show br 1 detail
    Write To Machine    agent_vpp_2_term    show br 1 detail
    Write To Machine    agent_vpp_3_term    show br 1 detail
    Write To Machine    agent_vpp_1_term    show vxlan tunnel
    Write To Machine    agent_vpp_2_term    show vxlan tunnel
    Write To Machine    agent_vpp_3_term    show vxlan tunnel
    Write To Machine    agent_vpp_1_term    show err
    Write To Machine    agent_vpp_2_term    show err
    Write To Machine    agent_vpp_3_term    show err
    vat_term: Interfaces Dump    agent_vpp_1
    vat_term: Interfaces Dump    agent_vpp_2
    vat_term: Interfaces Dump    agent_vpp_3
    Write To Machine    vpp_agent_ctl    vpp-agent-ctl ${AGENT_VPP_ETCD_CONF_PATH} -ps
    Execute In Container    agent_vpp_1    ip a
    Execute In Container    agent_vpp_2    ip a
    Execute In Container    agent_vpp_3    ip a
    Make Datastore Snapshots    before_resync

Check Ping Agent2 -> Agent3
    vpp_term: Check Ping    agent_vpp_2    ${IP_2}

Check Ping Agent3 -> Agent2
    vpp_term: Check Ping    agent_vpp_3    ${IP_1}

Remove Agent Nodes
    Remove All Nodes

Start Agent Nodes Again
    Add Agent VPP Node    agent_vpp_1    vswitch=${TRUE}
    Add Agent VPP Node    agent_vpp_2
    Add Agent VPP Node    agent_vpp_3
    Sleep    ${SYNC_SLEEP}

Check Memifs On Vswitch After Resync
    vat_term: Check Memif Interface State     agent_vpp_1  IF_MEMIF_VSWITCH_agent_vpp_2_vpp2_memif1  role=master  connected=1  enabled=1
    vat_term: Check Memif Interface State     agent_vpp_1  IF_MEMIF_VSWITCH_agent_vpp_3_vpp3_memif1  role=master  connected=1  enabled=1

Check Memif Interface On VPP2 After Resync
    vat_term: Check Memif Interface State     agent_vpp_2  vpp2_memif1  mac=02:02:02:02:02:02  role=slave  ipv6=${IP_1}/64  connected=1  enabled=1

Check Memif Interface On VPP3 After Resync
    vat_term: Check Memif Interface State     agent_vpp_3  vpp3_memif1  role=slave  ipv6=${IP_2}/64  connected=1  enabled=1

Show Interfaces And Other Objects After Resync
    vpp_term: Show Interfaces    agent_vpp_1
    vpp_term: Show Interfaces    agent_vpp_2
    vpp_term: Show Interfaces    agent_vpp_3
    Write To Machine    agent_vpp_1_term    show int addr
    Write To Machine    agent_vpp_2_term    show int addr
    Write To Machine    agent_vpp_3_term    show int addr
    Write To Machine    agent_vpp_1_term    show h
    Write To Machine    agent_vpp_2_term    show h
    Write To Machine    agent_vpp_3_term    show h
    Write To Machine    agent_vpp_1_term    show br
    Write To Machine    agent_vpp_2_term    show br
    Write To Machine    agent_vpp_3_term    show br
    Write To Machine    agent_vpp_1_term    show br 1 detail
    Write To Machine    agent_vpp_2_term    show br 1 detail
    Write To Machine    agent_vpp_3_term    show br 1 detail
    Write To Machine    agent_vpp_1_term    show vxlan tunnel
    Write To Machine    agent_vpp_2_term    show vxlan tunnel
    Write To Machine    agent_vpp_3_term    show vxlan tunnel
    Write To Machine    agent_vpp_1_term    show err
    Write To Machine    agent_vpp_2_term    show err
    Write To Machine    agent_vpp_3_term    show err
    vat_term: Interfaces Dump    agent_vpp_1
    vat_term: Interfaces Dump    agent_vpp_2
    vat_term: Interfaces Dump    agent_vpp_3
    Write To Machine    vpp_agent_ctl    vpp-agent-ctl ${AGENT_VPP_ETCD_CONF_PATH} -ps
    Execute In Container    agent_vpp_1    ip a
    Execute In Container    agent_vpp_2    ip a
    Execute In Container    agent_vpp_3    ip a

Check Ping Agent2 -> Agent3 After Resync
    vpp_term: Check Ping    agent_vpp_2    ${IP_2}

Check Ping Agent3 -> Agent2 After Resync
    vpp_term: Check Ping    agent_vpp_3    ${IP_1}

Done
    [Tags]    debug
    No Operation

Final Sleep For Manual Checking
    [Tags]    debug
    Sleep   ${FINAL_SLEEP}

*** Keywords ***
Suite Cleanup
    Stop SFC Controller Container
    Testsuite Teardown
