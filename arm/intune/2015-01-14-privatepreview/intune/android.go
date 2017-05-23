package intune

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

// AndroidClient is the microsoft.Intune Resource provider Api features in the
// swagger-2.0 specification
type AndroidClient struct {
	ManagementClient
}

// NewAndroidClient creates an instance of the AndroidClient client.
func NewAndroidClient() AndroidClient {
	return NewAndroidClientWithBaseURI(DefaultBaseURI)
}

// NewAndroidClientWithBaseURI creates an instance of the AndroidClient client.
func NewAndroidClientWithBaseURI(baseURI string) AndroidClient {
	return AndroidClient{NewWithBaseURI(baseURI)}
}

// AddAppForMAMPolicy add app to an AndroidMAMPolicy.
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy appName is application unique Name parameters is parameters
// supplied to the Create or update app to an android policy operation.
func (client AndroidClient) AddAppForMAMPolicy(hostName string, policyName string, appName string, parameters MAMPolicyAppIDOrGroupIDPayload) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Properties.URL", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "intune.AndroidClient", "AddAppForMAMPolicy")
	}

	req, err := client.AddAppForMAMPolicyPreparer(hostName, policyName, appName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "AddAppForMAMPolicy", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddAppForMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "AddAppForMAMPolicy", resp, "Failure sending request")
		return
	}

	result, err = client.AddAppForMAMPolicyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "AddAppForMAMPolicy", resp, "Failure responding to request")
	}

	return
}

// AddAppForMAMPolicyPreparer prepares the AddAppForMAMPolicy request.
func (client AndroidClient) AddAppForMAMPolicyPreparer(hostName string, policyName string, appName string, parameters MAMPolicyAppIDOrGroupIDPayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":    autorest.Encode("path", appName),
		"hostName":   autorest.Encode("path", hostName),
		"policyName": autorest.Encode("path", policyName),
	}

	const APIVersion = "2015-01-14-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}/apps/{appName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// AddAppForMAMPolicySender sends the AddAppForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client AndroidClient) AddAppForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// AddAppForMAMPolicyResponder handles the response to the AddAppForMAMPolicy request. The method always
// closes the http.Response Body.
func (client AndroidClient) AddAppForMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// AddGroupForMAMPolicy add group to an AndroidMAMPolicy.
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy groupID is group Id parameters is parameters supplied to the
// Create or update app to an android policy operation.
func (client AndroidClient) AddGroupForMAMPolicy(hostName string, policyName string, groupID string, parameters MAMPolicyAppIDOrGroupIDPayload) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Properties.URL", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "intune.AndroidClient", "AddGroupForMAMPolicy")
	}

	req, err := client.AddGroupForMAMPolicyPreparer(hostName, policyName, groupID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "AddGroupForMAMPolicy", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddGroupForMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "AddGroupForMAMPolicy", resp, "Failure sending request")
		return
	}

	result, err = client.AddGroupForMAMPolicyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "AddGroupForMAMPolicy", resp, "Failure responding to request")
	}

	return
}

// AddGroupForMAMPolicyPreparer prepares the AddGroupForMAMPolicy request.
func (client AndroidClient) AddGroupForMAMPolicyPreparer(hostName string, policyName string, groupID string, parameters MAMPolicyAppIDOrGroupIDPayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":    autorest.Encode("path", groupID),
		"hostName":   autorest.Encode("path", hostName),
		"policyName": autorest.Encode("path", policyName),
	}

	const APIVersion = "2015-01-14-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}/groups/{groupId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// AddGroupForMAMPolicySender sends the AddGroupForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client AndroidClient) AddGroupForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// AddGroupForMAMPolicyResponder handles the response to the AddGroupForMAMPolicy request. The method always
