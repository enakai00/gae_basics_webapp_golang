package greetings

import (
	"net/http"

	"github.com/labstack/echo"
)

func Register(e *echo.Echo) {
	e.GET("/api/greetings", greetings)
}

type userData struct {
	ID      int    `json:"id"`
	Name    string `json:"author"`
	Message string `json:"message"`
}

type response struct {
	Greetings []userData `json:"greetings"`
}

// e.GET("/api/greetings", home)
func greetings(c echo.Context) error {
	igarashi := userData{
		ID:      1,
		Name:    "Tuyushi Igarashi",
		Message: "Hello",
	}
	miyayama := userData{
		ID:      2,
		Name:    "Ryutaro Miyayama",
		Message: "Looks good to me",
	}
	data := response{Greetings: []userData{igarashi, miyayama}}
	return c.JSON(http.StatusOK, data)
}
