package greeting

import "fmt"

// Salutation type.
type Salutation struct {
	Name     string
	Greeting string
}

// Vertex type.
type Vertex struct {
	X int
	Y int
}

// Printer type of function that accepts a string and returns nothing.
type Printer func(string)

// Greet function.
func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	if prefix := GetPrefix(salutation.Name); isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}

}

// GetPrefix function that returns a string.
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

// Calculate function that returns two integer values.
func Calculate(number1, number2 int) (num1 int, num2 int) {
	num1 += number1 + 1
	num2 += number2 + 11
	return
}

// CreateMessage function that returns two string values.
func CreateMessage(name, greeting string) (message string, alternate string) {
	fmt.Println(len(greeting))
	message = greeting + " " + name
	alternate = "Hey " + name
	return
}

// Print function that accepts a string.
func Print(s string) {
	fmt.Print(s)
}

// PrintLine function that accepts a string.
func PrintLine(s string) {
	fmt.Println(s)
}

// CreatePrintFunction function that accepts a string and returns a Printer.
func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

// PrintCustom function that accepts two strings.
func PrintCustom(s string, custom string) {
	fmt.Println(s, custom)
}

// Looper function that accepts 0 arguments and returns an integer.
func Looper() (numbers int) {
	numbers = 0
	for i := 0; i < 15; i++ {
		numbers += i
	}
	return numbers

}

// Deferred function.
func Deferred() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// Pointers function.
func Pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// PointerToStruct function.
func PointerToStruct() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func IfStatement(lim int) {
	if fmt.Println("line"); 5 < lim {
		fmt.Println("Lim is greater then 5")
	} else {
		fmt.Println("It's not greater")
	}
}
