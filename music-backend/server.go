package main

import (
	"fmt"
	"log"
	"my_gql_server/database"
	"my_gql_server/graph"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)

func init() {
	db := database.Connect()
	defer db.Close()

	err := db.Ping()

	if err != nil {
		fmt.Println("データベース接続失敗")
		return
	} else {
		fmt.Println("データベース接続成功")
	}

	err = database.CreateTableIfNotExists(db)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

const defaultPort = "8083"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001"},
		AllowCredentials: true,
		Debug:            true,
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
