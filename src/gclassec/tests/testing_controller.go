//package Testing_of_Controllers


package main
import (
    	"fmt"
	"os"
	"encoding/json"
	"strings"
	"github.com/verdverm/frisby"
	"runtime"
)

type Configration struct {
	Protocol string
	Host     string
	PortVal  string
}
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

	frisby.Global.PrintReport()
}
