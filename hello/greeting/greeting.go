package greeting

import "fmt"

//dsdsds
type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	if prefix := GetPrefix(salutation.Name); isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}

}

func GetPrefix(name string) (prefix string) {
	switch name {
	case "Bob":
		prefix = "Mr "
	case "Joe":
		prefix = "Dr "
	case "Mary":
		prefix = "Mrs "
	default:
		prefix = "Dude "

	}
	return
}

func Calculate(number_1, number_2 int) (num1 int, num2 int) {
	num1 += number_1 + 1
	num2 += number_2 + 11
	return
}

func CreateMessage(name, greeting string) (message string, alternate string) {
	fmt.Println(len(greeting))
	message = greeting + " " + name
	alternate = "Hey " + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func PrintCustom(s string, custom string) {
	fmt.Println(s, custom)
}

func Looper() (numbers int) {
	numbers = 0
	for i := 0; i < 15; i++ {
		numbers += i
	}
	return numbers

}

func Deferred() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
