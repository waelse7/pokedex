package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(){
  for {
    fmt.Print(">> ")
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    
    input := cleanInput(sc.Text())
    
    if len(input) == 0 {
      continue
    }
  
    commandName := input[0]
    
    commands := getCommands()
    
    command, ok := commands[commandName]
    if !ok {
      fmt.Println("Invalid command")
      continue
    }
    command.callback()
  
  }
}

type cliCommand struct {
  name string
  description string
  callback func() error
}


func getCommands() map[string]cliCommand {
  return map[string]cliCommand{
    "help": {
      name: "help",
      description: "Displays help menu",
      callback: callbackHelp,
    },
    "exit": {
      name: "exit",
      description: "exits pokedex",
      callback: callbackExit,
    },
    "map": {
      name: "map",
      description: "shows the next 20 locations",
      callback: callbackMap,
    },
    "mapb": {
     name: "mapb",
      description: "shows the previous 20 locations",
      callback: callbackMapB,
    },
  }
}

func cleanInput(str string) []string {
  lower := strings.ToLower(str)
  words := strings.Fields(lower)
  return words
}














