package db

import (
	"database/sql"
	"fmt"
	"go_api_echo/utils"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var Db *sql.DB

func InitDB() {
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
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	Db = db
	log.Println("Connected to database")
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		panic(err.Error())
	}

	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err.Error())
	}
}

func Mysql(query string) (*sql.Rows, error) {

	insert, err := Db.Query(query)
	if err != nil {
		return nil, err
	}

	return insert, nil
}
