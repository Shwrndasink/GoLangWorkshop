package main

import (
	"fmt"

	"fem.com/go/types/data"
)

func main() {
	marc := data.Instructor{Id: 0, FirstName: "Marc", LastName: "White", Score: 99}
	var joe data.Instructor = data.NewInstructor("Joe", "Martinez")
	fmt.Println(joe.PrettyPrint())
	fmt.Println(marc.PrettyPrint())

	goCourse := data.Course{Id: 2, Name: "Go Fundamentals", Instructor: joe}

	fmt.Printf("%v", goCourse)

	swiftWorkshop := data.NewWorkshop("Switft with iOS", joe)
	fmt.Printf("%v", swiftWorkshop)

	var courses [2]data.SignUpToCourse
	courses[0] = goCourse
	courses[1] = swiftWorkshop

	for _, course := range courses {
		fmt.Println(course)
	}
}
