package conditionals

import (
	"fmt"
)

func CourseLengths() int {
	gsc := 275
	gi := 300

	fmt.Println("Getting Started With CLI will take ", gsc, " mins to complete")
	fmt.Println("Go Init will take ", gi, " mins to complete")

	return gsc
}
