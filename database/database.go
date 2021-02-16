package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "", "localhost", 3306, "test_nodejs")
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to database opened!")

	return db
}
