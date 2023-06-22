package liveoutputs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LiveOutputsClient struct {
	Client *resourcemanager.Client
}

func NewLiveOutputsClientWithBaseURI(api environments.Api) (*LiveOutputsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "liveoutputs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LiveOutputsClient: %+v", err)
	}

	return &LiveOutputsClient{
		Client: client,
	}, nil
}