// closes the http.Response Body.
func (client AndroidClient) AddGroupForMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdateMAMPolicy creates or updates AndroidMAMPolicy.
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy parameters is parameters supplied to the Create or update an
// android policy operation.
func (client AndroidClient) CreateOrUpdateMAMPolicy(hostName string, policyName string, parameters AndroidMAMPolicy) (result AndroidMAMPolicy, err error) {
	req, err := client.CreateOrUpdateMAMPolicyPreparer(hostName, policyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "CreateOrUpdateMAMPolicy", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "CreateOrUpdateMAMPolicy", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateMAMPolicyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "CreateOrUpdateMAMPolicy", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateMAMPolicyPreparer prepares the CreateOrUpdateMAMPolicy request.
func (client AndroidClient) CreateOrUpdateMAMPolicyPreparer(hostName string, policyName string, parameters AndroidMAMPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":   autorest.Encode("path", hostName),
		"policyName": autorest.Encode("path", policyName),
	}

	const APIVersion = "2015-01-14-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateMAMPolicySender sends the CreateOrUpdateMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client AndroidClient) CreateOrUpdateMAMPolicySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateMAMPolicyResponder handles the response to the CreateOrUpdateMAMPolicy request. The method always
// closes the http.Response Body.
func (client AndroidClient) CreateOrUpdateMAMPolicyResponder(resp *http.Response) (result AndroidMAMPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteAppForMAMPolicy delete App for Android Policy
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy appName is application unique Name
func (client AndroidClient) DeleteAppForMAMPolicy(hostName string, policyName string, appName string) (result autorest.Response, err error) {
	req, err := client.DeleteAppForMAMPolicyPreparer(hostName, policyName, appName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "DeleteAppForMAMPolicy", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteAppForMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "DeleteAppForMAMPolicy", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteAppForMAMPolicyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "DeleteAppForMAMPolicy", resp, "Failure responding to request")
	}

	return
}

// DeleteAppForMAMPolicyPreparer prepares the DeleteAppForMAMPolicy request.
func (client AndroidClient) DeleteAppForMAMPolicyPreparer(hostName string, policyName string, appName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":    autorest.Encode("path", appName),
		"hostName":   autorest.Encode("path", hostName),
		"policyName": autorest.Encode("path", policyName),
	}

	const APIVersion = "2015-01-14-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}/apps/{appName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteAppForMAMPolicySender sends the DeleteAppForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client AndroidClient) DeleteAppForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteAppForMAMPolicyResponder handles the response to the DeleteAppForMAMPolicy request. The method always
// closes the http.Response Body.
func (client AndroidClient) DeleteAppForMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteGroupForMAMPolicy delete Group for Android Policy
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy groupID is application unique Name
func (client AndroidClient) DeleteGroupForMAMPolicy(hostName string, policyName string, groupID string) (result autorest.Response, err error) {
	req, err := client.DeleteGroupForMAMPolicyPreparer(hostName, policyName, groupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "DeleteGroupForMAMPolicy", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteGroupForMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "DeleteGroupForMAMPolicy", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteGroupForMAMPolicyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "DeleteGroupForMAMPolicy", resp, "Failure responding to request")
	}

	return
}

// DeleteGroupForMAMPolicyPreparer prepares the DeleteGroupForMAMPolicy request.
func (client AndroidClient) DeleteGroupForMAMPolicyPreparer(hostName string, policyName string, groupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":    autorest.Encode("path", groupID),
		"hostName":   autorest.Encode("path", hostName),
		"policyName": autorest.Encode("path", policyName),
	}

	const APIVersion = "2015-01-14-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}/groups/{groupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteGroupForMAMPolicySender sends the DeleteGroupForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client AndroidClient) DeleteGroupForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteGroupForMAMPolicyResponder handles the response to the DeleteGroupForMAMPolicy request. The method always
// closes the http.Response Body.
func (client AndroidClient) DeleteGroupForMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteMAMPolicy delete Android Policy
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy
func (client AndroidClient) DeleteMAMPolicy(hostName string, policyName string) (result autorest.Response, err error) {
	req, err := client.DeleteMAMPolicyPreparer(hostName, policyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "DeleteMAMPolicy", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "DeleteMAMPolicy", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteMAMPolicyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "DeleteMAMPolicy", resp, "Failure responding to request")
	}

	return
}

// DeleteMAMPolicyPreparer prepares the DeleteMAMPolicy request.
func (client AndroidClient) DeleteMAMPolicyPreparer(hostName string, policyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":   autorest.Encode("path", hostName),
		"policyName": autorest.Encode("path", policyName),
	}

	const APIVersion = "2015-01-14-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteMAMPolicySender sends the DeleteMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client AndroidClient) DeleteMAMPolicySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteMAMPolicyResponder handles the response to the DeleteMAMPolicy request. The method always
