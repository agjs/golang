package main

// import "github.com/agjs/hello/greeting"
import "github.com/agjs/playground"
import "fmt"

func main() {
	// var s = greeting.Salutation{}
	// s.Name = "Mary"
	// s.Greeting = "Batman"
	// greeting.Greet(s, greeting.CreatePrintFunction("?"), true)
	// fmt.Println(greeting.Calculate(7, 5))
	// fmt.Println(greeting.Looper())
	//
	// defer fmt.Println("world")
	//
	// fmt.Println("hello")
	// // greeting.Deferred()
	// greeting.Pointers()
	// greeting.PointerToStruct()
	// greeting.IfStatement(7)

	fmt.Print("Welcome to World of Golang! What's your name ? ")
	user := playground.Player{}
	fmt.Scanln(&user.Nickname)
	fmt.Println(user.Nickname)

}
