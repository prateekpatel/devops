// Copyright (c) 2014 Hewlett-Packard Development Company, L.P.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

/*
Package image implements a client library for accessing OpenStack Image V1 service

Images and ImageDetails can be retrieved using the api.

In addition more complex filtering and sort queries can by using the ImageQueryParameters.

*/
package compute
import (
	"encoding/json"
	"errors"

	"io/ioutil"

	"net/http"
	"net/url"
	"gclassec/goClientCompute/openstack"
	"gclassec/goClientCompute/util"
	//"fmt"
)
// Service is a client service that can make
// requests against a OpenStack version 2 Compute service.
// Below is an example on creating an Compute service and getting Instances:
// 	computeService := compute.ComputeService{Client: *http.DefaultClient, TokenId: tokenId, Url: "http://computeservicelocation"}
//  instance:= computeService.Instances()
type Service struct {
	Session openstack.Session
	Client  http.Client
	URL     string
}

// Response is a structure for all properties of
// an instance for a non detailed query
type Response struct {
	Network   			string `json:"network"`
	diskConfig			string `json:"diskConfig"`
	availability_zone      		string `json:"availability_zone"`
	host              		string `json:"host"`
	instance_name            	string `json:"instance_name"`
	power_state   			int64  `json:"power_state"`
	vm_state	      		string `json:"vm_state"`
	flavor				string `json:"flavor"`
	hostId		      		string `json:"hostId"`
	id	       			string `json:"id"`
	image	            		string `json:"image"`
	key_name            		string `json:"key_name"`
	name	       			string `json:"name"`
	volumes_attached            	string `json:"volumes_attached"`
	security_groups	            	string `json:"security_groups"`
	status            		string `json:"status"`
	tenant_id   			string `json:"tenant_id"`
}
// DetailResponse is a structure for all properties of
// an instance for a detailed query
type DetailResponse struct {
	Name	       			string `json:"name"`
	progress            		int64  `json:"progress"`
	security_groups	            	string `json:"security_groups"`
	status            		string `json:"status"`
	tenant_id   			string `json:"tenant_id"`
	updated				string `json:"updated"`
	user_id		      		string `json:"user_id"`
}
// QueryParameters is a structure that
// contains the filter, sort, and paging parameters for
// an instance or computedetail query.
type QueryParameters struct {
	instance_name           string
}
// SortDirection of the sort, ascending or descending.
type SortDirection string

const (
	// Desc specifies the sort direction to be descending.
	Desc SortDirection = "desc"
	// Asc specifies the sort direction to be ascending.
	Asc SortDirection = "asc"
)



// Compute will issue a get request to OpenStack to retrieve the list of instances.
func (computeService Service) Instances() (instance []Response, err error) {
	return computeService.QueryInstances(nil)
}

// InstancesDetail will issue a get request to OpenStack to retrieve the list of Instances complete with
// additional details.
func (computeService Service) InstancesDetail() (instance []DetailResponse, err error) {
	return computeService.QueryInstancesDetail(nil)
}


// QueryInstances will issue a get request with the specified InstanceQueryParameters to retrieve the list of
// instances.
func (computeService Service) QueryInstances(queryParameters *QueryParameters) ([]Response, error) {
	computeContainer := computeResponse{}
	err := computeService.queryInstances(false /*includeDetails*/, &computeContainer, queryParameters)
	if err != nil {
		return nil, err
	}

	return computeContainer.Instances, nil
}

// QueryImagesDetail will issue a get request with the specified QueryParameters to retrieve the list of
// images with additional details.
func (computeService Service) QueryInstancesDetail(queryParameters *QueryParameters) ([]DetailResponse, error) {
	computeDetailContainer := computeDetailResponse{}
	err := computeService.queryInstances(true /*includeDetails*/, &computeDetailContainer, queryParameters)
	if err != nil {
		return nil, err
	}

	return computeDetailContainer.Instances, nil
}

func (computeService Service) queryInstances(includeDetails bool, computeResponseContainer interface{}, queryParameters *QueryParameters) error {
	urlPostFix := "/servers"
	if includeDetails {
		urlPostFix = urlPostFix + "/detail"
	}

	reqURL, err := buildQueryURL(computeService, queryParameters, urlPostFix)
	//fmt.Printf("***********************", reqURL)
	if err != nil {
		return err
	}

	var headers http.Header = http.Header{}
	headers.Set("Accept", "application/json")
	resp, err := computeService.Session.Get(reqURL.String(), nil, &headers)
	if err != nil {
		return err
	}

	err = util.CheckHTTPResponseStatusCode(resp)
	if err != nil {
		return err
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("aaa")
	}
	if err = json.Unmarshal(rbody, &computeResponseContainer); err != nil {
		return err
	}
	return nil
}

func buildQueryURL(computeService Service, queryParameters *QueryParameters, computePartialURL string) (*url.URL, error) {
	reqURL, err := url.Parse(computeService.URL)
	if err != nil {
		return nil, err
	}

	if queryParameters != nil {
		values := url.Values{}
		if queryParameters.instance_name != "" {
			values.Set("instance_name", queryParameters.instance_name)
		}
		if len(values) > 0 {
			reqURL.RawQuery = values.Encode()
		}
	}
	reqURL.Path += computePartialURL

	return reqURL, nil
}

type computeDetailResponse struct {
	Instances []DetailResponse `json:"servers"`
}

type computeResponse struct {
	Instances []Response `json:"servers"`
}