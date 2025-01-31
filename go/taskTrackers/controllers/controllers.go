package controllers

import (
	"fmt"
	"os"
	"time"
)

// Task status constants
const (
	StatusToDo       = "To Do"
	StatusInProgress = "In Progress"
	StatusDone       = "Done"
)

// Date format for task timestamps
const dateFormat = "2006-01-02 15:04:05"

// printTask prints a single task in a formatted way
func printTask(task Task) {
	fmt.Printf("%d. %s - %s - Created: %s",
		task.Id,
		task.Description,
		task.Status,
		task.CreatedAt,
	)
	if task.UpdatedAt != "" {
		fmt.Printf(" - Updated: %s", task.UpdatedAt)
	}
	fmt.Println()
}

// findTask finds a task by ID and returns its index
func findTask(id int) int {
	tasks := GetAll()
	for i, task := range tasks {
		if task.Id == id {
			return i - 1
		}
	}
	fmt.Println("No task found with ID:", id)
	os.Exit(1)
	return -1
}

// AddTask creates a new task
func AddTask(description string) {
	task := Task{
		Id:          getLatestId(),
		Description: description,
		Status:      StatusToDo,
		CreatedAt:   time.Now().Format(dateFormat),
	}
	Create(task)
	fmt.Println("Task added successfully")
}

// UpdateTask updates an existing task's description
func UpdateTask(id int, newDescription string) {
	tasks := GetAll()
	index := findTask(id)

	tasks[index].Description = newDescription
	tasks[index].UpdatedAt = time.Now().Format(dateFormat)

	Update(tasks)
	fmt.Println("Task updated successfully")
}

// DeleteTask removes a task
func DeleteTask(id int) {
	tasks := GetAll()
	index := findTask(id)

	// Remove task at index
	tasks = append(tasks[:index], tasks[index+1:]...)

	Update(tasks)
	fmt.Println("Task deleted successfully")
}

// MarkTaskInProgress marks a task as in progress
func MarkTaskInProgress(id int) {
	updateTaskStatus(id, StatusInProgress)
}

// MarkTaskDone marks a task as done
func MarkTaskDone(id int) {
	updateTaskStatus(id, StatusDone)
}

// updateTaskStatus updates the status of a task
func updateTaskStatus(id int, status string) {
	tasks := GetAll()
	index := findTask(id)

	tasks[index].Status = status
	tasks[index].UpdatedAt = time.Now().Format(dateFormat)

	Update(tasks)
	fmt.Printf("Task marked as %s\n", status)
}

// ListTasks prints all tasks
func ListTasks() {
	fmt.Println("All Tasks:")
	for _, task := range GetAll() {
		printTask(task)
	}
}

// ListTasksWithStatus prints tasks filtered by status
func ListTasksWithStatus(status string) {
	fmt.Printf("Tasks with status: %s\n", status)
	for _, task := range GetAll() {
		if task.Status == status {
			printTask(task)
		}
	}
}

// Convenience methods for listing tasks by status
func ListTasksInProgress() {
	ListTasksWithStatus(StatusInProgress)
}

func ListTasksDone() {
	ListTasksWithStatus(StatusDone)
}

func ListTaskToDo() {
	ListTasksWithStatus(StatusToDo)
}
