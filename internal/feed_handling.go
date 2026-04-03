package internal

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"strconv"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequest("GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	client := http.Client{
		Timeout: 20 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status code while fetchin RSS feed: %s", strconv.Itoa((res.StatusCode)))
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var theFeed RSSFeed
	if err = xml.Unmarshal(data, &theFeed); err != nil {
		return nil, err
	}
	fmt.Printf("Title before: %s\n", theFeed.Channel.Title)
	theFeed.Channel.Title = html.UnescapeString(theFeed.Channel.Title)
	theFeed.Channel.Description = html.UnescapeString(theFeed.Channel.Description)
	for i, item := range theFeed.Channel.Item {
		theFeed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		theFeed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	return &theFeed, nil
}
