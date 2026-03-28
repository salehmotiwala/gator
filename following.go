package main

import (
	"context"
	"fmt"
	"log"

	"github.com/salehmotiwala/gator/internal/database"
)

func handleFollowing(state *state, cmd command, user database.User) error {

	follows, err := state.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		log.Fatalf("Cannot get following users. %v", err)
	}

	fmt.Println("You are following:")
	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}
