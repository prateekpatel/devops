package readopenstackconf

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"os"
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
	filename := "confmanagement/readopenstackconf/openstack_conf.go"
       _, filePath, _, _ := runtime.Caller(0)
       fmt.Println("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/osazureconf.json", 1))
       fmt.Println("ABSPATH:==",ConfigFilePath)
	file, _ := os.Open(ConfigFilePath)
	//dir, _ := os.Getwd()
	//file, _ := os.Open(dir + "/src/gclassec/conf/osazureconf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}
