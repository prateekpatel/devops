package main

import (
    // Standard library packages
    "net/http"

    // Third party packages
    "gclassec/controllers"
    "github.com/gorilla/mux"
)

func main() {
    // Instantiate a new router
    mx := mux.NewRouter()

    // Get a InstController instance
    uc := controllers.NewInstController()

    // Get a instance resource
    mx.HandleFunc("/dbaas/list", uc.GetDetails).Methods("GET")

    mx.HandleFunc("/dbaas/list/{id}", uc.GetDetailsById).Methods("GET")

    mx.HandleFunc("/dbaas/get", uc.GetDB).Methods("GET")

    mx.HandleFunc("/dbaas/pricing", uc.GetPrice).Methods("GET")

    // Fire up the server
    http.ListenAndServe("localhost:9009", mx)
}