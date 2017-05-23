package insightsmanagementclientautoscale

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// AutoscaleSettingsClient is the client for the AutoscaleSettings methods of
// the Insightsmanagementclientautoscale service.
type AutoscaleSettingsClient struct {
	ManagementClient
}

// NewAutoscaleSettingsClient creates an instance of the
// AutoscaleSettingsClient client.
func NewAutoscaleSettingsClient(subscriptionID string) AutoscaleSettingsClient {
	return NewAutoscaleSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAutoscaleSettingsClientWithBaseURI creates an instance of the
// AutoscaleSettingsClient client.
func NewAutoscaleSettingsClientWithBaseURI(baseURI string, subscriptionID string) AutoscaleSettingsClient {
	return AutoscaleSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an autoscale setting.
//
// resourceGroupName is the name of the resource group. autoscaleSettingName is
// the autoscale setting name. parameters is parameters supplied to the
// operation.
func (client AutoscaleSettingsClient) CreateOrUpdate(resourceGroupName string, autoscaleSettingName string, parameters AutoscaleSettingResource) (result AutoscaleSettingResource, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.AutoscaleSetting", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.AutoscaleSetting.Profiles", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.AutoscaleSetting.Profiles", Name: validation.MaxItems, Rule: 20, Chain: nil}}},
					{Target: "parameters.AutoscaleSetting.Name", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, autoscaleSettingName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AutoscaleSettingsClient) CreateOrUpdatePreparer(resourceGroupName string, autoscaleSettingName string, parameters AutoscaleSettingResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"autoscaleSettingName": autorest.Encode("path", autoscaleSettingName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/autoscalesettings/{autoscaleSettingName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AutoscaleSettingsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AutoscaleSettingsClient) CreateOrUpdateResponder(resp *http.Response) (result AutoscaleSettingResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes and autoscale setting
//
// resourceGroupName is the name of the resource group. autoscaleSettingName is
// the autoscale setting name.
func (client AutoscaleSettingsClient) Delete(resourceGroupName string, autoscaleSettingName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, autoscaleSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AutoscaleSettingsClient) DeletePreparer(resourceGroupName string, autoscaleSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"autoscaleSettingName": autorest.Encode("path", autoscaleSettingName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/autoscalesettings/{autoscaleSettingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AutoscaleSettingsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AutoscaleSettingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an autoscale setting
//
// resourceGroupName is the name of the resource group. autoscaleSettingName is
// the autoscale setting name.
func (client AutoscaleSettingsClient) Get(resourceGroupName string, autoscaleSettingName string) (result AutoscaleSettingResource, err error) {
	req, err := client.GetPreparer(resourceGroupName, autoscaleSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AutoscaleSettingsClient) GetPreparer(resourceGroupName string, autoscaleSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"autoscaleSettingName": autorest.Encode("path", autoscaleSettingName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/autoscalesettings/{autoscaleSettingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AutoscaleSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AutoscaleSettingsClient) GetResponder(resp *http.Response) (result AutoscaleSettingResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup lists the autoscale settings for a resource group
//
// resourceGroupName is the name of the resource group. filter is the filter to
// apply on the operation. For more information please see
// https://msdn.microsoft.com/en-us/library/azure/dn931934.aspx
func (client AutoscaleSettingsClient) ListByResourceGroup(resourceGroupName string, filter string) (result AutoscaleSettingResourceCollection, err error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client AutoscaleSettingsClient) ListByResourceGroupPreparer(resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/autoscalesettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AutoscaleSettingsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client AutoscaleSettingsClient) ListByResourceGroupResponder(resp *http.Response) (result AutoscaleSettingResourceCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client AutoscaleSettingsClient) ListByResourceGroupNextResults(lastResults AutoscaleSettingResourceCollection) (result AutoscaleSettingResourceCollection, err error) {
	req, err := lastResults.AutoscaleSettingResourceCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "ListByResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "ListByResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insightsmanagementclientautoscale.AutoscaleSettingsClient", "ListByResourceGroup", resp, "Failure responding to next results request")
	}

	return
}
