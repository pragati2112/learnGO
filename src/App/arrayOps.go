package main

import (
	"fmt"
)

func main() {
	//creating an array
	nums := [5]int{1, 2, 3, 4, 5}
	// copy by value
	newNums := nums
	fmt.Println("before changing", nums, newNums)
	nums[1] = 50
	// copy by references
	//numsRef := &nums
	fmt.Println("after changing", nums, newNums)

	my_arr1 := [5]string{"Scala", "Perl", "Java", "Python", "Go"}

	// Copying the array into new variable
	// Here, the elements are passed by value
	my_arr2 := my_arr1

	fmt.Println("Array_1: ", my_arr1)
	fmt.Println("Array_2:", my_arr2)

	my_arr1[0] = "C++"

	// Here, when we copy an array
	// into another array by value
	// then the changes made in original
	// array do not reflect in the
	// copy of that array
	fmt.Println("\nArray_1: ", my_arr1)
	fmt.Println("Array_2: ", my_arr2)

	//slicing same as python---
	fmt.Println(nums[:])
	fmt.Println(nums[1:3])

	//2D array
	//it will add 2 more rows initialized with zeroes automatically--as we gave size 5 for row
	identityMatrix := [5][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(identityMatrix)

}
