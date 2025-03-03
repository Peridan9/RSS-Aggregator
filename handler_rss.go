package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/peridan9/RSS-Aggregator/internal/database"
)

// handlerAgg is a command that starts the aggregator
func handlerAgg(s *state, cmd command) error {

	// check if there is a time_between_reqs argument
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	// parse the time between requests
	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't parse time duration: %w", err)
	}

	log.Printf("Starting aggregator with time between requests %v", timeBetweenReqs)

	// Create a ticker that will fire every timeBetweenReqs
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

// scrapeFeeds is a function that scrapes feeds
func scrapeFeeds(s *state) {
	// get the next feed to fetch
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("couldn't get next feed to fetch: %v", err)
		return
	}
	log.Printf("fetching feed %s", feed.Url)
	scrapeFeed(s.db, feed)
}

// scrapeFeed is a function that scrapes a feed
func scrapeFeed(db *database.Queries, feed database.Feed) {
	// mark the feed as fetched
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("couldn't mark feed fetched: %v", err)
		return
	}

	// fetch the feed
	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("couldn't fetch feed: %v", err)
		return
	}

	// create a post for each item in the feed
	for _, item := range feedData.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{Time: t, Valid: true}
		}

		_, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID:    feed.ID,
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
		})

		if err != nil {
			// ignore unique constraint errors
			if strings.Contains(err.Error(), "unique constraint") {
				continue
			}
			log.Printf("couldn't create post: %v", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Url, len(feedData.Channel.Item))
}
