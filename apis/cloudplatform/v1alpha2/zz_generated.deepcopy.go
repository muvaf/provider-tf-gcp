//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

package v1alpha2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccount) DeepCopyInto(out *ServiceAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccount.
func (in *ServiceAccount) DeepCopy() *ServiceAccount {
	if in == nil {
		return nil
	}
	out := new(ServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountKey) DeepCopyInto(out *ServiceAccountKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountKey.
func (in *ServiceAccountKey) DeepCopy() *ServiceAccountKey {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceAccountKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountKeyList) DeepCopyInto(out *ServiceAccountKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceAccountKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountKeyList.
func (in *ServiceAccountKeyList) DeepCopy() *ServiceAccountKeyList {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceAccountKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountKeyObservation) DeepCopyInto(out *ServiceAccountKeyObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(string)
		**out = **in
	}
	if in.ValidAfter != nil {
		in, out := &in.ValidAfter, &out.ValidAfter
		*out = new(string)
		**out = **in
	}
	if in.ValidBefore != nil {
		in, out := &in.ValidBefore, &out.ValidBefore
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountKeyObservation.
func (in *ServiceAccountKeyObservation) DeepCopy() *ServiceAccountKeyObservation {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountKeyParameters) DeepCopyInto(out *ServiceAccountKeyParameters) {
	*out = *in
	if in.Keepers != nil {
		in, out := &in.Keepers, &out.Keepers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KeyAlgorithm != nil {
		in, out := &in.KeyAlgorithm, &out.KeyAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.PrivateKeyType != nil {
		in, out := &in.PrivateKeyType, &out.PrivateKeyType
		*out = new(string)
		**out = **in
	}
	if in.PublicKeyData != nil {
		in, out := &in.PublicKeyData, &out.PublicKeyData
		*out = new(string)
		**out = **in
	}
	if in.PublicKeyType != nil {
		in, out := &in.PublicKeyType, &out.PublicKeyType
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountIDRef != nil {
		in, out := &in.ServiceAccountIDRef, &out.ServiceAccountIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ServiceAccountIDSelector != nil {
		in, out := &in.ServiceAccountIDSelector, &out.ServiceAccountIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountKeyParameters.
func (in *ServiceAccountKeyParameters) DeepCopy() *ServiceAccountKeyParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountKeySpec) DeepCopyInto(out *ServiceAccountKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountKeySpec.
func (in *ServiceAccountKeySpec) DeepCopy() *ServiceAccountKeySpec {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountKeyStatus) DeepCopyInto(out *ServiceAccountKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountKeyStatus.
func (in *ServiceAccountKeyStatus) DeepCopy() *ServiceAccountKeyStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountList) DeepCopyInto(out *ServiceAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountList.
func (in *ServiceAccountList) DeepCopy() *ServiceAccountList {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountObservation) DeepCopyInto(out *ServiceAccountObservation) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.UniqueID != nil {
		in, out := &in.UniqueID, &out.UniqueID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountObservation.
func (in *ServiceAccountObservation) DeepCopy() *ServiceAccountObservation {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountParameters) DeepCopyInto(out *ServiceAccountParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountParameters.
func (in *ServiceAccountParameters) DeepCopy() *ServiceAccountParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountSpec) DeepCopyInto(out *ServiceAccountSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountSpec.
func (in *ServiceAccountSpec) DeepCopy() *ServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountStatus) DeepCopyInto(out *ServiceAccountStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountStatus.
func (in *ServiceAccountStatus) DeepCopy() *ServiceAccountStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountStatus)
	in.DeepCopyInto(out)
	return out
}
