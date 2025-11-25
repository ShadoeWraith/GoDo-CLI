package cmd

import (
	"fmt"
	"godo-cli/model"
	"time"

	"github.com/google/uuid"
)

func Create(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: godo create -title <title> -description <description>")
		return
	}

	var task model.Task

	task.ID = uuid.New()
	task.UserID = 1
	task.CreatedAt = time.Now().Format(time.RFC3339Nano)
	task.UpdatedAt = time.Now().Format(time.RFC3339Nano)

	for i, arg := range args {
		if i+1 >= len(args) {
			continue
		}

		if arg == "-title" || arg == "-t" {
			fmt.Println("\nTitle: ", args[i+1])
			task.Title = args[i+1]
		}
		if arg == "-description" || arg == "-desc" || arg == "-d" {
			fmt.Println("Description: ", args[i+1])
			task.Description = args[i+1]
		}
	}

	if task.Title == "" {
		fmt.Println("Error: Task title cannot be empty.")
		return
	}

	if err := task.Save(); err != nil {
		fmt.Printf("Error saving task: %v\n", err)
	} else {
		fmt.Println("Task created successfully.")
	}
}
