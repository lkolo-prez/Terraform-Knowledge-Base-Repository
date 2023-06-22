package resourceguards

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceGuardsClient struct {
	Client *resourcemanager.Client
}

func NewResourceGuardsClientWithBaseURI(api environments.Api) (*ResourceGuardsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "resourceguards", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ResourceGuardsClient: %+v", err)
	}

	return &ResourceGuardsClient{
		Client: client,
	}, nil
}
