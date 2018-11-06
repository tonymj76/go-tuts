package main

import "fmt"

type student struct {
	id        int
	firstName string
	lastName  string
	country   string
	age       int
	grade     string
}

// student in f is a type of struct
func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, k := range s {
		if f(k) {
			r = append(r, k)
		}
	}
	return r
}

func main() {
	c1 := student{id: 1,
		firstName: "Mary",
		lastName:  "john",
		country:   "India",
		age:       34,
		grade:     "B"}

	c2 := student{id: 2,
		firstName: "Mater",
		lastName:  "peter",
		country:   "Nigeria",
		age:       20,
		grade:     "A"}

	c3 := student{id: 3,
		firstName: "Esther",
		lastName:  "okon",
		country:   "Ghana",
		age:       28,
		grade:     "B"}

	gradeSlice := []student{c1, c2, c3}
	// func call
	f := filter(gradeSlice, func(x student) bool {
		if x.grade == "B" {
			return true
		}
		return false
	})

	fmt.Println(f)
	c := filter(gradeSlice, func(t student) bool {
		if t.age > 19 && t.age < 30 {
			return true
		}
		return false
	})

	fmt.Println(c)
}
