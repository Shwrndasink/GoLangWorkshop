package main

import (
	"fmt"

	"fem.com/go/types/data"
)

func main() {
	marc := data.Instructor{Id: 0, FirstName: "Marc", LastName: "White", Score: 99}

	fmt.Println(marc.PrettyPrint())
}
