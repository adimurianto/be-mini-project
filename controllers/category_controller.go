package controllers

import (
	"net/http"

	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

// @Tags			Category
// @Produce			json
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/category/ [get]
func (ctrl *CategoryController) GetData(ctx *gin.Context) {
	var category []*models.Category

	err := repository.Get(&category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Status:  false,
			Message: "Error fetching data",
		})
		return
	}

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    &category,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @Tags			Category
// @Produce			json
// @Security		BearerAuth
// @Param		category	body		models.CategoryBase	true	"Category object to be created"
// @Success			201 {object} helpers.Response{}
// @Router			/api/v1/category/ [post]
func (ctrl *CategoryController) CreateData(ctx *gin.Context) {
	var body models.CategoryBase

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newCategory models.Category
	newCategory.Name = body.Name
	newCategory.Icon = body.Icon
	newCategory.Status = body.Status

	errSave := repository.Save(&newCategory)
	if errSave != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Status:  false,
			Message: "Error creating data",
		})
		return
	}

	webResponse := helpers.Response{
		Code:    http.StatusCreated,
		Status:  true,
		Data:    &newCategory,
		Message: "Success",
	}
	ctx.JSON(http.StatusCreated, webResponse)
}
