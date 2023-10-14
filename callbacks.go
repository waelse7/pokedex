package main

import (
    "os"
    "fmt"
    "github.com/waelse7/pokedex/infra"
)

func callbackExit() error {
    os.Exit(0)
    return nil
}
func callbackHelp() error {
    commands := getCommands()
    for _, command := range commands {
        fmt.Printf("%s: %s\n", command.name, command.description)
    }
    return nil
}
func callbackMap() error {
    locations := infra.GetLocation(GetNextPage())

    SetNextPage(locations.Next)
    SetPrevPage(locations.Previous)

    for _, loc := range locations.Results{
        fmt.Printf("->%v\n", loc.Name)
    }
    return nil
}
func callbackMapB() error {
    locations := infra.GetLocation(GetPrevPage())

    SetNextPage(locations.Next)
    SetPrevPage(locations.Previous)

    for _, loc := range locations.Results{
        fmt.Printf("->%v\n", loc.Name)
    }
    return nil
}
    