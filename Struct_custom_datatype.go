package main

import "fmt"

type Person struct {
	name string
	id   int
}

func print_student_info(student Person) {
	fmt.Println("student id:", student.id)
	fmt.Println("student name:", student.name)
}

func main() {
	// instantiate or object creation 
	var student1 = Person{
		name: "sojun",
		id:   15,
	}

	print_student_info(student1)
	
	// instantiate or object creation
	var student2 = Person{
		name: "sourov",
		id:   17,
	}

	print_student_info(student2)
}
