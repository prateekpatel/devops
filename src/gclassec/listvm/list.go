package main


import (
	"os"
	"log"

	"github.com/Azure/go-autorest/autorest/azure"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/arm/compute"
	"github.com/Azure/azure-sdk-for-go/arm/examples/helpers"


_ "github.com/go-sql-driver/mysql"

	"strings"
	"github.com/jinzhu/gorm"
	"gclassec/readazureconf"

	"gclassec/azurestruct"
)

type ls struct {

}
type (
    // UserController represents the controller for operating on the User resource
    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}


var dbcredentials = readazureconf.Configurtion()
var dbtype string = dbcredentials.Dbtype
var dbname  string = dbcredentials.Dbname
var dbusername string = dbcredentials.Dbusername
var dbpassword string = dbcredentials.Dbpassword
var dbhostname string = dbcredentials.Dbhostname
var dbport string = dbcredentials.Dbport
var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

var c string = (strings.Join(b,""))
var db,err  = gorm.Open(dbtype, c)



func checkEnvVar(envVars *map[string]string) error {
	var missingVars []string
	for varName, value := range *envVars {
		if value == "" {
			missingVars = append(missingVars, varName)
		}
	}
	if len(missingVars) > 0 {
		return fmt.Errorf("Missing environment variables %v", missingVars)
	}
	return nil
}

func main()/* (result compute.VirtualMachineListResult, err error)*/{
	resourceGroup := "test"
	os.Setenv("AZURE_CLIENT_ID", "2db3b1e3-b551-4e7a-b6cd-193042323f6a")
	os.Setenv("AZURE_CLIENT_SECRET", "S0aY9oF0L0RGGfUEGoT/HSdqypxXKh7lmaTawlekrxw=")
	os.Setenv("AZURE_SUBSCRIPTION_ID", "96782a8b-5f88-48ac-ac3c-91679baeb0ad")
	os.Setenv("AZURE_TENANT_ID", "db72859f-dc89-46f4-9134-30e8d982ba21")
	c := map[string]string{
		"AZURE_CLIENT_ID":       os.Getenv("AZURE_CLIENT_ID"),
		"AZURE_CLIENT_SECRET":   os.Getenv("AZURE_CLIENT_SECRET"),
		"AZURE_SUBSCRIPTION_ID": os.Getenv("AZURE_SUBSCRIPTION_ID"),
		"AZURE_TENANT_ID":       os.Getenv("AZURE_TENANT_ID")}
	if err := checkEnvVar(&c); err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	spt, err := helpers.NewServicePrincipalTokenFromCredentials(c, azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	ac := compute.NewVirtualMachinesClient(c["AZURE_SUBSCRIPTION_ID"])
	ac.Authorizer = spt
	/*if ls, err = ac.List(resourceGroup); err != nil {
		fmt.Printf("Failed to list virtual machines: %v\n", err)
		return
	}*/

	ls, _ := ac.List(resourceGroup)

	//u := &compute.VirtualMachineListResult{}



	//json.Unmarshal([]byte(ls),&u)

	//user:=azurestruct.AzureInstances{VmName:ls}
//	_ = json.NewEncoder(os.Stdout).Encode(&user)

	for _, element := range *ls.Value{
		//println(element.Name,element.ID,element.Status,element.Progress)
		//user :=	azurestruct.AzureInstances{VmName:element.NextLink}

		user:=azurestruct.AzureInstances{VmName:*element.Name,Type:*element.Type,Location:*element.Location,VmId:*element.VMID}

		db.Create(&user)



	}


	//_ = json.NewEncoder(os.Stdout).Encode(&ls)


//	println(u.Value)
//	println(u)

		//return ls
	//return ls
}
