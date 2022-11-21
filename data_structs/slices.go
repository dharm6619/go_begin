package datastructs

import (
	"fmt"
)

func DeclareSliceInfo() {
	courses := make([]string, 5, 10)
	fmt.Printf("Length of the slice is %d and capacity is %d\n", len(courses), cap(courses))
}
