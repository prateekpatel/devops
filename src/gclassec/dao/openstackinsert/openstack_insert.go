package openstackinsert

import (

	"gclassec/confmanagement/readopenstackconf"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"github.com/jinzhu/gorm"
	"gclassec/structs/openstackInstance"
	"gclassec/goclientopenstack"
)
type (
    // UserController represents the controller for operating on the User resource
    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}

var dbcredentials = readopenstackconf.Configurtion()
var dbtype string = dbcredentials.Dbtype
var dbname  string = dbcredentials.Dbname
var dbusername string = dbcredentials.Dbusername
var dbpassword string = dbcredentials.Dbpassword
var dbhostname string = dbcredentials.Dbhostname
var dbport string = dbcredentials.Dbport
var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

var c string = (strings.Join(b,""))
var db,err  = gorm.Open(dbtype, c)

func InsertInstances(){
	//println(examples.ComputeFunc())
	computeDetails:= goclientcompute.FinalCompute()
	println(computeDetails)
	for _, element := range computeDetails {
		//println(element.Name,element.ID,element.Status,element.Progress)
		/*user :=	openstackInstance.Instances{Name:element.Name,InstanceID:element.ID,Status:element.Status,AvailabilityZone:element.Availability_zone,CreationTime:element.Created,
		Volumes:element.Volumes_attached,KeyPairName:element.Key_name}*/


		user:=openstackInstance.Instances{Name:element.Name,InstanceID:element.ID,Status:element.Status,RAM:element.Flavor.RAM,VCPU:element.Flavor.VCPU,Flavor:element.Flavor.Name,Storage:element.Flavor.Disk,AvailabilityZone:element.Availability_zone,CreationTime:element.Created,
		FlavorID:element.Flavor.FlavorID,IPAddress:element.IPV4,KeyPairName:element.Key_name,ImageName:element.Image.ID}
		db.Create(&user)
	}

}