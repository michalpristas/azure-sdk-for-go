package job

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
	"github.com/satori/uuid"
	"net/http"
)

// GroupClient is the creates an Azure Data Lake Analytics job client.
type GroupClient struct {
	ManagementClient
}

// NewGroupClient creates an instance of the GroupClient client.
func NewGroupClient() GroupClient {
	return GroupClient{New()}
}

// Build builds (compiles) the specified job in the specified Data Lake
// Analytics account for job correctness and validation.
//
// accountName is the Azure Data Lake Analytics account to execute job
// operations on. parameters is the parameters to build a job.
func (client GroupClient) Build(accountName string, parameters Information) (result Information, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Properties", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.Properties.Script", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "job.GroupClient", "Build")
	}

	req, err := client.BuildPreparer(accountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "Build", nil, "Failure preparing request")
		return
	}

	resp, err := client.BuildSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.GroupClient", "Build", resp, "Failure sending request")
		return
	}

	result, err = client.BuildResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "Build", resp, "Failure responding to request")
	}

	return
}

// BuildPreparer prepares the Build request.
func (client GroupClient) BuildPreparer(accountName string, parameters Information) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPath("/BuildJob"),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// BuildSender sends the Build request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) BuildSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// BuildResponder handles the response to the Build request. The method always
// closes the http.Response Body.
func (client GroupClient) BuildResponder(resp *http.Response) (result Information, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Cancel cancels the running job specified by the job ID.
//
// accountName is the Azure Data Lake Analytics account to execute job
// operations on. jobIdentity is jobInfo ID to cancel.
func (client GroupClient) Cancel(accountName string, jobIdentity uuid.UUID) (result autorest.Response, err error) {
	req, err := client.CancelPreparer(accountName, jobIdentity)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "job.GroupClient", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "Cancel", resp, "Failure responding to request")
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client GroupClient) CancelPreparer(accountName string, jobIdentity uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/Jobs/{jobIdentity}/CancelJob", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) CancelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client GroupClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Create submits a job to the specified Data Lake Analytics account.
//
// accountName is the Azure Data Lake Analytics account to execute job
// operations on. jobIdentity is the job ID (a GUID) for the job being
// submitted. parameters is the parameters to submit a job.
func (client GroupClient) Create(accountName string, jobIdentity uuid.UUID, parameters Information) (result Information, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Properties", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.Properties.Script", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "job.GroupClient", "Create")
	}

	req, err := client.CreatePreparer(accountName, jobIdentity, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.GroupClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client GroupClient) CreatePreparer(accountName string, jobIdentity uuid.UUID, parameters Information) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/Jobs/{jobIdentity}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client GroupClient) CreateResponder(resp *http.Response) (result Information, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the job information for the specified job ID.
//
// accountName is the Azure Data Lake Analytics account to execute job
// operations on. jobIdentity is jobInfo ID.
func (client GroupClient) Get(accountName string, jobIdentity uuid.UUID) (result Information, err error) {
	req, err := client.GetPreparer(accountName, jobIdentity)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.GroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client GroupClient) GetPreparer(accountName string, jobIdentity uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/Jobs/{jobIdentity}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GroupClient) GetResponder(resp *http.Response) (result Information, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDebugDataPath gets the U-SQL job debug data information specified by the
// job ID.
//
// accountName is the Azure Data Lake Analytics account to execute job
// operations on. jobIdentity is jobInfo ID.
func (client GroupClient) GetDebugDataPath(accountName string, jobIdentity uuid.UUID) (result DataPath, err error) {
	req, err := client.GetDebugDataPathPreparer(accountName, jobIdentity)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "GetDebugDataPath", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDebugDataPathSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.GroupClient", "GetDebugDataPath", resp, "Failure sending request")
		return
	}

	result, err = client.GetDebugDataPathResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "GetDebugDataPath", resp, "Failure responding to request")
	}

	return
}

// GetDebugDataPathPreparer prepares the GetDebugDataPath request.
func (client GroupClient) GetDebugDataPathPreparer(accountName string, jobIdentity uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/Jobs/{jobIdentity}/GetDebugDataPath", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetDebugDataPathSender sends the GetDebugDataPath request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetDebugDataPathSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetDebugDataPathResponder handles the response to the GetDebugDataPath request. The method always
// closes the http.Response Body.
func (client GroupClient) GetDebugDataPathResponder(resp *http.Response) (result DataPath, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetStatistics gets statistics of the specified job.
//
// accountName is the Azure Data Lake Analytics account to execute job
// operations on. jobIdentity is jobInfo ID.
func (client GroupClient) GetStatistics(accountName string, jobIdentity uuid.UUID) (result Statistics, err error) {
	req, err := client.GetStatisticsPreparer(accountName, jobIdentity)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "GetStatistics", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetStatisticsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.GroupClient", "GetStatistics", resp, "Failure sending request")
		return
	}

	result, err = client.GetStatisticsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "GetStatistics", resp, "Failure responding to request")
	}

	return
}

// GetStatisticsPreparer prepares the GetStatistics request.
func (client GroupClient) GetStatisticsPreparer(accountName string, jobIdentity uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/Jobs/{jobIdentity}/GetStatistics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetStatisticsSender sends the GetStatistics request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetStatisticsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetStatisticsResponder handles the response to the GetStatistics request. The method always
// closes the http.Response Body.
func (client GroupClient) GetStatisticsResponder(resp *http.Response) (result Statistics, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the jobs, if any, associated with the specified Data Lake
// Analytics account. The response includes a link to the next page of results,
// if any.
//
// accountName is the Azure Data Lake Analytics account to execute job
// operations on. filter is oData filter. Optional. top is the number of items
// to return. Optional. skip is the number of items to skip over before
// returning elements. Optional. expand is oData expansion. Expand related
// resources in line with the retrieved resources, e.g.
// Categories/$expand=Products would expand Product data in line with each
// Category entry. Optional. selectParameter is oData Select statement. Limits
// the properties on each entry to just those requested, e.g.
// Categories?$select=CategoryName,Description. Optional. orderby is orderBy
// clause. One or more comma-separated expressions with an optional "asc" (the
// default) or "desc" depending on the order you'd like the values sorted, e.g.
// Categories?$orderby=CategoryName desc. Optional. count is the Boolean value
// of true or false to request a count of the matching resources included with
// the resources in the response, e.g. Categories?$count=true. Optional. search
// is a free form search. A free-text search expression to match for whether a
// particular entry should be included in the feed, e.g.
// Categories?$search=blue OR green. Optional. formatParameter is the return
// format. Return the response in particular formatxii without access to
// request headers for standard content-type negotiation (e.g
// Orders?$format=json). Optional.
func (client GroupClient) List(accountName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result InfoListResult, err error) {
	req, err := client.ListPreparer(accountName, filter, top, skip, expand, selectParameter, orderby, count, search, formatParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.GroupClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client GroupClient) ListPreparer(accountName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if count != nil {
		queryParameters["$count"] = autorest.Encode("query", *count)
	}
	if len(search) > 0 {
		queryParameters["$search"] = autorest.Encode("query", search)
	}
	if len(formatParameter) > 0 {
		queryParameters["$format"] = autorest.Encode("query", formatParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPath("/Jobs"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client GroupClient) ListResponder(resp *http.Response) (result InfoListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client GroupClient) ListNextResults(lastResults InfoListResult) (result InfoListResult, err error) {
	req, err := lastResults.InfoListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "job.GroupClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "job.GroupClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.GroupClient", "List", resp, "Failure responding to next results request")
	}

	return
}
