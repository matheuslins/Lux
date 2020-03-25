package main

import (
	"testing"
)


func Test1Concatenation(t *testing.T) {
	permutations := getPermutations([]string{"ab", "bc", "a"})
	result := checkConcatenation(permutations, "abc")
	if !result {
		t.Error("Result failed!")
	}
}

func Test2Concatenation(t *testing.T) {
	permutations := getPermutations([]string{"a", "ab", "bc"})
	result := checkConcatenation(permutations, "ab")
	if !result {
		t.Error("Result failed!")
	}
}


func Test3Concatenation(t *testing.T) {
	permutations := getPermutations([]string{"ab", "bc"})
	result := checkConcatenation(permutations, "abc")
	if result {
		t.Error("Result failed!")
	}
}

