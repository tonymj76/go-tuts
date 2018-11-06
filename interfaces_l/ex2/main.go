package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	// sort.Slice(studyGroup, func(i, j int)bool{return studyGroup[i] < studyGroup[j]})
	// Note that Sort fn takes in interface
	// And for interface to take in interface it must have len, swap and Less
	// which means that people(user type) is an interface of interface which i declared
	sort.Sort(people(studyGroup))
	fmt.Println(studyGroup)

	s := []string{"Zeno", "John", "Al", "Jenny"}
	t := []string{"Zeno", "John", "Al", "Jenny"}
	x := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
	fmt.Println(s)
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Slice(n, func(i, j int) bool { return n[i] > n[j] })
	fmt.Println(n)

	// sort.Sort(sort.Reverse(sort.StringSlice(t)))
	// or sort.Sort(StringSlice(t))
	// sort.Sort(sort.Reverse(sort.IntSlice(t)))
	// or sort.Sort(IntSlice(t))
	sort.StringSlice(t).Sort()
	// sort.Ints(n)
	sort.Strings(x)
	fmt.Println(t)
	fmt.Println(x)

}
