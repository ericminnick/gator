package main

import (
	"context"
	"fmt"
	"time"
	"log"
	"strings"
	"database/sql"
	"github.com/ericminnick/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <time between requests e.g. 1m for 1 minute>", cmd.Name)
	}
	
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	layout := "Mon, 2 Jan 2006 15:04:05 +0000"

	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("couldn't find next feed to fetch", err)
		return 
	}

	log.Println("Found a feed to fetch")

	s.db.MarkFeedFetched(context.Background(), feed.ID) 
	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("couldn't fetch feed %s: %w", feed.Name, err)
		return
	}
	
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)

		var pubTime sql.NullTime
		if item.PubDate != "" {
			pubTime.Time, err = time.Parse(layout, item.PubDate)
			if err != nil {
				log.Printf("couldn't parse publish date: %w", err)
			} else {
				pubTime.Valid = true
			}
		} else {
			pubTime.Valid = false
		}

		
		var nullDescription sql.NullString
		if item.Description != "" {
			nullDescription.String = item.Description
			nullDescription.Valid = true
		} else {
			nullDescription.Valid = false
		}
		
		post, err := s.db.CreatePosts(context.Background(), database.CreatePostsParams{
			ID: 			uuid.New(),
			CreatedAt:		time.Now().UTC(),
			UpdatedAt: 		time.Now().UTC(),
			Title:			item.Title,
			Url:			item.Link,
			Description:	nullDescription,
			PublishedAt:	pubTime,
			FeedID:			feed.ID,

		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique") {
				continue
			} else {
				log.Printf("couldn't insert post %s into database: %w", post.ID, err)
			}

		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
