package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

var students []Student

func AddStudent(name string, age int) {
	students = append(students, Student{Name: name, Age: age})
	fmt.Println("Student added successfully!")
}

func ListStudents() {
	if len(students) == 0 {
		fmt.Println("No students available.")
		return
	}
	fmt.Println("List of Students:")
	for i, student := range students {
		fmt.Printf("%d. Name: %s, Age: %d\n", i+1, student.Name, student.Age)
	}
}

func UpdateStudent(index int, newName string, newAge int) {
	if index < 0 || index >= len(students) {
		fmt.Println("Invalid index!")
		return
	}
	students[index].Name = newName
	students[index].Age = newAge
	fmt.Println("Student updated successfully!")
}

func DeleteStudent(index int) {
	if index < 0 || index >= len(students) {
		fmt.Println("Invalid index!")
		return
	}
	students = append(students[:index], students[index+1:]...)
	fmt.Println("Student deleted successfully!")
}

func main() {
	for {
		fmt.Println("\nTo-Do List - Student Management")
		fmt.Println("1. Add Student")
		fmt.Println("2. List Students")
		fmt.Println("3. Update Student")
		fmt.Println("4. Delete Student")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name string
			var age int
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter age: ")
			fmt.Scan(&age)
			AddStudent(name, age)
		case 2:
			ListStudents()
		case 3:
			var index int
			var newName string
			var newAge int
			ListStudents()
			fmt.Print("Enter student index to update: ")
			fmt.Scan(&index)
			index-- // Convert to zero-based index
			fmt.Print("Enter new name: ")
			fmt.Scan(&newName)
			fmt.Print("Enter new age: ")
			fmt.Scan(&newAge)
			UpdateStudent(index, newName, newAge)
		case 4:
			var index int
			ListStudents()
			fmt.Print("Enter student index to delete: ")
			fmt.Scan(&index)
			index-- // Convert to zero-based index
			DeleteStudent(index)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

