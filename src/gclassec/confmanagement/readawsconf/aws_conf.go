package readawsconf

import (
	"os"
	"encoding/json"
	"fmt"
)

type Configuration struct {
    Dbtype    string
    Dbname   string
	Dbusername   string
    Dbpassword   string
	Dbhostname   string
	Dbport   string
}

func Configurtion() Configuration{
	dir, _ := os.Getwd()
	file, _ := os.Open(dir + "/src/gclassec/conf/awsconf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}

