package main

import (
    // Standard library packages
    "net/http"

    // Third party packages
    "gclassec/controllers"
    "github.com/gorilla/mux"
    "fmt"
)

func main() {
    // Instantiate a new router
    mx := mux.NewRouter()

    // Get a InstController instance
    uc := controllers.NewInstController()

    mx.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(fmt.Sprintf("%s URL is Not Valid Please Enter Valid URL\n", r.URL)))
    })

    // Get a instance resource
    mx.HandleFunc("/dbaas/list", uc.GetDetails).Methods("GET")  // 'http://localhost:9009/dbaas/list'

    mx.HandleFunc("/dbaas/list/{id}", uc.GetDetailsById).Methods("GET")  // 'http://localhost:9009/dbaas/list/dev01-a-tky-customerorderpf'

    mx.HandleFunc("/dbaas/get", uc.GetDB).Methods("GET")  // 'http://localhost:9009/dbaas/get?CPUUtilization_max=5&DatabaseConnections_max=0'

    mx.HandleFunc("/dbaas/pricing", uc.GetPrice).Methods("GET")  // 'http://localhost:9009/dbaas/pricing'

    //using handle
    http.Handle("/", mx)

    // Fire up the server
    fmt.Println("Server is on Port 9009")
    fmt.Println("Listening .....")
    http.ListenAndServe("localhost:9009", nil)
}