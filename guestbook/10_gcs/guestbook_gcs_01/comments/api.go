package comments

import (
	"net/http"
	"time"

	"github.com/enakai00/gae_basics_webapp_golang/guestbook/10_gcs/guestbook_gcs_01/ds"
	"github.com/labstack/echo"
)

func Register(e *echo.Echo) {
	e.GET("/api/comments", getComments)
	e.POST("/api/comments", addComment)
}

type CommentData struct {
	Message string    `json:"message"`
	Created time.Time `json:"created"`
	ID      int64     `json:"id"`
}

// e.GET("/api/comments", getComments)
func getComments(c echo.Context) error {
	type queryData struct {
		ParentID int64 `json:"parent_id" form:"parent_id" query:"parent_id"`
	}
	type response struct {
		Comments []CommentData `json:"comments"`
	}

	data := new(queryData)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}
	entities := ds.GetComments(data.ParentID)
	comments := []CommentData{}
	for _, entity := range entities {
		item := CommentData{
			Message: entity.Message,
			Created: entity.Created,
			ID:      entity.Key.ID,
		}
		comments = append(comments, item)
	}
	return c.JSON(http.StatusOK, response{Comments: comments})
}

// e.POST("/api/comments", addComment)
func addComment(c echo.Context) error {
	type postData struct {
		ParentID int64  `json:"parent_id" form:"parent_id" query:"parent_id"`
		Message  string `json:"message" form:"message" query:"message"`
	}

	data := new(postData)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}
	entity := ds.InsertComment(data.ParentID, data.Message)
	response := CommentData{
		Message: entity.Message,
		Created: entity.Created,
		ID:      entity.Key.ID,
	}
	return c.JSON(http.StatusCreated, response)
}
