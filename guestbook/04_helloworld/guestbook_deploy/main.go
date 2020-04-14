package main

import (
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

func init() {
	e.GET("/", home)
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
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
