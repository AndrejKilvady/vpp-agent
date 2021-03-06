// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package ipsec

import "reflect"

var Types = map[string]reflect.Type{
	"Ikev2InitiateDelChildSa":        reflect.TypeOf((*Ikev2InitiateDelChildSa)(nil)).Elem(),
	"Ikev2InitiateDelChildSaReply":   reflect.TypeOf((*Ikev2InitiateDelChildSaReply)(nil)).Elem(),
	"Ikev2InitiateDelIkeSa":          reflect.TypeOf((*Ikev2InitiateDelIkeSa)(nil)).Elem(),
	"Ikev2InitiateDelIkeSaReply":     reflect.TypeOf((*Ikev2InitiateDelIkeSaReply)(nil)).Elem(),
	"Ikev2InitiateRekeyChildSa":      reflect.TypeOf((*Ikev2InitiateRekeyChildSa)(nil)).Elem(),
	"Ikev2InitiateRekeyChildSaReply": reflect.TypeOf((*Ikev2InitiateRekeyChildSaReply)(nil)).Elem(),
	"Ikev2InitiateSaInit":            reflect.TypeOf((*Ikev2InitiateSaInit)(nil)).Elem(),
	"Ikev2InitiateSaInitReply":       reflect.TypeOf((*Ikev2InitiateSaInitReply)(nil)).Elem(),
	"Ikev2ProfileAddDel":             reflect.TypeOf((*Ikev2ProfileAddDel)(nil)).Elem(),
	"Ikev2ProfileAddDelReply":        reflect.TypeOf((*Ikev2ProfileAddDelReply)(nil)).Elem(),
	"Ikev2ProfileSetAuth":            reflect.TypeOf((*Ikev2ProfileSetAuth)(nil)).Elem(),
	"Ikev2ProfileSetAuthReply":       reflect.TypeOf((*Ikev2ProfileSetAuthReply)(nil)).Elem(),
	"Ikev2ProfileSetID":              reflect.TypeOf((*Ikev2ProfileSetID)(nil)).Elem(),
	"Ikev2ProfileSetIDReply":         reflect.TypeOf((*Ikev2ProfileSetIDReply)(nil)).Elem(),
	"Ikev2ProfileSetTs":              reflect.TypeOf((*Ikev2ProfileSetTs)(nil)).Elem(),
	"Ikev2ProfileSetTsReply":         reflect.TypeOf((*Ikev2ProfileSetTsReply)(nil)).Elem(),
	"Ikev2SetEspTransforms":          reflect.TypeOf((*Ikev2SetEspTransforms)(nil)).Elem(),
	"Ikev2SetEspTransformsReply":     reflect.TypeOf((*Ikev2SetEspTransformsReply)(nil)).Elem(),
	"Ikev2SetIkeTransforms":          reflect.TypeOf((*Ikev2SetIkeTransforms)(nil)).Elem(),
	"Ikev2SetIkeTransformsReply":     reflect.TypeOf((*Ikev2SetIkeTransformsReply)(nil)).Elem(),
	"Ikev2SetLocalKey":               reflect.TypeOf((*Ikev2SetLocalKey)(nil)).Elem(),
	"Ikev2SetLocalKeyReply":          reflect.TypeOf((*Ikev2SetLocalKeyReply)(nil)).Elem(),
	"Ikev2SetResponder":              reflect.TypeOf((*Ikev2SetResponder)(nil)).Elem(),
	"Ikev2SetResponderReply":         reflect.TypeOf((*Ikev2SetResponderReply)(nil)).Elem(),
	"Ikev2SetSaLifetime":             reflect.TypeOf((*Ikev2SetSaLifetime)(nil)).Elem(),
	"Ikev2SetSaLifetimeReply":        reflect.TypeOf((*Ikev2SetSaLifetimeReply)(nil)).Elem(),
	"IpsecInterfaceAddDelSpd":        reflect.TypeOf((*IpsecInterfaceAddDelSpd)(nil)).Elem(),
	"IpsecInterfaceAddDelSpdReply":   reflect.TypeOf((*IpsecInterfaceAddDelSpdReply)(nil)).Elem(),
	"IpsecSaDetails":                 reflect.TypeOf((*IpsecSaDetails)(nil)).Elem(),
	"IpsecSaDump":                    reflect.TypeOf((*IpsecSaDump)(nil)).Elem(),
	"IpsecSaSetKey":                  reflect.TypeOf((*IpsecSaSetKey)(nil)).Elem(),
	"IpsecSaSetKeyReply":             reflect.TypeOf((*IpsecSaSetKeyReply)(nil)).Elem(),
	"IpsecSadAddDelEntry":            reflect.TypeOf((*IpsecSadAddDelEntry)(nil)).Elem(),
	"IpsecSadAddDelEntryReply":       reflect.TypeOf((*IpsecSadAddDelEntryReply)(nil)).Elem(),
	"IpsecSpdAddDel":                 reflect.TypeOf((*IpsecSpdAddDel)(nil)).Elem(),
	"IpsecSpdAddDelEntry":            reflect.TypeOf((*IpsecSpdAddDelEntry)(nil)).Elem(),
	"IpsecSpdAddDelEntryReply":       reflect.TypeOf((*IpsecSpdAddDelEntryReply)(nil)).Elem(),
	"IpsecSpdAddDelReply":            reflect.TypeOf((*IpsecSpdAddDelReply)(nil)).Elem(),
	"IpsecSpdDetails":                reflect.TypeOf((*IpsecSpdDetails)(nil)).Elem(),
	"IpsecSpdDump":                   reflect.TypeOf((*IpsecSpdDump)(nil)).Elem(),
	"IpsecTunnelIfAddDel":            reflect.TypeOf((*IpsecTunnelIfAddDel)(nil)).Elem(),
	"IpsecTunnelIfAddDelReply":       reflect.TypeOf((*IpsecTunnelIfAddDelReply)(nil)).Elem(),
	"IpsecTunnelIfSetKey":            reflect.TypeOf((*IpsecTunnelIfSetKey)(nil)).Elem(),
	"IpsecTunnelIfSetKeyReply":       reflect.TypeOf((*IpsecTunnelIfSetKeyReply)(nil)).Elem(),
	"IpsecTunnelIfSetSa":             reflect.TypeOf((*IpsecTunnelIfSetSa)(nil)).Elem(),
	"IpsecTunnelIfSetSaReply":        reflect.TypeOf((*IpsecTunnelIfSetSaReply)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"NewIkev2InitiateDelChildSa":        reflect.ValueOf(NewIkev2InitiateDelChildSa),
	"NewIkev2InitiateDelChildSaReply":   reflect.ValueOf(NewIkev2InitiateDelChildSaReply),
	"NewIkev2InitiateDelIkeSa":          reflect.ValueOf(NewIkev2InitiateDelIkeSa),
	"NewIkev2InitiateDelIkeSaReply":     reflect.ValueOf(NewIkev2InitiateDelIkeSaReply),
	"NewIkev2InitiateRekeyChildSa":      reflect.ValueOf(NewIkev2InitiateRekeyChildSa),
	"NewIkev2InitiateRekeyChildSaReply": reflect.ValueOf(NewIkev2InitiateRekeyChildSaReply),
	"NewIkev2InitiateSaInit":            reflect.ValueOf(NewIkev2InitiateSaInit),
	"NewIkev2InitiateSaInitReply":       reflect.ValueOf(NewIkev2InitiateSaInitReply),
	"NewIkev2ProfileAddDel":             reflect.ValueOf(NewIkev2ProfileAddDel),
	"NewIkev2ProfileAddDelReply":        reflect.ValueOf(NewIkev2ProfileAddDelReply),
	"NewIkev2ProfileSetAuth":            reflect.ValueOf(NewIkev2ProfileSetAuth),
	"NewIkev2ProfileSetAuthReply":       reflect.ValueOf(NewIkev2ProfileSetAuthReply),
	"NewIkev2ProfileSetID":              reflect.ValueOf(NewIkev2ProfileSetID),
	"NewIkev2ProfileSetIDReply":         reflect.ValueOf(NewIkev2ProfileSetIDReply),
	"NewIkev2ProfileSetTs":              reflect.ValueOf(NewIkev2ProfileSetTs),
	"NewIkev2ProfileSetTsReply":         reflect.ValueOf(NewIkev2ProfileSetTsReply),
	"NewIkev2SetEspTransforms":          reflect.ValueOf(NewIkev2SetEspTransforms),
	"NewIkev2SetEspTransformsReply":     reflect.ValueOf(NewIkev2SetEspTransformsReply),
	"NewIkev2SetIkeTransforms":          reflect.ValueOf(NewIkev2SetIkeTransforms),
	"NewIkev2SetIkeTransformsReply":     reflect.ValueOf(NewIkev2SetIkeTransformsReply),
	"NewIkev2SetLocalKey":               reflect.ValueOf(NewIkev2SetLocalKey),
	"NewIkev2SetLocalKeyReply":          reflect.ValueOf(NewIkev2SetLocalKeyReply),
	"NewIkev2SetResponder":              reflect.ValueOf(NewIkev2SetResponder),
	"NewIkev2SetResponderReply":         reflect.ValueOf(NewIkev2SetResponderReply),
	"NewIkev2SetSaLifetime":             reflect.ValueOf(NewIkev2SetSaLifetime),
	"NewIkev2SetSaLifetimeReply":        reflect.ValueOf(NewIkev2SetSaLifetimeReply),
	"NewIpsecInterfaceAddDelSpd":        reflect.ValueOf(NewIpsecInterfaceAddDelSpd),
	"NewIpsecInterfaceAddDelSpdReply":   reflect.ValueOf(NewIpsecInterfaceAddDelSpdReply),
	"NewIpsecSaDetails":                 reflect.ValueOf(NewIpsecSaDetails),
	"NewIpsecSaDump":                    reflect.ValueOf(NewIpsecSaDump),
	"NewIpsecSaSetKey":                  reflect.ValueOf(NewIpsecSaSetKey),
	"NewIpsecSaSetKeyReply":             reflect.ValueOf(NewIpsecSaSetKeyReply),
	"NewIpsecSadAddDelEntry":            reflect.ValueOf(NewIpsecSadAddDelEntry),
	"NewIpsecSadAddDelEntryReply":       reflect.ValueOf(NewIpsecSadAddDelEntryReply),
	"NewIpsecSpdAddDel":                 reflect.ValueOf(NewIpsecSpdAddDel),
	"NewIpsecSpdAddDelEntry":            reflect.ValueOf(NewIpsecSpdAddDelEntry),
	"NewIpsecSpdAddDelEntryReply":       reflect.ValueOf(NewIpsecSpdAddDelEntryReply),
	"NewIpsecSpdAddDelReply":            reflect.ValueOf(NewIpsecSpdAddDelReply),
	"NewIpsecSpdDetails":                reflect.ValueOf(NewIpsecSpdDetails),
	"NewIpsecSpdDump":                   reflect.ValueOf(NewIpsecSpdDump),
	"NewIpsecTunnelIfAddDel":            reflect.ValueOf(NewIpsecTunnelIfAddDel),
	"NewIpsecTunnelIfAddDelReply":       reflect.ValueOf(NewIpsecTunnelIfAddDelReply),
	"NewIpsecTunnelIfSetKey":            reflect.ValueOf(NewIpsecTunnelIfSetKey),
	"NewIpsecTunnelIfSetKeyReply":       reflect.ValueOf(NewIpsecTunnelIfSetKeyReply),
	"NewIpsecTunnelIfSetSa":             reflect.ValueOf(NewIpsecTunnelIfSetSa),
	"NewIpsecTunnelIfSetSaReply":        reflect.ValueOf(NewIpsecTunnelIfSetSaReply),
}

var Variables = map[string]reflect.Value{}

var Consts = map[string]reflect.Value{}
