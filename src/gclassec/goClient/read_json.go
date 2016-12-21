package goClient

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
<<<<<<< HEAD
	file, _ := os.Open("C:/goclassec/conf.json")
=======
	file, _ := os.Open("C:/Git/goclassec/conf.json")
>>>>>>> 642043858044df1f382c847bbd98812c14fde101
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}

