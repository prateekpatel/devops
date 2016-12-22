package main

import (
    // Standard library packages
    "net/http"
    // Third party packages
    "gclassec/Controller"
    "github.com/gorilla/mux"
    "fmt"
    "gclassec/openstackController"
    "gclassec/validation"
    "gclassec/OpenstackInsert"
)

func main() {
    mx := mux.NewRouter()

    OpenstackInsert.InsertInstances()
    // Get a InstController instance
    uc :=Controller.NewUserController()

    op := openstackController.UserController{}

    mx.NotFoundHandler = http.HandlerFunc(validation.ValidateWrongURL)

    // Get a instance resource
    mx.HandleFunc("/dbaas/list", uc.GetDetails).Methods("GET")  // 'http://localhost:9009/dbaas/list'

    mx.HandleFunc("/dbaas/list/{id}", uc.GetDetailsById).Methods("GET")  // 'http://localhost:9009/dbaas/list/dev01-a-tky-customerorderpf'

    mx.HandleFunc("/dbaas/get", uc.GetDB).Methods("GET")  // 'http://localhost:9009/dbaas/get?CPUUtilization_max=5&DatabaseConnections_max=0'

    mx.HandleFunc("/dbaas/pricing", uc.GetPrice).Methods("GET")  // 'http://localhost:9009/dbaas/pricing'

    mx.HandleFunc("/dbaas/openstackDetail", op.GetDetailsOpenstack).Methods("GET")

    //using handle
    http.Handle("/", mx)

    // Fire up the server
    fmt.Println("Server is on Port 8000")
    fmt.Println("Listening .....")
    http.ListenAndServe("0.0.0.0:8000", nil)
}
