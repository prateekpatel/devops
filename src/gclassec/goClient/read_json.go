package goClient

import (
    "encoding/json"
    "os"
    "fmt"
	//"strings"
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
	file, _ := os.Open("C:/Projects/Go/src/Go_Project/Go_Project/conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration

}