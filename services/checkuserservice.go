package services

import (
	"goddamnnoob/Zocket-assignment/integrations"
)

func CheckUserPresent(Userid int32) (ispresent bool, err error) {
	connection, er := integrations.GetNewPostgresConnection()
	if er != nil {
		return false, er
	}
	defer connection.Close()
	ispresent = true
	//TODO Check for user in DB
	return ispresent, nil
}
