package main

import (
	"context"
	"log"
	"testing"

	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
	_ "github.com/lib/pq"
)

func TestUserCreate(t *testing.T) {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=milky dbname=blog password=1234 sslmode=disable")

	if err != nil {
		log.Fatalf("Failed to connect to the postgresql : %v", err)
	}

	defer client.Close()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	searchUser, err := client.User.
		Query().
		Where(user.ID(1)).
		Only(ctx)

	if err != nil {
		log.Fatalf("can't find user : %v", err)
	}

	if searchUser.Username != "Mil" {
		t.Fatalf("This User's username is %v, not Milky.", searchUser.Username)
	}
}
