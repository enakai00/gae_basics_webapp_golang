package greetings

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func Register(e *echo.Echo) {
	e.GET("/api/greetings", getAllGuests)
}

type GuestData struct {
	Name    string    `json:"author"`
	Message string    `json:"message"`
	Created time.Time `json:"created"`
	ID      int64     `json:"id"`
}

// e.GET("/api/greetings", getAllGuests)
func getAllGuests(c echo.Context) error {
	type response struct {
		Guests []GuestData `json:"greetings"`
	}

	igarashi := GuestData{
		ID:      1,
		Name:    "Tuyushi Igarashi",
		Message: "Hello",
	}
	miyayama := GuestData{
		ID:      2,
		Name:    "Ryutaro Miyayama",
		Message: "Looks good to me",
	}
	data := response{Guests: []GuestData{igarashi, miyayama}}
	return c.JSON(http.StatusOK, data)
}
