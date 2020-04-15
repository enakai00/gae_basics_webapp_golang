package greetings

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func Register(e *echo.Echo) {
	e.GET("/api/greetings", getAllGuests)
	e.POST("/api/greetings", addGuest)
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

// e.POST("/api/greetings", addGuest)
func addGuest(c echo.Context) error {
	type postData struct {
		Name    string `json:"author" form:"author" query:"author"`
		Message string `json:"message" form:"message" query:"message"`
	}

	data := new(postData)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	response := GuestData{
		ID:      999,
		Name:    data.Name,
		Message: data.Message,
	}
	return c.JSON(http.StatusCreated, response)
}
