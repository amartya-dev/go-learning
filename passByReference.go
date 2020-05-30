package main

import (
	"fmt"
)

func main() {

	name := "Amartya"
	course := "GO"

	fmt.Println("\nHi ", name, "you're in ", course)

	changeCourse(&course)

	fmt.Println("\nHi ", name, "you're in ", course)
}

func changeCourse(course *string) string {
	*course = "Course 2"
	fmt.Println("\n Course changed to", course)

	return *course
}
