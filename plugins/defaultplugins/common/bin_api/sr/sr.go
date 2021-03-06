// Code generated by govpp binapi-generator DO NOT EDIT.
// Package sr represents the VPP binary API of the 'sr' VPP module.
// Generated from '/usr/share/vpp/api/sr.api.json'
package sr

import "git.fd.io/govpp.git/api"

// SrLocalsidAddDel represents the VPP binary API message 'sr_localsid_add_del'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 30:
//
//            "sr_localsid_add_del",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "is_del"
//            ],
//            [
//                "u8",
//                "localsid_addr",
//                16
//            ],
//            [
//                "u8",
//                "end_psp"
//            ],
//            [
//                "u8",
//                "behavior"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u32",
//                "vlan_index"
//            ],
//            [
//                "u32",
//                "fib_table"
//            ],
//            [
//                "u8",
//                "nh_addr",
//                16
//            ],
//            {
//                "crc": "0xa833a891"
//            }
//
type SrLocalsidAddDel struct {
	IsDel        uint8
	LocalsidAddr []byte `struc:"[16]byte"`
	EndPsp       uint8
	Behavior     uint8
	SwIfIndex    uint32
	VlanIndex    uint32
	FibTable     uint32
	NhAddr       []byte `struc:"[16]byte"`
}

func (*SrLocalsidAddDel) GetMessageName() string {
	return "sr_localsid_add_del"
}
func (*SrLocalsidAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrLocalsidAddDel) GetCrcString() string {
	return "a833a891"
}
func NewSrLocalsidAddDel() api.Message {
	return &SrLocalsidAddDel{}
}

// SrLocalsidAddDelReply represents the VPP binary API message 'sr_localsid_add_del_reply'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 82:
//
//            "sr_localsid_add_del_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type SrLocalsidAddDelReply struct {
	Retval int32
}

func (*SrLocalsidAddDelReply) GetMessageName() string {
	return "sr_localsid_add_del_reply"
}
func (*SrLocalsidAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrLocalsidAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func NewSrLocalsidAddDelReply() api.Message {
	return &SrLocalsidAddDelReply{}
}

// SrPolicyAdd represents the VPP binary API message 'sr_policy_add'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 100:
//
//            "sr_policy_add",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "bsid_addr",
//                16
//            ],
//            [
//                "u32",
//                "weight"
//            ],
//            [
//                "u8",
//                "is_encap"
//            ],
//            [
//                "u8",
//                "type"
//            ],
//            [
//                "u32",
//                "fib_table"
//            ],
//            [
//                "u8",
//                "n_segments"
//            ],
//            [
//                "u8",
//                "segments",
//                0
//            ],
//            {
//                "crc": "0xc8c2222d"
//            }
//
type SrPolicyAdd struct {
	BsidAddr  []byte `struc:"[16]byte"`
	Weight    uint32
	IsEncap   uint8
	Type      uint8
	FibTable  uint32
	NSegments uint8      `struc:"sizeof=Segments"` // MANUALLY ADDED TO FIX MARSHALLING (BAD VPP API)
	Segments  []IPv6type // MANUALLY ADDED TO FIX MARSHALLING (BAD VPP API)
}

func (*SrPolicyAdd) GetMessageName() string {
	return "sr_policy_add"
}
func (*SrPolicyAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrPolicyAdd) GetCrcString() string {
	return "c8c2222d"
}
func NewSrPolicyAdd() api.Message {
	return &SrPolicyAdd{}
}

// SrPolicyAddReply represents the VPP binary API message 'sr_policy_add_reply'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 148:
//
//            "sr_policy_add_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type SrPolicyAddReply struct {
	Retval int32
}

func (*SrPolicyAddReply) GetMessageName() string {
	return "sr_policy_add_reply"
}
func (*SrPolicyAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrPolicyAddReply) GetCrcString() string {
	return "e8d4e804"
}
func NewSrPolicyAddReply() api.Message {
	return &SrPolicyAddReply{}
}

