package main

import (
	"fmt"
)

func handlerHelp(s *state, cmd command) error {
	maxCommandLength := 9
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("Usage: help (without parameters)")
	}

	fmt.Println("Gator supports the following commands:")
	for cmd, desc := range s.commands.commandDescriptionMap {
		var spaces string
		for i := 0; i < maxCommandLength-len(cmd); i++ {
			spaces = spaces + " "
		}
		fmt.Printf("%s%s: %s\n", cmd, spaces, desc)
	}
	return nil
}
