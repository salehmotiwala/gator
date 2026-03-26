package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	list map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	command, cmdExists := c.list[cmd.name]

	if !cmdExists {
		return fmt.Errorf("Command %s does not exist.", cmd.name)
	}

	err := command(s, cmd)

	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	if c.list == nil {
		c.list = make(map[string]func(*state, command) error)
	}

	c.list[name] = f
}
