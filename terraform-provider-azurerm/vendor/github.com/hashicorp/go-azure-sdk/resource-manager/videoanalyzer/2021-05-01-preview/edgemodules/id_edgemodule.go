package edgemodules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = EdgeModuleId{}

// EdgeModuleId is a struct representing the Resource ID for a Edge Module
type EdgeModuleId struct {
	SubscriptionId    string
	ResourceGroupName string
	VideoAnalyzerName string
	EdgeModuleName    string
}

// NewEdgeModuleID returns a new EdgeModuleId struct
func NewEdgeModuleID(subscriptionId string, resourceGroupName string, videoAnalyzerName string, edgeModuleName string) EdgeModuleId {
	return EdgeModuleId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		VideoAnalyzerName: videoAnalyzerName,
		EdgeModuleName:    edgeModuleName,
	}
}

// ParseEdgeModuleID parses 'input' into a EdgeModuleId
func ParseEdgeModuleID(input string) (*EdgeModuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(EdgeModuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EdgeModuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VideoAnalyzerName, ok = parsed.Parsed["videoAnalyzerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "videoAnalyzerName", *parsed)
	}

	if id.EdgeModuleName, ok = parsed.Parsed["edgeModuleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "edgeModuleName", *parsed)
	}

	return &id, nil
}

// ParseEdgeModuleIDInsensitively parses 'input' case-insensitively into a EdgeModuleId
// note: this method should only be used for API response data and not user input
func ParseEdgeModuleIDInsensitively(input string) (*EdgeModuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(EdgeModuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EdgeModuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VideoAnalyzerName, ok = parsed.Parsed["videoAnalyzerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "videoAnalyzerName", *parsed)
	}

	if id.EdgeModuleName, ok = parsed.Parsed["edgeModuleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "edgeModuleName", *parsed)
	}

	return &id, nil
}

// ValidateEdgeModuleID checks that 'input' can be parsed as a Edge Module ID
func ValidateEdgeModuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEdgeModuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Edge Module ID
func (id EdgeModuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Media/videoAnalyzers/%s/edgeModules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VideoAnalyzerName, id.EdgeModuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Edge Module ID
func (id EdgeModuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMedia", "Microsoft.Media", "Microsoft.Media"),
		resourceids.StaticSegment("staticVideoAnalyzers", "videoAnalyzers", "videoAnalyzers"),
		resourceids.UserSpecifiedSegment("videoAnalyzerName", "videoAnalyzerValue"),
		resourceids.StaticSegment("staticEdgeModules", "edgeModules", "edgeModules"),
		resourceids.UserSpecifiedSegment("edgeModuleName", "edgeModuleValue"),
	}
}

// String returns a human-readable description of this Edge Module ID
func (id EdgeModuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Video Analyzer Name: %q", id.VideoAnalyzerName),
		fmt.Sprintf("Edge Module Name: %q", id.EdgeModuleName),
	}
	return fmt.Sprintf("Edge Module (%s)", strings.Join(components, "\n"))
}
