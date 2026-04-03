package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: login <username>")
	}
	userName := cmd.arguments[0]

	bgrd := context.Background()

	_, err := s.db.GetUser(bgrd, userName)
	if err != nil {
		return err
	}
	if err := s.config.SetUser(userName); err != nil {
		return err
	}
	fmt.Printf("User set to %s\n", userName)
	return nil
}
