package main

import "github.com/agjs/hello/greeting"

func main() {
	var s = greeting.Salutation{"bob", "hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("?"), false)
}
