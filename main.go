package main

import (
	"bufio"
	"encoding/json"
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
var nextID = 1

func loadTasks() error {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No tasks file exists, start with an empty list
		}
		return err
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return err
	}
	for i := range tasks {
		if tasks[i].ID >= nextID {
			nextID = tasks[i].ID + 1
		}
	}
	return nil
}

func saveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func addTask(title string) {
	tasks = append(tasks, Task{ID: nextID, Title: title, Completed: false})
	nextID++
}

func deleteTask(id int) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Task %d deleted.\n", id)
			return
		}
	}
	fmt.Printf("Task %d not found.\n", id)
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

func getIntInput(scanner *bufio.Scanner, prompt string) (int, error) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		return 0, fmt.Errorf("failed to read input")
	}

	return strconv.Atoi(scanner.Text())
}

func main() {
	err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("======TO-DO LIST======")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		choice, err := getIntInput(scanner, "Enter your choice: ")
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
			err := saveTasks()
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			}

		case 2:
			viewTasks()
		case 3:
			if len(tasks) == 0 {
				fmt.Println("No tasks to complete.")
				continue
			}
			viewTasks()
			id, err := getIntInput(scanner, "Enter task ID to mark as completed: ")
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			completeTask(id)
			err = saveTasks()
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			}
		case 4:
			if len(tasks) == 0 {
				fmt.Println("No tasks to delete.")
				continue
			}
			viewTasks()
			id, err := getIntInput(scanner, "Enter task ID to delete: ")
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			deleteTask(id)
			err = saveTasks()
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			}

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
