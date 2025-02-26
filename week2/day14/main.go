package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tasks, err := LoadTasks("tasks.txt")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}

	for {
		fmt.Println("\nTo-Do List")
		fmt.Println("1. View Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			viewTasks(tasks)
		case 2:
			tasks = addTask(tasks)
		case 3:
			tasks = removeTask(tasks)
		case 4:
			SaveTasks("tasks.txt", tasks)
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func viewTasks(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("Your Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func addTask(tasks []string) []string {
	fmt.Print("Enter new task: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	task := scanner.Text()

	if task != "" {
		tasks = append(tasks, task)
		SaveTasks("tasks.txt", tasks)
		fmt.Println("Task added successfully!")
	} else {
		fmt.Println("Task cannot be empty.")
	}
	return tasks
}

func removeTask(tasks []string) []string {
	viewTasks(tasks)
	if len(tasks) == 0 {
		return tasks
	}

	fmt.Print("Enter task number to remove: ")
	var index int
	fmt.Scanln(&index)

	if index < 1 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return tasks
	}

	tasks = append(tasks[:index-1], tasks[index:]...)
	SaveTasks("tasks.txt", tasks)
	fmt.Println("Task removed successfully!")
	return tasks
}
