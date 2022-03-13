package db

import (
	"database/sql"
	"fmt"
	"go_api_echo/utils"

	_ "github.com/go-sql-driver/mysql"
)

func Mysql(query string) (*sql.Rows, error) {
	var (
		host, _     = utils.DotEnvGenerator("DB_HOST")
		port, _     = utils.DotEnvGenerator("DB_PORT")
		user, _     = utils.DotEnvGenerator("DB_USER")
		password, _ = utils.DotEnvGenerator("DB_PASSWORD")
		database, _ = utils.DotEnvGenerator("DB_DATABASE")
	)

	mySqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	db, err := sql.Open("mysql", mySqlInfo)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	insert, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return insert, nil
}
