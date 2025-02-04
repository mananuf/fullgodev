package main

import "fmt"

type Details struct {
	firstname  string
	lastname   string
	age     uint
}

type KeyVal map[uint]string;

func main () {
	usernames := map[byte]string{
		0x12: "mano",
		0x13: "kc",
		0x14: "shaaibu",
	}

	fmt.Printf("%+v\n",usernames)
	fmt.Println(usernames[20], usernames[0x12])

	userDetails := map[byte]map[string]Details{
		0x12: {"mano" : Details{
			"mananaf",
			"bankat",
			20,
		}},
		0x13: {"kc" : Details{
			"Kingsley",
			"Kefas",
			27,
		}},
		0x14: {"shaaiibu" : Details{
			"Shaaibu",
			"Suleiman",
			28,
		}},
	}

	fmt.Printf("%+v\n%T\n", userDetails, userDetails)
	fmt.Println(userDetails[20])

	fmt.Println(getMapping(make(map[uint]string)))
}

func getMapping(k KeyVal) KeyVal {
	k[1] = "one"
	k[2] = "two"

	return k
}