package confcontroller

import (
	"net/http"
	"fmt"
	"os"
	"encoding/json"
	"bufio"
	"runtime"
	"strings"
)

var redirectTarget string

 const indexPage = `
  <h1>Select Provider</h1>
  <form method="post" action="/providers">
      <label for="provider">Provider</label>
      <input type="text" id="provider" name="provider">
      </br></br>
      <button type="submit">Select</button>
  </form>`

 const osPage = `
  <h1>Openstack Credentials</h1>
  <form method="post" action="/providers/openstack">
      <label for="host">Host</label>
      <input type="text" id="host" name="host"></br></br>

      <label for="username">Username</label>
      <input type="text" id="username" name="username"></br></br>

      <label for="password">Password</label>
      <input type="text" id="password" name="password"></br></br>

      <label for="projectid">ProjectID</label>
      <input type="text" id="projectid" name="projectid"></br></br>

      <label for="projectname">ProjectName</label>
      <input type="text" id="projectname" name="projectname"></br></br>

      <label for="container">Container</label>
      <input type="text" id="container" name="container"></br></br>

      <label for="imageregion">ImageRegion</label>
      <input type="text" id="imageregion" name="imageregion"></br></br>

      <label for="controller">Controller</label>
      <input type="text" id="controller" name="controller"></br></br>

      <button type="submit">Submit</button>
  </form>`

const azurePage = `
  <h1>Azure Credentials</h1>
  <form method="post" action="/providers/azure">
      <label for="clientid">Client ID</label>
      <input type="text" id="clientid" name="clientid"></br></br>

      <label for="clientsecret">Client Secret</label>
      <input type="text" id="clientsecret" name="clientsecret"></br></br>

      <label for="subscriptionid">Subscription ID</label>
      <input type="text" id="subscriptionid" name="subscriptionid"></br></br>

      <label for="tenantid">Tenant ID</label>
      <input type="text" id="tenantid" name="tenantid"></br></br>

      <button type="submit">Submit</button>
  </form>`

type (
    // UserController represents the controller for operating on the User resource
    UserController struct{}
)

func NewUserController() *UserController {
    return &UserController{}
}

func (uc UserController) SelectProvider(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexPage)
}

func (uc UserController) OpenstackCreds(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, osPage)
}

func (uc UserController) AzureCreds(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, azurePage)
}

func (uc UserController) ProviderHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, indexPage)
	fmt.Println("--------In Provider Handler--------")
	provider := r.FormValue("provider")
	fmt.Println("Provider : ")
	fmt.Println(provider)

	if provider == "openstack"{
		//setSession(provider, w)
		redirectTarget = "/selectedOs"
	}

	if provider == "azure"{
		//setSession(provider, w)
		redirectTarget = "/selectedAzure"
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func (uc UserController) ProviderOpenstack(w http.ResponseWriter, r *http.Request) {
	//host := r.FormValue("host")
	c := map[string]string{
		"host":       r.FormValue("host"),
		"username":   r.FormValue("username"),
		"password": r.FormValue("password"),
		"projectid": r.FormValue("projectid"),
		"projectname": r.FormValue("projectname"),
		"container": r.FormValue("container"),
		"imageregion": r.FormValue("imageregion"),
		"controller": r.FormValue("controller")}

  	outputjson,_:=json.Marshal(c)

	filename := "controllers/confcontroller/userconf.go"
       _, filePath, _, _ := runtime.Caller(0)
       fmt.Println("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/computeVM.json", 1))
       fmt.Println("ABSPATH:==",ConfigFilePath)
	f, err := os.Create(ConfigFilePath)

	//f, err := os.OpenFile("C:/goclassec/src/gclassec/conf/dependencies.env", os.O_APPEND | os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	/*for _, line := range c {
		if _, err = f.WriteString(line); err != nil {
			panic(err)
		}
	}*/

	//define the 'string writer'
  	filewriter:=bufio.NewWriter(f)

  	//write the JSON string. First we need to convert the outputjson to string, and then write it
  	filewriter.WriteString(string(outputjson))
  	filewriter.Flush()
}

func (uc UserController) ProviderAzure(w http.ResponseWriter, r *http.Request) {
	c := map[string]string{
		"clientid": r.FormValue("clientid"),
		"clientsecret": r.FormValue("clientsecret"),
		"subscriptionid": r.FormValue("subscriptionid"),
		"tenantid": r.FormValue("tenantid")}

	outputjson,_:=json.Marshal(c)

	filename := "controllers/confcontroller/userconf.go"
       _, filePath, _, _ := runtime.Caller(0)
       fmt.Println("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/azurecred.json", 1))
       fmt.Println("ABSPATH:==",ConfigFilePath)
	f, err := os.Create(ConfigFilePath)

	//f, err := os.OpenFile("C:/goclassec/src/gclassec/conf/dependencies.env", os.O_APPEND | os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//define the 'string writer'
  	filewriter:=bufio.NewWriter(f)

  	//write the JSON string. First we need to convert the outputjson to string, and then write it
  	filewriter.WriteString(string(outputjson))
  	filewriter.Flush()
}