// SrPolicyMod represents the VPP binary API message 'sr_policy_mod'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 166:
//
//            "sr_policy_mod",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "bsid_addr",
//                16
//            ],
//            [
//                "u32",
//                "sr_policy_index"
//            ],
//            [
//                "u32",
//                "fib_table"
//            ],
//            [
//                "u8",
//                "operation"
//            ],
//            [
//                "u32",
//                "sl_index"
//            ],
//            [
//                "u32",
//                "weight"
//            ],
//            [
//                "u8",
//                "n_segments"
//            ],
//            [
//                "u8",
//                "segments",
//                0
//            ],
//            {
//                "crc": "0x596a4682"
//            }
//
type SrPolicyMod struct {
	BsidAddr      []byte `struc:"[16]byte"`
	SrPolicyIndex uint32
	FibTable      uint32
	Operation     uint8
	SlIndex       uint32
	Weight        uint32
	NSegments     uint8      `struc:"sizeof=Segments"` // MANUALLY ADDED TO FIX MARSHALLING (BAD VPP API)
	Segments      []IPv6type // MANUALLY ADDED TO FIX MARSHALLING (BAD VPP API)
}

func (*SrPolicyMod) GetMessageName() string {
	return "sr_policy_mod"
}
func (*SrPolicyMod) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrPolicyMod) GetCrcString() string {
	return "596a4682"
}
func NewSrPolicyMod() api.Message {
	return &SrPolicyMod{}
}

// SrPolicyModReply represents the VPP binary API message 'sr_policy_mod_reply'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 218:
//
//            "sr_policy_mod_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type SrPolicyModReply struct {
	Retval int32
}

func (*SrPolicyModReply) GetMessageName() string {
	return "sr_policy_mod_reply"
}
func (*SrPolicyModReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrPolicyModReply) GetCrcString() string {
	return "e8d4e804"
}
func NewSrPolicyModReply() api.Message {
	return &SrPolicyModReply{}
}

// SrPolicyDel represents the VPP binary API message 'sr_policy_del'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 236:
//
//            "sr_policy_del",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "bsid_addr",
//                16
//            ],
//            [
//                "u32",
//                "sr_policy_index"
//            ],
//            {
//                "crc": "0x0388e561"
//            }
//
type SrPolicyDel struct {
	BsidAddr      []byte `struc:"[16]byte"`
	SrPolicyIndex uint32
}

func (*SrPolicyDel) GetMessageName() string {
	return "sr_policy_del"
}
func (*SrPolicyDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrPolicyDel) GetCrcString() string {
	return "0388e561"
}
func NewSrPolicyDel() api.Message {
	return &SrPolicyDel{}
}

// SrPolicyDelReply represents the VPP binary API message 'sr_policy_del_reply'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 263:
//
//            "sr_policy_del_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type SrPolicyDelReply struct {
	Retval int32
}

func (*SrPolicyDelReply) GetMessageName() string {
	return "sr_policy_del_reply"
}
func (*SrPolicyDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrPolicyDelReply) GetCrcString() string {
	return "e8d4e804"
}
func NewSrPolicyDelReply() api.Message {
	return &SrPolicyDelReply{}
}

// SrSetEncapSource represents the VPP binary API message 'sr_set_encap_source'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 281:
//
//            "sr_set_encap_source",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "encaps_source",
//                16
//            ],
//            {
//                "crc": "0xd05bb4de"
//            }
//
type SrSetEncapSource struct {
	EncapsSource []byte `struc:"[16]byte"`
}

func (*SrSetEncapSource) GetMessageName() string {
	return "sr_set_encap_source"
}
func (*SrSetEncapSource) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrSetEncapSource) GetCrcString() string {
	return "d05bb4de"
}
func NewSrSetEncapSource() api.Message {
	return &SrSetEncapSource{}
}

// SrSetEncapSourceReply represents the VPP binary API message 'sr_set_encap_source_reply'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 304:
//
//            "sr_set_encap_source_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type SrSetEncapSourceReply struct {
	Retval int32
}

