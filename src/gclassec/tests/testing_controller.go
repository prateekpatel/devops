//package Testing_of_Controllers


package main
import (
    "fmt"
    "github.com/verdverm/frisby"
)

func main() {
	fmt.Println("Frisby")
	frisby.Create("Connect to the server").
		Get("http://110.110.110.233:9009/dbaas/list").
		Send().
		ExpectStatus(200)
		//ExpectContent(" list")

	frisby.Create("Test to the server (Should fails in casesentive and post and delete aplhabatic character )").
		Get("http://110.110.110.233:9009/dbaas/listw").
		Send().
		ExpectStatus(404)
	frisby.Create("display the dbaas of the  id").
		Get("http://110.110.110.233:9009/dbaas/list/dev01-a-tky-customerorderpf").
		Send().
		ExpectStatus(200)

	frisby.Create("display the dbaas id(server falis at numeric character)").
		Get("http://110.110.110.233:9009/dbaas1/list/dev01-a-tky-customerorderpf").
		Send().
		ExpectStatus(404)
	frisby.Create("display the cpu utalization server").
		Get("http://110.110.110.233:9009/dbaas/get?CPUUtilization_max=5&DatabaseConnections_max=0").
		Send().
		ExpectStatus(200)

	frisby.Create("display the cpu utalization server (should falis alphanumeric charecetr )").
		Get("http://110.110.110.233:9009/dbaas1/get?CPUUtilization_max=acb2&DatabaseConnections_max=string").
		Send().
		ExpectStatus(404)
	frisby.Create("display the price of dbass").
		Get("http://110.110.110.233:9009/dbaas/pricing").
		Send().
		ExpectStatus(200)

	frisby.Create("display the price of dabbas (should falis in  post)").
		Post("http://110.110.110.233:9009/dbaas/pricing").
		Send().
		ExpectStatus(404)
		//ExpectContent("A string which won't be found")
	frisby.Create("display the openstack details").
		Get("http://110.110.110.233:9009/dbaas/openstackDetail").
		Send().
		ExpectStatus(200)

	frisby.Create("display the openstack details (should fails in case senstive)").
		Get("http://110.110.110.233:9009/dbaas/openstackdetail").
		Send().
		ExpectStatus(404)
	frisby.Global.PrintReport()
}
