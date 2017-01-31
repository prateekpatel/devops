package main

import ("github.com/verdverm/frisby"
	"strings"
	"os"
	"encoding/json"
	"runtime"
	"fmt"
	"github.com/bitly/go-simplejson"
)
type Configration struct {
	Protocol string		//`json:"protocol"`
	Hostname     string	//`json:"hostname"`
	PortValue  string	//`json:"portValue"`
}
func main() {
	//file, _ :=  os.Open("C:\\Git\\goclassec\\src\\gclassec\\conf\\controllertesting.json")
	var filename string = "tests/test_controller_hos.go"
       _, filePath, _, _ := runtime.Caller(0)
       fmt.Println("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/hos_controller_test.json", 1))
       fmt.Println("ABSPATH:==",ConfigFilePath)
	file, _ := os.Open(ConfigFilePath)

	decoder := json.NewDecoder(file)
	tempconfig := Configration{}
	fmt.Println(tempconfig)
	err := decoder.Decode(&tempconfig)
	if err != nil {
		//fmt.Println("eror")
		fmt.Println("error:", err)
	}
	//fmt.Println("tempconfig:=========")
	//fmt.Printf("%+v",tempconfig)
	Protocol := tempconfig.Protocol
	Host := tempconfig.Hostname
	PortVal := tempconfig.PortValue
var b []string = []string{Protocol,"://",Host,":",PortVal}

var LINK string = (strings.Join(b,""))
	println(LINK)
	fmt.Println("Frisby")
	//connect to the compute server of hos..................
	frisby.Create("Connect to the hos server").
		Get(LINK+"/hos/computedetails").
		Send().
		ExpectStatus(200).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			})

	frisby.Create("Test to the compute server (Should fails in casesentive and post and delete aplhabatic character )").
		Get(LINK +"/hos/computedetailss").
		Send().
		ExpectStatus(404)
	//falvor details..........................................
	frisby.Create("Display the hos flavors details server").
		Get(LINK+"/hos/flavorsdetails").
		Send().
		ExpectStatus(200).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			})
	frisby.Create("Test to the  hos flavorsdetails server (Should fails in  aplhabatic character )").
		Get(LINK +"/hos/flavorsdetailss").
		Send().
		ExpectStatus(404)

	//hos/cpu_utilization[/hos/cpu_utilization/{id}].......................
	frisby.Create("Display the hos cpu_utilization").
		Get(LINK+"/hos/cpu_utilization/8d874af2-5cb8-49a4-90f6-378a0a2633bc").
		Send().
		ExpectStatus(200).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			})
	frisby.Create("Test to the hos cpu_utilization (Should fails in numeric character )").
		Get(LINK +"/hos/cpu_utilizations/8d874af2-5cb8-49a4-90f6-378a0a2634bc").
		Send().
		ExpectStatus(404)
	///hos/index.................

	frisby.Create("Display the   hos index ").
		Get(LINK+"/hos/index").
		Send().
		ExpectStatus(200).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			})
	frisby.Create("Test to the  hos index (Should fails in casesentive  )").
		Get(LINK +"/hos/Index").
		Send().
		ExpectStatus(404)
	//AWS dbaas list.............................................................................................
	//fmt.Println("AWS DETAILS......")
	frisby.Create("Connect to the dbaas list").
		Get(LINK+"/dbaas/list").
		Send().
		ExpectStatus(200).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			})
	frisby.Create("Test to the  dbaas list (Should fails  aplhabatic character )").
		Get(LINK +"/dbaass/listS").
		Send().
		ExpectStatus(404)
	//dbaas list id [/dbaas/list/{id}].........................
	frisby.Create("Connect to the dbaas list id").
		Get(LINK+"/dbaas/list/dev01-a-tky-customerorderpf").
		Send().
		ExpectStatus(200).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			})
	frisby.Create("Test to the  dbaas list id (Should fails in numeric and aplhabatic character )").
		Get(LINK +"/dbaas/lists/dev012-a-tkys-customerorderpfs").
		Send().
		ExpectStatus(404)
	//dbass get......................
	frisby.Create("Connect to the dbaas get").
		Get(LINK+"/dbaas/get").
		Send().
		ExpectStatus(200).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			})
	frisby.Create("Test to the  dbaas get (Should fails in casesentive and post and delete aplhabatic character )").
		Get(LINK +"/dbaas/gets").
		Send().
		ExpectStatus(404)
	//dbass pricing ....................
	frisby.Create("Connect to the dbaas pricing").
		Get(LINK+"/dbaas/pricing").
		Send().
		ExpectStatus(200).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			})
	frisby.Create("Test to the  dbaas pricing (Should fails in alphanumeric character )").
		Get(LINK +"/dbaas/pricingS1").
		Send().
		ExpectStatus(404)
	//dbass openstack details...........................
	frisby.Create("Connect to the dbaas openstack list").
		Get(LINK+"/dbaas/openstackDetail").
		Send().
		ExpectStatus(200).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			})
	frisby.Create("Test to the  openstack dbaas list (Should fails in casesentive and post and delete aplhabatic character )").
		Get(LINK +"/dbaas/openstackDetailS").
		Send().
		ExpectStatus(404)
	//dbass azure details ....................
	frisby.Create("Connect to the dbaas azure details").
		Get(LINK+"/dbaas/azureDetail").
		Send().
		ExpectStatus(200).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			})
	frisby.Create("Test to the  dbaas azure details (Should fails in casesentive and post and delete aplhabatic character )").
		Get(LINK +"/dbaas/azureDetailS").
		Send().
		ExpectStatus(404)
	//azure cpu utalization for specifed instances[/dbaas/azureDetail/percentCPU/{resourceGroup}/{name}]..........................
	frisby.Create("Connect to the dbaas list").
		Get(LINK+"/dbaas/azureDetail/percentCPU/test/testGo").
		Send().
		ExpectStatus(200).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			})
	frisby.Create("Test to the  dbaas list (Should fails in casesentive and post and delete aplhabatic character )").
		Get(LINK +"/dbaass/azureDetail/percentCPU/test/testGo").
		Send().
		ExpectStatus(404)
	frisby.Global.PrintReport()
}
//
//
