package main

import (
	"fmt"
	"os"

	"fem.com/go/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	contents, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")
	if err == nil {
		fmt.Println(contents)
	} else {
		fmt.Printf("Error Panic!! %v", err)
	}
}
