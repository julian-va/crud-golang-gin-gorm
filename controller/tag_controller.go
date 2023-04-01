package controller

import (
	"crud-golang-gin-gorm/data/dto"
	"crud-golang-gin-gorm/helper"
	"crud-golang-gin-gorm/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagService service.TagsService
}

func NewTagController(service service.TagsService) *TagController {
	return &TagController{
		tagService: service,
	}
}

func (controller *TagController) Create(ctx *gin.Context) {
	var dtoTag dto.TagDto
	err := ctx.ShouldBindJSON(&dtoTag)
	helper.ErrorPanic(err)
	ctx.JSON(http.StatusCreated, controller.tagService.Save(dtoTag))
}

func (controller *TagController) FindAll(ctx *gin.Context) {
	dtoList := controller.tagService.FindAll()
	if len(dtoList) > 0 {
		ctx.JSON(http.StatusOK, dtoList)
	} else {
		ctx.JSON(http.StatusNotFound, "Not Found")
	}
}
