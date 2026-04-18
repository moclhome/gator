package main

import (
	"bootdev/go/gator/internal/database"
	"fmt"
	"os/exec"
)

func handlerOpen(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: open <url>")
	}
	url := cmd.arguments[0]
	err := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", url).Start()
	if err != nil {
		return fmt.Errorf("failed to open URL: %w", err)
	}

	return nil
}