// closes the http.Response Body.
func (client AndroidClient) DeleteMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetAppForMAMPolicy get apps for an AndroidMAMPolicy.
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy filter is the filter to apply on the operation. selectParameter
// is select specific fields in entity.
func (client AndroidClient) GetAppForMAMPolicy(hostName string, policyName string, filter string, top *int32, selectParameter string) (result ApplicationCollection, err error) {
	req, err := client.GetAppForMAMPolicyPreparer(hostName, policyName, filter, top, selectParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetAppForMAMPolicy", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAppForMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetAppForMAMPolicy", resp, "Failure sending request")
		return
	}

	result, err = client.GetAppForMAMPolicyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetAppForMAMPolicy", resp, "Failure responding to request")
	}

	return
}

// GetAppForMAMPolicyPreparer prepares the GetAppForMAMPolicy request.
func (client AndroidClient) GetAppForMAMPolicyPreparer(hostName string, policyName string, filter string, top *int32, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":   autorest.Encode("path", hostName),
		"policyName": autorest.Encode("path", policyName),
	}

	const APIVersion = "2015-01-14-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Intune/locations/{hostName}/AndroidPolicies/{policyName}/apps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetAppForMAMPolicySender sends the GetAppForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client AndroidClient) GetAppForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetAppForMAMPolicyResponder handles the response to the GetAppForMAMPolicy request. The method always
// closes the http.Response Body.
func (client AndroidClient) GetAppForMAMPolicyResponder(resp *http.Response) (result ApplicationCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAppForMAMPolicyNextResults retrieves the next set of results, if any.
func (client AndroidClient) GetAppForMAMPolicyNextResults(lastResults ApplicationCollection) (result ApplicationCollection, err error) {
	req, err := lastResults.ApplicationCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune.AndroidClient", "GetAppForMAMPolicy", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetAppForMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune.AndroidClient", "GetAppForMAMPolicy", resp, "Failure sending next results request")
	}

	result, err = client.GetAppForMAMPolicyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetAppForMAMPolicy", resp, "Failure responding to next results request")
	}

	return
}

// GetGroupsForMAMPolicy returns groups for a given AndroidMAMPolicy.
//
// hostName is location hostName for the tenant policyName is policy name for
// the tenant
func (client AndroidClient) GetGroupsForMAMPolicy(hostName string, policyName string) (result GroupsCollection, err error) {
	req, err := client.GetGroupsForMAMPolicyPreparer(hostName, policyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetGroupsForMAMPolicy", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetGroupsForMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetGroupsForMAMPolicy", resp, "Failure sending request")
		return
	}

	result, err = client.GetGroupsForMAMPolicyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetGroupsForMAMPolicy", resp, "Failure responding to request")
	}

	return
}

// GetGroupsForMAMPolicyPreparer prepares the GetGroupsForMAMPolicy request.
func (client AndroidClient) GetGroupsForMAMPolicyPreparer(hostName string, policyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":   autorest.Encode("path", hostName),
		"policyName": autorest.Encode("path", policyName),
	}

	const APIVersion = "2015-01-14-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}/groups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetGroupsForMAMPolicySender sends the GetGroupsForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client AndroidClient) GetGroupsForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetGroupsForMAMPolicyResponder handles the response to the GetGroupsForMAMPolicy request. The method always
// closes the http.Response Body.
func (client AndroidClient) GetGroupsForMAMPolicyResponder(resp *http.Response) (result GroupsCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetGroupsForMAMPolicyNextResults retrieves the next set of results, if any.
func (client AndroidClient) GetGroupsForMAMPolicyNextResults(lastResults GroupsCollection) (result GroupsCollection, err error) {
	req, err := lastResults.GroupsCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune.AndroidClient", "GetGroupsForMAMPolicy", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetGroupsForMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune.AndroidClient", "GetGroupsForMAMPolicy", resp, "Failure sending next results request")
	}

	result, err = client.GetGroupsForMAMPolicyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetGroupsForMAMPolicy", resp, "Failure responding to next results request")
	}

	return
}

// GetMAMPolicies returns Intune Android policies.
//
// hostName is location hostName for the tenant filter is the filter to apply
// on the operation. selectParameter is select specific fields in entity.
func (client AndroidClient) GetMAMPolicies(hostName string, filter string, top *int32, selectParameter string) (result AndroidMAMPolicyCollection, err error) {
	req, err := client.GetMAMPoliciesPreparer(hostName, filter, top, selectParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetMAMPolicies", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMAMPoliciesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetMAMPolicies", resp, "Failure sending request")
		return
	}

	result, err = client.GetMAMPoliciesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetMAMPolicies", resp, "Failure responding to request")
	}

	return
}

