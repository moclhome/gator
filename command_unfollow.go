package main

import (
	"bootdev/go/gator/internal/database"
	"context"
	"fmt"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: unfollow <url>")
	}
	url := cmd.arguments[0]
	bgrd := context.Background()

	feed, err := s.db.GetFeed(bgrd, url)
	if err != nil {
		return err
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(bgrd, user.Name)

	isFollowingTheFeed := false
	for _, follow := range feedFollows {
		if follow.FeedName == feed.Name {
			isFollowingTheFeed = true
		}
	}
	if !isFollowingTheFeed {
		return fmt.Errorf("Unfollow not possible because user %s is not following this feed.\n", user.Name)
	}

	unfollowParams := database.DeleteFeedFollowParams{
		FollowingUserID: user.ID,
		FollowedFeedID:  feed.ID,
	}

	err = s.db.DeleteFeedFollow(bgrd, unfollowParams)
	if err != nil {
		return err
	}
	fmt.Printf("%s is no longer following the feed %s.\n", user.Name, feed.Name)
	return nil
}
