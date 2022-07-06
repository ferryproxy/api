//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Shiming Zhang.

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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hub) DeepCopyInto(out *Hub) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hub.
func (in *Hub) DeepCopy() *Hub {
	if in == nil {
		return nil
	}
	out := new(Hub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Hub) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubList) DeepCopyInto(out *HubList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Hub, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubList.
func (in *HubList) DeepCopy() *HubList {
	if in == nil {
		return nil
	}
	out := new(HubList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HubList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubSpec) DeepCopyInto(out *HubSpec) {
	*out = *in
	in.Gateway.DeepCopyInto(&out.Gateway)
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = make(map[string]HubSpecGateway, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubSpec.
func (in *HubSpec) DeepCopy() *HubSpec {
	if in == nil {
		return nil
	}
	out := new(HubSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubSpecGateway) DeepCopyInto(out *HubSpecGateway) {
	*out = *in
	if in.Navigation != nil {
		in, out := &in.Navigation, &out.Navigation
		*out = make(HubSpecGatewayWays, len(*in))
		copy(*out, *in)
	}
	if in.Reception != nil {
		in, out := &in.Reception, &out.Reception
		*out = make(HubSpecGatewayWays, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubSpecGateway.
func (in *HubSpecGateway) DeepCopy() *HubSpecGateway {
	if in == nil {
		return nil
	}
	out := new(HubSpecGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubSpecGatewayWay) DeepCopyInto(out *HubSpecGatewayWay) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubSpecGatewayWay.
func (in *HubSpecGatewayWay) DeepCopy() *HubSpecGatewayWay {
	if in == nil {
		return nil
	}
	out := new(HubSpecGatewayWay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in HubSpecGatewayWays) DeepCopyInto(out *HubSpecGatewayWays) {
	{
		in := &in
		*out = make(HubSpecGatewayWays, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubSpecGatewayWays.
func (in HubSpecGatewayWays) DeepCopy() HubSpecGatewayWays {
	if in == nil {
		return nil
	}
	out := new(HubSpecGatewayWays)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubStatus) DeepCopyInto(out *HubStatus) {
	*out = *in
	in.LastSynchronizationTimestamp.DeepCopyInto(&out.LastSynchronizationTimestamp)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubStatus.
func (in *HubStatus) DeepCopy() *HubStatus {
	if in == nil {
		return nil
	}
	out := new(HubStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Route) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteList) DeepCopyInto(out *RouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteList.
func (in *RouteList) DeepCopy() *RouteList {
	if in == nil {
		return nil
	}
	out := new(RouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutePolicy) DeepCopyInto(out *RoutePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutePolicy.
func (in *RoutePolicy) DeepCopy() *RoutePolicy {
	if in == nil {
		return nil
	}
	out := new(RoutePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoutePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutePolicyList) DeepCopyInto(out *RoutePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoutePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutePolicyList.
func (in *RoutePolicyList) DeepCopy() *RoutePolicyList {
	if in == nil {
		return nil
	}
	out := new(RoutePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoutePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutePolicySpec) DeepCopyInto(out *RoutePolicySpec) {
	*out = *in
	if in.Exports != nil {
		in, out := &in.Exports, &out.Exports
		*out = make([]RoutePolicySpecRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Imports != nil {
		in, out := &in.Imports, &out.Imports
		*out = make([]RoutePolicySpecRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutePolicySpec.
func (in *RoutePolicySpec) DeepCopy() *RoutePolicySpec {
	if in == nil {
		return nil
	}
	out := new(RoutePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutePolicySpecRule) DeepCopyInto(out *RoutePolicySpecRule) {
	*out = *in
	in.Service.DeepCopyInto(&out.Service)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutePolicySpecRule.
func (in *RoutePolicySpecRule) DeepCopy() *RoutePolicySpecRule {
	if in == nil {
		return nil
	}
	out := new(RoutePolicySpecRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutePolicySpecRuleService) DeepCopyInto(out *RoutePolicySpecRuleService) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutePolicySpecRuleService.
func (in *RoutePolicySpecRuleService) DeepCopy() *RoutePolicySpecRuleService {
	if in == nil {
		return nil
	}
	out := new(RoutePolicySpecRuleService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutePolicyStatus) DeepCopyInto(out *RoutePolicyStatus) {
	*out = *in
	in.LastSynchronizationTimestamp.DeepCopyInto(&out.LastSynchronizationTimestamp)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutePolicyStatus.
func (in *RoutePolicyStatus) DeepCopy() *RoutePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(RoutePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
	out.Import = in.Import
	out.Export = in.Export
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpecRule) DeepCopyInto(out *RouteSpecRule) {
	*out = *in
	out.Service = in.Service
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpecRule.
func (in *RouteSpecRule) DeepCopy() *RouteSpecRule {
	if in == nil {
		return nil
	}
	out := new(RouteSpecRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpecRuleService) DeepCopyInto(out *RouteSpecRuleService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpecRuleService.
func (in *RouteSpecRuleService) DeepCopy() *RouteSpecRuleService {
	if in == nil {
		return nil
	}
	out := new(RouteSpecRuleService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteStatus) DeepCopyInto(out *RouteStatus) {
	*out = *in
	in.LastSynchronizationTimestamp.DeepCopyInto(&out.LastSynchronizationTimestamp)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteStatus.
func (in *RouteStatus) DeepCopy() *RouteStatus {
	if in == nil {
		return nil
	}
	out := new(RouteStatus)
	in.DeepCopyInto(out)
	return out
}
