package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func (i Instructor) PrettyPrint() string {
	return fmt.Sprintf("%v %v, Score is: (%d)", i.FirstName, i.LastName, i.Score)
}
