package main

import (
	"fmt"
	"os"

	"github.com/salehmotiwala/gator/internal/config"
	"github.com/salehmotiwala/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func mapCommands() commands {
	cmds := commands{}

	cmds.register("login", handleLogin)
	cmds.register("register", handleRegister)

	return cmds
}

func getUserCommands() (command, error) {
	args := os.Args

	if len(args) < 2 {
		return command{}, fmt.Errorf("gator expects a command. No command found")
	}

	cmd := command{
		name: args[1],
		args: args[2:],
	}

	return cmd, nil
}
