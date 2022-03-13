package repositories

import (
	"go_api_echo/db"
	"go_api_echo/entities"
	"go_api_echo/queries"
	"go_api_echo/transport"
)

func CreateUser(data transport.RegisterRequest) error {
	createUserQuery := queries.CreateUserQuery(data)
	_, err := db.Mysql(createUserQuery)
	if err != nil {
		return err
	}
	return err
}

func ReadUser(data transport.LoginRequest)(*entities.UserLoginEntity, error) {
	readUserQuery := queries.ReadUserQuery(data)
	results, err := db.Mysql(readUserQuery)
	if err != nil {
		return nil, err
	}
	result := new(entities.UserLoginEntity)
	for results.Next() {
		err := results.Scan(&result.UserId, &result.Username, &result.Email, &result.FirstName, &result.LastName, &result.Password, )
		if err != nil {
			return nil, err
		}
	}

	defer results.Close()
	return result, nil
}