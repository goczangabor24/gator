package main

import (
	"context"
	"fmt"

	"github.com/goczangabor24/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	if len(cmd.args) != 0 {
		return fmt.Errorf("Error: Following command doesn't take arguments")
	}

	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range followedFeeds {
		fmt.Println(feed.Feedname)
	}

	return nil
}
