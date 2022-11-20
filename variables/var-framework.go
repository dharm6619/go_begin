package variables

import (
	"fmt"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CourseDetails() (name string, course string) {
	name = os.Getenv("USER")
	course = "go init"
	updateCourse(&course)
	return name, course
}

func updateCourse(course *string) string {
	courseName := cases.Title(language.English)
	*course = courseName.String("getting started with CLI")
	fmt.Println("Updated Course to ", *course)
	return *course
}
