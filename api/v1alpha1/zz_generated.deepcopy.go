//go:build !ignore_autogenerated

// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerPolicy) DeepCopyInto(out *ContainerPolicy) {
	*out = *in
	if in.PercentageIncrease != nil {
		in, out := &in.PercentageIncrease, &out.PercentageIncrease
		*out = new(PercentageIncrease)
		**out = **in
	}
	if in.FixedResources != nil {
		in, out := &in.FixedResources, &out.FixedResources
		*out = new(FixedResources)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerPolicy.
func (in *ContainerPolicy) DeepCopy() *ContainerPolicy {
	if in == nil {
		return nil
	}
	out := new(ContainerPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DurationPolicy) DeepCopyInto(out *DurationPolicy) {
	*out = *in
	if in.Fixed != nil {
		in, out := &in.Fixed, &out.Fixed
		*out = new(FixedDurationPolicy)
		**out = **in
	}
	if in.PodCondition != nil {
		in, out := &in.PodCondition, &out.PodCondition
		*out = new(PodConditionDurationPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DurationPolicy.
func (in *DurationPolicy) DeepCopy() *DurationPolicy {
	if in == nil {
		return nil
	}
	out := new(DurationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixedDurationPolicy) DeepCopyInto(out *FixedDurationPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixedDurationPolicy.
func (in *FixedDurationPolicy) DeepCopy() *FixedDurationPolicy {
	if in == nil {
		return nil
	}
	out := new(FixedDurationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixedResources) DeepCopyInto(out *FixedResources) {
	*out = *in
	out.Requests = in.Requests.DeepCopy()
	out.Limits = in.Limits.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixedResources.
func (in *FixedResources) DeepCopy() *FixedResources {
	if in == nil {
		return nil
	}
	out := new(FixedResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PercentageIncrease) DeepCopyInto(out *PercentageIncrease) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PercentageIncrease.
func (in *PercentageIncrease) DeepCopy() *PercentageIncrease {
	if in == nil {
		return nil
	}
	out := new(PercentageIncrease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodConditionDurationPolicy) DeepCopyInto(out *PodConditionDurationPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodConditionDurationPolicy.
func (in *PodConditionDurationPolicy) DeepCopy() *PodConditionDurationPolicy {
	if in == nil {
		return nil
	}
	out := new(PodConditionDurationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicy) DeepCopyInto(out *ResourcePolicy) {
	*out = *in
	if in.ContainerPolicies != nil {
		in, out := &in.ContainerPolicies, &out.ContainerPolicies
		*out = make([]ContainerPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicy.
func (in *ResourcePolicy) DeepCopy() *ResourcePolicy {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartupCPUBoost) DeepCopyInto(out *StartupCPUBoost) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Selector.DeepCopyInto(&out.Selector)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartupCPUBoost.
func (in *StartupCPUBoost) DeepCopy() *StartupCPUBoost {
	if in == nil {
		return nil
	}
	out := new(StartupCPUBoost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StartupCPUBoost) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartupCPUBoostList) DeepCopyInto(out *StartupCPUBoostList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StartupCPUBoost, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartupCPUBoostList.
func (in *StartupCPUBoostList) DeepCopy() *StartupCPUBoostList {
	if in == nil {
		return nil
	}
	out := new(StartupCPUBoostList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StartupCPUBoostList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartupCPUBoostSpec) DeepCopyInto(out *StartupCPUBoostSpec) {
	*out = *in
	in.ResourcePolicy.DeepCopyInto(&out.ResourcePolicy)
	in.DurationPolicy.DeepCopyInto(&out.DurationPolicy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartupCPUBoostSpec.
func (in *StartupCPUBoostSpec) DeepCopy() *StartupCPUBoostSpec {
	if in == nil {
		return nil
	}
	out := new(StartupCPUBoostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartupCPUBoostStatus) DeepCopyInto(out *StartupCPUBoostStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartupCPUBoostStatus.
func (in *StartupCPUBoostStatus) DeepCopy() *StartupCPUBoostStatus {
	if in == nil {
		return nil
	}
	out := new(StartupCPUBoostStatus)
	in.DeepCopyInto(out)
	return out
}
