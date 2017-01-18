package OpenStack_API_Function

import (
	"fmt"

	"net/http"
	"io/ioutil"
	"gclassec/hos/GetAuthToken"
)

func GetCpuUtilDetails(id string) string {

	fmt.Println("This to get Nothing")
	var auth = GetAuthToken.GetOpenStackAuthToken()
	fmt.Println("Auth Token in Compute.go:=====\n", auth)

	var reqURL string =  "http://140.140.140.41:8777/v2/meters/cpu_util/statistics?q.field=resource_id&q.field=timestamp&q.op=eq&q.op=gt&q.type=&q.type=&q.value="+id+"&q.value=2017-01-18T08%3A55%3A00"


	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("x-auth-token", auth)
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	fmt.Println("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, _ := ioutil.ReadAll(res.Body)

	fmt.Print("respBody:==\n",respBody)
	respBodyInString:= string(respBody)
	//fmt.Println("\nrespBodyInString:==\n",respBodyInString)
	return respBodyInString
}
