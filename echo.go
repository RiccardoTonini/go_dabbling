// Prints command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	/**
	The := symbol is part of a short variable declaration,
	a statement that declares one or more variables and gives them
	appropriate types based on the initializer values
	/**
	i++ These are statements, not expressions as they are in most languages in
	the C family, so j=i++ is illegal, and they are postfix only,
	so --i is not legal either
	**/
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
/**
The for loop is the only loop statement in Go:
	for initialization; condition; post {
		// zero or more statements
	}

// a traditional "while" loop
	for condition {
		// ...
	}

// a traditional infinite loop
	for {
		// ...
	}

This loop is infinite, though loops of this form may be terminated in some
other way, like a break or return statement.
**/