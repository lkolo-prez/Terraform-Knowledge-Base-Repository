package prometheusrulegroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrometheusRuleGroupResourceCollection
}

// ListByResourceGroup ...
func (c PrometheusRuleGroupsClient) ListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result ListByResourceGroupOperationResponse, err error) {
	req, err := c.preparerForListByResourceGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prometheusrulegroups.PrometheusRuleGroupsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "prometheusrulegroups.PrometheusRuleGroupsClient", "ListByResourceGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByResourceGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prometheusrulegroups.PrometheusRuleGroupsClient", "ListByResourceGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByResourceGroup prepares the ListByResourceGroup request.
func (c PrometheusRuleGroupsClient) preparerForListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.AlertsManagement/prometheusRuleGroups", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByResourceGroup handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (c PrometheusRuleGroupsClient) responderForListByResourceGroup(resp *http.Response) (result ListByResourceGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
