package main

import (
	"fmt"
)

func main() {
	name := "Amartya"
	module := 3.2
	ptr := &module

	fmt.Println("Name is ", name)
	fmt.Println("Module is", module, "Memory address of variable is", ptr, "the value is", *ptr)
}
