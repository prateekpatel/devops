package controllers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)
type (
    // InstController represents the controller for operating on the User resource
    InstController struct{}
)

func NewInstController() *InstController {
    return &InstController{}
}

type rds_dynamic struct{
	DBInstanceIdentifier string `gorm:"column:DBInstanceIdentifier"`
	CPUUtilization_max float64 `gorm:"column:CPUUtilization_max"`
	DatabaseConnections_max float64 `gorm:"column:DatabaseConnections_max"`
}

type vw_rds struct {
	API_Name string `gorm:"column:API_Name"`
	Linux_On_Demand_cost string `gorm:"column:Linux_On_Demand_cost"`
	Linux_Reserved_cost string `gorm:"column:Linux_Reserved_cost"`
	Windows_On_Demand_cost string `gorm:"column:Windows_On_Demand_cost"`
	Windows_Reserved_cost string `gorm:"column:Windows_Reserved_cost"`
}

/*type Ec2_dynamic struct {
	Inst_id string `gorm:"column:inst_id"`
	EC2_cpu_util_max float64 `gorm:"column:ec2_cpu_util_max"`
	EC2_cpu_util_min float64 `gorm:"column:ec2_cpu_util_min"`
	EC2_cpu_util_avg float64 `gorm:"column:ec2_cpu_util_avg"`
	EC2_start_time string `sql:"type:date" gorm:"column:ec2_start_time"`
	EC2_end_time string `sql:"type:date" gorm:"column:ec2_end_time"`
}*/

/*
Get instance details
*/
func (uc InstController) GetDetails(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	dbObj := []rds_dynamic{}

    db, err := gorm.Open("mysql", "root:root@tcp(110.110.110.170:3306)/cloud_assessment")
	println(db)
	println(err)

	db.SingularTable(true)

	/*for _, v := range dbObj {
		//fmt.Println("Id : ", v.Id)
		//fmt.Println("Username : ", v.Name)
		fmt.Println(v)
	        _ = json.NewEncoder(w).Encode(db.Find(&dbObj))
        }*/

	_ = json.NewEncoder(w).Encode(db.Find(&dbObj))

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	db.Close()
}

/*
Get details of instance for the given id
*/
func (uc InstController) GetDetailsById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	dbObj := []rds_dynamic{}

	id := p.ByName("id")

    db, err := gorm.Open("mysql", "root:root@tcp(110.110.110.170:3306)/cloud_assessment")
	println(db)
	println(err)

	db.SingularTable(true)

	_ = json.NewEncoder(w).Encode(db.Find(&dbObj, "DBInstanceIdentifier = ?", id))

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	db.Close()
}

/*
Get details of instances based on CPU Utilization and DB Connections
*/
func (uc InstController) GetDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	dbObj := []rds_dynamic{}

	queryValue1 := r.URL.Query().Get("CPUUtilization_max")

	queryValue2 := r.URL.Query().Get("DatabaseConnections_max")

	println(queryValue1)
	println(queryValue2)

    db, err := gorm.Open("mysql", "root:root@tcp(110.110.110.170:3306)/cloud_assessment")
	println(db)
	println(err)

	db.SingularTable(true)

	//_ = json.NewEncoder(w).Encode(db.Find(&dbObj, "CPUUtilization_max<5 && DatabaseConnections_max=0"))

	_ = json.NewEncoder(w).Encode(db.Where("CPUUtilization_max < ? AND DatabaseConnections_max = ?", queryValue1, queryValue2).Find(&dbObj))

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	db.Close()
}

/*
Get pricing depending on instance type
*/
func (uc InstController) GetPrice(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	dbObj := []vw_rds{}

    db, err := gorm.Open("mysql", "root:root@tcp(110.110.110.170:3306)/cloud_assessment")
	println(db)
	println(err)

	db.SingularTable(true)

	_ = json.NewEncoder(w).Encode(db.Find(&dbObj))

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	db.Close()
}
