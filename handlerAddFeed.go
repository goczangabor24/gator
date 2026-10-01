package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/goczangabor24/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {

	if len(cmd.args) != 2 {
		return fmt.Errorf("Error: The addFeed function needs at least two parameters: name and url")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	currentUser, err := s.db.GetUser(context.Background(), s.configData.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error: couldn't find current user: %w", err)
	}

	params := database.CreatedFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	}

	feed, err := s.db.CreatedFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	fmt.Printf("%+v\n", feed)

	//ADDING FEED TO FOLLOWED FEEDS

	feedID, err := s.db.FeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	params2 := database.FeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feedID,
	}

	follow, err := s.db.FeedFollows(context.Background(), params2)
	if err != nil {
		return err
	}

	fmt.Println(follow.FeedName)
	fmt.Println(currentUser.Name)

	return nil
}
