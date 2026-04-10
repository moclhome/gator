package main

import (
	"bootdev/go/gator/internal"
	"bootdev/go/gator/internal/database"
	"context"
	"fmt"
	"log"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: agg <Duration>, like agg 1m")
	}
	time_between_reqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Wrong duration format: %v", err)
	}
	fmt.Printf("Collecting feeds every %v...", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state) {
	bgrd := context.Background()

	nextFeed, err := s.db.GetNextFeedToFetch(bgrd)
	if err != nil {
		log.Println("Couldn't get next feed to fetch: ", err)
	}
	log.Printf("Found a new feed: %s\n", nextFeed.Name)

	err = s.db.MarkFeedFetched(bgrd, nextFeed.ID)
	if err != nil {
		log.Println("Couldn't mark feed as fetched: ", err)
	}

	fetchFeed(nextFeed)
}

func fetchFeed(feed database.Feed) {
	rssFeed, err := internal.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Println("Couldn't collect the feed: ", err)
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("  * Item Title: %v\n", item.Title)
		fmt.Printf("         Link: %v\n", item.Link)
	}
}
