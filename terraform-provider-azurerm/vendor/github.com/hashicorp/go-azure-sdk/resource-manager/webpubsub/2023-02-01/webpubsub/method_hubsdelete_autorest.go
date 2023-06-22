package webpubsub

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HubsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// HubsDelete ...
func (c WebPubSubClient) HubsDelete(ctx context.Context, id HubId) (result HubsDeleteOperationResponse, err error) {
	req, err := c.preparerForHubsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webpubsub.WebPubSubClient", "HubsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForHubsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webpubsub.WebPubSubClient", "HubsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// HubsDeleteThenPoll performs HubsDelete then polls until it's completed
func (c WebPubSubClient) HubsDeleteThenPoll(ctx context.Context, id HubId) error {
	result, err := c.HubsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing HubsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after HubsDelete: %+v", err)
	}

	return nil
}

// preparerForHubsDelete prepares the HubsDelete request.
func (c WebPubSubClient) preparerForHubsDelete(ctx context.Context, id HubId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForHubsDelete sends the HubsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c WebPubSubClient) senderForHubsDelete(ctx context.Context, req *http.Request) (future HubsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
