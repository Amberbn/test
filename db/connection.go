package db

import (
	 "database/sql"
	 "fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3360)/test")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		// fmt.Println("test")
	}
}
