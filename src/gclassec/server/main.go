package main

import (
    // Standard library packages
    "net/http"

    // Third party packages
    "github.com/julienschmidt/httprouter"
    "gclassec/controllers"
)

func main() {
    // Instantiate a new router
    r := httprouter.New()

    // Get a InstController instance
    uc := controllers.NewInstController()

    // Get a instance resource
    r.GET("/dbaas/list", uc.GetDetails)

    // Fire up the server
    http.ListenAndServe("localhost:9000", r)
}