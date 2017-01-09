package main

import (
	"os"
	"log"

	"github.com/Azure/go-autorest/autorest/azure"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/arm/compute"
	"github.com/Azure/azure-sdk-for-go/arm/examples/helpers"
	/*"encoding/json"
	//"golang.org/x/tools/go/gcimporter15/testdata"
	"gclassec/azurestruct"*/


)

type ls struct {

}


//type Members azurestruct.AzureInstances
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

func main() {
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


	ls, _ := ac.List(resourceGroup)


}