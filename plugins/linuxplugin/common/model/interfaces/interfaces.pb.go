// Code generated by protoc-gen-gogo.
// source: interfaces/interfaces.proto
// DO NOT EDIT!

/*
Package interfaces is a generated protocol buffer package.

It is generated from these files:
	interfaces/interfaces.proto

It has these top-level messages:
	LinuxInterfaces
*/
package interfaces

import proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type LinuxInterfaces_InterfaceType int32

const (
	LinuxInterfaces_VETH     LinuxInterfaces_InterfaceType = 0
	LinuxInterfaces_AUTO_TAP LinuxInterfaces_InterfaceType = 1
)

var LinuxInterfaces_InterfaceType_name = map[int32]string{
	0: "VETH",
	1: "AUTO_TAP",
}
var LinuxInterfaces_InterfaceType_value = map[string]int32{
	"VETH":     0,
	"AUTO_TAP": 1,
}

func (x LinuxInterfaces_InterfaceType) String() string {
	return proto.EnumName(LinuxInterfaces_InterfaceType_name, int32(x))
}

type LinuxInterfaces_Interface_Namespace_NamespaceType int32

const (
	LinuxInterfaces_Interface_Namespace_PID_REF_NS          LinuxInterfaces_Interface_Namespace_NamespaceType = 0
	LinuxInterfaces_Interface_Namespace_MICROSERVICE_REF_NS LinuxInterfaces_Interface_Namespace_NamespaceType = 1
	LinuxInterfaces_Interface_Namespace_NAMED_NS            LinuxInterfaces_Interface_Namespace_NamespaceType = 2
	LinuxInterfaces_Interface_Namespace_FILE_REF_NS         LinuxInterfaces_Interface_Namespace_NamespaceType = 3
)

var LinuxInterfaces_Interface_Namespace_NamespaceType_name = map[int32]string{
	0: "PID_REF_NS",
	1: "MICROSERVICE_REF_NS",
	2: "NAMED_NS",
	3: "FILE_REF_NS",
}
var LinuxInterfaces_Interface_Namespace_NamespaceType_value = map[string]int32{
	"PID_REF_NS":          0,
	"MICROSERVICE_REF_NS": 1,
	"NAMED_NS":            2,
	"FILE_REF_NS":         3,
}

func (x LinuxInterfaces_Interface_Namespace_NamespaceType) String() string {
	return proto.EnumName(LinuxInterfaces_Interface_Namespace_NamespaceType_name, int32(x))
}

type LinuxInterfaces struct {
	Interface []*LinuxInterfaces_Interface `protobuf:"bytes,1,rep,name=interface" json:"interface,omitempty"`
}

func (m *LinuxInterfaces) Reset()         { *m = LinuxInterfaces{} }
func (m *LinuxInterfaces) String() string { return proto.CompactTextString(m) }
func (*LinuxInterfaces) ProtoMessage()    {}

func (m *LinuxInterfaces) GetInterface() []*LinuxInterfaces_Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

type LinuxInterfaces_Interface struct {
	Name        string                               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string                               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type        LinuxInterfaces_InterfaceType        `protobuf:"varint,3,opt,name=type,proto3,enum=interfaces.LinuxInterfaces_InterfaceType" json:"type,omitempty"`
	Enabled     bool                                 `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	IpAddresses []string                             `protobuf:"bytes,8,rep,name=ip_addresses" json:"ip_addresses,omitempty"`
	PhysAddress string                               `protobuf:"bytes,5,opt,name=phys_address,proto3" json:"phys_address,omitempty"`
	Mtu         uint32                               `protobuf:"varint,6,opt,name=mtu,proto3" json:"mtu,omitempty"`
	HostIfName  string                               `protobuf:"bytes,7,opt,name=host_if_name,proto3" json:"host_if_name,omitempty"`
	Namespace   *LinuxInterfaces_Interface_Namespace `protobuf:"bytes,50,opt,name=namespace" json:"namespace,omitempty"`
	Veth        *LinuxInterfaces_Interface_Veth      `protobuf:"bytes,51,opt,name=veth" json:"veth,omitempty"`
}

func (m *LinuxInterfaces_Interface) Reset()         { *m = LinuxInterfaces_Interface{} }
func (m *LinuxInterfaces_Interface) String() string { return proto.CompactTextString(m) }
func (*LinuxInterfaces_Interface) ProtoMessage()    {}

func (m *LinuxInterfaces_Interface) GetNamespace() *LinuxInterfaces_Interface_Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *LinuxInterfaces_Interface) GetVeth() *LinuxInterfaces_Interface_Veth {
	if m != nil {
		return m.Veth
	}
	return nil
}

type LinuxInterfaces_Interface_Namespace struct {
	Type         LinuxInterfaces_Interface_Namespace_NamespaceType `protobuf:"varint,1,opt,name=type,proto3,enum=interfaces.LinuxInterfaces_Interface_Namespace_NamespaceType" json:"type,omitempty"`
	Pid          uint32                                            `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Microservice string                                            `protobuf:"bytes,3,opt,name=microservice,proto3" json:"microservice,omitempty"`
	Name         string                                            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Filepath     string                                            `protobuf:"bytes,5,opt,name=filepath,proto3" json:"filepath,omitempty"`
}

func (m *LinuxInterfaces_Interface_Namespace) Reset()         { *m = LinuxInterfaces_Interface_Namespace{} }
func (m *LinuxInterfaces_Interface_Namespace) String() string { return proto.CompactTextString(m) }
func (*LinuxInterfaces_Interface_Namespace) ProtoMessage()    {}

type LinuxInterfaces_Interface_Veth struct {
	PeerIfName string `protobuf:"bytes,1,opt,name=peer_if_name,proto3" json:"peer_if_name,omitempty"`
}

func (m *LinuxInterfaces_Interface_Veth) Reset()         { *m = LinuxInterfaces_Interface_Veth{} }
func (m *LinuxInterfaces_Interface_Veth) String() string { return proto.CompactTextString(m) }
func (*LinuxInterfaces_Interface_Veth) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("interfaces.LinuxInterfaces_InterfaceType", LinuxInterfaces_InterfaceType_name, LinuxInterfaces_InterfaceType_value)
	proto.RegisterEnum("interfaces.LinuxInterfaces_Interface_Namespace_NamespaceType", LinuxInterfaces_Interface_Namespace_NamespaceType_name, LinuxInterfaces_Interface_Namespace_NamespaceType_value)
}
