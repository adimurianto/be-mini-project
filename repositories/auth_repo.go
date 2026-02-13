package repository

import (
	"be-mini-project/infra/database"
	"be-mini-project/models"
)

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	database.DB = database.DB.Debug()
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
