package ds

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/datastore"
)

var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

type Entity struct {
	Name    string    `datastore:"author"`
	Message string    `datastore:"message"`
	Created time.Time `datastore:"created"`
	ID      int64     `datastore:"__key__"`
}

func Insert(author, message string) Entity {
	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, projectID)

	data := Entity{
		Name:    author,
		Message: message,
		Created: time.Now(),
	}

	kind := "Greeting"
	key := datastore.IncompleteKey(kind, nil)

	keyWithID, err := client.Put(ctx, key, &data)
	if err != nil {
		log.Fatalf("Failed to store data: %v", err)
	}
	data.ID = keyWithID.ID

	return data
}
