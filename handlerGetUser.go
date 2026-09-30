package main

import (
	"context"
	"fmt"
)

func handlerGetUser(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("name parameter required")
	}

	nameToGet := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), nameToGet)
	if err != nil {
		return fmt.Errorf("no record of user: '%v' in the database", user)
	}

	return nil
}
