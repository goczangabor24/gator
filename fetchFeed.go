package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	var fetchedFeed RSSFeed
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	if err := xml.Unmarshal(data, &fetchedFeed); err != nil {
		return &RSSFeed{}, err
	}

	for i := range fetchedFeed.Channel.Item {
		fetchedFeed.Channel.Item[i].Title = html.UnescapeString(fetchedFeed.Channel.Item[i].Title)
		fetchedFeed.Channel.Item[i].Description = html.UnescapeString(fetchedFeed.Channel.Item[i].Description)
	}

	return &fetchedFeed, nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("The 'agg' needs a time between requests interval, e.g.: 5s")
	}

	interval := cmd.args[0]

	timeBetweenReqs, err := time.ParseDuration(interval)
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %.0f seconds\n\n", timeBetweenReqs.Seconds())

	ticker := time.NewTicker(timeBetweenReqs)

	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			fmt.Printf("\n%s\n", err)
		}
	}
}
