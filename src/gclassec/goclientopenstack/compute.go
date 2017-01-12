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

package goclientcompute
import (
	"fmt"
	"time"
	"git.openstack.org/openstack/golang-client.git/openstack"
	"os"
	"encoding/json"
	"net/http"
	"gclassec/goclientopenstack/flavor"
	"gclassec/goclientopenstack/compute"
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

func Compute() []compute.DetailResponse {
	//config := getConfig()
	dir, _ := os.Getwd()
	file, _ := os.Open(dir + "/src/gclassec/conf/computeVM.json")
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
		computeIDs = append(computeIDs, element.ID)

	}
	fmt.Println(computeIDs)
	if len(computeIDs) == 0 {
		panicString := fmt.Sprint("No instances found, check to make sure access is correct")
		panic(panicString)
	}
	return computeDetails
}

func FinalCompute() []compute.DetailResponse {
	var flvObj []flavor.DetailResponse
	flvObj = flavor.Flavor()
	fmt.Println("&**********Showing FLVOBJ&************")
	fmt.Println(flvObj)
	fmt.Println("*********************")
	fmt.Println("flvObj.FlavorID::", flvObj[1].FlavorID)

	var obj []compute.DetailResponse
	obj = Compute()
	fmt.Println("77778888899999", obj[1].Flavor.FlavorID)
	for i:=0; i<len(obj); i++{
		tempFID :=obj[i].Flavor.FlavorID
		for j:=0; j<len(flvObj); j++{
			if tempFID==flvObj[j].FlavorID{
				obj[i].Flavor.Name=flvObj[j].Name
				obj[i].Flavor.Disk=flvObj[j].Disk
				obj[i].Flavor.RAM=flvObj[j].RAM
				obj[i].Flavor.VCPU=flvObj[j].VCPU
			}
		}
	}
	out, err := json.Marshal(obj)
	if err != nil {
        	panic (err)
    	}
	fmt.Println("Out Sritng")
	fmt.Println(string(out))
	temp := string(out)
	temp1 := strings.TrimPrefix(temp, "[{")
	tempstr:= strings.TrimSuffix(temp1, "}]")
	tempVar := strings.Split(tempstr,"},{")
	fmt.Println("-----------TempVar----------")
	for i:=0; i<len(tempVar);i++{
		fmt.Println(tempVar[i])
	}
	for i:=0; i<len(tempVar);i++{
		nevVar := string(tempVar[i])
		tempVar1 := strings.Split(nevVar,",")
		fmt.Println("-----------TempVar1.----------",i)
		for j:=0; j<len(tempVar1);j++{
			fmt.Println(tempVar1[j])
		}
	}
	return obj
}