package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {

	users, err := s.db.Users(context.Background())
	if err != nil {
		return err
	}

	if len(users) < 1 {
		return fmt.Errorf("Error: No users in the database")
	}

	for _, user := range users {
		if user == s.configData.CurrentUserName {
			fmt.Printf("%v (current)\n", user)
		} else {
			fmt.Println(user)
		}
	}

	return nil
}
