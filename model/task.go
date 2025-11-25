package model

import (
	"encoding/json"
	"fmt"
	"godo-cli/pkg/common"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   string    `json:"createdAt"`
	UpdatedAt   string    `json:"updatedAt"`
	DeletedAt   string    `json:"deletedAt"`
	UserID      int       `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	DueDate     string    `json:"dueDate"`
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

func Pull() ([]Task, error) {
	userID := 1
	userIDStr := strconv.Itoa(userID)
	var tasks []Task

	req, err := http.NewRequest("GET", "http://localhost:8080/api/tasks", nil)

	if err != nil {
		log.Fatalf("Error creating request\n\n Error: %v", err)
	}

	req.Header.Set("User-ID", userIDStr)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error executing request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Recieved non-OK HTTP status: %d", res.StatusCode)
		return []Task{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &tasks)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	fmt.Println("Response Status:", res.StatusCode)
	fmt.Println("Response Body:")
	fmt.Println(string(body))

	return tasks, nil
}
