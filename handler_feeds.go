package main

import (
	"context"
	"fmt"
	"time"	
	"github.com/ericminnick/gator/internal/database"
	"github.com/google/uuid"
)



func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("Usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't retrieve feeds %w:", err)
	}	

	printFeeds(s, feeds)

	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Usage: %s <name> <url>", cmd.Name)
	}

	feedName, feedURL := cmd.Args[0], cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt:	time.Now().UTC(),
		Name:		feedName,
		Url:		feedURL,
		UserID:		user.ID,		
	})
	if err != nil {
		return fmt.Errorf("couldn't not create feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: 		uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		UserID:		user.ID,
		FeedID:		feed.ID,

	})

	fmt.Println("Successfully added feed:")
	printFeed(feed)
	return nil

}


func printFeeds(s *state, feeds []database.Feed) {
	for i, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			fmt.Errorf("couldn't retrieve user %w", err)
		}
		fmt.Printf("Feed %v\n", i+1)
		fmt.Printf(" * Feed Name: 		%s\n", feed.Name)	
		fmt.Printf(" * Feed URL:  		%s\n", feed.Url)
		fmt.Printf(" * Feed UserName:	%s\n", user.Name)
	}
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * Feed ID: 		%s\n", feed.ID)
	fmt.Printf(" * Feed Created:	%v\n", feed.CreatedAt)
	fmt.Printf(" * Feed Updated:	%v\n", feed.UpdatedAt)
	fmt.Printf(" * Feed Name: 		%s\n", feed.Name)
	fmt.Printf(" * Feed URL:  		%s\n", feed.Url)
	fmt.Printf(" * Feed UserID		%s\n", feed.UserID)
}
