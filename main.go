package main

import (
	"bufio"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"NIWT/backends"
	"NIWT/frontends"
)

func LoadConfig(filePath string) (map[string]string, error) {
	config := make(map[string]string)
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filepath.Join(usr.HomeDir, filePath[1:]))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore empty lines and comments
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		// Split key and value by '='
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Printf("Warning: Skipping malformed line: %s", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		config[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return config, nil
}

func main() {
	//Load config
	config, err := LoadConfig("~/.niwtrc")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Get city from command-line argument
	if len(os.Args) < 2 {
		log.Fatal("Please provide a city name")
	}
	city := os.Args[1]

	// Fetch and display weather data
	weather, err := backends.GetWeatherData(config["OWM_API_KEY"], city)
	if err != nil {
		log.Fatalf("Error fetching weather data: %v", err)
	}

	frontends.AsciiDraw(*weather)
}
