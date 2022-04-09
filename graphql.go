package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/graphql-go/graphql"
)

type Tutorial struct {
	ID int
	Title string
	Author Author
	Comments []Comment
}

type Author struct {
	Name string
	Tutorials []Tutorial
}

type Comment struct {
	Body string
}

func populate() []Tutorial {
	author := &Author{Name: "Haruki Murakami", Tutorials: []int{1}}
	tutorial := Tutorial{
		ID: 1,
		Title: "Golang GraphQL TUTORIAL",
		Author: author,
		Comments: []Comment{
			Comment{Body: "such an awesome course!!"},
		},
	}
	var tutorials := []Tutorial
	tutorials = append(tutorials, tutorial)

	return tutorials
}

func main() {
	fmt.Println("GraphQL Tutorial")

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "World", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create new GraphQL schema, err: %v", err)
	}

	query := `
		{
			hello
		}
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("Failed to execute graphql operation, err: %v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}