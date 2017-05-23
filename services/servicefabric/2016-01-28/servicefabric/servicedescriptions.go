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

// ServiceDescriptionsClient is the client for the ServiceDescriptions methods
// of the Servicefabric service.
type ServiceDescriptionsClient struct {
	ManagementClient
}

// NewServiceDescriptionsClient creates an instance of the
// ServiceDescriptionsClient client.
func NewServiceDescriptionsClient(timeout *int32) ServiceDescriptionsClient {
	return NewServiceDescriptionsClientWithBaseURI(DefaultBaseURI, timeout)
}

// NewServiceDescriptionsClientWithBaseURI creates an instance of the
// ServiceDescriptionsClient client.
func NewServiceDescriptionsClientWithBaseURI(baseURI string, timeout *int32) ServiceDescriptionsClient {
	return ServiceDescriptionsClient{NewWithBaseURI(baseURI, timeout)}
}

// Get get service descriptions
//
// serviceName is the name of the service
func (client ServiceDescriptionsClient) Get(serviceName string) (result ServiceDescription, err error) {
	req, err := client.GetPreparer(serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceDescriptionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceDescriptionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceDescriptionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServiceDescriptionsClient) GetPreparer(serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceName": serviceName,
	}

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
		autorest.WithPathParameters("/Services/{serviceName}/$/GetDescription", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceDescriptionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServiceDescriptionsClient) GetResponder(resp *http.Response) (result ServiceDescription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
