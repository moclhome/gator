package main

import (
	"bootdev/go/gator/internal/database"
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("Usage: following (without parameters)")
	}
	bgrd := context.Background()

	feedFollows, err := s.db.GetFeedFollowsForUser(bgrd, user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("User %s is following these feeds:\n", user.Name)
	for _, feedFollow := range feedFollows {
		fmt.Printf("* %s\n", feedFollow.FeedName)
	}
	return nil
}
