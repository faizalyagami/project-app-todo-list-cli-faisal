package service

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"project-app-todo-list-cli/data"
	"project-app-todo-list-cli/model"
	"project-app-todo-list-cli/utils"
)

func AddTask(title, priority string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return errors.New("judul tugas tidak boleh kosong")
	}

	if !utils.IsValidPriority(priority) {
		return errors.New("prioritas tidak valid, gunakan: low, medium, atau high")
	}

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

	if len(tasks) == 0 {
		fmt.Println("Tidak ada tugas")
		return nil
	}

	fmt.Println("========List Task========")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "No\tTask\tStatus\tPriority")

	for i, task := range tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", i+1, task.Title, task.Status, task.Priority)
	}
	return w.Flush()
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

func MarkTaskDone(index int) error {
	tasks, err := data.LoadTask()
	if err != nil {
		return err
	}

	if index < 1 || index > len(tasks) {
		return fmt.Errorf("index tidak valid")
	}

	tasks[index-1].Status = "completed"
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

func SearchTask(k string) error {
	tasks, err := data.LoadTask()
	if err != nil {
		return err
	}
	fmt.Println("========Search Result========")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "No\tTask\tStatus\tPriority\n")

	found := false
	for i, task := range tasks {
		if strings.Contains(strings.ToLower(task.Title), strings.ToLower(k)) {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", i+1, task.Title, task.Status, task.Priority)
			found = true
		}
	}

	if !found {
		fmt.Println("No match tasks.")
		return nil
	}
	return w.Flush()
}
