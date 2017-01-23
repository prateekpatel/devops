//package Testing_of_Controllers


package main
import (
    	"fmt"
	"os"
	"encoding/json"
	"strings"
	"github.com/verdverm/frisby"
	"runtime"
	//"github.com/mozillazg/request"
	"github.com/mozillazg/request"
	//"reflect"

	"github.com/bitly/go-simplejson"
)

type Configration struct {
	Protocol string
	Host     string
	PortVal  string
	//username string
	//password string
	//Resp *request.Response
}
type Request struct {
	Resp *request.Response
}
type Frisby struct {

}

type Post int



func main() {
	//file, _ :=  os.Open("C:\\Git\\goclassec\\src\\gclassec\\conf\\controllertesting.json")
	filename := "tests/testing_controller.go"
       _, filePath, _, _ := runtime.Caller(0)
       fmt.Println("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/controllertesting.json", 1))
       fmt.Println("ABSPATH:==",ConfigFilePath)
	file, _ := os.Open(ConfigFilePath)

	decoder := json.NewDecoder(file)
	configration := Configration{}
	err := decoder.Decode(&configration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configration.PortVal)
	Protocol := configration.Protocol
	Host := configration.Host
	PortVal := configration.PortVal
	//fmt.Println(Protocol)
	//fmt.Println(Host)
	//fmt.Println(PortVal)
var b []string = []string{Protocol,"://",Host,":",PortVal}

var URL string = (strings.Join(b,""))
	println(URL)
	fmt.Println("Frisby")
	//connect to the server..................
	frisby.Create("Connect to the server").
		Get(URL+"/dbaas/list").
		Send().
		ExpectStatus(200)
		//ExpectContent(" list")
	frisby.Create("Test to the server (Should fails in casesentive and post and delete aplhabatic character )").
		Get(URL +"/dbaas/listw").
		Send().
		ExpectStatus(404)
	//dbaas of the id...............
	frisby.Create("display the dbaas of the  id").
		Get(URL +"/dbaas/list/dev01-a-tky-customerorderpf").
		Send().
		ExpectStatus(200)

	frisby.Create("display the dbaas id(server falis at numeric character)").
		Get(URL +"/dbaas1/list/dev01-a-tky-customerorderpf").
		Send().
		ExpectStatus(404)
	//dbaas cpu utalization.............................
	frisby.Create("display the cpu utalization server").
		Get(URL +"/dbaas/get?CPUUtilization_max=5&DatabaseConnections_max=0").
		Send().
		ExpectStatus(200)

	frisby.Create("display the cpu utalization server (should falis alphanumeric charecetr )").
		Get(URL +"/dbaas1/get?CPUUtilization_max=acb2&DatabaseConnections_max=string").
		Send().
		ExpectStatus(404)
	//dbass price detalis.............................
	frisby.Create("display the price of dbass").
		Get(URL +"/dbaas/pricing").
		Send().
		ExpectStatus(200)

	frisby.Create("display the price of dabbas (should falis in  post)").
		Post(URL +"/dbaas/pricing").
		Send().
		ExpectStatus(404)
		//ExpectContent("A string which won't be found")
	//openstack details..................................
	frisby.Create("display the openstack details").
		Get(URL +"/dbaas/openstackDetail").
		Send().
		ExpectStatus(200)

	frisby.Create("display the openstack details (should fails in case senstive)").
		Get(URL +"/dbaas/openstackdetail").
		Send().
		ExpectStatus(404)
	//azure static detalis.........................
	frisby.Create(" Display the list of  static details of Azure instances").
		Get(URL +"/dbaas/azureDetail").
		Send().
		ExpectStatus(200)
	frisby.Create("Display the list of  static details of Azure instances (should fails in case senstive)").
		Get(URL +"/dbaas/AzureDetail").
		Send().
		ExpectStatus(404)
	//azure dynamic details..................
	frisby.Create("Display the list of  dynamic details of Azure instances ").
		Get(URL +"/dbaas/azureDetail/percentCPU/test/testGo").
		Send().
		ExpectStatus(200)
	frisby.Create("Display the list of  dynamic details of Azure instances (should fails in alphabatic character)").
		Get(URL +"/dbaas/azureDetails/percentCPU/test/testGo").
		Send().
		ExpectStatus(404)
	//vcenterDetail...................
	frisby.Create("Display the list of  dynamic vcenter details ").
		Get(URL +"/dbaas/vcenterDetail").
		Send().
		ExpectStatus(200)
	frisby.Create("Display the list of  dynamic details of vcenterdetails (should fails in alphabatic character)").
		Get(URL +"/dbaas/vcenterDetails").
		Send().
		ExpectStatus(404)
	//providers post ...................................

	//regi := "{username: swathi, password: atmecs@123}"
	//frisby.Create("Test POST").
	//	Post(URL+"/providers").
	//	Send().
	//	SetJson(regi).
	//	ExpectStatus(404).
	//AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
	//	})
	//frisby.Create("Test POST").
	//	Post("http://110.110.110.233:9009/Providers1@123we").
	//	Send().
	//	ExpectStatus(404)
	//frisby.Create("Test POST").
	//	Post("http://110.110.110.233:9009//providers").
	//	Send().
	//	ExpectStatus(301)
	//openstack providers details.....................
		Regi := "{username: swathi, password: atmecs@123}"
		frisby.Create("Testing the opestack details..").
			Post(URL+"/providers/openstack").
			SetJson(Regi).
			Send().
			ExpectStatus(200).
			AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			})
		//values := "{username: swathi, password: atmecs@123}"
		frisby.Create("Testing the opestack details..(should fails in the alphatic").
			Post(URL+"/providers/openstacks").
			//SetJson(values).
			Send().
			ExpectStatus(404)
		frisby.Create("Testing the opestack details..(should fails in(/) ").
			Post(URL+"/providers//openstacks").
			//SetJson(values).
			Send().
			ExpectStatus(301)

			//AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			//})
	//Azure detalis...........................
		regi := "{username: werty, password: atmecs@123}"
		frisby.Create("Display the providers details of azure ").
			Post(URL +"/providers/azure").
			SetJson(regi).
			Send().
	 		ExpectStatus(200).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			})
		frisby.Create("Display the providers details of azure(should fails in alphabatic character)").
			Post(URL +"/providers/azures").
			Send().
			ExpectStatus(404)

	frisby.Global.PrintReport()
}
