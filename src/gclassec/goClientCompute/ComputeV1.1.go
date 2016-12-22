// +build !unit

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

package main
import (
	"fmt"
	"time"
	"os"
	"encoding/json"
	"net/http"
	"gclassec/goClientCompute/compute"
	"gclassec/goClientCompute/openstack"
	"strings"
)
type Configuration struct {
    Host    string
    Username   string
    Password   string
    ProjectID   string
    ProjectName   string
    Container   string
    ImageRegion string
}

func main() {
	//config := getConfig()
	file, _ := os.Open("C:/Chaitrali/Git/goclassec/computeVM.json")
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}

	// Authenticate with a username, password, tenant id.
	creds := openstack.AuthOpts{
		AuthUrl:     config.Host,
		ProjectName: config.ProjectName,
		Username:    config.Username,
		Password:    config.Password,
	}
	auth, err := openstack.DoAuthRequest(creds)

	//fmt.Printf("**** Thia is an AUTH Token ***** ::: ", auth)
	if err != nil {
		panicString := fmt.Sprint("There was an error authenticating:", err)
		panic(panicString)
	}
	if !auth.GetExpiration().After(time.Now()) {
		panic("There was an error. The auth token has an invalid expiration.")
	}
	fmt.Println(auth)
	// Find the endpoint for the Nova Compute service.

	url, err := auth.GetEndpoint("compute", "")

	if url == "" || err != nil {
		panic("EndPoint Not Found.")
		panic(err)
	}
	// Make a new client with these creds
	sess, err := openstack.NewSession(nil, auth, nil)
	if err != nil {
		panicString := fmt.Sprint("Error creating new Session:", err)
		panic(panicString)
	}
	url = strings.Replace(url,"controller",config.Container,1)
	fmt.Println(url)
	computeService := compute.Service{
		Session: *sess,
		Client:  *http.DefaultClient,
		URL:     url, // We're forcing Volume v2 for now
	}
	computeDetails, err := computeService.InstancesDetail()
	if err != nil {
		panicString := fmt.Sprint("Cannot access Compute:", err)
		panic(panicString)
	}
	fmt.Println("computedetails printing..")
	fmt.Println(computeDetails)
	var computeIDs = make([]string, 0)
	for _, element := range computeDetails {

		computeIDs = append(computeIDs, element.Name)
	}

	if len(computeIDs) == 0 {
		panicString := fmt.Sprint("No instances found, check to make sure access is correct")
		panic(panicString)
	}
}