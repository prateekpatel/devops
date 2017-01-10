package main

import (
    // Standard library packages
    "net/http"
    // Third party packages
    "gclassec/awscontroller"
    "github.com/gorilla/mux"
    "fmt"
    "gclassec/openstackcontroller"
    "gclassec/validation"
    "gclassec/openstackinsert"

    "gclassec/azureinsert"
    "gclassec/azurecontroller"
)

func main() {
    mx := mux.NewRouter()

    openstackinsert.InsertInstances()
    azureinsert.AzureInsert()

    uc := awscontroller.NewUserController()

    op := openstackcontroller.NewUserController()
    op1 := azurecontroller.NewUserController()

    mx.NotFoundHandler = http.HandlerFunc(validation.ValidateWrongURL)

    // Get a instance resource
    mx.HandleFunc("/dbaas/list", uc.GetDetails).Methods("GET")  // 'http://localhost:9009/dbaas/list'

    mx.HandleFunc("/dbaas/list/{id}", uc.GetDetailsById).Methods("GET")  // 'http://localhost:9009/dbaas/list/dev01-a-tky-customerorderpf'

    mx.HandleFunc("/dbaas/get", uc.GetDB).Methods("GET")  // 'http://localhost:9009/dbaas/get?CPUUtilization_max=5&DatabaseConnections_max=0'

    mx.HandleFunc("/dbaas/pricing", uc.GetPrice).Methods("GET")  // 'http://localhost:9009/dbaas/pricing'

    mx.HandleFunc("/dbaas/openstackDetail", op.GetDetailsOpenstack).Methods("GET")

    mx.HandleFunc("/dbaas/azureDetail", op1.GetAzureDetails).Methods("GET") // http://localhost:9009/dbaas/azureDetail
    http.Handle("/", mx)

    // Fire up the server
    fmt.Println("Server is on Port 9009")
    fmt.Println("Listening .....")



    http.ListenAndServe("0.0.0.0:9009", nil)
}