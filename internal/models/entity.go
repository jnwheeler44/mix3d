package models

import (
	"log"
	"parksport-go/internal/data"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

type Entity struct {
	ID         int64  `json:"id"`
	ParkId     int64  `json:"parkId"`
	ExternalId string `json:"externalId"`
	Name       string `json:"name"`
	EntityType string `json:"entityType"`
	Timezone   string `json:"timezone"`
	UpdatedAt  string `json:"updatedAt"`

	// Potentially nil values
	ParentId  int64   `json:"parentId"`
	Major     bool    `json:"major"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func Entities(c *gin.Context) []Entity {
	var ValidSearchParams = []string{"id", "park_id", "external_id", "name", "entity_type", "major", "lat", "lng", "updated_at", "parent_id"}

	documents := data.IndexQuery("entities", ValidSearchParams, c)
	var entities []Entity

	for {
		doc, err := documents.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v\n", err)
		}

		entity := BuildEntityModel(doc.Data())
		entities = append(entities, entity)
	}

	return entities
}

func EntityById(id string) Entity {
	var entity Entity

	doc := data.ShowQuery("entities", id)
	entity = BuildEntityModel(doc.Data())

	return entity
}

func BuildEntityModel(data map[string]interface{}) Entity {
	entity := &Entity{
		ID:         data["id"].(int64),
		ParkId:     data["park_id"].(int64),
		ExternalId: data["external_id"].(string),
		Name:       data["name"].(string),
		EntityType: data["entity_type"].(string),
		Timezone:   data["timezone"].(string),
		UpdatedAt:  data["updated_at"].(string),
	}

	// Potentially nil values
	entity.ParentId, _ = data["parent_id"].(int64)
	entity.Major, _ = data["major"].(bool)
	entity.Latitude, _ = data["lat"].(float64)
	entity.Longitude, _ = data["lng"].(float64)

	return *entity
}
