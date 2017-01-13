package readazurecreds

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"os"
)

type Configuration struct {
    ClientId    string
    ClientSecret   string
    SubscriptionId   string
    TenantId   string

}

func Configurtion() Configuration{
	filename := "readcredentials/readCredentials.go"
       _, filePath, _, _ := runtime.Caller(0)
       fmt.Println("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/azurecred.json", 1))
       fmt.Println("ABSPATH:==",ConfigFilePath)
	file, _ := os.Open(ConfigFilePath)
	//dir, _ := os.Getwd()
	//file, _ := os.Open(dir + "/src/gclassec/conf/azurecred.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}
