package sqlx_connect

import (
	_ "github.com/lib/pq"

	_ "go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/jinzhu/gorm"
	"encoding/json"
	"net/http"
	"Go_Project/Go_Project/goClient"
	"strings"
	"Go_Project/Go_Project/table_structs"
	//"Go_Project/vw_dynamic_struct"
	//"fmt"
	//"Go_Project"
	//"fmt"
	"fmt"

)



type (
    // UserController represents the controller for operating on the User resource
    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}

		var dbcredentials = goClient.Configurtion()
		var dbtype string = dbcredentials.Dbtype
		var dbname  string = dbcredentials.Dbname
		var dbusername string = dbcredentials.Dbusername
		var dbpassword string = dbcredentials.Dbpassword
		var dbhostname string = dbcredentials.Dbhostname
		var dbport string = dbcredentials.Dbport

		var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

		var c string = (strings.Join(b,""))

		var db,err  = gorm.Open(dbtype, c)


/*
Get details of instance for the given id
*/

func (uc UserController) GetDetailsById(w http.ResponseWriter, r *http.Request, p httprouter.Params){

	tx := db.Begin()
	dbObj := []structs.Rds_dynamic{}
	id := p.ByName("id")
	db.SingularTable(true)

	err := db.Find(&dbObj, "DBInstanceIdentifier = ?", id).Error
	fmt.Println(err)
	if err != nil {

		tx.Rollback()
	}


	_ = json.NewEncoder(w).Encode(db.Find(&dbObj, "DBInstanceIdentifier = ?", id))

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	tx.Commit()
}

//Get details of instances based on CPU Utilization and DB Connections



func (uc UserController) GetDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tx := db.Begin()
	dbObj := []structs.Rds_dynamic{}

	queryValue1 := r.URL.Query().Get("CPUUtilization_max")

	queryValue2 := r.URL.Query().Get("DatabaseConnections_max")

	println(queryValue1)
	println(queryValue2)
	db.SingularTable(true)

	err := db.Where("CPUUtilization_max < ? AND DatabaseConnections_max = ?", queryValue1, queryValue2).Find(&dbObj).Error
if err != nil{
		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Where("CPUUtilization_max < ? AND DatabaseConnections_max = ?", queryValue1, queryValue2).Find(&dbObj))

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	tx.Commit()

}

//Get pricing depending on instance type



func (uc UserController) GetPrice(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tx := db.Begin()
	dbObj := []structs.Vw_rds{}


	db.SingularTable(true)
	err := db.Find(&dbObj).Error
if err != nil{
		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Find(&dbObj))

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	tx.Commit()
}

func (uc UserController) GetDetails(w http.ResponseWriter, r *http.Request, p httprouter.Params){

	tx := db.Begin()
	db.SingularTable(true)

	rds_dynamic := []structs.Rds_dynamic{}

	err := db.Find(&rds_dynamic).Error

	if err != nil{

		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Find(&rds_dynamic))

		if err != nil {
			println(err)
		}

	tx.Commit()
}




