package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

var e = createMux()

func createMux() *echo.Echo {
	e := echo.New()
	http.Handle("/", e)
	return e
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func init() {
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	e.GET("/", home)
}

func home(c echo.Context) error {
	type Data struct {
		Message string
	}
	data := Data{Message: "App Engine 勉強会 にようこそ"}
	return c.Render(http.StatusOK, "greeting", data)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
