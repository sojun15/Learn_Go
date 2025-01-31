package main

import "fmt"

func main() {
	// anonymous function is that type of function which haven't name and it calls after ending parenthesis
	// it needn't any memory allocation
	func(x int, y int) {
		fmt.Print("multiplication:", x*y)
	}(15, 31)
}
