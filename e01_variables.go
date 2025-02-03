package main

import (
	"fmt"
)

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

	// arrays
	// arrays in go are fixedSized by default
	var fixedSizeArray [5]uint8
	fmt.Println(fixedSizeArray)

	fixedSizeArray[0] = 1
	fixedSizeArray[1] = 2
	fixedSizeArray[2] = 3
	fixedSizeArray[3] = 4
	fixedSizeArray[4] = 5
	// fixedSizeArray[5] = 6 // out of bounds

	fmt.Println(fixedSizeArray)

	// slices
	// slices in go are used to create dynamicSized arrays
	fruits := []string{
		"Apple",
		"Banana",
		"Oranges",
		"papaya",
	}

	fmt.Println(fruits)

	slicedFruits := fruits[0:2]
	fmt.Println(slicedFruits)

	slicedFruits = append(slicedFruits, "mango", "strawberry")
	fmt.Println(slicedFruits)

	newFruits := []string{
		"pawpaw",
		"tangerine",
		"pineapple",
	}

	copy(slicedFruits, newFruits)
	fmt.Println(slicedFruits)
}