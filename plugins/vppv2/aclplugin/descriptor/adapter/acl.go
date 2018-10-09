// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/gogo/protobuf/proto"
	. "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/plugins/vppv2/aclplugin/aclidx"
	"github.com/ligato/vpp-agent/plugins/vppv2/model/acl"
)

////////// type-safe key-value pair with metadata //////////

type AclKVWithMetadata struct {
	Key      string
	Value    *acl.Acl
	Metadata *aclidx.AclMetadata
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type AclDescriptor struct {
	Name               string
	KeySelector        KeySelector
	ValueTypeName      string
	KeyLabel           func(key string) string
	ValueComparator    func(key string, oldValue, newValue *acl.Acl) bool
	NBKeyPrefix        string
	WithMetadata       bool
	MetadataMapFactory MetadataMapFactory
	Add                func(key string, value *acl.Acl) (metadata *aclidx.AclMetadata, err error)
	Delete             func(key string, value *acl.Acl, metadata *aclidx.AclMetadata) error
	Modify             func(key string, oldValue, newValue *acl.Acl, oldMetadata *aclidx.AclMetadata) (newMetadata *aclidx.AclMetadata, err error)
	ModifyWithRecreate func(key string, oldValue, newValue *acl.Acl, metadata *aclidx.AclMetadata) bool
	Update             func(key string, value *acl.Acl, metadata *aclidx.AclMetadata) error
	IsRetriableFailure func(err error) bool
	Dependencies       func(key string, value *acl.Acl) []Dependency
	DerivedValues      func(key string, value *acl.Acl) []KeyValuePair
	Dump               func(correlate []AclKVWithMetadata) ([]AclKVWithMetadata, error)
	DumpDependencies   []string /* descriptor name */
}

////////// Descriptor adapter //////////

type AclDescriptorAdapter struct {
	descriptor *AclDescriptor
}

func NewAclDescriptor(typedDescriptor *AclDescriptor) *KVDescriptor {
	adapter := &AclDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:               typedDescriptor.Name,
		KeySelector:        typedDescriptor.KeySelector,
		ValueTypeName:      typedDescriptor.ValueTypeName,
		KeyLabel:           typedDescriptor.KeyLabel,
		NBKeyPrefix:        typedDescriptor.NBKeyPrefix,
		WithMetadata:       typedDescriptor.WithMetadata,
		MetadataMapFactory: typedDescriptor.MetadataMapFactory,
		IsRetriableFailure: typedDescriptor.IsRetriableFailure,
		DumpDependencies:   typedDescriptor.DumpDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Add != nil {
		descriptor.Add = adapter.Add
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Modify != nil {
		descriptor.Modify = adapter.Modify
	}
	if typedDescriptor.ModifyWithRecreate != nil {
		descriptor.ModifyWithRecreate = adapter.ModifyWithRecreate
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	if typedDescriptor.Dump != nil {
		descriptor.Dump = adapter.Dump
	}
	return descriptor
}

func (da *AclDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castAclValue(key, oldValue)
	typedNewValue, err2 := castAclValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *AclDescriptorAdapter) Add(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castAclValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Add(key, typedValue)
}

func (da *AclDescriptorAdapter) Modify(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castAclValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castAclValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castAclMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Modify(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *AclDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castAclValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castAclMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *AclDescriptorAdapter) ModifyWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castAclValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castAclValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castAclMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.ModifyWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *AclDescriptorAdapter) Update(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castAclValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castAclMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Update(key, typedValue, typedMetadata)
}

func (da *AclDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castAclValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

func (da *AclDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castAclValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *AclDescriptorAdapter) Dump(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []AclKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castAclValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castAclMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			AclKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedDump, err := da.descriptor.Dump(correlateWithType)
	if err != nil {
		return nil, err
	}
	var dump []KVWithMetadata
	for _, typedKVWithMetadata := range typedDump {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		dump = append(dump, kvWithMetadata)
	}
	return dump, err
}

////////// Helper methods //////////

func castAclValue(key string, value proto.Message) (*acl.Acl, error) {
	typedValue, ok := value.(*acl.Acl)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castAclMetadata(key string, metadata Metadata) (*aclidx.AclMetadata, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(*aclidx.AclMetadata)
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
