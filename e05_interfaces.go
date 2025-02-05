package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float32) string
}

type Card struct{}
type Cash struct{}
type Transfer struct{}

func (c Card) Pay(amount float32) string {
	return fmt.Sprintf("%f paid using card\n", amount)
}

func (c Cash) Pay(amount float32) string {
	return fmt.Sprintf("%f paid using cash\n", amount)
}

func (c Transfer) Pay(amount float32) string {
	return fmt.Sprintf("%f paid using transfer\n", amount)
}

func processPayment(p PaymentMethod, amount float32) {
	fmt.Println(p.Pay(amount))
}

func main() {
	cardPayment := Card{}
	cashPayment := Cash{}
	transferPayment := Transfer{}

	processPayment(cardPayment, 2000)
	processPayment(cashPayment, 4000)
	processPayment(transferPayment, 3000)

	console := Console{}
	file := File{}

	setLog(console, "first console log")
	setLog(file, "first file log")
	setLog(file, "second file log")

	getType("hello")
	getType(12)
	getType(12.34)
}

var logs []string

type Logger interface {
	SetLog(message string) ([]string, error)
}

type Console struct{}
type File struct{}

func (c Console) SetLog(message string) ([]string, error) {
	logs = append(logs, message)

	return logs, nil
}

func (f File) SetLog(message string) ([]string, error) {
	logs = append(logs, message)

	return logs, nil
}

func setLog(l Logger, m string) {
	fmt.Println(l.SetLog(m))
}

func getLogs() []string {
	return logs
}

func getType(v interface{}) {
	switch v := v.(type) {
	case string:
		fmt.Println(v, ": string")

		return
	case int:
		fmt.Println(v, ": int")

		return
	case uint:
		fmt.Println(v, ": uint")

		return
	case uint8:
		fmt.Println(v, ": uint8")

		return
	case uint16:
		fmt.Println(v, ": uint16")

		return
	case uint32:
		fmt.Println(v, ": uint32")

		return
	case uint64:
		fmt.Println(v, ": uint64")

		return
	case float32:
		fmt.Println(v, ": float32")

		return
	case float64:
		fmt.Println(v, ": float64")

		return
	case []string:
		fmt.Println(v, ": []string")

		return
	case []float32:
		fmt.Println(v, ": []float32")

		return
	case []float64:
		fmt.Println(v, ": []float64")

		return
	//todo: custom types
	// case map:
	// 	fmt.Println(v, ": map")

	// 	return
	default:
		fmt.Println(v, ": type could not identified")
		return
	}
}
