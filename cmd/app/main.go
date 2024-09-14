package main

import (
	"go-rest/pkg/utils"
)
func main() {
	fg := utils.NewFileGenerator("./")


	err := fg.GenerateFile("Hello World")
	if err != nil {
		println(err.Error())
	}
	
}