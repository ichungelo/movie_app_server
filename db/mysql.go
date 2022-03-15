package db

import (
	"database/sql"
	"fmt"
	"go_api_echo/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func Mysql(query string) (*sql.Rows, error) {
	var (
		host, _     = utils.DotEnvGenerator("DB_HOST")
		port, _     = utils.DotEnvGenerator("DB_PORT")
		user, _     = utils.DotEnvGenerator("DB_USER")
		password, _ = utils.DotEnvGenerator("DB_PASSWORD")
		database, _ = utils.DotEnvGenerator("DB_DATABASE")
	)

	mySqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", user, password, host, port, database)

	db, err := sql.Open("mysql", mySqlInfo)
	if err != nil {
		return nil, err
	}
	driver, _ := mysql.WithInstance(db, &mysql.Config{
		MigrationsTable: "users",
		DatabaseName:    "movie_app",
		NoLock:          true,
	})

	m, _ := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?query", user, password, host, port, database),
		"mysql",
		driver,
	)

	m.Steps(2)

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
