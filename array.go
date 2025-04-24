package main

import "fmt"

func main() {
	var name [5]string
	fmt.Println("input any 5 names:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&name[i])
	}

	for i := 0; i < 5; i++ {
		fmt.Print(name[i], " ")
	}
}
