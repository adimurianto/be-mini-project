package controllers

import (
	"net/http"
	"strconv"

	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

// @Tags			Category
// @Produce			json
// @Param 			sort query string false "Sort >> column_name,order | order: pilih 'asc' atau 'desc'"
// @Param 			filter query string false "Filter >> [column_name,operator,value;.....]"
// @Param 			page query int false "Page"
// @Param 			perPage query int false "perPage"
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/category/ [get]
func (ctrl *CategoryController) GetData(ctx *gin.Context) {
	var category []*models.Category

	// query parameters
	sort := ctx.Query("sort")
	filter := ctx.Query("filter")
	page, _ := strconv.Atoi(ctx.Query("page"))
	perPage, _ := strconv.Atoi(ctx.Query("perPage"))

	// default perPage
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}

	model, tData, tPages, err := repository.GetWithFilter(&category, sort, filter, page, perPage)
	if err != nil {
		// Handle error
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Status:  false,
			Message: "Error fetching data",
		})
		return
	}

	webResponse := helpers.Response{
		Code:   http.StatusOK,
		Status: true,
		Info: helpers.Info{
			Page:       page,
			PerPage:    perPage,
			TotalPages: tPages,
			TotalData:  tData,
		},
		Data:    &model,
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

// @Tags		Category
// @Produce		json
// @Security	BearerAuth
// @Param		category	body		models.Category	true	"Category object to be updated"
// @Success		200		{object}	helpers.Response{}
// @Router		/api/v1/category/ [put]
func (ctrl *CategoryController) UpdateData(ctx *gin.Context) {
	var body models.Category

	if err := ctx.Bind(&body); err != nil {
		webResponse := helpers.Response{
			Code:    http.StatusBadRequest,
			Status:  true,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	var category models.Category
	repository.GetById(&category, body.ID)

	category.Name = body.Name
	category.Icon = body.Icon
	category.Status = body.Status

	repository.Update(&category)

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    &category,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @Tags		Category
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	true	"Category ID"
// @Success		200		{object}	helpers.Response{}
// @Router		/api/v1/category/{id}  [delete]
func (ctrl *CategoryController) DeleteData(ctx *gin.Context) {
	var category models.Category

	repository.GetById(&category, ctx.Param("id"))

	category.Status = false

	repository.Update(&category)

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    &category,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}
