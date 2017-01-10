package azurestruct



//import "github.com/Azure/go-autorest/autorest/date"

//import "github.com/Azure/go-autorest/autorest/date"

type AzureInstances struct{
	VmName string 			`gorm:"column:name"`
	Type string 			`gorm:"column:type"`
	Location string 		`gorm:"column:location"`
	VmSize string            	`gorm:"column:vmsize"`
	Publisher string 		`gorm:"column:publisher"`
	Offer string 			`gorm:"column:offer"`
	SKU string 			`gorm:"column:sku"`
	VmId string			`gorm:"column:vmid"`
	AvailabilitySetName string 	`gorm:"column:availabilitysetid"`
	Provisioningstate string	`sql:"type:decimal" gorm:"column:provisioningstate"`
	ResourcegroupName string	`gorm:"column:resourcegroupname"`

}



type AzureInstancesDynamic struct {
	Id	       			string `gorm:"column:id"`
	Timestamp 			string `gorm:"column:timestamp"`
	Average				float32 `gorm:"column:average"`



}