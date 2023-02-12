package solution

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCase := []struct {
		want int
		cs   string
	}{
		{cs: "abcabcbb", want: 3},
		{cs: "bbbbb", want: 1},
		{cs: "pwwkew", want: 3},
		{cs: "", want: 0},
		{cs: " ", want: 1},
		{cs: "avrd", want: 4},
	}
	for _, tc := range testCase {
		t.Run("length of longest sub string", func(t *testing.T) {
			got := LengthOfLongestSubstring(tc.cs)
			if tc.want != got {
				t.Fatalf("want %v; got %v", tc.want, got)
			}
		})
	}
}

func TestCharacterReplacement(t *testing.T) {
	testCase := []struct {
		want, k int
		cs      string
	}{
		{want: 4, cs: "ABAB", k: 2},
		{want: 4, cs: "AABABBA", k: 1},
	}

	for _, tc := range testCase {
		t.Run("character replacement", func(t *testing.T) {
			got := characterReplacement(tc.cs, tc.k)
			if tc.want != got {
				t.Fatalf("want %v; got %v", tc.want, got)
			}
		})
	}
}
