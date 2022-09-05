// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package statusv1

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using GroupStatusRequest within kubernetes types, where deepcopy-gen is used.
func (in *GroupStatusRequest) DeepCopyInto(out *GroupStatusRequest) {
	p := proto.Clone(in).(*GroupStatusRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupStatusRequest. Required by controller-gen.
func (in *GroupStatusRequest) DeepCopy() *GroupStatusRequest {
	if in == nil {
		return nil
	}
	out := new(GroupStatusRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GroupStatusRequest. Required by controller-gen.
func (in *GroupStatusRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GroupStatus within kubernetes types, where deepcopy-gen is used.
func (in *GroupStatus) DeepCopyInto(out *GroupStatus) {
	p := proto.Clone(in).(*GroupStatus)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupStatus. Required by controller-gen.
func (in *GroupStatus) DeepCopy() *GroupStatus {
	if in == nil {
		return nil
	}
	out := new(GroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GroupStatus. Required by controller-gen.
func (in *GroupStatus) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Status within kubernetes types, where deepcopy-gen is used.
func (in *Status) DeepCopyInto(out *Status) {
	p := proto.Clone(in).(*Status)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status. Required by controller-gen.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Status. Required by controller-gen.
func (in *Status) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ErrorDetails within kubernetes types, where deepcopy-gen is used.
func (in *ErrorDetails) DeepCopyInto(out *ErrorDetails) {
	p := proto.Clone(in).(*ErrorDetails)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorDetails. Required by controller-gen.
func (in *ErrorDetails) DeepCopy() *ErrorDetails {
	if in == nil {
		return nil
	}
	out := new(ErrorDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ErrorDetails. Required by controller-gen.
func (in *ErrorDetails) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}