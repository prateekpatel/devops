package controllers

import(
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
	_"github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"fmt"
	"time"
	"github.com/jmoiron/sqlx"
)
type (
    // InstController represents the controller for operating on the Instance resource
    InstController struct{}
)

func NewInstController() *InstController {
    return &InstController{}
}

type ec2_dynamic struct {
	InsT_id string
	EC2_cpu_util_max float64
	EC2_cpu_util_min float64
	EC2_cpu_util_avg float64
	EC2_start_time time.Time
	EC2_end_time time.Time
}

func (uc InstController) GetDetails(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	db,err := sqlx.Connect("mysql", "root:root@tcp(110.110.110.170:3306)/cloud_assessment")

	fmt.Print(db)

	inst := ec2_dynamic{}

	rows, _ := db.Queryx("SELECT * FROM ec2_dynamic where ec2_cpu_util_max>90")

	for rows.Next() {
		rows.StructScan(&inst)
	    _ = json.NewEncoder(w).Encode(inst)
    	}

	if err != nil {
		println(err)
	}
}