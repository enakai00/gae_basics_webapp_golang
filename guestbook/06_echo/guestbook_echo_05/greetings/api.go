package greetings

import (
	"net/http"

	"github.com/labstack/echo"
)

func Register(e *echo.Echo) {
	e.GET("/api/greetings", greetings)
	e.POST("/api/greetings", addUser)
}

type userData struct {
	ID      int    `json:"ID"`
	Name    string `json:"author"`
	Message string `json:"Message"`
}

type postData struct {
	Name    string `json:"author" form:"author" query:"author"`
	Message string `json:"Message" form:"Message" query:"Message"`
}

// e.GET("/api/greetings", home)
func greetings(c echo.Context) error {
	type response struct {
		Greetings []userData `json:"greetings"`
	}

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

// e.POST("/api/greetings", addUser)
func addUser(c echo.Context) (err error) {
	user := new(postData)
	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	data := userData{
		ID:      999,
		Name:    user.Name,
		Message: user.Message,
	}
	return c.JSON(http.StatusOK, data)
}
