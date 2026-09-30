package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/goczangabor24/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username required")
	}

	if handlerGetUser(s, cmd) == nil {
		return fmt.Errorf("'%v' is already registered", cmd.args[0])
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	err = s.configData.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User %s was created!\n", user.Name)
	fmt.Printf("%+v\n", user)

	return nil
}
