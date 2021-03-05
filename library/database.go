package library

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3305)/think2")
	if err != nil {
		fmt.Println("Connect database error")
	}

	return db
}