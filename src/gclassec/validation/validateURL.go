package validation

import (
	"net/http"
	"fmt"
)

func ValidateWrongURL(w http.ResponseWriter, r *http.Request){
	 w.WriteHeader(http.StatusNotFound)
	 w.Write([]byte(fmt.Sprintf("%s URL is Not Valid Please Enter Valid URL\n", r.URL)))
}
