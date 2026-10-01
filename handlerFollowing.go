package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {

	if len(cmd.args) != 0 {
		return fmt.Errorf("Error: Following command doesn't take arguments")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.configData.CurrentUserName)
	if err != nil {
		return err
	}

	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return err
	}

	for _, feed := range followedFeeds {
		fmt.Println(feed.Feedname)
	}

	return nil
}
