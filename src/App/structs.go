package main

import "fmt"

type person struct {
	name    string
	age     int
	hobbies []string
}

func main() {
	personObj := person{
		"John",
		24,
		[]string{"cricket", "reading", "painting"},
	}
	personObj.age = 26
	fmt.Println(personObj)
	fmt.Println(personObj.name)
	fmt.Println(personObj.age)
}
