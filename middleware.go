package main

import (
	"context"

	"github.com/peridan9/RSS-Aggregator/internal/database"
)

// middlewareLoggedIn is a middleware that checks if the user is logged in
// and calls the handler if they are
func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
