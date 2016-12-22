package Controller

import (
	"net/http"
	//_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	"gclassec/goClient"
	"strings"
	//"github.com/jinzhu/gorm"
	"gclassec/table_structs"
	"encoding/json"
	"github.com/gorilla/mux"
	//"src/github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm"
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

func (uc UserController) GetDetailsById(w http.ResponseWriter, r *http.Request) {
	dbObj := []structs.Rds_dynamic{}

	//id := p.ByName("id")

	vars := mux.Vars(r)
	id := vars["id"]

	db.SingularTable(true)

	_ = json.NewEncoder(w).Encode(db.Find(&dbObj, "DBInstanceIdentifier = ?", id))

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	//db.Close()
}

func (uc UserController) GetDB(w http.ResponseWriter, r *http.Request) {
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



func (uc UserController) GetPrice(w http.ResponseWriter, r *http.Request) {
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

func (uc UserController) GetDetails(w http.ResponseWriter, r *http.Request){

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
