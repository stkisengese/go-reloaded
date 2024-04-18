package main

import (
	"testing"

	gg "go_reloaded/modifytext"
)

type testCase struct {
	name     string
	function func([]string) []string
	input    []string
	want     []string
}

func TestModifications(t *testing.T) {
	testCases := []testCase{
		{"Hexadecimal", gg.Hexadecimal, []string{"1E", "(hex)", "files", "were", "added"}, []string{"30", "files", "were", "added"}},
		{"Binary", gg.Binary, []string{"It", "has", "been", "10", "(bin)", "years"}, []string{"It", "has", "been", "2", "years"}},
		{"Capitalize", gg.Capitalize, []string{"Ready,", "set,", "go", "(cap,", "2)", "!"}, []string{"Ready,", "Set,", "Go", "!"}},
		{"UpperCase", gg.UpperCase, []string{"Ready,", "set,", "go", "(up)", "!"}, []string{"Ready,", "set,", "GO", "!"}},
		{"LowerCase", gg.LowerCase, []string{"Ready,", "SET,", "GO", "(low,", "2)", "!"}, []string{"Ready,", "set,", "go", "!"}},
		{"VowelArticle", gg.VowelArticle, []string{"There", "it", "was.", "A", "amazing", "rock!"}, []string{"There", "it", "was.", "An", "amazing", "rock!"}},
		{"FixPunctuation", gg.FixPunctuation, []string{"There", "it", "was", "...", "A", "amazing", "rock", "!?"}, []string{"There", "it", "was...", "A", "amazing", "rock!?"}},
		{"FixApostrophe", gg.FixApostrophe, []string{"There", "it", "was:", "'", "A", "amazing", "rock", "'"}, []string{"There", "it", "was:", "'A", "amazing", "rock'"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.function(tc.input)
			if !equal(got, tc.want) {
				t.Fatalf("got %v, wanted %v", got, tc.want)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
