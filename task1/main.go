package main

import (
	"fmt"
	"strings"
)

func getPermutations(tokens []string) []string {
	permutations := []string{}
	if len(tokens) == 1 {
		permutations = []string{strings.Join(tokens[:], "")}
		return permutations
	}
	for i := range tokens {
		elements := make([]string, len(tokens))
		copy(elements, tokens)

		for _, perm := range getPermutations(append(elements[0:i], elements[i+1:]...)) {
			item := append([]string{tokens[i]}, perm)
			permutations = append(permutations, strings.Join(item[:], ""))
		}
	}
	return permutations
}

func checkConcatenation(tokens []string, input string) bool {

	if input == "" {
		return true
	}
	for _, element := range tokens {
		if strings.Contains(element, input){
			return true
		}
	}
	return false
}


func main() {
	permutations := getPermutations([]string{  "ab", "bc", "c" })

	for _, input := range []string{"abc", "ab", "c", "", "acba"} {
		fmt.Printf("%s: %v\n", input, checkConcatenation(permutations, input))
	}
}