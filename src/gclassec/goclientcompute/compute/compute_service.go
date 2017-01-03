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
	"git.openstack.org/openstack/golang-client.git/openstack"
	"net/http"
	"git.openstack.org/openstack/golang-client.git/util"
	"net/url"
	"fmt"
	"git.openstack.org/openstack/golang-client.git/flavor"
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
	DiskConfig			string `json:"diskConfig"`
	Host              		string `json:"host"`
	Instance_name            	string `json:"instance_name"`
	Power_state   			int64  `json:"power_state"`
	vm_state	      		string `json:"vm_state"`
	Flavor				string `json:"flavor"`
	HostId		      		string `json:"hostId"`
	Image	            		string `json:"image"`
	Key_name            		string `json:"key_name"`
	Name	       			string `json:"name"`
	Volumes_attached            	string `json:"volumes_attached"`
	Security_groups	            	string `json:"security_groups"`
	Status            		string `json:"status"`
	Tenant_id   			string `json:"tenant_id"`
}
// DetailResponse is a structure for all properties of
// an instance for a detailed query
type DetailResponse struct {
	Name	       			string `json:"name"`
	ID	       			string `json:"id"`
	Status            		string `json:"status"`
	Availability_zone      		string `json:"OS-EXT-AZ:availability_zone"`
	Created				string `json:"created"`
	Flavor				flavor.DetailResponse `json:"flavor"`
	Addresses			address `json:"addresses"`
	Security_groups 		[]security_groups `json:"security_groups"`
	Key_name           		string `json:"key_name"`
	Image           		Image  `json:"image"`
	Tenant_id   			string `json:"tenant_id"`
	Updated				string `json:"updated"`
	User_id		      		string `json:"user_id"`
	HostId				string `json:"hostId"`
	Task_state   			string `json:"OS-EXT-STS:task_state"`
	Vm_state			string `json:"OS-EXT-STS:vm_state"`
	Launched_at		      	string `json:"OS-SRV-USG:launched_at"`
	Volumes_attached   		[]Volume `json:"os-extended-volumes:volumes_attached"`
	Progress            		int64  `json:"progress"`
	IPV4				string	`json:"accessIPv4"`
	IPV6				string	`json:"accessIPv6"`
	Power_State			int64	`json:"OS-EXT-STS:power_state"`
	ConfigDrive			string	`json:"config_drive"`
	DiskConfig			string	`json:"OS-DCF:diskConfig"`
}
type Volume struct{
	Vol	string `json:"os-extended-volumes:volumes_attached"`
}
type Image struct{
	ID	string `json:"id"`
}
/*type Flavor struct{
	ID	string `json:"id"`
}*/
type security_groups struct{
	Name	string `json:"name"`
}
type address struct{
	Mac_addr	string	`json:"OS-EXT-IPS-MAC:mac_addr"`
	Version		string	`json:"version"`
	Addr		string	`json:"addr"`
	Type		string	`json:"OS-EXT-IPS:type"`
}
// QueryParameters is a structure that
// contains the filter, sort, and paging parameters for
// an instance or computedetail query.
type QueryParameters struct {
	Id			int64
	Name			string
	InstanceID		string
	Status 			string
	AvailabilityZone	string
	CreationTime		string
	Flavor			[]flavor.DetailResponse
	FlavorID		int64
	RAM			string
	VCPU 			string
	Storage 		string
	IPAddress		string
	SecurityGroup		string
	KeyPairName		string
	ImageName		string
	Volumes			string
	InsertionDate		string
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

	return computeContainer.Servers, nil
}

// QueryImagesDetail will issue a get request with the specified QueryParameters to retrieve the list of
// images with additional details.
func (computeService Service) QueryInstancesDetail(queryParameters *QueryParameters) ([]DetailResponse, error) {
	computeDetailContainer := computeDetailResponse{}
	err := computeService.queryInstances(true /*includeDetails*/, &computeDetailContainer, queryParameters)
	if err != nil {
		return nil, err
	}

	return computeDetailContainer.Servers, nil
}

func (computeService Service) queryInstances(includeDetails bool, computeResponseContainer interface{}, queryParameters *QueryParameters) error {
	urlPostFix := "/servers"
	if includeDetails {
		urlPostFix = urlPostFix + "/detail"
	}
	fmt.Println(urlPostFix)
	reqURL, err := buildQueryURL(computeService, queryParameters, urlPostFix)
	if err != nil {
		return err
	}

	fmt.Println("cghvgjvgjgjvgj",reqURL)

	var headers http.Header = http.Header{}
	headers.Set("Accept", "application/json")
	resp, err := computeService.Session.Get(reqURL.String(), nil, &headers)
	fmt.Println("********************",headers)
	if err != nil {
		return err
	}

	err = util.CheckHTTPResponseStatusCode(resp)
	if err != nil {
		return err
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response body printng.")
	fmt.Println(rbody)
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
		if queryParameters.Name != "" {
			values.Set("instance_name", queryParameters.Name)
		}
		if len(values) > 0 {
			reqURL.RawQuery = values.Encode()
		}
	}
	reqURL.Path += computePartialURL

	return reqURL, nil
}

type computeDetailResponse struct {
	Servers []DetailResponse `json:"servers"`
}

type computeResponse struct {
	Servers []Response `json:"servers"`
}



type FinalDetailResponse struct {
	Name	       			string `json:"name"`
	ID	       			string `json:"id"`
	Status            		string `json:"status"`
	Availability_zone      		string `json:"OS-EXT-AZ:availability_zone"`
	Created				string `json:"created"`
	Flavor				flavor.DetailResponse `json:"flavor"`
	Addresses			address `json:"addresses"`
	Security_groups 		[]security_groups `json:"security_groups"`
	Key_name           		string `json:"key_name"`
	Image           		Image  `json:"image"`
	Tenant_id   			string `json:"tenant_id"`
	Updated				string `json:"updated"`
	User_id		      		string `json:"user_id"`
	HostId				string `json:"hostId"`
	Task_state   			string `json:"OS-EXT-STS:task_state"`
	Vm_state			string `json:"OS-EXT-STS:vm_state"`
	Launched_at		      	string `json:"OS-SRV-USG:launched_at"`
	Volumes_attached   		[]Volume `json:"os-extended-volumes:volumes_attached"`
	Progress            		int64  `json:"progress"`
	IPV4				string	`json:"accessIPv4"`
	IPV6				string	`json:"accessIPv6"`
	Power_State			int64	`json:"OS-EXT-STS:power_state"`
	ConfigDrive			string	`json:"config_drive"`
	DiskConfig			string	`json:"OS-DCF:diskConfig"`
}