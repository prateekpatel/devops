package main

import (
    // Standard library packages
    "net/http"

    // Third party packages
    "github.com/julienschmidt/httprouter"

    "CloudApp/controllers"
)

func main() {
    // Instantiate a new router
    r := httprouter.New()

    // Get a InstController instance
    uc := controllers.NewInstController()

    // Get a instance resource
    r.GET("/dbaas/list", uc.GetDetails)

    r.GET("/dbaas/list/:id", uc.GetDetailsById)

    r.GET("/dbaas/get", uc.GetDB)

    r.GET("/dbaas/pricing", uc.GetPrice)

    // Fire up the server
    http.ListenAndServe("localhost:9009", r)
}