package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/salehmotiwala/gator/internal/database"
	"github.com/salehmotiwala/gator/internal/feeds"
)

func handleAgg(state *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Command `agg` has only one argument: <time_between_reqs>.")
	}

	duration, err := time.ParseDuration(cmd.args[0])

	if err != nil {
		log.Fatalf("Cannot parse duration: %v", err)
	}

	fmt.Printf("Collecting feeds every %s\n", duration.String())

	ticker := time.NewTicker(duration)

	for ; ; <-ticker.C {
		scrapeFeeds(state)
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())

	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:     time.Now(),
		ID:            feed.ID,
	})

	if err != nil {
		return err
	}

	rssFeed, err := feeds.FetchFeed(feed.Url)

	for _, item := range rssFeed.Channel.Item {
		postDate, err := time.Parse(time.RFC1123, item.PubDate)

		if err != nil {
			log.Printf("Cannot parse time: %v", err)
		}

		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: postDate,
			FeedID:      feed.ID,
		}

		post, err := s.db.CreatePost(context.Background(), postParams)

		if err != nil {
			if !strings.Contains(err.Error(), "23505") {
				log.Println(err)
			}

			continue
		}

		fmt.Printf("Post %s created!\n", post.Title)
	}

	return nil
}
