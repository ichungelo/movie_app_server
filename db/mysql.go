package db

import (
	"database/sql"
	"fmt"
	"go_api_echo/utils"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Mysql(query string) (result *sql.Rows) {
	var (
		host     = utils.DotEnvGenerator("DB_HOST")
		port     = utils.DotEnvGenerator("DB_PORT")
		user     = utils.DotEnvGenerator("DB_USER")
		password = utils.DotEnvGenerator("DB_PASSWORD")
		database = utils.DotEnvGenerator("DB_DATABASE")
	)

	mySqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port,database)

	db, err := sql.Open("mysql", mySqlInfo)
	utils.CheckError(err)

	defer db.Close()

	err = db.Ping()
	utils.CheckError(err)
	log.Println("DB Connected")

	insert, err := db.Query(query)
	utils.CheckError(err)

	// defer insert.Close()
	return insert
}
