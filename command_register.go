package main

import (
	"bootdev/go/gator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: register <username>")
	}
	userName := cmd.arguments[0]
	bgrd := context.Background()

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName}
	newUser, err := s.db.CreateUser(bgrd, userParams)
	if err != nil {
		return err
	}
	if err := s.config.SetUser(userName); err != nil {
		return err
	}
	fmt.Printf("User %s has been created.\nValues: %v", userName, newUser)
	return nil
}
