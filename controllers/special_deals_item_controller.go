package controllers

import (
	"net/http"

	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"

	"github.com/gin-gonic/gin"
)

type SpecialDealsItemController struct{}

// @Tags			Special Deals Item
// @Produce			json
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/special-deals-item/ [get]
func (ctrl *SpecialDealsItemController) GetData(ctx *gin.Context) {
	var specialDealsItem []*models.SpecialDealsItem

	err := repository.Get(&specialDealsItem)
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
		Data:    &specialDealsItem,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @Tags			Special Deals Item
// @Produce			json
// @Security		BearerAuth
// @Param		special_deals_item	body		models.SpecialDealsItemBase	true	"Special Deals Item object to be created"
// @Success			201 {object} helpers.Response{}
// @Router			/api/v1/special-deals-item/ [post]
func (ctrl *SpecialDealsItemController) CreateData(ctx *gin.Context) {
	var body models.SpecialDealsItemBase

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newSpecialDealsItem models.SpecialDealsItem
	newSpecialDealsItem.ProductID = body.ProductID
	newSpecialDealsItem.SpecialDealsID = body.SpecialDealsID
	newSpecialDealsItem.Quantity = body.Quantity
	newSpecialDealsItem.Status = body.Status

	errSave := repository.Save(&newSpecialDealsItem)
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
		Data:    &newSpecialDealsItem,
		Message: "Success",
	}
	ctx.JSON(http.StatusCreated, webResponse)
}
