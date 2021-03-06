package photos

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo"
	"google.golang.org/api/iterator"
)

var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
var ctx = context.Background()
var client, _ = storage.NewClient(ctx)
var bucketName = projectID

func Register(e *echo.Echo) {
	e.GET("/photos", getPhotos)
	e.POST("/photos", addPhoto)
}

// e.GET("/photos", getPhotos)
func getPhotos(c echo.Context) error {
	type PhotoData struct {
		PublicURL string
		Name      string
	}

	it := client.Bucket(bucketName).Objects(c.Request().Context(), nil)
	data := []PhotoData{}
	baseURL := "https://storage.cloud.google.com/" + bucketName + "/"
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read bucket: %v", err)
		}

		item := PhotoData{
			PublicURL: baseURL + attrs.Name,
			Name:      attrs.Name,
		}
		data = append(data, item)
	}
	return c.Render(http.StatusOK, "photos", data)
}

// e.POST("/photos", addPhoto)
func addPhoto(c echo.Context) error {
	file, _ := c.FormFile("file")
	src, err := file.Open()
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer src.Close()

	bucket := client.Bucket(bucketName)
	dst := bucket.Object(file.Filename).NewWriter(c.Request().Context())
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}
	return c.Render(http.StatusOK, "complete", nil)
}
