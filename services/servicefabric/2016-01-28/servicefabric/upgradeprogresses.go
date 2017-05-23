package servicefabric

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

// UpgradeProgressesClient is the client for the UpgradeProgresses methods of
// the Servicefabric service.
type UpgradeProgressesClient struct {
	ManagementClient
}

// NewUpgradeProgressesClient creates an instance of the
// UpgradeProgressesClient client.
func NewUpgradeProgressesClient(timeout *int32) UpgradeProgressesClient {
	return NewUpgradeProgressesClientWithBaseURI(DefaultBaseURI, timeout)
}

// NewUpgradeProgressesClientWithBaseURI creates an instance of the
// UpgradeProgressesClient client.
func NewUpgradeProgressesClientWithBaseURI(baseURI string, timeout *int32) UpgradeProgressesClient {
	return UpgradeProgressesClient{NewWithBaseURI(baseURI, timeout)}
}

// Get get upgrade progresses
func (client UpgradeProgressesClient) Get() (result ClusterUpgradeProgress, err error) {
	req, err := client.GetPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.UpgradeProgressesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.UpgradeProgressesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.UpgradeProgressesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client UpgradeProgressesClient) GetPreparer() (*http.Request, error) {
	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/$/GetUpgradeProgress"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client UpgradeProgressesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client UpgradeProgressesClient) GetResponder(resp *http.Response) (result ClusterUpgradeProgress, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
