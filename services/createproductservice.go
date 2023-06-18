package services

import (
	"goddamnnoob/RabbitMQ-ProductAPI/models"
	"goddamnnoob/RabbitMQ-ProductAPI/repositories"
	"time"

	"github.com/google/uuid"
)

func AddProduct(product *models.Product, user *models.User) (result string, err error) {

	isvalid, err := CheckUserValid(user.Userid) // Check whethre the user_id is valid
	if err != nil {
		return "Error while checking validity of user", err
	}
	if !isvalid {
		return "User is not a valid user please check the user id again", nil
	}
	product.Product_id, err = uuid.NewRandom()
	if err != nil {
		return "Error while getting new UUID", err
	}

	product.Createdat = time.Now()
	product.Updatedat = time.Now()
	err = repositories.AddProduct(product) // Add Product to DB
	if err != nil {
		return "Error while adding product to the database", err
	}

	err = AddProductToMessageQueue(product.Product_id)
	if err != nil {
		return "Error While adding data to message queue ", err
	}

	return "Product Added Successfully ", nil

}
