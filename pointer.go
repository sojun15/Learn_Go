package main

import "fmt"

func passing_value(values [4]string) {
	fmt.Println(values)
}

func passing_address(address *[4]string) {
	fmt.Println(address)
}

func main() {
	var id int
	id = 15
	address := &id
	*address = 17
	fmt.Println("Address of id variable:", &id)
	fmt.Println("value of id variable:", id)

	array := [4]string{"my", "name", "is", "sojun"}

	// In pass by value there will be copy only values into another array or variable or memory address 
	passing_value(array)
	// In pass by address there will go only one address of the array or variable there save memory space
	passing_address(&array)
}
