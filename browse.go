package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/salehmotiwala/gator/internal/database"
)

func handleBrowse(state *state, cmd command, user database.User) error {
	limit := 2

	if len(cmd.args) > 0 {
		i, err := strconv.Atoi(cmd.args[0])

		if err != nil {
			return fmt.Errorf("Invalid limit. Limit must be a number. %w", err)
		}

		limit = i
	}

	posts, err := state.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		log.Fatalf("Cannot fetch posts: %v", err)
	}

	if len(posts) == 0 {
		fmt.Println("There are no posts!")
		return nil
	}

	for i, post := range posts {
		fmt.Printf("Title: %s\n\n", post.Title)

		fmt.Println(post.Description)

		fmt.Printf("===	END OF POST %d ===\n\n\n", i+1)
	}

	return nil
}
