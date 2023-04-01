package helper

import (
	"crud-golang-gin-gorm/data/dto"
	"crud-golang-gin-gorm/data/entity"
)

func DtoTagToEntityTagList(dtos []dto.TagDto) []entity.Tags {
	var temp []entity.Tags
	for _, v := range dtos {
		temp = append(temp, DtoTagToEntityTag(v))
	}
	return temp
}

func DtoTagToEntityTag(dtos dto.TagDto) entity.Tags {
	temp := entity.Tags{
		Id:   dtos.Id,
		Name: dtos.Name,
	}
	return temp
}

func EntityTagToDtoTagList(entity []entity.Tags) []dto.TagDto {
	var temp []dto.TagDto
	for _, v := range entity {
		temp = append(temp, EntityTagToDtoTag(v))
	}
	return temp
}

func EntityTagToDtoTag(entity entity.Tags) dto.TagDto {
	temp := dto.TagDto{
		Id:   entity.Id,
		Name: entity.Name,
	}
	return temp
}
