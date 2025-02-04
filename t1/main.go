package main

import "fmt"

var (
	basic       float32
	allowance   float32
	temperature Temperature
	students    StudentGrade
)

const (
	_ DaysOfTheWeek = iota
	Sunday
	Monday
	Tuesday
	Wednessday
	Thursday
	Friday
	Saturday
)

func (day DaysOfTheWeek) String() string {
	switch day {
	case Sunday:
		return "SUNDAY"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednessday:
		return "Wednessday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	default:
		panic("invalid day")
	}
}

type Temperature float64
type Fahrenheit float64
type StudentGrade map[string][]float32
type DaysOfTheWeek uint

type Book struct {
	Title  string
	Author string
	Price  float64
}

func main() {
	// fmt.Println("Enter Basic Salary and allowance: ")
	// fmt.Scanf("%f %f", &basic, &allowance)
	// fmt.Printf("\n basic: %f\n allowance: %f\n take home: %f", basic, allowance, basicSalary(basic, allowance))

	// fmt.Println("\nEnter Temperature: ")
	// fmt.Scanf("%f", &temperature)
	// fmt.Printf("%f °C is equivalent to %f °F", temperature, celciusToFahrenheit(temperature))

	book1 := createBook("Women Of Owu", "Chinua Achebe", 2000)
	book2 := createBook("Things Fall Apart", "Chinua Achebe", 2500)
	book3 := createBook("Purple Hibiscus", "Helen Opara", 4000)

	fmt.Printf("%+v\n", book1)
	fmt.Printf("%+v\n", book2)
	fmt.Printf("%+v\n", book3)
	fmt.Println(">>>>>>>>>>>> deleted a book")
	fmt.Printf("%+v\n", book3)
	book3.updateBookTitle("Protocol Basics")
	book3.updateBookAuthor("devlongs")
	fmt.Println(">>>>>>>>>>>> updated a book")
	fmt.Printf("%+v\n", book3)
	fmt.Println(">>>>>>>>>>>> discounted price")
	fmt.Println(book3.calculateDiscount(4, 10))

	storeStudentGrade("mano", 1)
	storeStudentGrade("bash", 5)

	retrieveStudentScore("mano")
	retrieveStudentScore("mani")
	retrieveStudentScore("bash")

	tellTimeOfTheWeek(Saturday)
	tellTimeOfTheWeek(Monday)
	tellTimeOfTheWeek(Friday)
	tellTimeOfTheWeek(Sunday)
}

func basicSalary(basic float32, allowance float32) float32 {
	totalSalary := basic + allowance
	takeHome := totalSalary - (totalSalary * 0.075)

	return takeHome
}

func celciusToFahrenheit(temperature Temperature) Fahrenheit {
	t := temperature * (9 / 5)
	return Fahrenheit(t + 32)
}

func createBook(title string, author string, price float64) Book {
	book := Book{
		Title:  title,
		Author: author,
		Price:  price,
	}

	return book
}

func (book *Book) deleteBook() *Book {
	return nil
}

func (book *Book) updateBookTitle(title string) Book {
	book.Title = title

	return *book
}

func (book *Book) updateBookAuthor(author string) Book {
	book.Author = author

	return *book
}

func (book *Book) updateBookPrice(price float64) Book {
	book.Price = price

	return *book
}

func (book *Book) calculateDiscount(qty float64, discount float64) float64 {
	totalPrice := qty * book.Price
	discountedPrice := totalPrice * (discount / 100)

	return totalPrice - discountedPrice
}

func storeStudentGrade(student string, score float32) {
	if students == nil {
		students = make(StudentGrade)
	}

	students[student] = append(students[student], score)
}

func retrieveStudentScore(student string) {
	if students == nil {
		fmt.Println("no student records yet")
	}

	_, ok := students[student]

	if !ok {
		fmt.Println("student doesnt exists")
	}

	var total float32
	for _, score := range students[student] {
		total += score
	}

	fmt.Printf("Student Score: %v\nAgrregated Score: %f\n", students[student], total)
}

func tellTimeOfTheWeek(day DaysOfTheWeek) {
	if day > 1 && day < 7 {
		fmt.Printf("\n%s is a weekday", day)
	} else {
		fmt.Printf("\n%s is a weekend", day)
	}
}
