// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// GCPCSIDriverConfigSpecApplyConfiguration represents a declarative configuration of the GCPCSIDriverConfigSpec type for use
// with apply.
type GCPCSIDriverConfigSpecApplyConfiguration struct {
	KMSKey *GCPKMSKeyReferenceApplyConfiguration `json:"kmsKey,omitempty"`
}

// GCPCSIDriverConfigSpecApplyConfiguration constructs a declarative configuration of the GCPCSIDriverConfigSpec type for use with
// apply.
func GCPCSIDriverConfigSpec() *GCPCSIDriverConfigSpecApplyConfiguration {
	return &GCPCSIDriverConfigSpecApplyConfiguration{}
}

// WithKMSKey sets the KMSKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KMSKey field is set to the value of the last call.
func (b *GCPCSIDriverConfigSpecApplyConfiguration) WithKMSKey(value *GCPKMSKeyReferenceApplyConfiguration) *GCPCSIDriverConfigSpecApplyConfiguration {
	b.KMSKey = value
	return b
}