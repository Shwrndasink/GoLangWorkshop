package main

import (
	"fmt"
	"os"

	"fem.com/go/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data"
	contents, err := fileutils.ReadTextFile(filePath + "/text.txt")
	if err == nil {
		fmt.Println(contents)

		newContent := fmt.Sprintf("Original: %v\nDouble Original: %v%v", contents, contents, contents)
		fileutils.WriteTextFile(filePath+"/output.txt", newContent)
	} else {
		fmt.Printf("Error Panic!! %v", err)
	}
}
