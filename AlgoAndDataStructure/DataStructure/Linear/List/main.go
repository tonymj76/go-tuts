package main

import (
	list2 "container/list"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println()
	var list list2.List
	list.PushBack(11)
	list.PushBack(23)
	list.PushBack(33)

	for listItem := list.Front(); listItem != nil; listItem = listItem.Next() {
		fmt.Println(listItem.Value.(int))
	}
	fmt.Println(yakFun(400, 14) + yakFun(800, 14) + yakFun(950, 14))
	fmt.Println(yakFun(400, 13))
	fmt.Println(yakFun(950, 13))
	fmt.Println()
	fmt.Println(ShaveYak(400, 14))

}

func yakFun(yakAge, days int) float32 {
	var total float32
	for i := 0; i < days; i++ {
		total += 50.0 - float32(yakAge+i)*0.03
	}
	return total
}

func ShaveYak(yakAge, days float64) float64 {
	var total float64
	var shavedDays float64
	if yakAge < 100 {
		return 0
	}
	for i := 0; i < int(days); i++ {
		ageInDay := yakAge + float64(i)
		if ageInDay < 1000 {
			shavedDays = 8 + ageInDay*0.01
		}
	}
	fmt.Println(shavedDays)
	return shavedYakCount(shavedDays, days-1, total)
}

func shavedYakCount(shavedDays, day, total float64) float64 {
	fmt.Println(day)
	fmt.Println(shavedDays)
	if day <= 0 || shavedDays < 1 {
		return total
	}
	return shavedYakCount(day, day-shavedDays, total+1)
}
