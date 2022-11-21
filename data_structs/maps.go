package datastructs

func GetCourseCost(course_name string) float32 {

	courseCost := make(map[string]float32)
	courseCost["GoLang Init"] = 100.0
	courseCost["Getting Started With CLI"] = 150.50
	courseCost["Data Structures and Algorithms"] = 250.75
	courseCost["Cracking the Coding Interview: Interview Prep"] = 325.60

	return courseCost[course_name]

}
