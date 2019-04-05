[Documentation]     Keywords for working with PAPI terminal
*** Settings ***
Library      papi_term.py

*** Variables ***
${terminal_timeout}=      30s
${bd_timeout}=            15s

*** Keywords ***

papi_term: Check PAPI Terminal
    [Arguments]           ${node}
    [Documentation]       Check PAPI terminal on node ${node}
    ${out}=               Write To Machine    ${node}    ${DOCKER_COMMAND} exec -it ${node} /bin/bash
    #SSHLibrary.Put_file   ${CURDIR}/vpp_api.py	    /tmp/
    #Execute On Machine     ${node}    ${DOCKER_COMMAND} cp /tmp/vpp_api.py ${node}:/
    #${command}=        Set Variable        ${VAT_START_COMMAND}
    #${out}=            Write To Machine    ${node}    ${command}
    #Should Contain     ${out}              ${${node}_VPP_VAT_PROMPT}
    ${out}=            Write To Machine Until String    ${node}_vat    from vpp_papi import VPP    ${${node}_VPP_VAT_PROMPT}    delay=${SSH_READ_DELAY}s
    ${out}=            Write To Machine Until String    ${node}_vat    vapi = VPP()    ${${node}_VPP_VAT_PROMPT}    delay=${SSH_READ_DELAY}s
    ${out}=            Write To Machine Until String    ${node}_vat    r = vapi.connect('papi-example')    ${${node}_VPP_VAT_PROMPT}    delay=${SSH_READ_DELAY}s
   [Return]           ${out}

papi_term: Open VAT Terminal
    [Arguments]    ${node}
    [Documentation]    Wait for VAT terminal on node ${node} or timeout
    wait until keyword succeeds  ${terminal_timeout}    5s   papi_term: Check PAPI Terminal    ${node}

papi_term: Exit VAT Terminal
    [Arguments]        ${node}
    ${ctrl_c}          Evaluate    chr(int(3))
    ${command}=        Set Variable       ${ctrl_c}
    ${out}=            Write To Machine   ${node}_vat    ${command}
    [Return]           ${out}

papi_term: Issue Command
    [Arguments]        ${node}     ${command}    ${delay}=${SSH_READ_DELAY}s
    #${out}=            Write To Machine Until String    ${node}_vat    ${command}    ${${node}_VPP_VAT_PROMPT}    delay=${delay}
    ${out}=            Write To Machine Until String    ${node}_vat    ret = vapi.api.${command}()    ${${node}_VPP_VAT_PROMPT}    delay=${delay}
    ${out}=            Write To Machine Until String    ${node}_vat    print(ret)    ${${node}_VPP_VAT_PROMPT}    delay=${delay}
    Should Contain     ${out}             ${${node}_VPP_VAT_PROMPT}
    [Return]           ${out}

papi_term: Interfaces Dump
    [Arguments]        ${node}
    [Documentation]    Executing command sw_interface_dump
    ${out}=            papi_term: Issue Command  ${node}  sw_interface_dump
    ${out}=            papi_term.Process_Reply_2    ${out}
    [Return]           ${out}


papi_term: Check Loopback Interface State
    [Arguments]          ${node}    ${name}    @{desired_state}
    ${internal_name}=    Get Interface Internal Name    ${node}    ${name}
    ${internal_index}=   papi_term: Get Interface Index    ${node}    ${internal_name}
    ${interfaces}=       papi_term: Interfaces Dump    ${node}
    ${int_state}=        papi_term.Papi Get Interface State    ${interfaces}    ${internal_name}
    ${ipv4_list}=        vpp_term: Get Interface IPs    ${node}    ${internal_name}
    ${ipv6_list}=        vpp_term: Get Interface IP6 IPs    ${node}    ${internal_name}
    ${enabled}=          Set Variable    ${int_state}   #${int_state["admin_up_down"]}
    ${mtu}=              Papi_Get_Mtu    ${interfaces}    ${internal_name}    #Set Variable    ${int_state["mtu"]}
    #${dec_mac}=          Set Variable    ${int_state["l2_address"]}
    ${mac}=              Papi_Get_Mac    ${interfaces}    ${internal_name}    #Convert Dec MAC To Hex    ${dec_mac}
    ${actual_state}=     Create List    enabled=${enabled}    mtu=${mtu}    mac=${mac}
    :FOR    ${ip}    IN    @{ipv4_list}
    \    Append To List    ${actual_state}    ipv4=${ip}
    :FOR    ${ip}    IN    @{ipv6_list}
    \    Append To List    ${actual_state}    ipv6=${ip}
    List Should Contain Sub List    ${actual_state}    ${desired_state}
    [Return]             ${actual_state}

papi_term: Get Interface Name
    [Arguments]        ${node}     ${index}
    [Documentation]    Return interface with specified index name
    ${out}=            papi_term: Interfaces Dump    ${node}
    ${name}=           Get Interface Name    ${out}    ${index}
    [Return]           ${name}

papi_term: Get Interface Index
    [Arguments]        ${node}     ${name}
    [Documentation]    Return interface index with specified name
    ${out}=            papi_term: Interfaces Dump    ${node}
    ${index}=          Vpp Get Interface Index    ${out}    ${name}
    [Return]           ${index}