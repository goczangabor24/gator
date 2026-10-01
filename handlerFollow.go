package main

import (
	"context"
	"fmt"
	"time"

	"github.com/goczangabor24/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: Follow command only takes a 'url' argument")
	}

	url := cmd.args[0]

	feedID, err := s.db.FeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	currentUser, err := s.db.GetUser(context.Background(), s.configData.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error: couldn't find current user: %w", err)
	}

	params := database.FeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feedID,
	}

	follow, err := s.db.FeedFollows(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println(follow.FeedName)
	fmt.Println(currentUser.Name)

	return nil
}
