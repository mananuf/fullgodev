package main

import "fmt"

type Student struct {
	firstname  string
	lastname   string
	cohort     uint
	track      string
	assessment []float32
}

func main() {
	student1 := Student{}
	student1.firstname = "mananaf"
	student1.lastname = "bankat"
	student1.cohort = 1
	student1.track = "web3"
	student1.assessment = append(student1.assessment, 5, 10, 5)

	fmt.Println(student1)
	//fmt.Printf("first student:", student1) // first student:%!(EXTRA main.Student={mananaf bankat 1 web3 [5 10 5]})%
	fmt.Printf("\n first student: %v", student1)
	fmt.Printf("\n name of student is %s %s a student of cohort %d %s", student1.firstname, student1.lastname, student1.cohort, student1.track)
	fmt.Printf("\n %+v", student1)

	fmt.Println("\n", student1.fullname())

	student2 := Student{}
	student2.firstname = "amigo"
	student2.lastname = "banderas"

	fmt.Println("\n", student2.fullname())
	fmt.Printf("\n Final score for %s is %d", student1.firstname, student1.finalScore())

}

// Methods
func (student Student) fullname() string {
	return student.firstname + " " + student.lastname
}

func (student Student) finalScore() int {
	var total int;

	for _, score := range student.assessment {
		total += int(score)
	}

	return total
}
