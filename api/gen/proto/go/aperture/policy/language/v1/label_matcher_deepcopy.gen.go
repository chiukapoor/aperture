// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package languagev1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using LabelMatcher within kubernetes types, where deepcopy-gen is used.
func (in *LabelMatcher) DeepCopyInto(out *LabelMatcher) {
	p := proto.Clone(in).(*LabelMatcher)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelMatcher. Required by controller-gen.
func (in *LabelMatcher) DeepCopy() *LabelMatcher {
	if in == nil {
		return nil
	}
	out := new(LabelMatcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LabelMatcher. Required by controller-gen.
func (in *LabelMatcher) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using K8SLabelMatcherRequirement within kubernetes types, where deepcopy-gen is used.
func (in *K8SLabelMatcherRequirement) DeepCopyInto(out *K8SLabelMatcherRequirement) {
	p := proto.Clone(in).(*K8SLabelMatcherRequirement)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8SLabelMatcherRequirement. Required by controller-gen.
func (in *K8SLabelMatcherRequirement) DeepCopy() *K8SLabelMatcherRequirement {
	if in == nil {
		return nil
	}
	out := new(K8SLabelMatcherRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new K8SLabelMatcherRequirement. Required by controller-gen.
func (in *K8SLabelMatcherRequirement) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using MatchExpression within kubernetes types, where deepcopy-gen is used.
func (in *MatchExpression) DeepCopyInto(out *MatchExpression) {
	p := proto.Clone(in).(*MatchExpression)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchExpression. Required by controller-gen.
func (in *MatchExpression) DeepCopy() *MatchExpression {
	if in == nil {
		return nil
	}
	out := new(MatchExpression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MatchExpression. Required by controller-gen.
func (in *MatchExpression) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using MatchExpression_List within kubernetes types, where deepcopy-gen is used.
func (in *MatchExpression_List) DeepCopyInto(out *MatchExpression_List) {
	p := proto.Clone(in).(*MatchExpression_List)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchExpression_List. Required by controller-gen.
func (in *MatchExpression_List) DeepCopy() *MatchExpression_List {
	if in == nil {
		return nil
	}
	out := new(MatchExpression_List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MatchExpression_List. Required by controller-gen.
func (in *MatchExpression_List) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using EqualsMatchExpression within kubernetes types, where deepcopy-gen is used.
func (in *EqualsMatchExpression) DeepCopyInto(out *EqualsMatchExpression) {
	p := proto.Clone(in).(*EqualsMatchExpression)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EqualsMatchExpression. Required by controller-gen.
func (in *EqualsMatchExpression) DeepCopy() *EqualsMatchExpression {
	if in == nil {
		return nil
	}
	out := new(EqualsMatchExpression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new EqualsMatchExpression. Required by controller-gen.
func (in *EqualsMatchExpression) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using MatchesMatchExpression within kubernetes types, where deepcopy-gen is used.
func (in *MatchesMatchExpression) DeepCopyInto(out *MatchesMatchExpression) {
	p := proto.Clone(in).(*MatchesMatchExpression)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchesMatchExpression. Required by controller-gen.
func (in *MatchesMatchExpression) DeepCopy() *MatchesMatchExpression {
	if in == nil {
		return nil
	}
	out := new(MatchesMatchExpression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MatchesMatchExpression. Required by controller-gen.
func (in *MatchesMatchExpression) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
