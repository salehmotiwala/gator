package main

import (
	"context"
	"fmt"
	"log"

	"github.com/salehmotiwala/gator/internal/feeds"
)

func handleAgg(state *state, cmd command) error {
	feed, err := feeds.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		log.Fatalf("Failed to fetch feed: %v", err)
	}

	fmt.Println(feed)

	return nil
}
