package main

import (
	"fmt"
	"os"

	"github.com/salehmotiwala/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func mapCommands() commands {
	cmds := commands{}

	cmds.register("login", handleLogin)

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
