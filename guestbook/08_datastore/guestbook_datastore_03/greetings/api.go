package greetings

import (
	"gae_basics_webapp_golang/guestbook/06_echo/guestbook_echo_01/ds"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func Register(e *echo.Echo) {
	e.GET("/api/greetings", greetings)
	e.GET("/api/greetings/:id", greetingsWithId)
	e.POST("/api/greetings", addUser)
}

type Greeting struct {
	Name    string    `json:"author"`
	Message string    `json:"message"`
	Created time.Time `json:"created"`
	ID      int64     `json:"id"`
}

// e.GET("/api/greetings", greetings)
func greetings(c echo.Context) error {
	type response struct {
		Greetings []Greeting `json:"greetings"`
	}

	entities := ds.GetAll()
	greetings := []Greeting{}
	for _, entity := range entities {
		greet := Greeting{
			Name:    entity.Name,
			Message: entity.Message,
			Created: entity.Created,
			ID:      entity.Key.ID,
		}
		greetings = append(greetings, greet)
	}
	return c.JSON(http.StatusOK, response{Greetings: greetings})
}

// e.GET("/api/greetings/:id", greetingsWithId)
func greetingsWithId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}
	entity := ds.GetByID(int64(id))
	if entity == (ds.MessageEntity{}) {
		return echo.NewHTTPError(http.StatusNotFound, "No existing key")
	}
	greet := Greeting{}
	if entity != (ds.MessageEntity{}) {
		greet = Greeting{
			Name:    entity.Name,
			Message: entity.Message,
			Created: entity.Created,
			ID:      entity.Key.ID,
		}
	}
	return c.JSON(http.StatusOK, greet)
}

// e.POST("/api/greetings", addUser)
func addUser(c echo.Context) (err error) {
	type postData struct {
		Name    string `json:"author" form:"author" query:"author"`
		Message string `json:"message" form:"message" query:"message"`
	}

	user := new(postData)
	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}
	entity := ds.Insert(user.Name, user.Message)
	response := Greeting{
		Name:    entity.Name,
		Message: entity.Message,
		Created: entity.Created,
		ID:      entity.Key.ID,
	}
	return c.JSON(http.StatusOK, response)
}
