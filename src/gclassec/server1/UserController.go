package main
import (
    // Standard library packages
    "net/http"

    // Third party packages
    "github.com/julienschmidt/httprouter"
    //"github.com/swhite24/go-rest-tutorial/controllers"

    "Go_Project/Go_Project/Controller"


)

func main() {
    // Instantiate a new router
    r := httprouter.New()

    // Get a UserController instance
    uc := sqlx_connect.NewUserController()

    // Get a user resource
    r.GET("/dbaas/list", uc.GetDetails)

    r.GET("/dbaas/list/:id", uc.GetDetailsById)

    r.GET("/dbaas/get", uc.GetDB)

     r.GET("/dbaas/pricing", uc.GetPrice)
    //r.POST("/user", uc.CreateUser)

    //r.DELETE("/user/:id", uc.RemoveUser)

    // Fire up the server
    http.ListenAndServe("localhost:9090", r)
}