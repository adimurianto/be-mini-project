package controllers

import (
	"net/http"
	"strconv"

	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"

	"github.com/gin-gonic/gin"
)

type BannerController struct{}

// @Tags			Banner
// @Produce			json
// @Param 			sort query string false "Sort >> column_name,order | order: pilih 'asc' atau 'desc'"
// @Param 			filter query string false "Filter >> [column_name,operator,value;.....]"
// @Param 			page query int false "Page"
// @Param 			perPage query int false "perPage"
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/banner/ [get]
func (ctrl *BannerController) GetData(ctx *gin.Context) {
	var banner []*models.Banner

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

	model, tData, tPages, err := repository.GetWithFilter(&banner, sort, filter, page, perPage)
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

// @Tags		Banner
// @Produce		json
// @Security	BearerAuth
// @Param		banner	body		models.Banner	true	"Banner object to be updated"
// @Success		200		{object}	helpers.Response{}
// @Router		/api/v1/banner/ [put]
func (ctrl *BannerController) UpdateData(ctx *gin.Context) {
	var body models.Banner

	if err := ctx.Bind(&body); err != nil {
		webResponse := helpers.Response{
			Code:    http.StatusBadRequest,
			Status:  true,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	var banner models.Banner
	repository.GetById(&banner, body.ID)

	banner.Title = body.Title
	banner.Link = body.Link
	banner.PrimaryImage = body.PrimaryImage
	banner.SecondaryImage = body.SecondaryImage
	banner.Status = body.Status

	repository.Update(&banner)

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    &banner,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @Tags		Banner
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	true	"Banner ID"
// @Success		200		{object}	helpers.Response{}
// @Router		/api/v1/banner/{id}  [delete]
func (ctrl *BannerController) DeleteData(ctx *gin.Context) {
	var banner models.Banner

	repository.GetById(&banner, ctx.Param("id"))

	banner.Status = false

	repository.Update(&banner)

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    &banner,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}
