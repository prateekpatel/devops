package flavor
import (
	"fmt"
	"time"
	"git.openstack.org/openstack/golang-client.git/openstack"
	"os"
	"encoding/json"
	"net/http"
)
type Configuration struct {
    Host	string
    Username	string
    Password	string
    ProjectID	string
    ProjectName	string
    Container	string
    ImageRegion	string
}

func Flavor() []DetailResponse{
	//config := getConfig()
	file, _ := os.Open("C:/Project/Go/src/git.openstack.org/openstack/golang-client.git/examples/config.json")
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
	flavorService := Service{
		Session: *sess,
		Client:  *http.DefaultClient,
		URL:     url, // We're forcing Volume v2 for now
	}
	flavorDetails, err := flavorService.FlavorsDetail()
	fmt.Println(flavorDetails,"00000000000000000000000000")
	if err != nil {
		panicString := fmt.Sprint("Cannot access Compute:", err)
		panic(panicString)
	}
	fmt.Println("computedetails printing..")
	fmt.Println(flavorDetails)
	var flavorIDs = make([]string, 0)
	for _, element := range flavorDetails {
		flavorIDs = append(flavorIDs, element.FlavorID)
	}
	fmt.Println(flavorIDs)
	if len(flavorIDs) == 0 {
		panicString := fmt.Sprint("No instances found, check to make sure access is correct")
		panic(panicString)
	}
	return flavorDetails
}