package repository

import (
	"crud-golang-gin-gorm/data/entity"
	"crud-golang-gin-gorm/helper"
	"errors"

	"gorm.io/gorm"
)

type TagRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagRepositoryImpl(Db *gorm.DB) TagRepository {
	return &TagRepositoryImpl{Db: Db}
}

// Delete implements TagRepository
func (t *TagRepositoryImpl) Delete(tagsId int) {
	var tags entity.Tags
	result := t.Db.Where("id=?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagRepository
func (t *TagRepositoryImpl) FindAll() []entity.Tags {
	var tagList []entity.Tags
	result := t.Db.Find(&tagList)
	helper.ErrorPanic(result.Error)
	return tagList
}

// FindById implements TagRepository
func (t *TagRepositoryImpl) FindById(tagsId int) (tags entity.Tags, err error) {
	var tag entity.Tags
	result := t.Db.Find(&tag, tagsId)

	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// Save implements TagRepository
func (t *TagRepositoryImpl) Save(tags entity.Tags) entity.Tags {

	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
	return tags

}

// Update implements TagRepository
func (t *TagRepositoryImpl) Update(tags entity.Tags) entity.Tags {
	result := t.Db.Save(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}
