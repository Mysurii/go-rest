package main

import (
	"go-rest/pkg/utils"
)
func main() {
	fg := utils.NewFileGenerator("./")


	err := fg.GenerateFile("testing.go", "hello world!")
	if err != nil {
		println(err.Error())
	}
	
}