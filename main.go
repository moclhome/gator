package main

import (
	"bootdev/go/gator/internal/config"
	"bootdev/go/gator/internal/database"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	db       *database.Queries
	config   *config.Config
	commands commands
}

func main() {
	// fetch db URL and (if existing) userName from config file
	theConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	theState := state{
		config: &theConfig,
	}

	// open db connection
	db, err := sql.Open("postgres", theState.config.Db_url)
	dbQueries := database.New(db)
	theState.db = dbQueries

	theCommands := commands{
		commandMap:            make(map[string]func(*state, command) error),
		commandDescriptionMap: make(map[string]string),
	}

	registerCommands(&theCommands)

	theState.commands = theCommands

	arguments := os.Args
	if len(arguments) < 2 {
		log.Fatal("There must be at least one argument used as a command for the gator program to work")
	}
	commandName := arguments[1]
	parameters := arguments[2:]
	theCommand := command{
		name:      commandName,
		arguments: parameters,
	}
	if err := theCommands.run(&theState, theCommand); err != nil {
		log.Fatalf("Error running command %s: %v", theCommand.name, err)
	}

	//fmt.Printf("%v\n", theConfig)
}

func registerCommands(cmds *commands) {
	cmds.register("login", handlerLogin, "Usage: login <username>. The user has to be registerd before and is set as current user.")
	cmds.register("register", handlerRegister, "Usage: register <username>. A user with this name will be created and logged in.")
	cmds.register("reset", handlerReset, "Usage: reset. All users and feeds are deleted.")
	cmds.register("users", handlerUsers, "Usage: users. A list of all users is displayed. The current user is marked.")
	cmds.register("agg", handlerAgg, "Usage: agg. TODO")
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed), "Usage: addFeed <feedName> <url>. Saves a feed with this name and url and lets the current user follow it.")
	cmds.register("feeds", handlerFeeds, "Usage: feeds. A list of all feeds is displayed, together with the user who has added this feed.")
	cmds.register("follow", middlewareLoggedIn(handlerFollow), "Usage: follow <url>. Adds the current user to the followers of this feed.")
	cmds.register("following", middlewareLoggedIn(handlerFollowing), "Usage: following. Lists all feeds that are followed by the current user.")
	cmds.register("help", handlerHelp, "Usage: help. Lists all upported commands with their ddescriptions.")
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow), "Usage: unfollow <url>. Removes the current user from the followers of this feed.")
	cmds.register("browse", middlewareLoggedIn(handlerBrowse), "Usage: browsw [limit]. Lists the posts of the current user. Optionally restricted to \"limit\" posts.")
}
