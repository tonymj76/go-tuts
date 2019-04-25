package main

import (
	"fmt"
	"strings"
)

func filterDoubleSlice(slice []string) []string {
	var newslice []string
	for i := 0; i < len(slice); i++ {
		if strings.Index(strings.Join(newslice, " "), slice[i]) == -1 {
			newslice = append(newslice, slice[i])
		}
	}
	return newslice
}

func main() {
	fmt.Println(filterDoubleSlice([]string{"name", "name", "set", "set"}))
}

func filterDouble(str string) string {
	var newStr string
	for i := 0; i < len(str); i++ {
		if strings.Index(newStr, string(str[i])) == -1 {
			newStr += string(str[i])
		}
	}
	return newStr
}
