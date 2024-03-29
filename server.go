package main // import "github.com/shobhitsharma/go-gql-starter"

import (
	"context"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/shobhitsharma/go-gql-starter/db"
	"github.com/shobhitsharma/go-gql-starter/handler"
	"github.com/shobhitsharma/go-gql-starter/resolvers"
	"github.com/shobhitsharma/go-gql-starter/schema"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	context.Background()

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(schema.GetSchema(), &resolvers.Resolvers{DB: db}, opts...)

	mux := http.NewServeMux()
	mux.Handle("/", handler.GraphiQL{})
	mux.Handle("/query", handler.Authenticate(&handler.GraphQL{Schema: schema}))

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Listening to... port 8080")
	if err = s.ListenAndServe(); err != nil {
		panic(err)
	}
}
