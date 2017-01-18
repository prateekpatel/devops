package OpenStack_API_Function

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"gclassec/hos/GetAuthToken"
)

func Flavors() string {

	//fmt.Println("This to get Nothing")
	var auth = GetAuthToken.GetOpenStackAuthToken()
	fmt.Println("Auth Token in Flavors.go:=====\n", auth)

	var reqURL string = "http://140.140.140.41:8774/v2.1/13a46539e3a146f68fc5b105655403fa/flavors/detail"
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("x-auth-token", auth)
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	fmt.Println("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, _ := ioutil.ReadAll(res.Body)

	//fmt.Print("respBody:==\n",respBody)
	respBodyInString:= string(respBody)
	fmt.Println("\nrespBodyInString:==\n",respBodyInString)
	return respBodyInString
}

