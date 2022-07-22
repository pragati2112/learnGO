package main

import (
	"fmt"
	"strconv"
)

// globally declared variable
var i int = 32

//	declaring global variable in camelCase -- it will be available only in this package
//  if you are declaring variable in PascalCase --- it will be available across all the packages

func main() {

	// variable declaration
	// this variable with the same name is shadowing the globally declared variable
	// println always print i = 12
	var i int = 12

	// variable assignment with the declaration
	j := 10

	fmt.Println(i)
	fmt.Println(j)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Println("Hello GoLand")

	//type conversion--
	//we have to explicitly tell the compiler to type cast the variable
	var x float32 = float32(i)
	fmt.Printf("%v, %T\n", x, x)

	// we will get the ASCII value of 12
	var foo string = string(i)
	fmt.Printf("%v, %T\n", foo, foo)

	// string conversion of 12
	bar := strconv.Itoa(i)
	fmt.Printf("%v, %T\n", bar, bar)

	// bool data type
	xyz := true
	fmt.Printf("%v, %T\n", xyz, xyz)

	// complex type
	var a complex64 = 1 + 2i
	fmt.Println(a)
	fmt.Printf("%v, %T\n", real(a), real(a))
	fmt.Printf("%v, %T\n", imag(a), imag(a))

	var b complex64 = complex(2, 4)
	fmt.Println(b)

	// Arithmetic operations
	fmt.Println(a + b)
	fmt.Println(a - b)

	// string operations-----

	s := "This is my house"
	fmt.Println(s[1])
	fmt.Println(string(s[1]))

	//replacing 'h' to 'u'
	//first convert the whole string to bytes array
	// in go string stored as a byte array
	s1 := []byte(s)
	fmt.Println(s1)

	s1[1] = 'u'
	fmt.Println(s1)
	s = string(s1)
	fmt.Println(s)

}
