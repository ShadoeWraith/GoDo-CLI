package model

import (
	"encoding/json"
	"fmt"
	"godo-cli/pkg/common"
	"io/fs"
	"os"
)

type Task struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
	UserID      int    `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	DueDate     string `json:"dueDate"`
}

func LoadTasks() ([]Task, error) {
	data, err := os.ReadFile(common.TasksFile)

	if err != nil {
		if os.IsNotExist(err) || err == fs.ErrNotExist {
			fmt.Printf("File %s not found. Starting with an empty list.\n", common.TasksFile)
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("error parsing JSON data: %w", err)
	}

	return tasks, nil
}

func (t *Task) Save() error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	tasks = append(tasks, *t)

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(common.TasksFile, data, 0644)
}
