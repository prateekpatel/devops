//package openstackInstance
package main
import (
	_ "github.com/lib/pq"
	_ "go-sql-driver/mysql"
	//"github.com/julienschmidt/httprouter"
	"github.com/jinzhu/gorm"
	//"encoding/json"
	//"net/http"
	//"Go_Project/goClient"
	//"strings"

	//"Go_Project/vw_dynamic_struct"
	//"fmt"
	//"Go_Project"
	//"fmt"
	//"fmt"

	"encoding/json"
	"os"
)

type Instances struct{
	Id int 			`gorm:"column:id"`
	Name string 		`gorm:"column:Name"`
	InstanceID string 	`gorm:"column:InstanceID"`
	Status string 		`gorm:"column:Status"`
	AvailabilityZone string `gorm:"column:AvailabilityZone"`
	CreationTime string 	`gorm:"column:CreationTime"`
	Flavor string 		`gorm:"column:Flavor"`
	FlavorID int 		`gorm:"column:FlavorID"`
	RAM string 		`gorm:"column:RAM"`
	VCPU string 		`gorm:"column:VCPU"`
	Storage string 		`gorm:"column:Storage"`
	IPAddress string	`sql:"type:decimal" gorm:"column:IPAddress"`
	SecurityGroup string 	`gorm:"column:SecurityGroup"`
	KeyPairName string 	`gorm:"column:KeyPairName"`
	ImageName string 	`gorm:"column:ImageName"`
	Volumes string 		`gorm:"column:Volumes"`
	InsertionDate string 	`sql:"type:date" gorm:"column:InsertionDate"`
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(110.110.110.170:3306)/GO-DB-Testing")
	//db.CreateTable(&TTest{})
	user := []Instances{}
	db.SingularTable(true)
	//db.Find(&user)

	_ = json.NewEncoder(os.Stdout).Encode(db.Find(&user))

	if err != nil {
		println(err)
	}

}