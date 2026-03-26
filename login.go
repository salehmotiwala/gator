package main

import "fmt"

func handleLogin(state *state, cmd command) error {
	if len(cmd.args) == 0 || len(cmd.args) > 1 {
		return fmt.Errorf("The login command must have a single argument, i.e., username.")
	}

	username := cmd.args[0]
	err := state.cfg.SetUser(username)

	if err != nil {
		return err
	}

	fmt.Printf("User `%s` has been set.", username)

	return nil
}
