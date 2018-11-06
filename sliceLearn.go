package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

//making a method

type areaOfCircle struct {
	r float64
}

func (param *areaOfCircle) areaOfCircle() float64 {
	return math.Round(math.Pi * math.Pow(param.r, 2))
}

func areas() {
	circle1 := areaOfCircle{r: 3.0}
	circle2 := areaOfCircle{553.3}
	circle3 := areaOfCircle{3.43}
	circle4 := areaOfCircle{89.5}
	fmt.Printf("%v\n%v\n%v\n%v\n", circle1.areaOfCircle(), circle2.areaOfCircle(), circle3.areaOfCircle(), circle4.areaOfCircle())
}

type path []byte

func (p *path) truncatePath() {
	holder := bytes.LastIndex(*p, []byte("/"))
	if holder > 0 {
		*p = (*p)[:holder] // why put in (*p)
		//becus to dereference and call a function to it
	}
}

func mainfortrunc() {
	pathName := path("c:/dkd/dkd/we/takeme")
	pathName.truncatePath()
	fmt.Printf("%s\n", pathName)
	fmt.Println(os.Getwd())
}

// learning user define type like interface and method set
type user struct {
	username string
	position string
}

type admin struct {
	username string
	position string
}

type notice interface {
	notify()
}

func (p *user) notify() {
	fmt.Printf("username is %v and position is %v\n", p.username, p.position)
}

func (p *admin) notify() {
	fmt.Printf("username is %v and position is %v\n", p.username, p.position)
}

func usernotification(n notice) {
	n.notify() // will expect a location/address
}

func main() {
	u := user{"tony", "manager"}
	a := admin{"Mr Cross", "MD"}

	usernotification(&u)
	usernotification(&a)

	dict := map[string]int{"one": 1}
	dict2 := make(map[int]string)
	dict2[2] = "two"
	fmt.Println(dict["one"], dict2[2])
}
