package main

import (
	"api-test/db"
	"database/sql"
)

func main() {
	defer func(SqlDB *sql.DB) {
		err := SqlDB.Close()
		if err != nil {

		}
	}(db.SqlDB)
	router := initRouter()
	err := router.Run(":8806")
	if err != nil {
		return
	}
}
