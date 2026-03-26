package main

import (
	"context"
	"fmt"
	"log"
)

func handleFeeds(state *state, cmd command) error {
	feeds, err := state.db.GetFeeds(context.Background())

	if err != nil {
		log.Fatalf("Cannot fetch feeds: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %s	|	Link: %s	|	Username: %s\n", feed.Name, feed.Url, feed.Username)
	}

	return nil
}
