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
var	ctx = context.Background()
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

func GetByID(id int64) GuestEntity {
	key := datastore.IDKey("Greeting", id, nil)
	query := datastore.NewQuery("Greeting").Filter("__key__ =", key)
	it := client.Run(ctx, query)

	var entity GuestEntity
	_, err := it.Next(&entity)
	if err != iterator.Done && err != nil {
		log.Fatalf("Error fetching next entity: %v", err)
	}
	return entity
}

func Update(entity GuestEntity) GuestEntity {
	_, err := client.Put(ctx, entity.Key, &entity)
	if err != nil {
		log.Fatalf("Failed to store data: %v", err)
	}
	return entity
}

func Delete(id int64) {
	key := datastore.IDKey("Greeting", id, nil)
	err := client.Delete(ctx, key)
	if err != nil {
		log.Fatalf("Failed to delete data: %v", err)
	}
}

func InsertComment(parentID int64, message string) CommentEntity {
	parentKey := datastore.IDKey("Greeting", parentID, nil)
	key := datastore.IncompleteKey("Comment", parentKey)
	data := CommentEntity{
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

func GetComments(parentID int64) []CommentEntity {
	entities := []CommentEntity{}
	ancestor := datastore.IDKey("Greeting", parentID, nil)
	query := datastore.NewQuery("Comment").Ancestor(ancestor)
	it := client.Run(ctx, query)
	for {
		var entity CommentEntity
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
