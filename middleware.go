package main

import (
	"context"
	"fmt"
	"log"

	"github.com/salehmotiwala/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		if s.cfg.CurrentUsername == "" {
			return fmt.Errorf("You are not logged in. Please use command gator register <name>.")
		}

		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)

		if err != nil {
			log.Fatalf("Could not get user. %v", err)
		}

		return handler(s, c, user)
	}
}
