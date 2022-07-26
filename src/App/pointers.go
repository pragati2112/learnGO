package main

import "fmt"

type Foo struct {
	bar int
}

func main() {
	a := 23
	//create a pointer and pass value by reference
	var b *int = &a
	fmt.Println(a)

	//de-referencing the pointer
	fmt.Println(*b)

	//create a pointer of Foo type
	var foo *Foo
	fmt.Println(foo)

	//provide memory location
	foo = new(Foo)
	//this will print 0---because bar in foo initialized with default value
	fmt.Println(foo)

	//assign value
	foo.bar = 10
	//you can do this way also
	(*foo).bar = 20

	//de-ref pointer
	fmt.Println((*foo).bar)
	//you can write this way also
	//Go automatically handles the de-ref pointer for structs
	fmt.Println(foo.bar)

}
