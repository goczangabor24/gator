package main

import (
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(s *state, cmd command) error {
	var limit int32
	limit = 2

	if len(cmd.args) > 1 {
		return fmt.Errorf("Browse takes only an optional 'limit' argument for the number of posts to be shown")
	} else if len(cmd.args) == 1 {
		parsedNumber, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return err
		}

		limit = int32(parsedNumber)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), limit)
	if err != nil {
		return err
	}

	postNumber := 1

	for _, post := range posts {
		fmt.Printf("%d. Title: %v\n\n", postNumber, post.Title)
		fmt.Printf("Body: %v\n\n", post.Description)
		postNumber++
	}
	return nil
}
