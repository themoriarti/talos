// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type DefaultActionConfigV1Alpha1 -type KubespanEndpointsConfigV1Alpha1 -type RuleConfigV1Alpha1 -pointer-receiver -header-file ../../../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package network

import (
	"net/netip"
)

// DeepCopy generates a deep copy of *DefaultActionConfigV1Alpha1.
func (o *DefaultActionConfigV1Alpha1) DeepCopy() *DefaultActionConfigV1Alpha1 {
	var cp DefaultActionConfigV1Alpha1 = *o
	return &cp
}

// DeepCopy generates a deep copy of *KubespanEndpointsConfigV1Alpha1.
func (o *KubespanEndpointsConfigV1Alpha1) DeepCopy() *KubespanEndpointsConfigV1Alpha1 {
	var cp KubespanEndpointsConfigV1Alpha1 = *o
	if o.ExtraAnnouncedEndpointsConfig != nil {
		cp.ExtraAnnouncedEndpointsConfig = make([]netip.AddrPort, len(o.ExtraAnnouncedEndpointsConfig))
		copy(cp.ExtraAnnouncedEndpointsConfig, o.ExtraAnnouncedEndpointsConfig)
	}
	return &cp
}

// DeepCopy generates a deep copy of *RuleConfigV1Alpha1.
func (o *RuleConfigV1Alpha1) DeepCopy() *RuleConfigV1Alpha1 {
	var cp RuleConfigV1Alpha1 = *o
	if o.PortSelector.Ports != nil {
		cp.PortSelector.Ports = make([]PortRange, len(o.PortSelector.Ports))
		copy(cp.PortSelector.Ports, o.PortSelector.Ports)
	}
	if o.Ingress != nil {
		cp.Ingress = make([]IngressRule, len(o.Ingress))
		copy(cp.Ingress, o.Ingress)
	}
	return &cp
}
