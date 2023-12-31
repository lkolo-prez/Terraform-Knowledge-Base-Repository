package client

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/communication/2023-03-31/communicationservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/communication/2023-03-31/emailservices"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	ServiceClient *communicationservices.CommunicationServicesClient

	EmailServicesClient *emailservices.EmailServicesClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	servicesClient, err := communicationservices.NewCommunicationServicesClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Service client: %+v", err)
	}
	o.Configure(servicesClient.Client, o.Authorizers.ResourceManager)

	emailServicesClient, err := emailservices.NewEmailServicesClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Email Service client: %+v", err)
	}
	o.Configure(emailServicesClient.Client, o.Authorizers.ResourceManager)

	return &Client{
		ServiceClient:       servicesClient,
		EmailServicesClient: emailServicesClient,
	}, nil
}
