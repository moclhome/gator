package main

import (
	"bootdev/go/gator/internal/config"
	"fmt"
	"log"
	"os"
)

func main() {
	theConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	theState := config.State{
		Config: &theConfig,
	}
	theCommands := commands{
		commandMap: make(map[string]func(*config.State, command) error),
	}
	theCommands.register("login", handlerLogin)
	arguments := os.Args
	if len(arguments) < 2 {
		log.Fatal("There must be at least one argument used as a command for the gator program to work")
	}
	commandName := arguments[1]
	parameters := arguments[2:len(arguments)]
	theCommand := command{
		name:      commandName,
		arguments: parameters,
	}
	if err := theCommands.run(&theState, theCommand); err != nil {
		log.Fatalf("Error running command %s: %v", theCommand.name, err)
	}

	/*	err = theConfig.SetUser("Lauli")
		if err != nil {
			log.Fatal(err)
		}
		theConfig, err = config.Read()
		if err != nil {
			log.Fatal(err)
		}*/
	fmt.Printf("%v\n", theConfig)
}
