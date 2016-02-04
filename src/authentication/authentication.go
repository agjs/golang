package authentication

import "../database"
import "fmt"

// MainWindow - main menu options when user starts the program
func MainWindow(option int) {
	switch option {
	case 1:
		CreateUser()

	}
}

// CreateUser function
func CreateUser() {
	var username, password string
	var user = database.User{}

	fmt.Println("Please choose a unique username")
	fmt.Scanln(&username)
	fmt.Println("Please enter a password")
	fmt.Scanln(&password)

	if len(username) > 1 && len(password) > 1 {
		user.Username = username
		user.Password = password
		fmt.Println("User " + username + " created!")
	} else {
		fmt.Println("Something went wrong..")

	}

}
