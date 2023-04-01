package repository

import "crud-golang-gin-gorm/data/entity"

type TagRepository interface {
	Save(tags entity.Tags) entity.Tags
	Update(tags entity.Tags) entity.Tags
	Delete(tagsId int)
	FindById(tagsId int) (tags entity.Tags, err error)
	FindAll() []entity.Tags
}
