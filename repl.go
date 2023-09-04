package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func StartRepl(cfg *config){
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

		availiableCommands := getCommands()

		command, ok := availiableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue 
		}

		err := command.callback(cfg)
		if err != nil{
			fmt.Println(err)
		}

	}
	
}

type cliCommand struct{
	name string
	description string
	callback func(*config) error
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
			description: "Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
			callback:    commandMapb,
		},
	}
}

func cleanInput(str string) []string{
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}