package deployments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentProvisioningState string

const (
	DeploymentProvisioningStateAccepted  DeploymentProvisioningState = "Accepted"
	DeploymentProvisioningStateCreating  DeploymentProvisioningState = "Creating"
	DeploymentProvisioningStateDeleting  DeploymentProvisioningState = "Deleting"
	DeploymentProvisioningStateFailed    DeploymentProvisioningState = "Failed"
	DeploymentProvisioningStateMoving    DeploymentProvisioningState = "Moving"
	DeploymentProvisioningStateSucceeded DeploymentProvisioningState = "Succeeded"
)

func PossibleValuesForDeploymentProvisioningState() []string {
	return []string{
		string(DeploymentProvisioningStateAccepted),
		string(DeploymentProvisioningStateCreating),
		string(DeploymentProvisioningStateDeleting),
		string(DeploymentProvisioningStateFailed),
		string(DeploymentProvisioningStateMoving),
		string(DeploymentProvisioningStateSucceeded),
	}
}

func (s *DeploymentProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeploymentProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeploymentProvisioningState(input string) (*DeploymentProvisioningState, error) {
	vals := map[string]DeploymentProvisioningState{
		"accepted":  DeploymentProvisioningStateAccepted,
		"creating":  DeploymentProvisioningStateCreating,
		"deleting":  DeploymentProvisioningStateDeleting,
		"failed":    DeploymentProvisioningStateFailed,
		"moving":    DeploymentProvisioningStateMoving,
		"succeeded": DeploymentProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentProvisioningState(input)
	return &out, nil
}

type DeploymentScaleType string

const (
	DeploymentScaleTypeManual   DeploymentScaleType = "Manual"
	DeploymentScaleTypeStandard DeploymentScaleType = "Standard"
)

func PossibleValuesForDeploymentScaleType() []string {
	return []string{
		string(DeploymentScaleTypeManual),
		string(DeploymentScaleTypeStandard),
	}
}

func (s *DeploymentScaleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeploymentScaleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeploymentScaleType(input string) (*DeploymentScaleType, error) {
	vals := map[string]DeploymentScaleType{
		"manual":   DeploymentScaleTypeManual,
		"standard": DeploymentScaleTypeStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentScaleType(input)
	return &out, nil
}