// GetMAMPoliciesPreparer prepares the GetMAMPolicies request.
func (client AndroidClient) GetMAMPoliciesPreparer(hostName string, filter string, top *int32, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": autorest.Encode("path", hostName),
	}

	const APIVersion = "2015-01-14-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetMAMPoliciesSender sends the GetMAMPolicies request. The method will close the
// http.Response Body if it receives an error.
func (client AndroidClient) GetMAMPoliciesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetMAMPoliciesResponder handles the response to the GetMAMPolicies request. The method always
// closes the http.Response Body.
func (client AndroidClient) GetMAMPoliciesResponder(resp *http.Response) (result AndroidMAMPolicyCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMAMPoliciesNextResults retrieves the next set of results, if any.
func (client AndroidClient) GetMAMPoliciesNextResults(lastResults AndroidMAMPolicyCollection) (result AndroidMAMPolicyCollection, err error) {
	req, err := lastResults.AndroidMAMPolicyCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune.AndroidClient", "GetMAMPolicies", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetMAMPoliciesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune.AndroidClient", "GetMAMPolicies", resp, "Failure sending next results request")
	}

	result, err = client.GetMAMPoliciesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetMAMPolicies", resp, "Failure responding to next results request")
	}

	return
}

// GetMAMPolicyByName returns AndroidMAMPolicy with given name.
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy selectParameter is select specific fields in entity.
func (client AndroidClient) GetMAMPolicyByName(hostName string, policyName string, selectParameter string) (result AndroidMAMPolicy, err error) {
	req, err := client.GetMAMPolicyByNamePreparer(hostName, policyName, selectParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetMAMPolicyByName", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMAMPolicyByNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetMAMPolicyByName", resp, "Failure sending request")
		return
	}

	result, err = client.GetMAMPolicyByNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "GetMAMPolicyByName", resp, "Failure responding to request")
	}

	return
}

// GetMAMPolicyByNamePreparer prepares the GetMAMPolicyByName request.
func (client AndroidClient) GetMAMPolicyByNamePreparer(hostName string, policyName string, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":   autorest.Encode("path", hostName),
		"policyName": autorest.Encode("path", policyName),
	}

	const APIVersion = "2015-01-14-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetMAMPolicyByNameSender sends the GetMAMPolicyByName request. The method will close the
// http.Response Body if it receives an error.
func (client AndroidClient) GetMAMPolicyByNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetMAMPolicyByNameResponder handles the response to the GetMAMPolicyByName request. The method always
// closes the http.Response Body.
func (client AndroidClient) GetMAMPolicyByNameResponder(resp *http.Response) (result AndroidMAMPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PatchMAMPolicy patch AndroidMAMPolicy.
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy parameters is parameters supplied to the Create or update an
// android policy operation.
func (client AndroidClient) PatchMAMPolicy(hostName string, policyName string, parameters AndroidMAMPolicy) (result AndroidMAMPolicy, err error) {
	req, err := client.PatchMAMPolicyPreparer(hostName, policyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "PatchMAMPolicy", nil, "Failure preparing request")
		return
	}

	resp, err := client.PatchMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "PatchMAMPolicy", resp, "Failure sending request")
		return
	}

	result, err = client.PatchMAMPolicyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune.AndroidClient", "PatchMAMPolicy", resp, "Failure responding to request")
	}

	return
}

// PatchMAMPolicyPreparer prepares the PatchMAMPolicy request.
func (client AndroidClient) PatchMAMPolicyPreparer(hostName string, policyName string, parameters AndroidMAMPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":   autorest.Encode("path", hostName),
		"policyName": autorest.Encode("path", policyName),
	}

	const APIVersion = "2015-01-14-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// PatchMAMPolicySender sends the PatchMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client AndroidClient) PatchMAMPolicySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PatchMAMPolicyResponder handles the response to the PatchMAMPolicy request. The method always
// closes the http.Response Body.
func (client AndroidClient) PatchMAMPolicyResponder(resp *http.Response) (result AndroidMAMPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
