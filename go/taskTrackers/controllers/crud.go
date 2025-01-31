package controllers

import (
	"encoding/json"
	"os"
)

const filePath = "data.json"

func GetAll() []Task {
	var tasks []Task
	byteValue, _ := os.ReadFile(filePath)
	json.Unmarshal(byteValue, &tasks)
	return tasks
}

func getLatestId() int {
	var tasks []Task = GetAll()
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].Id + 1
}

func Create(task Task) {
	tasks := GetAll()
	tasks = append(tasks, task)
	file, _ := json.MarshalIndent(tasks, "", "    ")
	_ = os.WriteFile(filePath, file, 0644)
}

func Update(tasks []Task) {
	file, _ := json.MarshalIndent(tasks, "", "    ")
	_ = os.WriteFile(filePath, file, 0644)
}
