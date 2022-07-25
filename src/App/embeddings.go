package main

import "fmt"

type Processor struct {
	processorName string
	cores         int
}

type MemoryCap struct {
	memoryCapacity int
	memoryType     string
}

//composition relationship-----"HAS A" relationship
//Go doesn't have inheritance
type computer struct {
	Processor
	MemoryCap
	price int
}

func main() {
	//you can assign the value by accessing all the data types of structs
	comp := computer{}
	comp.price = 50000
	comp.processorName = "intel"
	comp.cores = 6
	comp.memoryCapacity = 500
	comp.memoryType = "DDR4"

	fmt.Println(comp)

	//there is another way-- "Embedded struct"
	comp1 := computer{
		Processor{
			"intel i10",
			6,
		},
		MemoryCap{
			256,
			"DDR4",
		},
		40000,
	}

	fmt.Println(comp1)
}
