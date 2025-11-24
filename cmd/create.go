package cmd

import (
	"fmt"
	"godo-cli/model"
)

func Create(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: godo create -title <title> -description <description>")
		return
	}

	var task model.Task

	for i, arg := range args {
		if arg == "-title" || arg == "-t" {
			fmt.Println("\nTitle: ", args[i+1])
			task.Title = args[i+1]
		}
		if arg == "-description" || arg == "-desc" || arg == "-d" {
			fmt.Println("Description: ", args[i+1])
			task.Description = args[i+1]
		}
	}

	task.Save()
}
