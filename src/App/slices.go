package main

import (
	"fmt"
)

func main() {
	//creating a slice
	//size not defined
	nums := []int{1, 2, 3, 4, 5}
	newNums := nums
	fmt.Println("before changing", nums, newNums)
	nums[1] = 50

	//value of nums and newNums both changes
	//because slices works on references
	fmt.Println("after changing", nums, newNums)

	//capacity and len of nums
	fmt.Println(len(nums), cap(nums))

	//make a slice by using make function

	a := make([]int, 4, 10)
	fmt.Println(a, len(a), cap(a))

	b := append(nums, a...)
	fmt.Println(b)

	c := append(nums, 23)
	fmt.Println(c)
	fmt.Println(nums)

}
