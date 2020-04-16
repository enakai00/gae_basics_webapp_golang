package greetings

import (
	"net/http"
	"strconv"
	"time"

	"github.com/enakai00/gae_basics_webapp_golang/guestbook/10_gcs/guestbook_gcs_02/ds"
	"github.com/labstack/echo"
)

func Register(e *echo.Echo) {
	e.GET("/api/greetings", getAllGuests)
	e.POST("/api/greetings", addGuest)
	e.GET("/api/greetings/:id", getGuest)
	e.PUT("/api/greetings/:id", updateGuest)
	e.DELETE("/api/greetings/:id", deleteGuest)
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

	entities := ds.GetAll()
	guests := []GuestData{}
	for _, entity := range entities {
		item := GuestData{
			Name:    entity.Name,
			Message: entity.Message,
			Created: entity.Created,
			ID:      entity.Key.ID,
		}
		guests = append(guests, item)
	}
	return c.JSON(http.StatusOK, response{Guests: guests})
}

// e.GET("/api/greetings/:id", getGuest)
func getGuest(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}
	entity := ds.GetByID(int64(id))
	if entity == (ds.GuestEntity{}) {
		return echo.NewHTTPError(http.StatusNotFound, "No existing key")
	}
	greet := GuestData{
		Name:    entity.Name,
		Message: entity.Message,
		Created: entity.Created,
		ID:      entity.Key.ID,
	}
	return c.JSON(http.StatusOK, greet)
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

// e.PUT("/api/greetings/:id", updateGuest)
func updateGuest(c echo.Context) error {
	type postData struct {
		Name    string `json:"author" form:"author" query:"author"`
		Message string `json:"message" form:"message" query:"message"`
	}

	data := new(postData)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	entity := ds.GetByID(int64(id))
	if entity == (ds.GuestEntity{}) {
		return echo.NewHTTPError(http.StatusNotFound, "No existing key")
	}

	entity.Name = data.Name
	entity.Message = data.Message
	entity = ds.Update(entity)

	item := GuestData{
		Name:    data.Name,
		Message: data.Message,
		Created: entity.Created,
		ID:      entity.Key.ID,
	}
	return c.JSON(http.StatusOK, item)
}

// e.DELETE("/api/greetings/:id", deleteGuest)
func deleteGuest(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}
	ds.Delete(int64(id))
	return c.String(http.StatusNoContent, "")
}
