//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NSXT) DeepCopyInto(out *NSXT) {
	*out = *in
	if in.PodRoutingEnabled != nil {
		in, out := &in.PodRoutingEnabled, &out.PodRoutingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(NSXTRouteConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CredentialLocalObjRef != nil {
		in, out := &in.CredentialLocalObjRef, &out.CredentialLocalObjRef
		*out = new(v1.TypedLocalObjectReference)
		(*in).DeepCopyInto(*out)
	}
	if in.APIHost != nil {
		in, out := &in.APIHost, &out.APIHost
		*out = new(string)
		**out = **in
	}
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(bool)
		**out = **in
	}
	if in.RemoteAuth != nil {
		in, out := &in.RemoteAuth, &out.RemoteAuth
		*out = new(bool)
		**out = **in
	}
	if in.VMCAccessToken != nil {
		in, out := &in.VMCAccessToken, &out.VMCAccessToken
		*out = new(string)
		**out = **in
	}
	if in.VMCAuthHost != nil {
		in, out := &in.VMCAuthHost, &out.VMCAuthHost
		*out = new(string)
		**out = **in
	}
	if in.ClientCertKeyData != nil {
		in, out := &in.ClientCertKeyData, &out.ClientCertKeyData
		*out = new(string)
		**out = **in
	}
	if in.ClientCertData != nil {
		in, out := &in.ClientCertData, &out.ClientCertData
		*out = new(string)
		**out = **in
	}
	if in.RootCAData != nil {
		in, out := &in.RootCAData, &out.RootCAData
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NSXT.
func (in *NSXT) DeepCopy() *NSXT {
	if in == nil {
		return nil
	}
	out := new(NSXT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NSXTRouteConfig) DeepCopyInto(out *NSXTRouteConfig) {
	*out = *in
	if in.RouterPath != nil {
		in, out := &in.RouterPath, &out.RouterPath
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NSXTRouteConfig.
func (in *NSXTRouteConfig) DeepCopy() *NSXTRouteConfig {
	if in == nil {
		return nil
	}
	out := new(NSXTRouteConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonParavirtualConfig) DeepCopyInto(out *NonParavirtualConfig) {
	*out = *in
	if in.VCenterAPIEndpoint != nil {
		in, out := &in.VCenterAPIEndpoint, &out.VCenterAPIEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.VSphereCredentialLocalObjRef != nil {
		in, out := &in.VSphereCredentialLocalObjRef, &out.VSphereCredentialLocalObjRef
		*out = new(v1.TypedLocalObjectReference)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSThumbprint != nil {
		in, out := &in.TLSThumbprint, &out.TLSThumbprint
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(bool)
		**out = **in
	}
	if in.IPFamily != nil {
		in, out := &in.IPFamily, &out.IPFamily
		*out = new(string)
		**out = **in
	}
	if in.VMNetwork != nil {
		in, out := &in.VMNetwork, &out.VMNetwork
		*out = new(VMNetwork)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSCipherSuites != nil {
		in, out := &in.TLSCipherSuites, &out.TLSCipherSuites
		*out = new(string)
		**out = **in
	}
	if in.NSXT != nil {
		in, out := &in.NSXT, &out.NSXT
		*out = new(NSXT)
		(*in).DeepCopyInto(*out)
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(Proxy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonParavirtualConfig.
func (in *NonParavirtualConfig) DeepCopy() *NonParavirtualConfig {
	if in == nil {
		return nil
	}
	out := new(NonParavirtualConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParavirtualConfig) DeepCopyInto(out *ParavirtualConfig) {
	*out = *in
	if in.AntreaNSXPodRoutingEnabled != nil {
		in, out := &in.AntreaNSXPodRoutingEnabled, &out.AntreaNSXPodRoutingEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParavirtualConfig.
func (in *ParavirtualConfig) DeepCopy() *ParavirtualConfig {
	if in == nil {
		return nil
	}
	out := new(ParavirtualConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Proxy) DeepCopyInto(out *Proxy) {
	*out = *in
	if in.HTTPProxy != nil {
		in, out := &in.HTTPProxy, &out.HTTPProxy
		*out = new(string)
		**out = **in
	}
	if in.HTTPSProxy != nil {
		in, out := &in.HTTPSProxy, &out.HTTPSProxy
		*out = new(string)
		**out = **in
	}
	if in.NoProxy != nil {
		in, out := &in.NoProxy, &out.NoProxy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Proxy.
func (in *Proxy) DeepCopy() *Proxy {
	if in == nil {
		return nil
	}
	out := new(Proxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMNetwork) DeepCopyInto(out *VMNetwork) {
	*out = *in
	if in.Internal != nil {
		in, out := &in.Internal, &out.Internal
		*out = new(string)
		**out = **in
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(string)
		**out = **in
	}
	if in.ExcludeInternalSubnetCidr != nil {
		in, out := &in.ExcludeInternalSubnetCidr, &out.ExcludeInternalSubnetCidr
		*out = new(string)
		**out = **in
	}
	if in.ExcludeExternalSubnetCidr != nil {
		in, out := &in.ExcludeExternalSubnetCidr, &out.ExcludeExternalSubnetCidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMNetwork.
func (in *VMNetwork) DeepCopy() *VMNetwork {
	if in == nil {
		return nil
	}
	out := new(VMNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereCPI) DeepCopyInto(out *VSphereCPI) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.NonParavirtualConfig != nil {
		in, out := &in.NonParavirtualConfig, &out.NonParavirtualConfig
		*out = new(NonParavirtualConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ParavirtualConfig != nil {
		in, out := &in.ParavirtualConfig, &out.ParavirtualConfig
		*out = new(ParavirtualConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereCPI.
func (in *VSphereCPI) DeepCopy() *VSphereCPI {
	if in == nil {
		return nil
	}
	out := new(VSphereCPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereCPIConfig) DeepCopyInto(out *VSphereCPIConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereCPIConfig.
func (in *VSphereCPIConfig) DeepCopy() *VSphereCPIConfig {
	if in == nil {
		return nil
	}
	out := new(VSphereCPIConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereCPIConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereCPIConfigList) DeepCopyInto(out *VSphereCPIConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VSphereCPIConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereCPIConfigList.
func (in *VSphereCPIConfigList) DeepCopy() *VSphereCPIConfigList {
	if in == nil {
		return nil
	}
	out := new(VSphereCPIConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereCPIConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereCPIConfigSpec) DeepCopyInto(out *VSphereCPIConfigSpec) {
	*out = *in
	in.VSphereCPI.DeepCopyInto(&out.VSphereCPI)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereCPIConfigSpec.
func (in *VSphereCPIConfigSpec) DeepCopy() *VSphereCPIConfigSpec {
	if in == nil {
		return nil
	}
	out := new(VSphereCPIConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereCPIConfigStatus) DeepCopyInto(out *VSphereCPIConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereCPIConfigStatus.
func (in *VSphereCPIConfigStatus) DeepCopy() *VSphereCPIConfigStatus {
	if in == nil {
		return nil
	}
	out := new(VSphereCPIConfigStatus)
	in.DeepCopyInto(out)
	return out
}