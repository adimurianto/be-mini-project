package controllers

import (
	"net/http"

	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"

	"github.com/gin-gonic/gin"
)

type SpecialDealsController struct{}

// @Tags			Special Deals
// @Produce			json
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/special-deals/ [get]
func (ctrl *SpecialDealsController) GetData(ctx *gin.Context) {
	var specialDeals []*models.SpecialDeals

	err := repository.Get(&specialDeals)
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
		Data:    &specialDeals,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @Tags			Special Deals
// @Produce			json
// @Security		BearerAuth
// @Param		special_deals	body		models.SpecialDealsBase	true	"Special Deals object to be created"
// @Success			201 {object} helpers.Response{}
// @Router			/api/v1/special-deals/ [post]
func (ctrl *SpecialDealsController) CreateData(ctx *gin.Context) {
	var body models.SpecialDealsBase

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newSpecialDeals models.SpecialDeals
	newSpecialDeals.Name = body.Name
	newSpecialDeals.Price = body.Price
	newSpecialDeals.Discount = body.Discount
	newSpecialDeals.Image = body.Image
	newSpecialDeals.Status = body.Status

	errSave := repository.Save(&newSpecialDeals)
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
		Data:    &newSpecialDeals,
		Message: "Success",
	}
	ctx.JSON(http.StatusCreated, webResponse)
}
