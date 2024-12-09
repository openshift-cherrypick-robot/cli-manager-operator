// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ServiceCatalogControllerManagerStatusApplyConfiguration represents a declarative configuration of the ServiceCatalogControllerManagerStatus type for use
// with apply.
type ServiceCatalogControllerManagerStatusApplyConfiguration struct {
	OperatorStatusApplyConfiguration `json:",inline"`
}

// ServiceCatalogControllerManagerStatusApplyConfiguration constructs a declarative configuration of the ServiceCatalogControllerManagerStatus type for use with
// apply.
func ServiceCatalogControllerManagerStatus() *ServiceCatalogControllerManagerStatusApplyConfiguration {
	return &ServiceCatalogControllerManagerStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *ServiceCatalogControllerManagerStatusApplyConfiguration) WithObservedGeneration(value int64) *ServiceCatalogControllerManagerStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ServiceCatalogControllerManagerStatusApplyConfiguration) WithConditions(values ...*OperatorConditionApplyConfiguration) *ServiceCatalogControllerManagerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *ServiceCatalogControllerManagerStatusApplyConfiguration) WithVersion(value string) *ServiceCatalogControllerManagerStatusApplyConfiguration {
	b.Version = &value
	return b
}

// WithReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyReplicas field is set to the value of the last call.
func (b *ServiceCatalogControllerManagerStatusApplyConfiguration) WithReadyReplicas(value int32) *ServiceCatalogControllerManagerStatusApplyConfiguration {
	b.ReadyReplicas = &value
	return b
}

// WithLatestAvailableRevision sets the LatestAvailableRevision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LatestAvailableRevision field is set to the value of the last call.
func (b *ServiceCatalogControllerManagerStatusApplyConfiguration) WithLatestAvailableRevision(value int32) *ServiceCatalogControllerManagerStatusApplyConfiguration {
	b.LatestAvailableRevision = &value
	return b
}

// WithGenerations adds the given value to the Generations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Generations field.
func (b *ServiceCatalogControllerManagerStatusApplyConfiguration) WithGenerations(values ...*GenerationStatusApplyConfiguration) *ServiceCatalogControllerManagerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGenerations")
		}
		b.Generations = append(b.Generations, *values[i])
	}
	return b
}