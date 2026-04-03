package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("Usage: following (without parameters)")
	}
	bgrd := context.Background()
	user, err := s.db.GetUser(bgrd, s.config.Current_user_name)
	if err != nil {
		return err
	}

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
