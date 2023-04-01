package service

import "crud-golang-gin-gorm/data/dto"

type TagsService interface {
	Save(tags dto.TagDto) dto.TagDto
	Update(tags dto.TagDto) dto.TagDto
	Delete(tagsId int)
	FindById(tagsId int) (tags dto.TagDto, err error)
	FindAll() []dto.TagDto
}
