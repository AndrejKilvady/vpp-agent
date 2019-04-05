import json
import re
import binascii
import ast
from socket import inet_ntoa
#from ipaddress import IPv6Address

from robot.api import logger

# input - json output from vxlan_tunnel_dump, src ip, dst ip, vni
# output - true if tunnel exists, false if not, interface index
def Check_VXLan_Tunnel_Presence(out, src, dst, vni):
    out =  out[out.find('['):out.rfind(']')+1]
    data = json.loads(out)
    present = False
    if_index = -1
    for iface in data:
        if iface["src_address"] == src and iface["dst_address"] == dst and iface["vni"] == int(vni):
            present = True
            if_index  = iface["sw_if_index"]
    return present, if_index

# input - json output from sw_interface_dump, index
# output - interface name
def Get_Interface_Name(out, index):
    out =  out[out.find('['):out.rfind(']')+1]
    data = json.loads(out)
    name = "x"
    for iface in data:
        if iface["sw_if_index"] == int(index):
            name = iface["interface_name"]
    return name

def replace_rn(mytext):
    if mytext=="":
        return ""
    mytext=mytext.replace("\r\n", "")
    #print mytext
    return mytext

def replace_rrn(mytext):
    if mytext=="":
        return ""
    mytext=mytext.replace("\r\r\n", "")
    #print mytext
    return mytext

def replace_spaces_to_space(mytext):
    mytext = re.sub(
           r" +",
           " ", mytext
           )
    return mytext

# input - json output from sw_interface_dump, interface name
# output - index
def Get_Interface_Index(out, name):
    #out = process_reply_2(out)
    #out = json.loads(out)

    #out =  out[out.find('['):out.rfind(']')+1]
    #data = json.loads(out)
    data = out
    index = -1
    for detail in data:
        print data
        print detail['sw_interface_details']['interface_name']
        if detail['sw_interface_details']['interface_name']  == name:
            index = detail['sw_interface_details']['sw_if_index']
    return index

# input - output from sh int, interface name
# output - index
def Vpp_Get_Interface_Index(out, name):
    swifindex = str('sw_if_index')
    out = replace_rrn(out)
    out = replace_spaces_to_space(out)
    data =  out[out.rfind(name)-110:out.rfind(name)-50]
    data = data.replace("=", ' ')
    data = data.replace(",", '')
    print data
    data = data[data.find(swifindex):data.find(swifindex)+15]
    print data
    index = -1
    numbers = [int(s) for s in data.split() if s.isdigit()]
    print data
    print numbers
    if len(numbers) > 0:
       index = numbers[0]
    else:
       print "Index Not Found"
    return index

def Process_Reply_2(out):
    #data = json.dumps(out)
    #print data
    #out = out.replace("'", '"')
    out = out.replace(">>>", '')
    #out = out.replace("]", '')
    #out = out.replace("[", '')
    print out
    out = out.replace("\\x00", '')
    out = replace_rrn(out)
    out = replace_rn(out)
    print out
    out = replace_spaces_to_space(out)
    # reply_converted = _convert_reply(out)
    # print reply_converted

    # data = json.loads(out.replace("\'", "\""))
    # #data = json.loads(out)
    # #data = reply_converted

    # for detail in out:
    #     detail['sw_interface_details'] = detail['sw_interface_details']['interface_name'].decode(
    #         "hex").replace("\x00", "")
    #     detail['sw_interface_details']['tag'] = detail['sw_interface_details']['tag'].decode("hex").replace("\x00", "")
    #     detail['sw_interface_details']['l2_address'] = ":".join("{:02x}".format(ord(c)) for c in detail['sw_interface_details']['l2_address'].decode("hex").replace("\x00", ""))

    # for detail in data:
    #     detail['sw_interface_details']['interface_name'] = detail['sw_interface_details']['interface_name'].decode(
    #         "hex").replace("\x00", "")
    #     detail['sw_interface_details']['tag'] = detail['sw_interface_details']['tag'].decode("hex").replace("\x00", "")
    #     detail['sw_interface_details']['l2_address'] = ":".join("{:02x}".format(ord(c)) for c in detail['sw_interface_details']['l2_address'].decode("hex").replace("\x00", ""))
    #
    # print data
    return out

