package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// os.mkdir
func main() {
	userInput := os.Args[1:]

	if len(userInput) < 1 {
		help()
		return
	}

	cmd := userInput[0]
	var task string
	if len(userInput) > 1 {
		task = strings.Join(userInput[1:], " ")
	}

	switch cmd {
	case "add":
		addTask(task)
	case "list":
		listTasks()
	case "delete":
		deleteTask(task)
	default:
		help()
	}

}

func initTodoList() {
	// Check if the file exists
	if _, err := os.Stat("tasks.txt"); err != nil {
		// File not exists
		_, err := os.Create("tasks.txt")
		// Check if there is an error
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
}

func addTask(task string) {
	initTodoList()
	if task == "" {
		fmt.Println("Please provide a task")
		return
	}
	// Open the file
	file, err := os.OpenFile("tasks.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString(task + "\n")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	listTasks()
}

func listTasks() {
	// Check if the file exists
	initTodoList()
	// Open the file
	file, err := os.Open("tasks.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Read the file
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		fmt.Printf("%d. %s\n", i, scanner.Text())
		i++
	}
	// Close the file
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func deleteTask(task string) {
	// Open the file
	file, err := os.Open("tasks.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the file
	var lines []string
	scanner := bufio.NewScanner(file)
	taskFound := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == task {
			taskFound = true
			continue
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	if !taskFound {
		fmt.Println("Task not found")
		return
	}

	// Write the remaining lines back to the file
	file, err = os.Create("tasks.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	writer.Flush()

	fmt.Println("Task deleted:", task)
}

func help() {
	fmt.Println("Available commands:")
	fmt.Println("  add <task>  - Add a new task")
	fmt.Println("  list        - List all tasks")
}
