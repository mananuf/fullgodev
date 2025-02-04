package main

import "fmt"

type Track uint
type Weekday uint

const (
	web2 Track = 1
	web3 Track = 2
)

const (
	_ Weekday = iota
	Sunday
	Monday
	Tuesday
	Wednessday
	Thursday
	_ //increment using iota
	Friday
	Saturday
)

func getDay(day Weekday) Weekday {
	switch day {
	case Sunday:
		return Sunday
	case Monday:
		return Monday
	case Wednessday:
		return Wednessday
	case Thursday:
		return Thursday
	case Friday:
		return Friday
	default:
		panic("not a weekday")
	}
}

func getCohortTrack(track Track) Track {
	switch track {
	case web2:
		return web2
	case web3:
		return web3
	default:
		panic("track does not exist")
	}
}

func (track Track) String() string {
	switch track {
	case web2:
		return "WEB2"
	case web3:
		return "WEB3"
	default:
		panic("track does not exist")
	}
}

func (day Weekday) String() string {
	switch day {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Wednessday:
		return "Wednessday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	default:
		panic("not a weekday")
	}
}

func main() {
	fmt.Println("Track: ", getCohortTrack(web2))
	fmt.Println("Track: ", getCohortTrack(web3))
	fmt.Println("Track: ", getDay(Sunday))
	fmt.Println("Track: ", getDay(Friday))

}
