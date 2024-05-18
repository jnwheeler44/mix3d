package models

import (
	"log"
	"parksport-go/internal/data"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

type Visit struct {
	ID            int64  `json:"id"`
	UserId        int64  `json:"userId"`
	DestinationId int64  `json:"destinationId"`
	TripId        int64  `json:"tripId"`
	ParkId        int64  `json:"parkId"`
	VisitedAt     string `json:"visitedAt"`
	UpdatedAt     string `json:"updatedAt"`

	// Potentially nil values
	EntityId         int64 `json:"entityId"`
	LocationVerified bool  `json:"locationVerified"`
	ParentId         int64 `json:"parentId"`
}

func Visits(c *gin.Context) []Visit {
	var ValidSearchParams = []string{"id", "destination_id", "entity_id", "trip_id", "user_id", "park_id", "parent_id", "visited_at", "location_verified"}

	documents := data.IndexQuery("visits", ValidSearchParams, c)
	var visits []Visit

	for {
		doc, err := documents.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v\n", err)
		}

		visit := BuildVisitModel(doc.Data())
		visits = append(visits, visit)
	}

	return visits
}

func VisitById(id string) Visit {
	var visit Visit

	doc := data.ShowQuery("visits", id)
	visit = BuildVisitModel(doc.Data())

	return visit
}

func BuildVisitModel(data map[string]interface{}) Visit {
	visit := &Visit{
		ID:            data["id"].(int64),
		DestinationId: data["destination_id"].(int64),
		TripId:        data["trip_id"].(int64),
		UserId:        data["user_id"].(int64),
		ParkId:        data["park_id"].(int64),
		VisitedAt:     data["visited_at"].(string),
		UpdatedAt:     data["updated_at"].(string),
	}

	// Potentially nil values
	visit.EntityId, _ = data["entity_id"].(int64)
	visit.LocationVerified, _ = data["location_verified"].(bool)
	visit.ParentId, _ = data["parent_id"].(int64)

	return *visit
}
