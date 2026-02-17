package controllers

import (
	"net/http"

	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

// @Tags			Product
// @Produce			json
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/product/ [get]
func (ctrl *ProductController) GetData(ctx *gin.Context) {
	var product []*models.Product

	err := repository.Get(&product)
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
		Data:    &product,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @Tags			Product
// @Produce			json
// @Security		BearerAuth
// @Param		product	body		models.ProductBase	true	"Product object to be created"
// @Success			201 {object} helpers.Response{}
// @Router			/api/v1/product/ [post]
func (ctrl *ProductController) CreateData(ctx *gin.Context) {
	var body models.ProductBase

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newProduct models.Product
	newProduct.Name = body.Name
	newProduct.Price = body.Price
	newProduct.Image = body.Image
	newProduct.CategoryID = body.CategoryID
	newProduct.Description = body.Description
	newProduct.Status = body.Status

	errSave := repository.Save(&newProduct)
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
		Data:    &newProduct,
		Message: "Success",
	}
	ctx.JSON(http.StatusCreated, webResponse)
}
