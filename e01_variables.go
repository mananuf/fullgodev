package main

import (
	"fmt"
)

// """
// NB: 
// 	- 1. In Go, strings are actually made up of sequences of bytes, not a sequence of rune.
//"""

var version = 1

// or

var (
	firstname       string = "mananaf"
	lastname        string = "bankat"
	age             uint8  = 20
	dateOfBirth     string = "10th October"
	maritalStatus   bool   = false
	mixedFillenings bool   = true
	libraryuse      string = "hello"
	secretMystery   rune   = 'a' // used to represent unicode character. Unicode is actually the collection of all possible characters
	byteVar         byte   = 0x12
)

const gender string = "male"

func main() {

	//format specifiers
	fmt.Printf("%v\n %#v\n", firstname, age)

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

	emoji := '😀'
	var freeze rune = '🥶'
	fmt.Printf("type of %c is %T", emoji, emoji)
	fmt.Printf("\ntype of %c is %T", freeze, freeze)

	byteSlice := []byte(dateOfBirth)
	runeSlice := []rune(dateOfBirth)

	fmt.Printf("\n%v\n%T", byteSlice, byteSlice)
	fmt.Printf("\n%v\n%T", runeSlice, byteSlice)
}
