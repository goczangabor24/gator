package main

import (
	"context"
	"fmt"
	"time"

	"github.com/goczangabor24/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: Follow command only takes a 'url' argument")
	}

	url := cmd.args[0]

	feedID, err := s.db.FeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	params := database.FeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedID,
	}

	follow, err := s.db.FeedFollows(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println(follow.FeedName)
	fmt.Println(user.Name)

	return nil
}
