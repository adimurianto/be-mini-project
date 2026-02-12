package controllers

import (
	"net/http"

	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// @Tags			User
// @Produce			json
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/user/ [get]
func (ctrl *UserController) GetData(ctx *gin.Context) {
	var user []*models.User

	err := repository.Get(&user)
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
		Data:    &user,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @Tags			User
// @Produce			json
// @Param		user	body		models.UserBase	true	"User object to be created"
// @Success			201 {object} helpers.Response{}
// @Router			/api/v1/user/ [post]
func (ctrl *UserController) CreateData(ctx *gin.Context) {
	var body models.UserBase

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newUser models.User
	newUser.Fullname = body.Fullname
	newUser.Username = body.Username
	newUser.Password = body.Password
	newUser.Role = body.Role
	newUser.Status = body.Status

	err := repository.Save(&newUser)
	if err != nil {
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
		Data:    &newUser,
		Message: "Success",
	}
	ctx.JSON(http.StatusCreated, webResponse)
}
