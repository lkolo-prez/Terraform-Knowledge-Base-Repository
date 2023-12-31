package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type NetworkSecurityGroupId struct {
	SubscriptionId string
	ResourceGroup  string
	Name           string
}

func NewNetworkSecurityGroupID(subscriptionId, resourceGroup, name string) NetworkSecurityGroupId {
	return NetworkSecurityGroupId{
		SubscriptionId: subscriptionId,
		ResourceGroup:  resourceGroup,
		Name:           name,
	}
}

func (id NetworkSecurityGroupId) String() string {
	segments := []string{
		fmt.Sprintf("Name %q", id.Name),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Network Security Group", segmentsStr)
}

func (id NetworkSecurityGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/networkSecurityGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.Name)
}

// NetworkSecurityGroupID parses a NetworkSecurityGroup ID into an NetworkSecurityGroupId struct
func NetworkSecurityGroupID(input string) (*NetworkSecurityGroupId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("parsing %q as an NetworkSecurityGroup ID: %+v", input, err)
	}

	resourceId := NetworkSecurityGroupId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	if resourceId.Name, err = id.PopSegment("networkSecurityGroups"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}

// NetworkSecurityGroupIDInsensitively parses an NetworkSecurityGroup ID into an NetworkSecurityGroupId struct, insensitively
// This should only be used to parse an ID for rewriting, the NetworkSecurityGroupID
// method should be used instead for validation etc.
//
// Whilst this may seem strange, this enables Terraform have consistent casing
// which works around issues in Core, whilst handling broken API responses.
func NetworkSecurityGroupIDInsensitively(input string) (*NetworkSecurityGroupId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := NetworkSecurityGroupId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	// find the correct casing for the 'networkSecurityGroups' segment
	networkSecurityGroupsKey := "networkSecurityGroups"
	for key := range id.Path {
		if strings.EqualFold(key, networkSecurityGroupsKey) {
			networkSecurityGroupsKey = key
			break
		}
	}
	if resourceId.Name, err = id.PopSegment(networkSecurityGroupsKey); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
