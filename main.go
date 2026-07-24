package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

var tasks []Task

func addTask(title string) {
	tasks = append(tasks, Task{ID: len(tasks) + 1, Title: title, Completed: false})
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
	} else {
		for _, task := range tasks {
			if task.Completed {
				fmt.Printf(" %d. [X] %s\n", task.ID, task.Title)
			}
			fmt.Printf(" %d. [ ] %s\n", task.ID, task.Title)

		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("======TO-DO LIST======")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Exit")

		fmt.Print("Enter your choice: ")
		scanner.Scan()
		text := scanner.Text()
		choice, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if choice == 3 {
			fmt.Println("Exiting...")
			break
		}

		switch choice {
		case 1:
			var title string
			fmt.Print("Enter task title: ")
			scanner.Scan()
			title = scanner.Text()
			addTask(title)
		case 2:
			viewTasks()

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
