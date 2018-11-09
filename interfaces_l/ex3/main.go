package main

import (
	"fmt"
	"log"
)

type user struct {
	Name  string
	Email string
}

type Admin struct {
	// this is not inheritance but composition
	user
	Level string
}

// Notify takes in struct and return nill if successful
func (u *user) Notify() error {
	log.Printf("user: sending User Email To %s<%s>\n",
		u.Name, u.Email)
	return nil
}

type notifier interface {
	Notify() error
}
// The method set of the corresponding pointer type *T is the set of all methods with receiver *T or T
// The method set of any other type T consists of all methods with receiver type T.
func sendNotification(notifier notifier) error {
	return notifier.Notify()
}

func main() {
	p1 := user{"Tony", "tony@gmail.com"}
	a1 := &Admin{user:p1, Level:"advance"}
	sendNotification(&p1)
	sendNotification(a1)
	fmt.Println(a1.user.Notify())
}