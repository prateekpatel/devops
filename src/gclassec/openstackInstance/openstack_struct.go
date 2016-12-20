package openstackInstance

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