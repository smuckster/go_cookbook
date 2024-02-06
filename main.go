package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func connectionHandle() *sql.DB {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	return db
}

func handleIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index", struct{}{})
}

func main() {
	connectionHandle()

	// Register templates
	t := &Template{
		templates: template.Must(template.ParseGlob("index.html")),
	}

	e := echo.New()
	e.Use(
		middleware.Recover(),
		middleware.Logger(),
	)
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		fmt.Println(c.Path(), c.QueryParams(), err.Error())
		e.DefaultHTTPErrorHandler(err, c)
	}
	e.Static("/static", "assets")
	e.Renderer = t
	e.GET("/", handleIndex)
	e.Logger.Fatal(e.Start(":8080"))
}
