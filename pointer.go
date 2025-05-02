package main

import "fmt"

func passing_value(values [5]string) {
	fmt.Println(values)
}

func passing_address(address *[5]string) {
	fmt.Println(address)
	fmt.Println("id:", address[4])
}

//? create custom datatype of student
type student struct {
	name       string
	discipline string
	id         int64
}

//? here works on my orginal object address values, not copy my object
func passing_struct(object_student *student) {
	fmt.Println("total object : ", object_student)
	fmt.Println("id:", object_student.id)
	fmt.Println("discipline", object_student.discipline)
	fmt.Println("name:", object_student.name)
	object_student.name = "sojun chandra"
}

func main() {

	var id int
	id = 15
	address := &id
	*address = 17
	fmt.Println("Address of id variable:", &id)
	fmt.Println("value of id variable:", id)

	array := [5]string{"my", "name", "is", "sojun", "15"}

	//? In pass by value there will be copy only values into another array or variable or memory address
	passing_value(array)
	//? In pass by address there will go only one address of the array or variable there save memory space
	passing_address(&array)

	object := student{
		name:       "sojun",
		discipline: "cse",
		id:         15,
	}

	//? passing my struct object address
	passing_struct(&object)

	//? if i change any object variable values into passing_struct function then it will be change my original object because here i pass my object address
	fmt.Println("Updated object name : ", object.name)
}
