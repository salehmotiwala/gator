package main

import (
	"context"
	"fmt"
	"log"
)

func handleLogin(state *state, cmd command) error {
	if len(cmd.args) == 0 || len(cmd.args) > 1 {
		return fmt.Errorf("The login command must have a single argument, i.e., username.")
	}

	username := cmd.args[0]

	user, err := state.db.GetUser(context.Background(), username)

	if err != nil {
		log.Fatal("User does not exist.")
	}

	err = state.cfg.SetUser(user.Name)

	if err != nil {
		return err
	}

	fmt.Printf("User `%s` has been set.", username)

	return nil
}
