package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ericminnick/gator/internal/database"
	"github.com/google/uuid"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("Usage: %s", cmd.Name)
	}
	
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't find any users: %w", err)
	}
	printUsers(users, s)
	return nil
	
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User switched successfully")
	return nil

}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: 		uuid.New(),
		CreatedAt: 	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		Name:		name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}
	
	err = s.cfg.SetUser(user.Name)
	if err != nil { 
		return fmt.Errorf("couldn's set current user: %w", err)		
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:		%v\n", user.ID)
	fmt.Printf(" * Name:	%v\n", user.Name)
}

func printUsers(users []string, s *state) {
	for _, user := range users {
		if s.cfg.CurrentUserName == user {
			fmt.Printf("* %v (current)\n", user)
		} else {
			fmt.Printf("* %v\n", user)
		}
	}
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("Usage: %s", cmd.Name)
	}
	
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset users table")
	}
	
	fmt.Println("Users deleted")
	return nil

}
