package main

import "fmt"

type people struct {
	name string
	age  int
}

type worker struct {
	people
	salary   int
	position string
}

type company struct {
	worker
	name string
	job  []string
}

type giveName interface {
	myName()
}

func (n *people) myName() {
	fmt.Println(n.name)
}
func (n *company) myName() {
	fmt.Println("company: ", n.name)
}
func giveNames(n giveName) {
	n.myName()
}
func main() {
	p1 := people{
		name: "Jerry",
		age:  43,
	}

	w1 := worker{
		people:   p1,
		salary:   5500,
		position: "operator",
	}

	c := company{
		worker: w1,
		name:   "Build2day",
		job:    []string{"computer op", "dish washer", "something"},
	}
	c.people.myName()
	c.worker.myName()
	fmt.Println(c.job[2])
	giveNames(&c)
}
