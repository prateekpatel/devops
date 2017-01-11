package readopenstackconf

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
	file, _ := os.Open("C:\\goclassec\\src\\gclassec\\conf\\osazureconf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}
