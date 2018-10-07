package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

//making a method

type AreaOfCircle struct {
	r float64
}

func (param *AreaOfCircle) areaOfCircle() float64 {
	return math.Round(math.Pi * math.Pow(param.r, 2))
}

func Areas() {
	circle1 := AreaOfCircle{r: 3.0}
	circle2 := AreaOfCircle{553.3}
	circle3 := AreaOfCircle{3.43}
	circle4 := AreaOfCircle{89.5}
	fmt.Printf("%v\n%v\n%v\n%v\n", circle1.areaOfCircle(), circle2.areaOfCircle(), circle3.areaOfCircle(), circle4.areaOfCircle())
}

type path []byte

func (p *path) truncatePath() {
	holder := bytes.LastIndex(*p, []byte("/"))
	if holder > 0 {
		*p = (*p)[:holder] // why put in (*p)
	}
}
func main() {
	pathName := path("c:/dkd/dkd/we/takeme")
	pathName.truncatePath()
	fmt.Printf("%s\n", pathName)
	fmt.Println(os.Getwd())
}
