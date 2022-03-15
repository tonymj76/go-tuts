package main

import (
	"fmt"
)

func main() {
	//name := "hello test"
	fmt.Println(finMaxSum("Codility we test coders", 14))
}

func Solution(message string, K int) string {
	var currentText string
	for i, msg := range message {
		currentText += string(msg)
		if i >= K-1 {
			if string(message[len(currentText)]) == " " {
				currentText = currentText[:len(currentText)-1]
			}
		}
	}
	return currentText
}

func finMaxSum(m string, k int) string {
	var currentText string
	for i := 0; i < len(m); i++ {
		currentText += string(m[i])
		if i >= k-1 {
			currentText = currentText[:len(currentText)-1]
		}
	}
	return currentText
}
