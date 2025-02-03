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

const gender string = "male"

func main () {

	// var variableName type = value;
	var newVersion uint8 = 2

	//or
	description := "this is the story of my go dev life. a full paid course by my mentor devlongs"

	fmt.Println(description)
	fmt.Println(version, newVersion, gender, lastname, firstname, age, maritalStatus, dateOfBirth)
}