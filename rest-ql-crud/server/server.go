package main // import "server"

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"dbbooks"
)

const table = "books"

var dbname = "./ql.db"

// ResultResponse : Create, Read 결과 반환용
type ResultResponse struct{ Message string }

func index(c echo.Context) error {
	data, _ := dbbooks.SelectData(0, table, dbname)

	return c.JSON(http.StatusOK, data)
}

func create(c echo.Context) error {
	var book dbbooks.Book
	book.Title = c.FormValue("Title")
	book.Author = c.FormValue("Author")

	dbbooks.InsertData(&book, table, dbname)
	return c.JSON(http.StatusOK, &ResultResponse{Message: "Create done"})
}

func read(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	data, _ := dbbooks.SelectData(uint(id), table, dbname)

	return c.JSON(http.StatusOK, data)
}

func update(c echo.Context) error {
	b := new(dbbooks.Book)
	if err := c.Bind(b); err != nil {
		panic(err)
	}

	dbbooks.UpdateData(b, table, dbname)

	return c.JSON(http.StatusOK, &ResultResponse{Message: "Update done"})
}

func delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	dbbooks.DeleteData(uint(id), table, dbname)
	return c.JSON(http.StatusOK, &ResultResponse{Message: "Delete done"})
}

func main() {
	if electronSite := os.Getenv("ELECTRONVUESITE"); len(electronSite) > 3 {
		dbname = electronSite + "/ql.db"
	}

	echo.NotFoundHandler = func(c echo.Context) error {
		errorResult := &ResultResponse{
			Message: "Contents not found",
		}
		return c.JSON(http.StatusNotFound, errorResult)
	}

	e := echo.New()

	e.Use(middleware.CORS())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*", "http://localhost:8080"},
	// 	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	// }))

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/books", index)

	e.POST("/books", create)
	e.GET("/books/:id", read)
	e.PUT("/books", update)
	e.DELETE("/books/:id", delete)

	dbbooks.CreateTable("books", dbname)

	e.Logger.Fatal(e.Start("127.0.0.1:1323")) // Electron에 붙일 때 방화벽 경고가 안 뜬다.
	// e.Logger.Fatal(e.Start(":1323")) // 0.0.0.0 은 Electron spawn시 방화벽 경고가 뜬다.
	// e.Logger.Fatal(e.Start(":80")) // 80포트 사용 가능. 안되면 다른 웹서버가 돌고 있는지 확인한다.
}
