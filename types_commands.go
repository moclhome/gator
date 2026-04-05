package main

import (
	"fmt"
)

type command struct {
	name        string
	arguments   []string
	description string
}

type commands struct {
	commandMap            map[string]func(*state, command) error
	commandDescriptionMap map[string]string
}

func (c *commands) run(s *state, cmd command) error {
	function := c.commandMap[cmd.name]
	if function == nil {
		return fmt.Errorf("Unknown command \"%s\".", cmd.name)
	}
	return function(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error, desc string) {
	c.commandMap[name] = f
	c.commandDescriptionMap[name] = desc
}
