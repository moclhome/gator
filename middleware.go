package main

import (
	"bootdev/go/gator/internal/database"
	"context"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	bgrd := context.Background()
	return func(s *state, cmd command) error {
		currentUserName := s.config.Current_user_name
		currentUser, err := s.db.GetUser(bgrd, currentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, currentUser)
	}

}
