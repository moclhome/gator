package main

import (
	"bootdev/go/gator/internal"
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("Usage: agg (without parameters)")
	}
	rssFeed, err := internal.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println("Fetched the following feed:")
	fmt.Printf("Title: %v\n", rssFeed.Channel.Title)
	fmt.Printf("Description: %v\n", rssFeed.Channel.Description)
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("* Item Title: %v\n", item.Title)
		fmt.Printf("* Item Description: %v\n", item.Description)
	}
	return nil
}
