package main

import (
	"bootdev/go/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	return listPosts(s, cmd, user, false)

}
