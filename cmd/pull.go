package cmd

import (
	"encoding/json"
	"fmt"
	"godo-cli/model"
	"godo-cli/pkg/common"
	"log"
	"os"

	"github.com/google/uuid"
)

func PullTasks() {
	tasks, err := model.Pull()
	if err != nil {
		log.Fatalf("Error pulling tasks: %v", err)
	}

	currentTasks, err := model.LoadTasks()
	if err != nil {
		log.Fatalf("Error loading tasks: %v", err)
	}

	existingIDs := make(map[uuid.UUID]bool)
	for _, task := range tasks {
		existingIDs[task.ID] = true
	}

	for _, currentTask := range currentTasks {
		if _, exists := existingIDs[currentTask.ID]; !exists {
			tasks = append(tasks, currentTask)
		}
	}

	formattedJSON, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		log.Fatalf("Error marshaling data to JSON: %v", err)
	}

	err = os.WriteFile(common.TasksFile, formattedJSON, 0644)
	if err != nil {
		log.Fatalf("Error writing data to file %s: %v", common.TasksFile, err)
	}

	fmt.Printf("Successfully pulled tasks to **%s**\n", common.TasksFile)
}
