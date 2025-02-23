package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/peridan9/RSS-Aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

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
	printFeed(feed)
	fmt.Println()
	fmt.Println("====================================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("ID:         %s\n", feed.ID)
	fmt.Printf("Created At: %s\n", feed.CreatedAt)
	fmt.Printf("Updated At: %s\n", feed.UpdatedAt)
	fmt.Printf("Name:       %s\n", feed.Name)
	fmt.Printf("URL:        %s\n", feed.Url)
	fmt.Printf("User ID:    %s\n", feed.UserID)
}
