package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("Usage: users (without parameters)")
	}
	bgrd := context.Background()

	users, err := s.db.GetUsers(bgrd)
	if err != nil {
		return err
	}
	for _, user := range users {
		fmt.Printf("* %s", user.Name)
		if user.Name == s.config.Current_user_name {
			fmt.Println(" (current)")
		} else {
			fmt.Println("")
		}
	}
	return nil
}
