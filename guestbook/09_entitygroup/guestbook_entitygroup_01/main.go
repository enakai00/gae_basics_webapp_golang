package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/enakai00/gae_basics_webapp_golang/guestbook/09_entitygroup/guestbook_entitygroup_01/comments"
	"github.com/enakai00/gae_basics_webapp_golang/guestbook/09_entitygroup/guestbook_entitygroup_01/greetings"
	"github.com/enakai00/gae_basics_webapp_golang/guestbook/09_entitygroup/guestbook_entitygroup_01/handler"
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
	handler.Register(e)
	greetings.Register(e)
	comments.Register(e)

	e.GET("/", home)
	e.GET("/err500", err500)
}

// e.GET("/", home)
func home(c echo.Context) error {
	type Data struct {
		Message string
	}
	data := Data{Message: "App Engine 勉強会 にようこそ"}
	return c.Render(http.StatusOK, "index", data)
}

// e.GET("/err500", err500)
func err500(c echo.Context) error {
	return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
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
