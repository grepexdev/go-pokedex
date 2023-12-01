package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu.",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the application.",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call will display the next 20 locations.",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Similar to the map command but it dispays the previous 20 locations. Basically, it's a way to go back.",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore",
			description: "explore {location_area}",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch {pokemon_name}",
			callback:    callbackCatch,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
