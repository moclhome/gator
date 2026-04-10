package main

import (
	"bootdev/go/gator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Usage: addFeed <feedName> <url>")
	}
	feedName := cmd.arguments[0]
	url := cmd.arguments[1]
	bgrd := context.Background()
	currentUserName := user.Name

	feedParams := database.CreateFeedParams{
		ID:              uuid.New(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Name:            feedName,
		CreatedByUserID: user.ID,
		Url:             url,
	}
	newFeed, err := s.db.CreateFeed(bgrd, feedParams)
	if err != nil {
		return err
	}

	followParams := database.CreateFeedFollowParams{
		ID:              uuid.New(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		FollowingUserID: user.ID,
		FollowedFeedID:  newFeed.ID,
	}
	_, err = s.db.CreateFeedFollow(bgrd, followParams)
	if err != nil {
		return err
	}

	fmt.Printf("Feed %s has been created by user %s.\nValues: %v\n", feedName, currentUserName, newFeed)
	return nil
}
