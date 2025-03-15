package main

import "fmt"

type Animal struct{
	name string
	category string 
}

func (animal Animal) print_animal_info(){
	fmt.Println("Animal name:",animal.name)
	fmt.Println("Animal catagory:",animal.category)
}

func (animal Animal) Color(color string){
	fmt.Println("Color is",color)
}

func main(){
	var animal1 = Animal{
		name: "parot",
		category: "bird",
	}

	animal1.print_animal_info()
	animal1.Color("green")

	var animal2 = Animal{
		name: "doyel",
		category: "bird",
	}

	animal2.print_animal_info()
	animal2.Color("gray")
}