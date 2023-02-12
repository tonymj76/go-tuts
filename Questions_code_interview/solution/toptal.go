package main

import (
	"fmt"
	"math"
	"strings"
)

func maxProfit(prices []int) int {
	maxProfit := math.Inf(-1)
	point := len(prices)
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < point {
			point = prices[i]
		}
		maxProfit = math.Max(maxProfit, float64(prices[i+1]-point))
	}
	return int(maxProfit)
}

func solution(message string, k int) string {
	if k > len(message) {
		return strings.TrimSpace(message)
	}
	if k+1 >= len(message) {
		return strings.TrimSpace(message)
	}

	kstring := message[:k+1]

	if string(message[k+1]) == " " || string(message[k]) == " " {
		return strings.TrimSpace(kstring)
	}

	str := strings.Fields(kstring)
	return strings.Join(str[:len(str)-1], " ")
}

func main() {
	toptalTestCase(solution)
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))
}

func toptalTestCase(fn func(m string, k int) string) {
	cases := []struct {
		t    string
		k    int
		want string
	}{
		{"Codility We test coders", 14, "Codility We"},
		{"Codility We test coders ", 140, "Codility We test coders"},
		{"Codility We test coders", 16, "Codility We test"},
		{"Codility We test coders", 17, "Codility We test"},
		{"Codility We test coders", 22, "Codility We test coders"},
	}

	for _, tc := range cases {
		got := fn(tc.t, tc.k)
		if got == tc.want {
			fmt.Println("PASS")
		} else {
			fmt.Printf("want %s got %s\n", tc.want, got)
		}
	}
}
