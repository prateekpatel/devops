package controllers

import (
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"encoding/json"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
	"gclassec/goClient"
)
type (
    // InstController represents the controller for operating on the User resource
    InstController struct{}
)

func NewInstController() *InstController {
    return &InstController{}
}

type rds_dynamic struct{
	DBInstanceIdentifier string     `gorm:"column:DBInstanceIdentifier"`
	StartTime string                `sql:"type:datetime" gorm:"column:StartTime"`
	EndTime string                 `sql:"type:datetime" gorm:"column:EndTime"`
	Period int64          `sql:"type:bigint" gorm:"column:Period"`
	BinLogDiskUsage_min float64 `gorm:"column:BinLogDiskUsage_min"`
	BinLogDiskUsage_max float64   `gorm:"column:BinLogDiskUsage_max"`
	BinLogDiskUsage_avg float64   `gorm:"column:BinLogDiskUsage_avg"`
	CPUUtilization_min float64    `gorm:"column:CPUUtilization_min"`
	CPUUtilization_max float64   `gorm:"column:CPUUtilization_max"`
	CPUUtilization_avg float64   `gorm:"column:CPUUtilization_avg"`
	CPUCreditUsage_min float64   `gorm:"column:CPUCreditUsage_min"`
	CPUCreditUsage_max float64   `gorm:"column:CPUCreditUsage_max"`
	CPUCreditUsage_avg float64   `gorm:"column:CPUCreditUsage_avg"`
	CPUCreditBalance_min float64 `gorm:"column:CPUCreditBalance_min"`
	CPUCreditBalance_max float64  `gorm:"column:CPUCreditBalance_max"`
	CPUCreditBalance_avg float64  `gorm:"column:CPUCreditBalance_avg"`
	DatabaseConnections_min float64  `gorm:"column:DatabaseConnections_min"`
	DatabaseConnections_max float64   `gorm:"column:DatabaseConnections_max"`
	DatabaseConnections_avg float64 `gorm:"column:DatabaseConnections_avg"`
	DiskQueueDepth_min float64     `gorm:"column:DiskQueueDepth_min"`
	DiskQueueDepth_max float64     `gorm:"column:DiskQueueDepth_max"`
	DiskQueueDepth_avg float64     `gorm:"column:DiskQueueDepth_avg"`
	FreeableMemory_min float64     `gorm:"column:FreeableMemory_min"`
	FreeableMemory_max float64     `gorm:"column:FreeableMemory_max"`
	FreeableMemory_avg float64	`gorm:"column:FreeableMemory_avg"`
	FreeStorageSpace_min float64   `gorm:"column:FreeStorageSpace_min"`
	FreeStorageSpace_max float64   `gorm:"column:FreeStorageSpace_max"`
	FreeStorageSpace_avg float64   `gorm:"column:FreeStorageSpace_avg"`
	ReplicaLag_min float64         `gorm:"column:ReplicaLag_min"`
	ReplicaLag_max float64		`gorm:"column:ReplicaLag_max"`
	ReplicaLag_avg float64     	`gorm:"column:ReplicaLag_avg"`
	SwapUsage_min float64		`gorm:"column:SwapUsage_min"`
	SwapUsage_max float64		`gorm:"column:SwapUsage_max"`
	SwapUsage_avg float64		`gorm:"column:SwapUsage_avg"`
	ReadIOPS_min float64		`gorm:"column:ReadIOPS_min"`
	ReadIOPS_max float64 		`gorm:"column:ReadIOPS_max"`
	ReadIOPS_avg float64		`gorm:"column:ReadIOPS_avg"`
	WriteIOPS_min float64		`gorm:"column:WriteIOPS_min"`
	WriteIOPS_max float64		`gorm:"column:WriteIOPS_max"`
	WriteIOPS_avg float64		`gorm:"column:WriteIOPS_avg"`
	ReadLatency_min float64		`gorm:"column:ReadLatency_min"`
	ReadLatency_max float64		`gorm:"column:ReadLatency_max"`
	ReadLatency_avg float64		`gorm:"column:ReadLatency_avg"`
	WriteLatency_min float64	`gorm:"column:WriteLatency_min"`
	WriteLatency_max float64	`gorm:"column:WriteLatency_max"`
	WriteLatency_avg float64	`gorm:"column:WriteLatency_avg"`
	ReadThroughput_min float64	`gorm:"column:ReadThroughput_min"`
	ReadThroughput_max float64	`gorm:"column:ReadThroughput_max"`
	ReadThroughput_avg float64	`gorm:"column:ReadThroughput_avg"`
	WriteThroughput_min float64	`gorm:"column:WriteThroughput_min"`
	WriteThroughput_max float64	`gorm:"column:WriteThroughput_max"`
	WriteThroughput_avg float64	`gorm:"column:WriteThroughput_avg"`
	NetworkReceiveThroughput_min float64	`gorm:"column:NetworkReceiveThroughput_min"`
	NetworkReceiveThroughput_max float64  	`gorm:"column:NetworkReceiveThroughput_max"`
	NetworkReceiveThroughput_avg float64	`gorm:"column:NetworkReceiveThroughput_avg"`
	NetworkTransmitThroughput_min float64	`gorm:"column:NetworkTransmitThroughput_min"`
	NetworkTransmitThroughput_max float64	`gorm:"column:NetworkTransmitThroughput_max"`
	NetworkTransmitThroughput_avg float64   `gorm:"column:NetworkTransmitThroughput_avg"`
}

