package controllers

import (
	"net/http"
	"strconv"

	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

// @Tags			User
// @Produce			json
// @Security		BearerAuth
// @Param 			sort query string false "Sort >> column_name,order | order: pilih 'asc' atau 'desc'"
// @Param 			filter query string false "Filter >> [column_name,operator,value;.....]"
// @Param 			page query int false "Page"
// @Param 			perPage query int false "perPage"
// @Success			200 {object} helpers.Response{}
// @Router			/api/v1/user/ [get]
func (ctrl *UserController) GetData(ctx *gin.Context) {
	var user []*models.User
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

	model, tData, tPages, err := repository.GetWithFilter(&user, sort, filter, page, perPage)
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

// @Tags			User
// @Produce			json
// @Security		BearerAuth
// @Param		user	body		models.UserBase	true	"User object to be created"
// @Success			201 {object} helpers.Response{}
// @Router			/api/v1/user/ [post]
func (ctrl *UserController) CreateData(ctx *gin.Context) {
	var body models.UserBase

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	var newUser models.User
	newUser.Fullname = body.Fullname
	newUser.Username = body.Username
	newUser.Password = string(hashedPassword)
	newUser.Role = body.Role
	newUser.Status = body.Status

	errSave := repository.Save(&newUser)
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
		Data:    &newUser,
		Message: "Success",
	}
	ctx.JSON(http.StatusCreated, webResponse)
}

// @Tags		User
// @Produce		json
// @Security	BearerAuth
// @Param		user	body		models.User	true	"User object to be updated"
// @Success		200		{object}	helpers.Response{}
// @Router		/api/v1/user/ [put]
func (ctrl *UserController) UpdateData(ctx *gin.Context) {
	var body models.User

	if err := ctx.Bind(&body); err != nil {
		webResponse := helpers.Response{
			Code:    http.StatusBadRequest,
			Status:  true,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	var user models.User
	repository.GetById(&user, body.ID)

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Fullname = body.Fullname
	user.Username = body.Username
	user.Password = string(hashedPassword)
	user.Role = body.Role
	user.Status = body.Status

	repository.Update(&user)

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    &user,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// @Tags		User
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	true	"User ID"
// @Success		200		{object}	helpers.Response{}
// @Router		/api/v1/user/{id}  [delete]
func (ctrl *UserController) DeleteData(ctx *gin.Context) {
	var user models.User

	repository.GetById(&user, ctx.Param("id"))

	user.Status = false

	repository.Update(&user)

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    &user,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}
