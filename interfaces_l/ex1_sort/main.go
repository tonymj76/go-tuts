package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

// this is more as __str__(self) in python
func (p person) String() string {
	return fmt.Sprintf("%s: %d,", p.Name, p.Age)
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []person{
		{Name: "Jerry", Age: 54},
		{Name: "bobby", Age: 67},
		{Name: "joe", Age: 23},
	}

	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool { return people[i].Age > people[j].Age })
	fmt.Println(people)
}
