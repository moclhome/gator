package main

import (
	"fmt"
)

func handlerHelp(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("Usage: help (without parameters)")
	}

	fmt.Printf("Help")
	return nil
}
