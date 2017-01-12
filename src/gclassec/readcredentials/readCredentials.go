package readazurecreds

import (
	"os"
	"encoding/json"
	"fmt"
)

type Configuration struct {
    ClientId    string
    ClientSecret   string
    SubscriptionId   string
    TenantId   string

}

func Configurtion() Configuration{
	dir, _ := os.Getwd()
	file, _ := os.Open(dir + "\\src\\gclassec\\conf\\azurecred.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}