def _convert_reply(api_r):
    """Process API reply / a part of API reply for smooth converting to
    JSON string.
    It is used only with 'request' and 'dump' methods.
    Apply binascii.hexlify() method for string values.
    TODO: Implement complex solution to process of replies.
    :param api_r: API reply.
    :type api_r: Vpp_serializer reply object (named tuple)
    :returns: Processed API reply / a part of API reply.
    :rtype: dict
    """
    unwanted_fields = ['count', 'index', 'context']

    reply_dict = dict()
    reply_key = repr(api_r).split('(')[0]
    reply_value = dict()
    for item in dir(api_r):
        if not item.startswith('_') and item not in unwanted_fields:
            attr_value = getattr(api_r, item)
            if isinstance(attr_value, list) or isinstance(attr_value, dict):
                value = attr_value
            elif hasattr(attr_value, '__int__'):
                value = int(attr_value)
            elif hasattr(attr_value, '__str__'):
                value = binascii.hexlify(str(attr_value))
            # Next handles parameters not supporting preferred integer or string
            # representation to get it logged
            elif hasattr(attr_value, '__repr__'):
                value = repr(attr_value)
            else:
                value = attr_value
            reply_value[item] = value
    reply_dict[reply_key] = reply_value
    return reply_dict


# input - json output from sw_interface_dump, index
# output - whole interface state
def Papi_Get_Interface_State(out, name):
    #out =  out[out.find('['):out.rfind(']')+1]
    #data = json.loads(out)
    state = -1
    data = out[out.rfind(name):out.rfind(name)+30]
    data = data.replace("=", ' ')
    data = data.replace(",", '')
    print data
    data = data[data.find('admin_up_down'):data.find('admin_up_down') + 15]
    print data
    numbers = [int(s) for s in data.split() if s.isdigit()]
    print data
    print numbers
    if len(numbers) > 0:
        state = numbers[0]
    else:
        print "State Not Found"
    # for detail in out:
    #     if detail['sw_interface_details']['sw_if_index']  == index:
    #         state = detail['sw_interface_details']['admin_up_down']
    return state

# input - json output from sw_interface_dump, index
# output - whole interface state
def Papi_Get_Mac(out, name):
    #out =  out[out.find('['):out.rfind(']')+1]
    #data = json.loads(out)
    mac = ''

    data = out[out.rfind(name) - 55:out.rfind(name) - 15]
    data = data.replace("=", ' ')
    data = data.replace(",", '')
    #data = data.replace("\'\\'", '\\')
    print data
    data = data[data.find('l2_address')+10:data.find('l2_address') + 38]
    print data
    data = data[data.find('\\'):data.find('\' ')]
    print data
    mac2 = mac_ntop2("\x12\x91\x91\x11\x11\x11")
    print mac2
    mac = mac_ntop2('\"'+data+'\"')
    print mac
    if len(mac) > 0:
        print mac
    else:
        print "Mac address Not Found"


    # for detail in out:
    #     if detail['sw_interface_details']['sw_if_index']  == index:
    #         mac = detail['sw_interface_details']['l2_address']   #.decode("hex")
    #         #mac = ":".join("{:02x}".format(ord(c)) for c in detail['sw_interface_details']['l2_address'])
    #         print mac
    return mac

# input - json output from sw_interface_dump, index
# output - whole interface state
def Papi_Get_Mtu(out, name):
    #out =  out[out.find('['):out.rfind(']')+1]
    #data = json.loads(out)
    mtu = ''
    data = out[out.rfind(name)+50:out.rfind(name)+100]
    data = data.replace("=", ' ')
    data = data.replace(",", '')
    print data
    data = data[data.find('link_mtu'):data.find('link_mtu') + 15]
    print data
    numbers = [int(s) for s in data.split() if s.isdigit()]
    print data
    print numbers
    if len(numbers) > 0:
        mtu = numbers[0]
    else:
        print "State Not Found"
    # for detail in out:
    #     if detail['sw_interface_details']['sw_if_index']  == index:
    #         mtu = detail['sw_interface_details']['mtu']
    #         mtu = mtu[0]
    return mtu

