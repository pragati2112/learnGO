package main

import (
	"fmt"
)

func main() {
	shoppingCart := map[string]int{
		"keyboard": 100,
		"mouse":    20,
		"laptops":  30,
	}

	//first check if it exists or not then apply condition
	if _, ok := shoppingCart["monitor"]; ok {
		fmt.Println("laptops are present")
	} else {
		fmt.Println("Sorry")
	}

	// '&&' and operator, '||' or operator

	switch 2 {
	case 1:
		fmt.Println("hi i am from case1")
	case 2:
		fmt.Println("hi i am from case2")
	default:
		fmt.Println("Now a default answer")
	}

	switch 4 {
	case 1, 3, 5, 7, 9:
		fmt.Println("hi i am from case1 ODD")
	case 2, 4, 6, 8:
		fmt.Println("hi i am from case2 EVEN")
	default:
		fmt.Println("Now a default answer")

	}

	//you can use break and fallthrough keywords in switch blocks

	switch i := 2 + 3; i {
	case 1, 3, 5, 7, 9:
		fmt.Println("hi i am from case1 ODD")
	case 2, 4, 6, 8:
		fmt.Println("hi i am from case2 EVEN")
	default:
		fmt.Println("Now a default answer")

	}

	//use interface to get the type of variable

	var a interface{} = 5.6

	switch a.(type) {
	case string:
		fmt.Println("I am string")
	case int:
		fmt.Println("I am int")
	case float64:
		fmt.Println("I am float")
	default:
		fmt.Println("I am default")
	}

}
