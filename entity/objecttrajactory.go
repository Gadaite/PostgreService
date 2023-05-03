package entity

type Objecttrajactory struct {
	LastappearedId int    `gorm:"primary_key;column:lastappeared_id;type:int;not null" json:"lastappeared_id"`
	GpsLine        string `gorm:"column:gps_line;type:string;not null json" json:"gps_line"`
}
