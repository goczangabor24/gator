package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Login handler expects a username")
	}

	if handlerGetUser(s, cmd) != nil {
		return fmt.Errorf("'%v' doesn't exist in the databse. You can't login to an account that doesn't exist", cmd.args[0])
	}

	err := s.configData.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to %v\n", cmd.args[0])
	return nil
}
