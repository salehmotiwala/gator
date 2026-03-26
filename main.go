package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/salehmotiwala/gator/internal/config"
)

func main() {
	cmds := mapCommands()

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	_, dbQueries := setupDb(&cfg)

	state := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmd := parseUserInput()

	err = cmds.run(state, cmd)

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
