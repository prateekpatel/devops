package OpenStack_API_Function

import (
	"fmt"

	"net/http"
	"io/ioutil"

	"gclassec/hos/GetAuthToken"
)



type Response struct {
	Counter_name		string	`json:"counter_name"`
	User_id			string	`json:"user_id"`
	Resource_id		string	`json:"resource_id"`
	Timestamp		string	`json:"timestamp"`
	Recorded_at		string	`json:"recorded_at"`
	Message_id		string	`json:"message_id"`
	Source			string	`json:"source"`
	Counter_unit		string	`json:"counter_unit"`
	Counter_volume		string	`json:"counter_volume"`
	Project_id		string 	`json:"project_id"`
	Resource_metadata	string	`json:"resource_metadata"`
	Counter_type		string	`json:"counter_type"`

}


func GetCeilometerDetail() string{

	fmt.Println("This to get Nothing")
	var auth = GetAuthToken.GetOpenStackAuthToken()
	fmt.Println("Auth Token in Compute.go:=====\n", auth)

	var reqURL string =  "http://140.140.140.41:8777/v2/meters"
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("x-auth-token", auth)
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	fmt.Println("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, _ := ioutil.ReadAll(res.Body)

	fmt.Print("respBody:==\n",respBody)
	respBodyInString:= string(respBody)
//	fmt.Println("\nrespBodyInString:==\n",respBodyInString)

	return respBodyInString
}
