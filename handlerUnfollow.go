package main

import (
	"context"
	"fmt"

	"github.com/goczangabor24/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: Unfollow takes a single 'url' argument")
	}

	params := database.UnfollowParams{
		Url: cmd.args[0],
		ID:  user.ID,
	}

	s.db.Unfollow(context.Background(), params)

	return nil
}
