package hoscontroller

import (
	"net/http"
	"github.com/gorilla/mux"
	"gclassec/hos/OpenStack_API_Function"
	"fmt"
)

type (
    // UserController represents the controller for operating on the User resource
    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}


func (uc UserController) CpuUtilDetails(w http.ResponseWriter, r *http.Request){
        vars := mux.Vars(r)
        id := vars["id"]
        res := OpenStack_API_Function.GetCpuUtilDetails(id)
        fmt.Fprintf(w,res)

}
func (uc UserController) GetComputeDetails(w http.ResponseWriter, r *http.Request){

	res := OpenStack_API_Function.Compute()
        fmt.Fprintf(w,res)

}

func (uc UserController) GetFlavorsDetails(w http.ResponseWriter, r *http.Request){

	res := OpenStack_API_Function.Flavors()
        fmt.Fprintf(w,res)

}

func (uc UserController) GetCeilometerStatitics(w http.ResponseWriter, r *http.Request){

	res := OpenStack_API_Function.GetCpuUtilStatistics()
        fmt.Fprintf(w,res)

}

func (uc UserController) GetCeilometerDetails(w http.ResponseWriter, r *http.Request){

	res := OpenStack_API_Function.GetCeilometerDetail()
        fmt.Fprintf(w,res)

}

func (uc UserController) Index(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Hi You Just tested Server ping.")
}