package data

import (
	"context"
	"log"
	"slices"
	"strconv"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// func Where()

func DB() *firestore.Client {
	ctx := context.Background()
	return createClient(ctx)
}

func createClient(ctx context.Context) *firestore.Client {
	sa := option.WithCredentialsFile("C:\\Users\\jnwhe\\dev\\google.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

func IndexQuery(collectionName string, validParams []string, c *gin.Context) *firestore.DocumentIterator {
	query := DB().Collection(collectionName).Query

	for k, v := range c.Request.URL.Query() {
		if slices.Contains(validParams, k) {
			idInt, err := strconv.Atoi(v[0])
			if err != nil {
				query = query.Where(k, ">=", v[0]).Where(k, "<=", v[0]+"~")
			} else {
				query = query.Where(k, "==", idInt)
			}
		} else if k == "sort" {
			order := strings.Split(v[0], ":")
			col := order[0]
			dir := 1
			if order[1] == "desc" {
				dir = 2
			}
			query = query.OrderBy(col, firestore.Direction(dir))
		} else if k == "limit" {
			limit, _ := strconv.Atoi(v[0])
			query = query.Limit(limit)
		}
	}

	return query.Documents(c.Request.Context())
}

func ShowQuery(collectionName string, id string) *firestore.DocumentSnapshot {
	query, _ := DB().Collection(collectionName).Doc(id).Get(context.Background())

	return query
}

func UpsertQuery(collectionName string, data map[string]interface{}) map[string]interface{} {
	id, _ := data["id"].(string)

	_, err := DB().Collection(collectionName).Doc(id).Set(context.Background(), data)

	if err != nil {
		data = map[string]interface{}{"status": "error", "message": err}
	}

	return data
}

func CreateQuery(collectionName string, data map[string]interface{}) map[string]interface{} {
	var id int64 = 0
	var textId string
	docs := DB().Collection(collectionName).Query.OrderBy("id", firestore.Desc).Limit(1).Documents(context.Background())
	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v\n", err)
		}
		d := doc.Data()
		id = d["id"].(int64)
	}

	newId := id + 1
	data["id"] = newId
	textId = strconv.Itoa(int(newId))

	_, err := DB().Collection(collectionName).Doc(textId).Set(context.Background(), data)

	if err != nil {
		data = map[string]interface{}{"status": "error", "message": err}
	}

	return data
}
