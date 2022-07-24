package main

import "fmt"

func main() {
	//make function is optional
	shoppingCart := make(map[string]int)
	shoppingCart = map[string]int{
		"keyboard": 100,
		"mouse":    20,
		"laptops":  30,
	}
	fmt.Println(shoppingCart, len(shoppingCart))
	fmt.Println(shoppingCart["keyboard"])

	//check if value exists--2ways
	//compiler ignores underscores
	isPresent, ok := shoppingCart["mobile"]
	_, ok1 := shoppingCart["headphones"]
	fmt.Println(isPresent, ok)
	fmt.Println(ok1)

	delete(shoppingCart, "mouse")
}
