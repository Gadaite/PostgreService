package service

import (
	"PostgreService/entity"
	"PostgreService/mapper"
)

func ObjectTrajactoryApi(id string) []entity.Objecttrajactory {
	return mapper.ObjectTrajactoryApi(id)
}

func MachinetypeSearch(lastappearedIds []string, machinetype string) []entity.Machinetype {
	return mapper.MachinetypeSearch(lastappearedIds, machinetype)
}

func MachinetypeUpdate(dto *entity.Machinetype) {
	mapper.MachinetypeUpdate(dto.LastappearedId, dto.Machinetype)
}

func MachinetypeAdd(dto *entity.Machinetype) string {
	existsData := MachinetypeSearch([]string{string(rune(dto.LastappearedId))}, dto.Machinetype)
	if len(existsData) != 0 {
		return "该数据已经存在"
	} else {
		mapper.MachinetypeAdd(dto)
		return ""
	}
}
