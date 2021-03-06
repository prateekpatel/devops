package compute
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
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	//"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
	"fmt"
)

// UsageOperationsClient is the the Compute Management Client.
type DynamicUsageOperationsClient struct {
	ManagementClient
}

// NewUsageOperationsClient creates an instance of the UsageOperationsClient
// client.
func NewDynamicUsageOperationsClient(subscriptionID string) DynamicUsageOperationsClient {
	return NewDynamicUsageOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsageOperationsClientWithBaseURI creates an instance of the
// UsageOperationsClient client.
func NewDynamicUsageOperationsClientWithBaseURI(baseURI string, subscriptionID string) DynamicUsageOperationsClient {
	return DynamicUsageOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets, for the specified location, the current compute resource usage
// information as well as the limits for compute resources under the
// subscription.
//
// location is the location for which resource usage is queried.
func (client DynamicUsageOperationsClient) ListDynamic(name string, resourceGroupName string) (result ListDynamicUsagesResult, err error) {
	/*if err := validation.Validate([]validation.Validation{
		{TargetValue: name,
			Constraints: []validation.Constraint{{Target: "name", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "compute.UsageOperationsClient", "ListDynamic")
	}*/

	req, err := client.ListPreparer(name, resourceGroupName)
	fmt.Println(name, resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.UsageOperationsClient", "ListDynamic", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.UsageOperationsClient", "ListDynamic", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.UsageOperationsClient", "ListDynamic", resp, "Failure responding to request")
	}
	//fmt.Println(result)
	return
}

// ListPreparer prepares the List request.
func (client DynamicUsageOperationsClient) ListPreparer(name string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":       autorest.Encode("path", name),
		"resourceGroupName":	autorest.Encode("path", resourceGroupName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": "2016-09-01",
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{name}/providers/microsoft.insights/metrics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DynamicUsageOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DynamicUsageOperationsClient) ListResponder(resp *http.Response) (result ListDynamicUsagesResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	//fmt.Println(result.Response)
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client DynamicUsageOperationsClient) ListNextResults(lastResults ListUsagesResult) (result ListDynamicUsagesResult, err error) {
	req, err := lastResults.ListUsagesResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.UsageOperationsClient", "ListDynamic", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.UsageOperationsClient", "ListDynamic", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.UsageOperationsClient", "ListDynamic", resp, "Failure responding to next results request")
	}

	return
}
