package azureinsert


import (
	"os"
	"log"
	"github.com/Azure/go-autorest/autorest/azure"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/arm/compute"
	"github.com/Azure/azure-sdk-for-go/arm/examples/helpers"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"gclassec/readazureconf"
	"gclassec/azurestruct"
	"github.com/jinzhu/gorm"
	"encoding/json"
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

func AzureInsert() {
	//resourceGroup := "test"
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

	ls, _ := ac.ListAll()
	_ = json.NewEncoder(os.Stdout).Encode(&ls)

	//var drggroup string
	for _, element := range *ls.Value {
		//println(element.Name,element.ID,element.Status,element.Progress)
		//user :=	azurestruct.AzureInstances{VmName:element.NextLink}
	rgroup:=*(element.AvailabilitySet.ID)
	resourcegroupname := strings.Split(rgroup, "/")
		//drggroup= resourcegroupname[4]
	user := azurestruct.AzureInstances{VmName:*element.Name, Type:*element.Type, Location:*element.Location,VmSize:element.VirtualMachineProperties.HardwareProfile.VMSize, VmId:*element.VMID, Publisher:*(element.StorageProfile.ImageReference.Publisher), Offer:*(element.StorageProfile.ImageReference.Offer), SKU:*(element.StorageProfile.ImageReference.Sku), AvailabilitySetName:*(element.AvailabilitySet.ID), Provisioningstate:*element.ProvisioningState,ResourcegroupName:resourcegroupname[4]}
	db.Create(&user)
	}
	//Get dynamic details (i.e. Percent CPU Utilization)
	// of Azure Virtual Machine
	/*dc := compute.NewDynamicUsageOperationsClient(c["AZURE_SUBSCRIPTION_ID"])
	dc.Authorizer = spt

	dlist, _ := dc.ListDynamic("testGo",drggroup )
	fmt.Println(dlist)

	_ = json.NewEncoder(os.Stdout).Encode(&dlist)*/


}