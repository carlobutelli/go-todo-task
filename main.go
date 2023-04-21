// Main.go
package main

import (
	"github.com/labstack/echo"
	"net/http"
	_ "github.com/lib/pq"

	"fmt"
	"database/sql"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "task"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
	//db := initDB("storage.db")
	//migrate(db)
}
func initDB(filepath string) interface{} {

}

func todo() {
	e := echo.New()

	e.GET("/", hello)
	e.GET("/tasks", func(c echo.Context) error {
		return c.JSON(200, "GET Tasks")
	})
	e.PUT("/tasks", func(c echo.Context) error {
		return c.JSON(200, "PUT Tasks")
	})
	e.DELETE("/tasks/:id", func(c echo.Context) error {
		return c.JSON(200, "DELETE Task" + c.Param("id"))
	})
	e.Logger.Fatal(e.Start(":8000"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}