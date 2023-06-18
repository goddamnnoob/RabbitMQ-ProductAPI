package repositories

import (
	"database/sql"
	"goddamnnoob/RabbitMQ-ProductAPI/integrations"
	"goddamnnoob/RabbitMQ-ProductAPI/models"
)

func CheckUserValid(user *models.User) (isvalid bool, err error) {
	isvalid = true
	checkUserValidQuery := "SELECT user_id FROM users WHERE user_id = $1"
	connection, err := integrations.GetNewPostgresConnection()
	if err != nil {
		return isvalid, err
	}
	err = connection.QueryRow(checkUserValidQuery, user.Userid).Scan(&user.Userid)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return isvalid, err
	}
	defer connection.Close()
	return isvalid, nil
}
