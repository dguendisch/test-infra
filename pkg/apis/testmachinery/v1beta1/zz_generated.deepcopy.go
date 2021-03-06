// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigElement) DeepCopyInto(out *ConfigElement) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(ConfigSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigElement.
func (in *ConfigElement) DeepCopy() *ConfigElement {
	if in == nil {
		return nil
	}
	out := new(ConfigElement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSource) DeepCopyInto(out *ConfigSource) {
	*out = *in
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSource.
func (in *ConfigSource) DeepCopy() *ConfigSource {
	if in == nil {
		return nil
	}
	out := new(ConfigSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestDefMetadata) DeepCopyInto(out *TestDefMetadata) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestDefMetadata.
func (in *TestDefMetadata) DeepCopy() *TestDefMetadata {
	if in == nil {
		return nil
	}
	out := new(TestDefMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestDefSpec) DeepCopyInto(out *TestDefSpec) {
	*out = *in
	if in.RecipientsOnFailure != nil {
		in, out := &in.RecipientsOnFailure, &out.RecipientsOnFailure
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Behavior != nil {
		in, out := &in.Behavior, &out.Behavior
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigElement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestDefSpec.
func (in *TestDefSpec) DeepCopy() *TestDefSpec {
	if in == nil {
		return nil
	}
	out := new(TestDefSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestDefinition) DeepCopyInto(out *TestDefinition) {
	*out = *in
	out.Metadata = in.Metadata
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestDefinition.
func (in *TestDefinition) DeepCopy() *TestDefinition {
	if in == nil {
		return nil
	}
	out := new(TestDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in TestFlow) DeepCopyInto(out *TestFlow) {
	{
		in := &in
		*out = make(TestFlow, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]TestflowStep, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestFlow.
func (in TestFlow) DeepCopy() TestFlow {
	if in == nil {
		return nil
	}
	out := new(TestFlow)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestLocation) DeepCopyInto(out *TestLocation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestLocation.
func (in *TestLocation) DeepCopy() *TestLocation {
	if in == nil {
		return nil
	}
	out := new(TestLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestflowStep) DeepCopyInto(out *TestflowStep) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigElement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestflowStep.
func (in *TestflowStep) DeepCopy() *TestflowStep {
	if in == nil {
		return nil
	}
	out := new(TestflowStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestflowStepStatus) DeepCopyInto(out *TestflowStepStatus) {
	*out = *in
	in.TestDefinition.DeepCopyInto(&out.TestDefinition)
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestflowStepStatus.
func (in *TestflowStepStatus) DeepCopy() *TestflowStepStatus {
	if in == nil {
		return nil
	}
	out := new(TestflowStepStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestflowStepStatusTestDefinition) DeepCopyInto(out *TestflowStepStatusTestDefinition) {
	*out = *in
	out.Location = in.Location
	if in.RecipientsOnFailure != nil {
		in, out := &in.RecipientsOnFailure, &out.RecipientsOnFailure
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Position != nil {
		in, out := &in.Position, &out.Position
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestflowStepStatusTestDefinition.
func (in *TestflowStepStatusTestDefinition) DeepCopy() *TestflowStepStatusTestDefinition {
	if in == nil {
		return nil
	}
	out := new(TestflowStepStatusTestDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Testrun) DeepCopyInto(out *Testrun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Testrun.
func (in *Testrun) DeepCopy() *Testrun {
	if in == nil {
		return nil
	}
	out := new(Testrun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Testrun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestrunKubeconfigs) DeepCopyInto(out *TestrunKubeconfigs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestrunKubeconfigs.
func (in *TestrunKubeconfigs) DeepCopy() *TestrunKubeconfigs {
	if in == nil {
		return nil
	}
	out := new(TestrunKubeconfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestrunList) DeepCopyInto(out *TestrunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Testrun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestrunList.
func (in *TestrunList) DeepCopy() *TestrunList {
	if in == nil {
		return nil
	}
	out := new(TestrunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestrunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestrunSpec) DeepCopyInto(out *TestrunSpec) {
	*out = *in
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	if in.TestLocations != nil {
		in, out := &in.TestLocations, &out.TestLocations
		*out = make([]TestLocation, len(*in))
		copy(*out, *in)
	}
	out.Kubeconfigs = in.Kubeconfigs
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigElement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TestFlow != nil {
		in, out := &in.TestFlow, &out.TestFlow
		*out = make(TestFlow, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]TestflowStep, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
		}
	}
	if in.OnExit != nil {
		in, out := &in.OnExit, &out.OnExit
		*out = make(TestFlow, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]TestflowStep, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestrunSpec.
func (in *TestrunSpec) DeepCopy() *TestrunSpec {
	if in == nil {
		return nil
	}
	out := new(TestrunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestrunStatus) DeepCopyInto(out *TestrunStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([][]*TestflowStepStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]*TestflowStepStatus, len(*in))
				for i := range *in {
					if (*in)[i] != nil {
						in, out := &(*in)[i], &(*out)[i]
						*out = new(TestflowStepStatus)
						(*in).DeepCopyInto(*out)
					}
				}
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestrunStatus.
func (in *TestrunStatus) DeepCopy() *TestrunStatus {
	if in == nil {
		return nil
	}
	out := new(TestrunStatus)
	in.DeepCopyInto(out)
	return out
}
