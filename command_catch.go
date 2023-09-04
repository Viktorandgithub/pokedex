package main

import (
	"errors"
	"fmt"
	"math/rand"
)


func commandCatch(cfg *config, args ...string) error{
	if len(args) != 1{
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	catchPokemon, err := cfg.pokeapiClient.CatchPokemon(pokemonName)
	if err != nil {
		return err
	}
	const threshold = 50
	randNum := rand.Intn(catchPokemon.BaseExperience)

	if randNum > threshold {
		return fmt.Errorf("Failed to catch %s!\n", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = catchPokemon
	fmt.Printf("%s was caught!\n", pokemonName)
	
 	return nil
}