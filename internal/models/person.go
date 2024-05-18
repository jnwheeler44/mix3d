package models

import (
	"log"
	"mix3d/internal/data"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

type Person struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func People(c *gin.Context) []Person {
	var ValidSearchParams = []string{"email", "first_name", "last_name", "name", "id"}

	documents := data.IndexQuery("people", ValidSearchParams, c)
	var people []Person

	for {
		doc, err := documents.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v\n", err)
		}

		data := doc.Data()
		person := BuildPersonModel(data)
		people = append(people, person)
	}

	return people
}

func PersonById(id string) Person {
	var Person Person

	data := data.ShowQuery("People", id).Data()
	if data == nil {
		return Person
	}
	Person = BuildPersonModel(data)

	return Person
}

func CreatePerson(Person Person) map[string]interface{} {
	PersonData := map[string]interface{}{
		"email":      Person.Email,
		"name":       Person.Name,
		"first_name": Person.FirstName,
		"last_name":  Person.LastName,
	}
	result := data.CreateQuery("People", PersonData)

	return result
}

func BuildPersonModel(data map[string]interface{}) Person {
	Person := Person{
		ID:        data["id"].(int64),
		Email:     data["email"].(string),
		Name:      data["name"].(string),
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
	}

	return Person
}
