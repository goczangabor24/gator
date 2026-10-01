package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {

	feeds, err := s.db.FeedsDisplay(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) < 1 {
		return fmt.Errorf("Error: No feeds in the database")
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(feed.Name_2)
		fmt.Println()
	}

	return nil
}
