package clusters

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckNameAvailabilityResourceType string

const (
	CheckNameAvailabilityResourceTypeMicrosoftPointDBforPostgreSQLServerGroupsvTwo CheckNameAvailabilityResourceType = "Microsoft.DBforPostgreSQL/serverGroupsv2"
)

func PossibleValuesForCheckNameAvailabilityResourceType() []string {
	return []string{
		string(CheckNameAvailabilityResourceTypeMicrosoftPointDBforPostgreSQLServerGroupsvTwo),
	}
}

func parseCheckNameAvailabilityResourceType(input string) (*CheckNameAvailabilityResourceType, error) {
	vals := map[string]CheckNameAvailabilityResourceType{
		"microsoft.dbforpostgresql/servergroupsv2": CheckNameAvailabilityResourceTypeMicrosoftPointDBforPostgreSQLServerGroupsvTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CheckNameAvailabilityResourceType(input)
	return &out, nil
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateEndpointServiceConnectionStatus() []string {
	return []string{
		string(PrivateEndpointServiceConnectionStatusApproved),
		string(PrivateEndpointServiceConnectionStatusPending),
		string(PrivateEndpointServiceConnectionStatusRejected),
	}
}

func parsePrivateEndpointServiceConnectionStatus(input string) (*PrivateEndpointServiceConnectionStatus, error) {
	vals := map[string]PrivateEndpointServiceConnectionStatus{
		"approved": PrivateEndpointServiceConnectionStatusApproved,
		"pending":  PrivateEndpointServiceConnectionStatusPending,
		"rejected": PrivateEndpointServiceConnectionStatusRejected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateEndpointServiceConnectionStatus(input)
	return &out, nil
}
