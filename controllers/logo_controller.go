package controllers

import (
	"net/http"

	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"

	"github.com/gin-gonic/gin"
)

type LogoController struct{}

// @Tags			Logo
// @Produce			json
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/logo/ [get]
func (ctrl *LogoController) GetData(ctx *gin.Context) {
	var logo []*models.Logo

	err := repository.Get(&logo)
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
		Data:    &logo,
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
