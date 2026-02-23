package controllers

import (
	"net/http"
	"strconv"

	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

// @Tags			Product
// @Produce			json
// @Param 			sort query string false "Sort >> column_name,order | order: pilih 'asc' atau 'desc'"
// @Param 			filter query string false "Filter >> [column_name,operator,value;.....]"
// @Param 			page query int false "Page"
// @Param 			perPage query int false "perPage"
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/product/ [get]
func (ctrl *ProductController) GetData(ctx *gin.Context) {
	var product []*models.Product
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

	model, tData, tPages, err := repository.GetWithFilter(&product, sort, filter, page, perPage)
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

// @Tags		Product
// @Produce		json
// @Security	BearerAuth
// @Param		product	body		models.Product	true	"Product object to be updated"
// @Success		200		{object}	helpers.Response{}
// @Router		/api/v1/product/ [put]
func (ctrl *ProductController) UpdateData(ctx *gin.Context) {
	var body models.Product

	if err := ctx.Bind(&body); err != nil {
		webResponse := helpers.Response{
			Code:    http.StatusBadRequest,
			Status:  true,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	var product models.Product
	repository.GetById(&product, body.ID)

	product.Name = body.Name
	product.Price = body.Price
	product.Image = body.Image
	product.CategoryID = body.CategoryID
	product.Description = body.Description
	product.Status = body.Status

	repository.Update(&product)

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    &product,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @Tags		Product
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	true	"Product ID"
// @Success		200		{object}	helpers.Response{}
// @Router		/api/v1/product/{id}  [delete]
func (ctrl *ProductController) DeleteData(ctx *gin.Context) {
	var product models.Product

	repository.GetById(&product, ctx.Param("id"))

	product.Status = false

	repository.Update(&product)

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    &product,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}
