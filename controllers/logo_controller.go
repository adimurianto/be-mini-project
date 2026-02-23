package controllers

import (
	"net/http"
	"strconv"

	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"

	"github.com/gin-gonic/gin"
)

type LogoController struct{}

// @Tags			Logo
// @Produce			json
// @Param 			sort query string false "Sort >> column_name,order | order: pilih 'asc' atau 'desc'"
// @Param 			filter query string false "Filter >> [column_name,operator,value;.....]"
// @Param 			page query int false "Page"
// @Param 			perPage query int false "perPage"
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/logo/ [get]
func (ctrl *LogoController) GetData(ctx *gin.Context) {
	var logo []*models.Logo
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

	model, tData, tPages, err := repository.GetWithFilter(&logo, sort, filter, page, perPage)
	if err != nil {
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

// @Tags			Logo
// @Produce			json
// @Security		BearerAuth
// @Param		logo	body		models.LogoBase	true	"Logo object to be created"
// @Success			201 {object} helpers.Response{}
// @Router			/api/v1/logo/ [post]
func (ctrl *LogoController) CreateData(ctx *gin.Context) {
	var body models.LogoBase

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newLogo models.Logo
	newLogo.Title = body.Title
	newLogo.Logo = body.Logo
	newLogo.Status = body.Status

	errSave := repository.Save(&newLogo)
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
		Data:    &newLogo,
		Message: "Success",
	}
	ctx.JSON(http.StatusCreated, webResponse)
}

// @Tags		Logo
// @Produce		json
// @Security	BearerAuth
// @Param		logo	body		models.Logo	true	"Logo object to be updated"
// @Success		200		{object}	helpers.Response{}
// @Router		/api/v1/logo/ [put]
func (ctrl *LogoController) UpdateData(ctx *gin.Context) {
	var body models.Logo

	if err := ctx.Bind(&body); err != nil {
		webResponse := helpers.Response{
			Code:    http.StatusBadRequest,
			Status:  true,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	var logo models.Logo
	repository.GetById(&logo, body.ID)

	logo.Title = body.Title
	logo.Logo = body.Logo
	logo.Status = body.Status

	repository.Update(&logo)

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    &logo,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @Tags		Logo
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	true	"Logo ID"
// @Success		200		{object}	helpers.Response{}
// @Router		/api/v1/logo/{id}  [delete]
func (ctrl *LogoController) DeleteData(ctx *gin.Context) {
	var logo models.Logo

	repository.GetById(&logo, ctx.Param("id"))

	logo.Status = false

	repository.Update(&logo)

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    &logo,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}
