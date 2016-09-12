// +build find_dupes

package main

import (
	"bufio";
	"fmt";
	"os";
)

func main () {
	p := fmt.Printf
	/**
		- A map holds a set of key/value pairs
		- provides constant-time operations to store, retrieve, or test for an item in the set.
		- key may be of any type whose values can compare d with ==
		- value can be any type
	**/
	// make creates a new empty map: keys are string values are int
	counts := make(map[string]int)
	// bufio package exposes a type called Scanner
	// that reads standard input and breaks it into lines or words
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// Not a problem if the map doesn’t yet contain that key:
		// expression counts[line] on the right-hand side evaluates
		// to the zero value for its type, which is 0 for int
		line := input.Text()
    	counts[line] = counts[line] + 1
    	// More compact:
    	// counts[input.Text()]++

	}

	for line, n := range counts {
		if n > 1 {
			p("%d\t%s\n", n, line)
		}
	}
}