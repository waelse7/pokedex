package main

import "fmt"

func callbackHelp() error {
  commands := getCommands()
  for _, command := range commands {
    fmt.Printf("%s: %s\n", command.name, command.description)
  }
  return nil
}
