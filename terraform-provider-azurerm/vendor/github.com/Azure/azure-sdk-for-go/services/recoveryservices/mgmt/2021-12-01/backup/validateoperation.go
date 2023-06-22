package backup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ValidateOperationClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ValidateOperationClient struct {
	BaseClient
}

// NewValidateOperationClient creates an instance of the ValidateOperationClient client.
func NewValidateOperationClient(subscriptionID string) ValidateOperationClient {
	return NewValidateOperationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewValidateOperationClientWithBaseURI creates an instance of the ValidateOperationClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewValidateOperationClientWithBaseURI(baseURI string, subscriptionID string) ValidateOperationClient {
	return ValidateOperationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Trigger validate operation for specified backed up item in the form of an asynchronous operation. Returns tracking
// headers which can be tracked using GetValidateOperationResult API.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// parameters - resource validate operation request
func (client ValidateOperationClient) Trigger(ctx context.Context, vaultName string, resourceGroupName string, parameters BasicValidateOperationRequest) (result ValidateOperationTriggerFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ValidateOperationClient.Trigger")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TriggerPreparer(ctx, vaultName, resourceGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ValidateOperationClient", "Trigger", nil, "Failure preparing request")
		return
	}

	result, err = client.TriggerSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ValidateOperationClient", "Trigger", result.Response(), "Failure sending request")
		return
	}

	return
}

// TriggerPreparer prepares the Trigger request.
func (client ValidateOperationClient) TriggerPreparer(ctx context.Context, vaultName string, resourceGroupName string, parameters BasicValidateOperationRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupTriggerValidateOperation", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TriggerSender sends the Trigger request. The method will close the
// http.Response Body if it receives an error.
func (client ValidateOperationClient) TriggerSender(req *http.Request) (future ValidateOperationTriggerFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// TriggerResponder handles the response to the Trigger request. The method always
// closes the http.Response Body.
func (client ValidateOperationClient) TriggerResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
