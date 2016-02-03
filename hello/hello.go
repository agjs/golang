package main

import "github.com/agjs/hello/greeting"
import "fmt"

func main() {
	var s = greeting.Salutation{}
	s.Name = "Mary"
	s.Greeting = "Batman"
	greeting.Greet(s, greeting.CreatePrintFunction("?"), true)
	fmt.Println(greeting.Calculate(7, 5))

}
