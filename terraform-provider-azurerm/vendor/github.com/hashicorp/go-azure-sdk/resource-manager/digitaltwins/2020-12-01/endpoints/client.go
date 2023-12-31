package endpoints

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointsClient struct {
	Client *resourcemanager.Client
}

func NewEndpointsClientWithBaseURI(api environments.Api) (*EndpointsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "endpoints", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EndpointsClient: %+v", err)
	}

	return &EndpointsClient{
		Client: client,
	}, nil
}
