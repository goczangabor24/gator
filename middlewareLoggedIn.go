package main

import (
	"context"
	"fmt"

	"github.com/goczangabor24/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.configData.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Error: couldn't find current user: %w", err)
		}
		return handler(s, cmd, currentUser)
	}
}
