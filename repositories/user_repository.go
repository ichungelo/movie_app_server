package repositories

import (
	"go_api_echo/db"
	"go_api_echo/entities"
	"go_api_echo/queries"
	"go_api_echo/transport"
	"go_api_echo/utils"
)

func CreateUser(data transport.RegisterRequest) {
	createUserQuery := queries.CreateUserQuery(data)
	db.Mysql(createUserQuery)
}

func ReadUser(data transport.LoginRequest)(result entities.UserLoginEntity) {
	readUserQuery := queries.ReadUserQuery(data)
	results := db.Mysql(readUserQuery)
	
	for results.Next() {
		err := results.Scan(&result.UserId, &result.Username, &result.Email, &result.FirstName, &result.LastName, &result.Password, )
		utils.CheckError(err)
	}

	defer results.Close()
	return
}