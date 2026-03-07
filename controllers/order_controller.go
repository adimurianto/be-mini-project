package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"
)

type OrderController struct{}

// @Tags			Order
// @Produce			json
// @Param 			sort query string false "Sort >> column_name,order | order: pilih 'asc' atau 'desc'"
// @Param 			filter query string false "Filter >> [column_name,operator,value;.....]"
// @Param 			page query int false "Page"
// @Param 			perPage query int false "perPage"
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/order/ [get]
func (ctrl *OrderController) GetData(ctx *gin.Context) {
	var orders []models.OrderResponse
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

	// Get data with filter, sort, and pagination
	model, totalData, totalPages, err := repository.GetWithFilter(&orders, sort, filter, page, perPage, "OrderDetails", "OrderDetails.Product")
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
			TotalPages: totalPages,
			TotalData:  totalData,
		},
		Data:    model,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @Tags			Order
// @Produce			json
// @Param			order	body		models.OrderRequest	true	"Order object to be created"
// @Success			201 {object} helpers.Response{}
// @Router			/api/v1/order/ [post]
func (ctrl *OrderController) CreateData(ctx *gin.Context) {
	var request models.OrderRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Response{
			Code:    http.StatusBadRequest,
			Status:  false,
			Message: "Invalid request body",
		})
		return
	}

	// Start transaction
	tx := repository.BeginTransaction()

	// Create new order
	var newOrder models.Order
	newOrder.Table = request.Table
	newOrder.Fullname = request.Fullname
	newOrder.Phone = request.Phone
	newOrder.IpAddress = request.IpAddress
	newOrder.TotalPrice = request.TotalPrice

	// Save order
	if err := tx.Create(&newOrder).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Status:  false,
			Message: "Error creating order",
			Info:    err.Error(),
		})
		return
	}

	// Create order details
	for _, detail := range request.OrderDetails {
		var orderDetail models.OrderDetail
		orderDetail.OrderID = newOrder.ID
		orderDetail.ProductID = detail.ProductID
		orderDetail.Quantity = detail.Quantity
		orderDetail.Price = detail.Price

		if err := tx.Create(&orderDetail).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, helpers.Response{
				Code:    http.StatusInternalServerError,
				Status:  false,
				Message: "Error creating order details",
				Info:    err.Error(),
			})
			return
		}
	}

	// Commit transaction
	tx.Commit()

	// Load order details for response
	repository.GetById(&newOrder, newOrder.ID)

	webResponse := helpers.Response{
		Code:    http.StatusCreated,
		Status:  true,
		Data:    newOrder,
		Message: "Success",
	}
	ctx.JSON(http.StatusCreated, webResponse)
}

// @Tags			Order
// @Produce			json
// @Security		BearerAuth
// @Param			id	path		string	true	"Order ID"
// @Param			order	body		models.OrderBase	true	"Order object to be updated"
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/order/{id} [put]
func (ctrl *OrderController) UpdateData(ctx *gin.Context) {
	id := ctx.Param("id")

	// Validate UUID
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Response{
			Code:    http.StatusBadRequest,
			Status:  false,
			Message: "Invalid order ID",
		})
		return
	}

	var body models.OrderBase
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Response{
			Code:    http.StatusBadRequest,
			Status:  false,
			Message: "Invalid request body",
		})
		return
	}

	// Check if order exists
	var existingOrder models.Order
	if err := repository.GetById(&existingOrder, id); err != nil {
		ctx.JSON(http.StatusNotFound, helpers.Response{
			Code:    http.StatusNotFound,
			Status:  false,
			Message: "Order not found",
		})
		return
	}

	// Update order fields
	existingOrder.Table = body.Table
	existingOrder.Fullname = body.Fullname
	existingOrder.Phone = body.Phone
	existingOrder.TotalPrice = body.TotalPrice
	existingOrder.StatusPayment = body.StatusPayment
	existingOrder.Status = body.Status

	// Save updates
	if err := repository.Update(&existingOrder); err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Status:  false,
			Message: "Error updating data",
		})
		return
	}

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    existingOrder,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @Tags			Order
// @Produce			json
// @Security		BearerAuth
// @Param			id	path		string	true	"Order ID"
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/order/{id} [delete]
func (ctrl *OrderController) DeleteData(ctx *gin.Context) {
	id := ctx.Param("id")

	// Validate UUID
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Response{
			Code:    http.StatusBadRequest,
			Status:  false,
			Message: "Invalid order ID",
		})
		return
	}

	// Check if order exists
	var existingOrder models.Order
	if err := repository.GetById(&existingOrder, id); err != nil {
		ctx.JSON(http.StatusNotFound, helpers.Response{
			Code:    http.StatusNotFound,
			Status:  false,
			Message: "Order not found",
		})
		return
	}

	// Soft delete - set status to false
	existingOrder.Status = false
	if err := repository.Update(&existingOrder); err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Status:  false,
			Message: "Error deleting data",
		})
		return
	}

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Order deleted successfully",
	}
	ctx.JSON(http.StatusOK, webResponse)
}
