package main

import "fmt"

func commandHelp(cfg *config, args ...string) error{
	fmt.Println("Welcom to the Pokedex help menu!")
	fmt.Println("Here is a list of availiable commands:")
	availiableCommands := getCommands()

	for _, cmd := range availiableCommands{
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	
	fmt.Println("")
	return nil
}