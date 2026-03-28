package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/salehmotiwala/gator/internal/database"
)

func handleFollow(state *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Command `follow` must have only one argument <url>")
	}

	feedUrl := cmd.args[0]

	feed, err := state.db.GetFeed(context.Background(), feedUrl)

	if err != nil {
		log.Fatalf("User not found: %v", err)
	}

	params := database.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	row, err := state.db.CreateFeedFollow(context.Background(), params)

	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println("Feed followed!")
	fmt.Printf("Name: %s	|	Username: %s", row.FeedName, user.Name)

	return nil
}
