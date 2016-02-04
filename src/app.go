package main

import "./authentication"

// import "./database"

import "fmt"

func main() {
	var option int
	fmt.Println("Hi, please choose an option:")
	fmt.Println("1 - Register, 2 - Login")
	fmt.Scan(&option)
	authentication.MainWindow(option)

}
