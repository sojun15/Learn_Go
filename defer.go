package main

import "fmt"

func defer_func(x int) {
	fmt.Println("First value of x: ", x)
	defer fmt.Println("Second value of x: ", x)
	x += 2
	fmt.Println("Third value of x: ", x)
	defer fmt.Println("Fourth value of x: ", x)
}

func main() {
	var x int = 15

	defer_func(x)
}
