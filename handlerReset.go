package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {

	s.db.DeleteDatabase(context.Background())

	fmt.Println("Database records successfully deleted")
	os.Exit(0)
	return nil
}
