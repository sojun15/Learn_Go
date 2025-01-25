package main

import "fmt"

func summation(num1 int, num2 int) int {
	fmt.Println("Summation of two numbers is:", num1+num2)
	return num1 + num2
}

func init() {
	fmt.Println("Init function call auto-metically as like constructor before main function")
}

func main() {
	var num1 int = 10
	var num2 int = 20
	var sum int = summation(num1, num2)
	fmt.Println("Summation of two numbers is:", sum)
}
