package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ericminnick/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 0 || len(cmd.Args) > 2 {
		return fmt.Errorf("Usage: %s <optional limit parameter>", cmd.Name)
	}
	var limit int32
	limit = 2	
	if len(cmd.Args) == 1 {
		temp, err := strconv.Atoi(cmd.Args[0]) 
		if err != nil {
			return fmt.Errorf("Usage: %s <optional limit paramter> (must be an integer)", cmd.Name)
		}
		limit = int32(temp)
	} 
	browsing, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:	limit,
	})
	if err != nil {
		return fmt.Errorf("couldn't browse posts for user %s: %w", user.Name, err)
	}

	printPosts(browsing)
	return nil
}

func printPosts(browsing []database.GetPostsForUserRow){
	for _, post := range browsing {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)		
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("===================================")
	}
}
