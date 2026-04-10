package main

import (
	"bootdev/go/gator/internal"
	"bootdev/go/gator/internal/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: agg <Duration>, like agg 1m")
	}
	time_between_reqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Wrong duration format: %v", err)
	}
	fmt.Printf("Collecting feeds every %v...\n", time_between_reqs)
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

	err = s.db.MarkFeedFetched(bgrd, nextFeed.ID)
	if err != nil {
		log.Println("Couldn't mark feed as fetched: ", err)
	}

	scrapeFeed(s.db, nextFeed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	rssFeed, err := internal.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Println("Couldn't collect the feed: ", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		pubDate := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			pubDate = sql.NullTime{Time: t, Valid: true}
		} else if t, err := time.Parse(time.RFC1123, item.PubDate); err == nil {
			pubDate = sql.NullTime{Time: t, Valid: true}
		} else {
			log.Printf("Couldn't parse publication date for feed %s: %v", feed.Name, err)
		}
		postParams := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		}
		_, err = db.CreatePost(context.Background(), postParams)
		if err != nil && !strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint \"posts_url_key\"") {
			// We ignore duplicate errors
			log.Println("Couldn't insert post: ", err)
		}
		if err == nil {
			log.Printf("Found a new post for feed %s!\n", feed.Name)
		}
		/*		fmt.Printf("  * Item Title: %v\n", item.Title)
				fmt.Printf("         Link: %v\n", item.Link)
				fmt.Printf("         Published: %v\n", item.PubDate)*/
	}
}
