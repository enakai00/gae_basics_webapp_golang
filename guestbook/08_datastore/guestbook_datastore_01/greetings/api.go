package greetings

import (
	"net/http"
	"strconv"
	"time"

	"github.com/enakai00/gae_basics_webapp_golang/guestbook/08_datastore/guestbook_datastore_01/ds"
	"github.com/labstack/echo"
)

func Register(e *echo.Echo) {
	e.GET("/api/greetings", getAllGuests)
	e.POST("/api/greetings", addGuest)
	e.GET("/api/greetings/:id", getGuest)
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

// e.GET("/api/greetings/:id", getGuest)
func getGuest(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}
	igarashi := GuestData{
		ID:      int64(id),
		Name:    "Tuyushi Igarashi",
		Message: "Hello",
	}
	return c.JSON(http.StatusOK, igarashi)
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
	entity := ds.Insert(data.Name, data.Message)
	response := GuestData{
		Name:    entity.Name,
		Message: entity.Message,
		Created: entity.Created,
		ID:      entity.Key.ID,
	}
	return c.JSON(http.StatusCreated, response)
}
