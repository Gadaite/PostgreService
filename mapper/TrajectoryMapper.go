package mapper

import (
	"PostgreService/common"
	pgDB "PostgreService/config"
	"PostgreService/entity"
	"github.com/google/uuid"
)

func ObjectTrajactoryApi(id string) []entity.Objecttrajactory {
	var res []entity.Objecttrajactory
	seRes := pgDB.PostDB.Model(&entity.Objecttrajactory{}).Table("objecttrajactory").Select("lastappeared_id,st_asewkt(gps_line) as gps_line").Where("lastappeared_id = ?", id)
	seRes.Find(&res)
	rows, _ := seRes.Rows()
	_ = common.PrintQueryResults(rows)
	return res
}

func MachinetypeSearch(lastappearedIds []string, machinetype string) []entity.Machinetype {
	var res []entity.Machinetype
	condition := pgDB.PostDB.Model(&entity.Machinetype{}).Table("machinetype")
	if len(lastappearedIds) != 0 {
		condition = condition.Where("lastappeared_id IN (?)", lastappearedIds)
	}
	if machinetype != "" {
		condition = condition.Where("machinetype = ?", machinetype)
	}
	condition.Find(&res)
	//rows, _ := condition.Rows()
	//common.PrintQueryResults(rows)
	return res
}

func MachinetypeUpdate(lastappeared_id int, machinetype string) {
	condition := pgDB.PostDB.Model(&entity.Machinetype{}).Table("machinetype")
	condition.Where("lastappeared_id = ?", lastappeared_id).Updates(map[string]interface{}{
		"lastappeared_id": lastappeared_id,
		"machinetype":     machinetype,
	})
}

func MachinetypeAdd(dto *entity.Machinetype) {
	lastappeared := entity.Lastappeared{
		Id:       dto.LastappearedId,
		ObjectId: uuid.NewString(),
	}
	pgDB.PostDB.Model(&entity.Lastappeared{}).Table("lastappeared").Create(&lastappeared)
	pgDB.PostDB.Model(&entity.Machinetype{}).Table("machinetype").Save(&dto)
}
