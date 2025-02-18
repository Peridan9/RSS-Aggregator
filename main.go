package main

import (
	"fmt"
	"log"

	"github.com/peridan9/RSS-Aggregator/internal/config" // Importing the config package to handle configuration operations.
)

func main() {

	// Read the existing configuration from the JSON file.
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	// Set the current user to "Daniel" and save the updated config.
	err = cfg.SetUser("Daniel")
	if err != nil {
		log.Fatalf("error seting user: %v", err)
	}

	// Read the config again to verify that the user was successfully updated.
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg) // Print the updated configuration to confirm changes.
}
