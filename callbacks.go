package main

import (
	"errors"
	"fmt"
	"os"
    "math/rand"
	"github.com/waelse7/pokedex/api"
)

func callbackExit(args ...string) error {
    os.Exit(0)
    return nil
}
func callbackHelp(args ...string) error {
    commands := getCommands()
    for _, command := range commands {
        fmt.Printf("%s: %s\n", command.name, command.description)
    }
    return nil
}
func callbackMap(args ...string) error {
    locations, err := api.GetLocationList(GetNextPage())
    if err != nil {
        return err
    }

    SetNextPage(locations.Next)
    SetPrevPage(locations.Previous)

    for _, loc := range locations.Results{
        fmt.Printf("->%v\n", loc.Name)
    }
    return nil
}
func callbackMapB(args ...string) error {
    locations, err := api.GetLocationList(GetPrevPage())
    if err != nil {
        return err
    }
    SetNextPage(locations.Next)
    SetPrevPage(locations.Previous)

    for _, loc := range locations.Results{
        fmt.Printf("->%v\n", loc.Name)
    }
    return nil
}

func callbackExplore(args ...string) error{
    if len(args) != 1 {
        return errors.New("error: wrong number of arguements") 
    }
    pokemons, err := api.GetLocationArea(args[0])
    if err != nil {
        return err
    }
    for _, pokemon := range pokemons.PokemonEncounters {
        fmt.Println(pokemon.Pokemon.Name)
    }
    return nil
}
func callbackCatch(args ...string) error{
    if len(args) != 1 {
        return errors.New("error: wrong number of arguements") 
    }
    pokemon, err := api.GetPokemon(args[0])
    if err != nil {
        return err
    }
    const threshhold = 50
    randint := rand.Intn(pokemon.BaseExperience)
    if randint < threshhold {
        return fmt.Errorf("failed to catch %s", args[0])
    }
    fmt.Printf("You cought %s.\n", args[0])
    return nil
}