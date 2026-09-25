package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Login handler expects a username")
	}

	err := s.configData.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
}
