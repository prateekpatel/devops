package main

import (
    // Standard library packages
    "net/http"
    // Third party packages
    "gclassec/controllers/awscontroller"
    "github.com/gorilla/mux"
    "fmt"
    "gclassec/controllers/openstackcontroller"
    "gclassec/validation"
    "gclassec/dao/openstackinsert"
    "gclassec/dao/azureinsert"
    "gclassec/controllers/azurecontroller"
    "os"
    "gclassec/controllers/confcontroller"
)

func main() {
    mx := mux.NewRouter()

    uc := awscontroller.NewUserController()
    op := openstackcontroller.NewUserController()
    op1 := azurecontroller.NewUserController()
    uc1 := confcontroller.NewUserController()

    openstackinsert.InsertInstances()
    azureinsert.AzureInsert()

    mx.NotFoundHandler = http.HandlerFunc(validation.ValidateWrongURL)

    // Get a instance resource
    mx.HandleFunc("/dbaas/list", uc.GetDetails).Methods("GET")  // 'http://localhost:9009/dbaas/list'

    mx.HandleFunc("/dbaas/list/{id}", uc.GetDetailsById).Methods("GET")  // 'http://localhost:9009/dbaas/list/dev01-a-tky-customerorderpf'

    mx.HandleFunc("/dbaas/get", uc.GetDB).Methods("GET")  // 'http://localhost:9009/dbaas/get?CPUUtilization_max=5&DatabaseConnections_max=0'

    mx.HandleFunc("/dbaas/pricing", uc.GetPrice).Methods("GET")  // 'http://localhost:9009/dbaas/pricing'

    mx.HandleFunc("/dbaas/openstackDetail", op.GetDetailsOpenstack).Methods("GET")

    mx.HandleFunc("/dbaas/azureDetail", op1.GetAzureDetails).Methods("GET") // http://localhost:9009/dbaas/azureDetail

    mx.HandleFunc("/dbaas/azureDetail/percentCPU/{resourceGroup}/{name}", op1.GetDynamicAzureDetails).Methods("GET")

    mx.HandleFunc("/selectProvider", uc1.SelectProvider)

    mx.HandleFunc("/selectedOs", uc1.OpenstackCreds)

	mx.HandleFunc("/selectedAzure", uc1.AzureCreds)

    mx.HandleFunc("/providers", uc1.ProviderHandler).Methods("POST")

    mx.HandleFunc("/providers/openstack", uc1.ProviderOpenstack).Methods("POST")

	mx.HandleFunc("/providers/azure", uc1.ProviderAzure).Methods("POST")

	http.Handle("/", mx)

    // Fire up the server
    fmt.Println("Server is on Port 9009")
    fmt.Println("Listening .....")

    fmt.Println(os.Getwd())

    http.ListenAndServe("0.0.0.0:9009", nil)
}