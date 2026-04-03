package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("Usage: reset (without parameters)")
	}
	bgrd := context.Background()

	err := s.db.DeleteAllUsers(bgrd)
	if err != nil {
		return err
	}
	if err := s.config.SetUser(""); err != nil {
		return err
	}
	fmt.Println("All users have been deleted.")
	return nil
}