func (*SrSetEncapSourceReply) GetMessageName() string {
	return "sr_set_encap_source_reply"
}
func (*SrSetEncapSourceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrSetEncapSourceReply) GetCrcString() string {
	return "e8d4e804"
}
func NewSrSetEncapSourceReply() api.Message {
	return &SrSetEncapSourceReply{}
}

// SrSteeringAddDel represents the VPP binary API message 'sr_steering_add_del'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 322:
//
//            "sr_steering_add_del",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "is_del"
//            ],
//            [
//                "u8",
//                "bsid_addr",
//                16
//            ],
//            [
//                "u32",
//                "sr_policy_index"
//            ],
//            [
//                "u32",
//                "table_id"
//            ],
//            [
//                "u8",
//                "prefix_addr",
//                16
//            ],
//            [
//                "u32",
//                "mask_width"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u8",
//                "traffic_type"
//            ],
//            {
//                "crc": "0x28b5dcab"
//            }
//
type SrSteeringAddDel struct {
	IsDel         uint8
	BsidAddr      []byte `struc:"[16]byte"`
	SrPolicyIndex uint32
	TableID       uint32
	PrefixAddr    []byte `struc:"[16]byte"`
	MaskWidth     uint32
	SwIfIndex     uint32
	TrafficType   uint8
}

func (*SrSteeringAddDel) GetMessageName() string {
	return "sr_steering_add_del"
}
func (*SrSteeringAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrSteeringAddDel) GetCrcString() string {
	return "28b5dcab"
}
func NewSrSteeringAddDel() api.Message {
	return &SrSteeringAddDel{}
}

// SrSteeringAddDelReply represents the VPP binary API message 'sr_steering_add_del_reply'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 374:
//
//            "sr_steering_add_del_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type SrSteeringAddDelReply struct {
	Retval int32
}

func (*SrSteeringAddDelReply) GetMessageName() string {
	return "sr_steering_add_del_reply"
}
func (*SrSteeringAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrSteeringAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func NewSrSteeringAddDelReply() api.Message {
	return &SrSteeringAddDelReply{}
}

// SrLocalsidsDump represents the VPP binary API message 'sr_localsids_dump'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 392:
//
//            "sr_localsids_dump",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type SrLocalsidsDump struct {
}

func (*SrLocalsidsDump) GetMessageName() string {
	return "sr_localsids_dump"
}
func (*SrLocalsidsDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrLocalsidsDump) GetCrcString() string {
	return "51077d14"
}
func NewSrLocalsidsDump() api.Message {
	return &SrLocalsidsDump{}
}

// SrLocalsidsDetails represents the VPP binary API message 'sr_localsids_details'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 410:
//
//            "sr_localsids_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "address",
//                16
//            ],
//            [
//                "u8",
//                "end_psp"
//            ],
//            [
//                "u16",
//                "behavior"
//            ],
//            [
//                "u32",
//                "fib_table"
//            ],
//            [
//                "u8",
//                "xconnect_next_hop",
//                16
//            ],
//            [
//                "u32",
//                "xconnect_iface_or_vrf_table"
//            ],
//            {
//                "crc": "0xb6556a9c"
//            }
//
type SrLocalsidsDetails struct {
	Address                 []byte `struc:"[16]byte"`
	EndPsp                  uint8
	Behavior                uint16
	FibTable                uint32
	XconnectNextHop         []byte `struc:"[16]byte"`
	XconnectIfaceOrVrfTable uint32
}

func (*SrLocalsidsDetails) GetMessageName() string {
	return "sr_localsids_details"
}
func (*SrLocalsidsDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrLocalsidsDetails) GetCrcString() string {
	return "b6556a9c"
}
func NewSrLocalsidsDetails() api.Message {
	return &SrLocalsidsDetails{}
}

type IPv6type struct{ Value [16]byte } // MANUALLY ADDED TO FIX MARSHALLING (BAD VPP API)
