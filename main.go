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
	validCommands.register("delete", handlerDelete)
	validCommands.register("reset", handlerReset)
	validCommands.register("users", handlerUsers)
	validCommands.register("agg", handlerAgg)
	validCommands.register("addfeed", handlerAddFeed)
	validCommands.register("feeds", handlerFeeds)
	validCommands.register("follow", handlerFollow)
	validCommands.register("following", handlerFollowing)

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Error: Too few arguments, program exits")
		os.Exit(1)
	}

	var currentCommand command

	currentCommand.name = args[0]
	for _, arg := range args[1:] {
		currentCommand.args = append(currentCommand.args, arg)
	}

	fmt.Println()
	err = validCommands.run(&currentState, currentCommand)
	fmt.Println()
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		os.Exit(1)
	}
}
