package main

import (
	"bootdev/go/gator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: follow <url>")
	}
	url := cmd.arguments[0]
	bgrd := context.Background()
	currentUserName := s.config.Current_user_name
	currentUser, err := s.db.GetUser(bgrd, currentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeed(bgrd, url)
	if err != nil {
		return err
	}

	followParams := database.CreateFeedFollowParams{
		ID:              uuid.New(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		FollowingUserID: currentUser.ID,
		FollowedFeedID:  feed.ID,
	}
	newFeedFollow, err := s.db.CreateFeedFollow(bgrd, followParams)
	if err != nil {
		return err
	}
	fmt.Printf("%s is now following the feed %s.\n", newFeedFollow.UserName, newFeedFollow.FeedName)
	return nil
}
