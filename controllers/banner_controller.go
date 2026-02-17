package controllers

import (
	"net/http"

	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"

	"github.com/gin-gonic/gin"
)

type BannerController struct{}

// @Tags			Banner
// @Produce			json
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/banner/ [get]
func (ctrl *BannerController) GetData(ctx *gin.Context) {
	var banner []*models.Banner

	err := repository.Get(&banner)
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
		Data:    &banner,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @Tags			Banner
// @Produce			json
// @Security		BearerAuth
// @Param		banner	body		models.BannerBase	true	"Banner object to be created"
// @Success			201 {object} helpers.Response{}
// @Router			/api/v1/banner/ [post]
func (ctrl *BannerController) CreateData(ctx *gin.Context) {
	var body models.BannerBase

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newBanner models.Banner
	newBanner.Title = body.Title
	newBanner.Link = body.Link
	newBanner.PrimaryImage = body.PrimaryImage
	newBanner.SecondaryImage = body.SecondaryImage
	newBanner.Status = body.Status

	errSave := repository.Save(&newBanner)
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
		Data:    &newBanner,
		Message: "Success",
	}
	ctx.JSON(http.StatusCreated, webResponse)
}
