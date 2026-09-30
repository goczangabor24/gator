package main

import (
	"context"
	"fmt"
)

func handlerDelete(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("name to be deleted required")
	}

	nameToBeDeleted := cmd.args[0]

	if handlerGetUser(s, cmd) != nil {
		return fmt.Errorf("'%v' doesn't exist in the databse. You can't delete an account that doesn't exist", cmd.args[0])
	}

	// if s.configData.CurrentUserName == nameToBeDeleted {
	// 	s.configData.SetUser("No user is logged in")
	// 	return nil
	// }

	user, err := s.db.DeleteUser(context.Background(), nameToBeDeleted)
	if err != nil {
		return err
	}

	fmt.Printf("User '%s' was deleted!\n", user.Name)

	return nil
}
