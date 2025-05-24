package data

import (
	"encoding/json"
	"os"

	"project-app-todo-list-cli/model"
)

var FileName = "tasks.json"

func LoadTask() ([]model.Task, error) {
	data, err := os.ReadFile(FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Task{}, nil
		}
		return nil, err
	}
	var tasks []model.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func SaveTasks(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(FileName, data, 0644)
}
