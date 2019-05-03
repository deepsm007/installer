package web

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
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

// DeletedWebAppsClient is the webSite Management Client
type DeletedWebAppsClient struct {
	BaseClient
}

// NewDeletedWebAppsClient creates an instance of the DeletedWebAppsClient client.
func NewDeletedWebAppsClient(subscriptionID string) DeletedWebAppsClient {
	return NewDeletedWebAppsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDeletedWebAppsClientWithBaseURI creates an instance of the DeletedWebAppsClient client.
func NewDeletedWebAppsClientWithBaseURI(baseURI string, subscriptionID string) DeletedWebAppsClient {
	return DeletedWebAppsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List get all deleted apps for a subscription.
func (client DeletedWebAppsClient) List(ctx context.Context) (result DeletedWebAppCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedWebAppsClient.List")
		defer func() {
			sc := -1
			if result.dwac.Response.Response != nil {
				sc = result.dwac.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dwac.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "List", resp, "Failure sending request")
		return
	}

	result.dwac, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DeletedWebAppsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/deletedSites", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DeletedWebAppsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DeletedWebAppsClient) ListResponder(resp *http.Response) (result DeletedWebAppCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DeletedWebAppsClient) listNextResults(ctx context.Context, lastResults DeletedWebAppCollection) (result DeletedWebAppCollection, err error) {
	req, err := lastResults.deletedWebAppCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeletedWebAppsClient) ListComplete(ctx context.Context) (result DeletedWebAppCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedWebAppsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}