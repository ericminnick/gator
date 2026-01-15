package main

import (
	"fmt"
	"time"
	"context"

	"github.com/ericminnick/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <URL>", cmd.Name)
	}

	URL := cmd.Args[0]
	
	feed, err := s.db.FeedByURL(context.Background(), URL)	
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: 		uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		UserID:		user.ID,
		FeedID:		feed.ID,

	})
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Printf("Successfully created follow:\n")
	fmt.Printf(" * Feed Name: %s\n", feedFollow.FeedName)
	fmt.Printf(" * User Name: %s\n", feedFollow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("Usage: %s" , cmd.Name)
	}

	userFollowing, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get following: %w", err)
	}
	
	fmt.Printf("Showing Feeds followed by %s:\n", user.Name)
	for i, feed := range userFollowing {
		fmt.Printf(" * Feed %v Name: %s\n", i+1,feed.Name)
	}
	

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <URL>", cmd.Name)
	}

	URL := cmd.Args[0]
	
	feed, err := s.db.FeedByURL(context.Background(), URL)	
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	err = s.db.Unfollow(context.Background(), database.UnfollowParams{user.ID, feed.ID})
	if err != nil {
		return fmt.Errorf("couldn't unfollow feed: %w", err)
	}
	
	fmt.Printf("Feed: %s unfollowed\n", feed.Name)
	return nil
}
