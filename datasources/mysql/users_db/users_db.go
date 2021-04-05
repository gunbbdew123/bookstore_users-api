package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Client *sql.DB
)

func init() {
	Loaderr := godotenv.Load("environment_variable.env")
	if Loaderr != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("mysql_users_username")
	password := os.Getenv("mysql_users_password")
	host := os.Getenv("mysql_users_host")
	schema := os.Getenv("mysql_users_schema")

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
