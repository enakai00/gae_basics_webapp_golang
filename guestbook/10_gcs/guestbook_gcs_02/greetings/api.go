package greetings

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gae_basics_webapp_golang/guestbook/10_gcs/guestbook_gcs_02/ds"
	"github.com/labstack/echo"
)

func Register(e *echo.Echo) {
	e.GET("/api/greetings", getAllMessages)
	e.POST("/api/greetings", addMessage)
	e.GET("/api/greetings/:id", getMessage)
	e.PUT("/api/greetings/:id", updateMessage)
	e.DELETE("/api/greetings/:id", deleteMessage)
}

type MessageData struct {
	Name    string    `json:"author"`
	Message string    `json:"message"`
	Created time.Time `json:"created"`
	ID      int64     `json:"id"`
}

// e.GET("/api/greetings", getAllMessages)
func getAllMessages(c echo.Context) error {
	type response struct {
		Messages []MessageData `json:"greetings"`
	}

	entities := ds.GetAll()
	messages := []MessageData{}
	for _, entity := range entities {
		item := MessageData{
			Name:    entity.Name,
			Message: entity.Message,
			Created: entity.Created,
			ID:      entity.Key.ID,
		}
		messages = append(messages, item)
	}
	return c.JSON(http.StatusOK, response{Messages: messages})
}

// e.GET("/api/greetings/:id", getMessage)
func getMessage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}
	entity := ds.GetByID(int64(id))
	if entity == (ds.MessageEntity{}) {
		return echo.NewHTTPError(http.StatusNotFound, "No existing key")
	}
	greet := MessageData{
		Name:    entity.Name,
		Message: entity.Message,
		Created: entity.Created,
		ID:      entity.Key.ID,
	}
	return c.JSON(http.StatusOK, greet)
}

// e.POST("/api/greetings", addMessage)
func addMessage(c echo.Context) error {
	type postData struct {
		Name    string `json:"author" form:"author" query:"author"`
		Message string `json:"message" form:"message" query:"message"`
	}

	user := new(postData)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}
	entity := ds.Insert(user.Name, user.Message)
	response := MessageData{
		Name:    entity.Name,
		Message: entity.Message,
		Created: entity.Created,
		ID:      entity.Key.ID,
	}
	return c.JSON(http.StatusCreated, response)
}

// e.PUT("/api/greetings/:id", updateMessage)
func updateMessage(c echo.Context) error {
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
	if entity == (ds.MessageEntity{}) {
		return echo.NewHTTPError(http.StatusNotFound, "No existing key")
	}

	entity.Name = data.Name
	entity.Message = data.Message
	entity = ds.Update(entity)

	item := MessageData{
		Name:    data.Name,
		Message: data.Message,
		Created: entity.Created,
		ID:      entity.Key.ID,
	}
	return c.JSON(http.StatusOK, item)
}

// e.DELETE("/api/greetings/:id", deleteMessage)
func deleteMessage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}
	ds.Delete(int64(id))
	return c.String(http.StatusNoContent, "")
}
