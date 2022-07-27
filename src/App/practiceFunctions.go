package main

import "fmt"

type rectangle struct {
	height, width int
}

//create method for particular type
//this is not normal function
func (r rectangle) area() int {
	return r.width * r.height
}

func main() {
	//important----
	rect := rectangle{
		10, 20,
	}
	area := rect.area()
	fmt.Println("Area of a rectangle", area)

	//------------------------
	fun := func() {
		fmt.Println("This is anonymous function declaration")
	}
	//invoking anonymous function
	fun()
	//--------------------------------

	total := sum("sum is: ", 1, 2, 3, 4, 5)
	fmt.Println(total)
	//----------------------------

	msg := "This is me!"
	//we have passed the reference
	demo(&msg)
	fmt.Println("****", msg)
	//--------------------------------

	result, err := divide(3, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	//------------------------

}

func sum(str string, values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func demo(msg *string) {
	fmt.Println(*msg)
	*msg = "This is you!"
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
