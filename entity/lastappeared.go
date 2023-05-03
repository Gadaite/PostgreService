package entity

type Lastappeared struct {
	Id               int    `gorm:"primary_key;column:id;type:int;not null" json:"id"`
	ObjectId         string `gorm:"column:object_id;type:string" json:"object_id"`
	LastmodifiedDate string `gorm:"column:lastmodified_date;type:string" json:"lastmodified_date"`
	LastmodifiedTime string `gorm:"column:lastmodified_time;type:string" json:"lastmodified_time"`
	GpsPoint         string `gorm:"column:gps_point;type:string" json:"gps_point"`
}
