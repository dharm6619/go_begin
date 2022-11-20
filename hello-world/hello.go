package main

import (
	"fmt"

	"github.com/dharm6619/go_begin/conditionals"
	"github.com/dharm6619/go_begin/variables"
)

func main() {
	fmt.Println("Hello World: This is GO Init")
	fmt.Println("Current LoggedIn User is: ", variables.GetLoggedInUser())
	fmt.Println("Constant Value is: ", variables.ListConstants())
	fmt.Println("Course Details are as Follows:")
	name, course := variables.CourseDetails()
	fmt.Println("User ", name, " has opt the course ", course)
	fmt.Println(course, " will take ", conditionals.CourseLengths(), " mins to complete")
	fmt.Println("Full Name for CLI is : ", conditionals.GetFullName("cli"))
}
