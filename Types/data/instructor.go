package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func NewInstructor(name string, lastName string) Instructor {
	return Instructor{FirstName: name, LastName: lastName}
}

func (i Instructor) PrettyPrint() string {
	return fmt.Sprintf("%v %v, Score is: (%d)", i.FirstName, i.LastName, i.Score)
}
