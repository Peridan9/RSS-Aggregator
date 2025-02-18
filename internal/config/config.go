package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Name of the configuration file that will be stored in the user's home directory.
const configFileName = ".gatorconfig.json"

// Config represents the structure of the JSON configuration file.
type Config struct {
	DBUrl       string `json:"db_url"`            // Database connection URL.
	CurrentUser string `json:"current_user_name"` // Currently logged-in user
}

// SetUser updates the CurrentUser field in the config and writes the updated configuration to the file.
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUser = username
	return write(*cfg) // Save the updated config to the JSON file.
}

// Read and loads the configuration from the JSON file located in the home directory.
func Read() (Config, error) {

	// Get the full path of the configuration file.
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// Open the configuration file.
	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err // Return error if the file cannot be opened.
	}

	defer file.Close() // Ensure the file is closed after reading.

	// Decode the JSON content into a Config struct.
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err // Return error if JSON decoding fails.
	}

	return cfg, nil // Successfully return the parsed config.
}

// getConfigFilePath constructs the full file path for the configuration file in the user's home directory.
func getConfigFilePath() (string, error) {

	// Get the user's home directory.
	homepath, err := os.UserHomeDir()
	if err != nil {
		return "", err // Return error if home directory retrieval fails.
	}

	// Construct the full path to the configuration file.
	fullPath := filepath.Join(homepath, configFileName)
	return fullPath, nil
}

// write writes the given config struct to the JSON file, overwriting any existing content.
func write(cfg Config) error {
	// Get the full path of the configuration file.
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err // Return error if file path retrieval fails.
	}

	// Create or overwrite the configuration file.
	file, err := os.Create(fullPath)
	if err != nil {
		return err // Return error if file creation fails.
	}
	defer file.Close() // Ensure the file is closed after writing.

	// Encode the config struct into JSON and write it to the file.
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err // Return error if encoding/writing fails.
	}

	return nil // Successfully written the config.
}
