package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/salehmotiwala/gator/internal/database"
)

func handleFollow(state *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Command `follow` must have only one argument <url>")
	}

	if state.cfg.CurrentUsername == "" {
		return fmt.Errorf("You must be registered. Use command `register <name>`.")
	}

	feedUrl := cmd.args[0]

	user, err := state.db.GetUser(context.Background(), state.cfg.CurrentUsername)

	if err != nil {
		log.Fatalf("User not found: %v", err)
	}

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
