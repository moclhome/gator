package main

import (
	"bootdev/go/gator/internal/config"
	"fmt"
)

func handlerLogin(s *config.State, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: login <username>")
	}
	if err := s.Config.SetUser(cmd.arguments[0]); err != nil {
		return err
	}
	fmt.Printf("User set to %s\n", cmd.arguments[0])
	return nil
}
