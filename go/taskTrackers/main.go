package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"taskTrackers/controllers"
)

func exitWithError(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func getIdAndBodyTask(input string) (int, string) {
	inputs := strings.Split(input, " ")
	if len(inputs) < 2 {
		exitWithError("Invalid format: Requires 'index body'")
	}

	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		exitWithError("Invalid index: Must be a number")
	}

	body := strings.Join(inputs[1:], " ")
	if body == "" {
		exitWithError("Empty task body")
	}

	return id, body
}

func getIdTask(input string) int {
	id, err := strconv.Atoi(input)
	if err != nil {
		exitWithError("Invalid index: Must be a number")
	}
	return id
}

func parseArgs() (command, body string) {
	if len(os.Args) < 2 {
		exitWithError("Usage: taskTracker <command> [args]\n" +
			"Commands:\n" +
			"  add <task>\n" +
			"  update <index> <task>\n" +
			"  delete <index>\n" +
			"  mark-in-progress <index>\n" +
			"  mark-done <index>\n" +
			"  list [in-progress|done|not-done]")
	}

	command = os.Args[1]
	if len(os.Args) > 2 {
		body = strings.Join(os.Args[2:], " ")
	}
	return
}

func main() {
	command, body := parseArgs()

	switch command {
	case "add":
		if body == "" {
			exitWithError("Empty task")
		}
		controllers.AddTask(body)

	case "update":
		id, taskBody := getIdAndBodyTask(body)
		controllers.UpdateTask(id, taskBody)

	case "delete":
		id := getIdTask(body)
		controllers.DeleteTask(id)

	case "mark-in-progress":
		id := getIdTask(body)
		controllers.MarkTaskInProgress(id)

	case "mark-done":
		id := getIdTask(body)
		controllers.MarkTaskDone(id)

	case "list":
		switch body {
		case "in-progress":
			controllers.ListTasksInProgress()
		case "done":
			controllers.ListTasksDone()
		case "to-do":
			controllers.ListTaskToDo()
		default:
			controllers.ListTasks()
		}

	default:
		exitWithError("Invalid command")
	}
}
