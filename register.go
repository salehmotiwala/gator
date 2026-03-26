package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/salehmotiwala/gator/internal/database"
)

func handleRegister(state *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Command `register` must have only 1 argument.")
	}

	name := cmd.args[0]

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	user, err := state.db.CreateUser(context.Background(), params)

	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}

	if err := state.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("Could not update user: %w", err)
	}

	fmt.Println("User was created!")
	fmt.Println(user)

	return nil
}
