package ds

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
var ctx = context.Background()
var client, _ = datastore.NewClient(ctx, projectID)

type GuestEntity struct {
	Name    string         `datastore:"author"`
	Message string         `datastore:"message"`
	Created time.Time      `datastore:"created"`
	Key     *datastore.Key `datastore:"__key__"`
}

type CommentEntity struct {
	Message string         `datastore:"message"`
	Created time.Time      `datastore:"created"`
	Key     *datastore.Key `datastore:"__key__"`
}

func Insert(author, message string) GuestEntity {
	key := datastore.IncompleteKey("Greeting", nil)
	data := GuestEntity{
		Name:    author,
		Message: message,
		Created: time.Now(),
	}
	key, err := client.Put(ctx, key, &data)
	if err != nil {
		log.Fatalf("Failed to store data: %v", err)
	}
	data.Key = key

	return data
}

func GetAll() []GuestEntity {
	entities := []GuestEntity{}
	query := datastore.NewQuery("Greeting").Order("-created")
	it := client.Run(ctx, query)
	for {
		var entity GuestEntity
		_, err := it.Next(&entity)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next entity: %v", err)
		}
		entities = append(entities, entity)
	}
	return entities
}
