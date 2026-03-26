package main

import (
	"context"
	"fmt"
	"log"
)

func handleFollowing(state *state, cmd command) error {
	if state.cfg.CurrentUsername == "" {
		return fmt.Errorf("You must be registered. Use command `gator register <name>`\n")
	}

	user, err := state.db.GetUser(context.Background(), state.cfg.CurrentUsername)

	if err != nil {
		log.Fatalf("User not found: %v", err)
	}

	follows, err := state.db.GetFeedFollowsForUser(context.Background(), user.ID)

	fmt.Println("You are following:")
	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}
