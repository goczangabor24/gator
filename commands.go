package main

import (
	"fmt"

	"github.com/goczangabor24/gator/internal/config"
	"github.com/goczangabor24/gator/internal/database"
)

type state struct {
	db         *database.Queries
	configData *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	function, ok := c.registeredCommands[cmd.name]
	if !ok {
		return fmt.Errorf("Command doesn't exist")
	}

	err := function(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
