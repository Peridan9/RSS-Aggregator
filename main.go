package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/peridan9/RSS-Aggregator/internal/config" // Importing the config package to handle configuration operations.
	"github.com/peridan9/RSS-Aggregator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {

	// Read the existing configuration from the JSON file.
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Connect to the database.
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}

	// Defer the closing of the database connection.
	defer db.Close()

	// Create a new database queries object.
	dbQueries := database.New(db)

	// Create a new state object.
	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	// Create a new commands object.
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	// Register the commands.
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerGetFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollowing))
	cmds.register("following", middlewareLoggedIn(handlerFollowsPerUser))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	// Check if the user has provided a command.
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	// Get the command name and arguments.
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	// Run the command.
	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
