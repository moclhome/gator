package main

import (
	"bootdev/go/gator/internal/config"
	"fmt"
)

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandMap map[string]func(*config.State, command) error
}

func (c *commands) run(s *config.State, cmd command) error {
	function := c.commandMap[cmd.name]
	if function == nil {
		return fmt.Errorf("Unknown command \"%s\".", cmd.name)
	}
	return function(s, cmd)
}

func (c *commands) register(name string, f func(*config.State, command) error) {
	c.commandMap[name] = f
}
