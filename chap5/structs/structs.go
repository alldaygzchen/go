package main

import (
	"chap5/structs/user"
	"fmt"
)
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}


func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// testUser := user.User{}
	// fmt.Println("#####",testUser)
	
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	// fmt.Println(appUser)

	if err != nil {
		fmt.Println(err)
		return
	}
	// hidden
	// fmt.Println(appUser.firstName)
	admin := user.NewAdmin("test@example.com", "test123")
	
	admin.User.OutputUserDetails()
	admin.User.ClearUserName()
	admin.User.OutputUserDetails()
	fmt.Println("##############")
	admin = user.NewAdmin("test@example.com", "test123")
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()
	fmt.Println("##############")

	// ... do something awesome with that gathered data!

	// outputUserDetails(&appUser)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}


// func outputUserDetails(u *user) {
// 	// ...
// 	// normally we need to dereference inside
// 	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate)
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }




