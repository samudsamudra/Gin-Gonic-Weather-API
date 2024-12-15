package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	dsn := "root@tcp(127.0.0.1:3306)/cuaca"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Error pinging the database:", err)
	}
	fmt.Println("Connected to MySQL database!")
}
