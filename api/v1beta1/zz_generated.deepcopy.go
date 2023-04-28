//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Barbican) DeepCopyInto(out *Barbican) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Barbican.
func (in *Barbican) DeepCopy() *Barbican {
	if in == nil {
		return nil
	}
	out := new(Barbican)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Barbican) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanAPI) DeepCopyInto(out *BarbicanAPI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanAPI.
func (in *BarbicanAPI) DeepCopy() *BarbicanAPI {
	if in == nil {
		return nil
	}
	out := new(BarbicanAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BarbicanAPI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanAPIList) DeepCopyInto(out *BarbicanAPIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BarbicanAPI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanAPIList.
func (in *BarbicanAPIList) DeepCopy() *BarbicanAPIList {
	if in == nil {
		return nil
	}
	out := new(BarbicanAPIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BarbicanAPIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanAPISpec) DeepCopyInto(out *BarbicanAPISpec) {
	*out = *in
	out.BarbicanTemplate = in.BarbicanTemplate
	in.BarbicanAPITemplate.DeepCopyInto(&out.BarbicanAPITemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanAPISpec.
func (in *BarbicanAPISpec) DeepCopy() *BarbicanAPISpec {
	if in == nil {
		return nil
	}
	out := new(BarbicanAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanAPIStatus) DeepCopyInto(out *BarbicanAPIStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanAPIStatus.
func (in *BarbicanAPIStatus) DeepCopy() *BarbicanAPIStatus {
	if in == nil {
		return nil
	}
	out := new(BarbicanAPIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanAPITemplate) DeepCopyInto(out *BarbicanAPITemplate) {
	*out = *in
	in.BarbicanComponentTemplate.DeepCopyInto(&out.BarbicanComponentTemplate)
	if in.ExternalEndpoints != nil {
		in, out := &in.ExternalEndpoints, &out.ExternalEndpoints
		*out = make([]MetalLBConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanAPITemplate.
func (in *BarbicanAPITemplate) DeepCopy() *BarbicanAPITemplate {
	if in == nil {
		return nil
	}
	out := new(BarbicanAPITemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanComponentTemplate) DeepCopyInto(out *BarbicanComponentTemplate) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanComponentTemplate.
func (in *BarbicanComponentTemplate) DeepCopy() *BarbicanComponentTemplate {
	if in == nil {
		return nil
	}
	out := new(BarbicanComponentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanDebug) DeepCopyInto(out *BarbicanDebug) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanDebug.
func (in *BarbicanDebug) DeepCopy() *BarbicanDebug {
	if in == nil {
		return nil
	}
	out := new(BarbicanDebug)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanList) DeepCopyInto(out *BarbicanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Barbican, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanList.
func (in *BarbicanList) DeepCopy() *BarbicanList {
	if in == nil {
		return nil
	}
	out := new(BarbicanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BarbicanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanSpec) DeepCopyInto(out *BarbicanSpec) {
	*out = *in
	out.BarbicanTemplate = in.BarbicanTemplate
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.BarbicanAPI.DeepCopyInto(&out.BarbicanAPI)
	in.BarbicanWorker.DeepCopyInto(&out.BarbicanWorker)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanSpec.
func (in *BarbicanSpec) DeepCopy() *BarbicanSpec {
	if in == nil {
		return nil
	}
	out := new(BarbicanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanStatus) DeepCopyInto(out *BarbicanStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.ServiceIDs != nil {
		in, out := &in.ServiceIDs, &out.ServiceIDs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanStatus.
func (in *BarbicanStatus) DeepCopy() *BarbicanStatus {
	if in == nil {
		return nil
	}
	out := new(BarbicanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanTemplate) DeepCopyInto(out *BarbicanTemplate) {
	*out = *in
	out.PasswordSelectors = in.PasswordSelectors
	out.Debug = in.Debug
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanTemplate.
func (in *BarbicanTemplate) DeepCopy() *BarbicanTemplate {
	if in == nil {
		return nil
	}
	out := new(BarbicanTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanWorker) DeepCopyInto(out *BarbicanWorker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanWorker.
func (in *BarbicanWorker) DeepCopy() *BarbicanWorker {
	if in == nil {
		return nil
	}
	out := new(BarbicanWorker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BarbicanWorker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanWorkerList) DeepCopyInto(out *BarbicanWorkerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BarbicanWorker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanWorkerList.
func (in *BarbicanWorkerList) DeepCopy() *BarbicanWorkerList {
	if in == nil {
		return nil
	}
	out := new(BarbicanWorkerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BarbicanWorkerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanWorkerSpec) DeepCopyInto(out *BarbicanWorkerSpec) {
	*out = *in
	out.BarbicanTemplate = in.BarbicanTemplate
	in.BarbicanWorkerTemplate.DeepCopyInto(&out.BarbicanWorkerTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanWorkerSpec.
func (in *BarbicanWorkerSpec) DeepCopy() *BarbicanWorkerSpec {
	if in == nil {
		return nil
	}
	out := new(BarbicanWorkerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanWorkerStatus) DeepCopyInto(out *BarbicanWorkerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanWorkerStatus.
func (in *BarbicanWorkerStatus) DeepCopy() *BarbicanWorkerStatus {
	if in == nil {
		return nil
	}
	out := new(BarbicanWorkerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanWorkerTemplate) DeepCopyInto(out *BarbicanWorkerTemplate) {
	*out = *in
	in.BarbicanComponentTemplate.DeepCopyInto(&out.BarbicanComponentTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanWorkerTemplate.
func (in *BarbicanWorkerTemplate) DeepCopy() *BarbicanWorkerTemplate {
	if in == nil {
		return nil
	}
	out := new(BarbicanWorkerTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetalLBConfig) DeepCopyInto(out *MetalLBConfig) {
	*out = *in
	if in.LoadBalancerIPs != nil {
		in, out := &in.LoadBalancerIPs, &out.LoadBalancerIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetalLBConfig.
func (in *MetalLBConfig) DeepCopy() *MetalLBConfig {
	if in == nil {
		return nil
	}
	out := new(MetalLBConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordSelector) DeepCopyInto(out *PasswordSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordSelector.
func (in *PasswordSelector) DeepCopy() *PasswordSelector {
	if in == nil {
		return nil
	}
	out := new(PasswordSelector)
	in.DeepCopyInto(out)
	return out
}
