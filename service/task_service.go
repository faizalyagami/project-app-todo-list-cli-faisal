package service

import (
	"fmt"

	"project-app-todo-list-cli/data"
	"project-app-todo-list-cli/model"
)

func AddTask(title, priority string) error {
	tasks, err := data.LoadTask()
	if err != nil {
		return err
	}
	newTask := model.Task{
		Title:    title,
		Status:   "new",
		Priority: priority,
	}
	tasks = append(tasks, newTask)
	return data.SaveTasks(tasks)
}

func ListTasks() error {
	tasks, err := data.LoadTask()
	if err != nil {
		return err
	}

	fmt.Println("========List Task========")
	fmt.Println("No | Task              | Status    | Priority")
	for i, task := range tasks {
		fmt.Printf("%-3d| %-18s| %-10s| %s\n", i+1, task.Title, task.Status, task.Priority)
	}
	return nil
}

func UpdateTaskStatus(index int, newStatus string) error {
	tasks, err := data.LoadTask()
	if err != nil {
		return err
	}

	if index < 1 || index > len(tasks) {
		return fmt.Errorf("index tidak valid")
	}
	tasks[index-1].Status = newStatus
	return data.SaveTasks(tasks)
}

func DeleteTask(index int) error {
	tasks, err := data.LoadTask()
	if err != nil {
		return err
	}
	if index < 1 || index > len(tasks) {
		return fmt.Errorf("index tidak valid")
	}

	tasks = append(tasks[:index-1], tasks[index:]...)
	return data.SaveTasks(tasks)
}
