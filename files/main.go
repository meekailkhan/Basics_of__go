package main

import (
	"fmt"
	"os"

	"files.com/begin/fileutils"
)

func main() {
	rootPath, _ := os.Getwd() // getting persent working directory

	c, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")

	if err == nil {
		fmt.Println(c)
		newContent := fmt.Sprintf("Original Content is %v\nand Copied Conetnet is %v%v", c, c, c)
		fileutils.WriteTextFile(rootPath+"/data/test2.txt", newContent)
	} else {
		fmt.Printf("Error Panic!! %v", err)
	}
}
