package openstackController


import(
	"strings"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"


	"net/http"

	"encoding/json"
	"gclassec/opClient"
	"gclassec/openstackInstance"
)
type (
    // UserController represents the controller for operating on the User resource
    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}

		var dbcredentials1 = opClient.Configurtion()
		var dbtype string = dbcredentials1.Dbtype
		var dbname  string = dbcredentials1.Dbname
		var dbusername string = dbcredentials1.Dbusername
		var dbpassword string = dbcredentials1.Dbpassword
		var dbhostname string = dbcredentials1.Dbhostname
		var dbport string = dbcredentials1.Dbport

		var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

		var c string = (strings.Join(b,""))

		var db,err  = gorm.Open(dbtype, c)

func (uc UserController) GetDetailsOpenstack(w http.ResponseWriter, r *http.Request, p httprouter.Params){

	tx := db.Begin()
	db.SingularTable(true)

	openstack_struct := []openstackInstance.Instances{}

	err := db.Find(&openstack_struct).Error

	if err != nil{

		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Find(&openstack_struct))

		if err != nil {
			println(err)
		}

	tx.Commit()
}