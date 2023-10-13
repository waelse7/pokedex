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
  }
}

func cleanInput(str string) []string {
  lower := strings.ToLower(str)
  words := strings.Fields(lower)
  return words
}














