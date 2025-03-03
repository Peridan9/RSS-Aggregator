package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

// RSSFeed is a struct that represents an RSS feed
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

// RSSItem is a struct that represents an RSS item
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// fetchFeed fetches an RSS feed from a given URL
func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	// Create an HTTP client with a timeout
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	// Create a new request with the given URL
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	// Set the User-Agent header to identify the client
	req.Header.Set("User-Agent", "gator")
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Defer closing the response body
	defer res.Body.Close()

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the XML data into an RSSFeed struct
	var rssfeed RSSFeed
	err = xml.Unmarshal(data, &rssfeed)
	if err != nil {
		return nil, err
	}

	// Unescape HTML entities in the feed
	rssfeed.Channel.Title = html.UnescapeString(rssfeed.Channel.Title)
	rssfeed.Channel.Description = html.UnescapeString(rssfeed.Channel.Description)
	for i, item := range rssfeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssfeed.Channel.Item[i] = item
	}

	// Return the RSS feed
	return &rssfeed, nil

}
