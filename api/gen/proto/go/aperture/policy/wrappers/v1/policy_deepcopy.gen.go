// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package wrappersv1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using PolicyWrapper within kubernetes types, where deepcopy-gen is used.
func (in *PolicyWrapper) DeepCopyInto(out *PolicyWrapper) {
	p := proto.Clone(in).(*PolicyWrapper)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyWrapper. Required by controller-gen.
func (in *PolicyWrapper) DeepCopy() *PolicyWrapper {
	if in == nil {
		return nil
	}
	out := new(PolicyWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PolicyWrapper. Required by controller-gen.
func (in *PolicyWrapper) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PolicyWrappers within kubernetes types, where deepcopy-gen is used.
func (in *PolicyWrappers) DeepCopyInto(out *PolicyWrappers) {
	p := proto.Clone(in).(*PolicyWrappers)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyWrappers. Required by controller-gen.
func (in *PolicyWrappers) DeepCopy() *PolicyWrappers {
	if in == nil {
		return nil
	}
	out := new(PolicyWrappers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PolicyWrappers. Required by controller-gen.
func (in *PolicyWrappers) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using FluxMeterWrapper within kubernetes types, where deepcopy-gen is used.
func (in *FluxMeterWrapper) DeepCopyInto(out *FluxMeterWrapper) {
	p := proto.Clone(in).(*FluxMeterWrapper)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluxMeterWrapper. Required by controller-gen.
func (in *FluxMeterWrapper) DeepCopy() *FluxMeterWrapper {
	if in == nil {
		return nil
	}
	out := new(FluxMeterWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FluxMeterWrapper. Required by controller-gen.
func (in *FluxMeterWrapper) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ClassifierWrapper within kubernetes types, where deepcopy-gen is used.
func (in *ClassifierWrapper) DeepCopyInto(out *ClassifierWrapper) {
	p := proto.Clone(in).(*ClassifierWrapper)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierWrapper. Required by controller-gen.
func (in *ClassifierWrapper) DeepCopy() *ClassifierWrapper {
	if in == nil {
		return nil
	}
	out := new(ClassifierWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierWrapper. Required by controller-gen.
func (in *ClassifierWrapper) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConcurrencyLimiterWrapper within kubernetes types, where deepcopy-gen is used.
func (in *ConcurrencyLimiterWrapper) DeepCopyInto(out *ConcurrencyLimiterWrapper) {
	p := proto.Clone(in).(*ConcurrencyLimiterWrapper)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConcurrencyLimiterWrapper. Required by controller-gen.
func (in *ConcurrencyLimiterWrapper) DeepCopy() *ConcurrencyLimiterWrapper {
	if in == nil {
		return nil
	}
	out := new(ConcurrencyLimiterWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConcurrencyLimiterWrapper. Required by controller-gen.
func (in *ConcurrencyLimiterWrapper) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimiterWrapper within kubernetes types, where deepcopy-gen is used.
func (in *RateLimiterWrapper) DeepCopyInto(out *RateLimiterWrapper) {
	p := proto.Clone(in).(*RateLimiterWrapper)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiterWrapper. Required by controller-gen.
func (in *RateLimiterWrapper) DeepCopy() *RateLimiterWrapper {
	if in == nil {
		return nil
	}
	out := new(RateLimiterWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiterWrapper. Required by controller-gen.
func (in *RateLimiterWrapper) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
