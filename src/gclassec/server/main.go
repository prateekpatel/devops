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
    "gclassec/controllers/vmwarecontroller"
    "gclassec/controllers/hoscontroller"
)

func main() {
    mx := mux.NewRouter()

    uc := awscontroller.NewUserController()
    op := openstackcontroller.NewUserController()
    ac := azurecontroller.NewUserController()
    uc1 := confcontroller.NewUserController()
    vc := vmwarecontroller.NewUserController()
    hc := hoscontroller.NewUserController()

    openstackinsert.InsertInstances()
    azureinsert.AzureInsert()

    mx.NotFoundHandler = http.HandlerFunc(validation.ValidateWrongURL)

    // Get a instance resource
    mx.HandleFunc("/hos/computedetails",hc.GetComputeDetails).Methods("GET")
    mx.HandleFunc("/hos/flavorsdetails",hc.GetFlavorsDetails).Methods("GET")
    mx.HandleFunc("/hos/cpu_utilization/{id}",hc.CpuUtilDetails).Methods("GET")
	//mux.HandleFunc("/hos/ceilometerstatitics",GetCeilometerStatitics).Methods("GET")
	//mux.HandleFunc("/hos/ceilometerdetails",GetCeilometerDetails).Methods("GET")
    mx.HandleFunc("/hos/index",hc.Index).Methods("GET")

    mx.HandleFunc("/dbaas/list", uc.GetDetails).Methods("GET")  // 'http://localhost:9009/dbaas/list'

    mx.HandleFunc("/dbaas/list/{id}", uc.GetDetailsById).Methods("GET")  // 'http://localhost:9009/dbaas/list/dev01-a-tky-customerorderpf'

    mx.HandleFunc("/dbaas/get", uc.GetDB).Methods("GET")  // 'http://localhost:9009/dbaas/get?CPUUtilization_max=5&DatabaseConnections_max=0'

    mx.HandleFunc("/dbaas/pricing", uc.GetPrice).Methods("GET")  // 'http://localhost:9009/dbaas/pricing'

    mx.HandleFunc("/dbaas/openstackDetail", op.GetDetailsOpenstack).Methods("GET")

    mx.HandleFunc("/dbaas/azureDetail", ac.GetAzureDetails).Methods("GET") // http://localhost:9009/dbaas/azureDetail

    mx.HandleFunc("/dbaas/azureDetail/percentCPU/{resourceGroup}/{name}", ac.GetDynamicAzureDetails).Methods("GET")

    mx.HandleFunc("/dbaas/vcenterDetail", vc.GetDynamicVcenterDetails).Methods("GET")

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