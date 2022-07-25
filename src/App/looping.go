package main

import "fmt"

func main() {
	//looping with multiple variables
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	//you can do this way also
	i, j := 0, 0
	for i < 5 {
		fmt.Println(i, j)
		i, j = i+1, j+1
		if i == 3 && j == 3 {
			break
		}
	}
	fmt.Println("******", i, j)

	//nested loops

	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			//continue will skip the loop
			if a*b == 0 {
				continue
			}
			if a*b == 16 {
				break
			}
			fmt.Println(a * b)
		}
	}

	// Range method

	nums := []int{23, 45, 56, 67, 89, 0}
	for index, val := range nums {
		fmt.Println(index, val)
	}

	shoppingCart := map[string]int{
		"keyboard": 100,
		"mouse":    20,
		"laptops":  30,
	}

	for k, v := range shoppingCart {
		fmt.Println(k, v)
	}

}
