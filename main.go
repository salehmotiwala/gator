package main

import (
	"log"

	"github.com/salehmotiwala/gator/internal/config"
)

func main() {
	cmds := mapCommands()
	cfgFile, err := config.Read()

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	state := &state{
		cfg: &cfgFile,
	}

	cmd, err := getUserCommands()

	if err != nil {
		log.Fatalf("%v", err)
	}

	err = cmds.run(state, cmd)

	if err != nil {
		log.Fatalf("%v", err)
	}
}
