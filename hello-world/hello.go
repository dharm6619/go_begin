package main

import (
	"fmt"

	"github.com/dharm6619/go_begin/conditionals"
	datastructs "github.com/dharm6619/go_begin/data_structs"
	"github.com/dharm6619/go_begin/loops"
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
	fmt.Println("Calling a timer for 5 seconds")
	loops.SecondsTimer(5)
	fmt.Println("Following are the courses we offer")
	courseData := loops.CourseOptions()

	for index, data := range courseData {
		fmt.Printf("%v. %s - %v\n", index, data, datastructs.GetCourseCost(data))
		fmt.Println("Following are the details about the author ")
		fmt.Println(datastructs.GetCourseMeta(data))
	}
}
