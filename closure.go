package main

import "fmt"

func call(money int) func() {
	id := 105

	show := func() {
		fmt.Println("Closure function called with id:", id)
		id += money
		fmt.Println("Updated id in closure:", id)
		fmt.Println("anonymous show function executed")
	}

	//! another way to define the closure function
	// return func() {
	// 	fmt.Println("Closure function called with money:", money)
	// 	money += 10
	// 	fmt.Println("Updated money in closure:", money)
	// 	fmt.Println("anonymous show function executed")
	// }

	return show
}

func main() {
	save := call(15)
	save()
	save()

	save2 := call(17)
	save2()
	save2()
}
