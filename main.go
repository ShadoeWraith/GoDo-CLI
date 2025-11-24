package main

import (
	"fmt"
	"os"

	"godo-cli/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("A simple Go CLI To-Do application.")
		fmt.Println("\nUsage:")
		fmt.Println("  godo -help")
		return
	}

	command := os.Args[1]

	switch command {
	case "list", "ls":
		cmd.List()
	case "create", "c":
		cmd.Create(os.Args)
	case "update", "u":
		cmd.Update(os.Args)
	case "delete", "d":
		cmd.Delete(os.Args)
	case "-help", "-h":
		cmd.Help()
	case "-version", "-v":
		cmd.Version()
	case "-setuser", "-su":
		if len(os.Args) < 3 {
			fmt.Println("Usage: godo -setuser <userID>")
			return
		}
		cmd.SetUser(os.Args[2])
	default:
		cmd.Default(command)
	}
}
