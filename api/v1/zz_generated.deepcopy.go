// +build !ignore_autogenerated

/*
Copyright 2021 PolicyEngineEndpoint Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Arguments) DeepCopyInto(out *Arguments) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Arguments.
func (in *Arguments) DeepCopy() *Arguments {
	if in == nil {
		return nil
	}
	out := new(Arguments)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *General) DeepCopyInto(out *General) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new General.
func (in *General) DeepCopy() *General {
	if in == nil {
		return nil
	}
	out := new(General)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module) DeepCopyInto(out *Module) {
	*out = *in
	if in.Consumers != nil {
		in, out := &in.Consumers, &out.Consumers
		*out = make([]*PolicyEngineConsumer, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PolicyEngineConsumer)
				**out = **in
			}
		}
	}
	if in.General != nil {
		in, out := &in.General, &out.General
		*out = new(General)
		**out = **in
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(Arguments)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module.
func (in *Module) DeepCopy() *Module {
	if in == nil {
		return nil
	}
	out := new(Module)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyEngineConsumer) DeepCopyInto(out *PolicyEngineConsumer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyEngineConsumer.
func (in *PolicyEngineConsumer) DeepCopy() *PolicyEngineConsumer {
	if in == nil {
		return nil
	}
	out := new(PolicyEngineConsumer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyEngineEndpoint) DeepCopyInto(out *PolicyEngineEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyEngineEndpoint.
func (in *PolicyEngineEndpoint) DeepCopy() *PolicyEngineEndpoint {
	if in == nil {
		return nil
	}
	out := new(PolicyEngineEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyEngineEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyEngineEndpointList) DeepCopyInto(out *PolicyEngineEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyEngineEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyEngineEndpointList.
func (in *PolicyEngineEndpointList) DeepCopy() *PolicyEngineEndpointList {
	if in == nil {
		return nil
	}
	out := new(PolicyEngineEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyEngineEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyEngineEndpointSpec) DeepCopyInto(out *PolicyEngineEndpointSpec) {
	*out = *in
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]*Module, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Module)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyEngineEndpointSpec.
func (in *PolicyEngineEndpointSpec) DeepCopy() *PolicyEngineEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyEngineEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyEngineEndpointStatus) DeepCopyInto(out *PolicyEngineEndpointStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyEngineEndpointStatus.
func (in *PolicyEngineEndpointStatus) DeepCopy() *PolicyEngineEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyEngineEndpointStatus)
	in.DeepCopyInto(out)
	return out
}
