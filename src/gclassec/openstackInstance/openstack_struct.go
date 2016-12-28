package openstackInstance

type Instances struct{
	//Id int 			`gorm:"column:id"`
	Name string 		`gorm:"column:name"`
	InstanceID string 	`gorm:"column:instance_id"`
	Status string 		`gorm:"column:status"`
	AvailabilityZone string `gorm:"column:availability_zone"`
	Flavor string            `gorm:"column:flavor"`
	CreationTime string 	`gorm:"column:CreationTime"`
	FlavorID string 		`gorm:"column:flavor_id"`
	RAM int64 		`gorm:"column:ram"`
	VCPU int64 		`gorm:"column:vcpu"`
	Storage int64 		`gorm:"column:storage"`
	IPAddress string	`sql:"type:decimal" gorm:"column:ip_address"`
	SecurityGroup string 	`gorm:"column:security_group"`
	KeyPairName string 	`gorm:"column:keypair_name"`
	ImageName string 	`gorm:"column:image_name"`
	Volumes string 		`gorm:"column:volumes"`
	InsertionDate string 	`sql:"type:date" gorm:"column:insertion_date"`

}
