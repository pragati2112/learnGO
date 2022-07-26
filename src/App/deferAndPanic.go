package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("1")
	// defer keyword is used when we want to execute
	//certain operation in the last of all operations of the function
	defer fmt.Println("2")
	fmt.Println("3")

	//all defer statements stored in a program stack
	//and then execute in a manner of last in first out!!!
	//output will be ---- This is pragati
	defer fmt.Println("pragati")
	defer fmt.Println("is")
	defer fmt.Println("This")

	//Panic is a keyword used to throw an error
	fmt.Println("creating a file")
	f, err := os.Create("newFile.txt")
	if err != nil {
		panic(err)
	}
	defer closeFile(f)
	writeFile(f)
}

func closeFile(f *os.File) {
	fmt.Println("lets close this file now!")
	err := f.Close()
	if err != nil {
		panic(err)
	}
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	_, err := fmt.Fprintln(f, "hello there")
	if err != nil {
		panic(err)
	}
}
