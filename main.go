package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/salehmotiwala/gator/internal/config"
	"github.com/salehmotiwala/gator/internal/database"
)

func main() {
	cmds := mapCommands()
	cfgFile, err := config.Read()

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	db, err := sql.Open("postgres", cfgFile.DbUrl)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	dbQueries := database.New(db)

	state := &state{
		cfg: &cfgFile,
		db:  dbQueries,
	}

	cmd, err := getUserCommands()

	if err != nil {
		log.Fatalf("%v", err)
	}

	err = cmds.run(state, cmd)

	if err != nil {
		log.Fatalf("%v", err)
	}

	os.Exit(0)
}
