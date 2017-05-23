package trafficmanager

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// GeographicHierarchiesClient is the client for the GeographicHierarchies
// methods of the Trafficmanager service.
type GeographicHierarchiesClient struct {
	ManagementClient
}

// NewGeographicHierarchiesClient creates an instance of the
// GeographicHierarchiesClient client.
func NewGeographicHierarchiesClient(subscriptionID string) GeographicHierarchiesClient {
	return NewGeographicHierarchiesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGeographicHierarchiesClientWithBaseURI creates an instance of the
// GeographicHierarchiesClient client.
func NewGeographicHierarchiesClientWithBaseURI(baseURI string, subscriptionID string) GeographicHierarchiesClient {
	return GeographicHierarchiesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetDefault gets the default Geographic Hierarchy used by the Geographic
// traffic routing method.
func (client GeographicHierarchiesClient) GetDefault() (result GeographicHierarchy, err error) {
	req, err := client.GetDefaultPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.GeographicHierarchiesClient", "GetDefault", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDefaultSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "trafficmanager.GeographicHierarchiesClient", "GetDefault", resp, "Failure sending request")
		return
	}

	result, err = client.GetDefaultResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.GeographicHierarchiesClient", "GetDefault", resp, "Failure responding to request")
	}

	return
}

// GetDefaultPreparer prepares the GetDefault request.
func (client GeographicHierarchiesClient) GetDefaultPreparer() (*http.Request, error) {
	const APIVersion = "2017-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Network/trafficManagerGeographicHierarchies/default"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetDefaultSender sends the GetDefault request. The method will close the
// http.Response Body if it receives an error.
func (client GeographicHierarchiesClient) GetDefaultSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetDefaultResponder handles the response to the GetDefault request. The method always
// closes the http.Response Body.
func (client GeographicHierarchiesClient) GetDefaultResponder(resp *http.Response) (result GeographicHierarchy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
