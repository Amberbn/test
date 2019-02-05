package main

import (
	// "database/sql"
	"net/http"
	"github.com/labstack/echo"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	con "test/db"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})

	con.Connection()
	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3360)/test")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("db is connected")
	// }
	// defer db.Close()
	// // make sure connection is available
	// err = db.Ping()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	e.Logger.Fatal(e.Start(":1323"))
}