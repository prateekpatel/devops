//package Testing_of_Controllers


package main
import (
    "fmt"
    "github.com/verdverm/frisby"
)

func main() {
	fmt.Println("Frisby")
	frisby.Create("Connect to the localhost").
		Get("http://localhost:9009/dbaas/list").
		Send().
		ExpectStatus(200)
		//ExpectContent(" list")

	frisby.Create("Test to the local host (which fails)").
		Get("http://localhost:9009/dbaas/list").
		Send().
		ExpectStatus(400)
	frisby.Create("display the id").
		Get("http://localhost:9009/dbaas/list/dev01-a-tky-customerorderpf").
		Send().
		ExpectStatus(200)

	frisby.Create("display the id(Wich fails)").
		Get("http://localhost:9009/dbaas/list/dev01-a-tky-customerorderpf").
		Send().
		ExpectStatus(400)
	frisby.Create("display the cpu utalization").
		Get("http://localhost:9009/dbaas/get?CPUUtilization_max=5&DatabaseConnections_max=0").
		Send().
		ExpectStatus(200)

	frisby.Create("display the cpu utalization (Wich fails)").
		Get("http://localhost:9009/dbaas/get?CPUUtilization_max=5&DatabaseConnections_max=0").
		Send().
		ExpectStatus(400)
	frisby.Create("display the sprint price").
		Get("http://localhost:9009/dbaas/pricing").
		Send().
		ExpectStatus(200)

	frisby.Create("display the sprint price (Wich fails)").
		Get("http://localhost:9009/dbaas/pricing").
		Send().
		ExpectStatus(400)
		//ExpectContent("A string which won't be found")

	frisby.Global.PrintReport()
}
