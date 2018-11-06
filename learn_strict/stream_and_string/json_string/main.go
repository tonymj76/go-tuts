package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name  string
	Age   int `json:"your age"` //change
	Job   string
	Email string `json:"-"` // forget about email
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

//just like what william kenedy did in expost and unexport
type user struct {
	Name  string
	Email string
}

// Admin to be extport admin
type Admin struct {
	user
	Position string
}

func main() {
	// initialized with type zeroth value
	var p1 person
	p1.Name = "Anthony"
	p1.Age = 20
	p1.Job = "software programmer"
	p1.Email = "anthony@gmail.com"

	value, err := json.Marshal(&p1)
	check(err)
	fmt.Println(string(value))
	// to the export function or module it will look like this
	holder := Admin{ // this is like this because we can't asign to user due to it is unexported type
		Position: "admin",
	}
	holder.Name = "tony"
	holder.Email = "tony@gmail.com"
	var t trying

	json.Unmarshal([]byte(fmt.Sprintln(string(value))), &t)
	check(err)
	fmt.Println(t.Name)

}

type trying struct {
	Name  string
	Age   int //change
	Job   string
	Email string // forget about email
}