# input - json output from sw_interface_dump, index
# output - whole interface state
def Get_Interface_State(out, index):
    out =  out[out.find('['):out.rfind(']')+1]
    data = json.loads(out)
    state = -1
    for iface in data:
        if iface["sw_if_index"] == int(index):
            state = iface
    return state

def mac_ntop2(binary):
    '''Convert MAC address as binary to text'''
    x = b':'.join(binascii.hexlify(binary)[i:i + 2]
                  for i in range(0, 12, 2))
    return str(x.decode('ascii'))

# input - mac in dec from sw_interface_dump
# output - regular mac in hex
def Convert_Dec_MAC_To_Hex(mac):
    hexmac=[]
    for num in mac[:6]:
        hexmac.append("%02x" % num)
    hexmac = ":".join(hexmac)
    return hexmac

# input - output from show memif intf command
# output - state info list
def Parse_Memif_Info(info):
    state = []
    socket_id = ''
    sockets_line = []
    for line in info.splitlines():
        if line:
            try:
                _ = int(line.strip().split()[0])
                sockets_line.append(line)
            except ValueError:
                pass
            if (line.strip().split()[0] == "flags"):
                if "admin-up" in line:
                    state.append("enabled=1")
                if "slave" in line:
                    state.append("role=slave")
                if "connected" in line:
                    state.append("connected=1")
            if (line.strip().split()[0] == "socket-id"):
                try:
                    socket_id = int(line.strip().split()[1])
                    state.append("id="+line.strip().split()[3])
                    for sock_line in sockets_line:
                      try:
                           num = int(sock_line.strip().split()[0])
                           if (num == socket_id):
                               state.append("socket=" + sock_line.strip().split()[-1])
                      except ValueError:
                           pass
                except ValueError:
                    pass
    if "enabled=1" not in state:
        state.append("enabled=0")
    if "role=slave" not in state:
        state.append("role=master")
    if "connected=1" not in state:
        state.append("connected=0")
    return state

# input - output from show br br_id detail command
# output - state info list
def Parse_BD_Details(details):
    state = []
    details = "\n".join([s for s in details.splitlines(True) if s.strip("\r\n")])
    line = details.splitlines()[1]
    if (line.strip().split()[6]) in ("on", "flood"):
        state.append("unicast=1")
    else:
        state.append("unicast=0")
    if (line.strip().split()[8]) == "on":
        state.append("arp_term=1")
    else:
        state.append("arp_term=0")
    return state

# input - etcd dump
# output - etcd dump converted to json + key, node, name, type atributes
def Convert_ETCD_Dump_To_JSON(dump):
    etcd_json = '['
    key = ''
    data = ''
    firstline = True
    for line in dump.splitlines():
        if line.strip() != '':
            if line[0] == '/':
                if not firstline:
                    etcd_json += '{"key":"'+key+'","node":"'+node+'","name":"'+name+'","type":"'+type+'","data":'+data+'},'
                key = line
                node = key.split('/')[2]
                name = key.split('/')[-1]
                type = key.split('/')[4]
                data = ''
                firstline = False
            else:
                if line == "null":
                    line = '{"error":"null"}'
                data += line
    if not firstline:
        etcd_json += '{"key":"'+key+'","node":"'+node+'","name":"'+name+'","type":"'+type+'","data":'+data+'}'
    etcd_json += ']'
    return etcd_json

