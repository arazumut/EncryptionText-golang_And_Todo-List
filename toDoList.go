package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const fileName = "tasks.json"

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var tasks []Task

func main() {
	loadTasks()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nTo-Do List:")
		listTasks()
		fmt.Println("\nCommands: add, remove, done, quit")
		fmt.Print("Enter command: ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "add":
			addTask(reader)
		case "remove":
			removeTask(reader)
		case "done":
			markTaskAsDone(reader)
		case "quit":
			saveTasks()
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

func loadTasks() {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []Task{}
			return
		}
		fmt.Println("Error reading tasks file:", err)
		os.Exit(1)
	}
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println("Error parsing tasks file:", err)
		os.Exit(1)
	}
}

func saveTasks() {
	file, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling tasks:", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		fmt.Println("Error writing tasks file:", err)
		os.Exit(1)
	}
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks")
		return
	}
	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Name)
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Print("Enter task: ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)
	taskID := 1
	if len(tasks) > 0 {
		taskID = tasks[len(tasks)-1].ID + 1
	}
	newTask := Task{
		ID:   taskID,
		Name: taskName,
		Done: false,
	}
	tasks = append(tasks, newTask)
	saveTasks()
}

func removeTask(reader *bufio.Reader) {
	fmt.Print("Enter task ID to remove: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			return
		}
	}
	fmt.Println("Task not found")
}

func markTaskAsDone(reader *bufio.Reader) {
	fmt.Print("Enter task ID to mark as done: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			saveTasks()
			return
		}
	}
	fmt.Println("Task not found")
}
