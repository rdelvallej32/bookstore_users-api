package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MYSQL_USERNAME = "MYSQL_USERNAME"
	MYSQL_PASSWORD = "MYSQL_PASSWORD"
	MYSQL_HOST     = "MYSQL_HOST"
	MYSQL_SCHEMA   = "MYSQL_SCHEMA"
)

var (
	Client   *sql.DB
	username = os.Getenv(MYSQL_USERNAME)
	password = os.Getenv(MYSQL_PASSWORD)
	host     = os.Getenv(MYSQL_HOST)
	schema   = os.Getenv(MYSQL_SCHEMA)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema)

	var err error

	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
