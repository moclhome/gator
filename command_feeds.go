package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("Usage: feeds (without parameters)")
	}
	bgrd := context.Background()

	feeds, err := s.db.GetFeeds(bgrd)
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		time := feed.CreatedAt.Format("2006-01-02 15:04")
		fmt.Printf("* %s - %s (inserted by %s at %v)\n", feed.FeedName, feed.Url, feed.UserName, time)
	}
	return nil
}
