package services

import (
	"goddamnnoob/RabbitMQ-ProductAPI/models"
	"goddamnnoob/RabbitMQ-ProductAPI/repositories"
)

func CheckUserValid(Userid int32) (isvalid bool, err error) {
	user := models.User{Userid: Userid}
	isvalid, err = repositories.CheckUserValid(&user)
	if err != nil {
		return false, err
	}
	return isvalid, nil
}
