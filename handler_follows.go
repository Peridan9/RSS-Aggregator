package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/peridan9/RSS-Aggregator/internal/database"
)

// unfollow is a command that unfollows a feed
func handlerUnfollow(s *state, cmd command, user database.User) error {

	// expect exactly one argument
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	// get the url from the command arguments
	url := cmd.Args[0]

	// get the feed by url
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	// delete the feed follow from the database
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't delete follow: %w", err)
	}

	fmt.Println("Feed unfollowed successfully.")
	return nil
}

// handlerFollowsPerUser is a command that lists all the feed follows for a user
func handlerFollowsPerUser(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get follows: %w", err)
	}

	if len(follows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, ff := range follows {
		fmt.Printf("* %s\n", ff.FeedName)
	}

	return nil
}

// handlerFollowing is a command that follows a feed
func handlerFollowing(s *state, cmd command, user database.User) error {

	// expect exactly one argument
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Args)
	}

	// get the url from the command arguments
	url := cmd.Args[0]

	// get the feed by url
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	// create a feed follow in the database
	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create follow: %w", err)
	}

	fmt.Println("Feed follow created:")
	printFeedFollow(follow.UserName, follow.FeedName)
	return nil

}

// printFeedFollow is a function that prints a feed follow
func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
