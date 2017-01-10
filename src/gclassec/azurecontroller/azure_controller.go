package azurecontroller

import(
	"strings"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"

	"gclassec/azurestruct"


	"gclassec/readazureconf"

	"net/http"
)
type (

    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}


var dbcredentials1 = readazureconf.Configurtion()
var dbtype string = dbcredentials1.Dbtype
var dbname  string = dbcredentials1.Dbname
var dbusername string = dbcredentials1.Dbusername
var dbpassword string = dbcredentials1.Dbpassword
var dbhostname string = dbcredentials1.Dbhostname
var dbport string = dbcredentials1.Dbport

var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

var c string = (strings.Join(b,""))

var db,err  = gorm.Open(dbtype, c)

func   (uc UserController) GetAzureDetails(w http.ResponseWriter, r *http.Request)(){

	tx := db.Begin()
	db.SingularTable(true)

	azure_struct := []azurestruct.AzureInstances{}

	err := db.Find(&azure_struct).Error

	if err != nil{

		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Find(&azure_struct))

		if err != nil {
			println(err)
		}

	tx.Commit()
}