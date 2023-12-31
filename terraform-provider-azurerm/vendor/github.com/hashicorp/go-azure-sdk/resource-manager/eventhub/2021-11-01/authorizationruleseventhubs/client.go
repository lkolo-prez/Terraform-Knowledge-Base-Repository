package authorizationruleseventhubs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationRulesEventHubsClient struct {
	Client *resourcemanager.Client
}

func NewAuthorizationRulesEventHubsClientWithBaseURI(api environments.Api) (*AuthorizationRulesEventHubsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "authorizationruleseventhubs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthorizationRulesEventHubsClient: %+v", err)
	}

	return &AuthorizationRulesEventHubsClient{
		Client: client,
	}, nil
}