type vw_rds struct {
	API_Name string `gorm:"column:API_Name"`
	Linux_On_Demand_cost string `gorm:"column:Linux_On_Demand_cost"`
	Linux_Reserved_cost string `gorm:"column:Linux_Reserved_cost"`
	Windows_On_Demand_cost string `gorm:"column:Windows_On_Demand_cost"`
	Windows_Reserved_cost string `gorm:"column:Windows_Reserved_cost"`
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
Get instance details
*/
func (uc InstController) GetDetails(w http.ResponseWriter, r *http.Request) {
	dbObj := []rds_dynamic{}

	db.SingularTable(true)

	/*for _, v := range dbObj {
		//fmt.Println("Id : ", v.Id)
		//fmt.Println("Username : ", v.Name)
		fmt.Println(v)
	        _ = json.NewEncoder(w).Encode(db.Find(&dbObj))
        }*/

	_ = json.NewEncoder(w).Encode(db.Find(&dbObj))

	w.WriteHeader(http.StatusOK)

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	//db.Close()
}

/*
Get details of instance for the given id
*/
func (uc InstController) GetDetailsById(w http.ResponseWriter, r *http.Request) {
	dbObj := []rds_dynamic{}

	//id := p.ByName("id")

	vars := mux.Vars(r)
	id := vars["id"]

	db.SingularTable(true)

	_ = json.NewEncoder(w).Encode(db.Find(&dbObj, "DBInstanceIdentifier = ?", id))

	w.WriteHeader(http.StatusOK)

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	//db.Close()
}

/*
Get details of instances based on CPU Utilization and DB Connections
*/
func (uc InstController) GetDB(w http.ResponseWriter, r *http.Request) {
	dbObj := []rds_dynamic{}

	queryValue1 := r.URL.Query().Get("CPUUtilization_max")

	queryValue2 := r.URL.Query().Get("DatabaseConnections_max")

	println(queryValue1)
	println(queryValue2)

	db.SingularTable(true)

	//_ = json.NewEncoder(w).Encode(db.Find(&dbObj, "CPUUtilization_max<5 && DatabaseConnections_max=0"))

	_ = json.NewEncoder(w).Encode(db.Where("CPUUtilization_max < ? AND DatabaseConnections_max = ?", queryValue1, queryValue2).Find(&dbObj))

	w.WriteHeader(http.StatusOK)

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	//db.Close()
}

/*
Get pricing depending on instance type
*/
func (uc InstController) GetPrice(w http.ResponseWriter, r *http.Request) {
	dbObj := []vw_rds{}

	db.SingularTable(true)

	_ = json.NewEncoder(w).Encode(db.Find(&dbObj))

	w.WriteHeader(http.StatusOK)

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	//db.Close()
}
