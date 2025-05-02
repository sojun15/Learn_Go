package main

import "fmt"

func main() {
	student := [5]string{"today", "i", "am", "learning", "slice"}

	//? pointer ref-1, length=3, capacity=4('i' to last element of original array)
	slice1 := student[1:4]
	fmt.Println(slice1)
	fmt.Println("length : ", len(slice1))
	fmt.Println("capacity : ", cap(slice1))

	//? pointer ref-2, length=1, capacity=3('am' to last element of original array)
	slice2 := slice1[1:2]
	fmt.Println(slice2)
	fmt.Println("length : ", len(slice2))
	fmt.Println("capacity : ", cap(slice2))

	slices := []int64{1, 15, 17, 31}
	fmt.Println("slice : ", slices, "length : ", len(slices), "capacity : ", cap(slices))
}
