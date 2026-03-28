package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/salehmotiwala/gator/internal/database"
)

func handleAddFeed(state *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("Command `addfeed` must have only two arguments: <name> <url>\n")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := state.db.CreateFeed(context.Background(), params)

	if err != nil {
		log.Fatalf("Could not create feed: %v", err)
	}

	_, err = state.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		fmt.Println(fmt.Printf("Could not follow feed: %w", err))
	}

	fmt.Println("Feed created!")
	fmt.Println(feed)

	return nil
}
