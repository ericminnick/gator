package main

import (
	"context"
	"fmt"
	//"time"	
	//"github.com/ericminnick/gator/internal/database"
	//"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("Usage: %s", cmd.Name)
	}

	feedURL := "https://www.wagslane.dev/index.xml"
	rssFeed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", rssFeed)
	return nil 
}


