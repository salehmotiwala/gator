package main

import (
	"context"
	"fmt"
	"log"

	"github.com/salehmotiwala/gator/internal/database"
)

func handleUnfollow(state *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Command `unfollow` must have only one arguments: <feed_url>.")
	}

	feed_url := cmd.args[0]

	err := state.db.DeleteFeedFollowsForUser(context.Background(), database.DeleteFeedFollowsForUserParams{
		UserID: user.ID,
		Url:    feed_url,
	})

	if err != nil {
		log.Fatalf("Cannot unfollow feed: %v", err)
	}

	fmt.Println("Feed unfollowed!")

	return nil
}
