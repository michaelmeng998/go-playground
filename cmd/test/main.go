package main

import (
	"fmt"
)

//Conventions
//lower case private
//Upper case public

//Create error struct
type myError struct{}

//using the Error function from error interface
func (myErr *myError) Error() string {
	return "Error Handler Hit"
}

type userDetails struct {
	userName string
}

//...
func Greet(user userDetails) error {
	user.userName = "Michael \n"
	if len(user.userName) < 5 {
		errorDetails(5)
	}
	println(user.userName)
	// fmt.Fprintf(writer, "Greetings, %s", name)
	return nil
}

//...
func Sum(x int, y int) int {
	return x + y
}

//Error Handler
func errorDetails(enum int) {
	myErr := &myError{} // implements error
	//default return
	fmt.Println(myErr)
}

func main() {
	var user userDetails

	// user.userName = "A\n"

	// if len(user.userName) < 5 {
	// 	errorDetails(5)
	// } else {
	Greet(user)

	// sum := Sum(5, 5)

	//print sum to console

	fmt.Printf("End of Run!\n")
}
