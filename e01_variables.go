package main

import "fmt"

var version = 1

// or

var (
	firstname string = "mananaf"
	lastname string = "bankat"
	age uint8 = 20
	dateOfBirth string = "10th October"
	maritalStatus bool = false
)

func main () {

	// var variableName type = value;
	var newVersion uint8 = 2

	fmt.Println(version, newVersion)
}