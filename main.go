package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/goczangabor24/gator/internal/config"
	"github.com/goczangabor24/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	dbURL := cfg.DbUrl

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
	}

	dbQueries := database.New(db)

	var currentState state
	currentState.configData = &cfg
	currentState.db = dbQueries

	var validCommands commands
	validCommands.registeredCommands = make(map[string]func(*state, command) error)

	validCommands.register("login", handlerLogin)
	validCommands.register("register", handlerRegister)

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Error: Too few arguments, program exits")
		os.Exit(1)
	}

	var currentCommand command
	currentCommand.name = args[0]
	currentCommand.args = append(currentCommand.args, args[1])

	err = validCommands.run(&currentState, currentCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
