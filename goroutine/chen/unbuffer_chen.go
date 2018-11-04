/* package main


type person struct {
	name		[]byte
	age,pay  int
	rd			newName
}

type newName interface {
	giveName(n []byte) (num int, err error)
}

fun */

package main

import "fmt"

type people struct {
	name 			string
	age  			int
	inter			giveName
}

type limitedTries struct {
	I  	giveName
	name  string
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
	myName()(s string, err error)
}

// Note this method and why i don't use := in s, err
func (n *limitedTries) myName() (s string, err error){
	s, err = n.I.myName()
	return
}

func (n *company) myName() {
	fmt.Println("company: ", n.name)
}

// i have my struct type and interface then a function that will implement my interface...
func giveNames(inter giveName) {
	m, _ := inter.(*limitedTries)
	fmt.Println(m.name)
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
	
	fmt.Println(c.job[2])
	giveNames(p1.inter)
}
