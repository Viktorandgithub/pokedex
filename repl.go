package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func StartRepl(cfg *config, args ...string){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("pokedex >")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availiableCommands := getCommands()

		command, ok := availiableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue 
		}

		err := command.callback(cfg, args...)
		if err != nil{
			fmt.Println(err)
		}

	}
	
}

type cliCommand struct{
	name string
	description string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "The map command displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "It displays the previous 20 locations. It's a way to go back.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "List the pokemon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "Catches the pokemon",
			callback:    commandCatch,
		},
	}
}

func cleanInput(str string) []string{
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}