package main

import "fmt"

func welcome_message() {
	fmt.Println("Welcome to Go programming language")
}

func input_name() string {
	var name string
	fmt.Print("Enter your name:")
	fmt.Scan(&name)
	return name
}

func input_id() int {
	var id int
	fmt.Print("enter your id:")
	fmt.Scan(&id)
	return id
}

func input_discipline() string {
	var discipline string
	fmt.Print("enter your discipline:")
	fmt.Scan(&discipline)
	return discipline
}

func input_year_term() (int, int) {
	var year int
	var term int
	fmt.Print("enter your year:")
	fmt.Scan(&year)
	fmt.Print("enter your term:")
	fmt.Scan(&term)
	return year,term
}

func display(name string, id int, discipline string, year int, term int) {
	fmt.Println("My name is:", name)
	fmt.Println("My id is:", id)
	fmt.Println("My discipline is:", discipline)
	fmt.Println("My year is:", year)
	fmt.Println("My term is:", term)
}

func main() {
	welcome_message()

	var name string = input_name()

	var id int = input_id()

	var discipline string = input_discipline()

	year, term := input_year_term()

	display(name, id, discipline, year, term)
}
