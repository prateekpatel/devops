package readawsconf

import (
	"os"
	"encoding/json"
	"fmt"
)

type Configuration struct {
    Clientid    string
    ClientSecret   string
	Subscriptionid   string
    Tenantid   string

}

func Configurtion() Configuration{
	file, _ := os.Open("C:\azure-sdk-for-go-master\azure-sdk-for-go-master\azurecred.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}

