package main

import (
	"bootdev/go/gator/internal/database"
	"context"
	"fmt"
	"strconv"
)

func handlerRead(s *state, cmd command, user database.User) error {
	return listPosts(s, cmd, user, true)
}

func listPosts(s *state, cmd command, user database.User, withDescription bool) error {
	if len(cmd.arguments) > 1 {
		return fmt.Errorf("Usage: %s [limit]", cmd.name)
	}
	limit := 2
	var err error
	if len(cmd.arguments) == 1 {
		limit, err = strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return fmt.Errorf("The optional parameter \"limit\" must be an integer")
		}
	}
	bgrd := context.Background()

	posts, err := s.db.GetPostsbyUser(bgrd, user.ID)
	if err != nil {
		return err
	}
	oldFeedName := "empty"
	for i := 0; i < limit; i++ {
		nextPost := posts[i]
		currentFeedName := nextPost.FeedName
		if currentFeedName != oldFeedName {
			fmt.Printf("Feed: %s\n", nextPost.FeedName)
		}
		oldFeedName = currentFeedName
		fmt.Printf("  * Item Title: %v\n", nextPost.Title)
		fmt.Printf("         Link: %v\n", nextPost.Url)
		fmt.Printf("         Published: %v\n", nextPost.PublishedAt)
		if withDescription {
			fmt.Printf("%v\n", nextPost.Description)
		}
	}
	return nil
}