# input - node name, bd name, etcd dump converted to json, bridge domain dump
# output - list of interfaces (etcd names) in bd
def Parse_BD_Interfaces(node, bd, etcd_json, bd_dump):
    interfaces = []
    bd_dump = json.loads(bd_dump)
    etcd_json = json.loads(etcd_json)
    for int in bd_dump[0]["sw_if"]:
        bd_sw_if_index =  int["sw_if_index"]
        etcd_name = "none"
        for key_data in etcd_json:
            if key_data["node"] == node and key_data["type"] == "status" and "/interface/" in key_data["key"]:
                if "if_index" in key_data["data"]:
                    if key_data["data"]["if_index"] == bd_sw_if_index:
                        etcd_name = key_data["data"]["name"]
        interfaces.append("interface="+etcd_name)
    if bd_dump[0]["bvi_sw_if_index"] != 4294967295:
        bvi_sw_if_index = bd_dump[0]["bvi_sw_if_index"]
        etcd_name = "none"
        for key_data in etcd_json:
            if key_data["node"] == node and key_data["type"] == "status" and "/interface/" in key_data["key"]:
                if "if_index" in key_data["data"]:
                    if key_data["data"]["if_index"] == bvi_sw_if_index:
                        etcd_name = key_data["data"]["name"]
        interfaces.append("bvi_int="+etcd_name)
    else:
        interfaces.append("bvi_int=none")
    return interfaces

# input - bridge domain dump, interfaces indexes
# output - true if bd with int indexes exists, false id bd not exists
def Check_BD_Presence(bd_dump, indexes):
    bd_dump = json.loads(bd_dump)
    present = False
    for bd in bd_dump:
        bd_present = True
        for index in indexes:
            int_present = False
            for bd_int in bd["sw_if"]:
                if bd_int["sw_if_index"] == index:
                    int_present = True
            if int_present == False:
                bd_present = False
        if bd_present == True:
            present = True
    return present

# def _revert_api_reply(api_r):
#     """Process API reply / a part of API reply.
#     Apply binascii.unhexlify() method for unicode values.
#     TODO: Remove the disabled code when definitely not needed.
#     :param api_r: API reply.
#     :type api_r: dict
#     :returns: Processed API reply / a part of API reply.
#     :rtype: dict
#     """
#     reply_dict = dict()
#     reply_value = dict()
#     for reply_key, reply_v in api_r.iteritems():
#         for a_k, a_v in reply_v.iteritems():
#             value = binascii.unhexlify(a_v) if isinstance(a_v, unicode) \
#                 else a_v
#             reply_value[a_k] = value
#             reply_value[a_k] = a_v
#         reply_dict[reply_key] = reply_value
#     return reply_dict

def _revert_api_reply(api_r):
    """Process API reply / a part of API reply.
    Apply binascii.unhexlify() method for unicode values.
    TODO: Implement complex solution to process of replies.
    :param api_r: API reply.
    :type api_r: dict
    :returns: Processed API reply / a part of API reply.
    :rtype: dict
    """
    reply_dict = dict()
    reply_value = dict()
    for reply_key, reply_v in api_r.iteritems():
        for a_k, a_v in reply_v.iteritems():
            reply_value[a_k] = a_v
        reply_dict[reply_key] = reply_value
    return reply_dict

def _process_reply(api_reply):
    """Process API reply.
    :param api_reply: API reply.
    :type api_reply: dict or list of dict
    :returns: Processed API reply.
    :rtype: list or dict
    """
    if isinstance(api_reply, list):
        reverted_reply = [_revert_api_reply(a_r) for a_r in api_reply]
    else:
        reverted_reply = _revert_api_reply(api_reply)
    return reverted_reply

# def _process_reply_1(out)
#     papi_reply = list()
#     for out in json_data:
#         try:
#             json_data = json.loads(out)
#         except ValueError:
#             logger.error("An error occured while processing the PAPI "
#                          "request:\n{rqst}".format(rqst=local_list))
#             raise
#         for data in json_data:
#             try:
#                 api_reply_processed = dict(
#                     api_name=data["api_name"],
#                     api_reply=self._process_reply(data["api_reply"]))
#             except KeyError:
#                 if ignore_errors:
#                     continue
#                 else:
#                     raise
#         papi_reply.append(api_reply_processed)
