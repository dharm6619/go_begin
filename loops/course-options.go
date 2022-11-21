package loops

import (
	"fmt"
)

func CourseOptions() string {
	courseOptions := []string{
		"GoLang Init",
		"Getting Started With CLI",
		"Data Structures and Algorithms",
		"Cracking the Coding Interview: Interview Prep"}

	for index, data := range courseOptions {
		fmt.Println(index+1, ". ", data)
	}
	return "Thats All"
}
