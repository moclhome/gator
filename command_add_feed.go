package main

import (
	"bootdev/go/gator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Usage: addFeed <feedName> <url>")
	}
	feedName := cmd.arguments[0]
	url := cmd.arguments[1]
	bgrd := context.Background()
	currentUserName := s.config.Current_user_name
	currentUser, err := s.db.GetUser(bgrd, currentUserName)
	if err != nil {
		return err
	}

	feedParams := database.CreateFeedParams{
		ID:              uuid.New(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Name:            feedName,
		CreatedByUserID: currentUser.ID,
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
		FollowingUserID: currentUser.ID,
		FollowedFeedID:  newFeed.ID,
	}
	_, err = s.db.CreateFeedFollow(bgrd, followParams)
	if err != nil {
		return err
	}

	fmt.Printf("Feed %s has been created by user %s.\nValues: %v", feedName, currentUserName, newFeed)
	return nil
}
