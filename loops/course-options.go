package loops

import (
	"fmt"

	"github.com/dharm6619/go_begin/loops"
)

func CourseOptions() string {
	courseOptions := []string{
		"GoLang Init",
		"Getting Started With CLI",
		"Data Structures and Algorithms",
		"Cracking the Coding Interview: Interview Prep"}

	for index, data := range courseOptions {
		fmt.Printf("%d. %s - %f\n", index, data, loops.GetCourseCost(data))
	}
	return "Thats All"
}
