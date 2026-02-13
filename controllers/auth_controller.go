package controllers

import (
	"be-mini-project/helpers"
	"be-mini-project/models"
	repository "be-mini-project/repositories"
	"be-mini-project/utils"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

// @Tags	Auth
// @Produce	json
// @Param	user	body		models.AuthRequest	true	"User object to be created"
// @Success	200		{object}	helpers.Response{}
// @Router	/api/v1/auth/login [post]
func (ctrl *AuthController) Login(ctx *gin.Context) {
	var body models.AuthRequest

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get data user by username
	result, err := repository.GetUserByUsername(body.Username)
	if err != nil {
		if customErr, ok := err.(*repository.CustomError); ok {
			ctx.JSON(customErr.Code, gin.H{"error": customErr.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Check input password
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(body.Password)); err != nil {
		webResponse := helpers.Response{
			Code:    http.StatusUnauthorized,
			Status:  false,
			Message: "Invalid credentials",
		}
		ctx.JSON(http.StatusUnauthorized, webResponse)
		return
	}

	// Retrieve token
	token, err := utils.GenerateToken(result.ID, result.Fullname, result.Role)
	if err != nil {
		webResponse := helpers.Response{
			Code:    http.StatusBadRequest,
			Status:  false,
			Message: "Failed to generate token",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	firstString := randomString(3)
	lastString := randomString(5)

	var response models.AuthResponse
	response.ID = result.ID
	response.Fullname = result.Fullname
	response.Username = result.Username
	response.Role = result.Role
	response.Token = firstString + token + lastString

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    response,
		Message: "Success",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
