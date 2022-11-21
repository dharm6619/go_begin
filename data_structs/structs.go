package datastructs

type courseMeta struct {
	author string
	level  string
	rating float64
}

func GetCourseMeta(courseName string) courseMeta {

	author1 := courseMeta{
		author: "Dharmendra Mishra",
		level:  "Beginner",
		rating: 2,
	}
	author2 := courseMeta{
		author: "mdharmendra",
		level:  "Intermediate",
		rating: 4.5,
	}

	courseDetailsMeta := make(map[string]courseMeta)
	courseDetailsMeta["GoLang Init"] = author1
	courseDetailsMeta["Getting Started With CLI"] = author1
	courseDetailsMeta["Data Structures and Algorithms"] = author2
	courseDetailsMeta["Cracking the Coding Interview: Interview Prep"] = author2

	return courseDetailsMeta[courseName]
}
