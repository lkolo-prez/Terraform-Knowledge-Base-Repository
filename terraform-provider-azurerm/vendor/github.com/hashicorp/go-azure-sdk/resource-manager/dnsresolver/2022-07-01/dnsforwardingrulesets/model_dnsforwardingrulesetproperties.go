package dnsforwardingrulesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DnsForwardingRulesetProperties struct {
	DnsResolverOutboundEndpoints []SubResource      `json:"dnsResolverOutboundEndpoints"`
	ProvisioningState            *ProvisioningState `json:"provisioningState,omitempty"`
	ResourceGuid                 *string            `json:"resourceGuid,omitempty"`
}
