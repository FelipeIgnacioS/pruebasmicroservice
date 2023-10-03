package controllers

import (
	"encoding/json"

	db "BACKEND-MICROSERVICE-AUTHENTICATION/config"
	"BACKEND-MICROSERVICE-AUTHENTICATION/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := db.DB.Select("id, first_name, last_name, email").Find(&users).Error
	return users, err
}

func GetUser(userID string) (models.User, error) {
	var user models.User
	err := db.DB.Select("id, first_name, last_name, email").Where("id = ?", userID).First(&user).Error

	return user, err
}

func GetUserByEmail(body []byte) (models.User, error) {
	var loginData models.LoginData

	err := json.Unmarshal(body, &loginData)
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = db.DB.Select("id, first_name, last_name, email").Where("email = ?", loginData.Email).First(&user).Error
	return user, err
}
