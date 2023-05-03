package entity

type Machinetype struct {
	LastappearedId int    `gorm:"primary_key;column:lastappeared_id;type:int;not null" json:"lastappeared_id"`
	Machinetype    string `gorm:"column:machinetype;type:string;not null json" json:"machinetype"`
}
