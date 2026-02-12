package repository

import (
	"be-mini-project/infra/database"
)

func Get(model interface{}) error {
	database.DB = database.DB.Debug()
	err := database.DB.Find(model).Error
	if err != nil {
		return err
	}
	return nil
}

func Save(model interface{}) interface{} {
	err := database.DB.Debug().Create(model).Error
	if err != nil {
		return err
	}
	return nil
}
