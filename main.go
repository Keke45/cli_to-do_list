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
			} else {
				fmt.Printf(" %d. [ ] %s\n", task.ID, task.Title)
			}
		}
	}
}

func completeTask(id int) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for i := range tasks {
		if tasks[i].ID == id {
			if tasks[i].Completed {
				fmt.Printf("Task %d is already completed.\n", id)
				return
			}
			tasks[i].Completed = true
			fmt.Printf("Task %d marked as completed.\n", id)
			return
		}
	}
	fmt.Printf("Task %d not found.\n", id)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("======TO-DO LIST======")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		scanner.Scan()
		text := scanner.Text()
		choice, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter task title: ")
			scanner.Scan()
			title := scanner.Text()
			addTask(title)
		case 2:
			viewTasks()
		case 3:
			viewTasks()
			if len(tasks) == 0 {
				continue
			}
			fmt.Print("Enter task ID to mark as completed: ")
			scanner.Scan()

			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			completeTask(id)
		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
