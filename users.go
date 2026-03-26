package main

import (
	"context"
	"fmt"
	"log"
)

func handleUsers(state *state, cmd command) error {
	users, err := state.db.GetUsers(context.Background())

	if err != nil {
		log.Fatalf("Could not fetch users: %v", err)
	}

	loggedUser := state.cfg.CurrentUsername

	for _, user := range users {
		current := ""

		if user.Name == loggedUser {
			current = " (current)"
		}

		fmt.Printf("* %s%s\n", user.Name, current)
	}

	return nil
}
