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
		commandMap: make(map[string]func(*state, command) error),
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
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFollowing)
	//cmds.register("help", handlerHelp)
}
