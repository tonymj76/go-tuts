package main

import (
	"sort"
)

type Node struct {
	val   string
	right *Node
	left  *Node
}

func main() {
	a := Node{"a", nil, nil}
	b := Node{"b", nil, nil}
	//c := Node{"c", nil, nil}
	//d := Node{"d", nil, nil}
	//e := Node{"e", nil, nil}
	//f := Node{"f", nil, nil}
	a.left = &b
}

func Consecutive(arr []int) int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)

	sort.Ints(newArr)

	smallInt := newArr[0]
	bigInt := newArr[len(newArr)-1]
	m := make(map[int]int)
	var count int

	for x := smallInt + 1; x < bigInt; x++ {
		m[x] = 0
	}

	for y := 0; y < len(newArr); y++ {
		if _, ok := m[newArr[y]]; ok {
			m[newArr[y]]++
		}
	}
	for _, v := range m {
		if v == 0 {
			count++
		}
	}
	// code goes here
	return count

}
