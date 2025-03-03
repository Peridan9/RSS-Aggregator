package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/peridan9/RSS-Aggregator/internal/database"
)

// handlerGetFeeds is a command that gets all the feeds
func handlerGetFeeds(s *state, cmd command) error {

	// get all the feeds
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}
	fmt.Println("Feeds:")

	// print each feed
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}
		// print the feed
		printFeed(feed, user)
		fmt.Println()
	}
	fmt.Println("====================================")
	return nil
}

// handlerAddFeed is a command that adds a feed
func handlerAddFeed(s *state, cmd command, user database.User) error {

	// expect exactly two arguments
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	// get the name and url from the command arguments
	name := cmd.Args[0]
	url := cmd.Args[1]

	// create the feed in the database
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return fmt.Errorf("couldn't create a feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed, user)
	fmt.Println()
	fmt.Println("====================================")

	// follow the feed for the current user
	err = handlerFollowing(s, command{Name: "follow", Args: []string{url}}, user)
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	return nil
}

// printFeed is a function that prints a feed
func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("ID:         %s\n", feed.ID)
	fmt.Printf("Created At: %s\n", feed.CreatedAt)
	fmt.Printf("Updated At: %s\n", feed.UpdatedAt)
	fmt.Printf("Name:       %s\n", feed.Name)
	fmt.Printf("URL:        %s\n", feed.Url)
	fmt.Printf("User ID:    %s\n", feed.UserID)
	fmt.Printf("User Name:  %s\n", user.Name)
	fmt.Printf("LastFetchedAt: %v\n", feed.LastFetchedAt)
}
