package main

import (
	"context"
	"fmt"
	"log"
)

func handleReset(state *state, cmd command) error {
	if err := state.db.DeleteAllUsers(context.Background()); err != nil {
		log.Fatal("Could not delete all users: %v", err)
	}

	fmt.Println("All users deleted.")

	return nil
}
