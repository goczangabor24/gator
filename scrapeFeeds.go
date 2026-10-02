package main

import (
	"context"
	"fmt"
	"time"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("%s\n\n", feed.Name)

	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range fetchedFeed.Channel.Item {
		fmt.Printf("%v\n\n", item.Title)
		time.Sleep(200 * time.Millisecond)
	}

	return nil
}
