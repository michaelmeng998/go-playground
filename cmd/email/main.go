package main

import (
	"fmt"
)

//CODE to send EMAIL and SMS notifications in Go
//Testing interfaces, DI, and mock unit tests

//Conventions...
//lower case private
//Upper case public

//User variables
type User struct {
	Name     string
	Email    string
	Notifier userNotifier //pass in interface in user struct, dynamically changed
}

//userNotifier interface for sending messages
type userNotifier interface {
	sendMessage(user *User, message string) error
}

//Now use a receiver for sendMessage function
type emailNotifier struct {
}

//sending email to user
func (email *emailNotifier) sendMessage(user *User, message string) error {
	_, err := fmt.Printf("Sending email to %s with content %s \n", user.Name, message)
	if err != nil {
		return err
	}
	return nil
}

//Send message through SMS

type smsNotifier struct{}

func (sms *smsNotifier) sendMessage(user *User, message string) error {
	_, err := fmt.Printf("Sending SMS to %s with content %s", user.Name, message)
	if err != nil {
		return err
	}
	return nil
}

//handler takes in specific  User with Notifier handle specified
func (user *User) notify(message string) error {
	return user.Notifier.sendMessage(user, message)
}

func main() {
	user1 := User{"Michael", "michaelmeng998@gmail.com", &emailNotifier{}}
	user2 := User{"Michael", "michaelmeng998@gmail.com", &smsNotifier{}}

	user1.notify("Body: This was sent using EMAIL system \n")
	user2.notify("Body: This was sent using SMS system \n")

}
