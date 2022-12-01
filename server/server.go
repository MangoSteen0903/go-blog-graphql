package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/ent/migrate"
	"github.com/MangoSteen0903/go-blog-graphql/resolvers"
	"github.com/MangoSteen0903/go-blog-graphql/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const PORT = "8000"

func main() {
	err := godotenv.Load(".env")
	ctx := context.Background()
	utils.HandleErr(err, "Can't Load .env File : ")

	portUrl := fmt.Sprintf(":%v", PORT)
	config := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USERNAME"),
		os.Getenv("DBNAME"),
		os.Getenv("PASSWORD"),
		"disable",
	)
	client, err := ent.Open("postgres", config)
	utils.HandleErr(err, "Can't Open ent Client : ")

	if err := client.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("Can't Open ent Client")
	}

	srv := handler.NewDefaultServer(resolvers.NewSchema(client))
	http.Handle("/", playground.Handler("User", "/query"))

	http.Handle("/query", srv)
	log.Printf("Listening on http://localhost%v", portUrl)

	if err := http.ListenAndServe(portUrl, nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
