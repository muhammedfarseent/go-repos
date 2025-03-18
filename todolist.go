package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the ToDoList application  ")

	scanner := bufio.NewScanner(os.Stdin)

	type todoTask struct {
		name string
		time string
	}

	var todoList []todoTask

	for {
		fmt.Println("Please enter EXIT to exit, press any key to continue")
		scanner.Scan()
		choice := scanner.Text()

		if choice == "EXIT" {
			break
		}

		fmt.Println("Please enter the name of the todo task:")
		scanner.Scan()
		name := scanner.Text()

		fmt.Println("Please enter the estimated time to complete the task:")
		scanner.Scan()
		time := scanner.Text()

		todo := todoTask{name: name, time: time}
		todoList = append(todoList, todo)

		fmt.Println(name, "and", time, todo, todoList)
	}

	for _, todo := range todoList {
		fmt.Println("The task name is", todo.name, "and the time for it is", todo.time)
	}
}