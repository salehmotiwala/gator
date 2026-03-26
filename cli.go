package main

import (
	"database/sql"
	"fmt"
	"log"
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
	cmds.register("reset", handleReset)
	cmds.register("users", handleUsers)
	cmds.register("agg", handleAgg)
	cmds.register("addfeed", handleAddFeed)
	cmds.register("feeds", handleFeeds)
	cmds.register("follow", handleFollow)
	cmds.register("following", handleFollowing)

	return cmds
}

func parseUserInput() command {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("gator expects a command. No command found")
		os.Exit(1)
	}

	cmd := command{
		name: args[1],
		args: args[2:],
	}

	return cmd
}

func setupDb(cfg *config.Config) (*sql.DB, *database.Queries) {
	db, err := sql.Open("postgres", cfg.DbUrl)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	dbQueries := database.New(db)

	return db, dbQueries
}
