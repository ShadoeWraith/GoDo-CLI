package cmd

import (
	"fmt"
	"godo-cli/model"
	"godo-cli/pkg/common"
	"os"
)

func List() {
	tasks, err := model.LoadTasks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError%s: Failed to load tasks: %v\n", common.ColorError, common.ColorReset, err)
		os.Exit(1)
	}

	fmt.Println("--- List of Tasks ---")
	for i, task := range tasks {
		status := "[ ]"
		if task.Completed {
			status = "\033[32m[X]\033[0m"
		}
		fmt.Printf("%s Task %d: %s\n", status, i+1, task.Title)
		fmt.Printf("    Description: %s\n\n", task.Description)
	}
}
