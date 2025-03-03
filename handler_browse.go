package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/peridan9/RSS-Aggregator/internal/database"
)

// handlerBrowse is a command that lists all the posts for a user by limit
func handlerBrowse(s *state, cmd command, user database.User) error {
	// check if there is a limit argument or set defult limit to 2
	limit := 2
	if len(cmd.Args) == 1 {
		// if limit is specified, parse it
		if specifiedLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	// get the posts for the user
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	// print the posts
	fmt.Printf("Found %d posts: for user %s\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
