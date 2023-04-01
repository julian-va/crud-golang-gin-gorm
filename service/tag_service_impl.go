package service

import (
	"crud-golang-gin-gorm/data/dto"
	"crud-golang-gin-gorm/data/repository"
	"crud-golang-gin-gorm/helper"
)

type TagsServiceImpl struct {
	TagRepository repository.TagRepository
}

// Delete implements TagsService
func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagRepository.Delete(tagsId)
}

// FindAll implements TagsService
func (t *TagsServiceImpl) FindAll() []dto.TagDto {
	result := t.TagRepository.FindAll()
	return helper.EntityTagToDtoTagList(result)
}

// FindById implements TagsService
func (t *TagsServiceImpl) FindById(tagsId int) (tags dto.TagDto, err error) {
	result, err := t.TagRepository.FindById(tagsId)
	helper.ErrorPanic(err)
	return helper.EntityTagToDtoTag(result), err
}

// Save implements TagsService
func (t *TagsServiceImpl) Save(tags dto.TagDto) dto.TagDto {
	result := t.TagRepository.Save(helper.DtoTagToEntityTag(tags))
	return helper.EntityTagToDtoTag(result)
}

// Update implements TagsService
func (t *TagsServiceImpl) Update(tags dto.TagDto) dto.TagDto {
	result := t.TagRepository.Save(helper.DtoTagToEntityTag(tags))
	return helper.EntityTagToDtoTag(result)
}

func NewTagsServiceImpl(TagRepository repository.TagRepository) TagsService {
	return &TagsServiceImpl{
		TagRepository: TagRepository,
	}
}
